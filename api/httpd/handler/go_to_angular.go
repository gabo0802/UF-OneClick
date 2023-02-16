package handler

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gabo0802/UF-OneClick/api/httpd/handler/MySQL"
	"github.com/gin-gonic/gin"

	"net/smtp"
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

const (
	emailHost = "smtp.gmail.com"
	emailPort = "587"
)

// Global Variables:
var currentDB *sql.DB
var currentID = -1

func SetDB(db *sql.DB) {
	currentDB = db
}

func sendReminders(rows *sql.Rows, message string, header string) bool {
	fmt.Println(header)

	var currentEmail string = ""
	var userMessage string = ""
	var emailSent bool = true

	for rows.Next() {
		var newEmail string
		var subName string
		var subPrice string

		rows.Scan(&newEmail, &subName, &subPrice)

		if currentEmail == "" {
			currentEmail = newEmail
			userMessage = message + "\n" + subName + ": $" + subPrice + "\n"

		} else if newEmail == currentEmail {
			userMessage += subName + ": $" + subPrice + "\n"

		} else {
			emailSent = sendEmail(currentEmail, header, userMessage)
			fmt.Println(userMessage)

			if !emailSent {
				return false
			}

			currentEmail = newEmail
			userMessage = message + "/n" + subName + ": $" + subPrice + "\n"
		}
	}

	fmt.Println(userMessage)
	if currentEmail != "" {
		emailSent = sendEmail(currentEmail, header, userMessage)
	}

	return emailSent
}

func sendEmail(toEmail string, emailSubject string, emailMessage string) bool {
	//Get Email from Database
	rows, err := currentDB.Query("SELECT EMAIL FROM USERS WHERE UserID = 1;")

	if err != nil {
		fmt.Println("Database Connection Issue")
		return false
	}

	var companyEmail string
	for rows.Next() {
		rows.Scan(&companyEmail)
	}

	//Get Password From .txt file
	code, missing := os.ReadFile("EmailCode.txt")
	if missing != nil {
		fmt.Println("Missing EmailCode.txt file")
		return false
	}
	emailSignInCode := string(code)
	emailSignInCode = strings.ReplaceAll(emailSignInCode, "\n", "")

	//Try to Send Email
	companyEmailAuthentication := smtp.PlainAuth("", companyEmail, emailSignInCode, emailHost)

	to := []string{toEmail}

	fullEmail := []byte("To: " + toEmail + "\r\n" +

		"Subject: " + emailSubject + "\r\n" +

		"\r\n" +

		emailMessage + "\r\n")

	err = smtp.SendMail(emailHost+":"+emailPort, companyEmailAuthentication, companyEmail, to, fullEmail)

	if err != nil {
		fmt.Println("Email: " + err.Error())
		return false
	}

	return true
}

func SendEmailToAllUsers(emailSubject string, emailMessage string) {
	rows, err := currentDB.Query("SELECT Email FROM Users WHERE UserID > 1;")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var currentEmail string
		rows.Scan(&currentEmail)

		emailSent := sendEmail(currentEmail, emailSubject, emailMessage)

		if !emailSent {
			fmt.Println("Email Not Sent!")
			return
		}
	}
}

