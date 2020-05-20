package field

import(
	"example.com/m/model"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var productType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"info": &graphql.Field{
				Type: graphql.String,

			},
			"url": &graphql.Field{
				Type: graphql.String,

			},
			"price": &graphql.Field{
				Type: graphql.Float,
			},

		},
	},
)

func GetProductByID(db *gorm.DB) *graphql.Field{
	return &graphql.Field{
		Type: productType,
		Description: "Get product by ID",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{},error){
			id, ok := p.Args["id"].(int)
			var product model.Product
			if ok{
				if err:= db.First(&product,id).Error; err !=nil{
					return nil,nil
				}
			}
			return product,nil 
		},
	}
}


func GetProductList(db *gorm.DB) *graphql.Field{
	return &graphql.Field{
		Type: graphql.NewList(productType),
		Description: "Get product list",
		Resolve: func(params graphql.ResolveParams) (interface{},error){
			var products []*model.Product
			if err := db.Find(&products).Error; err!=nil{
				return nil,nil
			}
			return products, nil
		},
	}
}

func CreateProduct(db *gorm.DB) *graphql.Field{
	return &graphql.Field{
		Type: productType,
		Description: "Create New Product",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			
			"name": &graphql.ArgumentConfig{
				Type: graphql.String,

			},
			"info": &graphql.ArgumentConfig{
				Type: graphql.String,

			},
			"url": &graphql.ArgumentConfig{
				Type: graphql.String,

			},
			"price": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
		},

		Resolve: func(params graphql.ResolveParams) (interface{}, error){
			id, _ := params.Args["id"].(int)
			name, _ := params.Args["name"].(string)
			info, _ :=params.Args["info"].(string)
			url, _ :=params.Args["url"].(string)
			price,_ :=params.Args["price"].(float64)
			product := model.Product{
				ID : id,
				Name : name,
				Info : info,
				Url : url,
				Price : price,
			}
	
			if err := db.Create(&product).Error; err !=nil{
				return nil,nil
			}
			return product,nil
		},
	}
}

func UpdateProductByID(db *gorm.DB) *graphql.Field{
	return &graphql.Field{
		Type: productType,
		Description: "Update Product by ID",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),

			},
			"name": &graphql.ArgumentConfig{
				Type: graphql.String,

			},
			"info": &graphql.ArgumentConfig{
				Type: graphql.String,

			},
			"url": &graphql.ArgumentConfig{
				Type: graphql.String,

			},
			"price": &graphql.ArgumentConfig{
				Type: graphql.Float,

			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{},error){
			id, _ := params.Args["id"].(int)
			name, nameok := params.Args["name"].(string)
			info, infook :=params.Args["info"].(string)
			url, urlok :=params.Args["url"].(string)
			price, priceok:=params.Args["price"].(float64)
			var product model.Product
			if err := db.First(&product,id).Error; err!=nil{
				return nil,nil
			}
			if nameok{
				product.Name=name
			}
			if infook{
				product.Info=info
			}
			if urlok{
				product.Url=url
			}
			if priceok{
				product.Price=price
			}
			if err := db.Save(&product).Error; err!=nil{
				return nil,nil
			}
			return product, nil
				
			

		},
	}
}

func DeleteProductByID(db *gorm.DB) *graphql.Field{
	return &graphql.Field{
		Type: productType,
		Description: "Delete Product by ID",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),

			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{},error){
			id, _ := params.Args["id"].(int)
			var product model.Product
			if err := db.First(&product,id).Error; err!=nil{
				return nil,nil
			}
			db.Delete(&product)
			return product,nil
		},
	}
}