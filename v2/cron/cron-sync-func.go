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
	"strings"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/andypangaribuan/gmod/gm"
	"github.com/jackc/pgx/v5"
)

func argsInternal(e *entity.Internal) []any {
	if app.Env.DbDestination.Type == "clickhouse" {
		createdAt := gm.Conv.Time.ToStrFull(e.CreatedAt)

		if ls := strings.Split(createdAt, " +"); len(ls) == 2 {
			createdAt = ls[0]
		}

		return []any{
			createdAt, e.Uid, e.ExecPath, e.ExecFunction, e.Data,
			e.ErrorMessage, e.StackTrace,
		}
	}

	return []any{
		e.CreatedAt, e.Uid, e.ExecPath, e.ExecFunction, e.Data,
		e.ErrorMessage, e.StackTrace,
	}
}

func argsNoteV1(e *entity.NoteV1) []any {
	if app.Env.DbDestination.Type == "clickhouse" {
		createdAt := gm.Conv.Time.ToStrFull(e.CreatedAt)

		if ls := strings.Split(createdAt, " +"); len(ls) == 2 {
			createdAt = ls[0]
		}

		return []any{
			createdAt, e.Uid, e.UserId, e.PartnerId, e.SvcName,
			e.SvcVersion, e.ExecPath, e.ExecFunction, e.Key, e.SubKey,
			e.Data,
		}
	}

	return []any{
		e.CreatedAt, e.Uid, e.UserId, e.PartnerId, e.SvcName,
		e.SvcVersion, e.ExecPath, e.ExecFunction, e.Key, e.SubKey,
		e.Data,
	}
}

func argsServicePieceV1(e *entity.ServicePieceV1) []any {
	if app.Env.DbDestination.Type == "clickhouse" {
		createdAt := gm.Conv.Time.ToStrFull(e.CreatedAt)
		startedAt := gm.Conv.Time.ToStrFull(e.StartedAt)

		if ls := strings.Split(createdAt, " +"); len(ls) == 2 {
			createdAt = ls[0]
		}

		if ls := strings.Split(startedAt, " +"); len(ls) == 2 {
			startedAt = ls[0]
		}

		return []any{
			createdAt, e.Uid, e.SvcName, e.SvcVersion, e.SvcParentName,
			e.SvcParentVersion, e.Endpoint, e.Url, e.ReqVersion, e.ReqSource,
			e.ReqHeader, e.ReqParam, e.ReqQuery, e.ReqForm, e.ReqBody,
			e.ClientIp, startedAt,
		}
	}

	return []any{
		e.CreatedAt, e.Uid, e.SvcName, e.SvcVersion, e.SvcParentName,
		e.SvcParentVersion, e.Endpoint, e.Url, e.ReqVersion, e.ReqSource,
		e.ReqHeader, e.ReqParam, e.ReqQuery, e.ReqForm, e.ReqBody,
		e.ClientIp, e.StartedAt,
	}
}

