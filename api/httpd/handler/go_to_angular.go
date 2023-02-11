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
	var login userData
	c.BindJSON(&login)

	username := login.Username
	if username == "" {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"Output": "No Username Entered!"})
		return
	}

	password := login.Password
	if password == "" {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"Output": "No Password Entered!"})
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
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"Output": "Incorrect Username or Password!"})

	} else if currentID == -502 { //server error
		currentID = -1
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"Output": "Error: Database Connection Error!"})

	} else {
		//c.JSON(http.StatusOK, gin.H{"ID": strconv.Itoa(currentID)})
		c.Redirect(http.StatusTemporaryRedirect, "/api/subscriptions")
	}
}

func NewUser(c *gin.Context) {
	//Trys to Get Cookie called postOutput
	_, err := c.Cookie("signupOutput")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"Output": "Cookie Does Not Exist!"})
		return
	}

	//Trys to Get username, password, and email
	var login userData
	c.BindJSON(&login)

	username := login.Username
	if username == "" {
		c.SetCookie("signupOutput", "Error: No Username Entered!", 60, "/", "localhost", false, false)
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"Output": "Error: No Username Entered!"})
		return
	}

	password := login.Password
	if password == "" {
		c.SetCookie("signupOutput", "Error: No Password Entered!", 60, "/", "localhost", false, false)
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"Output": "Error: No Password Entered!"})
		return
	}

	email := login.Email
	if email == "" {
		c.SetCookie("signupOutput", "Error: No Email Entered!", 60, "/", "localhost", false, false)
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"Output": "Error: No Email Entered!"})
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
		c.SetCookie("signupOutput", "Error: Username Already Exists!", 60, "/", "localhost", false, false)
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"Output": "Error: Username Already Exists!"})

	} else if rowsAffected == (-223 - 2) {
		c.SetCookie("signupOutput", "Error: Email "+email+" Already In Use!", 60, "/", "localhost", false, false)
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"Output": "Error: Email " + email + " Already In Use!"})

	} else if rowsAffected == -502 { //bad gateway
		c.SetCookie("signupOutput", "Error: Database Connection Error!", 60, "/", "localhost", false, false)
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"Output": "Error: Database Connection Error!"})

	} else if rowsAffected == -204 { //no content
		c.SetCookie("signupOutput", "Enter Value Into All Columns!", 60, "/", "localhost", false, false)
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"Output": "Enter Value Into All Columns!"})

	} else {
		c.SetCookie("signupOutput", "New User "+username+" Has Been Created! Enter Username and Password!", 60, "/", "localhost", false, false)
		c.JSON(http.StatusOK, gin.H{"Output": "New User " + username + " Has Been Created! Enter Username and Password!"})
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
				c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"message": "Error"})
			}

			var index = 0
			for rows.Next() {
				var newUserSub userData
				rows.Scan(&newUserSub.Name, &newUserSub.Price, &newUserSub.DateAdded, &newUserSub.DateRemoved)
				usersubInfo = append(usersubInfo, newUserSub)

				c.SetCookie("outputSubscriptions"+strconv.Itoa(currentID)+"-"+strconv.Itoa(index), newUserSub.Name+" "+newUserSub.Price+" "+newUserSub.DateAdded+" "+newUserSub.DateRemoved, 60*5, "/", "localhost", false, false)
				index += 1
			}

			c.IndentedJSON(http.StatusOK, usersubInfo)

		} else {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
		}
	}
}

/*
func HomePage() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Testing Redirects:

		c.Redirect(http.StatusTemporaryRedirect, "/signup") //redirect to front-end component
		//c.Redirect(http.StatusTemporaryRedirect, "/api/accountcreation")                        //redirect to back-end component
		//c.Redirect(http.StatusTemporaryRedirect, "https://www.youtube.com/watch?v=dQw4w9WgXcQ") //redirect to external website
	}
}

func Logout(message string) gin.HandlerFunc {
	return func(c *gin.Context) {
		setDefaultMessage()
		currentID = -1

		c.Redirect(http.StatusTemporaryRedirect, "/api/login")
	}
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
		combinedData, doesExist := c.GetPostForm("data")
		if !doesExist {
			combinedData = c.Param("data")
			fmt.Println("Get: " + combinedData)
		} else {
			fmt.Println("Post: " + combinedData)
		}

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
}*/

func ResetDatabase(c *gin.Context) {
	if currentID == 1 {
		MySQL.ResetAllTables(currentDB)
		MySQL.SetUpTables(currentDB)
		MySQL.CreateAdminUser(currentDB)

		//funcOutput = "Admin Database Reset! Enter Username and Password!"
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
			var username string
			var password string
			var subid int
			var name string
			var price string
			var dateadded string
			var dateremoved string
			var email string

			rows, err := currentDB.Query("SELECT SubID, Name, Price FROM Subscriptions;")

			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"message": "Error"})
			}

			for rows.Next() {
				rows.Scan(&subid, &name, &price)

				var newData userData
				newData.SubID = strconv.Itoa(subid)
				newData.Price = price
				newData.Name = name

				allUserData = append(allUserData, newData)
			}

			rows, err = currentDB.Query("SELECT UserID, Username, Password, Email FROM Users;")

			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"message": "Error"})
			}

			for rows.Next() {
				rows.Scan(&id, &username, &password, &email)

				var newData userData
				newData.UserID = strconv.Itoa(id)
				newData.Username = username
				newData.Password = password
				newData.Email = email
				allUserData = append(allUserData, newData)
			}

			rows, err = currentDB.Query("SELECT UserSubs.UserID, Username, Password, Email, UserSubs.SubID, Name, Price, DateAdded, DateRemoved FROM UserSubs INNER JOIN Subscriptions ON UserSubs.SubID = Subscriptions.SubID INNER JOIN Users ON UserSubs.UserID = Users.UserID;")

			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"message": "Error"})
			}

			for rows.Next() {
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

			c.IndentedJSON(http.StatusOK, allUserData)
		} else {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			//c.Redirect(http.StatusTemporaryRedirect, "/api/subscriptions")
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
