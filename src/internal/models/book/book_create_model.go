package book

// CreateModel - Модель создания книги
type CreateModel struct {
	Name string `json:"name" form:"name"`
}
