package handler

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"github.com/gabo0802/UF-OneClick/api/httpd/handler/MySQL"
	"github.com/gin-gonic/gin"
)

// New Data Types
type Cookie struct {
	First  string `json:"first"`
	Second string `json:"second"`
	Third  string `json:"third"`
}

type userData struct {
	UserID      string `json:"userid"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	SubID       string `json:"subid"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	DateAdded   string `json:"dateadded"`
	DateRemoved string `json:"dateremoved"`
}

type userSubscriptions struct {
	Name        string `json:"name"`
	Price       string `json:"price"`
	DateAdded   string `json:"dateadded"`
	DateRemoved string `json:"dateremoved"`
}

// Global Variables:
var currentDB *sql.DB
var currentID = -1
var currentCookie = Cookie{First: "Default Message", Second: ""}

var usersubInfo = []userSubscriptions{
	//{Name: "", Price: "", DateAdded: "", DateRemoved: ""},
}

var allUserData = []userData{}

func SetDB(db *sql.DB) {
	currentDB = db
}

// JSON:
func tryDefaultMessage(newMessage string) {
	if strings.Contains(currentCookie.First, "Default Message") {
		currentCookie.First = "Message: " + newMessage
	}
}

func setDefaultMessage() {
	currentCookie.First = "Default Message"
	currentCookie.Second = ""
	currentCookie.Third = ""
}

func changeMessage(newMessage string) {
	currentCookie.First = "Message: " + newMessage
}

func printMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Message": strings.Replace(currentCookie.First, "Message: ", "", 1)})
}

// GET and POST Functions:
func TryLogin(c *gin.Context) { // gin.Context parameter.
	tryDefaultMessage("Enter Username and Password!")

	if !strings.Contains(currentCookie.First, "Message: ") {
		currentID = MySQL.Login(currentDB, currentCookie.First, currentCookie.Second)

		if currentID == -1 {
			changeMessage("Incorrect Username or Password!")
		} else if currentID == -2 {
			changeMessage("Error")
		} else {
			setDefaultMessage()
			c.Redirect(http.StatusTemporaryRedirect, "/api/subscriptions")
		}
	}

	printMessage(c)
	setDefaultMessage()
}

func HomePage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "")
	}
}

func Logout(message string) gin.HandlerFunc {
	return func(c *gin.Context) {
		setDefaultMessage()
		currentID = -1

		c.Redirect(http.StatusTemporaryRedirect, "/api/login")
	}
}

func NewUser(c *gin.Context) {
	tryDefaultMessage("Enter Username, Password, and Email to Create New User!")

	if !strings.Contains(currentCookie.First, "Message: ") {
		rowsAffected := MySQL.CreateNewUser(currentDB, currentCookie.First, currentCookie.Second, currentCookie.Third)

		if rowsAffected == 0 {
			changeMessage("Error: Username " + currentCookie.First + " Already Exists!")
		} else if rowsAffected == 10 {
			changeMessage("Error: Email " + currentCookie.Third + " Already In Use!")
		} else if rowsAffected == -1 {
			changeMessage("Error")
		} else if rowsAffected == -2 {
			changeMessage("Enter Value Into All Columns!")
		} else {
			changeMessage("New User " + currentCookie.First + " Has Been Created! Enter Username and Password!")
			c.Redirect(http.StatusTemporaryRedirect, "/api/login")
		}
	}

	printMessage(c)
	setDefaultMessage()
}

func ChangeUserPassword(c *gin.Context) {
	if currentID != -1 {
		tryDefaultMessage("Enter Old Password and New Password!")

		if !strings.Contains(currentCookie.First, "Message: ") {
			rowsAffected := MySQL.ChangePassword(currentDB, currentID, currentCookie.First, currentCookie.Second)

			if rowsAffected == 0 {
				changeMessage("Incorrect Old Password!")
			} else if rowsAffected == -1 {
				changeMessage("Error")
			} else if rowsAffected == -2 {
				changeMessage("Enter Value Into All Columns!")
			} else {
				changeMessage("Password Has Been Changed! Re-enter Username and Password!")
				currentID = -1
				c.Redirect(http.StatusTemporaryRedirect, "/api/login")
			}
		}

		printMessage(c)
		setDefaultMessage()

	} else {
		setDefaultMessage()
		c.Redirect(http.StatusTemporaryRedirect, "/api/login")
	}
}

