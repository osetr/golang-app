package service

import (
	"errors"
	"testing"
	"time"

	"github.com/osetr/app/internal/domain"
	"github.com/osetr/app/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestPostCreateInputValidate(t *testing.T) {
	testTable := []struct {
		title              string
		description        string
		expctedValidStatus bool
		expectedMessage    map[string]interface{}
	}{
		{
			title:              "qwertyuiopqwe",
			description:        "qwertyuiopqwe",
			expctedValidStatus: true,
			expectedMessage:    map[string]interface{}{},
		},
		{
			title:              "qwe",
			description:        "qwe",
			expctedValidStatus: false,
			expectedMessage: map[string]interface{}{
				"description": []string{"this field must have length in range (10,500)"},
			},
		},
		{
			title:              "q",
			description:        "qwertyuiopqwe",
			expctedValidStatus: false,
			expectedMessage: map[string]interface{}{
				"title": []string{"this field must have length in range (2,36)"},
			},
		},
		{
			title:              "q",
			description:        "q",
			expctedValidStatus: false,
			expectedMessage: map[string]interface{}{
				"description": []string{"this field must have length in range (10,500)"},
				"title":       []string{"this field must have length in range (2,36)"},
			},
		},
	}

	postService := NewPostService(nil)

	for _, testCase := range testTable {
		i := postService.GetPostCreateInput()
		i.Description = testCase.description
		i.Title = testCase.title
		message, valid := i.Validate()
		assert.Equal(t, testCase.expctedValidStatus, valid)
		assert.Equal(t, testCase.expectedMessage, message)
	}
}

func TestPostCreate(t *testing.T) {
	testTable := []struct {
		title         string
		description   string
		expctedPost   *domain.Post
		expectedError error
	}{
		{
			title:       "qwertyuiop",
			description: "qwertyuiopqwe",
			expctedPost: &domain.Post{
				Id:          1,
				Title:       "qwertyuiop",
				Description: "qwertyuiopqwe",
				CreatedDate: time.Time{},
			},
			expectedError: nil,
		},
		{
			title:         "q",
			description:   "q",
			expctedPost:   nil,
			expectedError: errors.New("first you need validate input"),
		},
	}

	for _, testCase := range testTable {
		mockRepo := new(repository.MockPostRepostitory)
		mockRepo.On("Save").Return(testCase.expctedPost, nil)
		postService := NewPostService(mockRepo)
		i := postService.GetPostCreateInput()
		i.Description = testCase.description
		i.Title = testCase.title
		post, err := postService.PostCreate(i)
		assert.Equal(t, testCase.expctedPost, post)
		assert.Equal(t, testCase.expectedError, err)
	}
}

func TestPostList(t *testing.T) {
	testTable := []struct {
		title         string
		description   string
		expctedPosts  []domain.Post
		expectedError error
	}{
		{
			expctedPosts: []domain.Post{
				{
					Id:          1,
					Title:       "qwertyuiop",
					Description: "qwertyuiopqwe",
					CreatedDate: time.Time{},
				},
			},
			expectedError: nil,
		},
	}

	for _, testCase := range testTable {
		mockRepo := new(repository.MockPostRepostitory)
		mockRepo.On("GetAllPosts").Return(testCase.expctedPosts, testCase.expectedError)
		postService := NewPostService(mockRepo)
		posts, err := postService.PostList()
		assert.Equal(t, testCase.expctedPosts, posts)
		assert.Equal(t, testCase.expectedError, err)
	}
}

func TestPostGet(t *testing.T) {
	testTable := []struct {
		title         string
		description   string
		expctedPost   *domain.Post
		expectedError error
	}{
		{
			expctedPost: &domain.Post{
				Id:          1,
				Title:       "qwertyuiop",
				Description: "qwertyuiopqwe",
				CreatedDate: time.Time{},
			},
			expectedError: nil,
		},
	}

	for _, testCase := range testTable {
		mockRepo := new(repository.MockPostRepostitory)
		mockRepo.On("GetSinglePost").Return(testCase.expctedPost, testCase.expectedError)
		postService := NewPostService(mockRepo)
		post, err := postService.PostGet(1)
		assert.Equal(t, testCase.expctedPost, post)
		assert.Equal(t, testCase.expectedError, err)
	}
}
