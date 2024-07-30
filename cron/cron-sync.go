/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package cron

import (
	"clog-sync/app"
	"clog-sync/db/entity"
	"clog-sync/db/repo"
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/andypangaribuan/gmod/core/db"
	"github.com/andypangaribuan/gmod/gm"
	"github.com/jackc/pgx/v5"
)

func doSync(tableName string, optAction string, callback func()) {
	defer callback()

	var (
		ctx          = context.Background()
		_lastSync, _ = gm.Conv.Time.ToTime("1990-01-01", "yyyy-MM-dd")
		lastSync     = *_lastSync
		endQuery     = db.FetchOpt().EndQuery(fmt.Sprintf("ORDER BY created_at LIMIT %v", app.Env.FetchLimit))
		stm          = "stm:" + tableName
	)

	internalSyncLog, err := repo.InternalSyncLog.Fetch("table_name=?", tableName)
	if err != nil {
		log.Printf("[internal_sync_log] error when fetch: %+v\n", err)
		return
	}

	if internalSyncLog == nil {
		internalSyncLog = &entity.InternalSyncLog{
			TableName: tableName,
			LastSync:  lastSync,
		}

		err = repo.InternalSyncLog.Insert(internalSyncLog)
		if err != nil {
			log.Printf("[internal_sync_log] error when insert: %+v\n", err)
			return
		}
	}

	lastSync = internalSyncLog.LastSync

	switch {
	case tableName == "info_log":
		exec(app.DbDestInfo, ctx, tableName, stm, qInsertInfoLog, &lastSync, stmLoopInfoLog,
			func(lastSync *time.Time) ([]*entity.InfoLog, error) {
				return repo.SourceInfoLog.Fetches("created_at>?", lastSync, endQuery)
			})

	case tableName == "service_log":
		exec(app.DbDestService, ctx, tableName, stm, qInsertServiceLog, &lastSync, stmLoopServiceLog,
			func(lastSync *time.Time) ([]*entity.ServiceLog, error) {
				return repo.SourceServiceLog.Fetches("created_at>?", lastSync, endQuery)
			})

	case tableName == "dbq_log" && optAction == "":
		exec(app.DbDestDbq, ctx, tableName, stm, qInsertDbqLog, &lastSync, stmLoopDbqLog,
			func(lastSync *time.Time) ([]*entity.DbqLog, error) {
				return repo.SourceDbqLog.Fetches("created_at>?", lastSync, endQuery)
			})

	case tableName == "dbq_log":
		var (
			opt, _      = strconv.Atoi(optAction)
			secondRange = 10
			start       = opt * secondRange
			seconds     = make([]int, 0)
		)

		for i := 0; i < secondRange; i++ {
			seconds = append(seconds, start+i)
		}

		stm = fmt.Sprintf("%v-%v", stm, opt)
		exec(app.DbDestDbq, ctx, tableName, stm, qInsertDbqLog, &lastSync, stmLoopDbqLog,
			func(lastSync *time.Time) ([]*entity.DbqLog, error) {
				return repo.SourceDbqLog.Fetches("created_at>?", lastSync, endQuery)
			})
	}
}

func exec[T any](dbConn *pgx.Conn, ctx context.Context, tableName string, stm string, qry string, lastSync *time.Time, loopExec func([]*T, *time.Time, *pgx.Conn, context.Context, string) error, fetches func(*time.Time) ([]*T, error)) {
	var (
		isPrepared  = false
		startedTime time.Time
		oneSecond   = float64(1000)
		oneMinute   = float64(1000 * 60)
		oneHour     = float64(1000 * 60 * 60)
	)

	for {
		startedTime = gm.Util.Timenow()

		ls, err := fetches(lastSync)
		if err != nil {
			log.Printf("[%v] error when fetches: %+v\n", tableName, err)
			return
		}

		total := len(ls)
		if total == 0 {
			log.Printf("[%v] doesn't have new data\n", tableName)
			return
		}

		log.Printf("[%v] have %v new data\n", tableName, total)

		if !isPrepared {
			isPrepared = true
			_, err = dbConn.Prepare(ctx, stm, qry)
			if err != nil {
				log.Printf("[db-destination] error when prepare: %+v\n", err)
				return
			}
		}

		tx, err := dbConn.Begin(ctx)
		if err != nil {
			log.Printf("[db-destination] error when begin: %+v\n", err)
			return
		}

		err = loopExec(ls, lastSync, dbConn, ctx, stm)
		if err != nil {
			return
		}

		err = tx.Commit(ctx)
		if err != nil {
			log.Printf("[%v] error when commit: %+v\n", tableName, err)
			return
		}

		err = repo.InternalSyncLog.Update(db.Update().Set("last_sync=?", lastSync).Where("table_name=?", tableName).AutoUpdatedAt(false))
		if err != nil {
			log.Printf("[internal_sync_log] error when update: %+v\n", err)
			return
		}

		log.Printf("[%v] last sync: %v\n", tableName, gm.Conv.Time.ToStrFull(*lastSync))
		durationMs := float64(time.Since(startedTime).Milliseconds())

		switch {
		case durationMs >= oneHour:
			log.Printf("[%v] duration: %.2f h\n", tableName, durationMs/oneHour)
		case durationMs >= 3*oneMinute:
			log.Printf("[%v] duration: %.2f m\n", tableName, durationMs/oneMinute)
		case durationMs >= oneSecond:
			log.Printf("[%v] duration: %.2f s\n", tableName, durationMs/oneSecond)
		default:
			log.Printf("[%v] duration: %v ms\n", tableName, int64(durationMs))
		}
	}
}
