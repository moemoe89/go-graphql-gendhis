//
//  Practicing GraphQL
//
//  Copyright © 2020. All rights reserved.
//

package form

import (
	"github.com/asaskevich/govalidator"
)

// UserForm represent the user request model
type UserForm struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

// Validate represent the validation method from UserForm
func (v *UserForm) Validate() []string {
	errs := []string{}
	if len(v.Name) < 1 {
		errs = append(errs, "Name can't be empty")
	}

	if !govalidator.IsEmail(v.Email) {
		errs = append(errs, "Invalid email address")
	}

	return errs
}

// UserQueryForm represent the user request model
type UserQueryForm struct {
	Query string `json:"query"`
}

// Validate represent the validation method from UserQueryForm
func (v *UserQueryForm) Validate() []string {
	errs := []string{}
	if len(v.Query) < 1 {
		errs = append(errs, "Query can't be empty")
	}

	return errs
}
