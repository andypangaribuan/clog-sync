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

func SyncTableDbqLog(optAction string) {
	if optAction == "" {
		mxSyncDbqLog.Lock()
		defer mxSyncDbqLog.Unlock()

		if isSyncDbqLogRunning {
			return
		}

		isSyncDbqLogRunning = true
		go doSync("dbq_log", "", func() {
			mxSyncDbqLog.Lock()
			defer mxSyncDbqLog.Unlock()
			isSyncDbqLogRunning = false
		})

		return
	}

	opt, _ := strconv.Atoi(optAction)
	lsMxSyncDbqLog[opt].Lock()
	defer lsMxSyncDbqLog[opt].Unlock()

	if lsIsSyncDbqLogRunning[opt] {
		return
	}

	lsIsSyncDbqLogRunning[opt] = true
	go doSync("dbq_log", optAction, func() {
		lsMxSyncDbqLog[opt].Lock()
		defer lsMxSyncDbqLog[opt].Unlock()
		lsIsSyncDbqLogRunning[opt] = false
	})
}
