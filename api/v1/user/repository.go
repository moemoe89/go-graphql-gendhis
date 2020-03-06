//
//  Practicing GraphQL
//
//  Copyright © 2020. All rights reserved.
//

package user

import (
	"github.com/moemoe89/practicing-graphql-golang/api/v1/api_struct/model"

	"fmt"

	"github.com/jmoiron/sqlx"
)

// Repository represent the repositories
type Repository interface {
	Get(filter map[string]interface{}, where, orderBy, selectField string) ([]*model.UserModel, error)
	Count(filter map[string]interface{}, where string) (int, error)
	Create(userReq *model.UserModel) (*model.UserModel, error)
	GetByID(id, selectField string) (*model.UserModel, error)
	Update(userReq *model.UserModel) (*model.UserModel, error)
	Delete(id string) error
}

type mysqlRepository struct {
	DBRead  *sqlx.DB
	DBWrite *sqlx.DB
}

// NewMysqlRepository will create an object that represent the Repository interface
func NewMysqlRepository(DBRead *sqlx.DB, DBWrite *sqlx.DB) Repository {
	return &mysqlRepository{DBRead, DBWrite}
}

func (m *mysqlRepository) Get(filter map[string]interface{}, where, orderBy, selectField string) ([]*model.UserModel, error) {
	users := []*model.UserModel{}
	if len(selectField) == 0 {
		selectField = model.UserSelectField
	}
	query := fmt.Sprintf("SELECT %s FROM users %s ORDER BY %s LIMIT :limit OFFSET :offset", selectField, where, orderBy)
	namedQuery, args, _ := m.DBRead.BindNamed(query, filter)
	err := m.DBRead.Select(&users, namedQuery, args...)
	return users, err
}

func (m *mysqlRepository) Count(filter map[string]interface{}, where string) (int, error) {
	var count int
	query := fmt.Sprintf("SELECT COUNT(id) FROM users %s", where)
	namedQuery, args, _ := m.DBRead.BindNamed(query, filter)
	err := m.DBRead.Get(&count, namedQuery, args...)
	return count, err
}

func (m *mysqlRepository) Create(user *model.UserModel) (*model.UserModel, error) {
	_, err := m.DBWrite.NamedExec(`INSERT INTO users (id, name, email, phone, address, created_at, updated_at) VALUES (:id, :name, :email, :phone, :address, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`, user)
	return user, err
}

func (m *mysqlRepository) GetByID(id, selectField string) (*model.UserModel, error) {
	where := "WHERE deleted_at IS NULL"
	filter := map[string]interface{}{}

	where += " AND id = :id"
	filter["id"] = id

	user := &model.UserModel{}
	if len(selectField) == 0 {
		selectField = model.UserSelectField
	}
	query := fmt.Sprintf("SELECT %s FROM users %s", selectField, where)
	namedQuery, args, _ := m.DBRead.BindNamed(query, filter)
	err := m.DBRead.Get(user, namedQuery, args...)
	return user, err
}

func (m *mysqlRepository) Update(user *model.UserModel) (*model.UserModel, error) {
	_, err := m.DBWrite.NamedExec(`UPDATE users SET name = :name, email = :email, phone = :phone, address = :address, updated_at = CURRENT_TIMESTAMP WHERE id = :id`, user)
	return user, err
}

func (m *mysqlRepository) Delete(id string) error {
	user := &model.UserModel{
		ID: id,
	}
	_, err := m.DBWrite.NamedExec(`UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE id = :id`, user)
	return err
}
