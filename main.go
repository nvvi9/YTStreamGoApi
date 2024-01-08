package main

import (
	"YTStreamGoApi/config"
	"YTStreamGoApi/db"
	"YTStreamGoApi/handler"
	"YTStreamGoApi/router"
	"YTStreamGoApi/store"
	"log"
)

func main() {
	r := router.New()
	config, _ := config.LoadConfig(".")
	d := db.New(config)

	userStore := store.NewUserStore(d)
	h := handler.NewHandler(userStore, config)
	h.Register(r)
	err := r.Listen(":8585")
	if err != nil {
		log.Fatal(err.Error())
	}
}
