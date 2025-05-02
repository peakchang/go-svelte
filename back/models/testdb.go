package models

import (
	"database/sql"
	"go-svelte/utils"
)

type DbTestdb struct {
	ID          int            `db:"id"`
	Name        sql.NullString `db:"name"`
	Age         int            `db:"age"`
	ConfirmType bool           `db:"confirm_type"`
	CreatedAt   sql.NullString `db:"created_at"`
}

type ApiTestdb struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Age         int    `db:"age"`
	ConfirmType bool   `db:"confirm_type"`
	CreatedAt   string `db:"created_at"`
}

func ConvertToAPI(u DbTestdb) ApiTestdb {
	return ApiTestdb{
		ID:          u.ID,
		Age:         u.Age,
		Name:        utils.NullToStr(u.Name),
		ConfirmType: u.ConfirmType,
		CreatedAt:   utils.NullToStr(u.CreatedAt),
	}
}
