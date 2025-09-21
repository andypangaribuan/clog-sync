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
	"sync"
)

func init() {
	initServicePiece()
	initServiceV1()
	initDbqV1()
}

func initServicePiece() {
	lsMxSyncServicePieceV1 = make([]sync.Mutex, 0)
	lsIsSyncServicePieceRunning = make([]bool, 0)
	size := 60

	switch app.Env.ServicePieceV1Type {
	case "P1":
		size = 1

	case "P10":
		size = 10

	case "P60":
		size = 60
	}

	for range size {
		lsMxSyncServicePieceV1 = append(lsMxSyncServicePieceV1, sync.Mutex{})
		lsIsSyncServicePieceRunning = append(lsIsSyncServicePieceRunning, false)
	}
}

func initServiceV1() {
	lsMxSyncServiceV1 = make([]sync.Mutex, 0)
	lsIsSyncServiceV1Running = make([]bool, 0)
	size := 60

	switch app.Env.ServiceV1Type {
	case "P1":
		size = 1

	case "P10":
		size = 10

	case "P60":
		size = 60
	}

	for range size {
		lsMxSyncServiceV1 = append(lsMxSyncServiceV1, sync.Mutex{})
		lsIsSyncServiceV1Running = append(lsIsSyncServiceV1Running, false)
	}
}

func initDbqV1() {
	lsMxSyncDbqV1 = make([]sync.Mutex, 0)
	lsIsSyncDbqV1Running = make([]bool, 0)
	size := 60

	switch app.Env.DbqV1Type {
	case "P1":
		size = 1

	case "P10":
		size = 10

	case "P60":
		size = 60
	}

	for range size {
		lsMxSyncDbqV1 = append(lsMxSyncDbqV1, sync.Mutex{})
		lsIsSyncDbqV1Running = append(lsIsSyncDbqV1Running, false)
	}
}