func argsServiceV1(e *entity.ServiceV1) []any {
	if app.Env.DbDestination.Type == "clickhouse" {
		createdAt := gm.Conv.Time.ToStrFull(e.CreatedAt)
		startedAt := gm.Conv.Time.ToStrFull(e.StartedAt)
		finishedAt := gm.Conv.Time.ToStrFull(e.FinishedAt)

		if ls := strings.Split(createdAt, " +"); len(ls) == 2 {
			createdAt = ls[0]
		}

		if ls := strings.Split(startedAt, " +"); len(ls) == 2 {
			startedAt = ls[0]
		}

		if ls := strings.Split(finishedAt, " +"); len(ls) == 2 {
			finishedAt = ls[0]
		}

		return []any{
			createdAt, e.Uid, e.UserId, e.PartnerId, e.SvcName,
			e.SvcVersion, e.SvcParentName, e.SvcParentVersion, e.Endpoint, e.Url,
			e.Severity, e.ExecPath, e.ExecFunction, e.ReqVersion, e.ReqSource,
			e.ReqHeader, e.ReqParam, e.ReqQuery, e.ReqForm, e.ReqFiles,
			e.ReqBody, e.ResData, e.ResCode, e.ResSubCode, e.ErrorMessage,
			e.StackTrace, e.ClientIp, e.Duration, startedAt, finishedAt,
		}
	}

	return []any{
		e.CreatedAt, e.Uid, e.UserId, e.PartnerId, e.SvcName,
		e.SvcVersion, e.SvcParentName, e.SvcParentVersion, e.Endpoint, e.Url,
		e.Severity, e.ExecPath, e.ExecFunction, e.ReqVersion, e.ReqSource,
		e.ReqHeader, e.ReqParam, e.ReqQuery, e.ReqForm, e.ReqFiles,
		e.ReqBody, e.ResData, e.ResCode, e.ResSubCode, e.ErrorMessage,
		e.StackTrace, e.ClientIp, e.Duration, e.StartedAt, e.FinishedAt,
	}
}

func argsDbqV1(e *entity.DbqV1) []any {
	if app.Env.DbDestination.Type == "clickhouse" {
		createdAt := gm.Conv.Time.ToStrFull(e.CreatedAt)
		startedAt := gm.Conv.Time.ToStrFull(e.StartedAt)
		finishedAt := gm.Conv.Time.ToStrFull(e.FinishedAt)

		if ls := strings.Split(createdAt, " +"); len(ls) == 2 {
			createdAt = ls[0]
		}

		if ls := strings.Split(startedAt, " +"); len(ls) == 2 {
			startedAt = ls[0]
		}

		if ls := strings.Split(finishedAt, " +"); len(ls) == 2 {
			finishedAt = ls[0]
		}

		return []any{
			createdAt, e.Uid,
			e.UserId, e.PartnerId, e.SvcName, e.SvcVersion, e.SqlQuery,
			e.SqlArgs, e.Severity, e.ExecPath, e.ExecFunction, e.ErrorMessage,
			e.StackTrace, e.Host1, e.Host2, e.Duration1, e.Duration2,
			e.Duration, startedAt, finishedAt,
		}
	}

	return []any{
		e.CreatedAt, e.Uid,
		e.UserId, e.PartnerId, e.SvcName, e.SvcVersion, e.SqlQuery,
		e.SqlArgs, e.Severity, e.ExecPath, e.ExecFunction, e.ErrorMessage,
		e.StackTrace, e.Host1, e.Host2, e.Duration1, e.Duration2,
		e.Duration, e.StartedAt, e.FinishedAt,
	}
}

func argsGrpcV1(e *entity.GrpcV1) []any {
	if app.Env.DbDestination.Type == "clickhouse" {
		createdAt := gm.Conv.Time.ToStrFull(e.CreatedAt)

		if ls := strings.Split(createdAt, " +"); len(ls) == 2 {
			createdAt = ls[0]
		}

		return []any{
			createdAt, e.Uid, e.UserId, e.PartnerId, e.SvcName,
			e.SvcVersion, e.SvcParentName, e.SvcParentVersion, e.Destination, e.Severity,
			e.ExecPath, e.ExecFunction, e.ReqHeader, e.Data,
		}
	}

	return []any{
		e.CreatedAt, e.Uid, e.UserId, e.PartnerId, e.SvcName,
		e.SvcVersion, e.SvcParentName, e.SvcParentVersion, e.Destination, e.Severity,
		e.ExecPath, e.ExecFunction, e.ReqHeader, e.Data,
	}
}

