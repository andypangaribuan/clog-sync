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

func SyncTableInfoLog() {
	switch app.Env.InfoLogType {
	case "P1":
		go syncTableInfoLog(strings.ToLower(app.Env.ServiceLogType), "")

	case "P10":
		for i := range 10 {
			go syncTableInfoLog(strings.ToLower(app.Env.ServiceLogType), strconv.Itoa(i))
		}

	case "P60":
		for i := range 60 {
			go syncTableInfoLog(strings.ToLower(app.Env.ServiceLogType), strconv.Itoa(i))
		}
	}
}

func syncTableInfoLog(logType string, optAction string) {
	opt := 0
	if optAction != "" {
		opt, _ = strconv.Atoi(optAction)
	}

	lsMxSyncInfoLog[opt].Lock()
	defer lsMxSyncInfoLog[opt].Unlock()

	if lsIsSyncInfoLogRunning[opt] {
		return
	}

	lsIsSyncInfoLogRunning[opt] = true
	go doSync("info_log", logType, optAction, func() {
		lsMxSyncInfoLog[opt].Lock()
		defer lsMxSyncInfoLog[opt].Unlock()
		lsIsSyncInfoLogRunning[opt] = false
	})
}
