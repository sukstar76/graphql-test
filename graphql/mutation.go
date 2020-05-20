package graphql
import(
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	"example.com/m/graphql/field"

)

func NewMutation(db *gorm.DB) *graphql.Object{
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"create": field.CreateProduct(db),
			
			"update": field.UpdateProductByID(db),
	
			"delete": field.DeleteProductByID(db),
		},
	})
}