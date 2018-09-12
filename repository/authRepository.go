package repository

import "rest-api/models"

type AuthRepository interface {
	CreateUser(p *models.User) (models.User, error)
	FindUser(username string) (models.User, error)
}
