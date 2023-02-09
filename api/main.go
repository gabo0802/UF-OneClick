package main

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gabo0802/UF-OneClick/api/httpd/handler"
	"github.com/gabo0802/UF-OneClick/api/httpd/handler/MySQL"

	"github.com/gin-gonic/gin"
)

// UserID will normally be taken automatically from the database
// instead of specifying it directly, unlike in this test
func testBackend(db *sql.DB) {
	fmt.Println("Type -1 to quit the test.")
	var choice int
	for choice != -1 {
		fmt.Print("Enter a number from 1 - 10: ")
		fmt.Scanln(&choice)
		if choice == 1 {
			MySQL.ResetAllTables(db)
			fmt.Println("Choice 1: Clearing tables from database \"userdb.\"")
		} else if choice == 2 {
			MySQL.SetUpTables(db)
			fmt.Println("Choice 2: Setting up all tables for database \"userdb.\"")
		} else if choice == 3 {
			fmt.Println("Choice 3: Getting table sizes.")
			fmt.Println("Subscription table size: " + strconv.Itoa(MySQL.GetTableSize(db, "subscriptions")))
			fmt.Println("Users table size: " + strconv.Itoa(MySQL.GetTableSize(db, "users")))
			fmt.Println("UserSubs table size: " + strconv.Itoa(MySQL.GetTableSize(db, "usersubs")))
		} else if choice == 4 {
			fmt.Println("Choice 4: Creates new user.")
			fmt.Println("Enter a username, password, and email: ")
			var a, b, c string
			fmt.Scanln(&a, &b, &c)
			MySQL.CreateNewUser(db, a, b, c)
		} else if choice == 5 {
			fmt.Println("Choice 5: Creates new subscription.")
			fmt.Println("Enter a name and price: ")
			var a, b string
			fmt.Scanln(&a, &b)
			MySQL.CreateNewSub(db, a, b)
		} else if choice == 6 {
			fmt.Println("Choice 6: Subscribes to subscription service.")
			fmt.Println("Enter a UserID and Subscription Name: ")
			var a int
			var b string
			fmt.Scanln(&a, &b)
			MySQL.CreateNewUserSub(db, a, b)
		} else if choice == 7 {
			fmt.Println("Choice 7: Cancels subscription service.")
			fmt.Println("Enter a UserID and Subscription Name: ")
			var a int
			var b string
			fmt.Scanln(&a, &b)
			MySQL.CancelUserSub(db, a, b)
		} else if choice == 8 {
			fmt.Println("Choice 8: Resubscribes to subscription service.")
			fmt.Println("Enter a UserID, Subscription Name, and Date Added: ")
			var a int
			var b, c string
			fmt.Scanln(&a, &b, &c)
			MySQL.AddOldUserSub(db, a, b, c)
		} else if choice == 9 {
			fmt.Println("Choice 9: Deletes user that is specified.")
			fmt.Println("Enter a UserID: ")
			var a int
			fmt.Scanln(&a)
			MySQL.DeleteUser(db, a)
		} else if choice == 10 {
			fmt.Println("Choice 10: Changes the password of specified user.")
			fmt.Println("Enter a UserID, the Old Password, and the New Password: ")
			var a int
			var b, c string
			fmt.Scanln(&a, &b, &c)
			MySQL.ChangePassword(db, a, b, c)
		}
	}
}

func main() {
	//Establishes a connection to the remote MySQL server's database
	db := MySQL.MySQLConnect()

	//Defers the closing of the connection to the database until the end of main
	defer db.Close()

	testBackend(db)

	//MySQL.ResetAllTables(db)
	//MySQL.SetUpTables(db)
	//MySQL.CreateAdminUser(db)

	//Sets pointer in "handler" package to main.go's db
	handler.SetDB(db)

	//Angular Connection:
	r := gin.Default()

	api := r.Group("/api")
	{
		//api.GET("", handler.HomePage())

		//Account Management
		api.POST("/login", handler.TryLogin)

		api.POST("/accountcreation", handler.NewUser)
		//api.GET("/changepassword", handler.ChangeUserPassword)
		//api.GET("/changepassword/:data", handler.SetCookie("/api/changepassword"))

		//api.GET("/logout", handler.Logout("Enter"))

		//Subscription Management
		api.POST("/subscriptions", handler.GetAllUserSubscriptions())

		//api.GET("/subscriptions/createsubscription", handler.NewSubscriptionService)
		//api.GET("/subscriptions/createsubscription/:data", handler.SetCookie("/api/subscriptions/createsubscription"))

		//api.GET("/subscriptions/addsubscription", handler.NewUserSubscription)
		//api.GET("/subscriptions/addsubscription/:data", handler.SetCookie("/api/subscriptions/addsubscription"))

		//api.GET("/subscriptions/cancelsubscription", handler.CancelSubscriptionService)
		//api.GET("/subscriptions/cancelsubscription/:data", handler.SetCookie("/api/subscriptions/cancelsubscription"))

		//Admin Commands
		api.GET("/reset", handler.ResetDatabase)
		api.GET("/alldata", handler.GetAllUserData())
	}

	r.Run("0.0.0.0:5000") //http://127.0.0.1:5000
	//r.Run("localhost:4200")
	fmt.Println("End")
}

//Test
