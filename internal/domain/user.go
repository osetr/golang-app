package domain

import (
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type User struct {
	Id       int    `pg:"id"`
	Name     string `pg:"name,unique"`
	Email    string `pg:"email,unique"`
	Password string `pg:"password"`
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
