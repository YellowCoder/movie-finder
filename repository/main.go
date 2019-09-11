package repository

import "github.com/YellowCoder/movie-finder/config"

var MovieRepository = &movieRepository{config.DatabaseConnection}
