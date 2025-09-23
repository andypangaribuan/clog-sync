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

func SyncTableHttpCallV1() {
	switch app.Env.HttpCallV1Type {
	case "P1":
		go syncTableHttpCallV1(strings.ToLower(app.Env.HttpCallV1Type), "")

	case "P10":
		for i := range 10 {
			go syncTableHttpCallV1(strings.ToLower(app.Env.HttpCallV1Type), strconv.Itoa(i))
		}

	case "P60":
		for i := range 60 {
			go syncTableHttpCallV1(strings.ToLower(app.Env.HttpCallV1Type), strconv.Itoa(i))
		}
	}
}

func syncTableHttpCallV1(logType string, optAction string) {
	opt := 0
	if optAction != "" {
		opt, _ = strconv.Atoi(optAction)
	}

	lsMxSyncHttpCallV1[opt].Lock()
	defer lsMxSyncHttpCallV1[opt].Unlock()

	if lsIsSyncHttpCallV1Running[opt] {
		return
	}

	lsIsSyncHttpCallV1Running[opt] = true
	go doSync("http_call_v1", logType, optAction, func() {
		lsMxSyncHttpCallV1[opt].Lock()
		defer lsMxSyncHttpCallV1[opt].Unlock()
		lsIsSyncHttpCallV1Running[opt] = false
	})
}
