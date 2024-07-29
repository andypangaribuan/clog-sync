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
	"strings"
	"time"

	"github.com/andypangaribuan/gmod/core/db"
	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/gm"
	"github.com/jackc/pgx/v5"
)

func doSync(tableName string, callback func()) {
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

	switch tableName {
	case "dbq_log":
		var (
			qry = strings.TrimSpace(`
INSERT INTO dbq_log (
	id, uid, user_id, partner_id, xid,
	svc_name, svc_version, svc_parent, sql_query, sql_pars,
	severity, path, function, error, stack_trace,
	duration_ms, start_at, finish_at, created_at
) VALUES (
	$1, $2, $3, $4, $5,
	$6, $7, $8, $9, $10,
	$11, $12, $13, $14, $15,
	$16, $17, $18, $19
)`)
			args = func(e *entity.DbqLog) []any {
				return []any{
					e.Id, e.Uid, e.UserId, e.PartnerId, e.Xid,
					e.SvcName, e.SvcVersion, e.SvcParent, trim(e.SqlQuery), ptrTrim(e.SqlPars),
					e.Severity, e.Path, e.Function, ptrTrim(e.Error), ptrTrim(e.StackTrace),
					e.DurationMs, e.StartAt, e.FinishAt, e.CreatedAt,
				}
			}
		)

		exec(app.DbDestDbq, ctx, tableName, stm, qry, &lastSync,
			func(lastSync *time.Time) ([]*entity.DbqLog, error) {
				return repo.SourceDbqLog.Fetches("created_at>?", lastSync, endQuery)
			},
			func(entities []*entity.DbqLog, lastSync *time.Time) error {
				for i := 0; i < len(entities); i++ {
					e := entities[i]
					if lastSync.Before(e.CreatedAt) {
						*lastSync = e.CreatedAt
					}

					_, err = app.DbDestDbq.Exec(ctx, stm, args(e)...)
					if err != nil {
						log.Printf("[db-destination] error when exec statement %v: %+v\n", stm, err)
						return err
					}
				}

				return nil
			})

	case "info_log":
		var (
			qry = strings.TrimSpace(`
INSERT INTO info_log (
	id, uid, user_id, partner_id, xid,
	svc_name, svc_version, svc_parent, message, severity,
	path, function, data, created_at
) VALUES (
	$1, $2, $3, $4, $5,
	$6, $7, $8, $9, $10,
	$11, $12, $13, $14
)`)
			args = func(e *entity.InfoLog) []any {
				return []any{
					e.Id, e.Uid, e.UserId, e.PartnerId, e.Xid,
					e.SvcName, e.SvcVersion, e.SvcParent, e.Message, e.Severity,
					e.Path, e.Function, ptrTrim(e.Data), e.CreatedAt,
				}
			}
		)

		exec(app.DbDestInfo, ctx, tableName, stm, qry, &lastSync,
			func(lastSync *time.Time) ([]*entity.InfoLog, error) {
				return repo.SourceInfoLog.Fetches("created_at>?", lastSync, endQuery)
			},
			func(entities []*entity.InfoLog, lastSync *time.Time) error {
				for i := 0; i < len(entities); i++ {
					e := entities[i]
					if lastSync.Before(e.CreatedAt) {
						*lastSync = e.CreatedAt
					}

					_, err = app.DbDestInfo.Exec(ctx, stm, args(e)...)
					if err != nil {
						log.Printf("[db-destination] error when exec statement %v: %+v\n", stm, err)
						return err
					}
				}

				return nil
			})

	case "service_log":
		var (
			qry = strings.TrimSpace(`
INSERT INTO service_log (
	id, uid, user_id, partner_id, xid,
	svc_name, svc_version, svc_parent, endpoint, version,
	message, severity, path, function, req_header,
	req_body, req_par, res_data, res_code, data,
	error, stack_trace, client_ip, duration_ms, start_at,
	finish_at, created_at
) VALUES (
	$1, $2, $3, $4, $5,
	$6, $7, $8, $9, $10,
	$11, $12, $13, $14, $15,
	$16, $17, $18, $19, $20,
	$21, $22, $23, $24, $25,
	$26, $27
)`)
			args = func(e *entity.ServiceLog) []any {
				var resCode *string
				if e.ResCode != nil {
					resCode = fm.Ptr(strconv.Itoa(*e.ResCode))
				}

				return []any{
					e.Id, e.Uid, e.UserId, e.PartnerId, e.Xid,
					e.SvcName, e.SvcVersion, e.SvcParent, e.Endpoint, e.Version,
					e.Message, e.Severity, e.Path, e.Function, ptrTrim(e.ReqHeader),
					ptrTrim(e.ReqBody), ptrTrim(e.ReqPar), ptrTrim(e.ResData), resCode, ptrTrim(e.Data),
					e.Error, ptrTrim(e.StackTrace), e.ClientIp, e.DurationMs, e.StartAt,
					e.FinishAt, e.CreatedAt,
				}
			}
		)

		exec(app.DbDestService, ctx, tableName, stm, qry, &lastSync,
			func(lastSync *time.Time) ([]*entity.ServiceLog, error) {
				return repo.SourceServiceLog.Fetches("created_at>?", lastSync, endQuery)
			},
			func(entities []*entity.ServiceLog, lastSync *time.Time) error {
				for i := 0; i < len(entities); i++ {
					e := entities[i]
					if lastSync.Before(e.CreatedAt) {
						*lastSync = e.CreatedAt
					}

					_, err = app.DbDestService.Exec(ctx, stm, args(e)...)
					if err != nil {
						log.Printf("[db-destination] error when exec statement %v: %+v\n", stm, err)
						return err
					}
				}

				return nil
			})
	}
}

func exec[T any](dbConn *pgx.Conn, ctx context.Context, tableName string, stm string, qry string, lastSync *time.Time, fetches func(*time.Time) ([]*T, error), loopExec func([]*T, *time.Time) error) {
	var (
		isPrepared  = false
		total       = 0
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

		total = len(ls)
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

		err = loopExec(ls, lastSync)
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

func ptrTrim(val *string) *string {
	if val == nil {
		return nil
	}

	v := trim(*val)
	return &v
}

func trim(val string) string {
	max := 10000
	if len(val) <= max {
		return val
	}

	v := val[0 : max-1]
	return v
}
