package schemas

import (
	// "github.com/avachen2005/taxonomy_go/model/v1/entity"
	"github.com/avachen2005/taxonomy_go/model/v1/type"
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
				if v, ok := p.Args[FLD_PER_PAGE]; ok {
					res = v
				}
				return
			},
		},
		FLD_PAGE: &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (res interface{}, err error) {
				if v, ok := p.Args[FLD_PAGE]; ok {
					res = v
				}
				return
			},
		},
	},
})

var TotalType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "total",
	Description: "total",
	Fields: graphql.Fields{
		"entity": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (res interface{}, err error) {

				return
			},
		},
		"type": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (res interface{}, err error) {

				fmt.Println(p.Args)
				t := &model_v1_type.Type{}

				err, total := t.GetTotal()

				res = int(total)

				fmt.Println(fmt.Printf("total: %d", total))
				return
			},
		},
	},
})
