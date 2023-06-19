package main

import (
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

type Tutorial struct {
	ID       int
	Title    string
	Author   *Author
	Comments []Comment
}

type Author struct {
	Name      string
	Tutorials []int
}

type Comment struct {
	Body string
}

func populate() []Tutorial {

	author := Author{Name: "Razvan Sima", Tutorials: []int{2}}
	tutorials := []Tutorial{{
		ID:     1,
		Title:  "Go graphql",
		Author: &author,
		Comments: []Comment{
			{Body: "First comment"},
		},
	},
	}

	return tutorials
}

func main() {
	fmt.Println("GraphQL tutorial")

	//define comment  graphQl
	var commentType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Comment",
			Fields: graphql.Fields{
				"body": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)

	//define author type
	var authorType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "author",
			Fields: graphql.Fields{
				"Name": &graphql.Field{
					Type: graphql.String,
				},
				"Tutorials": &graphql.Field{
					Type: graphql.NewList(graphql.Int),
				},
			},
		},
	)

	//define tutorial type
	var tutorialType = graphql.NewObject(
		graphql.ObjectConfig{ //
			Name: "tutorial",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"Title": &graphql.Field{
					Type: graphql.String,
				},
				"Author": &graphql.Field{
					Type: authorType,
				},
				"Comments": &graphql.Field{
					Type: graphql.NewList(commentType),
				},
			},
		},
	)

	tutorials := populate()

	fields := graphql.Fields{
		"tutorials": &graphql.Field{
			Type:        tutorialType,
			Description: "get tutorial by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(int)
				if ok {
					for _, tutorial := range tutorials {
						if int(tutorial.ID) == id {
							return tutorial, nil
						}
					}
				}
				return nil, nil
			},
		},
		"list": &graphql.Field{
			Type:        graphql.NewList(tutorialType),
			Description: "get all tutorials",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return tutorials, nil
			},
		},
	}

	//create the schema base on defined fields
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		panic(err)
	}

	query := `
		{
			list {
				id
				Title
			}
		} `

	//Query execution
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		fmt.Printf("Failed to execute graphql operation, errors: %+v", r.Errors)
	}
	//Query result
	rJSON, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", rJSON)

}
