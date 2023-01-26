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

	//M.ResetTables(db)
	MySQL.SetUpTables(db)

	//fmt.Println("Users Size:", MySQL.GetTableSize(db, "Users"))
	//fmt.Println("Subscriptions Size:", MySQL.GetTableSize(db, "Subscriptions"))

	//MySQL.DeleteUser(db, "root", "password") //might not need
	//MySQL.DeleteUser(db, 2)

	//MySQL.CreateNewUser(db, "test", "testing")
	//MySQL.Login(db, "root", "password")

	handler.SetDB(db)

	//Angular Connection:
	r := gin.Default()

	//r.GET("/login", handler.PingGet("User_ID", strconv.Itoa(MySQL.Login(db, "root", "password")))) //Don't get why it's GET and not POST
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
