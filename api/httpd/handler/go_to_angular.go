package handler

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"net/http"
	"strconv"

	"github.com/gabo0802/UF-OneClick/api/httpd/handler/MySQL"
	"github.com/gin-gonic/gin"
)

// New Data Types
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

// Global Variables:
var currentDB *sql.DB
var currentID = -1

func SetDB(db *sql.DB) {
	currentDB = db
}

// GET and POST Functions:

func TryLogin(c *gin.Context) { // gin.Context parameter.
	/*_, err := c.Cookie("currentUserID")
	if err == nil {
		fmt.Println("Logged In Already!")
		c.Redirect(http.StatusTemporaryRedirect, "/api/subscriptions")
		return
	} else {
		fmt.Println("Not Logged In Already!")
	}*/

	var login userData
	c.BindJSON(&login)

	username := login.Username
	if username == "" {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"Error": "No Username Entered"})
		return
	}

	password := login.Password
	if password == "" {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"Error": "No Password Entered"})
		return
	}
	login = userData{}

	//Encrypt Username and Password
	stringEncrypter := sha256.New()
	stringEncrypter.Write([]byte(username))
	username = base64.URLEncoding.EncodeToString(stringEncrypter.Sum(nil))

	stringEncrypter = sha256.New()
	stringEncrypter.Write([]byte(password))
	password = base64.URLEncoding.EncodeToString(stringEncrypter.Sum(nil))

	//Try Login
	currentID = MySQL.Login(currentDB, username, password)

	if currentID == -401 { //unauthorized
		currentID = -1
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"Error": "Incorrect Username or Password"})

	} else if currentID == -502 { //server error
		currentID = -1
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"Error": "Database Connection Issue"})

	} else {
		//c.JSON(http.StatusOK, gin.H{"ID": strconv.Itoa(currentID)})
		c.SetCookie("currentUserID", strconv.Itoa(currentID), 60*60, "/", "localhost", false, false)
		c.Redirect(http.StatusTemporaryRedirect, "/api/subscriptions")
	}
}

func NewUser(c *gin.Context) {
	//Trys to Get Cookie called postOutput
	// _, err := c.Cookie("signupOutput")

	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusOK, gin.H{"Error": "Cookie Does Not Exist"})
	// 	return
	// }

	//Trys to Get username, password, and email
	var login userData
	c.BindJSON(&login)

	username := login.Username
	if username == "" {
		//c.SetCookie("signupOutput", "Error: No Username Entered!", 60, "/", "localhost", false, false)
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"Error": "No Username Entered"})
		return
	}

	password := login.Password
	if password == "" {
		//c.SetCookie("signupOutput", "Error: No Password Entered!", 60, "/", "localhost", false, false)
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"Error": "No Password Entered"})
		return
	}

	email := login.Email
	if email == "" {
		//c.SetCookie("signupOutput", "Error: No Email Entered!", 60, "/", "localhost", false, false)
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"Error": "No Email Entered"})
		return
	}
	login = userData{}

	//Encrypt Username and Password
	stringEncrypter := sha256.New()
	stringEncrypter.Write([]byte(username))
	encryptedusername := base64.URLEncoding.EncodeToString(stringEncrypter.Sum(nil))

	stringEncrypter = sha256.New()
	stringEncrypter.Write([]byte(password))
	password = base64.URLEncoding.EncodeToString(stringEncrypter.Sum(nil))

	//Try Create New User
	rowsAffected := MySQL.CreateNewUser(currentDB, encryptedusername, password, email)

	if rowsAffected == (-223 - 0) { //already exists
		//c.SetCookie("signupOutput", "Error: Username Already Exists!", 60, "/", "localhost", false, false)
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"Error": "Username Already Exists"})

	} else if rowsAffected == (-223 - 2) {
		//c.SetCookie("signupOutput", "Error: Email "+email+" Already In Use!", 60, "/", "localhost", false, false)
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"Error": "Email " + email + " Already In Use"})

	} else if rowsAffected == -502 { //bad gateway
		//c.SetCookie("signupOutput", "Error: Database Connection Error!", 60, "/", "localhost", false, false)
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"Error": "Database Connection Issue"})

	} else if rowsAffected == -204 { //no content
		//c.SetCookie("signupOutput", "Enter Value Into All Columns!", 60, "/", "localhost", false, false)
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"Error": "Enter Value Into All Columns"})

	} else {
		//c.SetCookie("signupOutput", "New User "+username+" Has Been Created!", 60, "/", "localhost", false, false) //maybe add " Enter Username and Password!"
		c.JSON(http.StatusOK, gin.H{"Success": "New User " + username + " Has Been Created"}) //maybe add " Enter Username and Password!"
		username = ""
	}
}

