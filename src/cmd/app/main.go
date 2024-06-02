package main

import (
	"src/internal/app"
	"src/internal/config"
	"src/server"
)

func main() {
	// Вызываем подгрузку конфигурации
	config.LoadEnvironment()

	// Билдим наш контейре с зависимостями
	container := app.BuildContainer()

	// Запускаем наш HTTP-сервер
	err := container.Invoke(func(server *server.Server) {
		server.Run()
	})

	if err != nil {
		panic(err)
	}
}
