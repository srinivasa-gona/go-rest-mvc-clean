package services

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"rest-api/models"
	"rest-api/repository"
	"rest-api/utils"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	authRepo      repository.AuthRepository
	configuration *models.ConfigProps
}

func NewAuthServiceImpl(ar repository.AuthRepository, configuration *models.ConfigProps) *AuthServiceImpl {
	return &AuthServiceImpl{
		authRepo:      ar,
		configuration: configuration,
	}
}

func (asi *AuthServiceImpl) JwtAuthentication(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		notAuth := []string{"/create-user", "/login"} //List of endpoints that doesn't require auth
		requestPath := r.URL.Path                     //current request path

		//check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range notAuth {

			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		//response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization") //Grab the token from the header

		if tokenHeader == "" { //Token is missing, returns with error code 403 Unauthorized
			err := utils.RestError("Missing auth token", http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.RestRespond(w, nil, err)
			return
		}

		splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		if len(splitted) != 2 {
			err := utils.RestError("Invalid/Malformed auth token", http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.RestRespond(w, nil, err)
			return
		}

		tokenPart := splitted[1] //Grab the token part, what we are truly interested in
		tk := &models.JWTToken{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(asi.configuration.JWT_TOKEN_PASSWORD), nil
		})

		if err != nil { //Malformed token, returns with http code 403 as usual
			err := utils.RestError(err.Error(), http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.RestRespond(w, nil, err)
			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			err := utils.RestError("Token is not valid.", http.StatusForbidden)
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.RestRespond(w, nil, err)
			return
		}

		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		fmt.Printf("User %v  ", tk.Username) //Useful for monitoring
		ctx := context.WithValue(r.Context(), "user", tk.Username)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) //proceed in the middleware chain!
	})
}

func (asi *AuthServiceImpl) Login(username, password string) (models.User, error) {

	user, err := asi.authRepo.FindUser(username)
	if err != nil {
		return user, utils.RestError(err.Error(), http.StatusNotFound)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
			return user, utils.RestError("Invalid login credentials. Please try again", http.StatusForbidden)
		} else {
			return user, utils.RestError("Something went wrong with authentication", http.StatusForbidden)
		}
	}
	//Worked! Logged In
	user.Password = ""

	tk := &models.JWTToken{Username: username}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(asi.configuration.JWT_TOKEN_PASSWORD))
	user.Token = tokenString

	return user, nil
}

func (asi *AuthServiceImpl) CreateUser(u *models.User) (models.User, error) {

	response, err := asi.authRepo.CreateUser(u)
	log.Printf("received response : %v ", response)
	if err != nil {
		log.Println("Error in creating response  ", err.Error())
		return response, utils.RestError(err.Error(), 500)
	}
	log.Println("Returning result")
	return response, nil
}
