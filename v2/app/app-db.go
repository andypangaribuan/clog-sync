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
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/andypangaribuan/gmod/gm"
	"github.com/andypangaribuan/gmod/mol"
	"github.com/jackc/pgx/v5"
)

func initDb() {
	source := mol.DbConnection{
		AppName:  Env.AppName,
		Host:     Env.DbSource.Host,
		Port:     Env.DbSource.Port,
		Name:     Env.DbSource.Name,
		Username: Env.DbSource.User,
		Password: Env.DbSource.Pass,
	}

	destination := mol.DbConnection{
		AppName:  Env.AppName,
		Host:     Env.DbDestination.Host,
		Port:     Env.DbDestination.Port,
		Name:     Env.DbDestination.Name,
		Username: Env.DbDestination.User,
		Password: Env.DbDestination.Pass,
	}

	if Env.DbSource.Type == "postgres" {
		DbSource = gm.Db.Postgres(source)
	}

	switch Env.InternalType {
	case "P1":
		LsDbDestInternal, ChLsDbDestInternal = createDestMultiConnection(1, destination, Env.DbDestination.Type)

	case "P10":
		LsDbDestInternal, ChLsDbDestInternal = createDestMultiConnection(10, destination, Env.DbDestination.Type)

	case "P60":
		LsDbDestInternal, ChLsDbDestInternal = createDestMultiConnection(60, destination, Env.DbDestination.Type)
	}

	switch Env.NoteV1Type {
	case "P1":
		LsDbDestNoteV1, ChLsDbDestNoteV1 = createDestMultiConnection(1, destination, Env.DbDestination.Type)

	case "P10":
		LsDbDestNoteV1, ChLsDbDestNoteV1 = createDestMultiConnection(10, destination, Env.DbDestination.Type)

	case "P60":
		LsDbDestNoteV1, ChLsDbDestNoteV1 = createDestMultiConnection(60, destination, Env.DbDestination.Type)
	}

	switch Env.ServicePieceV1Type {
	case "P1":
		LsDbDestServicePieceV1, ChLsDbDestServicePieceV1 = createDestMultiConnection(1, destination, Env.DbDestination.Type)

	case "P10":
		LsDbDestServicePieceV1, ChLsDbDestServicePieceV1 = createDestMultiConnection(10, destination, Env.DbDestination.Type)

	case "P60":
		LsDbDestServicePieceV1, ChLsDbDestServicePieceV1 = createDestMultiConnection(60, destination, Env.DbDestination.Type)
	}

	switch Env.ServiceV1Type {
	case "P1":
		LsDbDestServiceV1, ChLsDbDestServiceV1 = createDestMultiConnection(1, destination, Env.DbDestination.Type)

	case "P10":
		LsDbDestServiceV1, ChLsDbDestServiceV1 = createDestMultiConnection(10, destination, Env.DbDestination.Type)

	case "P60":
		LsDbDestServiceV1, ChLsDbDestServiceV1 = createDestMultiConnection(60, destination, Env.DbDestination.Type)
	}

	switch Env.DbqV1Type {
	case "P1":
		LsDbDestDbqV1, ChLsDbDestDbqV1 = createDestMultiConnection(1, destination, Env.DbDestination.Type)

	case "P10":
		LsDbDestDbqV1, ChLsDbDestDbqV1 = createDestMultiConnection(10, destination, Env.DbDestination.Type)

	case "P60":
		LsDbDestDbqV1, ChLsDbDestDbqV1 = createDestMultiConnection(60, destination, Env.DbDestination.Type)
	}

	switch Env.GrpcV1Type {
	case "P1":
		LsDbDestGrpcV1, ChLsDbDestGrpcV1 = createDestMultiConnection(1, destination, Env.DbDestination.Type)

	case "P10":
		LsDbDestGrpcV1, ChLsDbDestGrpcV1 = createDestMultiConnection(10, destination, Env.DbDestination.Type)

	case "P60":
		LsDbDestGrpcV1, ChLsDbDestGrpcV1 = createDestMultiConnection(60, destination, Env.DbDestination.Type)
	}

	switch Env.HttpCallV1Type {
	case "P1":
		LsDbDestHttpCallV1, ChLsDbDestHttpCallV1 = createDestMultiConnection(1, destination, Env.DbDestination.Type)

	case "P10":
		LsDbDestHttpCallV1, ChLsDbDestHttpCallV1 = createDestMultiConnection(10, destination, Env.DbDestination.Type)

	case "P60":
		LsDbDestHttpCallV1, ChLsDbDestHttpCallV1 = createDestMultiConnection(60, destination, Env.DbDestination.Type)
	}
}

func createDestMultiConnection(total int, dbc mol.DbConnection, dbType string) ([]*pgx.Conn, []driver.Conn) {
	ls := make([]*pgx.Conn, total)
	chLs := make([]driver.Conn, total)

	for i := range total {
		if dbType == "clickhouse" {
			conn, err := clickhouse.Open(&clickhouse.Options{
				Addr: []string{fmt.Sprintf("%v:%v", dbc.Host, dbc.Port)},
				Auth: clickhouse.Auth{
					Database: dbc.Name,
					Username: dbc.Username,
					Password: dbc.Password,
				},
				ClientInfo: clickhouse.ClientInfo{
					Products: []struct {
						Name    string
						Version string
					}{
						{Name: dbc.AppName, Version: gm.Util.Env.GetString("APP_VERSION", "0.0.0")},
					},
				},
				DialContext: func(ctx context.Context, addr string) (net.Conn, error) {
					var d net.Dialer
					return d.DialContext(ctx, "tcp", addr)
				},
				Settings: clickhouse.Settings{
					"max_execution_time": 60,
				},
				Compression: &clickhouse.Compression{
					Method: clickhouse.CompressionLZ4,
				},
				DialTimeout:          time.Second * 30,
				MaxOpenConns:         3,
				MaxIdleConns:         3,
				ConnMaxLifetime:      time.Duration(10) * time.Minute,
				ConnOpenStrategy:     clickhouse.ConnOpenInOrder,
				BlockBufferSize:      10,
				MaxCompressionBuffer: 10240,
			})

			if err != nil {
				log.Fatalf("%+v\n", err)
			}

			if err := conn.Ping(context.Background()); err != nil {
				if ex, ok := err.(*clickhouse.Exception); ok {
					log.Fatalf("exception [%d] %s \n%s\n", ex.Code, ex.Message, ex.StackTrace)
				}

				log.Fatalf("error when ping\n%+v\n", err)
			}

			chLs[i] = conn
			continue
		}

		connStr := fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v", dbc.Username, dbc.Password, dbc.Host, dbc.Port, dbc.Name)
		conn, err := pgx.Connect(context.Background(), connStr)
		if err != nil {
			log.Fatalf("destination connection error\n%v\n", err)
		}

		ls[i] = conn
	}

	return ls, chLs
}
