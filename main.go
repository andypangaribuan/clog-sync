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
)

func main() {
	fm.CallOrderedInit()
	server.Cron(cr)
	server.FuseR(app.Env.RestPort, rest)
}

func cr(router server.RouterC) {
	startUpDelayed := fm.Ptr(time.Second * 3)

	switch app.Env.InfoLogType {
	case "P1":
		fn := func() { cron.SyncTableInfoLog(strings.ToLower(app.Env.InfoLogType), "") }
		router.Every(app.Env.CronRunEvery, fn, startUpDelayed)

	case "P10":
		for i := range 10 {
			fn := func() { cron.SyncTableInfoLog(strings.ToLower(app.Env.InfoLogType), strconv.Itoa(i)) }
			router.Every(app.Env.CronRunEvery, fn, startUpDelayed)
		}

	case "P60":
		for i := range 60 {
			fn := func() { cron.SyncTableInfoLog(strings.ToLower(app.Env.InfoLogType), strconv.Itoa(i)) }
			router.Every(app.Env.CronRunEvery, fn, startUpDelayed)
		}
	}

	switch app.Env.ServiceLogType {
	case "P1":
		fn := func() { cron.SyncTableServiceLog(strings.ToLower(app.Env.ServiceLogType), "") }
		router.Every(app.Env.CronRunEvery, fn, startUpDelayed)

	case "P10":
		for i := range 10 {
			fn := func() { cron.SyncTableServiceLog(strings.ToLower(app.Env.ServiceLogType), strconv.Itoa(i)) }
			router.Every(app.Env.CronRunEvery, fn, startUpDelayed)
		}

	case "P60":
		for i := range 60 {
			fn := func() { cron.SyncTableServiceLog(strings.ToLower(app.Env.ServiceLogType), strconv.Itoa(i)) }
			router.Every(app.Env.CronRunEvery, fn, startUpDelayed)
		}
	}

	switch app.Env.DbqLogType {
	case "P1":
		fn := func() { cron.SyncTableDbqLog(strings.ToLower(app.Env.DbqLogType), "") }
		router.Every(app.Env.CronRunEvery, fn, startUpDelayed)

	case "P10":
		for i := range 10 {
			fn := func() { cron.SyncTableDbqLog(strings.ToLower(app.Env.DbqLogType), strconv.Itoa(i)) }
			router.Every(app.Env.CronRunEvery, fn, startUpDelayed)
		}

	case "P60":
		for i := range 60 {
			fn := func() { cron.SyncTableDbqLog(strings.ToLower(app.Env.DbqLogType), strconv.Itoa(i)) }
			router.Every(app.Env.CronRunEvery, fn, startUpDelayed)
		}
	}
}

func rest(router server.RouterR) {
	router.AutoRecover(app.Env.AppAutoRecover)
	router.PrintOnError(app.Env.AppServerPrintOnError)
	router.NoLog([]string{"GET: /healthz"})

	router.Endpoints(nil, nil, map[string][]func(server.FuseRContext) any{
		"GET: /healthz": {handler.Private.Status},
	})
}
