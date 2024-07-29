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
	"clog-sync/app"

	"github.com/andypangaribuan/gmod/core/db"
	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/ice"
)

func init() {
	fm.OrderedInit(fm.OrderedLevelRepo, func() {
		for _, fn := range callbacksSource {
			fn(app.DbSource)
		}

		// for _, fn := range callbacksDestination {
		// 	fn(app.DbDestination)
		// }
	})
}

func addSource(callback func(dbi ice.DbInstance)) {
	mx.Lock()
	defer mx.Unlock()
	callbacksSource = append(callbacksSource, callback)
}

func addDestination(callback func(dbi ice.DbInstance)) {
	mx.Lock()
	defer mx.Unlock()
	callbacksDestination = append(callbacksDestination, callback)
}

func new[T any](dbi ice.DbInstance, tableName string, columns string, fn func(e *T) []any, opt ...db.RepoOptBuilder) *stuRepo[T] {
	repo := db.NewRepo[T](dbi, tableName, opt...)
	repo.SetInsert(columns, fn)

	stu := &stuRepo[T]{repo: repo}
	stu.xrepo.repo = repo
	return stu
}
