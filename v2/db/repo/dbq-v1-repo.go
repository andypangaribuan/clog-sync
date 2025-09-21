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
	SourceDbqV1      *stuRepo[entity.DbqV1]
	DestinationDbqV1 *stuRepo[entity.DbqV1]
)

func init() {
	tableName := "dbq_v1"
	columns := `
		created_at, uid, user_id, partner_id, svc_name,
		svc_version, sql_query, sql_args, severity, exec_path,
		exec_function, error_message, stack_trace, host1, host2,
		duration1, duration2, duration, started_at, finished_at`
	fn := func(e *entity.DbqV1) []any {
		return []any{
			e.CreatedAt, e.Uid, e.UserId, e.PartnerId, e.SvcName,
			e.SvcVersion, e.SqlQuery, e.SqlArgs, e.Severity, e.ExecPath,
			e.ExecFunction, e.ErrorMessage, e.StackTrace, e.Host1, e.Host2,
			e.Duration1, e.Duration2, e.Duration, e.StartedAt, e.FinishedAt,
		}
	}

	addSource(func(dbi ice.DbInstance) {
		SourceDbqV1 = new(dbi, tableName, columns, fn, db.RepoOpt().WithDeletedAtIsNull(false))
	})

	addDestination(func(dbi ice.DbInstance) {
		DestinationDbqV1 = new(dbi, tableName, columns, fn, db.RepoOpt().WithDeletedAtIsNull(false))
	})
}
