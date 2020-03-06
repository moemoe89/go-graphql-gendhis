//
//  Practicing GraphQL
//
//  Copyright © 2020. All rights reserved.
//

package model

import (
	"time"
)

// UserSelectField represent the default selected column for user model
const UserSelectField = "id,name,email,phone,address,created_at,updated_at"

// UserModel represent the user model
type UserModel struct {
	ID        string     `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	Email     string     `json:"email" db:"email"`
	Phone     string     `json:"phone" db:"phone"`
	Address   string     `json:"address" db:"address"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"-" db:"deleted_at"`
}
