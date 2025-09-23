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
	initInternal()
	initServicePiece()
	initServiceV1()
	initDbqV1()
	initGrpcV1()
	initHttpCallV1()
}

func initInternal() {
	lsMxSyncInternal = make([]sync.Mutex, 0)
	lsIsSyncInternalRunning = make([]bool, 0)
	size := 60

	switch app.Env.InternalType {
	case "P1":
		size = 1

	case "P10":
		size = 10

	case "P60":
		size = 60
	}

	for range size {
		lsMxSyncInternal = append(lsMxSyncInternal, sync.Mutex{})
		lsIsSyncInternalRunning = append(lsIsSyncInternalRunning, false)
	}
}

func initServicePiece() {
	lsMxSyncServicePieceV1 = make([]sync.Mutex, 0)
	lsIsSyncServicePieceV1Running = make([]bool, 0)
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
		lsIsSyncServicePieceV1Running = append(lsIsSyncServicePieceV1Running, false)
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

func initGrpcV1() {
	lsMxSyncGrpcV1 = make([]sync.Mutex, 0)
	lsIsSyncGrpcV1Running = make([]bool, 0)
	size := 60

	switch app.Env.GrpcV1Type {
	case "P1":
		size = 1

	case "P10":
		size = 10

	case "P60":
		size = 60
	}

	for range size {
		lsMxSyncGrpcV1 = append(lsMxSyncGrpcV1, sync.Mutex{})
		lsIsSyncGrpcV1Running = append(lsIsSyncGrpcV1Running, false)
	}
}

func initHttpCallV1() {
	lsMxSyncHttpCallV1 = make([]sync.Mutex, 0)
	lsIsSyncHttpCallV1Running = make([]bool, 0)
	size := 60

	switch app.Env.HttpCallV1Type {
	case "P1":
		size = 1

	case "P10":
		size = 10

	case "P60":
		size = 60
	}

	for range size {
		lsMxSyncHttpCallV1 = append(lsMxSyncHttpCallV1, sync.Mutex{})
		lsIsSyncHttpCallV1Running = append(lsIsSyncHttpCallV1Running, false)
	}
}