func GetAllUserSubscriptions() gin.HandlerFunc {
	return func(c *gin.Context) {
		var usersubInfo = []userData{}

		if currentID != -1 {
			rows, err := currentDB.Query("SELECT Name, Price, DateAdded, DateRemoved FROM UserSubs INNER JOIN Subscriptions ON UserSubs.SubID = Subscriptions.SubID WHERE UserID = ? ORDER BY DateAdded ASC", currentID)
			//can order by anything

			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"Error": "Database Connection Issue"})
			}

			var index = 0
			for rows.Next() {
				var newUserSub userData
				rows.Scan(&newUserSub.Name, &newUserSub.Price, &newUserSub.DateAdded, &newUserSub.DateRemoved)
				usersubInfo = append(usersubInfo, newUserSub)

				//c.SetCookie("subscriptionsOutput"+strconv.Itoa(currentID)+"-"+strconv.Itoa(index), newUserSub.Name+" "+newUserSub.Price+" "+newUserSub.DateAdded+" "+newUserSub.DateRemoved, 60*5, "/", "localhost", false, false)
				index += 1
			}

			c.IndentedJSON(http.StatusOK, usersubInfo)
			c.Redirect(http.StatusTemporaryRedirect, "/subscriptions") //change later

		} else {
			c.JSON(http.StatusOK, gin.H{"Error": "Invalid User ID"})
			//c.Redirect(http.StatusTemporaryRedirect, "/login")
		}
	}
}

func Logout(message string) gin.HandlerFunc {
	return func(c *gin.Context) {
		currentID = -1
		c.SetCookie("currentUserID", strconv.Itoa(currentID), -1, "/", "localhost", false, false)

		//c.SetCookie("logoutOutput", "Logged Out!"+message, 60, "/", "localhost", false, false)
		c.JSON(http.StatusOK, gin.H{"Success": "Logged Out" + message})
	}
}

func NewUserSubscription(c *gin.Context) {
	if currentID != -1 {
		var userSubscriptionData userData
		c.BindJSON(&userSubscriptionData)

		subscriptionName := userSubscriptionData.Name
		//subscriptionID := userSubscriptionData.ID
		userSubscriptionData = userData{}

		rowsAffected := MySQL.CreateNewUserSub(currentDB, currentID, subscriptionName)

		if rowsAffected == -223 {
			//c.SetCookie("newusersubOutput", "Subscription to "+subscriptionName+" Already Active!", 60, "/", "localhost", false, false)
			c.JSON(http.StatusOK, gin.H{"Error": "Subscription to " + subscriptionName + " Already Active"})

		} else if rowsAffected == -404 {
			//c.SetCookie("newusersubOutput", "Subscription to "+subscriptionName+" Does Not Exist!", 60, "/", "localhost", false, false)
			c.JSON(http.StatusOK, gin.H{"Error": "Subscription to " + subscriptionName + " Does Not Exist"})

		} else if rowsAffected == -502 {
			//c.SetCookie("newusersubOutput", "Error: Database Connection Error", 60, "/", "localhost", false, false)
			c.JSON(http.StatusOK, gin.H{"Error": "Database Connection Issue"})

		} else if rowsAffected == -204 {
			//c.SetCookie("newusersubOutput", "Error: Enter Value Into All Columns!", 60, "/", "localhost", false, false)
			c.JSON(http.StatusOK, gin.H{"Error": "Enter Value Into All Columns"})

		} else if rowsAffected == 223 {
			//c.SetCookie("newusersubOutput", "Subscription to "+subscriptionName+" Renewed!", 60, "/", "localhost", false, false)
			c.JSON(http.StatusOK, gin.H{"Success": "Subscription to " + subscriptionName + " Renewed"})

		} else {
			//c.SetCookie("newusersubOutput", "Subscription to "+subscriptionName+" Added!", 60, "/", "localhost", false, false)
			c.JSON(http.StatusOK, gin.H{"Success": "Subscription to " + subscriptionName + " Added"})

		}

	} else {
		//c.SetCookie("newusersubOutput", "Error: Invalid UserID", 60, "/", "localhost", false, false)
		//c.Redirect(http.StatusTemporaryRedirect, "/login")
		c.JSON(http.StatusOK, gin.H{"Error": "Invalid User ID"})
	}
}

