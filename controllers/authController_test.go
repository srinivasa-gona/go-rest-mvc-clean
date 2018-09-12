package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"rest-api/models"
	"rest-api/services/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_controllerConfig_CreateUser(t *testing.T) {

	mockAuthService := new(mocks.AuthService)
	mockContactsService := new(mocks.ContactsService)

	var jsonStr = []byte(`{"Username" : "user1", "Displayname" : "User One", "Password" : "password"}`)
	req, _ := http.NewRequest("POST", "/create-user", bytes.NewBuffer(jsonStr))

	mockUserInput := models.User{
		Username:    "user1",
		Displayname: "User One",
		Password:    "password",
	}
	mockUser := models.User{
		Userid:      1,
		Username:    "user1",
		Displayname: "User One",
		Password:    "$2a$10$g54dN45HjsLsh7ilp7yv7u.UXwDZGyXfG.MA0EeB9b7Rf.Rtlzz66",
		Token:       "",
	}
	mockAuthService.On("CreateUser", &mockUserInput).Return(mockUser, nil)
	cc := NewcontrollerConfig(mockAuthService, mockContactsService)
	rr := httptest.NewRecorder()
	cc.CreateUser(rr, req)
	body, _ := ioutil.ReadAll(rr.Result().Body)
	var responseUser models.User
	err := json.Unmarshal(body, &responseUser)
	assert.NoError(t, err, "Error in parsing the response ")
	assert.Equal(t, http.StatusOK, rr.Code, "Error code is not 200")

}
