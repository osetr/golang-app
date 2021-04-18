package domain

import (
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type User struct {
	Id       int    `pg:"id" json:"id"`
	Name     string `pg:"name,unique" json:"name"`
	Email    string `pg:"email,unique" json:"email"`
	Password string `pg:"password" json:"password"`
}

func CreateUserTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.Model(&User{}).CreateTable(opts)
	if createError != nil {
		log.Fatalf("Error occured while creating users table: %v", createError)
		return createError
	}
	return nil
}
