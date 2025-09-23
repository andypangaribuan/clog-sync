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

func SyncTableNoteV1() {
	switch app.Env.NoteV1Type {
	case "P1":
		go syncTableNoteV1(strings.ToLower(app.Env.NoteV1Type), "")

	case "P10":
		for i := range 10 {
			go syncTableNoteV1(strings.ToLower(app.Env.NoteV1Type), strconv.Itoa(i))
		}

	case "P60":
		for i := range 60 {
			go syncTableNoteV1(strings.ToLower(app.Env.NoteV1Type), strconv.Itoa(i))
		}
	}
}

func syncTableNoteV1(logType string, optAction string) {
	opt := 0
	if optAction != "" {
		opt, _ = strconv.Atoi(optAction)
	}

	lsMxSyncNoteV1[opt].Lock()
	defer lsMxSyncNoteV1[opt].Unlock()

	if lsIsSyncNoteV1Running[opt] {
		return
	}

	lsIsSyncNoteV1Running[opt] = true
	go doSync("note_v1", logType, optAction, func() {
		lsMxSyncNoteV1[opt].Lock()
		defer lsMxSyncNoteV1[opt].Unlock()
		lsIsSyncNoteV1Running[opt] = false
	})
}
