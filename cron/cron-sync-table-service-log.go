/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package cron

func SyncTableServiceLog() {
	mxSyncServiceLog.Lock()
	defer mxSyncServiceLog.Unlock()

	if isSyncServiceLogRunning {
		return
	}

	isSyncServiceLogRunning = true
	go doSync("service_log", "", func() {
		mxSyncServiceLog.Lock()
		defer mxSyncServiceLog.Unlock()
		isSyncServiceLogRunning = false
	})
}
