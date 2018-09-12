package repository

import (
	"database/sql"
	"fmt"
	"log"
	"rest-api/models"

	_ "github.com/lib/pq"
)

func GetCconnection(config *models.ConfigProps) (*sql.DB, error) {

	dsn := ""
	if config.DB_USERNAME != "" && config.DB_PASSWORD != "" {
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", config.DB_USERNAME, config.DB_PASSWORD, config.DB_HOSTNAME, config.DB_PORT, config.DB_SCHEMA,
			config.DB_SSL_MODE)
	} else {
		dsn = fmt.Sprintf("postgres://%s:%d/%s?sslmode=%s", config.DB_HOSTNAME, config.DB_PORT, config.DB_SCHEMA,
			config.DB_SSL_MODE)
	}

	log.Println("DB HOST NAME IS ::: ", config.DB_HOSTNAME )
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error in opening database connection : %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error in getting database connection : %v", err)
	}

	return db, nil

}
