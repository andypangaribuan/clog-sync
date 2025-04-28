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
	"time"

	"github.com/andypangaribuan/gmod/gm"
)

func initEnv() {
	Env = &stuEnv{
		AppName:               gm.Util.Env.GetString("APP_NAME"),
		AppEnv:                gm.Util.Env.GetAppEnv("APP_ENV"),
		AppTimezone:           gm.Util.Env.GetString("APP_TIMEZONE"),
		AppAutoRecover:        gm.Util.Env.GetBool("APP_AUTO_RECOVER"),
		AppServerPrintOnError: gm.Util.Env.GetBool("APP_SERVER_PRINT_ON_ERROR"),
		RestPort:              gm.Util.Env.GetInt("REST_PORT"),

		FetchInterval:     gm.Util.Env.GetString("FETCH_INTERVAL", "10s"),
		FetchDelayStartUp: gm.Util.Env.GetDurationMs("FETCH_DELAY_STARTUP_MS", time.Second*5),
		FetchLimit:        gm.Util.Env.GetInt("FETCH_LIMIT"),

		DbSource: &stuDb{
			Host: gm.Util.Env.GetString("SOURCE_DB_HOST"),
			Port: gm.Util.Env.GetInt("SOURCE_DB_PORT"),
			Name: gm.Util.Env.GetString("SOURCE_DB_NAME"),
			User: gm.Util.Env.GetString("SOURCE_DB_USER"),
			Pass: gm.Util.Env.GetString("SOURCE_DB_PASS"),
		},
		DbDestination: &stuDb{
			Host: gm.Util.Env.GetString("DESTINATION_DB_HOST"),
			Port: gm.Util.Env.GetInt("DESTINATION_DB_PORT"),
			Name: gm.Util.Env.GetString("DESTINATION_DB_NAME"),
			User: gm.Util.Env.GetString("DESTINATION_DB_USER"),
			Pass: gm.Util.Env.GetString("DESTINATION_DB_PASS"),
			Type: gm.Util.Env.GetString("DESTINATION_DB_TYPE"),
		},

		InfoLogType:    gm.Util.Env.GetString("INFO_LOG_TYPE"),
		ServiceLogType: gm.Util.Env.GetString("SERVICE_LOG_TYPE"),
		DbqLogType:     gm.Util.Env.GetString("DBQ_LOG_TYPE"),
	}
}
