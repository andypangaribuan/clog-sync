/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package cron

import (
	"strings"
	"sync"
)

var (
	mxSyncInfoLog    sync.Mutex
	mxSyncServiceLog sync.Mutex
	mxSyncDbqLog     sync.Mutex
	lsMxSyncDbqLog   []sync.Mutex

	isSyncInfoLogRunning    bool
	isSyncServiceLogRunning bool
	isSyncDbqLogRunning     bool
	lsIsSyncDbqLogRunning   []bool

	qInsertDbqLog = strings.TrimSpace(`
INSERT INTO dbq_log (
	id, uid, user_id, partner_id, xid,
	svc_name, svc_version, svc_parent, sql_query, sql_pars,
	severity, path, function, error, stack_trace,
	duration_ms, start_at, finish_at, created_at
) VALUES (
	$1, $2, $3, $4, $5,
	$6, $7, $8, $9, $10,
	$11, $12, $13, $14, $15,
	$16, $17, $18, $19
)`)

	qInsertInfoLog = strings.TrimSpace(`
INSERT INTO info_log (
	id, uid, user_id, partner_id, xid,
	svc_name, svc_version, svc_parent, message, severity,
	path, function, data, created_at
) VALUES (
	$1, $2, $3, $4, $5,
	$6, $7, $8, $9, $10,
	$11, $12, $13, $14
)`)

	qInsertServiceLog = strings.TrimSpace(`
INSERT INTO service_log (
	id, uid, user_id, partner_id, xid,
	svc_name, svc_version, svc_parent, endpoint, version,
	message, severity, path, function, req_header,
	req_body, req_par, res_data, res_code, data,
	error, stack_trace, client_ip, duration_ms, start_at,
	finish_at, created_at
) VALUES (
	$1, $2, $3, $4, $5,
	$6, $7, $8, $9, $10,
	$11, $12, $13, $14, $15,
	$16, $17, $18, $19, $20,
	$21, $22, $23, $24, $25,
	$26, $27
)`)
)
