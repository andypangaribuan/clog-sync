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
	"strconv"
	"strings"
)

func SyncTableServiceV1() {
	switch app.Env.ServiceV1Type {
	case "P1":
		go syncTableServiceV1(strings.ToLower(app.Env.ServiceV1Type), "")

	case "P10":
		for i := range 10 {
			go syncTableServiceV1(strings.ToLower(app.Env.ServiceV1Type), strconv.Itoa(i))
		}

	case "P60":
		for i := range 60 {
			go syncTableServiceV1(strings.ToLower(app.Env.ServiceV1Type), strconv.Itoa(i))
		}
	}
}

func syncTableServiceV1(logType string, optAction string) {
	opt := 0
	if optAction != "" {
		opt, _ = strconv.Atoi(optAction)
	}

	lsMxSyncServiceV1[opt].Lock()
	defer lsMxSyncServiceV1[opt].Unlock()

	if lsIsSyncServiceV1Running[opt] {
		return
	}

	lsIsSyncServiceV1Running[opt] = true
	go doSync("service_v1", logType, optAction, func() {
		lsMxSyncServiceV1[opt].Lock()
		defer lsMxSyncServiceV1[opt].Unlock()
		lsIsSyncServiceV1Running[opt] = false
	})
}
