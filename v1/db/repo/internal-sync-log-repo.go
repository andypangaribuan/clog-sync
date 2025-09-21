/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
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

var InternalSyncLog *stuRepo[entity.InternalSyncLog]

func init() {
	addSource(func(dbi ice.DbInstance) {
		InternalSyncLog = new(dbi, "internal_sync_log",
			"table_name, last_sync",
			func(e *entity.InternalSyncLog) []any {
				return []any{e.TableName, e.LastSync}
			},
			db.RepoOpt().WithDeletedAtIsNull(false))
	})
}
