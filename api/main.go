package main

import (
	"fmt"

	"./httpd/handler"

	"./httpd/handler/MySQL"
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
	r.GET("", handler.PingGet("Website", "*insert welcome page*"))

	r.GET("/login", handler.GetLogins)
	r.GET("/login/:credentials", handler.SetCredentials)
	//r.POST("/login", handler.PostLogins)

	r.GET("/subscriptions", handler.PingGet("Subscriptions", "*insert all of users subscriptions*"))

	/*
		api := r.Group("/api")	{
			api.GET("/ping", handler.PingGet())
		}
	*/

	r.Run("0.0.0.0:5000") //http://127.0.0.1:5000

	fmt.Println("End")
}
