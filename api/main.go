package main

import (
	"fmt"

	"./httpd/handler"

	M "./MySQL"
	"github.com/gin-gonic/gin"
)

func main() {
	//MySQL Connection:
	db := M.MySQLConnect()

	//M.ResetTables(db)
	M.SetUpTables(db)

	//fmt.Println("Users Size:", M.GetTableSize(db, "Users"))
	//fmt.Println("Subscriptions Size:", GetTableSize(db, "Subscriptions"))

	//M.DeleteUser(db, "root", "password") //might not need
	//M.DeleteUser(db, 2)

	//M.CreateNewUser(db, "test", "testing")
	M.Login(db, "root", "password")

	//Angular Connection:
	r := gin.Default()

	r.GET("/ping", handler.PingGet()) //IT WORKS!!!!!!

	/*
		api := r.Group("/api")	{
			api.GET("/ping", handler.PingGet())
		}
	*/

	r.Run("0.0.0.0:5000") //http://127.0.0.1:5000/ping

	fmt.Println("End")
}