func NewSubscriptionService(c *gin.Context) {
	if currentID != -1 {
		var subscriptionData userData
		c.BindJSON(&subscriptionData)

		subscriptionName := subscriptionData.Name
		subscriptionPrice := subscriptionData.Price
		subscriptionData = userData{}

		rowsAffected := MySQL.CreateNewSub(currentDB, subscriptionName, subscriptionPrice)

		if rowsAffected == -223 {
			//c.SetCookie("newsubOutput", "Error: Subscription Service of "+subscriptionName+" Already Exists!", 60, "/", "localhost", false, false)
			c.JSON(http.StatusOK, gin.H{"Error": "Subscription Service of " + subscriptionName + " Already Exists"})

		} else if rowsAffected == -502 {
			//c.SetCookie("newsubOutput", "Error: Database Connection Error", 60, "/", "localhost", false, false)
			c.JSON(http.StatusOK, gin.H{"Error": "Database Connection Issue"})

		} else if rowsAffected == -204 {
			//c.SetCookie("newsubOutput", "Error: Enter Value Into All Columns!", 60, "/", "localhost", false, false)
			c.JSON(http.StatusOK, gin.H{"Error": "Enter Value Into All Columns"})

		} else {
			//c.SetCookie("newsubOutput", "Subscription to "+subscriptionName+" Created!", 60, "/", "localhost", false, false)
			c.JSON(http.StatusOK, gin.H{"Success": "Subscription to " + subscriptionName + " Created"})
		}

	} else {
		//c.SetCookie("newsubOutput", "Error: Invalid UserID", 60, "/", "localhost", false, false)
		//c.Redirect(http.StatusTemporaryRedirect, "/login")

		c.JSON(http.StatusOK, gin.H{"Error": "Invalid User ID"})
	}
}

func CancelSubscriptionService(c *gin.Context) {
	if currentID != -1 {
		var userSubscriptionData userData
		c.BindJSON(&userSubscriptionData)

		subscriptionName := userSubscriptionData.Name
		//subscriptionID := userSubscriptionData.ID
		userSubscriptionData = userData{}

		rowsAffected := MySQL.CancelUserSub(currentDB, currentID, subscriptionName)

		if rowsAffected == -404 {
			//c.SetCookie("cancelsubOutput", "Subscription to "+subscriptionName+" Does Not Exist!", 60, "/", "localhost", false, false)
			c.JSON(http.StatusOK, gin.H{"Error": "Subscription to " + subscriptionName + " Does Not Exist"})

		} else if rowsAffected == -1 {
			//c.SetCookie("cancelsubOutput", "Error: Database Connection Error", 60, "/", "localhost", false, false)
			c.JSON(http.StatusOK, gin.H{"Error": " Database Connection Issue"})

		} else if rowsAffected == -204 {
			//c.SetCookie("cancelsubOutput", "Error: Enter Value Into All Columns!", 60, "/", "localhost", false, false)
			c.JSON(http.StatusOK, gin.H{"Error": "Enter Value Into All Columns"})

		} else {
			//c.SetCookie("cancelsubOutput", "Subscription to "+subscriptionName+" Canceled!", 60, "/", "localhost", false, false)
			c.JSON(http.StatusOK, gin.H{"Success": "Subscription to " + subscriptionName + " Canceled"})
		}

	} else {
		//c.SetCookie("cancelsubOutput", "Error: Invalid UserID", 60, "/", "localhost", false, false)
		c.Redirect(http.StatusTemporaryRedirect, "/api/login")
	}
}

/*
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
*/

func ResetDatabase(c *gin.Context) {
	if currentID == 1 {
		MySQL.ResetAllTables(currentDB)
		MySQL.SetUpTables(currentDB)
		MySQL.CreateAdminUser(currentDB)

		c.Redirect(http.StatusTemporaryRedirect, "/login")
	} else {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		//c.Redirect(http.StatusTemporaryRedirect, "/api/subscriptions")
	}
}

func GetAllUserData() gin.HandlerFunc {
	return func(c *gin.Context) {
		if currentID == 1 {
			var allUserData = []userData{}
			var id int
			var subid int

			rows, err := currentDB.Query("SELECT SubID, Name, Price FROM Subscriptions;")

			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"message": "Error"})
			}

			for rows.Next() {
				var newData userData
				rows.Scan(&subid, &newData.Name, &newData.Price)

				newData.SubID = strconv.Itoa(subid)

				allUserData = append(allUserData, newData)
			}

			rows, err = currentDB.Query("SELECT UserID, Username, Password, Email FROM Users;")

			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"message": "Error"})
			}

			for rows.Next() {
				var newData userData
				rows.Scan(&id, &newData.Username, &newData.Password, &newData.Email)

				newData.UserID = strconv.Itoa(id)

				allUserData = append(allUserData, newData)
			}

			rows, err = currentDB.Query("SELECT UserSubs.UserID, Username, Password, Email, UserSubs.SubID, Name, Price, DateAdded, DateRemoved FROM UserSubs INNER JOIN Subscriptions ON UserSubs.SubID = Subscriptions.SubID INNER JOIN Users ON UserSubs.UserID = Users.UserID;")

			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"message": "Error"})
			}

			for rows.Next() {
				var newData userData
				rows.Scan(&id, &newData.Username, &newData.Password, &newData.Email, &subid, &newData.Name, &newData.Price, &newData.DateAdded, &newData.DateRemoved)

				newData.UserID = strconv.Itoa(id)
				newData.SubID = strconv.Itoa(subid)

				allUserData = append(allUserData, newData)
			}

			c.IndentedJSON(http.StatusOK, allUserData)
		} else {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			//c.Redirect(http.StatusTemporaryRedirect, "/api/subscriptions")
		}
	}
}