func argsHttpCallV1(e *entity.HttpCallV1) []any {
	if app.Env.DbDestination.Type == "clickhouse" {
		createdAt := gm.Conv.Time.ToStrFull(e.CreatedAt)
		startedAt := gm.Conv.Time.ToStrFull(e.StartedAt)
		finishedAt := gm.Conv.Time.ToStrFull(e.FinishedAt)

		if ls := strings.Split(createdAt, " +"); len(ls) == 2 {
			createdAt = ls[0]
		}

		if ls := strings.Split(startedAt, " +"); len(ls) == 2 {
			startedAt = ls[0]
		}

		if ls := strings.Split(finishedAt, " +"); len(ls) == 2 {
			finishedAt = ls[0]
		}

		if ls := strings.Split(createdAt, " +"); len(ls) == 2 {
			createdAt = ls[0]
		}

		return []any{
			createdAt, e.Uid, e.UserId, e.PartnerId, e.SvcName,
			e.SvcVersion, e.Url, e.Severity, e.ReqHeader, e.ReqParam,
			e.ReqQuery, e.ReqForm, e.ReqFiles, e.ReqBody, e.ResData,
			e.ResCode, e.ErrorMessage, e.StackTrace, e.Duration, startedAt,
			finishedAt,
		}
	}

	return []any{
		e.CreatedAt, e.Uid, e.UserId, e.PartnerId, e.SvcName,
		e.SvcVersion, e.Url, e.Severity, e.ReqHeader, e.ReqParam,
		e.ReqQuery, e.ReqForm, e.ReqFiles, e.ReqBody, e.ResData,
		e.ResCode, e.ErrorMessage, e.StackTrace, e.Duration, e.StartedAt,
		e.FinishedAt,
	}
}

func stmLoopInternal(entities []*entity.Internal, lastSync *time.Time, dbConn *pgx.Conn, chDbConn driver.Conn, ctx context.Context, stm string) error {
	for i := range len(entities) {
		e := entities[i]
		if lastSync.Before(e.CreatedAt) {
			*lastSync = e.CreatedAt
		}

		if app.Env.DbDestination.Type == "clickhouse" {
			err := chDbConn.AsyncInsert(context.Background(), qchInsertInternal, false, argsInternal(e)...)
			if err != nil {
				log.Printf("[db-destination] error when exec async-insert %v: %+v\n", stm, err)
				return err
			}

			continue
		}

		_, err := dbConn.Exec(ctx, stm, argsInternal(e)...)
		if err != nil {
			log.Printf("[db-destination] error when exec statement %v: %+v\n", stm, err)
			return err
		}
	}

	return nil
}

func stmLoopNoteV1(entities []*entity.NoteV1, lastSync *time.Time, dbConn *pgx.Conn, chDbConn driver.Conn, ctx context.Context, stm string) error {
	for i := range len(entities) {
		e := entities[i]
		if lastSync.Before(e.CreatedAt) {
			*lastSync = e.CreatedAt
		}

		if app.Env.DbDestination.Type == "clickhouse" {
			err := chDbConn.AsyncInsert(context.Background(), qchInsertNoteV1, false, argsNoteV1(e)...)
			if err != nil {
				log.Printf("[db-destination] error when exec async-insert %v: %+v\n", stm, err)
				return err
			}

			continue
		}

		_, err := dbConn.Exec(ctx, stm, argsNoteV1(e)...)
		if err != nil {
			log.Printf("[db-destination] error when exec statement %v: %+v\n", stm, err)
			return err
		}
	}

	return nil
}

func stmLoopServicePieceV1(entities []*entity.ServicePieceV1, lastSync *time.Time, dbConn *pgx.Conn, chDbConn driver.Conn, ctx context.Context, stm string) error {
	for i := range len(entities) {
		e := entities[i]
		if lastSync.Before(e.CreatedAt) {
			*lastSync = e.CreatedAt
		}

		if app.Env.DbDestination.Type == "clickhouse" {
			err := chDbConn.AsyncInsert(context.Background(), qchInsertServicePieceV1, false, argsServicePieceV1(e)...)
			if err != nil {
				log.Printf("[db-destination] error when exec async-insert %v: %+v\n", stm, err)
				return err
			}

			continue
		}

		_, err := dbConn.Exec(ctx, stm, argsServicePieceV1(e)...)
		if err != nil {
			log.Printf("[db-destination] error when exec statement %v: %+v\n", stm, err)
			return err
		}
	}

	return nil
}

