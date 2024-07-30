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
	"github.com/go-co-op/gocron"
)

func main() {
	fm.CallOrderedInit()
	runCron()
	server.FuseR(app.Env.RestPort, rest)
}

func runCron() {
	var (
		isStartUp = true
		loc, _    = time.LoadLocation(app.Env.AppTimezone)
		scheduler = gocron.NewScheduler(loc)
	)

	_, _ = scheduler.Every(app.Env.FetchInterval).Do(func() {
		if isStartUp {
			isStartUp = false
			time.Sleep(app.Env.FetchDelayStartUp)
		}

		go cron.SyncTableInfoLog()
		go cron.SyncTableServiceLog()

		// normal mode
		// go cron.SyncTableDbqLog("")

		// parallel mode
		go cron.SyncTableDbqLog("0")
		go cron.SyncTableDbqLog("1")
		go cron.SyncTableDbqLog("2")
		go cron.SyncTableDbqLog("3")
		go cron.SyncTableDbqLog("4")
		go cron.SyncTableDbqLog("5")
		go cron.SyncTableDbqLog("6")
		go cron.SyncTableDbqLog("7")
		go cron.SyncTableDbqLog("8")
		go cron.SyncTableDbqLog("9")
	})

	scheduler.StartAsync()
}

func rest(router server.RouterR) {
	router.AutoRecover(app.Env.AppAutoRecover)
	router.PrintOnError(app.Env.AppServerPrintOnError)

	router.Endpoints(nil, nil, map[string][]func(server.FuseRContext) any{
		"GET: /private/status": {handler.Private.Status},
	})
}
