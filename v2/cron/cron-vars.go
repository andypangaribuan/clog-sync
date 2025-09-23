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
	lsMxSyncInternal       []sync.Mutex
	lsMxSyncNoteV1         []sync.Mutex
	lsMxSyncServicePieceV1 []sync.Mutex
	lsMxSyncServiceV1      []sync.Mutex
	lsMxSyncDbqV1          []sync.Mutex
	lsMxSyncGrpcV1         []sync.Mutex
	lsMxSyncHttpCallV1     []sync.Mutex

	lsIsSyncInternalRunning       []bool
	lsIsSyncNoteV1Running         []bool
	lsIsSyncServicePieceV1Running []bool
	lsIsSyncServiceV1Running      []bool
	lsIsSyncDbqV1Running          []bool
	lsIsSyncGrpcV1Running         []bool
	lsIsSyncHttpCallV1Running     []bool

	qInsertInternal = strings.TrimSpace(`
INSERT INTO internal (
  created_at, uid, exec_path, exec_function, data,
  error_message, stack_trace
) VALUES (
  $1, $2, $3, $4, $5,
  $6, $7
)`)

	qchInsertInternal = strings.TrimSpace(`
INSERT INTO internal (
  created_at, uid, exec_path, exec_function, data,
  error_message, stack_trace
) VALUES (
  ?, ?, ?, ?, ?,
  ?, ?
)`)

	qInsertNoteV1 = strings.TrimSpace(`
INSERT INTO note_v1 (
  created_at, uid, user_id, partner_id, svc_name,
  svc_version, exec_path, exec_function, key, sub_key,
  data
) VALUES (
  $1, $2, $3, $4, $5,
  $6, $7, $8, $9, $10,
  $11
)`)

	qchInsertNoteV1 = strings.TrimSpace(`
INSERT INTO note_v1 (
  created_at, uid, user_id, partner_id, svc_name,
  svc_version, exec_path, exec_function, key, sub_key,
  data
) VALUES (
  ?, ?, ?, ?, ?,
  ?, ?, ?, ?, ?,
  ?
)`)

	qInsertServicePieceV1 = strings.TrimSpace(`
INSERT INTO service_piece_v1 (
  created_at, uid, svc_name, svc_version, svc_parent_name,
  svc_parent_version, endpoint, url, req_version, req_source,
  req_header, req_param, req_query, req_form, req_body,
  client_ip, started_at
) VALUES (
  $1, $2, $3, $4, $5,
  $6, $7, $8, $9, $10,
  $11, $12
)`)

	qchInsertServicePieceV1 = strings.TrimSpace(`
INSERT INTO service_piece_v1 (
  created_at, uid, svc_name, svc_version, svc_parent_name,
  svc_parent_version, endpoint, url, req_version, req_source,
  req_header, req_param, req_query, req_form, req_body,
  client_ip, started_at
) VALUES (
  ?, ?, ?, ?, ?,
  ?, ?, ?, ?, ?,
  ?, ?, ?, ?, ?,
  ?, ?
)`)

	qInsertServiceV1 = strings.TrimSpace(`
INSERT INTO service_v1 (
  created_at, uid, user_id, partner_id, svc_name,
  svc_version, svc_parent_name, svc_parent_version, endpoint, url,
  severity, exec_path, exec_function, req_version, req_source,
  req_header, req_param, req_query, req_form, req_files,
  req_body, res_data, res_code, res_sub_code, error_message,
  stack_trace, client_ip, duration, started_at, finished_at
) VALUES (
  $1, $2, $3, $4, $5,
  $6, $7, $8, $9, $10,
  $11, $12, $13, $14, $15,
  $16, $17, $18, $19, $20,
  $21, $22, $23, $24, $25,
  $26, $27, $28, $29, $30
)`)

	qchInsertServiceV1 = strings.TrimSpace(`
INSERT INTO service_v1 (
  created_at, uid, user_id, partner_id, svc_name,
  svc_version, svc_parent_name, svc_parent_version, endpoint, url,
  severity, exec_path, exec_function, req_version, req_source,
  req_header, req_param, req_query, req_form, req_files,
  req_body, res_data, res_code, res_sub_code, error_message,
  stack_trace, client_ip, duration, started_at, finished_at
) VALUES (
  ?, ?, ?, ?, ?,
  ?, ?, ?, ?, ?,
  ?, ?, ?, ?, ?,
  ?, ?, ?, ?, ?,
  ?, ?, ?, ?, ?,
  ?, ?, ?, ?, ?
)`)

	qInsertDbqV1 = strings.TrimSpace(`
INSERT INTO dbq_v1 (
  created_at, uid, user_id, partner_id, svc_name,
  svc_version, sql_query, sql_args, severity, exec_path,
  exec_function, error_message, stack_trace, host1, host2,
  duration1, duration2, duration, started_at, finished_at
) VALUES (
  $1, $2, $3, $4, $5,
  $6, $7, $8, $9, $10,
  $11, $12, $13, $14, $15,
  $16, $17, $18, $19, $20
)`)

	qchInsertDbqV1 = strings.TrimSpace(`
INSERT INTO dbq_v1 (
  created_at, uid, user_id, partner_id, svc_name,
  svc_version, sql_query, sql_args, severity, exec_path,
  exec_function, error_message, stack_trace, host1, host2,
  duration1, duration2, duration, started_at, finished_at
) VALUES (
  ?, ?, ?, ?, ?,
  ?, ?, ?, ?, ?,
  ?, ?, ?, ?, ?,
  ?, ?, ?, ?, ?
)`)

	qInsertGrpcV1 = strings.TrimSpace(`
INSERT INTO grpc_v1 (
  created_at, uid, user_id, partner_id, svc_name,
  svc_version, svc_parent_name, svc_parent_version, destination, severity,
  exec_path, exec_function, req_header, data
) VALUES (
  $1, $2, $3, $4, $5,
  $6, $7, $8, $9, $10,
  $11, $12, $13, $14
)`)

	qchInsertGrpcV1 = strings.TrimSpace(`
INSERT INTO grpc_v1 (
  created_at, uid, user_id, partner_id, svc_name,
  svc_version, svc_parent_name, svc_parent_version, destination, severity,
  exec_path, exec_function, req_header, data
) VALUES (
  ?, ?, ?, ?, ?,
  ?, ?, ?, ?, ?,
  ?, ?, ?, ?
)`)

	qInsertHttpCallV1 = strings.TrimSpace(`
INSERT INTO http_call_v1 (
  created_at, uid, user_id, partner_id, svc_name,
  svc_version, url, severity, req_header, req_param,
  req_query, req_form, req_files, req_body, res_data,
  res_code, error_message, stack_trace, duration, started_at,
  finished_at
) VALUES (
  $1, $2, $3, $4, $5,
  $6, $7, $8, $9, $10,
  $11, $12, $13, $14, $15,
  $16, $17, $18, $19, $20,
  $21
)`)

	qchInsertHttpCallV1 = strings.TrimSpace(`
INSERT INTO http_call_v1 (
  created_at, uid, user_id, partner_id, svc_name,
  svc_version, url, severity, req_header, req_param,
  req_query, req_form, req_files, req_body, res_data,
  res_code, error_message, stack_trace, duration, started_at,
  finished_at
) VALUES (
  ?, ?, ?, ?, ?,
  ?, ?, ?, ?, ?,
  ?, ?, ?, ?, ?,
  ?, ?, ?, ?, ?,
  ?,
)`)
)
