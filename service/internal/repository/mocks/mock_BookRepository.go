// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	model "readcommend/internal/repository/model"

	mock "github.com/stretchr/testify/mock"
)

// MockBookRepository is an autogenerated mock type for the BookRepository type
type MockBookRepository struct {
	mock.Mock
}

type MockBookRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBookRepository) EXPECT() *MockBookRepository_Expecter {
	return &MockBookRepository_Expecter{mock: &_m.Mock}
}

// GetAuthors provides a mock function with given fields:
func (_m *MockBookRepository) GetAuthors() ([]model.Author, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAuthors")
	}

	var r0 []model.Author
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]model.Author, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []model.Author); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Author)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBookRepository_GetAuthors_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAuthors'
type MockBookRepository_GetAuthors_Call struct {
	*mock.Call
}

// GetAuthors is a helper method to define mock.On call
func (_e *MockBookRepository_Expecter) GetAuthors() *MockBookRepository_GetAuthors_Call {
	return &MockBookRepository_GetAuthors_Call{Call: _e.mock.On("GetAuthors")}
}

func (_c *MockBookRepository_GetAuthors_Call) Run(run func()) *MockBookRepository_GetAuthors_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBookRepository_GetAuthors_Call) Return(_a0 []model.Author, _a1 error) *MockBookRepository_GetAuthors_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBookRepository_GetAuthors_Call) RunAndReturn(run func() ([]model.Author, error)) *MockBookRepository_GetAuthors_Call {
	_c.Call.Return(run)
	return _c
}

// GetBooks provides a mock function with given fields: authors, genres, minPages, maxPages, minYear, maxYear, limit
func (_m *MockBookRepository) GetBooks(authors []int, genres []int, minPages int, maxPages int, minYear int, maxYear int, limit int) ([]model.Book, error) {
	ret := _m.Called(authors, genres, minPages, maxPages, minYear, maxYear, limit)

	if len(ret) == 0 {
		panic("no return value specified for GetBooks")
	}

	var r0 []model.Book
	var r1 error
	if rf, ok := ret.Get(0).(func([]int, []int, int, int, int, int, int) ([]model.Book, error)); ok {
		return rf(authors, genres, minPages, maxPages, minYear, maxYear, limit)
	}
	if rf, ok := ret.Get(0).(func([]int, []int, int, int, int, int, int) []model.Book); ok {
		r0 = rf(authors, genres, minPages, maxPages, minYear, maxYear, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Book)
		}
	}

	if rf, ok := ret.Get(1).(func([]int, []int, int, int, int, int, int) error); ok {
		r1 = rf(authors, genres, minPages, maxPages, minYear, maxYear, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBookRepository_GetBooks_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBooks'
type MockBookRepository_GetBooks_Call struct {
	*mock.Call
}

// GetBooks is a helper method to define mock.On call
//   - authors []int
//   - genres []int
//   - minPages int
//   - maxPages int
//   - minYear int
//   - maxYear int
//   - limit int
func (_e *MockBookRepository_Expecter) GetBooks(authors interface{}, genres interface{}, minPages interface{}, maxPages interface{}, minYear interface{}, maxYear interface{}, limit interface{}) *MockBookRepository_GetBooks_Call {
	return &MockBookRepository_GetBooks_Call{Call: _e.mock.On("GetBooks", authors, genres, minPages, maxPages, minYear, maxYear, limit)}
}

func (_c *MockBookRepository_GetBooks_Call) Run(run func(authors []int, genres []int, minPages int, maxPages int, minYear int, maxYear int, limit int)) *MockBookRepository_GetBooks_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]int), args[1].([]int), args[2].(int), args[3].(int), args[4].(int), args[5].(int), args[6].(int))
	})
	return _c
}

func (_c *MockBookRepository_GetBooks_Call) Return(_a0 []model.Book, _a1 error) *MockBookRepository_GetBooks_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBookRepository_GetBooks_Call) RunAndReturn(run func([]int, []int, int, int, int, int, int) ([]model.Book, error)) *MockBookRepository_GetBooks_Call {
	_c.Call.Return(run)
	return _c
}

// GetEras provides a mock function with given fields:
func (_m *MockBookRepository) GetEras() ([]model.Era, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetEras")
	}

	var r0 []model.Era
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]model.Era, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []model.Era); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Era)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBookRepository_GetEras_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEras'
type MockBookRepository_GetEras_Call struct {
	*mock.Call
}

// GetEras is a helper method to define mock.On call
func (_e *MockBookRepository_Expecter) GetEras() *MockBookRepository_GetEras_Call {
	return &MockBookRepository_GetEras_Call{Call: _e.mock.On("GetEras")}
}

func (_c *MockBookRepository_GetEras_Call) Run(run func()) *MockBookRepository_GetEras_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBookRepository_GetEras_Call) Return(_a0 []model.Era, _a1 error) *MockBookRepository_GetEras_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBookRepository_GetEras_Call) RunAndReturn(run func() ([]model.Era, error)) *MockBookRepository_GetEras_Call {
	_c.Call.Return(run)
	return _c
}

// GetGenres provides a mock function with given fields:
func (_m *MockBookRepository) GetGenres() ([]model.Genre, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetGenres")
	}

	var r0 []model.Genre
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]model.Genre, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []model.Genre); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Genre)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBookRepository_GetGenres_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGenres'
type MockBookRepository_GetGenres_Call struct {
	*mock.Call
}

// GetGenres is a helper method to define mock.On call
func (_e *MockBookRepository_Expecter) GetGenres() *MockBookRepository_GetGenres_Call {
	return &MockBookRepository_GetGenres_Call{Call: _e.mock.On("GetGenres")}
}

func (_c *MockBookRepository_GetGenres_Call) Run(run func()) *MockBookRepository_GetGenres_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBookRepository_GetGenres_Call) Return(_a0 []model.Genre, _a1 error) *MockBookRepository_GetGenres_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBookRepository_GetGenres_Call) RunAndReturn(run func() ([]model.Genre, error)) *MockBookRepository_GetGenres_Call {
	_c.Call.Return(run)
	return _c
}

// GetSizes provides a mock function with given fields:
func (_m *MockBookRepository) GetSizes() ([]model.Size, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSizes")
	}

	var r0 []model.Size
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]model.Size, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []model.Size); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Size)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBookRepository_GetSizes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSizes'
type MockBookRepository_GetSizes_Call struct {
	*mock.Call
}

// GetSizes is a helper method to define mock.On call
func (_e *MockBookRepository_Expecter) GetSizes() *MockBookRepository_GetSizes_Call {
	return &MockBookRepository_GetSizes_Call{Call: _e.mock.On("GetSizes")}
}

func (_c *MockBookRepository_GetSizes_Call) Run(run func()) *MockBookRepository_GetSizes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBookRepository_GetSizes_Call) Return(_a0 []model.Size, _a1 error) *MockBookRepository_GetSizes_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBookRepository_GetSizes_Call) RunAndReturn(run func() ([]model.Size, error)) *MockBookRepository_GetSizes_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockBookRepository creates a new instance of MockBookRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBookRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBookRepository {
	mock := &MockBookRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
