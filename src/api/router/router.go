package router

import (
	"github.com/gorilla/mux"
	"src/api/controllers/book"
)

// Router - структура описывающие маршрутизацию контроллеров в нашем приложении
type Router struct {
	BookRoutes *book.ControllerRoute
}

// NewRouter Метод для инициализации структуры в DI
func NewRouter(bookRoutes *book.ControllerRoute) *Router {
	return &Router{
		BookRoutes: bookRoutes,
	}
}

// InitRoutes - инициализация маршрутизации API
func (routes *Router) InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = routes.BookRoutes.Route(router)
	return router
}
