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

	router.Every(app.Env.CronRunEvery, cron.SyncTableInfoLog, startUpDelayed)
	router.Every(app.Env.CronRunEvery, cron.SyncTableServiceLog, startUpDelayed)
	router.Every(app.Env.CronRunEvery, cron.SyncTableDbqLog, startUpDelayed)
}

func rest(router server.RouterR) {
	router.AutoRecover(app.Env.AppAutoRecover)
	router.PrintOnError(app.Env.AppServerPrintOnError)
	router.NoLog([]string{"GET: /healthz"})

	router.Endpoints(nil, nil, map[string][]func(server.FuseRContext) any{
		"GET: /healthz": {handler.Private.Status},
	})
}
