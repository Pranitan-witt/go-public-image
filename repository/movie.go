package repository

import (
	"database/sql"
	"log"

	"go-public-image/model"
)

type MovieService interface {
	GetMovies() (*[]model.Movie, error)
}

type movieRepo struct {
	db *sql.DB
}

func NewMovieRepository(db *sql.DB) MovieService {
	return &movieRepo{
		db: db,
	}
}

func (h *movieRepo) GetMovies() (*[]model.Movie, error) {
	rows, err := h.db.Query("SELECT id, title, director FROM Movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the result set
	result := make([]model.Movie, 0)
	for rows.Next() {
		temp := model.Movie{}
		err := rows.Scan(&temp.Id, &temp.Title, &temp.Director)
		if err != nil {
			return nil, err
		}
		result = append(result, temp)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return &result, nil
}
