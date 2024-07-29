/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package cron

func SyncInfoLog() {
	mxSyncInfoLog.Lock()
	defer mxSyncInfoLog.Unlock()

	if isSyncInfoLogRunning {
		return
	}

	isSyncInfoLogRunning = true
	go doSync("info_log", func() {
		mxSyncInfoLog.Lock()
		defer mxSyncInfoLog.Unlock()
		isSyncInfoLogRunning = false
	})
}
