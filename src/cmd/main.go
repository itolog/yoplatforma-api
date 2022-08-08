package main

import (
	"api_platforma/src/internal/app"
	"log"
)

func main() {
	err := app.Start()
	if err != nil {
		log.Fatal(err)
	}
}