func SetCookie(url string) gin.HandlerFunc {
	return func(c *gin.Context) {
		combinedData := c.Param("data")
		splitData := strings.Split(combinedData, ";") //Usernames or Passwords cannot have special character ';' unless encryption used (future issue)

		currentCookie.First = splitData[0]

		if len(splitData) > 1 {
			currentCookie.Second = splitData[1]
		} else {
			currentCookie.Second = ""
		}

		if len(splitData) > 2 {
			currentCookie.Third = splitData[2]
		} else {
			currentCookie.Third = ""
		}

		c.Redirect(http.StatusTemporaryRedirect, url)
	}
}

func GetAllUserSubscriptions() gin.HandlerFunc {
	return func(c *gin.Context) {
		usersubInfo = []userSubscriptions{}

		if currentID != -1 {
			rows, err := currentDB.Query("SELECT Name, Price, DateAdded, DateRemoved FROM UserSubs INNER JOIN Subscriptions ON UserSubs.SubID = Subscriptions.SubID WHERE UserID = ? ORDER BY DateAdded ASC", currentID)
			//can order by anything

			if err != nil {
				c.JSON(http.StatusBadGateway, gin.H{"message": "Error"})
			}

			for rows.Next() {
				var newUserSub userSubscriptions
				rows.Scan(&newUserSub.Name, &newUserSub.Price, &newUserSub.DateAdded, &newUserSub.DateRemoved)
				usersubInfo = append(usersubInfo, newUserSub)
			}

			c.IndentedJSON(http.StatusOK, usersubInfo)

		} else {
			setDefaultMessage()
			c.Redirect(http.StatusTemporaryRedirect, "/api/login")
		}
	}
}

func NewUserSubscription(c *gin.Context) {
	if currentID != -1 {
		tryDefaultMessage("Choose Subscription to Add or Renew!")

		if !strings.Contains(currentCookie.First, "Message: ") {
			rowsAffected := MySQL.CreateNewUserSub(currentDB, currentID, currentCookie.First)

			if rowsAffected == 0 {
				changeMessage("Subscription to " + currentCookie.First + " Already Active!")

			} else if rowsAffected == -1 {
				changeMessage("Subscription to " + currentCookie.First + " Does Not Exist!")

			} else if rowsAffected == -2 {
				changeMessage("Error")

			} else if rowsAffected == -3 {
				changeMessage("Enter Value Into All Columns!")

			} else if rowsAffected > 1 {
				changeMessage("Subscription to " + currentCookie.First + " Renewed!")

			} else {
				changeMessage("Subscription to " + currentCookie.First + " Added!")
			}
		}

		printMessage(c)
		setDefaultMessage()

	} else {
		setDefaultMessage()
		c.Redirect(http.StatusTemporaryRedirect, "/api/login")
	}
}

func NewSubscriptionService(c *gin.Context) {
	if currentID != -1 {
		tryDefaultMessage("Enter Name and Pricing of New Subscription!")

		if !strings.Contains(currentCookie.First, "Message: ") {
			rowsAffected := MySQL.CreateNewSub(currentDB, currentCookie.First, currentCookie.Second)

			if rowsAffected == 0 {
				changeMessage("Error: Subscription to " + currentCookie.First + " Already Exists!")
			} else if rowsAffected == -1 {
				changeMessage("Error")

			} else if rowsAffected == -2 {
				changeMessage("Enter Value Into All Columns!")

			} else {
				changeMessage("Subscription to " + currentCookie.First + " Created!")
			}
		}

		printMessage(c)
		setDefaultMessage()

	} else {
		setDefaultMessage()
		c.Redirect(http.StatusTemporaryRedirect, "/api/login")
	}

}

