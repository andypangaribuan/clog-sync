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
	SourceInfoLog      *stuRepo[entity.InfoLog]
	DestinationInfoLog *stuRepo[entity.InfoLog]
)

func init() {
	tableName := "info_log"
	columns := `
		id, uid, user_id, partner_id, xid,
		svc_name, svc_version, svc_parent, message, severity,
		path, function, data, created_at`
	fn := func(e *entity.InfoLog) []any {
		return []any{
			e.Id, e.Uid, e.UserId, e.PartnerId, e.Xid,
			e.SvcName, e.SvcVersion, e.SvcParent, e.Message, e.Severity,
			e.Path, e.Function, e.Data, e.CreatedAt,
		}
	}

	addSource(func(dbi ice.DbInstance) {
		SourceInfoLog = new(dbi, tableName, columns, fn)
	})

	addDestination(func(dbi ice.DbInstance) {
		DestinationInfoLog = new(dbi, tableName, columns, fn)
	})
}
