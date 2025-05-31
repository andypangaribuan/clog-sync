/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package cron

import (
	"clog-sync/app"
	"strconv"
	"strings"
)

func SyncTableServiceLog() {
	switch app.Env.ServiceLogType {
	case "P1":
		go syncTableServiceLog(strings.ToLower(app.Env.ServiceLogType), "")

	case "P10":
		for i := range 10 {
			go syncTableServiceLog(strings.ToLower(app.Env.ServiceLogType), strconv.Itoa(i))
		}

	case "P60":
		for i := range 60 {
			go syncTableServiceLog(strings.ToLower(app.Env.ServiceLogType), strconv.Itoa(i))
		}
	}
}

func syncTableServiceLog(logType string, optAction string) {
	opt := 0
	if optAction != "" {
		opt, _ = strconv.Atoi(optAction)
	}

	lsMxSyncServiceLog[opt].Lock()
	defer lsMxSyncServiceLog[opt].Unlock()

	if lsIsSyncServiceLogRunning[opt] {
		return
	}

	lsIsSyncServiceLogRunning[opt] = true
	go doSync("service_log", logType, optAction, func() {
		lsMxSyncServiceLog[opt].Lock()
		defer lsMxSyncServiceLog[opt].Unlock()
		lsIsSyncServiceLogRunning[opt] = false
	})
}
