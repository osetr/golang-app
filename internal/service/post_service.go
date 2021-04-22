package service

import (
	"errors"
	"time"

	"github.com/osetr/app/internal/domain"
	"github.com/osetr/app/internal/repository"
)

// Post service implementation
type IPostService interface {
	PostCreate(PostCreateInput) (*domain.Post, error)
	GetPostCreateInput() PostCreateInput
	PostList() ([]domain.Post, error)
	PostGet(int) (*domain.Post, error)
	GetPostUpdateInput() PostUpdateInput
	PostUpdate(int, PostUpdateInput) (*domain.Post, error)
	PostDelete(int) error
}

type PostService struct {
	postRepo repository.IPostRepository
}

func NewPostService(postRepo repository.IPostRepository) IPostService {
	return &PostService{
		postRepo: postRepo,
	}
}

// Post create functionality
func (*PostService) GetPostCreateInput() PostCreateInput {
	return PostCreateInput{}
}

type PostCreateInput struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

func (i *PostCreateInput) Validate() (map[string]interface{}, bool) {
	message := make(map[string]interface{})
	valid := true

	// TODO: make valid validation
	if ul := len(i.Title); ul >= 36 || ul <= 2 {
		message["title"] = []string{"this field must have length in range (2,36)"}
		valid = false
	}
	if ul := len(i.Description); ul >= 500 || ul <= 10 {
		message["description"] = []string{"this field must have length in range (10,500)"}
		valid = false
	}

	return message, valid
}

func (ps *PostService) PostCreate(i PostCreateInput) (*domain.Post, error) {
	if _, valid := i.Validate(); !valid {
		return nil, errors.New("first you need validate input")
	}

	postRepo := ps.postRepo
	p, err := postRepo.Save(&domain.Post{
		Title:       i.Title,
		Description: i.Description,
		CreatedDate: time.Now(),
	})

	return p, err
}

// Post list functionality
func (ps *PostService) PostList() ([]domain.Post, error) {
	postRepo := ps.postRepo
	p, err := postRepo.GetAllPosts()

	return p, err
}

// Post get by id functionality
func (ps *PostService) PostGet(id int) (*domain.Post, error) {
	postRepo := ps.postRepo
	p, err := postRepo.GetSinglePost(id)

	return p, err
}

// Post update functionality
func (ps *PostService) GetPostUpdateInput() PostUpdateInput {
	return PostUpdateInput{}
}

type PostUpdateInput struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

func (i *PostUpdateInput) Validate() (map[string]interface{}, bool) {
	message := make(map[string]interface{})
	valid := true

	// TODO: make valid validation
	if ul := len(i.Title); ul >= 36 || ul <= 2 {
		message["title"] = []string{"this field must have length in range (2,36)"}
		valid = false
	}
	if ul := len(i.Description); ul >= 500 || ul <= 10 {
		message["description"] = []string{"this field must have length in range (10,500)"}
		valid = false
	}

	return message, valid
}

func (ps *PostService) PostUpdate(id int, i PostUpdateInput) (*domain.Post, error) {
	if _, valid := i.Validate(); !valid {
		return nil, errors.New("first you need validate input")
	}

	postRepo := ps.postRepo
	p, err := postRepo.UpdatePost(&domain.Post{
		Id:          id,
		Title:       i.Title,
		Description: i.Description,
	})

	return p, err
}

// Post delete by id functionality
func (ps *PostService) PostDelete(id int) error {
	postRepo := ps.postRepo
	err := postRepo.DeletePost(id)
	return err
}
