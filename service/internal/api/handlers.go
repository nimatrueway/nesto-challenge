package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-playground/validator/v10"
	"net/http"
	"readcommend/internal/service"
	"strconv"
	"strings"
)

type Server struct {
	service service.BookService
}

func NewServer(service service.BookService) *Server {
	return &Server{service: service}
}

type CsvInt struct {
	value []int
}

func (idl *CsvInt) UnmarshalParam(param string) error {
	parts := strings.Split(param, ",")
	for _, part := range parts {
		intPart, err := strconv.Atoi(part)
		if err != nil {
			return fmt.Errorf("\"%s\" is excepted to be a comma-separated list of integers; invalid integer \"%s\"", param, part)
		}
		idl.value = append(idl.value, intPart)
	}
	return nil
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type BookParams struct {
	Authors  CsvInt `form:"authors" binding:"omitempty"`
	Genres   CsvInt `form:"genres" binding:"omitempty"`
	MinPages int    `form:"min-pages" binding:"omitempty,min=1,max=10000"`
	MaxPages int    `form:"max-pages" binding:"omitempty,min=1,max=10000"`
	MinYear  int    `form:"min-year" binding:"omitempty,min=1800,max=2100"`
	MaxYear  int    `form:"max-year" binding:"omitempty,min=1800,max=2100"`
	Limit    int    `form:"limit" binding:"omitempty,min=1"`
}

func (s *Server) GetBooks(c *gin.Context) {
	var params BookParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	books, err := s.service.GetBooks(params.Authors.value, params.Genres.value, params.MinPages, params.MaxPages, params.MinYear, params.MaxYear, params.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "failed to find books"})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (s *Server) GetAuthors(c *gin.Context) {
	authors, err := s.service.GetAuthors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "failed to get authors"})
		return
	}

	c.JSON(http.StatusOK, authors)
}

func (s *Server) GetGenres(c *gin.Context) {
	genres, err := s.service.GetGenres()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "failed to get genres"})
		return
	}

	c.JSON(http.StatusOK, genres)
}

func (s *Server) GetSizes(c *gin.Context) {
	sizes, err := s.service.GetSizes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "failed to get sizes"})
		return
	}

	c.JSON(http.StatusOK, sizes)
}

func (s *Server) GetEras(c *gin.Context) {
	eras, err := s.service.GetEras()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "failed to get eras"})
		return
	}

	c.JSON(http.StatusOK, eras)
}
