//
//  Practicing GraphQL
//
//  Copyright © 2020. All rights reserved.
//

package graphql

import (
	"github.com/moemoe89/practicing-graphql-golang/api/v1/api_struct/form"
	"github.com/moemoe89/practicing-graphql-golang/api/v1/api_struct/model"
	usr "github.com/moemoe89/practicing-graphql-golang/api/v1/user"

	"errors"
	"math"
	"strings"

	"github.com/graphql-go/graphql"
	"github.com/moemoe89/go-helpers"
	"github.com/rs/xid"
)

// Resolver represent the graphql resolver
type Resolver interface {
	List(params graphql.ResolveParams) (interface{}, error)
	Detail(params graphql.ResolveParams) (interface{}, error)

	Update(params graphql.ResolveParams) (interface{}, error)
	Create(params graphql.ResolveParams) (interface{}, error)
	Delete(params graphql.ResolveParams) (interface{}, error)
}

// UserList holds information of users list.
type UserList struct {
	List      []*model.UserModel
	Page      int
	PerPage   int
	TotalPage int
	TotalData int
}

type resolver struct {
	svc usr.Service
}

func NewResolver(svc usr.Service) Resolver {
	return &resolver{
		svc: svc,
	}
}

func (r resolver) List(params graphql.ResolveParams) (interface{}, error) {
	userModel := model.UserModel{}

	offset, perPage, showPage, err := helpers.PaginationSetter(params.Args["per_page"].(string), params.Args["page"].(string))
	if err != nil {
		return nil, err
	}

	orderBy := params.Args["order_by"].(string)
	orderBy = helpers.OrderByHandler(orderBy, "db", userModel)
	if len(orderBy) < 1 {
		orderBy = "created_at DESC"
	}

	where := "WHERE deleted_at IS NULL"
	filter := map[string]interface{}{}

	name := params.Args["name"].(string)
	if len(name) > 0 {
		where += " AND name LIKE :name"
		filter["name"] = "%" + name + "%"
	}

	email := params.Args["email"].(string)
	if len(email) > 0 {
		where += " AND email LIKE :email"
		filter["email"] = "%" + email + "%"
	}

	phone := params.Args["phone"].(string)
	if len(phone) > 0 {
		where += " AND phone LIKE :phone"
		filter["phone"] = "%" + phone + "%"
	}

	createdAtStart := params.Args["created_at_start"].(string)
	if len(createdAtStart) > 0 {
		where += " AND created_at >= :created_at_start"
		filter["created_at_start"] = createdAtStart
	}

	createdAtEnd := params.Args["created_at_end"].(string)
	if len(createdAtEnd) > 0 {
		where += " AND created_at <= :created_at_end"
		filter["created_at_end"] = createdAtEnd
	}

	filterCount := filter
	filter["limit"] = perPage
	filter["offset"] = offset

	selectField := model.UserSelectField
	filterField := params.Args["select_field"].(string)
	if len(filterField) > 0 {
		res := helpers.CheckInTag(userModel, filterField, "db")
		if len(res) > 0 {
			selectField = strings.Join(res, ",")
		}
	}

	users, count, _, err := r.svc.List(filter, filterCount, where, orderBy, selectField)
	if err != nil {
		return nil, err
	}

	totalPage := int(math.Ceil(float64(count) / float64(perPage)))

	resp := &UserList{}
	resp.Page = showPage
	resp.PerPage = perPage
	resp.TotalPage = totalPage
	resp.TotalData = count
	resp.List = users

	return resp, nil
}

func (r resolver) Detail(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(string)

	user, _, err := r.svc.Detail(id, "")
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r resolver) Update(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(string)

	req := &form.UserForm{
		Name:    params.Args["name"].(string),
		Phone:   params.Args["phone"].(string),
		Email:   params.Args["email"].(string),
		Address: params.Args["address"].(string),
	}

	errs := req.Validate()
	if len(errs) > 0 {
		return nil, errors.New(errs[0])
	}

	user, _, err := r.svc.Update(req, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r resolver) Create(params graphql.ResolveParams) (interface{}, error) {
	req := &form.UserForm{
		ID:      xid.New().String(),
		Name:    params.Args["name"].(string),
		Phone:   params.Args["phone"].(string),
		Email:   params.Args["email"].(string),
		Address: params.Args["address"].(string),
	}

	errs := req.Validate()
	if len(errs) > 0 {
		return nil, errors.New(errs[0])
	}

	user, _, err := r.svc.Create(req)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r resolver) Delete(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(string)

	_, err := r.svc.Delete(id)
	if err != nil {
		return nil, err
	}

	return id, nil
}
