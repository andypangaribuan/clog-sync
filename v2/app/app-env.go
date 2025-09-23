/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package app

import "github.com/andypangaribuan/gmod/gm"

func initEnv() {
	Env = &stuEnv{
		AppName:               gm.Util.Env.GetString("APP_NAME"),
		AppEnv:                gm.Util.Env.GetAppEnv("APP_ENV"),
		AppTimezone:           gm.Util.Env.GetString("APP_TIMEZONE"),
		AppAutoRecover:        gm.Util.Env.GetBool("APP_AUTO_RECOVER"),
		AppServerPrintOnError: gm.Util.Env.GetBool("APP_SERVER_PRINT_ON_ERROR"),
		RestPort:              gm.Util.Env.GetInt("REST_PORT"),

		FetchLimit:   gm.Util.Env.GetInt("FETCH_LIMIT"),
		CronRunEvery: gm.Util.Env.GetString("CRON_RUN_EVERY", "10s"),
		SafeFetch:    gm.Util.Env.GetString("SAFE_FETCH", "10m"),

		DbSource: &stuDb{
			Host: gm.Util.Env.GetString("SOURCE_DB_HOST"),
			Port: gm.Util.Env.GetInt("SOURCE_DB_PORT"),
			Name: gm.Util.Env.GetString("SOURCE_DB_NAME"),
			User: gm.Util.Env.GetString("SOURCE_DB_USER"),
			Pass: gm.Util.Env.GetString("SOURCE_DB_PASS"),
			Type: gm.Util.Env.GetString("SOURCE_DB_TYPE"),
		},
		DbDestination: &stuDb{
			Host: gm.Util.Env.GetString("DESTINATION_DB_HOST"),
			Port: gm.Util.Env.GetInt("DESTINATION_DB_PORT"),
			Name: gm.Util.Env.GetString("DESTINATION_DB_NAME"),
			User: gm.Util.Env.GetString("DESTINATION_DB_USER"),
			Pass: gm.Util.Env.GetString("DESTINATION_DB_PASS"),
			Type: gm.Util.Env.GetString("DESTINATION_DB_TYPE"),
		},

		InternalType:       gm.Util.Env.GetString("INTERNAL_TYPE"),
		NoteV1Type:         gm.Util.Env.GetString("NOTE_V1_TYPE"),
		ServicePieceV1Type: gm.Util.Env.GetString("SERVICE_PIECE_V1_TYPE"),
		ServiceV1Type:      gm.Util.Env.GetString("SERVICE_V1_TYPE"),
		DbqV1Type:          gm.Util.Env.GetString("DBQ_V1_TYPE"),
		GrpcV1Type:         gm.Util.Env.GetString("GRPC_V1_TYPE"),
		HttpCallV1Type:     gm.Util.Env.GetString("HTTP_CALL_V1_TYPE"),
	}
}
