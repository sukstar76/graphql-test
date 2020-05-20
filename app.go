package main

import (
	"log"
	//"github.com/labstack/echo"
	//"github.com/labstack/echo/middleware"
	"example.com/m/db"
	"example.com/m/graphql"
	"net/http"
)

func main() {
	mysqldb, err := db.NewDB()
	if err != nil {
		log.Fatalln(err)
	}

	//mongodb, err := db.NewmognoDB()

	mysqldb.LogMode(true)
	defer mysqldb.Close()

	/*e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())*/

	/*e.GET("/",func(c echo.Context) error{
		return c.String(http.StatusOK,"Hello\n")
	})*/

	h, err := graphql.NewMysqlHandler(mysqldb)
	if err != nil {
		log.Fatalln(err)
	}

	http.Handle("/graphql", CORS(h))
	log.Fatal(http.ListenAndServe(":1323", nil))

	//e.POST("/graphql", echo.WrapHandler(h))

	//e.Logger.Fatal(e.Start(":3000"))

}

func CORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, Accept-Encoding")
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Max-Age", "86400")
			w.WriteHeader(http.StatusOK)
			return
		}
		h.ServeHTTP(w, r)
	})
}
