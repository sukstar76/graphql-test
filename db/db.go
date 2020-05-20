package db

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	/*"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"*/
)

func NewDB() (*gorm.DB,error){
	DBMS := "mysql"
	config := &mysql.Config{
		User: "root",
		Passwd: "root",
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		DBName: "graphql",
		AllowNativePasswords: true,
		Params: map[string]string{
			"parseTime": "true",
		},
	}

	return gorm.Open(DBMS,config.FormatDSN())
}

/*func NewmongoDB() (*mongo.Client,error){
	clientOptions := options.Client().ApplyURL("mongodb://172.18.0.2:27017")
	client, err := mongo.Connect(context.TODO(),clientOptions)

	if err!=nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(),nil)

	if err!=nil{
		log.Fatal(err)
	}

	return client, nil
}*/