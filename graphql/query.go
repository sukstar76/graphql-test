package graphql
import(
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	"example.com/m/graphql/field"

)

 func NewQuery(db *gorm.DB) *graphql.Object{
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"product": field.GetProductByID(db),
			"list": field.GetProductList(db),
		},
	})
}