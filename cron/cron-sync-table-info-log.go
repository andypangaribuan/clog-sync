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

func SyncTableInfoLog(logType string, optAction string) {
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
