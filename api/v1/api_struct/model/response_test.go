//
//  Practicing GraphQL
//
//  Copyright © 2020. All rights reserved.
//

package model_test

import (
	"github.com/moemoe89/go-graphql-gendhis/api/v1/api_struct/model"

	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGenericResponse(t *testing.T) {
	status := http.StatusOK
	success := 0
	messages := []string{"OK"}

	resp := model.NewGenericResponse(status, success, messages)

	assert.Equal(t, status, resp.Status)
	assert.Equal(t, true, resp.Success)
	assert.Equal(t, messages[0], resp.Messages[0])
}

func TestNewPaginationResponse(t *testing.T) {
	page := 1
	perPage := 2
	totalPage := 3
	totalData := 4

	resp := model.NewPaginationResponse(page, perPage, totalPage, totalData)

	if page != resp.Page {
		t.Errorf("Should return %b, got %b", page, resp.Page)
	}

	if perPage != resp.PerPage {
		t.Errorf("Should return %b, got %b", perPage, resp.PerPage)
	}

	if totalPage != resp.TotalPage {
		t.Errorf("Should return %b, got %b", totalPage, resp.TotalPage)
	}

	if totalData != resp.TotalData {
		t.Errorf("Should return %b, got %b", totalData, resp.TotalData)
	}

	assert.Equal(t, page, resp.Page)
	assert.Equal(t, perPage, resp.PerPage)
	assert.Equal(t, totalPage, resp.TotalPage)
	assert.Equal(t, totalData, resp.TotalData)
}
