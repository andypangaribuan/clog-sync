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
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/gm"
	"github.com/jackc/pgx/v5"
)

func argsDbqLog(e *entity.DbqLog) []any {
	if app.Env.DbDestination.Type == "clickhouse" {
		startAt := strings.ReplaceAll(gm.Conv.Time.ToStrFull(e.StartAt), " +07:00", "")
		finishAt := strings.ReplaceAll(gm.Conv.Time.ToStrFull(e.FinishAt), " +07:00", "")
		createdAt := strings.ReplaceAll(gm.Conv.Time.ToStrFull(e.CreatedAt), " +07:00", "")

		return []any{
			e.Id, e.Uid, e.UserId, e.PartnerId, e.Xid,
			e.SvcName, e.SvcVersion, e.SvcParent, trim(e.SqlQuery), ptrTrim(e.SqlPars),
			e.Severity, e.Path, e.Function, ptrTrim(e.Error), ptrTrim(e.StackTrace),
			e.DurationMs, startAt, finishAt, createdAt,
		}
	}

	return []any{
		e.Id, e.Uid, e.UserId, e.PartnerId, e.Xid,
		e.SvcName, e.SvcVersion, e.SvcParent, trim(e.SqlQuery), ptrTrim(e.SqlPars),
		e.Severity, e.Path, e.Function, ptrTrim(e.Error), ptrTrim(e.StackTrace),
		e.DurationMs, e.StartAt, e.FinishAt, e.CreatedAt,
	}
}

func argsInfoLog(e *entity.InfoLog) []any {
	if app.Env.DbDestination.Type == "clickhouse" {
		createdAt := strings.ReplaceAll(gm.Conv.Time.ToStrFull(e.CreatedAt), " +07:00", "")

		return []any{
			e.Id, e.Uid, e.UserId, e.PartnerId, e.Xid,
			e.SvcName, e.SvcVersion, e.SvcParent, e.Message, e.Severity,
			e.Path, e.Function, ptrTrim(e.Data), createdAt,
		}
	}

	return []any{
		e.Id, e.Uid, e.UserId, e.PartnerId, e.Xid,
		e.SvcName, e.SvcVersion, e.SvcParent, e.Message, e.Severity,
		e.Path, e.Function, ptrTrim(e.Data), e.CreatedAt,
	}
}

func argsServiceLog(e *entity.ServiceLog) []any {
	var resCode *string
	if e.ResCode != nil {
		resCode = fm.Ptr(strconv.Itoa(*e.ResCode))
	}

	if app.Env.DbDestination.Type == "clickhouse" {
		startAt := strings.ReplaceAll(gm.Conv.Time.ToStrFull(e.StartAt), " +07:00", "")
		finishAt := strings.ReplaceAll(gm.Conv.Time.ToStrFull(e.FinishAt), " +07:00", "")
		createdAt := strings.ReplaceAll(gm.Conv.Time.ToStrFull(e.CreatedAt), " +07:00", "")

		return []any{
			e.Id, e.Uid, e.UserId, e.PartnerId, e.Xid,
			e.SvcName, e.SvcVersion, e.SvcParent, e.Endpoint, e.Version,
			e.Message, e.Severity, e.Path, e.Function, ptrTrim(e.ReqHeader),
			ptrTrim(e.ReqBody), ptrTrim(e.ReqPar), ptrTrim(e.ResData), resCode, ptrTrim(e.Data),
			e.Error, ptrTrim(e.StackTrace), e.ClientIp, e.DurationMs, startAt,
			finishAt, createdAt,
		}
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

func stmLoopDbqLog(entities []*entity.DbqLog, lastSync *time.Time, dbConn *pgx.Conn, chDbConn driver.Conn, ctx context.Context, stm string) error {
	for i := 0; i < len(entities); i++ {
		e := entities[i]
		if lastSync.Before(e.CreatedAt) {
			*lastSync = e.CreatedAt
		}

		if app.Env.DbDestination.Type == "clickhouse" {
			err := chDbConn.AsyncInsert(context.Background(), qchInsertDbqLog, false, argsDbqLog(e)...)
			if err != nil {
				log.Printf("[db-destination] error when exec async-insert %v: %+v\n", stm, err)
				return err
			}

			continue
		}

		_, err := dbConn.Exec(ctx, stm, argsDbqLog(e)...)
		if err != nil {
			log.Printf("[db-destination] error when exec statement %v: %+v\n", stm, err)
			return err
		}
	}

	return nil
}

func stmLoopInfoLog(entities []*entity.InfoLog, lastSync *time.Time, dbConn *pgx.Conn, chDbConn driver.Conn, ctx context.Context, stm string) error {
	for i := 0; i < len(entities); i++ {
		e := entities[i]
		if lastSync.Before(e.CreatedAt) {
			*lastSync = e.CreatedAt
		}

		if app.Env.DbDestination.Type == "clickhouse" {
			err := chDbConn.AsyncInsert(context.Background(), qchInsertInfoLog, false, argsInfoLog(e)...)
			if err != nil {
				log.Printf("[db-destination] error when exec async-insert %v: %+v\n", stm, err)
				return err
			}

			continue
		}

		_, err := dbConn.Exec(ctx, stm, argsInfoLog(e)...)
		if err != nil {
			log.Printf("[db-destination] error when exec statement %v: %+v\n", stm, err)
			return err
		}
	}

	return nil
}

func stmLoopServiceLog(entities []*entity.ServiceLog, lastSync *time.Time, dbConn *pgx.Conn, chDbConn driver.Conn, ctx context.Context, stm string) error {
	for i := 0; i < len(entities); i++ {
		e := entities[i]
		if lastSync.Before(e.CreatedAt) {
			*lastSync = e.CreatedAt
		}

		if app.Env.DbDestination.Type == "clickhouse" {
			err := chDbConn.AsyncInsert(context.Background(), qchInsertServiceLog, false, argsServiceLog(e)...)
			if err != nil {
				log.Printf("[db-destination] error when exec async-insert %v: %+v\n", stm, err)
				return err
			}

			continue
		}

		_, err := dbConn.Exec(ctx, stm, argsServiceLog(e)...)
		if err != nil {
			log.Printf("[db-destination] error when exec statement %v: %+v\n", stm, err)
			return err
		}
	}

	return nil
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
