//
//  Practicing GraphQL
//
//  Copyright © 2020. All rights reserved.
//

package user_test

import (
	"github.com/moemoe89/practicing-graphql-golang/api/v1/api_struct/form"
	"github.com/moemoe89/practicing-graphql-golang/api/v1/api_struct/model"
	"github.com/moemoe89/practicing-graphql-golang/api/v1/user"
	"github.com/moemoe89/practicing-graphql-golang/api/v1/user/mocks"
	"github.com/moemoe89/practicing-graphql-golang/config"

	"database/sql"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestServiceCreate(t *testing.T) {
	log := config.InitLog()
	mockRepo := new(mocks.Repository)

	reqUser := &form.UserForm{
		ID:      xid.New().String(),
		Name:    "Momo",
		Email:   "momo@mail.com",
		Phone:   "085640",
		Address: "Indonesia",
	}

	mockUser := &model.UserModel{
		ID:      reqUser.ID,
		Name:    reqUser.Name,
		Email:   reqUser.Email,
		Phone:   reqUser.Phone,
		Address: reqUser.Address,
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("Create", mockUser).Return(mockUser, nil).Once()
		u := user.NewService(log, mockRepo)

		userRow, status, err := u.Create(reqUser)

		assert.NoError(t, err)
		assert.NotNil(t, userRow)
		assert.Equal(t, 0, status)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		mockRepo.On("Create", mockUser).Return(nil, errors.New("Unexpected database error")).Once()
		u := user.NewService(log, mockRepo)

		userRow, status, err := u.Create(reqUser)

		assert.Error(t, err)
		assert.Nil(t, userRow)
		assert.Equal(t, http.StatusInternalServerError, status)

		mockRepo.AssertExpectations(t)
	})
}

func TestServiceDelete(t *testing.T) {
	log := config.InitLog()
	mockRepo := new(mocks.Repository)
	mockUser := &model.UserModel{
		ID:        xid.New().String(),
		Name:      "Momo",
		Email:     "momo@mail.com",
		Phone:     "085640",
		Address:   "Indonesia",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetByID", mock.AnythingOfType("string"), "id").Return(mockUser, nil).Once()
		mockRepo.On("Delete", mock.AnythingOfType("string")).Return(nil).Once()
		u := user.NewService(log, mockRepo)

		status, err := u.Delete(mockUser.ID)

		assert.NoError(t, err)
		assert.Equal(t, 0, status)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed-delete", func(t *testing.T) {
		mockRepo.On("GetByID", mock.AnythingOfType("string"), "id").Return(mockUser, nil).Once()
		mockRepo.On("Delete", mock.AnythingOfType("string")).Return(errors.New("Unexpected database error")).Once()
		u := user.NewService(log, mockRepo)

		status, err := u.Delete(mockUser.ID)

		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed-get-not-found", func(t *testing.T) {
		mockRepo.On("GetByID", mock.AnythingOfType("string"), "id").Return(nil, sql.ErrNoRows).Once()

		u := user.NewService(log, mockRepo)

		status, err := u.Delete(mockUser.ID)

		assert.Error(t, err)
		assert.Equal(t, http.StatusNotFound, status)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed-get", func(t *testing.T) {
		mockRepo.On("GetByID", mock.AnythingOfType("string"), "id").Return(nil, errors.New("Unexpected database error")).Once()
		u := user.NewService(log, mockRepo)

		status, err := u.Delete(mockUser.ID)

		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)

		mockRepo.AssertExpectations(t)
	})

}

func TestServiceDetail(t *testing.T) {
	log := config.InitLog()
	mockRepo := new(mocks.Repository)
	mockUser := &model.UserModel{
		ID:        xid.New().String(),
		Name:      "Momo",
		Email:     "momo@mail.com",
		Phone:     "085640",
		Address:   "Indonesia",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetByID", mock.AnythingOfType("string"), "").Return(mockUser, nil).Once()
		u := user.NewService(log, mockRepo)

		userRow, status, err := u.Detail(mockUser.ID, "")

		assert.NoError(t, err)
		assert.NotNil(t, userRow)
		assert.Equal(t, 0, status)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed-not-found", func(t *testing.T) {
		mockRepo.On("GetByID", mock.AnythingOfType("string"), "").Return(nil, sql.ErrNoRows).Once()
		u := user.NewService(log, mockRepo)

		userRow, status, err := u.Detail(mockUser.ID, "")

		assert.Error(t, err)
		assert.Nil(t, userRow)
		assert.Equal(t, http.StatusNotFound, status)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		mockRepo.On("GetByID", mock.AnythingOfType("string"), "").Return(nil, errors.New("Unexpected database error")).Once()
		u := user.NewService(log, mockRepo)

		userRow, status, err := u.Detail(mockUser.ID, "")

		assert.Error(t, err)
		assert.Nil(t, userRow)
		assert.Equal(t, http.StatusInternalServerError, status)

		mockRepo.AssertExpectations(t)
	})
}

