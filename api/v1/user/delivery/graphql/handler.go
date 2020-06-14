//
//  Practicing GraphQL
//
//  Copyright © 2020. All rights reserved.
//

package graphql

import (
	"github.com/moemoe89/go-graphql-gendhis/api/v1/user"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// Handler initializes the graphql middleware.
func Handler(userSvc user.Service) gin.HandlerFunc {
	schema := NewSchema(NewResolver(userSvc))
	graphqlSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    schema.Query(),
		Mutation: schema.Mutation(),
	})
	if err != nil {
		panic(err)
	}

	h := handler.New(&handler.Config{
		Schema:   &graphqlSchema,
		GraphiQL: true,
		Pretty:   true,
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
