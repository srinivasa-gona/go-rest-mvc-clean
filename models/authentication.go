package models

import (
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Userid      int64
	Username    string
	Displayname string
	Password    string
	Token       string
}

type JWTToken struct {
	Username string
	jwt.StandardClaims
}
