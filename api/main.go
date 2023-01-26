package main

import (
	M "./MySQL"
	//"github.com/gin-gonic/gin"
)

func main() {
	//Testing
	db := M.MySQLConnect()

	//ResetTables(db)
	M.SetUpTables(db)

	//fmt.Println("Users Size:", M.GetTableSize(db, "Users"))
	//fmt.Println("Subscriptions Size:", getTableSize(db, "Subscriptions"))

	//M.DeleteUser(db, "root", "password") //might not need
	//M.DeleteUser(db, 2)

	//M.CreateNewUser(db, "test", "testing")
	M.Login(db, "root", "password")

	//r := gin.Default()
	/*api := r.Group("/api")
	{
		api.POST("/newsfeed", tableSize)
	}*/
	//r.Run("0.0.0.0:5000")
}
