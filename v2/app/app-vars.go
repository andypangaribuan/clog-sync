/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package app

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/andypangaribuan/gmod/ice"
	"github.com/jackc/pgx/v5"
)

var (
	Env      *stuEnv
	DbSource ice.DbInstance

	LsDbDestInternal       []*pgx.Conn
	LsDbDestNoteV1         []*pgx.Conn
	LsDbDestServicePieceV1 []*pgx.Conn
	LsDbDestServiceV1      []*pgx.Conn
	LsDbDestDbqV1          []*pgx.Conn
	LsDbDestGrpcV1         []*pgx.Conn
	LsDbDestHttpCallV1     []*pgx.Conn

	ChLsDbDestInternal       []driver.Conn
	ChLsDbDestNoteV1         []driver.Conn
	ChLsDbDestServicePieceV1 []driver.Conn
	ChLsDbDestServiceV1      []driver.Conn
	ChLsDbDestDbqV1          []driver.Conn
	ChLsDbDestGrpcV1         []driver.Conn
	ChLsDbDestHttpCallV1     []driver.Conn
)
