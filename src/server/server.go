package server

import (
	"github.com/codegangsta/negroni"
	"src/api/router"
	"src/internal/config"

	"net/http"
)

// Server - структура сервера
type Server struct {
	AppConfig *config.Config
	Router    *router.Router
}

// NewServer Метод для инициализации структуры в DI
func NewServer(appConfig *config.Config, router *router.Router) *Server {
	return &Server{
		AppConfig: appConfig,
		Router:    router,
	}
}

// Run - метод для запуска нашего http-сервера
func (server *Server) Run() {
	ngRouter := server.Router.InitRoutes()
	ngClassic := negroni.Classic()
	ngClassic.UseHandler(ngRouter)
	err := http.ListenAndServe(":5000", ngClassic)
	if err != nil {
		return
	}
}
