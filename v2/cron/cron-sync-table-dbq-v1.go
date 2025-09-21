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

func SyncTableDbqV1() {
	switch app.Env.DbqV1Type {
	case "P1":
		go syncTableDbqV1(strings.ToLower(app.Env.DbqV1Type), "")

	case "P10":
		for i := range 10 {
			go syncTableDbqV1(strings.ToLower(app.Env.DbqV1Type), strconv.Itoa(i))
		}

	case "P60":
		for i := range 60 {
			go syncTableDbqV1(strings.ToLower(app.Env.DbqV1Type), strconv.Itoa(i))
		}
	}
}

func syncTableDbqV1(logType string, optAction string) {
	opt := 0
	if optAction != "" {
		opt, _ = strconv.Atoi(optAction)
	}

	lsMxSyncDbqV1[opt].Lock()
	defer lsMxSyncDbqV1[opt].Unlock()

	if lsIsSyncDbqV1Running[opt] {
		return
	}

	lsIsSyncDbqV1Running[opt] = true
	go doSync("dbq_v1", logType, optAction, func() {
		lsMxSyncDbqV1[opt].Lock()
		defer lsMxSyncDbqV1[opt].Unlock()
		lsIsSyncDbqV1Running[opt] = false
	})
}
