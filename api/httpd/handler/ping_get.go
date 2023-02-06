package handler

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/gabo0802/UF-OneClick/api/httpd/handler/MySQL"
	"github.com/gin-gonic/gin"
)

type loginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type userSubscriptions struct {
	Name        string `json:"name"`
	Price       string `json:"price"`
	DateAdded   string `json:"dateadded"`
	DateRemoved string `json:"dateremoved"`
}

var currentDB *sql.DB
var currentID = -1
var loginInfo = []loginCredentials{
	{Username: "", Password: ""},
}
var usersubInfo = []userSubscriptions{
	//{Name: "", Price: "", DateAdded: "", DateRemoved: ""},
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

	loginInfo = []loginCredentials{
		{Username: "", Password: ""},
	}
}

func NewUser(c *gin.Context) {
	if loginInfo[0].Username != "" {
		rowsAffected := MySQL.CreateNewUser(currentDB, loginInfo[0].Username, loginInfo[0].Password)

		if rowsAffected == 0 {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Error: Username Already Exists!"})
		} else if rowsAffected == -1 {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Error"})
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "New User Created!"})
		}

	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Create New User!"})
	}

	loginInfo = []loginCredentials{
		{Username: "", Password: ""},
	}
}

func ChangePass(c *gin.Context) {
	if currentID != -1 {
		if loginInfo[0].Username != "" {
			rowsAffected := MySQL.ChangePassword(currentDB, currentID, loginInfo[0].Username, loginInfo[0].Password)

			if rowsAffected == 0 {
				c.IndentedJSON(http.StatusOK, gin.H{"message": "Error: Wrong Username or Password!"})
			} else if rowsAffected == -1 {
				c.IndentedJSON(http.StatusOK, gin.H{"message": "Error"})
			} else {
				c.IndentedJSON(http.StatusOK, gin.H{"message": "Password Changed!"})
			}
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Enter Old Password and New Password!"})
		}

	} else {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "Error: You Should Not Be Here"})
	}

	loginInfo = []loginCredentials{
		{Username: "", Password: ""},
	}
}

func SetCredentials(c *gin.Context) {
	combinedCredentials := c.Param("credentials")
	splitCredentials := strings.Split(combinedCredentials, ";") //Usernames or Passwords cannot have special character ';' unless encryption used (future issue)

	loginInfo[0].Username = splitCredentials[0]

	if len(splitCredentials) > 1 {
		loginInfo[0].Password = splitCredentials[1]
	} else {
		loginInfo[0].Password = ""
	}

	c.IndentedJSON(http.StatusOK, loginInfo)
}

func GetAllUserSubs(c *gin.Context) {
	usersubInfo = []userSubscriptions{}

	if currentID != -1 {
		rows, err := currentDB.Query("SELECT Name, Price, DateAdded, DateRemoved FROM UserSubs INNER JOIN Subscriptions ON UserSubs.SubID = Subscriptions.SubID WHERE UserID = ? ORDER BY DateAdded ASC", currentID)
		//can order by anything

		if err != nil {
			c.IndentedJSON(http.StatusBadGateway, gin.H{"message": "Error"})
		}

		for rows.Next() {
			var newUserSub userSubscriptions
			rows.Scan(&newUserSub.Name, &newUserSub.Price, &newUserSub.DateAdded, &newUserSub.DateRemoved)
			usersubInfo = append(usersubInfo, newUserSub)
		}

		c.IndentedJSON(http.StatusOK, usersubInfo)

	} else {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "Error: You Should Not Be Here"})
	}

}

func NewUserSub(c *gin.Context) {
	if currentID != -1 {
		if loginInfo[0].Username != "" {
			rowsAffected := MySQL.CreateNewUserSub(currentDB, currentID, loginInfo[0].Username)

			if rowsAffected == 0 {
				c.IndentedJSON(http.StatusOK, gin.H{"message": "Subscription to " + loginInfo[0].Username + " Already Active!"})
			} else if rowsAffected == -1 {
				c.IndentedJSON(http.StatusOK, gin.H{"message": "Error"})
			} else if rowsAffected > 1 {
				c.IndentedJSON(http.StatusOK, gin.H{"message": "Subscription to " + loginInfo[0].Username + " Renewed!"})
			} else {
				c.IndentedJSON(http.StatusOK, gin.H{"message": "Subscription to " + loginInfo[0].Username + " Added!"})
			}
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Choose Subscription to Add or Renew!"})
		}
	} else {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "Error: You Should Not Be Here"})
	}

	loginInfo = []loginCredentials{
		{Username: "", Password: ""},
	}
}

func CancelSub(c *gin.Context) {
	if currentID != -1 {
		if loginInfo[0].Username != "" {
			rowsAffected := MySQL.CancelUserSub(currentDB, currentID, loginInfo[0].Username)

			if rowsAffected == 0 {
				c.IndentedJSON(http.StatusOK, gin.H{"message": "Error: Subscription to " + loginInfo[0].Username + " Does Not Exist!"})
			} else if rowsAffected == -1 {
				c.IndentedJSON(http.StatusOK, gin.H{"message": "Error"})
			} else {
				c.IndentedJSON(http.StatusOK, gin.H{"message": "Subscription to " + loginInfo[0].Username + " Canceled!"})
			}
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Choose Subscription to Cancel!"})
		}
	} else {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "Error: You Should Not Be Here"})
	}

	loginInfo = []loginCredentials{
		{Username: "", Password: ""},
	}
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
