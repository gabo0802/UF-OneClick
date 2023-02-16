package main

import (
	"fmt"

	"github.com/gabo0802/UF-OneClick/api/httpd/handler"
	"github.com/gabo0802/UF-OneClick/api/httpd/handler/MySQL"
	"github.com/gin-gonic/gin"
)

func main() {
	//Establishes a connection to the remote MySQL server's database
	db := MySQL.MySQLConnect()

	//Defers the closing of the connection to the database until the end of main
	defer db.Close()

	//MySQL.TestBackend(db)

	//MySQL.ResetAllTables(db)
	MySQL.SetUpTables(db)

	if MySQL.GetTableSize(db, "Users") == 0 {
		MySQL.CreateAdminUser(db)
	}

	if MySQL.GetTableSize(db, "Subscriptions") == 0 {
		MySQL.CreateCommonSubscriptions(db)
	}

	//Sets pointer in "handler" package to main.go's db
	handler.SetDB(db)

	//Angular Connection:
	r := gin.Default()

	api := r.Group("/api")
	{
		//Background
		api.GET("/remind", handler.DailyReminder)
		api.POST("/remind", handler.DailyReminder)

		//Account Management
		api.POST("/login", handler.TryLogin)

		api.POST("/accountcreation", handler.NewUser)

		//api.POST("/changepassword", handler.ChangeUserPassword) //need to agree on how to get user input (maybe name could be old password)

		api.GET("/logout", handler.Logout(""))
		api.POST("/logout", handler.Logout(""))

		api.GET("/verify/:code", handler.VerifyEmail)

		//Subscription Management
		api.POST("/subscriptions", handler.GetAllUserSubscriptions())
		api.GET("/subscriptions", handler.GetAllUserSubscriptions())

		api.POST("/subscriptions/createsubscription", handler.NewSubscriptionService)

		api.POST("/subscriptions/addsubscription", handler.NewUserSubscription)

		api.POST("/subscriptions/cancelsubscription", handler.CancelSubscriptionService)

		//Admin Commands
		api.GET("/reset", handler.ResetALL)
		api.POST("/reset", handler.ResetALL)

		api.GET("/alldata", handler.GetAllUserData())
		api.POST("/alldata", handler.GetAllUserData())
	}

	r.Run("0.0.0.0:5000") //http://127.0.0.1:5000
	r.Run("localhost:4200")
	fmt.Println("End")
}

//Test
