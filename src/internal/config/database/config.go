package database

// Config - Конфигурация для подключения к БД
type Config struct {
	User     string
	Password string
	Host     string
	Port     int
	DbName   string
}
