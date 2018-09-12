// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import http "net/http"
import mock "github.com/stretchr/testify/mock"
import models "rest-api/models"

// AuthService is an autogenerated mock type for the AuthService type
type AuthService struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: u
func (_m *AuthService) CreateUser(u *models.User) (models.User, error) {
	ret := _m.Called(u)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(*models.User) models.User); ok {
		r0 = rf(u)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.User) error); ok {
		r1 = rf(u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// JwtAuthentication provides a mock function with given fields: next
func (_m *AuthService) JwtAuthentication(next http.Handler) http.Handler {
	ret := _m.Called(next)

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func(http.Handler) http.Handler); ok {
		r0 = rf(next)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// Login provides a mock function with given fields: username, password
func (_m *AuthService) Login(username string, password string) map[string]interface{} {
	ret := _m.Called(username, password)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(string, string) map[string]interface{}); ok {
		r0 = rf(username, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	return r0
}