// GET and POST Functions:
func DailyReminder(c *gin.Context) {
	_, err := c.Cookie("didReminder")

	if err != nil {
		currentTime := time.Now()
		currentDate := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, time.Local)

		currentMonth := strconv.Itoa(int(currentDate.Month()))
		if len(currentMonth) < 2 {
			currentMonth = "0" + currentMonth
		}
		currentYear := strconv.Itoa(currentDate.Year())

		nextDayDate := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day()+2, 0, 0, 0, 0, time.Local)
		nextWeekDate := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day()+8, 0, 0, 0, 0, time.Local)

		SQLString := currentYear + "-" + currentMonth + "-%d"
		stringDate := strconv.Itoa(int(currentTime.Month())) + "/" + strconv.Itoa(int(currentTime.Day())) + "/" + strconv.Itoa(int(currentTime.Year()))

		rows, err := currentDB.Query("SELECT Email, Name, Price FROM UserSubs INNER JOIN Subscriptions ON UserSubs.SubID = Subscriptions.SubID INNER JOIN Users ON UserSubs.UserID = Users.UserID WHERE UserSubs.UserID > 1 AND DateRemoved IS NULL AND DATE_FORMAT(DateAdded, ?) BETWEEN ? AND ? ORDER By Email;", SQLString, currentDate, nextDayDate)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"Error": "Database Connection Issue"})
			return
		}

		if !sendReminders(rows, "Subscriptions to Renew", "Subscriptions to Renew "+stringDate+" (1 Day Left)") {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"Error": "Emails Not Sent"})
			return
		}

		currentDate = currentDate.Add(time.Hour * 24)
		rows, err = currentDB.Query("SELECT Email, Name, Price FROM UserSubs INNER JOIN Subscriptions ON UserSubs.SubID = Subscriptions.SubID INNER JOIN Users ON UserSubs.UserID = Users.UserID WHERE UserSubs.UserID > 1 AND DateRemoved IS NULL AND DATE_FORMAT(DateAdded, ?) BETWEEN ? AND ? ORDER By Email;", SQLString, currentDate, nextWeekDate)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"Error": "Database Connection Issue"})
			return
		}

		if !sendReminders(rows, "Subscriptions to Renew", "Subscriptions to Renew "+stringDate+" (1 Week Left)") {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"Error": "Emails Not Sent"})
			return
		}

		fmt.Println("Sent Emails")
		c.SetCookie("didReminder", "yes", 60*60*24, "/", "localhost", false, true)

		c.JSON(http.StatusOK, gin.H{"Success": "Emails Were Sent!"})
	} else {
		//c.SetCookie("didReminder", "yes", -1, "/", "localhost", false, true) //for testing
		fmt.Println("Emails Already Sent")

		c.JSON(http.StatusOK, gin.H{"Success": "Emails Already Sent!"})
	}

	//c.Redirect(http.StatusTemporaryRedirect, "/login")
}

func deleteUnverified() {
	currentTime := time.Now()
	deleteUser, err := currentDB.Query("SELECT UserID FROM Verification WHERE Type = \"vE\" AND ExpireDate < ?;", currentTime)

	if err != nil {
		return
	}

	for deleteUser.Next() {
		var ID int
		deleteUser.Scan(&ID)

		fmt.Println("Current User ID:", ID, "Deleted!")
		MySQL.DeleteUser(currentDB, ID)
	}
}

func startVerifyCheck(username string, email string) {
	var ID int
	row, _ := currentDB.Query("SELECT UserID FROM Users WHERE Username = ?;", username)

	for row.Next() {
		row.Scan(&ID)
	}

	codeGenerator := sha256.New()
	codeGenerator.Write([]byte(strconv.Itoa(ID)))
	newCode := base64.URLEncoding.EncodeToString(codeGenerator.Sum(nil))

	currentTime := time.Now()
	//expireDate := currentTime.Add(time.Second) //for testing
	expireDate := currentTime.Add(time.Hour * 24)

	currentDB.Exec("INSERT INTO Verification (UserID, Code, ExpireDate, Type) VALUES (?, ?, ?, \"vE\");", ID, newCode, expireDate)
	sendEmail(email, "Verify Identity", "http://localhost:4200/api/verify/"+newCode)

	fmt.Println("http://localhost:4200/api/verify/" + newCode)
	username = ""
}

func VerifyEmail(c *gin.Context) {
	deleteUnverified()

	//Verify Current User
	currentTime := time.Now()
	possibleCode := c.Param("code")
	possibleCode = strings.ReplaceAll(possibleCode, "=", "")

	if len(possibleCode) != 43 || strings.Contains(possibleCode, "/*") || strings.Contains(possibleCode, " ") || strings.Contains(possibleCode, ";") {
		fmt.Println("Possible SQL Injection")
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}

	possibleCode = possibleCode + "="
	verifyUser, err := currentDB.Query("SELECT UserID FROM Verification WHERE Type = \"vE\" AND ExpireDate > ? AND Code = ?;", currentTime, possibleCode)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"Error": "Database Connection Issue"})
	}

	for verifyUser.Next() {
		var ID int
		verifyUser.Scan(&ID)

		fmt.Println("Current User ID:", ID, "Verified!")
		currentDB.Query("DELETE FROM Verification WHERE UserID = ? AND Type = \"vE\";", ID)
	}

	c.Redirect(http.StatusTemporaryRedirect, "/login")
}

