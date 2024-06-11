package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"readcommend/internal/controller"
	"readcommend/internal/repository/mocks"
	"readcommend/internal/repository/model"
	"readcommend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type APIV1BooksTestSuite struct {
	router     *gin.Engine
	repository *mocks.MockBookRepository
	suite.Suite
}

func TestApiV1BooksTestSuite(t *testing.T) {
	suite.Run(t, new(APIV1BooksTestSuite))
}

func (suite *APIV1BooksTestSuite) SetupTest() {
	suite.repository = mocks.NewMockBookRepository(suite.T())
	suite.router = NewRouter(controller.NewController(service.NewBookService(suite.repository)), nil)
}

func (suite *APIV1BooksTestSuite) TestDefaults() {
	suite.repository.EXPECT().GetBooks([]int(nil), []int(nil), 0, 0, 0, 0, 100).Return([]model.Book{
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

	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, "GET", "/api/v1/books", nil)

	suite.Require().NoError(err)
	resp := httptest.NewRecorder()
	suite.router.ServeHTTP(resp, req)

	suite.Equal(200, resp.Code)
	suite.Equal(`[{"id":1,"title":"Book 1","yearPublished":2021,"rating":4.5,"pages":100,"genre":{"id":1,"title":"Genre 1"},"author":{"id":1,"firstName":"Author 1","lastName":"Author 1"}}]`, resp.Body.String())
}

func (suite *APIV1BooksTestSuite) TestValidParameters() {
	suite.repository.EXPECT().GetBooks([]int{1, 2, 3}, []int{4, 5, 6}, 10, 100, 1900, 2000, 50).Return(nil, nil).Once()

	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, "GET", "/api/v1/books?authors=1,2,3&genres=4,5,6&min-pages=10&max-pages=100&min-year=1900&max-year=2000&limit=50", nil)
	suite.Require().NoError(err)
	resp := httptest.NewRecorder()
	suite.router.ServeHTTP(resp, req)

	suite.Equal(200, resp.Code)
	suite.Equal(`[]`, resp.Body.String())
}

func (suite *APIV1BooksTestSuite) TestInvalidParameters() {
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
			suite.Run(key+" = "+value, func() {
				ctx := context.Background()
				req, err := http.NewRequestWithContext(ctx, "GET", "/api/v1/books?"+key+"="+value, nil)
				suite.Require().NoError(err)

				resp := httptest.NewRecorder()
				suite.router.ServeHTTP(resp, req)
				suite.Equal(400, resp.Code)

				errorResponse := controller.ErrorResponse{}
				err = json.Unmarshal(resp.Body.Bytes(), &errorResponse)
				suite.Require().NoError(err)
				suite.NotNil(errorResponse.Message)
			})
		}
	}
}
