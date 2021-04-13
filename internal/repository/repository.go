package repository

type Authorization interface {
}

type Post interface {
}

type Repository struct {
	Authorization
	Post
}

func NewRepository() *Repository {
	return &Repository{}
}
