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
	SourceInternal      *stuRepo[entity.Internal]
	DestinationInternal *stuRepo[entity.Internal]
)

func init() {
	tableName := "internal"
	columns := `
		created_at, uid, exec_path, exec_function, data,
		error_message, stack_trace`
	fn := func(e *entity.Internal) []any {
		return []any{
			e.CreatedAt, e.Uid, e.ExecPath, e.ExecFunction, e.Data,
			e.ErrorMessage, e.StackTrace,
		}
	}

	addSource(func(dbi ice.DbInstance) {
		SourceInternal = new(dbi, tableName, columns, fn, db.RepoOpt().WithDeletedAtIsNull(false))
	})

	addDestination(func(dbi ice.DbInstance) {
		DestinationInternal = new(dbi, tableName, columns, fn, db.RepoOpt().WithDeletedAtIsNull(false))
	})
}
