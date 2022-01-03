// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"
)

type Unko struct {
	// 唯一に定まるやーつ
	ID int32 `json:"id"`
	// うんこの種類
	Type string `json:"type"`
	// うんこの数
	Num int32 `json:"num"`
	// 作成日時
	CreatedAt time.Time `json:"createdAt"`
	// 更新日時
	UpdatedAt time.Time `json:"updatedAt"`
	// 削除日時
	DeletedAt sql.NullTime `json:"deletedAt"`
	Color     string       `json:"color"`
}