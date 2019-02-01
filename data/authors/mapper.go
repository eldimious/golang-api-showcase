package authors

import (
	domain "github.com/eldimious/golang-api-showcase/domain/authors"
)

func toDBModel(entity *domain.Author) *Author {
	return &Author{
		Id:      entity.Id,
		Name:    entity.Name,
		Surname: entity.Surname,
	}
}

func toDomainModel(entity *Author) *domain.Author {
	return &domain.Author{
		Id:        entity.Id,
		Name:      entity.Name,
		Surname:   entity.Surname,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