func CancelSubscriptionService(c *gin.Context) {
	if currentID != -1 {
		tryDefaultMessage("Choose Subscription to Cancel!")

		if !strings.Contains(currentCookie.First, "Message: ") {
			rowsAffected := MySQL.CancelUserSub(currentDB, currentID, currentCookie.First)

			if rowsAffected == 0 {
				changeMessage("Subscription to " + currentCookie.First + " Does Not Exist!")
			} else if rowsAffected == -1 {
				changeMessage("Error")
			} else {
				changeMessage("Subscription to " + currentCookie.First + " Canceled!")
			}
		}

		printMessage(c)
		setDefaultMessage()

	} else {
		setDefaultMessage()
		c.Redirect(http.StatusTemporaryRedirect, "/api/login")
	}
}

func ResetDatabase(c *gin.Context) {
	if currentID == 1 {
		MySQL.ResetAllTables(currentDB)
		MySQL.SetUpTables(currentDB)
		MySQL.CreateAdminUser(currentDB)

		changeMessage("Admin Database Reset! Enter Username and Password!")
		c.Redirect(http.StatusTemporaryRedirect, "/api/login")
	} else {
		setDefaultMessage()
		c.Redirect(http.StatusTemporaryRedirect, "/api/subscriptions")
	}
}

func GetAllUserData() gin.HandlerFunc {
	return func(c *gin.Context) {
		if currentID == 1 {
			tryDefaultMessage("UserSubs")
			allUserData = []userData{}

			if strings.Contains(currentCookie.First, "Subscriptions") {
				rows, err := currentDB.Query("SELECT SubID, Name, Price FROM Subscriptions;")

				if err != nil {
					c.JSON(http.StatusBadGateway, gin.H{"message": "Error"})
				}

				for rows.Next() {
					var subid int
					var name string
					var price string

					rows.Scan(&subid, &name, &price)

					var newData userData
					newData.SubID = strconv.Itoa(subid)
					newData.Price = price
					newData.Name = name

					allUserData = append(allUserData, newData)
				}

			} else if strings.Contains(currentCookie.First, "Users") {
				rows, err := currentDB.Query("SELECT UserID, Username, Password, Email FROM Users;")

				if err != nil {
					c.JSON(http.StatusBadGateway, gin.H{"message": "Error"})
				}

				for rows.Next() {
					var id int
					var username string
					var password string
					var email string

					rows.Scan(&id, &username, &password, &email)

					var newData userData
					newData.UserID = strconv.Itoa(id)
					newData.Username = username
					newData.Password = password
					newData.Email = email

					allUserData = append(allUserData, newData)
				}
			} else {
				rows, err := currentDB.Query("SELECT UserSubs.UserID, Username, Password, Email, UserSubs.SubID, Name, Price, DateAdded, DateRemoved FROM UserSubs INNER JOIN Subscriptions ON UserSubs.SubID = Subscriptions.SubID INNER JOIN Users ON UserSubs.UserID = Users.UserID;")

				if err != nil {
					c.JSON(http.StatusBadGateway, gin.H{"message": "Error"})
				}

				for rows.Next() {
					var id int
					var username string
					var password string
					var subid int
					var name string
					var price string
					var dateadded string
					var dateremoved string
					var email string

					rows.Scan(&id, &username, &password, &email, &subid, &name, &price, &dateadded, &dateremoved)

					var newData userData
					newData.UserID = strconv.Itoa(id)
					newData.SubID = strconv.Itoa(subid)
					newData.Username = username
					newData.Password = password
					newData.DateAdded = dateadded
					newData.DateRemoved = dateremoved
					newData.Price = price
					newData.Name = name
					newData.Email = email

					allUserData = append(allUserData, newData)
				}
			}

			c.IndentedJSON(http.StatusOK, allUserData)
			setDefaultMessage()
		} else {
			setDefaultMessage()
			c.Redirect(http.StatusTemporaryRedirect, "/api/subscriptions")
		}
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
