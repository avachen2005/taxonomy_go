package schemas

import (
	"github.com/avachen2005/taxonomy_go/model/v1/entity"
	"github.com/avachen2005/taxonomy_go/model/v1/nature"
	"github.com/graphql-go/graphql"
)

const (
	FLD_ENTITY_ID          = "id"
	FLD_ENTITY_KEY         = "key"
	FLD_ENTITY_DESCRIPTION = "description"
	FLD_ENTITY_PARENT_ID   = "parent_id"
	FLD_ENTITY_NATURE      = "nature"
	FLD_ENTITY_ORDER       = "order"
	FLD_ENTITY_CREATED_AT  = "created_at"
	FLD_ENTITY_UPDATED_AT  = "updated_at"
	FLD_ENTITY_DELETED_AT  = "deleted"
)

type Entity struct {
	Id          int         `json:"id"`
	Key         string      `json:"key"`
	Description string      `json:"description"`
	ParentId    int         `json:"parent_id"`
	Nature      int         `json:"nature"`
	Order       int         `json:"order"`
	CreatedAt   int         `json:"created_at"`
	UpdatedAt   int         `json:"updated_at"`
	DeletedAt   int         `json:"deleted_at"`
	PageInfo    *Pagination `json: "page_info"`
}

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
		FLD_ENTITY_NATURE: &graphql.Field{
			Type:        graphql.String,
			Description: "",
			Resolve: func(p graphql.ResolveParams) (res interface{}, err error) {
				t := &model_v1_nature.Nature{}
				err, res = t.GetById(p.Source.(model_v1_entity.Entity).Nature)
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

func getEntity(p graphql.ResolveParams) (res interface{}, err error) {

	entity := &model_v1_entity.Entity{}
	newList := []Entity{}

	stringFilter := map[string]string{}
	intFilter := map[string]int64{}

	page := int64(default_page)
	if v, ok := p.Args[FLD_PAGE]; ok {
		page = int64(v.(int))

	}

	if page > 0 {
		page = page - 1
	}

	per_page := int64(default_per_page)
	if v, ok := p.Args[FLD_PER_PAGE]; ok {
		per_page = int64(v.(int))
	}

	for k, v := range p.Args {

		switch k {
		case FLD_ENTITY_ID, FLD_ENTITY_PARENT_ID, FLD_ENTITY_NATURE, FLD_ENTITY_CREATED_AT, FLD_ENTITY_UPDATED_AT:
			intFilter[k] = int64(v.(int))
		case FLD_ENTITY_KEY, FLD_ENTITY_DESCRIPTION:
			stringFilter[k] = v.(string)

		}
	}

	err, _, list := entity.Search(stringFilter, intFilter, per_page, page*per_page)

	for i, e := range list {
		newList = append(newList, Entity{
			Id:          int(e.Id),
			Key:         e.Key,
			Description: e.Description,
			ParentId:    int(e.ParentId),
			Nature:      int(e.Nature),
			Order:       int(e.Order),
			CreatedAt:   int(e.CreatedAt),
			UpdatedAt:   int(e.UpdatedAt),
			DeletedAt:   int(e.DeletedAt),
			PageInfo: &Pagination{
				Page:    int(page),
				PerPage: int(per_page),
				Cursor:  i,
			},
		})
	}

	res = newList
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

		case FLD_ENTITY_NATURE:

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

		case FLD_ENTITY_NATURE:

		}
	}

	newEntity := &model_v1_entity.Entity{}
	_, err = newEntity.Create(stringFilters, intFilters)
	return
}
