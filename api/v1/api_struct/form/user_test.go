//
//  Practicing GraphQL
//
//  Copyright © 2020. All rights reserved.
//

package form_test

import (
	"github.com/moemoe89/practicing-graphql-golang/api/v1/api_struct/form"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserNameEmpty(t *testing.T) {
	user := &form.UserForm{}

	expected := "Name can't be empty"
	errs := user.Validate()

	assert.Equal(t, expected, errs[0])
}

func TestUserInvalidEmail(t *testing.T) {
	user := &form.UserForm{
		Name:  "Momo",
		Email: "Baruno",
	}

	expected := "Invalid email address"
	errs := user.Validate()

	assert.Equal(t, expected, errs[0])
}
