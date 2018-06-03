package schemas

import (
	// "fmt"
	// "github.com/kr/pretty"

	"github.com/graphql-go/graphql"
)

var Schema, Err = graphql.NewSchema(graphql.SchemaConfig{
	Query: graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"entities": &graphql.Field{
				Type: EntityType,
				Args: graphql.FieldConfigArgument{
					FLD_ENTITY_ID:          entity_id,
					FLD_ENTITY_KEY:         entity_key,
					FLD_ENTITY_DESCRIPTION: entity_description,
					FLD_ENTITY_PARENT_ID:   entity_parent_id,
					FLD_ENTITY_NATURE:      entity_type,
					FLD_ENTITY_ORDER:       entity_order,
					FLD_ENTITY_UPDATED_AT:  entity_updated_at,
					FLD_ENTITY_DELETED_AT:  entity_deleted_at,
					FLD_ENTITY_CREATED_AT:  entity_created_at,
				},
				Resolve:     getEntity,
				Description: "Taxonomy is build basd on entity of different types",
			},
			"nature": &graphql.Field{
				Type: graphql.NewList(NatureType),
				Args: graphql.FieldConfigArgument{
					FLD_TYPE_ID:          arg_type_id,
					FLD_TYPE_NAME:        arg_type_name,
					FLD_TYPE_DESCRIPTION: arg_type_description,
					FLD_TYPE_DELETED_AT:  arg_type_deleted_at,
					FLD_PAGE:             arg_page,
					FLD_PER_PAGE:         arg_per_page,
				},
				Resolve:     getNatures,
				Description: "Type of entity",
			},
			"total": &graphql.Field{
				Type:        NatureTotal,
				Description: "total",
				Resolve: func(p graphql.ResolveParams) (res interface{}, err error) {
					res = 100
					return
				},
			},
		},
	}),
	Mutation: graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"entity": &graphql.Field{
				Type: EntityType,
				Args: graphql.FieldConfigArgument{
					FLD_ENTITY_ID:          entity_id,
					FLD_ENTITY_KEY:         entity_key,
					FLD_ENTITY_DESCRIPTION: entity_description,
					FLD_ENTITY_PARENT_ID:   entity_parent_id,
					FLD_ENTITY_NATURE:      entity_type,
					FLD_ENTITY_ORDER:       entity_order,
					FLD_ENTITY_DELETED_AT:  entity_deleted_at,
				},
				Resolve:     entityMutation,
				Description: "Update taxonomy entity",
			},
			"nature": &graphql.Field{
				Type: NatureType,
				Args: graphql.FieldConfigArgument{
					FLD_TYPE_ID:          arg_type_id,
					FLD_TYPE_NAME:        arg_type_name,
					FLD_TYPE_DESCRIPTION: arg_type_description,
					FLD_TYPE_DELETED_AT:  arg_type_deleted_at,
				},
				Resolve:     typeMutation,
				Description: "entity type manipulation",
			},
		},
	}),
})
