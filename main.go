package main

import (
	"YTStreamGoApi/extractor"
	"YTStreamGoApi/handler"
	"YTStreamGoApi/router"
	"log"
)

func main() {
	r := router.New()

	e := extractor.NewExtractor()
	h := handler.NewHandler(e)
	h.Register(r)
	err := r.Listen(":8080")
	if err != nil {
		log.Fatal(err.Error())
	}
}
