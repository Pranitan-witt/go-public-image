package service

import (
	"go-public-image/model"
	"go-public-image/repository"
)

type MovieService interface {
	GetMovies() (*[]model.Movie, error)
}

type movieService struct {
	repo repository.MovieService
}

func NewMovieService(repoParams repository.MovieService) MovieService {
	return &movieService{
		repo: repoParams,
	}
}

func (h *movieService) GetMovies() (*[]model.Movie, error) {
	// result := getMock()
	result, err := h.repo.GetMovies()
	if err != nil {
		return nil, err
	}
	return result, nil
}
