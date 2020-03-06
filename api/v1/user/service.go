//
//  Practicing GraphQL
//
//  Copyright © 2020. All rights reserved.
//

package user

import (
	"github.com/moemoe89/practicing-graphql-golang/api/v1/api_struct/form"
	"github.com/moemoe89/practicing-graphql-golang/api/v1/api_struct/model"

	"database/sql"
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Service represent the services
type Service interface {
	Create(req *form.UserForm) (*model.UserModel, int, error)
	Delete(id string) (int, error)
	Detail(id string, selectField string) (*model.UserModel, int, error)
	List(filter, filterCount map[string]interface{}, where, orderBy, selectField string) ([]*model.UserModel, int, int, error)
	Update(req *form.UserForm, id string) (*model.UserModel, int, error)
}

type implService struct {
	log        *logrus.Entry
	repository Repository
}

// NewService will create an object that represent the Service interface
func NewService(log *logrus.Entry, r Repository) Service {
	return &implService{log: log, repository: r}
}

func (u *implService) Create(req *form.UserForm) (*model.UserModel, int, error) {
	userReq := &model.UserModel{
		ID:      req.ID,
		Name:    req.Name,
		Email:   req.Email,
		Phone:   req.Phone,
		Address: req.Address,
	}

	user, err := u.repository.Create(userReq)
	if err != nil {
		u.log.Errorf("can't create user: %s", err.Error())
		return nil, http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	return user, 0, nil
}

func (u *implService) Delete(id string) (int, error) {

	_, err := u.repository.GetByID(id, "id")
	if err == sql.ErrNoRows {
		return http.StatusNotFound, errors.New("User not found")
	}

	if err != nil {
		u.log.Errorf("can't get user: %s with id %v", err.Error(), id)
		return http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	err = u.repository.Delete(id)
	if err != nil {
		u.log.Errorf("can't delete user: %s", err.Error())
		return http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	return 0, nil
}

func (u *implService) Detail(id string, selectField string) (*model.UserModel, int, error) {
	user, err := u.repository.GetByID(id, selectField)
	if err == sql.ErrNoRows {
		return nil, http.StatusNotFound, errors.New("User not found")
	}

	if err != nil {
		u.log.Errorf("can't get user: %s with id %v", err.Error(), id)
		return nil, http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	return user, 0, nil
}

func (u *implService) List(filter, filterCount map[string]interface{}, where, orderBy, selectField string) ([]*model.UserModel, int, int, error) {

	users, err := u.repository.Get(filter, where, orderBy, selectField)
	if err != nil {
		u.log.Errorf("can't get users: %s", err.Error())
		return nil, 0, http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	count, err := u.repository.Count(filterCount, where)
	if err != nil {
		u.log.Errorf("can't count users: %s", err.Error())
		return nil, 0, http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	return users, count, 0, nil
}

func (u *implService) Update(req *form.UserForm, id string) (*model.UserModel, int, error) {
	user := &model.UserModel{
		ID:      id,
		Name:    req.Name,
		Phone:   req.Phone,
		Email:   req.Email,
		Address: req.Address,
	}

	user, err := u.repository.Update(user)
	if err != nil {
		u.log.Errorf("can't update user: %s with id %v", err.Error(), id)
		return nil, http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	return user, 0, nil
}
