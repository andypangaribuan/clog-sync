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

	"github.com/andypangaribuan/gmod/ice"
)

var (
	SourceDbqLog      *stuRepo[entity.DbqLog]
	DestinationDbqLog *stuRepo[entity.DbqLog]
)

func init() {
	tableName := "dbq_log"
	columns := `
		id, uid, user_id, partner_id, xid,
		svc_name, svc_version, svc_parent, sql_query, sql_pars,
		severity, path, function, error, stack_trace,
		duration_ms, start_at, finish_at, created_at`
	fn := func(e *entity.DbqLog) []any {
		return []any{
			e.Id, e.Uid, e.UserId, e.PartnerId, e.Xid,
			e.SvcName, e.SvcVersion, e.SvcParent, e.SqlQuery, e.SqlPars,
			e.Severity, e.Path, e.Function, e.Error, e.StackTrace,
			e.DurationMs, e.StartAt, e.FinishAt, e.CreatedAt,
		}
	}

	addSource(func(dbi ice.DbInstance) {
		SourceDbqLog = new(dbi, tableName, columns, fn)
	})

	addDestination(func(dbi ice.DbInstance) {
		DestinationDbqLog = new(dbi, tableName, columns, fn)
	})
}
