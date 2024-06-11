package repository

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"readcommend/internal/repository/model"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	gormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type RepositoryTestSuite struct {
	db         *gorm.DB
	repository BookRepository
	suite.Suite
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (suite *RepositoryTestSuite) SetupSuite() {
	suite.db = createContainerDB(&suite.Suite)
	seed(suite.db)
	suite.repository = NewBookRepository(suite.db)
}

func (suite *RepositoryTestSuite) TestGetAllBooksRanked() {
	books, err := suite.repository.GetBooks(nil, nil, 0, 0, 0, 0, 0)
	suite.Require().NoError(err)
	suite.Len(books, 58)
	suite.Equal(model.Book{
		ID:            12,
		Title:         "Can I Be Honest?",
		YearPublished: 2007,
		Rating:        4.77,
		Pages:         542,
		GenreID:       1,
		AuthorID:      11,
		Genre: model.Genre{
			ID:    1,
			Title: "Young Adult",
		},
		Author: model.Author{
			ID:        11,
			FirstName: "Charles",
			LastName:  "Fenimore",
		},
	}, books[0])
}

func (suite *RepositoryTestSuite) TestGetBooksSingleConditional() {
	suite.Run("get books of author #1", func() {
		books, err := suite.repository.GetBooks([]int{1}, nil, 0, 0, 0, 0, 0)
		suite.Require().NoError(err)
		suite.Len(books, 1)
	})

	suite.Run("get books of author #4", func() {
		books, err := suite.repository.GetBooks([]int{4}, nil, 0, 0, 0, 0, 0)
		suite.Require().NoError(err)
		suite.Len(books, 2)
	})

	suite.Run("get books of author #1 and #4", func() {
		books, err := suite.repository.GetBooks([]int{1, 4}, nil, 0, 0, 0, 0, 0)
		suite.Require().NoError(err)
		suite.Len(books, 3)
	})

	suite.Run("get books of genre #1 and #2", func() {
		books, err := suite.repository.GetBooks(nil, []int{1, 2}, 0, 0, 0, 0, 0)
		suite.Require().NoError(err)
		suite.Len(books, 6+11)
	})

	suite.Run("get books of min-page 1000", func() {
		books, err := suite.repository.GetBooks(nil, nil, 875, 0, 0, 0, 0)
		suite.Require().NoError(err)
		suite.Len(books, 2)
	})

	suite.Run("get books of max-page 10", func() {
		books, err := suite.repository.GetBooks(nil, nil, 0, 50, 0, 0, 0)
		suite.Require().NoError(err)
		suite.Len(books, 6)
	})

	suite.Run("get books of min-year 2020", func() {
		books, err := suite.repository.GetBooks(nil, nil, 0, 0, 2020, 0, 0)
		suite.Require().NoError(err)
		suite.Len(books, 2)
	})

	suite.Run("get books of max-year 10", func() {
		books, err := suite.repository.GetBooks(nil, nil, 0, 0, 0, 1930, 0)
		suite.Require().NoError(err)
		suite.Len(books, 1)
	})

	suite.Run("get a maximum of 5 books of all available", func() {
		books, err := suite.repository.GetBooks(nil, nil, 0, 0, 0, 0, 5)
		suite.Require().NoError(err)
		suite.Len(books, 5)
	})
}

func (suite *RepositoryTestSuite) TestGetAllBooksMultiConditional() {
	suite.Run("get books of books of author #1 and #4 intersect with genre #1 and #2", func() {
		books, err := suite.repository.GetBooks([]int{1, 4}, []int{1, 2}, 0, 0, 0, 0, 0)
		suite.Require().NoError(err)
		suite.Len(books, 1)
	})

	suite.Run("get books of min-year 2000 limited to 5", func() {
		books, err := suite.repository.GetBooks(nil, nil, 0, 0, 2000, 0, 5)
		suite.Require().NoError(err)
		suite.Len(books, 5)
		for _, book := range books {
			suite.GreaterOrEqual(book.YearPublished, 2000)
		}
	})
}

func (suite *RepositoryTestSuite) TestGetAuthors() {
	suite.Run("get all authors", func() {
		authors, err := suite.repository.GetAuthors("", 0)
		suite.Require().NoError(err)
		suite.Len(authors, 41)
	})
	suite.Run("find 'Robert Plimpton' and 'Robert Milofsky' by searching 'rob'", func() {
		authors, err := suite.repository.GetAuthors("rob", 0)
		suite.Require().NoError(err)
		suite.Len(authors, 2)
	})

	suite.Run("find one of 'Robert Plimpton' and 'Robert Milofsky' by searching 'rob' and limiting results to 1", func() {
		authors, err := suite.repository.GetAuthors("rob", 1)
		suite.Require().NoError(err)
		suite.Len(authors, 1)
	})

	suite.Run("find 'Robert Plimpton' by searching 'rob pli'", func() {
		authors, err := suite.repository.GetAuthors("rob pli", 0)
		suite.Require().NoError(err)
		suite.Len(authors, 1)
	})

	suite.Run("find 'Robert Plimpton' by searching 'pli rob'", func() {
		authors, err := suite.repository.GetAuthors("pli rob", 0)
		suite.Require().NoError(err)
		suite.Len(authors, 1)
	})

	suite.Run("find 'Robert Plimpton' by searching 'pli'", func() {
		authors, err := suite.repository.GetAuthors("pli", 0)
		suite.Require().NoError(err)
		suite.Len(authors, 1)
	})
}

func (suite *RepositoryTestSuite) TestGetGenres() {
	genres, err := suite.repository.GetGenres()
	suite.Require().NoError(err)
	suite.Len(genres, 8)
}

func (suite *RepositoryTestSuite) TestGetSizes() {
	sizes, err := suite.repository.GetSizes()
	suite.Require().NoError(err)
	suite.Len(sizes, 7)
}

func (suite *RepositoryTestSuite) TestGetEras() {
	eras, err := suite.repository.GetEras()
	suite.Require().NoError(err)
	suite.Len(eras, 3)
}

func createContainerDB(suite *suite.Suite) *gorm.DB {
	ctx := context.Background()
	containerDb, err := postgres.RunContainer(ctx,
		testcontainers.WithLogger(testcontainers.TestLogger(suite.T())),
		testcontainers.WithWaitStrategy(
			wait.ForExposedPort(),
		),
	)
	suite.Require().NoError(err)

	suite.T().Cleanup(func() {
		if err := containerDb.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}
	})

	connString, err := containerDb.ConnectionString(ctx)
	suite.Require().NoError(err)

	config := gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}}
	db, err := gorm.Open(gormPostgres.Open(connString), &config)
	suite.Require().NoError(err)

	err = db.AutoMigrate(model.Author{}, model.Era{}, model.Genre{}, model.Size{}, model.Book{})
	suite.Require().NoError(err)
	return db
}

