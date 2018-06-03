package schemas

import (
	"time"

	"github.com/avachen2005/taxonomy_go/model/v1/nature"
	"github.com/graphql-go/graphql"

	"fmt"
	"github.com/kr/pretty"
)

const (
	FLD_TYPE_ID          = "id"
	FLD_TYPE_NAME        = "name"
	FLD_TYPE_DESCRIPTION = "description"
	FLD_TYPE_CREATED_AT  = "created_at"
	FLD_TYPE_UPDATED_AT  = "updated_at"
	FLD_TYPE_DELETED_AT  = "deleted"
	FLD_TYPE_CURSOR      = "cursor"
)

type Nature struct {
	Id          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	CreatedAt   int         `json:"created_at"`
	UpdatedAt   int         `json:"updated_at"`
	DeletedAt   int         `json:"deleted_at"`
	PageInfo    *Pagination `json: "page_info"`
}

var NatureType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "nature",
	Description: "",
	Fields: graphql.Fields{
		FLD_TYPE_ID: &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (res interface{}, err error) {
				res = p.Source.(Nature).Id
				return
			},
		},
		FLD_TYPE_NAME: &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (res interface{}, err error) {
				res = p.Source.(Nature).Name
				return
			},
		},
		FLD_TYPE_DESCRIPTION: &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (res interface{}, err error) {
				res = p.Source.(Nature).Description
				return
			},
		},
		FLD_PAGE_INFO: &graphql.Field{
			Type: PaginationType,
			Resolve: func(p graphql.ResolveParams) (res interface{}, err error) {

				fmt.Println("== PaginationType ==")
				fmt.Printf("%#v\n", pretty.Formatter(p.Source))
				res = p.Source.(Nature).PageInfo
				return
			},
		},
		FLD_TYPE_UPDATED_AT: &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (res interface{}, err error) {
				res = p.Source.(Nature).UpdatedAt
				return
			},
		},
		FLD_TYPE_CREATED_AT: &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (res interface{}, err error) {
				res = p.Source.(Nature).CreatedAt
				return
			},
		},
		FLD_TYPE_DELETED_AT: &graphql.Field{
			Type: graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (res interface{}, err error) {

				if p.Source.(Nature).DeletedAt != 0 {
					res = true
				}

				res = false

				return
			},
		},
	},
})

var arg_type_id = &graphql.ArgumentConfig{
	Type: graphql.Int,
}

var arg_type_name = &graphql.ArgumentConfig{
	Type: graphql.String,
}

var arg_type_description = &graphql.ArgumentConfig{
	Type: graphql.String,
}

var arg_type_deleted_at = &graphql.ArgumentConfig{
	Type: graphql.Int,
}

func getNatures(p graphql.ResolveParams) (res interface{}, err error) {

	newList := []Nature{}
	t := model_v1_nature.Nature{}

	stringFilter := map[string]string{}
	intFilter := map[string]int64{}

	for k, v := range p.Args {

		switch k {
		case FLD_TYPE_ID:
			intFilter[k] = int64(v.(int))
		case FLD_TYPE_NAME:
			stringFilter[k] = v.(string)
		case FLD_TYPE_DESCRIPTION:
			stringFilter[k] = v.(string)
		}
	}

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

	err, _, list := t.GetList(stringFilter, intFilter, per_page, page*per_page)

	for i, e := range list {
		newList = append(newList, Nature{
			Id:          int(e.Id),
			Name:        e.Name,
			Description: e.Description,
			CreatedAt:   int(e.CreatedAt),
			DeletedAt:   int(e.DeletedAt),
			UpdatedAt:   int(e.UpdatedAt),
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

func typeMutation(p graphql.ResolveParams) (res interface{}, err error) {

	stringFilter := map[string]string{}
	intFilter := map[string]int64{}
	newType := &model_v1_nature.Nature{}

	for k, v := range p.Args {

		switch k {
		case FLD_TYPE_NAME:

			stringFilter[k] = v.(string)
		case FLD_TYPE_DESCRIPTION:

			stringFilter[k] = v.(string)
		case FLD_TYPE_ID:

			intFilter[k] = int64(v.(int))
		case FLD_TYPE_DELETED_AT:

			intFilter[k] = time.Now().Unix()
		}
	}

	err, id := newType.CreateOrUpdate(stringFilter, intFilter)

	if err != nil {
		res = nil
	}

	err, res = newType.GetById(id)
	if err != nil {
		res = nil
	}

	return
}
