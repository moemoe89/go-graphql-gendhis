//
//  Practicing GraphQL
//
//  Copyright © 2020. All rights reserved.
//

package config_test

import (
	"github.com/moemoe89/go-graphql-gendhis/config"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLang(t *testing.T) {
	_, err := config.InitLang()

	assert.Nil(t, err)
}
