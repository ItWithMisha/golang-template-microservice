package book

import (
	"encoding/json"
	"net/http"
	bookModels "src/internal/models/book"
	"src/internal/services/book"
)

// Controller - контроллер для работы с книгами
type Controller struct {
	service *book.Service
}

// NewController - мето для регистрации контроллера DI
func NewController(service *book.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (controller *Controller) CreateBook(w http.ResponseWriter, r *http.Request) {
	request := new(bookModels.CreateModel)
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&request)

	controller.service.Create(request)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
