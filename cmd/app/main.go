package main

import (
	"log"

	v1 "github.com/osetr/app/internal/delivery/http/v1"
	"github.com/osetr/app/internal/server"
)

func main() {
	srv := new(server.Server)
	if err := srv.Run("8000", v1.InitRoute()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
