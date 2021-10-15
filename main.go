package main

import (
	"card_collection/route"
	"card_collection/storage"
)

func main() {
	storage.Init()
	e := route.Init()
	e.Logger.Fatal(e.Start(":8080"))
}
