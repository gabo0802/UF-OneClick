package main

import (
	"fmt"

	"github.com/gabo0802/UF-OneClick/api/httpd/handler"

	"github.com/gabo0802/UF-OneClick/api/httpd/handler/MySQL"
	"github.com/gin-gonic/gin"
)

func main() {
	//MySQL Connection:
	db := MySQL.MySQLConnect()
	//MySQL.ResetTables(db)
	MySQL.SetUpTables(db)
	handler.SetDB(db)

	//Angular Connection:
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("", handler.PingGet("Website", "*insert welcome page*"))
		api.GET("/login", handler.GetLogins)
		api.GET("/login/:credentials", handler.SetCredentials)
		//r.POST("/login", handler.PostLogins)
		api.GET("/subscriptions", handler.PingGet("Subscriptions", "*insert all of users subscriptions*"))
	}
	/*
		r.GET("", handler.PingGet("Website", "*insert welcome page*"))
		r.GET("/login", handler.GetLogins)
		r.GET("/login/:credentials", handler.SetCredentials)
		//r.POST("/login", handler.PostLogins)
		r.GET("/subscriptions", handler.PingGet("Subscriptions", "*insert all of users subscriptions*"))
	*/
	r.Run("0.0.0.0:5000") //http://127.0.0.1:5000

	fmt.Println("End")
}

//Test
