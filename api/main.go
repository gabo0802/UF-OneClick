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

	//MySQL.ResetAllTables(db)
	MySQL.SetUpTables(db)
	//MySQL.CreateAdminUser(db)

	//Sets pointer in "handler" package to main.go's db
	handler.SetDB(db)

	//Angular Connection:
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("", handler.HomePage())

		//Account Management
		api.GET("/login", handler.TryLogin)
		api.GET("/login/:data", handler.SetCookie("/api/login"))

		api.GET("/accountcreation", handler.NewUser)
		api.GET("/accountcreation/:data", handler.SetCookie("/api/accountcreation"))

		api.GET("/changepassword", handler.ChangeUserPassword)
		api.GET("/changepassword/:data", handler.SetCookie("/api/changepassword"))

		api.GET("/logout", handler.Logout("Enter"))

		//Subscription Management
		api.GET("/subscriptions", handler.GetAllUserSubscriptions())

		api.GET("/subscriptions/createsubscription", handler.NewSubscriptionService)
		api.GET("/subscriptions/createsubscription/:data", handler.SetCookie("/api/subscriptions/createsubscription"))

		api.GET("/subscriptions/addsubscription", handler.NewUserSubscription)
		api.GET("/subscriptions/addsubscription/:data", handler.SetCookie("/api/subscriptions/addsubscription"))

		api.GET("/subscriptions/cancelsubscription", handler.CancelSubscriptionService)
		api.GET("/subscriptions/cancelsubscription/:data", handler.SetCookie("/api/subscriptions/cancelsubscription"))

		//Admin Commands
		api.GET("/reset", handler.ResetDatabase)
		api.GET("/alldata", handler.GetAllUserData())
		api.GET("/alldata/:data", handler.SetCookie("/api/alldata"))

	}

	r.Run("0.0.0.0:5000") //http://127.0.0.1:5000

	fmt.Println("End")
}

//Test
