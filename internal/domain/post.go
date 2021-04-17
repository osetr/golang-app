package domain

import (
	"log"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type Post struct {
	Id          int       `pg:"id" json:"id"`
	Title       string    `pg:"title,unique" json:"title"`
	Description string    `pg:"description" json:"description"`
	CreatedDate time.Time `pg:"created_date" json:"createDate"`
}

func CreatePostTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.Model(&Post{}).CreateTable(opts)
	if createError != nil {
		log.Fatalf("Error occured while creating posts table: %v", createError)
		return createError
	}
	return nil
}
