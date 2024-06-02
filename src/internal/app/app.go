package app

import (
	"go.uber.org/dig"
	"src/api/controllers/book"
	"src/api/router"
	"src/internal/config"
	bookRepository "src/internal/repositories/book"
	bookService "src/internal/services/book"
	"src/server"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	_ = container.Provide(config.NewConfig)
	_ = container.Provide(server.NewServer)
	_ = container.Provide(router.NewRouter)

	buildBook(container)

	return container
}

func buildBook(container *dig.Container) {
	_ = container.Provide(book.NewController)
	_ = container.Provide(book.NewControllerRoute)
	_ = container.Provide(bookService.NewService)
	_ = container.Provide(bookRepository.NewRepository)
}
