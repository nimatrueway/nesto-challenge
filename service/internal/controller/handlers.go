package controller

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-playground/validator/v10"
	"net/http"
	"readcommend/internal/service"
)

type Controller struct {
	service service.BookService
}

func NewController(service service.BookService) *Controller {
	return &Controller{service: service}
}

type BookParams struct {
	Authors  CsvIds `form:"authors" binding:"omitempty"`
	Genres   CsvIds `form:"genres" binding:"omitempty"`
	MinPages int    `form:"min-pages" binding:"omitempty,min=1,max=10000"`
	MaxPages int    `form:"max-pages" binding:"omitempty,min=1,max=10000"`
	MinYear  int    `form:"min-year" binding:"omitempty,min=1800,max=2100"`
	MaxYear  int    `form:"max-year" binding:"omitempty,min=1800,max=2100"`
	Limit    int    `form:"limit,default=100" binding:"omitempty,min=1,max=100"`
}

func (s *Controller) GetBooks(c *gin.Context) {
	var params BookParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	if len(params.Authors.value) > 100 {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "only fewer that 100 authors are allowed to be used as filter"})
		return
	}

	if params.MinPages > params.MaxPages {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "min-pages must be less than or equal to max-pages"})
		return
	}

	if params.MinYear > params.MaxYear {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "min-year must be less than or equal to max-year"})
		return
	}

	books, err := s.service.GetBooks(params.Authors.value, params.Genres.value, params.MinPages, params.MaxPages, params.MinYear, params.MaxYear, params.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "failed to find books"})
		return
	}

	c.JSON(http.StatusOK, books)
}

type AuthorParams struct {
	Search string `binding:"omitempty"               form:"search"`
	Limit  int    `binding:"omitempty,min=1,max=100" form:"limit,default=100"`
}

func (s *Controller) GetAuthors(c *gin.Context) {
	var params AuthorParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	authors, err := s.service.GetAuthors(params.Search, params.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "failed to get authors"})
		return
	}

	c.JSON(http.StatusOK, authors)
}

func (s *Controller) GetGenres(c *gin.Context) {
	genres, err := s.service.GetGenres()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "failed to get genres"})
		return
	}

	c.JSON(http.StatusOK, genres)
}

func (s *Controller) GetSizes(c *gin.Context) {
	sizes, err := s.service.GetSizes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "failed to get sizes"})
		return
	}

	c.JSON(http.StatusOK, sizes)
}

func (s *Controller) GetEras(c *gin.Context) {
	eras, err := s.service.GetEras()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "failed to get eras"})
		return
	}

	c.JSON(http.StatusOK, eras)
}
