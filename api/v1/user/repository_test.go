//
//  Practicing GraphQL
//
//  Copyright © 2020. All rights reserved.
//

package user_test

import (
	"github.com/moemoe89/go-graphql-gendhis/api/v1/api_struct/model"
	"github.com/moemoe89/go-graphql-gendhis/api/v1/user"

	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "address", "created_at", "updated_at", "deleted_at"}).
		AddRow(xid.New().String(), "Momo", "momo@mail.com", "085640", "Indonesia", time.Now().UTC(), time.Now().UTC(), nil)

	query := "SELECT " + model.UserSelectField + " FROM users WHERE deleted_at IS NULL ORDER BY id ASC LIMIT \\? OFFSET \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	u := user.NewPostgresRepository(sqlxDB, sqlxDB)

	orderBy := "id ASC"
	where := "WHERE deleted_at IS NULL"
	filter := map[string]interface{}{}
	filter["limit"] = "10"
	filter["offset"] = "0"

	users, err := u.Get(filter, where, orderBy, "")

	assert.NotEmpty(t, users)
	assert.NoError(t, err)
	assert.Len(t, users, 1)
}

func TestCount(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	rows := sqlmock.NewRows([]string{"COUNT(id)"}).AddRow(1)

	query := "SELECT COUNT\\(id\\) FROM users WHERE deleted_at IS NULL"

	mock.ExpectQuery(query).WillReturnRows(rows)
	u := user.NewPostgresRepository(sqlxDB, sqlxDB)

	where := "WHERE deleted_at IS NULL"
	filter := map[string]interface{}{}

	count, err := u.Count(filter, where)

	assert.NoError(t, err)
	assert.Equal(t, count, 1)
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	req := &model.UserModel{
		ID:      xid.New().String(),
		Name:    "Momo",
		Email:   "momo@mail.com",
		Phone:   "085640",
		Address: "Indonesia",
	}

	query := "INSERT INTO users \\(id, name, email, phone, address, created_at, updated_at\\) VALUES \\(\\?, \\?, \\?, \\?, \\?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP"

	mock.ExpectExec(query).WithArgs(req.ID, req.Name, req.Email, req.Phone, req.Address).WillReturnResult(sqlmock.NewResult(0, 1))

	u := user.NewPostgresRepository(sqlxDB, sqlxDB)
	userRow, err := u.Create(req)

	assert.NoError(t, err)
	assert.Equal(t, req.ID, userRow.ID)
}

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	req := &model.UserModel{
		ID:        xid.New().String(),
		Name:      "Momo",
		Email:     "momo@mail.com",
		Phone:     "085640",
		Address:   "Indonesia",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "address", "created_at", "updated_at"}).
		AddRow(req.ID, req.Name, req.Email, req.Phone, req.Address, req.CreatedAt, req.UpdatedAt)

	query := "SELECT " + model.UserSelectField + " FROM users WHERE deleted_at IS NULL AND id = \\?"

	mock.ExpectQuery(query).WithArgs(req.ID).WillReturnRows(rows)
	u := user.NewPostgresRepository(sqlxDB, sqlxDB)

	userRow, err := u.GetByID(req.ID, "")

	assert.NoError(t, err)
	assert.NotNil(t, userRow)
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	req := &model.UserModel{
		ID:      xid.New().String(),
		Name:    "Momo",
		Email:   "momo@mail.com",
		Phone:   "085640",
		Address: "Indonesia",
	}

	query := "UPDATE users SET name = \\?, email = \\?, phone = \\?, address = \\?, updated_at = CURRENT_TIMESTAMP WHERE id = \\?"

	mock.ExpectExec(query).WithArgs(req.Name, req.Email, req.Phone, req.Address, req.ID).WillReturnResult(sqlmock.NewResult(0, 1))
	u := user.NewPostgresRepository(sqlxDB, sqlxDB)

	userRow, err := u.Update(req)

	assert.NoError(t, err)
	assert.Equal(t, req.ID, userRow.ID)
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	id := xid.New().String()

	query := "UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE id = \\?"

	mock.ExpectExec(query).WithArgs(id).WillReturnResult(sqlmock.NewResult(0, 1))
	u := user.NewPostgresRepository(sqlxDB, sqlxDB)

	err = u.Delete(id)

	assert.NoError(t, err)
}
