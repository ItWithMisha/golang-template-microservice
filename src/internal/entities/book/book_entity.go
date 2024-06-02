package book

import "github.com/google/uuid"

// Entity - модель в БД для нашей книги
type Entity struct {
	Uuid uuid.UUID
	Name string
}
