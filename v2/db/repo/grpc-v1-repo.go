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
	SourceGrpcV1      *stuRepo[entity.GrpcV1]
	DestinationGrpcV1 *stuRepo[entity.GrpcV1]
)

func init() {
	tableName := "grpc_v1"
	columns := `
		created_at, uid, user_id, partner_id, svc_name,
		svc_version, svc_parent_name, svc_parent_version, destination, severity,
		exec_path, exec_function, req_header, data`
	fn := func(e *entity.GrpcV1) []any {
		return []any{
			e.CreatedAt, e.Uid, e.UserId, e.PartnerId, e.SvcName,
			e.SvcVersion, e.SvcParentName, e.SvcParentVersion, e.Destination, e.Severity,
			e.ExecPath, e.ExecFunction, e.ReqHeader, e.Data,
		}
	}

	addSource(func(dbi ice.DbInstance) {
		SourceGrpcV1 = new(dbi, tableName, columns, fn, db.RepoOpt().WithDeletedAtIsNull(false))
	})

	addDestination(func(dbi ice.DbInstance) {
		DestinationGrpcV1 = new(dbi, tableName, columns, fn, db.RepoOpt().WithDeletedAtIsNull(false))
	})
}
