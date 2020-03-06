//
//  Practicing GraphQL
//
//  Copyright © 2020. All rights reserved.
//

package user

import (
	"github.com/moemoe89/go-helpers"
	"github.com/moemoe89/practicing-graphql-golang/api/v1/api_struct/form"
	"github.com/moemoe89/practicing-graphql-golang/api/v1/api_struct/model"
	cons "github.com/moemoe89/practicing-graphql-golang/constant"

	"math"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/moemoe89/go-localization"
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
)

type userCtrl struct {
	lang *language.Config
	log  *logrus.Entry
	svc  Service
}

// NewUserCtrl will create an object that represent the userCtrl struct
func NewUserCtrl(lang *language.Config, log *logrus.Entry, svc Service) *userCtrl {
	return &userCtrl{lang, log, svc}
}

//
// @Summary User Create
// @Description create user data
// @Accept  json
// @Produce  json
// @Param Accept-Language header string false "language message response"
// @Param body body form.UserForm true "Request Payload"
// @Success 201 {object} model.UserResponse
// @Failure 400 {object} model.GenericResponse
// @Failure 500 {object} model.GenericResponse
// @Router /user [post]
func (u *userCtrl) Create(c *gin.Context) {
	l := c.Request.Header.Get("Accept-Language")
	resp := &model.UserResponse{}

	req := &form.UserForm{}
	if err := c.ShouldBindJSON(&req); err != nil {
		u.log.Errorf("can't get json body: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewGenericResponse(http.StatusBadRequest, cons.ERR, []string{u.lang.Lookup(l, "Oops! Something went wrong with your request")}))
		return
	}

	errs := req.Validate()
	if len(errs) > 0 {
		errsLocale := []string{}
		for _, e := range errs {
			errsLocale = append(errsLocale, u.lang.Lookup(l, e))
		}
		c.JSON(http.StatusBadRequest, model.NewGenericResponse(http.StatusBadRequest, cons.ERR, errsLocale))
		return
	}

	req.ID = xid.New().String()
	user, status, err := u.svc.Create(req)
	if err != nil {
		c.JSON(status, model.NewGenericResponse(status, cons.ERR, []string{u.lang.Lookup(l, err.Error())}))
		return
	}

	resp.GenericResponse = model.NewGenericResponse(http.StatusCreated, cons.OK, []string{u.lang.Lookup(l, "Created data successful")})
	resp.Data = user
	c.JSON(http.StatusCreated, resp)
}

//
// @Summary User Detail
// @Description get user data by ID
// @Produce  json
// @Param Accept-Language header string false "language message response"
// @Param id path string true "User ID"
// @Success 200 {object} model.UserResponse
// @Failure 404 {object} model.GenericResponse
// @Failure 500 {object} model.GenericResponse
// @Router /user/{id} [get]
func (u *userCtrl) Detail(c *gin.Context) {
	l := c.Request.Header.Get("Accept-Language")
	resp := &model.UserResponse{}

	id := c.Param("id")

	user, status, err := u.svc.Detail(id, model.UserSelectField)
	if err != nil {
		c.JSON(status, model.NewGenericResponse(status, cons.ERR, []string{u.lang.Lookup(l, err.Error())}))
		return
	}

	resp.GenericResponse = model.NewGenericResponse(http.StatusOK, cons.OK, []string{u.lang.Lookup(l, "OK")})
	resp.Data = user
	c.JSON(http.StatusOK, resp)
}

//
// @Summary User List
// @Description get list of user data
// @Produce  json
// @Param Accept-Language header string false "Language message response"
// @Param per_page query int false "Limit per page"
// @Param page query int false "Page number"
// @Param select_field query string false "Select field"
// @Param order_by query string false "Sort by"
// @Success 200 {object} model.UsersResponse
// @Failure 500 {object} model.GenericResponse
// @Router /user [get]
func (u *userCtrl) List(c *gin.Context) {
	l := c.Request.Header.Get("Accept-Language")
	resp := &model.UsersResponse{}
	pagination := &model.UserPaginationResponse{}

	userModel := model.UserModel{}

	offset, perPage, showPage, err := helpers.PaginationSetter(c.Query("per_page"), c.Query("page"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewGenericResponse(http.StatusInternalServerError, cons.ERR, []string{u.lang.Lookup(l, err.Error())}))
		return
	}

	orderBy := c.Query("order_by")
	orderBy = helpers.OrderByHandler(orderBy, "db", userModel)
	if len(orderBy) < 1 {
		orderBy = "created_at DESC"
	}

	where := "WHERE deleted_at IS NULL"
	filter := map[string]interface{}{}

	name := c.Query("name")
	if len(name) > 0 {
		where += " AND name LIKE :name"
		filter["name"] = "%" + name + "%"
	}

	email := c.Query("email")
	if len(email) > 0 {
		where += " AND email LIKE :email"
		filter["email"] = "%" + email + "%"
	}

	phone := c.Query("phone")
	if len(phone) > 0 {
		where += " AND phone LIKE :phone"
		filter["phone"] = "%" + phone + "%"
	}

	createdAtStart := c.Query("created_at_start")
	if len(createdAtStart) > 0 {
		where += " AND created_at >= :created_at_start"
		filter["created_at_start"] = createdAtStart
	}

	createdAtEnd := c.Query("created_at_end")
	if len(createdAtEnd) > 0 {
		where += " AND created_at <= :created_at_end"
		filter["created_at_end"] = createdAtEnd
	}

	filterCount := filter
	filter["limit"] = perPage
	filter["offset"] = offset

	selectField := model.UserSelectField
	filterField := c.Query("select_field")
	if len(filterField) > 0 {
		res := helpers.CheckInTag(userModel, filterField, "db")
		if len(res) > 0 {
			selectField = strings.Join(res, ",")
		}
	}

	users, count, status, err := u.svc.List(filter, filterCount, where, orderBy, selectField)
	if err != nil {
		c.JSON(status, model.NewGenericResponse(status, cons.ERR, []string{u.lang.Lookup(l, "Oops! Something went wrong. Please try again later")}))
		return
	}

	totalPage := int(math.Ceil(float64(count) / float64(perPage)))
	pagination.PaginationResponse = model.NewPaginationResponse(showPage, perPage, totalPage, count)
	pagination.List = users

	resp.GenericResponse = model.NewGenericResponse(http.StatusOK, cons.OK, []string{u.lang.Lookup(l, "OK")})
	resp.Data = pagination
	c.JSON(http.StatusOK, resp)
}

//
// @Summary User Update
// @Description update user data
// @Accept  json
// @Produce  json
// @Param Accept-Language header string false "language message response"
// @Param id path string true "User ID"
// @Param body body form.UserForm true "Request Payload"
// @Success 200 {object} model.UserResponse
// @Failure 400 {object} model.GenericResponse
// @Failure 500 {object} model.GenericResponse
// @Router /user/{id} [put]
func (u *userCtrl) Update(c *gin.Context) {
	l := c.Request.Header.Get("Accept-Language")
	resp := &model.UserResponse{}

	id := c.Param("id")

	req := &form.UserForm{}
	if err := c.ShouldBindJSON(&req); err != nil {
		u.log.Errorf("can't get json body: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewGenericResponse(http.StatusBadRequest, cons.ERR, []string{u.lang.Lookup(l, "Oops! Something went wrong with your request")}))
		return
	}

	errs := req.Validate()
	if len(errs) > 0 {
		errsLocale := []string{}
		for _, e := range errs {
			errsLocale = append(errsLocale, u.lang.Lookup(l, e))
		}
		c.JSON(http.StatusBadRequest, model.NewGenericResponse(http.StatusBadRequest, cons.ERR, errsLocale))
		return
	}

	user, status, err := u.svc.Detail(id, "id")
	if err != nil {
		c.JSON(status, model.NewGenericResponse(status, cons.ERR, []string{u.lang.Lookup(l, err.Error())}))
		return
	}

	user, status, err = u.svc.Update(req, id)
	if err != nil {
		c.JSON(status, model.NewGenericResponse(status, cons.ERR, []string{u.lang.Lookup(l, err.Error())}))
		return
	}

	resp.GenericResponse = model.NewGenericResponse(http.StatusOK, cons.OK, []string{u.lang.Lookup(l, "Updated data successful")})
	resp.Data = user
	c.JSON(http.StatusOK, resp)
}

//
// @Summary User Delete
// @Description delete user data by ID
// @Produce  json
// @Param id path string true "User ID"
// @Param Accept-Language header string false "language message response"
// @Success 200 {object} model.GenericResponse
// @Failure 404 {object} model.GenericResponse
// @Failure 500 {object} model.GenericResponse
// @Router /user/{id} [delete]
func (u *userCtrl) Delete(c *gin.Context) {
	l := c.Request.Header.Get("Accept-Language")

	id := c.Param("id")

	status, err := u.svc.Delete(id)
	if err != nil {
		c.JSON(status, model.NewGenericResponse(status, cons.ERR, []string{u.lang.Lookup(l, err.Error())}))
		return
	}

	c.JSON(http.StatusOK, model.NewGenericResponse(http.StatusOK, cons.OK, []string{u.lang.Lookup(l, "Deleted data successful")}))
}
