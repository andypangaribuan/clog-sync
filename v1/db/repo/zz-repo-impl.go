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
	"github.com/andypangaribuan/gmod/core/db"
	"github.com/andypangaribuan/gmod/ice"
)

func (slf *xrepo[T]) Fetch(condition string, args ...any) (*T, error) {
	return slf.repo.Fetch(nil, condition, args...)
}

func (slf *xrepo[T]) Fetches(condition string, args ...any) ([]*T, error) {
	return slf.repo.Fetches(nil, condition, args...)
}

func (slf *xrepo[T]) Insert(e *T) error {
	return slf.repo.Insert(nil, e)
}

func (slf *xrepo[T]) Update(builder db.UpdateBuilder) error {
	return slf.repo.Update(nil, builder)
}

func (slf *xrepo[T]) TxBulkInsert(tx ice.DbTx, entities []*T, chunkSize ...int) error {
	return slf.repo.TxBulkInsert(nil, tx, entities, chunkSize...)
}
