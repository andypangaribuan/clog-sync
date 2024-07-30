/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package cron

func SyncTableDbqLog(optAction string) {
	mxSyncDbqLog.Lock()
	defer mxSyncDbqLog.Unlock()

	if isSyncDbqLogRunning {
		return
	}

	isSyncDbqLogRunning = true
	go doSync("dbq_log", optAction, func() {
		mxSyncDbqLog.Lock()
		defer mxSyncDbqLog.Unlock()
		isSyncDbqLogRunning = false
	})
}
