package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type MusicData struct {
	Releases []MusicItem
	Spending float32
}

type MusicItem struct {
	Id            int
	External_ids  sql.NullString
	Name          string
	Artist        string
	Price         float64
	Seller        sql.NullString
	Note          string
	Purchase_date sql.NullString
}
