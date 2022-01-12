package main

import (
	"log"

	"github.com/Qwepo/site.git"
	"github.com/Qwepo/site.git/cmd/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(site.Server)
	if err := srv.Run("8080", handlers.InitRouter()); err != nil {
		log.Fatalf("Какая то хуйня братишка, вот те ошибка: %s", err)
	}

}
