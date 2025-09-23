/*
 * Copyright (c) 2025.
 * Created by Andy Pangaribuan (iam.pangaribuan@gmail.com)
 * https://github.com/apangaribuan
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package repo

import (
	"clog-sync/db/entity"

	"github.com/andypangaribuan/gmod/core/db"
	"github.com/andypangaribuan/gmod/ice"
)

var (
	SourceNoteV1      *stuRepo[entity.NoteV1]
	DestinationNoteV1 *stuRepo[entity.NoteV1]
)

func init() {
	tableName := "note_v1"
	columns := `
		created_at, uid, user_id, partner_id, svc_name,
		svc_version, exec_path, exec_function, key, sub_key,
		data`
	fn := func(e *entity.NoteV1) []any {
		return []any{
			e.CreatedAt, e.Uid, e.UserId, e.PartnerId, e.SvcName,
			e.SvcVersion, e.ExecPath, e.ExecFunction, e.Key, e.SubKey,
			e.Data,
		}
	}

	addSource(func(dbi ice.DbInstance) {
		SourceNoteV1 = new(dbi, tableName, columns, fn, db.RepoOpt().WithDeletedAtIsNull(false))
	})

	addDestination(func(dbi ice.DbInstance) {
		DestinationNoteV1 = new(dbi, tableName, columns, fn, db.RepoOpt().WithDeletedAtIsNull(false))
	})
}
