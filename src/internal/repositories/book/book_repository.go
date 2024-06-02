package book

import (
	"fmt"
	"src/internal/entities/book"
)

// Repository - Структура репозитория
type Repository struct {
	database []book.Entity
}

// NewRepository - Метод для регистрации в DI
func NewRepository() *Repository {
	return &Repository{
		database: make([]book.Entity, 0),
	}
}

// Create - добавить книгу
func (repository *Repository) Create(entity book.Entity) {
	repository.database = append(repository.database, entity)
	fmt.Println(repository.database)
}