/*func start2FA(userID int) {
	//Verify 2FA Code
	currentTime := time.Now()
	expireDate := currentTime.Add(time.Minute * 15)

	newCode := 1
	currentDB.Exec("INSERT INTO Verification (UserID, Code, ExpireDate, Type) VALUES (?, ?, ?, \"vL\");", userID, newCode, expireDate)
}

func do2FA(possibleCode string, userID int) bool {
	//Verify 2FA Code
	currentTime := time.Now()
	verifyUser, err := currentDB.Query("SELECT UserID FROM Verification WHERE Type = \"vL\" AND ExpireDate > ? AND Code = ? AND UserID = ?;", currentTime, possibleCode, userID)

	if err != nil {
		return false
	}

	for verifyUser.Next() {
		var ID int
		verifyUser.Scan(&ID)

		fmt.Println("Current User ID:", ID, "Verified!")
		currentDB.Query("DELETE FROM Verification WHERE UserID = ? AND Type = \"vL\";", ID)

		return true
	}

	return false
}

func TwoFactorAuthentication(c *gin.Context) {

}*/

func TryLogin(c *gin.Context) { // gin.Context parameter.
	/*_, err := c.Cookie("currentUserID")
	if err == nil {
		fmt.Println("Logged In Already!")
		c.Redirect(http.StatusTemporaryRedirect, "/api/subscriptions")
		return
	} else {
		fmt.Println("Not Logged In Already!")
	}*/
	deleteUnverified()

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

	verifyUser, _ := currentDB.Query("SELECT * FROM Verification WHERE UserID = ? AND Type = \"vE\";", currentID)
	if verifyUser.Next() {
		currentID = -1
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"Error": "Unverified Username"})
		return
	}

	if currentID == -401 { //unauthorized
		currentID = -1
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"Error": "Incorrect Username or Password"})

	} else if currentID == -502 { //server error
		currentID = -1
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"Error": "Database Connection Issue"})

	} else {
		//c.JSON(http.StatusOK, gin.H{"ID": strconv.Itoa(currentID)})
		c.SetCookie("currentUserID", strconv.Itoa(currentID), 60*60, "/", "localhost", false, false)
		//c.Redirect(http.StatusTemporaryRedirect, "/api/subscriptions")
		c.JSON(http.StatusOK, gin.H{"Success": "Logged In"})
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

		//User Verification
		startVerifyCheck(encryptedusername, email)
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

/*func ChangeUserPassword(c *gin.Context) {
	if currentID != -1{
		var userInfo userData
		c.BindJSON(&userInfo)

		oldPassword	:= userInfo.Username
		newPassword := userInfo.Password
		userInfo = userData{}

		stringEncrypter := sha256.New()
		stringEncrypter.Write([]byte(oldPassword))
		oldPassword = base64.URLEncoding.EncodeToString(stringEncrypter.Sum(nil))

		stringEncrypter = sha256.New()
		stringEncrypter.Write([]byte(newPassword))
		newPassword = base64.URLEncoding.EncodeToString(stringEncrypter.Sum(nil))

		rowsAffected := MySQL.ChangePassword(currentDB, currentID, oldPassword, newPassword)

		if rowsAffected == 0 {
			c.JSON(http.StatusOK, gin.H{"Error": "Incorrect Old Password"})
		} else if rowsAffected == -502 {
			c.JSON(http.StatusOK, gin.H{"Error": "Database Connection Issue"})
		} else if rowsAffected == -204 {
			c.JSON(http.StatusOK, gin.H{"Error": "Enter Value Into All Columns"})
		} else {
			currentID = -1
			c.JSON(http.StatusOK, gin.H{"Success": "Password Changed"})
			//c.Redirect(http.StatusTemporaryRedirect, "/api/login")
		}
	}else{
		c.Redirect(http.StatusTemporaryRedirect, "/api/login")
	}
}*/

func resetCookies(c *gin.Context) {
	c.SetCookie("didReminder", "yes", -1, "/", "localhost", false, true)
	c.SetCookie("currentUserID", strconv.Itoa(currentID), -1, "/", "localhost", false, false)
}

func ResetALL(c *gin.Context) {
	if currentID == 1 {
		MySQL.ResetAllTables(currentDB)
		MySQL.SetUpTables(currentDB)
		MySQL.CreateAdminUser(currentDB)
		MySQL.CreateCommonSubscriptions(currentDB)
		resetCookies(c)

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
