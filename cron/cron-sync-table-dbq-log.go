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

func SyncTableDbqLog() {
	switch app.Env.DbqLogType {
	case "P1":
		go syncTableDbqLog(strings.ToLower(app.Env.DbqLogType), "")

	case "P10":
		for i := range 10 {
			go syncTableDbqLog(strings.ToLower(app.Env.DbqLogType), strconv.Itoa(i))
		}

	case "P60":
		for i := range 60 {
			go syncTableDbqLog(strings.ToLower(app.Env.DbqLogType), strconv.Itoa(i))
		}
	}
}

func syncTableDbqLog(logType string, optAction string) {
	opt := 0
	if optAction != "" {
		opt, _ = strconv.Atoi(optAction)
	}

	lsMxSyncDbqLog[opt].Lock()
	defer lsMxSyncDbqLog[opt].Unlock()

	if lsIsSyncDbqLogRunning[opt] {
		return
	}

	lsIsSyncDbqLogRunning[opt] = true
	go doSync("dbq_log", logType, optAction, func() {
		lsMxSyncDbqLog[opt].Lock()
		defer lsMxSyncDbqLog[opt].Unlock()
		lsIsSyncDbqLogRunning[opt] = false
	})
}
