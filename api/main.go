package main

import (
	"fmt"

	"github.com/gabo0802/UF-OneClick/api/httpd/handler"
	"github.com/gabo0802/UF-OneClick/api/httpd/handler/MySQL"
	"github.com/gin-gonic/gin"
)

func main() {
	//Establishes a connection to the remote MySQL server's database:
	db := MySQL.MySQLConnect()

	//Defers the closing of the connection to the database until the end of main:
	defer db.Close()

	//Sets Up Tables in Database:

	//MySQL.ResetAllTables(db)
	MySQL.SetUpTables(db)

	if MySQL.GetTableSize(db, "Users") == 0 {
		MySQL.CreateAdminUser(db)
	}

	if MySQL.GetTableSize(db, "Subscriptions") == 0 {
		MySQL.CreateCommonSubscriptions(db)
	}

	//Sets pointer in "handler" package to main.go's db:
	handler.SetDB(db)

	//Testing:
	//MySQL.TestBackend(db)
	//handler.SendAllReminders()

	//Angular Connection:
	r := gin.Default()

	api := r.Group("/api")
	{
		//Background
		api.GET("/remind", handler.DailyReminder)
		api.POST("/remind", handler.DailyReminder)

		//Account Management
		api.GET("/userinfo", handler.GetUserInfo)
		api.POST("/login", handler.TryLogin)
		//api.GET("/2FA")
		//api.GET("/2FA/:userCode", handler.TwoFactorAuthentication) //testing
		//api.POST("/2FA", handler.TwoFactorAuthentication()) //need to agree on how to send POST request
		api.POST("/accountcreation", handler.NewUser)
		api.PUT("/changepassword", handler.ChangeUserPassword)
		api.PUT("/changeusername", handler.ChangeUserUsername)
		api.PUT("/changeemail", handler.ChangeUserEmail)
		api.POST("/deleteuser", handler.DeleteUser)
		api.DELETE("/deleteuser", handler.DeleteUser)
		//api.GET("/logout/:valid", handler.Logout(""))
		//api.POST("/logout", handler.Logout(""))
		api.GET("/verify/:code", handler.VerifyEmail)

		//Subscription Management
		api.POST("/subscriptions", handler.GetAllUserSubscriptions())
		api.POST("/subscriptions/createsubscription", handler.NewSubscriptionService)
		api.POST("/subscriptions/addsubscription", handler.NewUserSubscription)
		api.POST("/subscriptions/addoldsubscription", handler.NewPreviousUserSubscription)
		api.POST("/subscriptions/cancelsubscription", handler.CancelSubscriptionService)

		//Admin Commands
		api.POST("/news", handler.NewsLetter) //need to agree on how to get user input (for now name is message)
		api.POST("/reset", handler.ResetALL)
		api.POST("/alldata", handler.GetAllUserData())
	}

	r.Run("0.0.0.0:5000") //http://127.0.0.1:5000
	fmt.Println("End")
}

//Test
