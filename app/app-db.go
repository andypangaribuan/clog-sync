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
	"github.com/andypangaribuan/gmod/gm"
	"github.com/andypangaribuan/gmod/mol"
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
	DbDestination = gm.Db.Postgres(destination)
}