//nolint:golint,funlen
func seed(db *gorm.DB) {
	const batchSize = 100

	db.CreateInBatches([]model.Era{
		{ID: 2, Title: "Classic", MaxYear: sql.NullInt16{Int16: 1969, Valid: true}},
		{ID: 1, Title: "Any"},
		{ID: 3, Title: "Modern", MinYear: sql.NullInt16{Int16: 1970, Valid: true}},
	}, batchSize)

	db.CreateInBatches([]model.Size{
		{ID: 1, Title: "Any"},
		{ID: 2, Title: "Short story – up to 35 pages", MaxPages: sql.NullInt16{Int16: 34, Valid: true}},
		{ID: 3, Title: "Novelette – 35 to 85 pages", MinPages: sql.NullInt16{Int16: 35, Valid: true}, MaxPages: sql.NullInt16{Int16: 84, Valid: true}},
		{ID: 4, Title: "Novella – 85 to 200 pages", MinPages: sql.NullInt16{Int16: 85, Valid: true}, MaxPages: sql.NullInt16{Int16: 199, Valid: true}},
		{ID: 5, Title: "Novel – 200 to 500 pages", MinPages: sql.NullInt16{Int16: 200, Valid: true}, MaxPages: sql.NullInt16{Int16: 499, Valid: true}},
		{ID: 6, Title: "Brick – 500 to 800 pages", MinPages: sql.NullInt16{Int16: 500, Valid: true}, MaxPages: sql.NullInt16{Int16: 799, Valid: true}},
		{ID: 7, Title: "Monument – 800 pages and up", MinPages: sql.NullInt16{Int16: 800, Valid: true}},
	}, batchSize)

	db.CreateInBatches([]model.Genre{
		{ID: 1, Title: "Young Adult"},
		{ID: 2, Title: "SciFi/Fantasy"},
		{ID: 3, Title: "Romance"},
		{ID: 4, Title: "Nonfiction"},
		{ID: 5, Title: "Mystery"},
		{ID: 6, Title: "Memoir"},
		{ID: 7, Title: "Fiction"},
		{ID: 8, Title: "Childrens"},
	}, batchSize)

	db.CreateInBatches([]model.Author{
		{ID: 1, FirstName: "Wendell", LastName: "Stackhouse"},
		{ID: 2, FirstName: "Amelia", LastName: "Wangerin, Jr."},
		{ID: 3, FirstName: "Anastasia", LastName: "Inez"},
		{ID: 4, FirstName: "Arthur", LastName: "McCrumb"},
		{ID: 5, FirstName: "Arturo", LastName: "Hijuelos"},
		{ID: 6, FirstName: "Bernard", LastName: "Hopf"},
		{ID: 7, FirstName: "Bianca", LastName: "Thompson"},
		{ID: 8, FirstName: "Bravig", LastName: "Lewisohn"},
		{ID: 9, FirstName: "Burton", LastName: "Malamud"},
		{ID: 10, FirstName: "Carolyn", LastName: "Segal"},
		{ID: 11, FirstName: "Charles", LastName: "Fenimore"},
		{ID: 12, FirstName: "Clifford", LastName: "Wolitzer"},
		{ID: 13, FirstName: "Darryl", LastName: "Fleischman"},
		{ID: 14, FirstName: "David", LastName: "Beam"},
		{ID: 15, FirstName: "Elizabeth", LastName: "Herbach"},
		{ID: 16, FirstName: "Elmer", LastName: "Komroff"},
		{ID: 17, FirstName: "Gloria", LastName: "Green"},
		{ID: 18, FirstName: "Grace", LastName: "Harrison"},
		{ID: 19, FirstName: "Hamlin", LastName: "Myrer"},
		{ID: 20, FirstName: "Hillary", LastName: "Barnhardt"},
		{ID: 21, FirstName: "Jill", LastName: "Hergesheimer"},
		{ID: 22, FirstName: "John W.", LastName: "Spanogle"},
		{ID: 23, FirstName: "Jonathan", LastName: "Kotzwinkle"},
		{ID: 24, FirstName: "Kathy", LastName: "Yglesias"},
		{ID: 25, FirstName: "Kenneth", LastName: "Douglas"},
		{ID: 26, FirstName: "Kris", LastName: "Elegant"},
		{ID: 27, FirstName: "Langston", LastName: "Lippman"},
		{ID: 28, FirstName: "Leonard", LastName: "Nabokov"},
		{ID: 29, FirstName: "Lori", LastName: "Kaan"},
		{ID: 30, FirstName: "Lynne", LastName: "Danticat"},
		{ID: 31, FirstName: "Malin", LastName: "Wolff"},
		{ID: 32, FirstName: "Oliver", LastName: "Lowry"},
		{ID: 33, FirstName: "Patricia", LastName: "Hazzard"},
		{ID: 34, FirstName: "Philip", LastName: "Antrim"},
		{ID: 35, FirstName: "Phoebe", LastName: "Brown"},
		{ID: 36, FirstName: "R.M.", LastName: "Larner"},
		{ID: 37, FirstName: "Robert", LastName: "Plimpton"},
		{ID: 38, FirstName: "Robert", LastName: "Milofsky"},
		{ID: 39, FirstName: "Ursula", LastName: "Karénine"},
		{ID: 40, FirstName: "Ward", LastName: "Haigh"},
		{ID: 41, FirstName: "Abraham", LastName: "Barton"},
	}, batchSize)

	db.CreateInBatches([]model.Book{
		{ID: 1, Title: "Alanna Saves the Day", YearPublished: 1972, Rating: 1.62, Pages: 169, GenreID: 8, AuthorID: 6},
		{ID: 2, Title: "Adventures of Kaya", YearPublished: 1999, Rating: 2.13, Pages: 619, GenreID: 1, AuthorID: 40},
		{ID: 3, Title: "A Horrible Human with the Habits of a Monster", YearPublished: 1976, Rating: 1.14, Pages: 258, GenreID: 7, AuthorID: 25},
		{ID: 4, Title: "And I Said Yes", YearPublished: 1954, Rating: 3.3, Pages: 183, GenreID: 7, AuthorID: 16},
		{ID: 5, Title: "Ballinby Boys", YearPublished: 1960, Rating: 1.88, Pages: 205, GenreID: 2, AuthorID: 4},
		{ID: 6, Title: "Banana Slug and the Lost Cow", YearPublished: 1983, Rating: 2.53, Pages: 527, GenreID: 8, AuthorID: 20},
		{ID: 7, Title: "Banana Slug and Xyr Friends", YearPublished: 1989, Rating: 3.64, Pages: 558, GenreID: 8, AuthorID: 20},
		{ID: 8, Title: "Banana Slug and the Glass Half Full", YearPublished: 1952, Rating: 4.51, Pages: 796, GenreID: 8, AuthorID: 17},
		{ID: 9, Title: "Banana Slug and the Mossy Rock", YearPublished: 2006, Rating: 4.43, Pages: 70, GenreID: 8, AuthorID: 31},
		{ID: 10, Title: "Burnished Silver", YearPublished: 1932, Rating: 1.2, Pages: 202, GenreID: 3, AuthorID: 30},
		{ID: 11, Title: "Cimornul", YearPublished: 1942, Rating: 1.08, Pages: 791, GenreID: 2, AuthorID: 21},
		{ID: 12, Title: "Can I Be Honest?", YearPublished: 2007, Rating: 4.77, Pages: 542, GenreID: 1, AuthorID: 11},
		{ID: 13, Title: "Concerning Prophecy", YearPublished: 1944, Rating: 3.8, Pages: 155, GenreID: 2, AuthorID: 18},
		{ID: 14, Title: "Don't Check your Ego", YearPublished: 1993, Rating: 3.02, Pages: 100, GenreID: 4, AuthorID: 36},
		{ID: 15, Title: "The Deep Grey", YearPublished: 1931, Rating: 3.94, Pages: 43, GenreID: 7, AuthorID: 37},
		{ID: 16, Title: "Dust on the Rim", YearPublished: 1946, Rating: 4.24, Pages: 38, GenreID: 2, AuthorID: 24},
		{ID: 17, Title: "Did You Hear?", YearPublished: 1954, Rating: 2.48, Pages: 887, GenreID: 7, AuthorID: 30},
		{ID: 18, Title: "Heliotrope Pajamas", YearPublished: 1952, Rating: 3.74, Pages: 16, GenreID: 8, AuthorID: 31},
		{ID: 19, Title: "Hashtag QuokkaSelfie", YearPublished: 1995, Rating: 3.42, Pages: 417, GenreID: 4, AuthorID: 27},
		{ID: 20, Title: "Interrobangs for All", YearPublished: 2011, Rating: 3.37, Pages: 677, GenreID: 7, AuthorID: 16},
		{ID: 21, Title: "Inconvenient Confessions: a 6", YearPublished: 1972, Rating: 4.11, Pages: 766, GenreID: 6, AuthorID: 32},
		{ID: 22, Title: "It's Never Just a Glass", YearPublished: 1956, Rating: 3.55, Pages: 305, GenreID: 1, AuthorID: 28},
		{ID: 23, Title: "Kalakalal Avenue", YearPublished: 2016, Rating: 4.27, Pages: 26, GenreID: 7, AuthorID: 16},
		{ID: 24, Title: "Lace and Brandy", YearPublished: 1967, Rating: 4.13, Pages: 158, GenreID: 3, AuthorID: 30},
		{ID: 25, Title: "Land Water Sky Space", YearPublished: 1983, Rating: 1.64, Pages: 320, GenreID: 4, AuthorID: 15},
		{ID: 26, Title: "(im)Mortality", YearPublished: 1985, Rating: 1.72, Pages: 214, GenreID: 1, AuthorID: 12},
		{ID: 27, Title: "Muddy Waters", YearPublished: 2020, Rating: 4.76, Pages: 594, GenreID: 3, AuthorID: 30},
		{ID: 28, Title: "Not to Gossip, But", YearPublished: 1958, Rating: 3.96, Pages: 537, GenreID: 7, AuthorID: 17},
		{ID: 29, Title: "Nothing But Capers", YearPublished: 2004, Rating: 3.87, Pages: 347, GenreID: 4, AuthorID: 1},
		{ID: 30, Title: "No More Lightning", YearPublished: 1978, Rating: 3.16, Pages: 99, GenreID: 7, AuthorID: 11},
		{ID: 31, Title: "Natural Pamplemousse", YearPublished: 1957, Rating: 4.66, Pages: 886, GenreID: 4, AuthorID: 35},
		{ID: 32, Title: "9803 North Millworks Road", YearPublished: 1935, Rating: 4.76, Pages: 449, GenreID: 5, AuthorID: 10},
		{ID: 33, Title: "Post Alley", YearPublished: 2014, Rating: 1.63, Pages: 374, GenreID: 7, AuthorID: 9},
		{ID: 34, Title: "Portmeirion", YearPublished: 2020, Rating: 2.11, Pages: 277, GenreID: 2, AuthorID: 7},
		{ID: 35, Title: "Quiddity and Quoddity", YearPublished: 2005, Rating: 2.42, Pages: 318, GenreID: 1, AuthorID: 21},
		{ID: 36, Title: "Rystwyth", YearPublished: 1930, Rating: 1.6, Pages: 59, GenreID: 2, AuthorID: 7},
		{ID: 37, Title: "Saint Esme", YearPublished: 1949, Rating: 1.84, Pages: 196, GenreID: 3, AuthorID: 30},
		{ID: 38, Title: "Some Eggs or Something?", YearPublished: 1997, Rating: 3.24, Pages: 12, GenreID: 7, AuthorID: 29},
		{ID: 39, Title: "Say it with Snap!", YearPublished: 1989, Rating: 3.77, Pages: 499, GenreID: 4, AuthorID: 22},
		{ID: 40, Title: "Soft, Pliable Truth", YearPublished: 1933, Rating: 3.28, Pages: 453, GenreID: 2, AuthorID: 38},
		{ID: 41, Title: "She Also Tottered", YearPublished: 2010, Rating: 2.09, Pages: 225, GenreID: 2, AuthorID: 38},
		{ID: 42, Title: "The Spark and The Ashes", YearPublished: 2000, Rating: 2.71, Pages: 721, GenreID: 1, AuthorID: 39},
		{ID: 43, Title: "Thatchwork Cottage", YearPublished: 1986, Rating: 2.43, Pages: 667, GenreID: 7, AuthorID: 9},
		{ID: 44, Title: "Tales of the Compass", YearPublished: 1945, Rating: 4.22, Pages: 570, GenreID: 2, AuthorID: 24},
		{ID: 45, Title: "The Elephant House", YearPublished: 1979, Rating: 3.95, Pages: 349, GenreID: 4, AuthorID: 22},
		{ID: 46, Title: "The Winchcombe Railway Museum Heist", YearPublished: 2004, Rating: 3.04, Pages: 731, GenreID: 5, AuthorID: 10},
		{ID: 47, Title: "The Startling End of Mr. Hidhoo", YearPublished: 1986, Rating: 1.59, Pages: 842, GenreID: 7, AuthorID: 23},
		{ID: 48, Title: "The Thing Is", YearPublished: 1988, Rating: 2.83, Pages: 115, GenreID: 7, AuthorID: 17},
		{ID: 49, Title: "The Mallemaroking", YearPublished: 1970, Rating: 1.95, Pages: 418, GenreID: 2, AuthorID: 7},
		{ID: 50, Title: "The Scent of Oranges", YearPublished: 2006, Rating: 2.37, Pages: 264, GenreID: 3, AuthorID: 30},
		{ID: 51, Title: "the life and times of an utterly inconsequential person", YearPublished: 1992, Rating: 1, Pages: 509, GenreID: 7, AuthorID: 4},
		{ID: 52, Title: "The Seawitch Sings", YearPublished: 1977, Rating: 4.62, Pages: 90, GenreID: 3, AuthorID: 30},
		{ID: 53, Title: "Turn Left Til You Get There", YearPublished: 1985, Rating: 4.54, Pages: 331, GenreID: 7, AuthorID: 26},
		{ID: 54, Title: "The Triscanipt", YearPublished: 2018, Rating: 2.26, Pages: 16, GenreID: 2, AuthorID: 39},
		{ID: 55, Title: "Whither Thou Goest", YearPublished: 1963, Rating: 4.44, Pages: 146, GenreID: 3, AuthorID: 30},
		{ID: 56, Title: "Who Did You Think You Were Kidding?", YearPublished: 1986, Rating: 4.6, Pages: 867, GenreID: 6, AuthorID: 34},
		{ID: 57, Title: "We're Sisters and We Kinda Like Each Other", YearPublished: 1989, Rating: 4.71, Pages: 67, GenreID: 6, AuthorID: 33},
		{ID: 58, Title: "Zero over Twelve", YearPublished: 1981, Rating: 1.01, Pages: 287, GenreID: 5, AuthorID: 9},
	}, batchSize)
}
