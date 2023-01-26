package main

import (
	"./handler"

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

	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	/*api := r.Group("/api")
	{
		api.POST("/newsfeed", tableSize)
	}*/
	r.GET("/ping", handler.PingGet())

	//r.Run()
	r.Run("0.0.0.0:5000")
}
