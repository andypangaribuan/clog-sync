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

	DbSource = gm.Db.Postgres(source)
	// DbDestination = gm.Db.Postgres(destination)
	// DbDestination = gm.Db.Postgres(destination)

	connStr := fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v", destination.Username, destination.Password, destination.Host, destination.Port, destination.Name)
	ls, chLs := createDestMultiConnection(3, connStr)

	DbDestInfo = ls[0]
	DbDestService = ls[1]
	DbDestDbq = ls[2]

	ChDbDestInfo = chLs[0]
	ChDbDestService = chLs[1]
	ChDbDestDbq = chLs[2]

	LsDbDestDbq, ChLsDbDestDbq = createDestMultiConnection(10, connStr)
}

func createDestMultiConnection(total int, connStr string) ([]*pgx.Conn, []driver.Conn) {
	ls := make([]*pgx.Conn, total)
	chLs := make([]driver.Conn, total)

	for i := 0; i < total; i++ {
		if Env.DbDestination.Type == "clickhouse" {
			conn, err := clickhouse.Open(&clickhouse.Options{
				Addr: []string{fmt.Sprintf("%v:%v", Env.DbDestination.Host, Env.DbDestination.Port)},
				Auth: clickhouse.Auth{
					Database: Env.DbDestination.Name,
					Username: Env.DbDestination.User,
					Password: Env.DbDestination.Pass,
				},
				ClientInfo: clickhouse.ClientInfo{
					Products: []struct {
						Name    string
						Version string
					}{
						{Name: Env.AppName, Version: gm.Util.Env.GetString("APP_VERSION", "0.0.0")},
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
				MaxOpenConns:         5,
				MaxIdleConns:         5,
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

		conn, err := pgx.Connect(context.Background(), connStr)
		if err != nil {
			log.Fatalf("destination connection error\n%v\n", err)
		}

		ls[i] = conn
	}

	return ls, chLs
}
