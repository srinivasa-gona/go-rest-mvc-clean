package models

type ConfigProps struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOSTNAME string
	DB_PORT     int
	DB_SCHEMA   string
	DB_SSL_MODE string
	JWT_TOKEN_PASSWORD string
}
