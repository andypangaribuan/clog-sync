/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package cron

import "sync"

func init() {
	lsMxSyncDbqLog = make([]sync.Mutex, 0)
	lsIsSyncDbqLogRunning = make([]bool, 0)

	for i := 0; i < 10; i++ {
		lsMxSyncDbqLog = append(lsMxSyncDbqLog, sync.Mutex{})
		lsIsSyncDbqLogRunning = append(lsIsSyncDbqLogRunning, false)
	}
}
