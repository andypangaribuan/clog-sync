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

	connStr := fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v", destination.Username, destination.Password, destination.Host, destination.Port, destination.Name)

	conn1, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("destination connection error\n%v\n", err)
	}

	conn2, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("destination connection error\n%v\n", err)
	}

	conn3, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("destination connection error\n%v\n", err)
	}

	DbDestDbq = conn1
	DbDestInfo = conn2
	DbDestService = conn3
}
