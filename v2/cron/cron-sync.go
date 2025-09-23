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

	switch tableName {
	case "internal":
		if logType == "p1" {
			exec(logType, optAction, app.LsDbDestInternal[0], app.ChLsDbDestInternal[0], ctx, tableName, stm, qInsertInternal, &lastSync, stmLoopInternal,
				func(safe string, lastSync *time.Time) ([]*entity.Internal, error) {
					return repo.SourceInternal.Fetches(safe+"created_at>?", lastSync, endQuery)
				})
			return
		}

		var (
			opt, _  = strconv.Atoi(optAction)
			start   = opt * secondRange
			seconds = make([]int, 0)
		)

		for i := range secondRange {
			seconds = append(seconds, start+i)
		}

		exec(logType, optAction, app.LsDbDestInternal[opt], app.ChLsDbDestInternal[opt], ctx, tableName, stm, qInsertInternal, &lastSync, stmLoopInternal,
			func(safe string, lastSync *time.Time) ([]*entity.Internal, error) {
				if len(seconds) == 1 {
					return repo.SourceInternal.Fetches(safe+"created_at>? AND FLOOR(EXTRACT(SECOND FROM created_at))::INTEGER=?", lastSync, seconds[0], endQuery)
				}

				return repo.SourceInternal.Fetches(safe+"created_at>? AND FLOOR(EXTRACT(SECOND FROM created_at))::INTEGER IN (?)", lastSync, seconds, endQuery)
			})

	case "note_v1":
		if logType == "p1" {
			exec(logType, optAction, app.LsDbDestNoteV1[0], app.ChLsDbDestNoteV1[0], ctx, tableName, stm, qInsertNoteV1, &lastSync, stmLoopNoteV1,
				func(safe string, lastSync *time.Time) ([]*entity.NoteV1, error) {
					return repo.SourceNoteV1.Fetches(safe+"created_at>?", lastSync, endQuery)
				})
			return
		}

		var (
			opt, _  = strconv.Atoi(optAction)
			start   = opt * secondRange
			seconds = make([]int, 0)
		)

		for i := range secondRange {
			seconds = append(seconds, start+i)
		}

		exec(logType, optAction, app.LsDbDestNoteV1[opt], app.ChLsDbDestNoteV1[opt], ctx, tableName, stm, qInsertNoteV1, &lastSync, stmLoopNoteV1,
			func(safe string, lastSync *time.Time) ([]*entity.NoteV1, error) {
				if len(seconds) == 1 {
					return repo.SourceNoteV1.Fetches(safe+"created_at>? AND FLOOR(EXTRACT(SECOND FROM created_at))::INTEGER=?", lastSync, seconds[0], endQuery)
				}

				return repo.SourceNoteV1.Fetches(safe+"created_at>? AND FLOOR(EXTRACT(SECOND FROM created_at))::INTEGER IN (?)", lastSync, seconds, endQuery)
			})

	case "service_piece_v1":
		if logType == "p1" {
			exec(logType, optAction, app.LsDbDestServicePieceV1[0], app.ChLsDbDestServicePieceV1[0], ctx, tableName, stm, qInsertServicePieceV1, &lastSync, stmLoopServicePieceV1,
				func(safe string, lastSync *time.Time) ([]*entity.ServicePieceV1, error) {
					return repo.SourceServicePieceV1.Fetches(safe+"created_at>?", lastSync, endQuery)
				})
			return
		}

		var (
			opt, _  = strconv.Atoi(optAction)
			start   = opt * secondRange
			seconds = make([]int, 0)
		)

		for i := range secondRange {
			seconds = append(seconds, start+i)
		}

		exec(logType, optAction, app.LsDbDestServicePieceV1[opt], app.ChLsDbDestServicePieceV1[opt], ctx, tableName, stm, qInsertServicePieceV1, &lastSync, stmLoopServicePieceV1,
			func(safe string, lastSync *time.Time) ([]*entity.ServicePieceV1, error) {
				if len(seconds) == 1 {
					return repo.SourceServicePieceV1.Fetches(safe+"created_at>? AND FLOOR(EXTRACT(SECOND FROM created_at))::INTEGER=?", lastSync, seconds[0], endQuery)
				}

				return repo.SourceServicePieceV1.Fetches(safe+"created_at>? AND FLOOR(EXTRACT(SECOND FROM created_at))::INTEGER IN (?)", lastSync, seconds, endQuery)
			})

	case "service_v1":
		if logType == "p1" {
			exec(logType, optAction, app.LsDbDestServiceV1[0], app.ChLsDbDestServiceV1[0], ctx, tableName, stm, qInsertServiceV1, &lastSync, stmLoopServiceV1,
				func(safe string, lastSync *time.Time) ([]*entity.ServiceV1, error) {
					return repo.SourceServiceV1.Fetches(safe+"created_at>? AND created_at < NOW() - INTERVAL '?'", lastSync, app.Env.SafeFetch, endQuery)
				})
			return
		}

		var (
			opt, _  = strconv.Atoi(optAction)
			start   = opt * secondRange
			seconds = make([]int, 0)
		)

		for i := range secondRange {
			seconds = append(seconds, start+i)
		}

		exec(logType, optAction, app.LsDbDestServiceV1[opt], app.ChLsDbDestServiceV1[opt], ctx, tableName, stm, qInsertServiceV1, &lastSync, stmLoopServiceV1,
			func(safe string, lastSync *time.Time) ([]*entity.ServiceV1, error) {
				if len(seconds) == 1 {
					return repo.SourceServiceV1.Fetches(safe+"created_at>? AND FLOOR(EXTRACT(SECOND FROM created_at))::INTEGER=?", lastSync, seconds[0], endQuery)
				}

				return repo.SourceServiceV1.Fetches(safe+"created_at>? AND FLOOR(EXTRACT(SECOND FROM created_at))::INTEGER IN (?)", lastSync, seconds, endQuery)
			})

	case "dbq_v1":
		if logType == "p1" {
			exec(logType, optAction, app.LsDbDestDbqV1[0], app.ChLsDbDestDbqV1[0], ctx, tableName, stm, qInsertDbqV1, &lastSync, stmLoopDbqV1,
				func(safe string, lastSync *time.Time) ([]*entity.DbqV1, error) {
					return repo.SourceDbqV1.Fetches(safe+"created_at>?", lastSync, endQuery)
				})
			return
		}

		var (
			opt, _  = strconv.Atoi(optAction)
			start   = opt * secondRange
			seconds = make([]int, 0)
		)

		for i := range secondRange {
			seconds = append(seconds, start+i)
		}

		exec(logType, optAction, app.LsDbDestDbqV1[opt], app.ChLsDbDestDbqV1[opt], ctx, tableName, stm, qInsertDbqV1, &lastSync, stmLoopDbqV1,
			func(safe string, lastSync *time.Time) ([]*entity.DbqV1, error) {
				if len(seconds) == 1 {
					return repo.SourceDbqV1.Fetches(safe+"created_at>? AND FLOOR(EXTRACT(SECOND FROM created_at))::INTEGER=?", lastSync, seconds[0], endQuery)
				}

				return repo.SourceDbqV1.Fetches(safe+"created_at>? AND FLOOR(EXTRACT(SECOND FROM created_at))::INTEGER IN (?)", lastSync, seconds, endQuery)
			})

	case "grpc_v1":
		if logType == "p1" {
			exec(logType, optAction, app.LsDbDestGrpcV1[0], app.ChLsDbDestGrpcV1[0], ctx, tableName, stm, qInsertGrpcV1, &lastSync, stmLoopGrpcV1,
				func(safe string, lastSync *time.Time) ([]*entity.GrpcV1, error) {
					return repo.SourceGrpcV1.Fetches(safe+"created_at>?", lastSync, endQuery)
				})
			return
		}

		var (
			opt, _  = strconv.Atoi(optAction)
			start   = opt * secondRange
			seconds = make([]int, 0)
		)

		for i := range secondRange {
			seconds = append(seconds, start+i)
		}

		exec(logType, optAction, app.LsDbDestGrpcV1[opt], app.ChLsDbDestGrpcV1[opt], ctx, tableName, stm, qInsertGrpcV1, &lastSync, stmLoopGrpcV1,
			func(safe string, lastSync *time.Time) ([]*entity.GrpcV1, error) {
				if len(seconds) == 1 {
					return repo.SourceGrpcV1.Fetches(safe+"created_at>? AND FLOOR(EXTRACT(SECOND FROM created_at))::INTEGER=?", lastSync, seconds[0], endQuery)
				}

				return repo.SourceGrpcV1.Fetches(safe+"created_at>? AND FLOOR(EXTRACT(SECOND FROM created_at))::INTEGER IN (?)", lastSync, seconds, endQuery)
			})

	case "http_call_v1":
		if logType == "p1" {
			exec(logType, optAction, app.LsDbDestHttpCallV1[0], app.ChLsDbDestHttpCallV1[0], ctx, tableName, stm, qInsertHttpCallV1, &lastSync, stmLoopHttpCallV1,
				func(safe string, lastSync *time.Time) ([]*entity.HttpCallV1, error) {
					return repo.SourceHttpCallV1.Fetches(safe+"created_at>?", lastSync, endQuery)
				})
			return
		}

		var (
			opt, _  = strconv.Atoi(optAction)
			start   = opt * secondRange
			seconds = make([]int, 0)
		)

		for i := range secondRange {
			seconds = append(seconds, start+i)
		}

		exec(logType, optAction, app.LsDbDestHttpCallV1[opt], app.ChLsDbDestHttpCallV1[opt], ctx, tableName, stm, qInsertHttpCallV1, &lastSync, stmLoopHttpCallV1,
			func(safe string, lastSync *time.Time) ([]*entity.HttpCallV1, error) {
				if len(seconds) == 1 {
					return repo.SourceHttpCallV1.Fetches(safe+"created_at>? AND FLOOR(EXTRACT(SECOND FROM created_at))::INTEGER=?", lastSync, seconds[0], endQuery)
				}

				return repo.SourceHttpCallV1.Fetches(safe+"created_at>? AND FLOOR(EXTRACT(SECOND FROM created_at))::INTEGER IN (?)", lastSync, seconds, endQuery)
			})
	}
}

func exec[T any](logType string, optAction string, dbConn *pgx.Conn, chDbConn driver.Conn, ctx context.Context, tableName string, stm string, qry string, lastSync *time.Time, loopExec func([]*T, *time.Time, *pgx.Conn, driver.Conn, context.Context, string) error, fetches func(string, *time.Time) ([]*T, error)) {
	var (
		safe        = fmt.Sprintf("created_at < NOW() - INTERVAL '%v' AND ", app.Env.SafeFetch)
		isPrepared  = false
		startedTime time.Time
		oneSecond   = float64(1000)
		oneMinute   = float64(1000 * 60)
		oneHour     = float64(1000 * 60 * 60)
	)

	for {
		startedTime = gm.Util.Timenow()

		ls, err := fetches(safe, lastSync)
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

			// insert into destination
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
			// insert into destination
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
