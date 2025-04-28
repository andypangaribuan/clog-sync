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

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/andypangaribuan/gmod/core/db"
	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/gm"
	"github.com/jackc/pgx/v5"
)

func doSync(tableName string, logType string, optAction string, callback func()) {
	defer callback()

	var (
		ctx          = context.Background()
		_lastSync, _ = gm.Conv.Time.ToTime("1990-01-01", "yyyy-MM-dd")
		lastSync     = *_lastSync
		endQuery     = db.FetchOpt().EndQuery(fmt.Sprintf("ORDER BY created_at LIMIT %v", app.Env.FetchLimit))
		stm          = fmt.Sprintf("stm:%v:%v%v", tableName, logType, fm.Ternary(optAction == "", "", ":"+optAction))
	)

	internalSyncLog, err := repo.InternalSyncLog.Fetch("table_name=?", fmt.Sprintf("%v:%v%v", tableName, logType, fm.Ternary(optAction == "", "", ":"+optAction)))
	if err != nil {
		log.Printf("[internal_sync_log] error when fetch: %+v\n", err)
		return
	}

	if internalSyncLog == nil {
		internalSyncLog = &entity.InternalSyncLog{
			TableName: fmt.Sprintf("%v:%v%v", tableName, logType, fm.Ternary(optAction == "", "", ":"+optAction)),
			LastSync:  lastSync,
		}

		err = repo.InternalSyncLog.Insert(internalSyncLog)
		if err != nil {
			log.Printf("[internal_sync_log] error when insert: %+v\n", err)
			return
		}
	}

	lastSync = internalSyncLog.LastSync

	// secondRange:
	// - 6 = 60 seconds / 6 secondRange = 10 concurrent connection
	// - 1 = 60 seconds / 1 secondRange = 60 concurrent connection
	secondRange := 1
	switch logType {
	case "p1":
		secondRange = 60
	case "p10":
		secondRange = 6
	case "p60":
		secondRange = 1
	}

	switch {
	case tableName == "info_log":
		if logType == "p1" {
			exec(logType, optAction, app.LsDbDestInfo[0], app.ChLsDbDestInfo[0], ctx, tableName, stm, qInsertInfoLog, &lastSync, stmLoopInfoLog,
				func(lastSync *time.Time) ([]*entity.InfoLog, error) {
					return repo.SourceInfoLog.Fetches("created_at>?", lastSync, endQuery)
				})
		}

		var (
			opt, _  = strconv.Atoi(optAction)
			start   = opt * secondRange
			seconds = make([]int, 0)
		)

		for i := range secondRange {
			seconds = append(seconds, start+i)
		}

		exec(logType, optAction, app.LsDbDestInfo[opt], app.ChLsDbDestInfo[opt], ctx, tableName, stm, qInsertInfoLog, &lastSync, stmLoopInfoLog,
			func(lastSync *time.Time) ([]*entity.InfoLog, error) {
				if len(seconds) == 1 {
					return repo.SourceInfoLog.Fetches("created_at>? AND FLOOR(EXTRACT(SECOND FROM created_at))::INTEGER=?", lastSync, seconds[0], endQuery)
				}

				return repo.SourceInfoLog.Fetches("created_at>? AND FLOOR(EXTRACT(SECOND FROM created_at))::INTEGER IN (?)", lastSync, seconds, endQuery)
			})

	case tableName == "service_log":
		if logType == "p1" {
			exec(logType, optAction, app.LsDbDestService[0], app.ChLsDbDestService[0], ctx, tableName, stm, qInsertServiceLog, &lastSync, stmLoopServiceLog,
				func(lastSync *time.Time) ([]*entity.ServiceLog, error) {
					return repo.SourceServiceLog.Fetches("created_at>?", lastSync, endQuery)
				})
		}

		var (
			opt, _  = strconv.Atoi(optAction)
			start   = opt * secondRange
			seconds = make([]int, 0)
		)

		for i := range secondRange {
			seconds = append(seconds, start+i)
		}

		exec(logType, optAction, app.LsDbDestService[opt], app.ChLsDbDestService[opt], ctx, tableName, stm, qInsertServiceLog, &lastSync, stmLoopServiceLog,
			func(lastSync *time.Time) ([]*entity.ServiceLog, error) {
				if len(seconds) == 1 {
					return repo.SourceServiceLog.Fetches("created_at>? AND FLOOR(EXTRACT(SECOND FROM created_at))::INTEGER=?", lastSync, seconds[0], endQuery)
				}

				return repo.SourceServiceLog.Fetches("created_at>? AND FLOOR(EXTRACT(SECOND FROM created_at))::INTEGER IN (?)", lastSync, seconds, endQuery)
			})

	case tableName == "dbq_log":
		if logType == "p1" {
			exec(logType, optAction, app.LsDbDestDbq[0], app.ChLsDbDestDbq[0], ctx, tableName, stm, qInsertDbqLog, &lastSync, stmLoopDbqLog,
				func(lastSync *time.Time) ([]*entity.DbqLog, error) {
					return repo.SourceDbqLog.Fetches("created_at>?", lastSync, endQuery)
				})
		}

		var (
			opt, _  = strconv.Atoi(optAction)
			start   = opt * secondRange
			seconds = make([]int, 0)
		)

		for i := range secondRange {
			seconds = append(seconds, start+i)
		}

		exec(logType, optAction, app.LsDbDestDbq[opt], app.ChLsDbDestDbq[opt], ctx, tableName, stm, qInsertDbqLog, &lastSync, stmLoopDbqLog,
			func(lastSync *time.Time) ([]*entity.DbqLog, error) {
				if len(seconds) == 1 {
					return repo.SourceDbqLog.Fetches("created_at>? AND FLOOR(EXTRACT(SECOND FROM created_at))::INTEGER=?", lastSync, seconds[0], endQuery)
				}

				return repo.SourceDbqLog.Fetches("created_at>? AND FLOOR(EXTRACT(SECOND FROM created_at))::INTEGER IN (?)", lastSync, seconds, endQuery)
			})
	}
}

func exec[T any](logType string, optAction string, dbConn *pgx.Conn, chDbConn driver.Conn, ctx context.Context, tableName string, stm string, qry string, lastSync *time.Time, loopExec func([]*T, *time.Time, *pgx.Conn, driver.Conn, context.Context, string) error, fetches func(*time.Time) ([]*T, error)) {
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

		if dbConn != nil {
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

			err = loopExec(ls, lastSync, dbConn, chDbConn, ctx, stm)
			if err != nil {
				return
			}

			err = tx.Commit(ctx)
			if err != nil {
				log.Printf("[%v] error when commit: %+v\n", tableName, err)
				return
			}
		}

		if chDbConn != nil {
			err = loopExec(ls, lastSync, dbConn, chDbConn, ctx, "")
			if err != nil {
				log.Printf("[db-destination] error when insert: %+v\n", err)
				return
			}
		}

		err = repo.InternalSyncLog.Update(db.Update().Set("last_sync=?", lastSync).Where("table_name=?", fmt.Sprintf("%v:%v%v", tableName, logType, fm.Ternary(optAction == "", "", ":"+optAction))).AutoUpdatedAt(false))
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
