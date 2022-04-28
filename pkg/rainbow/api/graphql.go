package api

import (
	"net/http"

	"github.com/friendsofgo/graphiql"
	graphql "github.com/graphql-go/graphql"
	graphhandler "github.com/graphql-go/handler"

	"github.com/teal-finance/rainbow/pkg/rainbow"
)

func (h APIHandler) GraphQLHandler() http.Handler {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: graphql.NewObject(graphql.ObjectConfig{
				Name:   "RootQuery",
				Fields: h.fieldsRoot(),
			}),
		},
	)
	if err != nil {
		panic(err)
	}
	return graphhandler.New(&graphhandler.Config{Schema: &schema})
}

func InteractiveGQLHandler(endpoint string) http.Handler {
	h, err := graphiql.NewGraphiqlHandler(endpoint)
	if err != nil {
		panic(err)
	}
	return h
}

// GetRootFields returns all the available queries.
func (h APIHandler) fieldsRoot() graphql.Fields {
	return graphql.Fields{
		"rows": h.fieldsCallPut(),
	}
}

// GetUserQuery returns the queries available against user type.
func (h APIHandler) fieldsCallPut() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(
			graphql.NewObject(graphql.ObjectConfig{
				Name: "CallPut",
				Fields: graphql.Fields{
					"date":     &graphql.Field{Type: graphql.String},
					"expiry":   &graphql.Field{Type: graphql.String},
					"provider": &graphql.Field{Type: graphql.String},
					"asset":    &graphql.Field{Type: graphql.String},
					"strike":   &graphql.Field{Type: graphql.Float},
					"call":     &graphql.Field{Type: limitType},
					"put":      &graphql.Field{Type: limitType},
				},
			}),
		),
		Args: graphql.FieldConfigArgument{
			"assets": &graphql.ArgumentConfig{
				Type: graphql.NewList(graphql.String),
			},
			"expiries": &graphql.ArgumentConfig{
				Type: graphql.NewList(graphql.String),
			},
			"providers": &graphql.ArgumentConfig{
				Type: graphql.NewList(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			// parse args
			args := rainbow.StoreArgs{}

			if v, ok := params.Args["assets"]; ok {
				assets := v.([]interface{})
				args.Assets = make([]string, len(assets))
				for i, a := range assets {
					args.Assets[i] = a.(string)
				}
			}

			if v, ok := params.Args["expiries"]; ok {
				expiries := v.([]interface{})
				args.Expiries = make([]string, len(expiries))
				for i, e := range expiries {
					args.Expiries[i] = e.(string)
				}
			}

			if v, ok := params.Args["providers"]; ok {
				providers := v.([]interface{})
				args.Providers = make([]string, len(providers))
				for i, p := range providers {
					args.Providers[i] = p.(string)
				}
			}

			// get options
			options, err := h.Service.Options(args)
			if err != nil {
				return nil, err
			}

			rows := buildCallPut(options)
			return rows, nil
		},
	}
}

var limitType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Limit",
	Fields: graphql.Fields{
		"bid": &graphql.Field{Type: orderType},
		"ask": &graphql.Field{Type: orderType},
	},
})

var orderType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Order",
	Fields: graphql.Fields{
		"px":   &graphql.Field{Type: graphql.String},
		"size": &graphql.Field{Type: graphql.String},
	},
})
