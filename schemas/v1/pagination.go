package schemas

import (
	// "github.com/avachen2005/taxonomy_go/model/v1/entity"
	"github.com/avachen2005/taxonomy_go/model/v1/nature"
	"github.com/graphql-go/graphql"

	"fmt"
	// "github.com/kr/pretty"
)

const (
	FLD_PAGE_INFO = "page_info"
	FLD_PAGE      = "page"
	FLD_PER_PAGE  = "per_page"
	FLD_CURSOR    = "cursor"
	FLD_TOTAL     = "total"
)

const (
	default_page     = 1
	default_per_page = 25
)

type Pagination struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
	Cursor  int `json:"cursor"`
}

var arg_page = &graphql.ArgumentConfig{
	Type: graphql.Int,
}

var arg_per_page = &graphql.ArgumentConfig{
	Type: graphql.Int,
}

var PaginationType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "page_info",
	Description: "page_info",
	Fields: graphql.Fields{
		FLD_PER_PAGE: &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (res interface{}, err error) {

				res = p.Source.(*Pagination).PerPage
				return
			},
		},
		FLD_PAGE: &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (res interface{}, err error) {

				res = p.Source.(*Pagination).Page
				return
			},
		},
		FLD_CURSOR: &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (res interface{}, err error) {

				res = p.Source.(*Pagination).Cursor
				return
			},
		},
	},
})

var NatureTotal = graphql.NewObject(graphql.ObjectConfig{
	Name:        "total",
	Description: "total",
	Fields: graphql.Fields{
		"entity": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (res interface{}, err error) {

				return
			},
		},
		"nature": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (res interface{}, err error) {

				fmt.Println(p.Args)
				t := &model_v1_nature.Nature{}

				err, total := t.GetTotal()

				res = int(total)

				fmt.Println(fmt.Printf("total: %d", total))
				return
			},
		},
	},
})