func stmLoopServiceV1(entities []*entity.ServiceV1, lastSync *time.Time, dbConn *pgx.Conn, chDbConn driver.Conn, ctx context.Context, stm string) error {
	for i := range len(entities) {
		e := entities[i]
		if lastSync.Before(e.CreatedAt) {
			*lastSync = e.CreatedAt
		}

		if app.Env.DbDestination.Type == "clickhouse" {
			err := chDbConn.AsyncInsert(context.Background(), qchInsertServiceV1, false, argsServiceV1(e)...)
			if err != nil {
				log.Printf("[db-destination] error when exec async-insert %v: %+v\n", stm, err)
				return err
			}

			continue
		}

		_, err := dbConn.Exec(ctx, stm, argsServiceV1(e)...)
		if err != nil {
			log.Printf("[db-destination] error when exec statement %v: %+v\n", stm, err)
			return err
		}
	}

	return nil
}

func stmLoopDbqV1(entities []*entity.DbqV1, lastSync *time.Time, dbConn *pgx.Conn, chDbConn driver.Conn, ctx context.Context, stm string) error {
	for i := range len(entities) {
		e := entities[i]
		if lastSync.Before(e.CreatedAt) {
			*lastSync = e.CreatedAt
		}

		if app.Env.DbDestination.Type == "clickhouse" {
			err := chDbConn.AsyncInsert(context.Background(), qchInsertDbqV1, false, argsDbqV1(e)...)
			if err != nil {
				log.Printf("[db-destination] error when exec async-insert %v: %+v\n", stm, err)
				return err
			}

			continue
		}

		_, err := dbConn.Exec(ctx, stm, argsDbqV1(e)...)
		if err != nil {
			log.Printf("[db-destination] error when exec statement %v: %+v\n", stm, err)
			return err
		}
	}

	return nil
}

func stmLoopGrpcV1(entities []*entity.GrpcV1, lastSync *time.Time, dbConn *pgx.Conn, chDbConn driver.Conn, ctx context.Context, stm string) error {
	for i := range len(entities) {
		e := entities[i]
		if lastSync.Before(e.CreatedAt) {
			*lastSync = e.CreatedAt
		}

		if app.Env.DbDestination.Type == "clickhouse" {
			err := chDbConn.AsyncInsert(context.Background(), qchInsertGrpcV1, false, argsGrpcV1(e)...)
			if err != nil {
				log.Printf("[db-destination] error when exec async-insert %v: %+v\n", stm, err)
				return err
			}

			continue
		}

		_, err := dbConn.Exec(ctx, stm, argsGrpcV1(e)...)
		if err != nil {
			log.Printf("[db-destination] error when exec statement %v: %+v\n", stm, err)
			return err
		}
	}

	return nil
}

func stmLoopHttpCallV1(entities []*entity.HttpCallV1, lastSync *time.Time, dbConn *pgx.Conn, chDbConn driver.Conn, ctx context.Context, stm string) error {
	for i := range len(entities) {
		e := entities[i]
		if lastSync.Before(e.CreatedAt) {
			*lastSync = e.CreatedAt
		}

		if app.Env.DbDestination.Type == "clickhouse" {
			err := chDbConn.AsyncInsert(context.Background(), qchInsertHttpCallV1, false, argsHttpCallV1(e)...)
			if err != nil {
				log.Printf("[db-destination] error when exec async-insert %v: %+v\n", stm, err)
				return err
			}

			continue
		}

		_, err := dbConn.Exec(ctx, stm, argsHttpCallV1(e)...)
		if err != nil {
			log.Printf("[db-destination] error when exec statement %v: %+v\n", stm, err)
			return err
		}
	}

	return nil
}
