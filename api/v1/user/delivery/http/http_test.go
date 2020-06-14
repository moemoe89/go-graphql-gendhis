//
//  Practicing GraphQL
//
//  Copyright © 2020. All rights reserved.
//

package http_test

import (
	"github.com/moemoe89/go-graphql-gendhis/api/v1/api_struct/form"
	"github.com/moemoe89/go-graphql-gendhis/api/v1/api_struct/model"
	"github.com/moemoe89/go-graphql-gendhis/api/v1/user/mocks"
	"github.com/moemoe89/go-graphql-gendhis/config"
	"github.com/moemoe89/go-graphql-gendhis/routers"

	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
)

func TestDeliveryCreateFail(t *testing.T) {
	lang, _ := config.InitLang()
	log := config.InitLog()

	userForm := &form.UserForm{
		Name:    "Momo",
		Email:   "momo@mail.com",
		Phone:   "085640",
		Address: "Indonesia",
	}

	j, err := json.Marshal(userForm)
	assert.NoError(t, err)

	mockService := new(mocks.Service)
	mockService.On("Create", userForm).Return(nil, http.StatusInternalServerError, errors.New("Unexpected database error"))

	router := routers.GetRouter(lang, log, mockService)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/api/v1/user", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.NotNil(t, w.Body)
}

func TestDeliveryCreateFailValidation(t *testing.T) {
	lang, _ := config.InitLang()
	log := config.InitLog()

	userForm := &form.UserForm{
		Name:    "",
		Email:   "momo@mail.com",
		Phone:   "085640",
		Address: "Indonesia",
	}

	j, err := json.Marshal(userForm)
	assert.NoError(t, err)

	mockService := new(mocks.Service)

	router := routers.GetRouter(lang, log, mockService)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/api/v1/user", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NotNil(t, w.Body)
}

func TestDeliveryCreateFailBindJSON(t *testing.T) {
	lang, _ := config.InitLang()
	log := config.InitLog()

	mockService := new(mocks.Service)

	router := routers.GetRouter(lang, log, mockService)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/api/v1/user", strings.NewReader(""))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NotNil(t, w.Body)
}

func TestDeliveryUpdate(t *testing.T) {
	id := xid.New().String()

	lang, _ := config.InitLang()
	log := config.InitLog()

	userForm := &form.UserForm{
		ID:      xid.New().String(),
		Name:    "Momo",
		Email:   "momo@mail.com",
		Phone:   "085640",
		Address: "Indonesia",
	}
	user := &model.UserModel{
		ID:      userForm.ID,
		Name:    userForm.Name,
		Email:   userForm.Email,
		Phone:   userForm.Phone,
		Address: userForm.Address,
	}

	j, err := json.Marshal(userForm)
	assert.NoError(t, err)

	mockService := new(mocks.Service)
	mockService.On("Detail", id, "id").Return(user, 0, nil)
	mockService.On("Update", userForm, id).Return(user, 0, nil)

	router := routers.GetRouter(lang, log, mockService)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "/api/v1/user/"+id, strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)
}

func TestDeliveryUpdateFail(t *testing.T) {
	id := xid.New().String()

	lang, _ := config.InitLang()
	log := config.InitLog()

	userForm := &form.UserForm{
		ID:      xid.New().String(),
		Name:    "Momo",
		Email:   "momo@mail.com",
		Phone:   "085640",
		Address: "Indonesia",
	}
	user := &model.UserModel{
		ID:      userForm.ID,
		Name:    userForm.Name,
		Email:   userForm.Email,
		Phone:   userForm.Phone,
		Address: userForm.Address,
	}

	j, err := json.Marshal(userForm)
	assert.NoError(t, err)

	mockService := new(mocks.Service)
	mockService.On("Detail", id, "id").Return(user, 0, nil)
	mockService.On("Update", userForm, id).Return(nil, http.StatusInternalServerError, errors.New("Unexpected database error"))

	router := routers.GetRouter(lang, log, mockService)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "/api/v1/user/"+id, strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.NotNil(t, w.Body)
}

func TestDeliveryUpdateFailDetail(t *testing.T) {
	id := xid.New().String()

	lang, _ := config.InitLang()
	log := config.InitLog()

	userForm := &form.UserForm{
		ID:      xid.New().String(),
		Name:    "Momo",
		Email:   "momo@mail.com",
		Phone:   "085640",
		Address: "Indonesia",
	}

	j, err := json.Marshal(userForm)
	assert.NoError(t, err)

	mockService := new(mocks.Service)
	mockService.On("Detail", id, "id").Return(nil, http.StatusInternalServerError, errors.New("Unexpected database error"))

	router := routers.GetRouter(lang, log, mockService)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "/api/v1/user/"+id, strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.NotNil(t, w.Body)
}

