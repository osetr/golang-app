package database

import (
	"context"
	"runtime"
	"time"

	"github.com/go-pg/pg/v10"
)

type IСonnectionFactory interface {
	GetConnection() (*pg.DB, error)
}

type СonnectionFactory struct {
	addr     string
	user     string
	pass     string
	database string
}

func NewConnectionFactory(addr string, user string, pass string, database string) IСonnectionFactory {
	return &СonnectionFactory{
		addr:     addr,
		user:     user,
		pass:     pass,
		database: database,
	}
}

func (conFact *СonnectionFactory) GetConnection() (*pg.DB, error) {

	db := pg.Connect(&pg.Options{
		Addr:        ":" + conFact.addr,
		User:        conFact.user,
		Password:    conFact.pass,
		Database:    conFact.database,
		PoolSize:    5 * runtime.NumCPU(),
		PoolTimeout: 30 * time.Second,
	})

	ctx := context.Background()

	if err := db.Ping(ctx); err != nil {
		panic(err)
	}

	return db, nil
}
