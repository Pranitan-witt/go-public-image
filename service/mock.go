package service

import "go-public-image/model"

func getMock() []model.Movie {
	return []model.Movie{
		{
			Id:       "1",
			Title:    "Avatar",
			Director: "Mr.A",
		},
		{
			Id:       "2",
			Title:    "Marvel",
			Director: "Stand Lees",
		}, {
			Id:       "3",
			Title:    "ABC",
			Director: "Director",
		},
	}
}
