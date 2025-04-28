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
	"sync"
)

func init() {
	initInfoLog()
	initServiceLog()
	initDbqLog()
}

func initInfoLog() {
	lsMxSyncInfoLog = make([]sync.Mutex, 0)
	lsIsSyncInfoLogRunning = make([]bool, 0)
	size := 60

	switch app.Env.InfoLogType {
	case "P1":
		size = 1

	case "P10":
		size = 10

	case "P60":
		size = 60
	}

	for range size {
		lsMxSyncInfoLog = append(lsMxSyncInfoLog, sync.Mutex{})
		lsIsSyncInfoLogRunning = append(lsIsSyncInfoLogRunning, false)
	}
}

func initServiceLog() {
	lsMxSyncServiceLog = make([]sync.Mutex, 0)
	lsIsSyncServiceLogRunning = make([]bool, 0)
	size := 60

	switch app.Env.ServiceLogType {
	case "P1":
		size = 1

	case "P10":
		size = 10

	case "P60":
		size = 60
	}

	for range size {
		lsMxSyncServiceLog = append(lsMxSyncServiceLog, sync.Mutex{})
		lsIsSyncServiceLogRunning = append(lsIsSyncServiceLogRunning, false)
	}
}

func initDbqLog() {
	lsMxSyncDbqLog = make([]sync.Mutex, 0)
	lsIsSyncDbqLogRunning = make([]bool, 0)
	size := 60

	switch app.Env.DbqLogType {
	case "P1":
		size = 1

	case "P10":
		size = 10

	case "P60":
		size = 60
	}

	for range size {
		lsMxSyncDbqLog = append(lsMxSyncDbqLog, sync.Mutex{})
		lsIsSyncDbqLogRunning = append(lsIsSyncDbqLogRunning, false)
	}
}
