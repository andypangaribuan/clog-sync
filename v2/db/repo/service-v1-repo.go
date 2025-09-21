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
	SourceServiceV1      *stuRepo[entity.ServiceV1]
	DestinationServiceV1 *stuRepo[entity.ServiceV1]
)

func init() {
	tableName := "service_v1"
	columns := `
		created_at, uid, user_id, partner_id, svc_name,
		svc_version, svc_parent_name, svc_parent_version, endpoint, url,
		severity, exec_path, exec_function, req_version, req_source,
		req_header, req_param, req_query, req_form, req_files,
		req_body, res_data, res_code, res_sub_code, error_message,
		stack_trace, client_ip, duration, started_at, finished_at`
	fn := func(e *entity.ServiceV1) []any {
		return []any{
			e.CreatedAt, e.Uid, e.UserId, e.PartnerId, e.SvcName,
			e.SvcVersion, e.SvcParentName, e.SvcParentVersion, e.Endpoint, e.Url,
			e.Severity, e.ExecPath, e.ExecFunction, e.ReqVersion, e.ReqSource,
			e.ReqHeader, e.ReqParam, e.ReqQuery, e.ReqForm, e.ReqFiles,
			e.ReqBody, e.ResData, e.ResCode, e.ResSubCode, e.ErrorMessage,
			e.StackTrace, e.ClientIp, e.Duration, e.StartedAt, e.FinishedAt,
		}
	}

	addSource(func(dbi ice.DbInstance) {
		SourceServiceV1 = new(dbi, tableName, columns, fn, db.RepoOpt().WithDeletedAtIsNull(false))
	})

	addDestination(func(dbi ice.DbInstance) {
		DestinationServiceV1 = new(dbi, tableName, columns, fn, db.RepoOpt().WithDeletedAtIsNull(false))
	})
}
