/*
 * Copyright (c) 2025.
 * Created by Andy Pangaribuan (iam.pangaribuan@gmail.com)
 * https://github.com/apangaribuan
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

func SyncTableInternal() {
	switch app.Env.InternalType {
	case "P1":
		go syncTableInternal(strings.ToLower(app.Env.InternalType), "")

	case "P10":
		for i := range 10 {
			go syncTableInternal(strings.ToLower(app.Env.InternalType), strconv.Itoa(i))
		}

	case "P60":
		for i := range 60 {
			go syncTableInternal(strings.ToLower(app.Env.InternalType), strconv.Itoa(i))
		}
	}
}

func syncTableInternal(logType string, optAction string) {
	opt := 0
	if optAction != "" {
		opt, _ = strconv.Atoi(optAction)
	}

	lsMxSyncInternal[opt].Lock()
	defer lsMxSyncInternal[opt].Unlock()

	if lsIsSyncInternalRunning[opt] {
		return
	}

	lsIsSyncInternalRunning[opt] = true
	go doSync("internal", logType, optAction, func() {
		lsMxSyncInternal[opt].Lock()
		defer lsMxSyncInternal[opt].Unlock()
		lsIsSyncInternalRunning[opt] = false
	})
}
