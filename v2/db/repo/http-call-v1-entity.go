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
	SourceHttpCallV1      *stuRepo[entity.HttpCallV1]
	DestinationHttpCallV1 *stuRepo[entity.HttpCallV1]
)

func init() {
	tableName := "http_call_v1"
	columns := `
		created_at, uid, user_id, partner_id, svc_name,
		svc_version, url, severity, req_header, req_param,
		req_query, req_form, req_files, req_body, res_data,
		res_code, error_message, stack_trace, duration, started_at,
		finished_at`
	fn := func(e *entity.HttpCallV1) []any {
		return []any{
			e.CreatedAt, e.Uid, e.UserId, e.PartnerId, e.SvcName,
			e.SvcVersion, e.Url, e.Severity, e.ReqHeader, e.ReqParam,
			e.ReqQuery, e.ReqForm, e.ReqFiles, e.ReqBody, e.ResData,
			e.ResCode, e.ErrorMessage, e.StackTrace, e.Duration, e.StartedAt,
			e.FinishedAt,
		}
	}

	addSource(func(dbi ice.DbInstance) {
		SourceHttpCallV1 = new(dbi, tableName, columns, fn, db.RepoOpt().WithDeletedAtIsNull(false))
	})

	addDestination(func(dbi ice.DbInstance) {
		DestinationHttpCallV1 = new(dbi, tableName, columns, fn, db.RepoOpt().WithDeletedAtIsNull(false))
	})
}
