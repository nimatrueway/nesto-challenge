package api

import (
	"log/slog"

	"readcommend/internal/controller"

	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

func NewRouter(s *controller.Controller, middlewares []gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(sloggin.New(slog.Default()))
	router.Use(gin.Recovery())
	for _, m := range middlewares {
		router.Use(m)
	}
	router.GET("/api/v1/books", s.GetBooks)
	router.GET("/api/v1/authors", s.GetAuthors)
	router.GET("/api/v1/genres", s.GetGenres)
	router.GET("/api/v1/sizes", s.GetSizes)
	router.GET("/api/v1/eras", s.GetEras)
	return router
}
