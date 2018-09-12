package services

import (
	"net/http"
	"rest-api/models"
)

type AuthService interface {
	JwtAuthentication(next http.Handler) http.Handler
	Login(username, password string) (models.User, error)
	CreateUser(u *models.User) (models.User, error)
}
