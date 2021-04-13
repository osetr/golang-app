package service

import "github.com/osetr/app/internal/repository"

type Authorization interface {
}

type Post interface {
}

type Service struct {
	Authorization
	Post
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
