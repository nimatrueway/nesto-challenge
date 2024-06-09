package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-playground/validator/v10"
	"net/http"
	"readcommend/internal/repository"
)

type Server struct {
	repo repository.Repository
}

func NewServer(repo repository.Repository) *Server {
	return &Server{repo: repo}
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type BookParams struct {
	Authors  []int `form:"authors" binding:"omitempty"`
	Genres   []int `form:"genres" binding:"omitempty"`
	MinPages int   `form:"min-pages" binding:"omitempty,min=1,max=10000"`
	MaxPages int   `form:"max-pages" binding:"omitempty,min=1,max=10000"`
	MinYear  int   `form:"min-year" binding:"omitempty,min=1800,max=2100"`
	MaxYear  int   `form:"max-year" binding:"omitempty,min=1800,max=2100"`
	Limit    int   `form:"limit" binding:"omitempty,min=1"`
}

func (s *Server) GetBooks(c *gin.Context) {
	var params BookParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	books, err := s.repo.GetBooks(params.Authors, params.Genres, params.MinPages, params.MaxPages, params.MinYear, params.MaxYear, params.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "failed to find books"})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (s *Server) GetAuthors(c *gin.Context) {
	authors, err := s.repo.GetAuthors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "failed to get authors"})
		return
	}

	c.JSON(http.StatusOK, authors)
}

func (s *Server) GetGenres(c *gin.Context) {
	genres, err := s.repo.GetGenres()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "failed to get genres"})
		return
	}

	c.JSON(http.StatusOK, genres)
}

func (s *Server) GetSizes(c *gin.Context) {
	sizes, err := s.repo.GetSizes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "failed to get sizes"})
		return
	}

	c.JSON(http.StatusOK, sizes)
}

func (s *Server) GetEras(c *gin.Context) {
	eras, err := s.repo.GetEras()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "failed to get eras"})
		return
	}

	c.JSON(http.StatusOK, eras)
}
