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
	"strconv"
	"strings"
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

		switch app.Env.InfoLogType {
		case "P1":
			go cron.SyncTableInfoLog(strings.ToLower(app.Env.InfoLogType), "")

		case "P10":
			for i := range 10 {
				go cron.SyncTableInfoLog(strings.ToLower(app.Env.InfoLogType), strconv.Itoa(i))
			}

		case "P60":
			for i := range 60 {
				go cron.SyncTableInfoLog(strings.ToLower(app.Env.InfoLogType), strconv.Itoa(i))
			}
		}

		switch app.Env.ServiceLogType {
		case "P1":
			go cron.SyncTableServiceLog(strings.ToLower(app.Env.ServiceLogType), "")

		case "P10":
			for i := range 10 {
				go cron.SyncTableServiceLog(strings.ToLower(app.Env.ServiceLogType), strconv.Itoa(i))
			}

		case "P60":
			for i := range 60 {
				go cron.SyncTableServiceLog(strings.ToLower(app.Env.ServiceLogType), strconv.Itoa(i))
			}
		}

		switch app.Env.DbqLogType {
		case "P1":
			go cron.SyncTableDbqLog(strings.ToLower(app.Env.DbqLogType), "")

		case "P10":
			for i := range 10 {
				go cron.SyncTableDbqLog(strings.ToLower(app.Env.DbqLogType), strconv.Itoa(i))
			}

		case "P60":
			for i := range 60 {
				go cron.SyncTableDbqLog(strings.ToLower(app.Env.DbqLogType), strconv.Itoa(i))
			}
		}
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
