package main

import (
	"card_collection/route"
	"card_collection/storage"
	"github.com/labstack/echo/v4"
)

type server struct {
	store  storage.Storage
	router *echo.Echo
}

func main() {
	storage.Init()
	e := route.Init()

	e.Logger.Fatal(e.Start(":8080"))
}

//func newServer() *server {
//	s := &server{
//		store:  storage.NewStoragePostgres(),
//		router: echo.New(),
//	}
//	s.routes()
//	return s
//}
//
//func (s *server) routes() {
//	s.router.GET("/info", s.handleInfo())
//	s.router.POST("/cards", s.handleAddCards())
//	s.router.GET("/cards", s.handleGetAllCards())
//	s.router.GET("/cards/:id", s.handleGetCardById())
//}
//
//
//
//func main() {
//	s := newServer()
//	s.router.Logger.Fatal(s.router.Start(":8080"))
//}
