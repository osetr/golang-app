package service

import (
	"time"

	"github.com/osetr/app/internal/domain"
	"github.com/osetr/app/internal/repository"
)

type PostCreateService struct {
	Input struct {
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
	}
}

func (s *PostCreateService) Validate() (bool, map[string]interface{}) {
	message := make(map[string]interface{})
	i := s.Input
	valid := true

	if ul := len(i.Title); ul >= 36 || ul <= 2 {
		message["title"] = []string{"this field must have length in range (2,36)"}
		valid = false
	}
	if ul := len(i.Description); ul >= 500 || ul <= 10 {
		message["description"] = []string{"this field must have length in range (10,500)"}
		valid = false
	}

	return valid, message
}

func (s *PostCreateService) Execute() (*domain.Post, error) {
	if valid, _ := s.Validate(); !valid {
		panic("First you need validate input")
	}

	postRepo := repository.NewRepository().PostRepository
	p, err := postRepo.Save(&domain.Post{
		Title:       s.Input.Title,
		Description: s.Input.Description,
		CreatedDate: time.Now(),
	})

	return p, err
}

type PostListService struct {
}

func (s *PostListService) Execute() ([]domain.Post, error) {
	postRepo := repository.NewRepository().PostRepository
	p, err := postRepo.GetAllPosts()

	return p, err
}

type PostGetService struct {
}

func (s *PostGetService) Execute(id int) (*domain.Post, error) {
	postRepo := repository.NewRepository().PostRepository
	p, err := postRepo.GetSinglePost(id)

	return p, err
}

type PostDeleteService struct {
}

func (s *PostDeleteService) Execute(id int) error {
	postRepo := repository.NewRepository().PostRepository
	err := postRepo.DeletePost(id)
	return err
}

type PostUpdateService struct {
	Input struct {
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
	}
}

func (s *PostUpdateService) Validate() (bool, map[string]interface{}) {
	message := make(map[string]interface{})
	i := s.Input
	valid := true

	if ul := len(i.Title); ul >= 36 || ul <= 2 {
		message["title"] = []string{"this field must have length in range (2,36)"}
		valid = false
	}
	if ul := len(i.Description); ul >= 500 || ul <= 10 {
		message["description"] = []string{"this field must have length in range (10,500)"}
		valid = false
	}

	return valid, message
}

func (s *PostUpdateService) Execute(id int) (*domain.Post, error) {
	if valid, _ := s.Validate(); !valid {
		panic("First you need validate input")
	}

	postRepo := repository.NewRepository().PostRepository
	p, err := postRepo.UpdatePost(&domain.Post{
		Id:          id,
		Title:       s.Input.Title,
		Description: s.Input.Description,
	})

	return p, err
}
