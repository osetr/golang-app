package main

import (
	"fmt"
	"log"
	"time"

	"github.com/osetr/app/configs"
	"github.com/osetr/app/internal/dao"
	v1 "github.com/osetr/app/internal/delivery/http/v1"
	"github.com/osetr/app/internal/domain"
	"github.com/osetr/app/internal/repository"
	"github.com/osetr/app/internal/server"
	"github.com/osetr/app/internal/service"
	"github.com/osetr/app/pkg/database"
	"github.com/spf13/viper"
)

func main() {
	// testServerStart()
	testPostDAO()
}

func testPostDAO() {
	postDAO := dao.NewPostDAO()
	fmt.Println(postDAO.UpdatePost(&domain.Post{
		Id:          2,
		Title:       "test1",
		Description: "test",
		CreatedDate: time.Now(),
	}))
}

func testServerStart() {
	configs.GetConfig("config")
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := v1.NewHandler(services)

	if db, err := new(database.СonnectionFactory).GetConnection(); err != nil {
		log.Fatalf("Error occured while connecting to database: %v", err)
	} else {
		domain.CreatePostTable(db)
	}

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoute()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
