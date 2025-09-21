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

var (
	SourceServiceLog      *stuRepo[entity.ServiceLog]
	DestinationServiceLog *stuRepo[entity.ServiceLog]
)

func init() {
	tableName := "service_log"
	columns := `
		id, uid, user_id, partner_id, xid,
		svc_name, svc_version, svc_parent, endpoint, version,
		message, severity, path, function, req_header,
		req_body, req_par, res_data, res_code, data,
		error, stack_trace, client_ip, duration_ms, start_at,
		finish_at, created_at`
	fn := func(e *entity.ServiceLog) []any {
		return []any{
			e.Id, e.Uid, e.UserId, e.PartnerId, e.Xid,
			e.SvcName, e.SvcVersion, e.SvcParent, e.Endpoint, e.Version,
			e.Message, e.Severity, e.Path, e.Function, e.ReqHeader,
			e.ReqBody, e.ReqPar, e.ResData, e.ResCode, e.Data,
			e.Error, e.StackTrace, e.ClientIp, e.DurationMs, e.StartAt,
			e.FinishAt, e.CreatedAt,
		}
	}

	addSource(func(dbi ice.DbInstance) {
		SourceServiceLog = new(dbi, tableName, columns, fn, db.RepoOpt().WithDeletedAtIsNull(false))
	})

	addDestination(func(dbi ice.DbInstance) {
		DestinationServiceLog = new(dbi, tableName, columns, fn, db.RepoOpt().WithDeletedAtIsNull(false))
	})
}
