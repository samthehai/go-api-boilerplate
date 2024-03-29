// Code generated by MockGen. DO NOT EDIT.
// Source: favorite.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/samthehai/ml-backend-test-samthehai/internal/entity"
	repository "github.com/samthehai/ml-backend-test-samthehai/internal/movie/usecase/repository"
)

// MockFavoriteRepository is a mock of FavoriteRepository interface.
type MockFavoriteRepository struct {
	ctrl     *gomock.Controller
	recorder *MockFavoriteRepositoryMockRecorder
}

// MockFavoriteRepositoryMockRecorder is the mock recorder for MockFavoriteRepository.
type MockFavoriteRepositoryMockRecorder struct {
	mock *MockFavoriteRepository
}

// NewMockFavoriteRepository creates a new mock instance.
func NewMockFavoriteRepository(ctrl *gomock.Controller) *MockFavoriteRepository {
	mock := &MockFavoriteRepository{ctrl: ctrl}
	mock.recorder = &MockFavoriteRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFavoriteRepository) EXPECT() *MockFavoriteRepositoryMockRecorder {
	return m.recorder
}

// AddFavoriteMovie mocks base method.
func (m *MockFavoriteRepository) AddFavoriteMovie(ctx context.Context, args repository.AddFavoriteMovieParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFavoriteMovie", ctx, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFavoriteMovie indicates an expected call of AddFavoriteMovie.
func (mr *MockFavoriteRepositoryMockRecorder) AddFavoriteMovie(ctx, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFavoriteMovie", reflect.TypeOf((*MockFavoriteRepository)(nil).AddFavoriteMovie), ctx, args)
}

// CheckIsFavoriteMovie mocks base method.
func (m *MockFavoriteRepository) CheckIsFavoriteMovie(ctx context.Context, args repository.CheckIsFavoriteMovieParams) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckIsFavoriteMovie", ctx, args)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckIsFavoriteMovie indicates an expected call of CheckIsFavoriteMovie.
func (mr *MockFavoriteRepositoryMockRecorder) CheckIsFavoriteMovie(ctx, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckIsFavoriteMovie", reflect.TypeOf((*MockFavoriteRepository)(nil).CheckIsFavoriteMovie), ctx, args)
}

// FindFavoriteMoviesByUserID mocks base method.
func (m *MockFavoriteRepository) FindFavoriteMoviesByUserID(ctx context.Context, userID uint64) ([]*entity.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindFavoriteMoviesByUserID", ctx, userID)
	ret0, _ := ret[0].([]*entity.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindFavoriteMoviesByUserID indicates an expected call of FindFavoriteMoviesByUserID.
func (mr *MockFavoriteRepositoryMockRecorder) FindFavoriteMoviesByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindFavoriteMoviesByUserID", reflect.TypeOf((*MockFavoriteRepository)(nil).FindFavoriteMoviesByUserID), ctx, userID)
}
