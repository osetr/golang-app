package database

import (
	"context"
	"runtime"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/osetr/app/configs"
	"github.com/spf13/viper"
)

type СonnectionFactory struct {
	addr     string
	user     string
	password string
	database string
}

func (conFact *СonnectionFactory) GetConnection() (*pg.DB, error) {
	configs.GetConfig("config")
	conFact.addr = viper.GetString("db.addr")
	conFact.user = viper.GetString("db.user")
	conFact.password = viper.GetString("db.password")
	conFact.database = viper.GetString("db.database")

	db := pg.Connect(&pg.Options{
		Addr:        ":" + conFact.addr,
		User:        conFact.user,
		Password:    conFact.password,
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
