package service

import (
	"bus-api/internal/model"
	"bus-api/internal/repository"
)

func GetAllLines() []model.Line {
	return repository.FindAllLines()
}
