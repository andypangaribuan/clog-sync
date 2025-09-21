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

func SyncTableServicePieceV1() {
	switch app.Env.ServicePieceV1Type {
	case "P1":
		go syncTableServicePieceV1(strings.ToLower(app.Env.ServicePieceV1Type), "")

	case "P10":
		for i := range 10 {
			go syncTableServicePieceV1(strings.ToLower(app.Env.ServicePieceV1Type), strconv.Itoa(i))
		}

	case "P60":
		for i := range 60 {
			go syncTableServicePieceV1(strings.ToLower(app.Env.ServicePieceV1Type), strconv.Itoa(i))
		}
	}
}

func syncTableServicePieceV1(logType string, optAction string) {
	opt := 0
	if optAction != "" {
		opt, _ = strconv.Atoi(optAction)
	}

	lsMxSyncServicePieceV1[opt].Lock()
	defer lsMxSyncServicePieceV1[opt].Unlock()

	if lsIsSyncServicePieceRunning[opt] {
		return
	}

	lsIsSyncServicePieceRunning[opt] = true
	go doSync("service_piece_v1", logType, optAction, func() {
		lsMxSyncServicePieceV1[opt].Lock()
		defer lsMxSyncServicePieceV1[opt].Unlock()
		lsIsSyncServicePieceRunning[opt] = false
	})
}
