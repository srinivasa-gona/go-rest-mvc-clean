package services_test

import (
	"net/http"
	"net/http/httptest"
	"rest-api/models"
	"rest-api/repository/mocks"
	"rest-api/services"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestAuthServiceImpl_Login_Success(t *testing.T) {
	mockAuthRepo := new(mocks.AuthRepository)
	mockUser := models.User{
		Userid:      1,
		Username:    "user1",
		Displayname: "User One",
		Password:    "$2a$10$g54dN45HjsLsh7ilp7yv7u.UXwDZGyXfG.MA0EeB9b7Rf.Rtlzz66",
		Token:       "",
	}
	mockAuthRepo.On("FindUser", mock.AnythingOfType("string")).Return(mockUser, nil)

	config := &models.ConfigProps{
		JWT_TOKEN_PASSWORD: "token_password",
	}
	mockAuthService := services.NewAuthServiceImpl(mockAuthRepo, config)
	_, err := mockAuthService.Login("user1", "password123")
	if err != nil {
		t.Errorf("Login Failed : %v", err)
	}
}

func TestAuthServiceImpl_Login_WrongPassword(t *testing.T) {
	mockAuthRepo := new(mocks.AuthRepository)
	mockUser := models.User{
		Userid:      1,
		Username:    "user1",
		Displayname: "User One",
		Password:    "$2a$10$g54dN45HjsLsh7ilp7yv7u.UXwDZGyXfG.MA0EeB9b7Rf.Rtlzz66",
		Token:       "",
	}
	mockAuthRepo.On("FindUser", mock.AnythingOfType("string")).Return(mockUser, nil)
	config := &models.ConfigProps{
		JWT_TOKEN_PASSWORD: "token_password",
	}

	mockAuthService := services.NewAuthServiceImpl(mockAuthRepo, config)
	_, err := mockAuthService.Login("user1", "password")
	if err.Error() != "Invalid login credentials. Please try again" {
		t.Errorf("Login with incorrect password test failed : %v", err.Error())
	}
}

func TestAuthServiceImpl_JwtAuthentication(t *testing.T) {
	mockAuthRepo := new(mocks.AuthRepository)
	config := &models.ConfigProps{
		JWT_TOKEN_PASSWORD: "token_password",
	}
	mockAuthService := services.NewAuthServiceImpl(mockAuthRepo, config)

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
	rr := httptest.NewRecorder()
	handler := mockAuthService.JwtAuthentication(nextHandler)
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6InNyaW9ub24zIn0.y8_kyOjsJXrdizxlmcrRBrMtaBiFdmGssB-_j3tIebw")
	handler.ServeHTTP(rr, req)
	if rr.Code != 200 {
		t.Errorf("Error in JMT Authentication : %v", rr.Body)
	}

}
