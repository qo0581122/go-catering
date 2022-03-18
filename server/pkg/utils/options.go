package utils

import (
	. "catering/proto/common"
)

type Options func(sql string) (string, interface{})

type OptionsNoParam func(sql string) string

func WithColumn(column string, data interface{}) Options {
	return func(sql string) (string, interface{}) {
		sql += " AND '" + column + "' = ?"
		return sql, data
	}
}

func WithOrderBy(orderByOptions OrderByOptions) Options {
	return func(sql string) (string, interface{}) {
		temp := " ORDER BY "
		for _, opt := range orderByOptions.GetOptions() {
			column := opt.Column
			desc := opt.Desc
			temp += column
			if desc == 1 {
				temp += " desc,"
			} else {
				temp += " asc,"
			}
		}
		return sql + temp, ""
	}
}
