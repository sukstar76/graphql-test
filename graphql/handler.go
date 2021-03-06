package graphql

import(
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/jinzhu/gorm"

)

func NewMysqlHandler(db *gorm.DB)(*handler.Handler, error){
	schema, err:= graphql.NewSchema(
		graphql.SchemaConfig{
			Query: NewQuery(db),
			Mutation: NewMutation(db),
		},
	)

	if err !=nil{
		return nil,err
	}

	return handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
		GraphiQL: true,
	}), nil
}
