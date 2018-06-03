package schemas

import (
	"github.com/avachen2005/taxonomy_go/model/v1/entity"
	"github.com/avachen2005/taxonomy_go/model/v1/type"
	"github.com/graphql-go/graphql"
)

const (
	FLD_ENTITY_ID          = "id"
	FLD_ENTITY_KEY         = "key"
	FLD_ENTITY_DESCRIPTION = "description"
	FLD_ENTITY_PARENT_ID   = "parent_id"
	FLD_ENTITY_TYPE        = "type"
	FLD_ENTITY_ORDER       = "order"
	FLD_ENTITY_CREATED_AT  = "created_at"
	FLD_ENTITY_UPDATED_AT  = "updated_at"
	FLD_ENTITY_DELETED_AT  = "deleted"
)

var EntityType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "entity",
	Description: "",
	Fields: graphql.Fields{
		FLD_ENTITY_ID: &graphql.Field{
			Type:        graphql.Int,
			Description: "",
		},
		FLD_ENTITY_KEY: &graphql.Field{
			Type:        graphql.String,
			Description: "",
		},
		FLD_ENTITY_DESCRIPTION: &graphql.Field{
			Type:        graphql.String,
			Description: "",
		},
		FLD_ENTITY_PARENT_ID: &graphql.Field{
			Type:        graphql.Int,
			Description: "",
		},
		FLD_ENTITY_TYPE: &graphql.Field{
			Type:        graphql.String,
			Description: "",
			Resolve: func(p graphql.ResolveParams) (res interface{}, err error) {
				t := &model_v1_type.Type{}
				err, res = t.GetById(p.Source.(model_v1_entity.Entity).Type)
				return
			},
		},
		FLD_ENTITY_ORDER: &graphql.Field{
			Type:        graphql.Int,
			Description: "",
		},
		FLD_ENTITY_CREATED_AT: &graphql.Field{
			Type:        graphql.Int,
			Description: "",
		},
		FLD_ENTITY_UPDATED_AT: &graphql.Field{
			Type:        graphql.Int,
			Description: "",
		},
		FLD_ENTITY_DELETED_AT: &graphql.Field{
			Type:        graphql.Boolean,
			Description: "",
			Resolve: func(p graphql.ResolveParams) (res interface{}, err error) {

				if p.Source.(model_v1_entity.Entity).DeletedAt != 0 {
					res = true
				}

				res = false

				return
			},
		},
	},
})

var entity_id = &graphql.ArgumentConfig{
	Type:        graphql.Int,
	Description: "Primary key for the entity",
}

var entity_key = &graphql.ArgumentConfig{
	Type:        graphql.String,
	Description: "Search key field for the entity",
}

var entity_description = &graphql.ArgumentConfig{
	Type:        graphql.String,
	Description: "Description of the key",
}

var entity_parent_id = &graphql.ArgumentConfig{
	Type:        graphql.Int,
	Description: "Parent id",
}

var entity_type = &graphql.ArgumentConfig{
	Type:        graphql.Int,
	Description: "Type of the entity",
}

var entity_order = &graphql.ArgumentConfig{
	Type:        graphql.Int,
	Description: "order of the type",
}

var entity_updated_at = &graphql.ArgumentConfig{
	Type:        graphql.Int,
	Description: "last updated date",
}

var entity_deleted_at = &graphql.ArgumentConfig{
	Type:        graphql.Int,
	Description: "soft delete",
}

var entity_created_at = &graphql.ArgumentConfig{
	Type:        graphql.Int,
	Description: "created date",
}

func getEntity(params graphql.ResolveParams) (res interface{}, err error) {

	entity := &model_v1_entity.Entity{}

	if len(params.Args) == 0 {

		err, _, res = entity.GetAll(25, 1, true)
		return
	}

	stringFilter := map[string]string{}
	intFilter := map[string]int64{}

	if val, ok := params.Args[FLD_ENTITY_ID]; ok {
		intFilter[FLD_ENTITY_ID] = int64(val.(int))
	}

	if val, ok := params.Args[FLD_ENTITY_KEY]; ok {
		stringFilter[FLD_ENTITY_KEY] = val.(string)
	}

	if val, ok := params.Args[FLD_ENTITY_DESCRIPTION]; ok {
		stringFilter[FLD_ENTITY_DESCRIPTION] = val.(string)
	}

	if val, ok := params.Args[FLD_ENTITY_PARENT_ID]; ok && int64(val.(int)) > 0 {
		intFilter[FLD_ENTITY_PARENT_ID] = int64(val.(int))
	}

	if val, ok := params.Args[FLD_ENTITY_TYPE]; ok {
		stringFilter[FLD_ENTITY_TYPE] = val.(string)
	}

	if val, ok := params.Args[FLD_ENTITY_CREATED_AT]; ok && int64(val.(int)) > 0 {
		intFilter[FLD_ENTITY_CREATED_AT] = int64(val.(int))
	}

	if val, ok := params.Args[FLD_ENTITY_UPDATED_AT]; ok && int64(val.(int)) > 0 {
		intFilter[FLD_ENTITY_UPDATED_AT] = int64(val.(int))
	}

	if val, ok := params.Args[FLD_ENTITY_DELETED_AT]; ok && int64(val.(int)) > 0 {
		intFilter[FLD_ENTITY_DELETED_AT] = int64(val.(int))
	}

	err, _, res = entity.Search(stringFilter, intFilter)

	return

}

func entityMutation(params graphql.ResolveParams) (res interface{}, err error) {

	if _, ok := params.Args[FLD_ENTITY_ID]; ok {

		return updateEntity(params)

	} else {

		return createEntity(params)
	}

	return
}

func updateEntity(params graphql.ResolveParams) (res interface{}, err error) {

	intFilters := map[string]int64{}
	stringFilters := map[string]string{}

	for k, v := range params.Args {

		switch k {

		case FLD_ENTITY_ID:
		case FLD_ENTITY_ORDER:
		case FLD_ENTITY_DELETED_AT:
		case FLD_ENTITY_PARENT_ID:

			intFilters[k] = int64(v.(int))

		case FLD_ENTITY_KEY:
		case FLD_ENTITY_DESCRIPTION:

			stringFilters[k] = v.(string)

		case FLD_ENTITY_TYPE:

		}
	}

	newEntity := &model_v1_entity.Entity{}
	err, _ = newEntity.UpdateById(stringFilters, intFilters)
	err, _, res = newEntity.ReadById(newEntity.Id)

	return
}

func createEntity(params graphql.ResolveParams) (res interface{}, err error) {

	intFilters := map[string]int64{}
	stringFilters := map[string]string{}

	for k, v := range params.Args {

		switch k {

		case FLD_ENTITY_ID:
		case FLD_ENTITY_ORDER:
		case FLD_ENTITY_DELETED_AT:
		case FLD_ENTITY_PARENT_ID:

			intFilters[k] = int64(v.(int))

		case FLD_ENTITY_KEY:
		case FLD_ENTITY_DESCRIPTION:

			stringFilters[k] = v.(string)

		case FLD_ENTITY_TYPE:

		}
	}

	newEntity := &model_v1_entity.Entity{}
	_, err = newEntity.Create(stringFilters, intFilters)
	return
}
