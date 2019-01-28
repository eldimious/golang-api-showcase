package authors

import (
	domain "github.com/eldimious/golang-api-showcase/domain/authors"
)

func toDBModel(entity *domain.Author) *Author {
	return &Author{
		ID:      entity.ID,
		Name:    entity.Name,
		Surname: entity.Surname,
	}
}

func toDomainModel(entity *Author) *domain.Author {
	return &domain.Author{
		ID:        entity.ID,
		Name:      entity.Name,
		Surname:   entity.Surname,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		DeletedAt: entity.DeletedAt,
	}
}
