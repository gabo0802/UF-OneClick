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

	MySQL.CreateNewUser(db, "root", "password")
	fmt.Println("Number of rows in table:", MySQL.GetTableSize(db, "Users"))
	MySQL.Login(db, "root", "password")

	fmt.Println("Tables:")
	MySQL.ShowDatabaseTables(db, "userdb")

	fmt.Println("Column Data:")
	fmt.Println("UserID -")
	MySQL.GetColumnData(db, "userdb", "Users", "UserID")
	fmt.Println("Username -")
	MySQL.GetColumnData(db, "userdb", "Users", "Username")
	fmt.Println("Password -")
	MySQL.GetColumnData(db, "userdb", "Users", "Password")

	//Sets pointer in "handler" function to main.go's db
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
