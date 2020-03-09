//
//  Practicing GraphQL
//
//  Copyright © 2020. All rights reserved.
//

package graphql

import (
	"github.com/graphql-go/graphql"
)

// UserGraphQL holds user information with graphql object
var UserGraphQL = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"phone": &graphql.Field{
				Type: graphql.String,
			},
			"address": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"updated_at": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)

// UserListGraphQL holds user list information with graphql object
var UserListGraphQL = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UserList",
		Fields: graphql.Fields{
			"list": &graphql.Field{
				Type: graphql.NewList(UserGraphQL),
			},
			"page": &graphql.Field{
				Type: graphql.Int,
			},
			"per_page": &graphql.Field{
				Type: graphql.Int,
			},
			"total_page": &graphql.Field{
				Type: graphql.Int,
			},
			"total_data": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

// Schema is struct which has method for Query and Mutation. Please init this struct using constructor function.
type Schema struct {
	userResolver Resolver
}

// NewSchema initializes Schema struct which takes resolver as the argument.
func NewSchema(userResolver Resolver) Schema {
	return Schema{
		userResolver: userResolver,
	}
}

// Query initializes config schema query for graphql server.
func (s Schema) Query() *graphql.Object {
	objectConfig := graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"List": &graphql.Field{
				Type:        UserListGraphQL,
				Description: "Get list user",
				Args: graphql.FieldConfigArgument{
					"per_page": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"page": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"order_by": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"phone": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"created_at_start": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"created_at_end": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"select_field": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: s.userResolver.List,
			},
			"Detail": &graphql.Field{
				Type:        UserGraphQL,
				Description: "Get detail user",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: s.userResolver.Detail,
			},
		},
	}

	return graphql.NewObject(objectConfig)
}

// Mutation initializes config schema mutation for graphql server.
func (s Schema) Mutation() *graphql.Object {
	objectConfig := graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"Update": &graphql.Field{
				Type:        UserGraphQL,
				Description: "Update an user",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"phone": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"address": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: s.userResolver.Update,
			},
			"Create": &graphql.Field{
				Type:        UserGraphQL,
				Description: "Create a new iser",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"phone": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"address": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: s.userResolver.Create,
			},
			"Delete": &graphql.Field{
				Type:        graphql.String,
				Description: "Delete an user",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: s.userResolver.Delete,
			},
		},
	}

	return graphql.NewObject(objectConfig)
}
