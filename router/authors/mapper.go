package authors

import (
	domain "github.com/eldimious/golang-api-showcase/domain/authors"
)

func toResponseModel(entity *domain.Author) *AuthorResponse {
	return &AuthorResponse{
		Id:        entity.Id,
		Name:      entity.Name,
		Surname:   entity.Surname,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
