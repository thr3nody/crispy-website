package repo

import (
	"crispy-website/internal/db"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(dbName string) (*Repository, error) {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		panic("couldn't open db")
	}

	return &Repository{DB: db}, nil
}

func (r *Repository) Close() error {
	return r.DB.Close()
}

func (r *Repository) LoadMusic() []db.MusicItem {
	rows, err := r.DB.Query("SELECT * FROM music ORDER BY purchase_date DESC")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var results []db.MusicItem
	for rows.Next() {
		var result db.MusicItem
		err := rows.Scan(
			&result.Id,
			&result.External_ids,
			&result.Name,
			&result.Artist,
			&result.Price,
			&result.Seller,
			&result.Note,
			&result.Purchase_date)
		if err != nil {
			log.Print(err)
		}
		results = append(results, result)
	}
	return results
}

func (r *Repository) GetSpending() (value float32) {
	err := r.DB.QueryRow("SELECT ROUND(SUM(price), 2) FROM music").Scan(&value)

	if err == sql.ErrNoRows {
		log.Println("No rows found")
	} else if err != nil {
		log.Fatal(err)
	}

	return value
}
