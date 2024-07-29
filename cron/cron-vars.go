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

var (
	mxSyncDbqLog     sync.Mutex
	mxSyncInfoLog    sync.Mutex
	mxSyncServiceLog sync.Mutex

	isSyncDbqLogRunning     bool
	isSyncInfoLogRunning    bool
	isSyncServiceLogRunning bool
)