func TestServiceList(t *testing.T) {
	log := config.InitLog()
	mockRepo := new(mocks.Repository)
	mockUser := &model.UserModel{
		ID:        xid.New().String(),
		Name:      "Momo",
		Email:     "momo@mail.com",
		Phone:     "085640",
		Address:   "Indonesia",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	mockListUser := make([]*model.UserModel, 0)
	mockListUser = append(mockListUser, mockUser)

	filter := map[string]interface{}{}
	filter["limit"] = "10"
	filter["offset"] = "0"
	orderBy := "id DESC"
	deletedNull := "WHERE deleted_at IS NULL"

	t.Run("success", func(t *testing.T) {
		mockRepo.On("Get", filter, deletedNull, orderBy, model.UserSelectField).Return(mockListUser, nil).Once()
		mockRepo.On("Count", filter, deletedNull).Return(1, nil).Once()
		u := user.NewService(log, mockRepo)

		users, count, status, err := u.List(filter, filter, deletedNull, orderBy, model.UserSelectField)

		assert.NoError(t, err)
		assert.NotNil(t, users)
		assert.Equal(t, 1, count)
		assert.Equal(t, 0, status)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed-get", func(t *testing.T) {
		mockRepo.On("Get", filter, deletedNull, orderBy, model.UserSelectField).Return(nil, errors.New("Unexpected database error")).Once()

		u := user.NewService(log, mockRepo)

		users, count, status, err := u.List(filter, filter, deletedNull, orderBy, model.UserSelectField)

		assert.Error(t, err)
		assert.Nil(t, users)
		assert.Equal(t, 0, count)
		assert.Equal(t, http.StatusInternalServerError, status)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed-count", func(t *testing.T) {
		mockRepo.On("Get", filter, deletedNull, orderBy, model.UserSelectField).Return(mockListUser, nil).Once()
		mockRepo.On("Count", filter, deletedNull).Return(0, errors.New("Unexpected database error")).Once()

		u := user.NewService(log, mockRepo)

		users, count, status, err := u.List(filter, filter, deletedNull, orderBy, model.UserSelectField)

		assert.Error(t, err)
		assert.Nil(t, users)
		assert.Equal(t, 0, count)
		assert.Equal(t, http.StatusInternalServerError, status)

		mockRepo.AssertExpectations(t)
	})
}

func TestServiceUpdate(t *testing.T) {
	log := config.InitLog()
	mockRepo := new(mocks.Repository)

	reqUser := &form.UserForm{
		Name:    "Momo",
		Email:   "momo@mail.com",
		Phone:   "085640",
		Address: "Indonesia",
	}

	mockUser := &model.UserModel{
		ID:      xid.New().String(),
		Name:    reqUser.Name,
		Email:   reqUser.Email,
		Phone:   reqUser.Phone,
		Address: reqUser.Address,
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("Update", mockUser).Return(mockUser, nil).Once()
		u := user.NewService(log, mockRepo)

		userRow, status, err := u.Update(reqUser, mockUser.ID)

		assert.NoError(t, err)
		assert.NotNil(t, userRow)
		assert.Equal(t, 0, status)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		mockRepo.On("Update", mockUser).Return(nil, errors.New("Unexpected database error")).Once()
		u := user.NewService(log, mockRepo)

		userRow, status, err := u.Update(reqUser, mockUser.ID)

		assert.Error(t, err)
		assert.Nil(t, userRow)
		assert.Equal(t, http.StatusInternalServerError, status)

		mockRepo.AssertExpectations(t)
	})
}
