// utils/mapper.go
package utils

import "database/sql"

func MapSlice[T any, R any](input []T, mapper func(T) R) []R {
	out := make([]R, 0, len(input))
	for _, item := range input {
		out = append(out, mapper(item))
	}
	return out
}

func NullToStr(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}
