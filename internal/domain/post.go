package domain

type Post struct {
	Id          int    `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
