/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package cron

import "strconv"

func SyncTableDbqLog(logType string, optAction string) {
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
