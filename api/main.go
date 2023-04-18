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

	if MySQL.GetTableSize(db, "Subscriptions") == 0 {
		MySQL.CreateCommonSubscriptions(db)
	}

	if MySQL.GetTableSize(db, "Users") == 0 {
		MySQL.CreateAdminUser(db)
		MySQL.CreateTestUser(db) //for testing
	}

	//Sets pointer in "handler" package to main.go's db:
	handler.SetDB(db)

	//Reminders:
	fmt.Println(handler.SendAllReminders())

	//Angular Connection:
	r := gin.Default()

	api := r.Group("/api")
	{
		//Background:
		api.GET("/remind", handler.DailyReminder)
		//api.POST("/remind", handler.DailyReminder)

		//Login and Verification:
		api.POST("/login", handler.TryLogin)
		api.GET("/verify/:code", handler.VerifyEmail)
		//api.GET("/2FA")
		//api.GET("/2FA/:userCode", handler.TwoFactorAuthentication) //testing
		//api.POST("/2FA", handler.TwoFactorAuthentication()) //need to agree on how to send POST request

		//Account Management:
		api.POST("/accountcreation", handler.NewUser)
		api.GET("/alluserinfo", handler.GetAllCurrentUserInfo)
		api.PUT("/changepassword", handler.ChangeUserPassword)
		api.PUT("/changeusername", handler.ChangeUserUsername)
		api.PUT("/changeemail", handler.ChangeUserEmail)
		api.DELETE("/deleteuser", handler.DeleteUser)

		api.GET("/alltimezones", handler.GetAllTimezones)
		api.PUT("/changetimezone", handler.ChangeTimezone)

		//Subscription Management:
		api.GET("/subscriptions/active", handler.GetAllCurrentUserSubscriptions(true))
		api.GET("/subscriptions", handler.GetAllCurrentUserSubscriptions(false))
		api.GET("/subscriptions/services", handler.GetAllSubscriptionServices())

		api.POST("/subscriptions/createsubscription", handler.NewSubscriptionService)
		api.POST("/subscriptions/addsubscription", handler.NewUserSubscription)
		api.POST("/subscriptions/addoldsubscription", handler.NewPreviousUserSubscription)
		api.POST("/subscriptions/cancelsubscription", handler.CancelSubscriptionService)
		api.DELETE("/subscriptions/:id", handler.DeleteUserSubID)

		api.GET("/longestsub", handler.GetMostUsedUserSubscription(false, false))
		api.GET("/longestcontinoussub", handler.GetMostUsedUserSubscription(true, false))
		api.GET("/longestactivesub", handler.GetMostUsedUserSubscription(false, true))

		api.GET("/avgpriceactivesub", handler.GetAvgPriceofAllCurrentUserSubscriptions(true))
		api.GET("/avgpriceallsubs", handler.GetAvgPriceofAllCurrentUserSubscriptions(false))

		api.GET("/avgageallsubs", handler.GetAvgAgeofAllCurrentUserSubscriptions(true, false))
		api.GET("/avgageactivesubs", handler.GetAvgAgeofAllCurrentUserSubscriptions(false, false))
		api.GET("/avgagecontinuoussubs", handler.GetAvgAgeofAllCurrentUserSubscriptions(false, false))

		api.POST("/getprice", handler.GetPriceForMonth())
		api.POST("/getallprices", handler.GetAllPricesInRange())

		//Admin Commands:
		api.POST("/news", handler.NewsLetter)
		api.DELETE("/reset", handler.ResetALL)
		api.GET("/alldata", handler.GetAllUserData())
	}

	r.Run("0.0.0.0:5000") //http://127.0.0.1:5000
	fmt.Println("End")
}

//Test
