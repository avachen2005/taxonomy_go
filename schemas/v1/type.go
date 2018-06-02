package schemas

import (
	"time"

	"github.com/avachen2005/taxonomy_go/model/v1/type"
	"github.com/graphql-go/graphql"
	// "fmt"
	// "github.com/kr/pretty"
)

const (
	FLD_TYPE_ID          = "id"
	FLD_TYPE_NAME        = "name"
	FLD_TYPE_DESCRIPTION = "description"
	FLD_TYPE_CREATED_AT  = "created_at"
	FLD_TYPE_UPDATED_AT  = "updated_at"
	FLD_TYPE_DELETED_AT  = "deleted_at"
)

var TypeType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "type",
	Description: "",
	Fields: graphql.Fields{
		FLD_TYPE_ID: &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (res interface{}, err error) {
				res = int(p.Source.(model_v1_type.Type).Id)
				return
			},
		},
		FLD_TYPE_NAME: &graphql.Field{
			Type: graphql.String,
		},
		FLD_TYPE_DESCRIPTION: &graphql.Field{
			Type: graphql.String,
		},
		FLD_TYPE_CREATED_AT: &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (res interface{}, err error) {
				res = int(p.Source.(model_v1_type.Type).CreatedAt)
				return
			},
		},
		FLD_TYPE_UPDATED_AT: &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (res interface{}, err error) {
				res = int(p.Source.(model_v1_type.Type).UpdatedAt)
				return
			},
		},
		FLD_TYPE_DELETED_AT: &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (res interface{}, err error) {
				res = int(p.Source.(model_v1_type.Type).DeletedAt)
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

func getType(p graphql.ResolveParams) (res interface{}, err error) {

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

	t := model_v1_type.Type{}

	err, _, res = t.GetList(stringFilter, intFilter)

	return
}

func typeMutation(p graphql.ResolveParams) (res interface{}, err error) {

	// fmt.Printf("%# v\n", pretty.Formatter(p.Args))

	stringFilter := map[string]string{}
	intFilter := map[string]int64{}
	newType := &model_v1_type.Type{}

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

	// fmt.Printf("%# v\n", pretty.Formatter(intFilter))
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
