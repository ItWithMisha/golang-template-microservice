package book

import (
	"github.com/google/uuid"
	bookEntities "src/internal/entities/book"
	bookService "src/internal/models/book"
	"src/internal/repositories/book"
)

// Service - для работы с книгами
type Service struct {
	repository *book.Repository
}

// NewService - метод для регистрации в DI
func NewService(repository *book.Repository) *Service {
	return &Service{repository: repository}
}

// Create - метод создания книги
func (service *Service) Create(model *bookService.CreateModel) {
	// Создаем сущность
	bookEntity := bookEntities.Entity{
		Uuid: uuid.New(),
		Name: model.Name,
	}

	service.repository.Create(bookEntity)
}
