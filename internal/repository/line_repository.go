
package repository

import "bus-api/internal/model"

func FindAllLines() []model.Line {
	return []model.Line{
		{ID: 1, Name: "T12"},
		{ID: 2, Name: "T13"},
		{ID: 3, Name: "Santa"},
	}
}
