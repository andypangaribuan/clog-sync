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
	SourceServicePieceV1      *stuRepo[entity.ServicePieceV1]
	DestinationServicePieceV1 *stuRepo[entity.ServicePieceV1]
)

func init() {
	tableName := "service_piece_v1"
	columns := `
		created_at, uid, svc_name, svc_version, svc_parent_name,
		svc_parent_version, endpoint, url, req_version, req_source,
		req_header, req_param, req_query, req_form, req_body,
		client_ip, started_at`
	fn := func(e *entity.ServicePieceV1) []any {
		return []any{
			e.CreatedAt, e.Uid, e.SvcName, e.SvcVersion, e.SvcParentName,
			e.SvcParentVersion, e.Endpoint, e.Url, e.ReqVersion, e.ReqSource,
			e.ReqHeader, e.ReqParam, e.ReqQuery, e.ReqForm, e.ReqBody,
			e.ClientIp, e.StartedAt,
		}
	}

	addSource(func(dbi ice.DbInstance) {
		SourceServicePieceV1 = new(dbi, tableName, columns, fn, db.RepoOpt().WithDeletedAtIsNull(false))
	})

	addDestination(func(dbi ice.DbInstance) {
		DestinationServicePieceV1 = new(dbi, tableName, columns, fn, db.RepoOpt().WithDeletedAtIsNull(false))
	})
}
