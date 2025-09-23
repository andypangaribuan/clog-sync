/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package main

import (
	"clog-sync/app"
	"clog-sync/cron"
	"clog-sync/handler"
	"time"

	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/server"
)

func main() {
	fm.CallOrderedInit()
	server.Cron(cr)
	server.FuseR(app.Env.RestPort, rest)
}

func cr(router server.RouterC) {
	startUpDelayed := fm.Ptr(time.Second * 3)

	router.Every(app.Env.CronRunEvery, cron.SyncTableInternal, startUpDelayed)
	router.Every(app.Env.CronRunEvery, cron.SyncTableNoteV1, startUpDelayed)
	router.Every(app.Env.CronRunEvery, cron.SyncTableServicePieceV1, startUpDelayed)
	router.Every(app.Env.CronRunEvery, cron.SyncTableServiceV1, startUpDelayed)
	router.Every(app.Env.CronRunEvery, cron.SyncTableDbqV1, startUpDelayed)
	router.Every(app.Env.CronRunEvery, cron.SyncTableGrpcV1, startUpDelayed)
	router.Every(app.Env.CronRunEvery, cron.SyncTableHttpCallV1, startUpDelayed)
}

func rest(router server.RouterR) {
	router.AutoRecover(app.Env.AppAutoRecover)
	router.PrintOnError(app.Env.AppServerPrintOnError)
	router.NoLog([]string{"GET: /healthz"})

	router.Endpoints(nil, nil, map[string][]func(server.FuseRContext) any{
		"GET: /healthz": {handler.Private.Status},
	})
}
