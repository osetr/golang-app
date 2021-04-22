package service

import (
	"errors"
	"testing"

	"github.com/osetr/app/internal/domain"
	"github.com/osetr/app/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestSignUpInputValidate(t *testing.T) {
	testTable := []struct {
		name               string
		email              string
		password           string
		expctedValidStatus bool
		expectedMessage    map[string]interface{}
	}{
		{
			name:               "qwertyuiopqwe",
			email:              "qwertyuiopqwe",
			password:           "123",
			expctedValidStatus: true,
			expectedMessage:    map[string]interface{}{},
		},
		{
			name:               "q",
			email:              "qwertyuiopqwe",
			password:           "123",
			expctedValidStatus: false,
			expectedMessage: map[string]interface{}{
				"name": []string{"this field must have length in range (2,36)"},
			},
		},
		{
			name:               "q",
			email:              "q",
			password:           "123",
			expctedValidStatus: false,
			expectedMessage: map[string]interface{}{
				"email": []string{"this field must have length in range (2,36)"},
				"name":  []string{"this field must have length in range (2,36)"},
			},
		},
		{
			name:               "q",
			email:              "q",
			password:           "1",
			expctedValidStatus: false,
			expectedMessage: map[string]interface{}{
				"email":    []string{"this field must have length in range (2,36)"},
				"name":     []string{"this field must have length in range (2,36)"},
				"password": []string{"this field must have length in range (2,36)"},
			},
		},
	}

	postService := NewAuthService(nil)

	for _, testCase := range testTable {
		i := postService.GetSignUpInput()
		i.Email = testCase.email
		i.Name = testCase.name
		i.Password = testCase.password
		message, valid := i.Validate()
		assert.Equal(t, testCase.expctedValidStatus, valid)
		assert.Equal(t, testCase.expectedMessage, message)
	}
}

func TestUserSignUp(t *testing.T) {
	testTable := []struct {
		name          string
		email         string
		password      string
		mockedUser    *domain.User
		expectedRes   map[string]interface{}
		expectedError error
	}{
		{
			name:     "qwertyuiop",
			email:    "qwertyuiopqwe",
			password: "123",
			mockedUser: &domain.User{
				Id:       1,
				Name:     "qwertyuiop",
				Email:    "qwertyuiopqwe",
				Password: "asd",
			},
			expectedRes: map[string]interface{}{
				"id":    1,
				"name":  "qwertyuiop",
				"email": "qwertyuiopqwe",
			},
			expectedError: nil,
		},
		{
			name:          "q",
			email:         "q",
			password:      "123",
			expectedRes:   nil,
			expectedError: errors.New("first you need validate input"),
		},
	}

	for _, testCase := range testTable {
		mockRepo := new(repository.MockUserRepostitory)
		mockRepo.On("Save").Return(testCase.mockedUser, nil)
		userService := NewAuthService(mockRepo)
		i := userService.GetSignUpInput()
		i.Email = testCase.email
		i.Name = testCase.name
		i.Password = testCase.password
		res, err := userService.SignUp(i)
		assert.Equal(t, testCase.expectedRes, res)
		assert.Equal(t, testCase.expectedError, err)
	}
}
