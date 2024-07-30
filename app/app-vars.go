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
	"github.com/andypangaribuan/gmod/ice"
	"github.com/jackc/pgx/v5"
)

var (
	DbSource      ice.DbInstance
	Env           *stuEnv
	DbDestInfo    *pgx.Conn
	DbDestService *pgx.Conn
	DbDestDbq     *pgx.Conn
	LsDbDestDbq   []*pgx.Conn
)
