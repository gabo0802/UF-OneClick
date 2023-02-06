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

	//Sets pointer in "handler" package to main.go's db
	handler.SetDB(db)

	//Angular Connection:
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("", handler.PingGet("Website", "*insert welcome page*"))

		api.GET("/login", handler.GetLogins)
		api.GET("/login/:credentials", handler.SetCredentials)
		//r.POST("/login", handler.PostLogins)

		api.GET("/accountcreation", handler.NewUser)
		api.GET("/accountcreation/:credentials", handler.SetCredentials)

		api.GET("/changepassword", handler.ChangePass)
		api.GET("/changepassword/:credentials", handler.SetCredentials)

		api.GET("/subscriptions", handler.GetAllUserSubs)

		api.GET("/subscriptions/addsubscription", handler.NewUserSub)
		api.GET("/subscriptions/addsubscription/:credentials", handler.SetCredentials)

		api.GET("/subscriptions/cancelsubscription", handler.CancelSub)
		api.GET("/subscriptions/cancelsubscription/:credentials", handler.SetCredentials)
	}

	r.Run("0.0.0.0:5000") //http://127.0.0.1:5000

	fmt.Println("End")
}

//Test
