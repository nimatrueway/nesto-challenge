package api

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"readcommend/internal/controller"
	"readcommend/internal/repository/mocks"
	"readcommend/internal/repository/model"
	"readcommend/internal/service"
	"testing"
)

func TestApiV1Books(t *testing.T) {
	repository := mocks.NewMockBookRepository(t)
	router := NewRouter(controller.NewController(service.NewBookService(repository)), nil)

	t.Run("GET /api/v1/books with no query parameter", func(t *testing.T) {
		repository.EXPECT().GetBooks([]int(nil), []int(nil), 0, 0, 0, 0, 0).Return([]model.Book{
			{
				ID:            1,
				Title:         "Book 1",
				YearPublished: 2021,
				Rating:        4.5,
				Pages:         100,
				Genre: model.Genre{
					ID:    1,
					Title: "Genre 1",
				},
				Author: model.Author{
					ID:        1,
					FirstName: "Author 1",
					LastName:  "Author 1",
				},
			},
		}, nil).Once()

		req, err := http.NewRequest("GET", "/api/v1/books", nil)
		assert.NoError(t, err)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, `[{"id":1,"title":"Book 1","yearPublished":2021,"rating":4.5,"pages":100,"genre":{"id":1,"title":"Genre 1"},"author":{"id":1,"firstName":"Author 1","lastName":"Author 1"}}]`, resp.Body.String())
	})

	t.Run("GET /api/v1/books with valid parameters", func(t *testing.T) {
		repository.EXPECT().GetBooks([]int{1, 2, 3}, []int{4, 5, 6}, 10, 100, 1900, 2000, 50).Return(nil, nil).Once()

		req, err := http.NewRequest("GET", "/api/v1/books?authors=1,2,3&genres=4,5,6&min-pages=10&max-pages=100&min-year=1900&max-year=2000&limit=50", nil)
		assert.NoError(t, err)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, `[]`, resp.Body.String())
	})

	t.Run("GET /api/v1/books with invalid parameters", func(t *testing.T) {
		invalidParameters := map[string][]string{
			"authors":   {"-1", "1,2,x"},
			"genres":    {"-1", "1,2,x"},
			"min-pages": {"x", "-1", "10001"},
			"max-pages": {"x", "-1", "10001"},
			"min-year":  {"x", "-1", "1799", "2101"},
			"max-year":  {"x", "-1", "1799", "2101"},
			"limit":     {"x", "-1"},
		}

		for key, values := range invalidParameters {
			for _, value := range values {
				t.Run(key+" = "+value, func(t *testing.T) {
					req, err := http.NewRequest("GET", "/api/v1/books?"+key+"="+value, nil)
					assert.NoError(t, err)
					resp := httptest.NewRecorder()
					router.ServeHTTP(resp, req)

					assert.Equal(t, 400, resp.Code)
					errorResponse := controller.ErrorResponse{}
					err = json.Unmarshal(resp.Body.Bytes(), &errorResponse)
					assert.NoError(t, err)
					assert.NotNil(t, errorResponse.Message)
				})
			}
		}
	})
}
