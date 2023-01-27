package handler

import (
	"database/sql"
	"net/http"
	"strings"

	"./MySQL"
	"github.com/gin-gonic/gin"
)

type loginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var currentDB *sql.DB
var currentID = -1
var loginInfo = []loginCredentials{
	{Username: "", Password: ""},
}

func SetDB(db *sql.DB) {
	currentDB = db
}

func GetLogins(c *gin.Context) { // gin.Context parameter.
	if loginInfo[0].Username != "" {
		currentID = MySQL.Login(currentDB, loginInfo[0].Username, loginInfo[0].Password)

		if currentID == -1 {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Incorrect Username or Password!"})
		} else if currentID == -2 {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Error"})
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{"message": currentID})
		}
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Enter Username and Password!"})
	}
}

func SetCredentials(c *gin.Context) {
	combinedCredentials := c.Param("credentials")
	splitCredentials := strings.Split(combinedCredentials, ";") //Usernames or Passwords cannot have special character ';' unless encryption used (future issue)

	loginInfo[0].Username = splitCredentials[0]
	loginInfo[0].Password = splitCredentials[1]

	c.IndentedJSON(http.StatusOK, loginInfo)
}

/*func PostLogins(c *gin.Context) {
	var newLogin loginCredentials
	// To bind the received JSON to newTshirt, call BindJSON
	if err := c.BindJSON(&newLogin); err != nil {
		return
	}
	// Add the new tshirt to the slice.
	loginInfo[0] = newLogin
	c.IndentedJSON(http.StatusCreated, newLogin)
}*/

func PingGet(firstString string, secondString string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			firstString: secondString,
		})
	}
}

/*
SQL Commands For Future Use:

MySQL.ResetTables(db)
MySQL.SetUpTables(db)

fmt.Println("Users Size:", MySQL.GetTableSize(db, "Users"))
fmt.Println("Subscriptions Size:", MySQL.GetTableSize(db, "Subscriptions"))

MySQL.DeleteUser(db, "root", "password") //might not need
MySQL.DeleteUser(db, 2)

MySQL.CreateNewUser(db, "test", "testing")
MySQL.Login(db, "root", "password")

*/
