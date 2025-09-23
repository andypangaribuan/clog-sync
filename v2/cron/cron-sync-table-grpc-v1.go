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

func SyncTableGrpcV1() {
	switch app.Env.GrpcV1Type {
	case "P1":
		go syncTableGrpcV1(strings.ToLower(app.Env.GrpcV1Type), "")

	case "P10":
		for i := range 10 {
			go syncTableGrpcV1(strings.ToLower(app.Env.GrpcV1Type), strconv.Itoa(i))
		}

	case "P60":
		for i := range 60 {
			go syncTableGrpcV1(strings.ToLower(app.Env.GrpcV1Type), strconv.Itoa(i))
		}
	}
}

func syncTableGrpcV1(logType string, optAction string) {
	opt := 0
	if optAction != "" {
		opt, _ = strconv.Atoi(optAction)
	}

	lsMxSyncGrpcV1[opt].Lock()
	defer lsMxSyncGrpcV1[opt].Unlock()

	if lsIsSyncGrpcV1Running[opt] {
		return
	}

	lsIsSyncGrpcV1Running[opt] = true
	go doSync("grpc_v1", logType, optAction, func() {
		lsMxSyncGrpcV1[opt].Lock()
		defer lsMxSyncGrpcV1[opt].Unlock()
		lsIsSyncGrpcV1Running[opt] = false
	})
}