func TestDeliveryUpdateFailValidation(t *testing.T) {
	id := xid.New().String()

	lang, _ := config.InitLang()
	log := config.InitLog()

	userForm := &form.UserForm{
		ID:      xid.New().String(),
		Name:    "",
		Email:   "momo@mail.com",
		Phone:   "085640",
		Address: "Indonesia",
	}

	j, err := json.Marshal(userForm)
	assert.NoError(t, err)

	mockService := new(mocks.Service)

	router := routers.GetRouter(lang, log, mockService)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "/api/v1/user/"+id, strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NotNil(t, w.Body)
}

func TestDeliveryUpdateFailBindJSON(t *testing.T) {
	id := xid.New().String()

	lang, _ := config.InitLang()
	log := config.InitLog()

	mockService := new(mocks.Service)

	router := routers.GetRouter(lang, log, mockService)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "/api/v1/user/"+id, strings.NewReader(""))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NotNil(t, w.Body)
}

func TestDeliveryList(t *testing.T) {
	lang, _ := config.InitLang()
	log := config.InitLog()

	user := &model.UserModel{
		ID:      xid.New().String(),
		Name:    "Momo",
		Email:   "momo@mail.com",
		Phone:   "085640",
		Address: "Indonesia",
	}
	users := []*model.UserModel{}
	users = append(users, user)

	name := "a"
	email := "b"
	phone := "0856"
	createdAtStart := "2020-01-01"
	createdAtEnd := "2020-01-31"

	filter := map[string]interface{}{}
	filter["name"] = "%" + name + "%"
	filter["email"] = "%" + email + "%"
	filter["phone"] = "%" + phone + "%"
	filter["created_at_start"] = createdAtStart
	filter["created_at_end"] = createdAtEnd
	filterCount := filter
	filter["limit"] = 10
	filter["offset"] = 0

	mockService := new(mocks.Service)
	mockService.On("List", filter, filterCount, "WHERE deleted_at IS NULL AND name LIKE :name AND email LIKE :email AND phone LIKE :phone AND created_at >= :created_at_start AND created_at <= :created_at_end", "created_at DESC", model.UserSelectField).Return(users, 1, 0, nil)

	router := routers.GetRouter(lang, log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/user?per_page=10&name="+name+"&email="+email+"&phone="+phone+"&created_at_start="+createdAtStart+"&created_at_end="+createdAtEnd, strings.NewReader(""))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)
}

func TestDeliveryListFail(t *testing.T) {
	lang, _ := config.InitLang()
	log := config.InitLog()

	filter := map[string]interface{}{}
	filterCount := filter
	filter["limit"] = 10
	filter["offset"] = 0

	mockService := new(mocks.Service)
	mockService.On("List", filter, filterCount, "WHERE deleted_at IS NULL", "created_at DESC", model.UserSelectField).Return(nil, 0, http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later"))

	router := routers.GetRouter(lang, log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/user?per_page=10", strings.NewReader(""))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.NotNil(t, w.Body)
}

func TestDeliveryListFailPagination(t *testing.T) {
	lang, _ := config.InitLang()
	log := config.InitLog()

	filter := map[string]interface{}{}
	filterCount := filter
	filter["limit"] = 10
	filter["offset"] = 0

	mockService := new(mocks.Service)
	mockService.On("List", filter, filterCount, "WHERE deleted_at IS NULL", "created_at DESC", model.UserSelectField).Return(nil, 0, http.StatusInternalServerError, errors.New("Invalid parameter per_page: not an int"))

	router := routers.GetRouter(lang, log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/user?per_page=a", strings.NewReader(""))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.NotNil(t, w.Body)
}

func TestDeliveryDetail(t *testing.T) {
	id := xid.New().String()

	lang, _ := config.InitLang()
	log := config.InitLog()

	user := &model.UserModel{
		ID:      xid.New().String(),
		Name:    "Momo",
		Email:   "momo@mail.com",
		Phone:   "085640",
		Address: "Indonesia",
	}

	mockService := new(mocks.Service)
	mockService.On("Detail", id, model.UserSelectField).Return(user, 0, nil)

	router := routers.GetRouter(lang, log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/user/"+id, strings.NewReader(""))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)
}

func TestDeliveryDetailFail(t *testing.T) {
	id := xid.New().String()

	lang, _ := config.InitLang()
	log := config.InitLog()
	mockService := new(mocks.Service)
	mockService.On("Detail", id, model.UserSelectField).Return(nil, http.StatusInternalServerError, errors.New("Unexpected database error"))

	router := routers.GetRouter(lang, log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/user/"+id, strings.NewReader(""))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestDeliveryDelete(t *testing.T) {
	id := xid.New().String()

	lang, _ := config.InitLang()
	log := config.InitLog()
	mockService := new(mocks.Service)
	mockService.On("Delete", id).Return(0, nil)

	router := routers.GetRouter(lang, log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/user/"+id, strings.NewReader(""))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeliveryDeleteFail(t *testing.T) {
	id := xid.New().String()

	lang, _ := config.InitLang()
	log := config.InitLog()
	mockService := new(mocks.Service)
	mockService.On("Delete", id).Return(http.StatusInternalServerError, errors.New("Unexpected database error"))

	router := routers.GetRouter(lang, log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/user/"+id, strings.NewReader(""))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
