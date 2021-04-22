package v1

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/osetr/app/internal/domain"
	"github.com/osetr/app/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestPostCreate(t *testing.T) {
	mockService := new(service.MockPostService)
	postCreateInput := service.PostCreateInput{}
	mockService.On("GetPostCreateInput").Return(postCreateInput)

	testTable := []struct {
		title        string
		description  string
		mockedPost   *domain.Post
		expectedCode int
	}{
		{
			title:       "qwertyuiop",
			description: "qwertyuiopqwe",
			mockedPost: &domain.Post{
				Id:          1,
				Title:       "qwertyuiop",
				Description: "qwertyuiopqwe",
				CreatedDate: time.Time{},
			},
			expectedCode: 201,
		},
	}

	for _, testCase := range testTable {
		mockService.On("PostCreate").Return(testCase.mockedPost, nil)

		var jsonReq = []byte(`{"title":"` + testCase.title + `","description":"` + testCase.description + `"}`)
		req, _ := http.NewRequest("POST", "tests/api/v1/posts", bytes.NewBuffer(jsonReq))

		handler := http.HandlerFunc(NewPostHandler(mockService).createPost)

		response := httptest.NewRecorder()

		handler.ServeHTTP(response, req)

		status := response.Code

		assert.Equal(t, testCase.expectedCode, status)
	}
}
