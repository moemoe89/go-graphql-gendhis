//
//  Practicing GraphQL
//
//  Copyright © 2020. All rights reserved.
//

package model

// NewGenericResponse will create an object that represent the GenericResponse struct
func NewGenericResponse(stsCd, isError int, messages []string) *GenericResponse {

	return &GenericResponse{
		Status:   stsCd,
		Success:  isError == 0,
		Messages: messages,
	}
}

// NewPaginationResponse will create an object that represent the PaginationResponse struct
func NewPaginationResponse(page, perPage, totalPage, totalData int) *PaginationResponse {

	return &PaginationResponse{
		Page:      page,
		PerPage:   perPage,
		TotalPage: totalPage,
		TotalData: totalData,
	}
}

// GenericResponse represent the generic response API
type GenericResponse struct {
	Status   int      `json:"status"`
	Success  bool     `json:"success"`
	Messages []string `json:"messages"`
}

// PaginationResponse represent the response API with pagination
type PaginationResponse struct {
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	TotalPage int `json:"total_page"`
	TotalData int `json:"total_data"`
}

// UserResponse represent the generic user response API
type UserResponse struct {
	*GenericResponse
	Data *UserModel `json:"data"`
}

// UserPaginationResponse represent the user response API with pagination
type UserPaginationResponse struct {
	*PaginationResponse
	List []*UserModel `json:"list"`
}

// UsersResponse represent the generic user response API with pagination
type UsersResponse struct {
	*GenericResponse
	Data *UserPaginationResponse `json:"data"`
}
