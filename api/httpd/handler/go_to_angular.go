package handler

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	"math/big"
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
	Timezone    string `json:"timezone"`
}
type passwordChange struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type newsLetterInfo struct {
	Message string `json:"message"`
}

type timezoneInfo struct {
	TimezoneDifference string `json:"timezonedifference"`
}

const (
	emailHost = "smtp.gmail.com"
	emailPort = "587"

	reference = "2006-01-02 15:04:05"

	oneYearInSeconds   = (365 * 24 * 60 * 60)
	oneMonthInSeconds  = 2629746
	oneWeekInSeconds   = (7 * 24 * 60 * 60)
	oneDayInSeconds    = (24 * 60 * 60)
	oneHourInSeconds   = (60 * 60)
	oneMinuteInSeconds = 60
)

// Global Variables:
var currentDB *sql.DB
var currentID = -1
var currentTimezone int = -500

func SetDB(db *sql.DB) {
	currentDB = db
}

func convert_timezone(timeString string, toUTC bool) (string, time.Time) {
	var reverse int

	if !toUTC {
		reverse = 1
	} else {
		reverse = -1
	}

	timezoneMinutes := currentTimezone
	timezoneHours := currentTimezone / 100

	timezoneMinutes -= timezoneHours * 100
	if timezoneMinutes < 0 {
		timezoneMinutes = 0
	}

	inputTime, _ := time.Parse(reference, timeString)
	convertedTime := time.Date(inputTime.Year(), inputTime.Month(), inputTime.Day(), inputTime.Hour()+(timezoneHours*reverse), inputTime.Minute()+(timezoneMinutes*reverse), inputTime.Second(), 0, time.Local)

	return string(convertedTime.String()[0 : len(convertedTime.String())-10]), convertedTime
}

func getReminderMessage(subName string, subPrice string, dateRenew string, dateAdded string) string {
	var userMessage string = ""

	if !strings.Contains(subName, "Yearly") && !strings.Contains(subName, "3 Months") {
		//fmt.Println("Monthly ", subName)

		dateRenew = strings.Replace(dateRenew, " 00:00:00", "", 1)
		userMessage = "[" + dateRenew + "] " + subName + ": $" + subPrice + "\n"

	} else if strings.Contains(subName, "Yearly") {
		dateRenewTime, _ := time.Parse(reference, dateRenew)
		dateAddedTime, _ := time.Parse(reference, dateAdded)
		dateRenew = strings.Replace(dateRenew, " 00:00:00", "", 1)

		//fmt.Println("Yearly ", subName)
		//fmt.Println(int(dateRenewTime.Month()), ",", int(dateAddedTime.Month()))

		if int(dateRenewTime.Month()) == int(dateAddedTime.Month()) {
			userMessage = "[" + dateRenew + "] " + subName + ": $" + subPrice + "\n"
		}
	} else if strings.Contains(subName, "3 Months") {
		dateRenewTime, _ := time.Parse(reference, dateRenew)
		dateAddedTime, _ := time.Parse(reference, dateAdded)
		dateRenew = strings.Replace(dateRenew, " 00:00:00", "", 1)

		//fmt.Println("3 Months ", subName)
		//fmt.Println(int(dateRenewTime.Month()), ",", int(dateAddedTime.Month()))

		if (int(dateRenewTime.Month())-int(dateAddedTime.Month()))%3 == 0 {
			userMessage = "[" + dateRenew + "] " + subName + ": $" + subPrice + "\n"
		}
	}

	return userMessage
}

func sendReminders(rows *sql.Rows, message string, header string) bool {
	var currentEmail string = ""
	var userMessage string = ""
	var emailSent bool = true

	for rows.Next() {
		var newEmail string
		var subName string
		var subPrice string
		var dateRenew string
		var dateAdded string

		rows.Scan(&newEmail, &subName, &subPrice, &dateRenew, &dateAdded)

		//fmt.Println(dateRenew, ",", dateAdded)

		if currentEmail == "" {
			currentEmail = newEmail
			userMessage = message + "\n"
			userMessage += getReminderMessage(subName, subPrice, dateRenew, dateAdded)

		} else if newEmail == currentEmail {
			userMessage += getReminderMessage(subName, subPrice, dateRenew, dateAdded)

		} else {
			emailSent = sendEmail(currentEmail, header, userMessage)
			//fmt.Println(userMessage)

			if !emailSent {
				return false
			}

			currentEmail = newEmail
			userMessage = message + "\n"
			userMessage += getReminderMessage(subName, subPrice, dateRenew, dateAdded)
		}
	}

	//fmt.Println(userMessage)
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

func SendEmailToAllUsers(emailSubject string, emailMessage string) bool {
	rows, err := currentDB.Query("SELECT Email FROM Users LEFT JOIN Verification ON Users.UserID = Verification.UserID WHERE Users.UserID > 1 AND Type IS NULL;")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var currentEmail string
		rows.Scan(&currentEmail)

		emailSent := sendEmail(currentEmail, emailSubject, emailMessage)

		if !emailSent {
			fmt.Println("Email Not Sent!")
			return false
		}
	}

	fmt.Println("Emails Sent")
	return true
}

func SendAllReminders() int {
	currentTime := time.Now()
	currentDate := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day()-1, 11, 59, 59, 0, time.Local)

	currentMonth := strconv.Itoa(int(currentTime.Month()))
	if len(currentMonth) < 2 {
		currentMonth = "0" + currentMonth
	}
	currentYear := strconv.Itoa(currentTime.Year())

	nextDayDate := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day()+1, 11, 59, 59, 0, time.Local)
	nextWeekDate := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day()+7, 11, 59, 59, 0, time.Local)

	SQLStringYearMonth := currentYear + "-" + currentMonth + "-%d 00:00:00"

	stringDate := strconv.Itoa(int(currentTime.Month())) + "/" + strconv.Itoa(int(currentTime.Day())) + "/" + strconv.Itoa(int(currentTime.Year()))

	rows, err := currentDB.Query("SELECT Email, Name, Price, DATE_FORMAT(DateAdded, ?), DateAdded FROM UserSubs INNER JOIN Subscriptions ON UserSubs.SubID = Subscriptions.SubID INNER JOIN Users ON UserSubs.UserID = Users.UserID WHERE UserSubs.UserID > 1 AND DateRemoved IS NULL AND DATE_FORMAT(DateAdded, ?) BETWEEN ? AND ? ORDER By Email, DATE_FORMAT(DateAdded, ?), UserSubs.SubID;", SQLStringYearMonth, SQLStringYearMonth, currentDate, nextDayDate, SQLStringYearMonth)
	if err != nil {
		return -502
	}

	if !sendReminders(rows, "Subscriptions to Renew", "Subscriptions to Renew "+stringDate+" (1 Day Left)") {
		return -401
	}

	currentDate = time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day()+1, 0, 0, 0, 0, time.Local)
	rows, err = currentDB.Query("SELECT Email, Name, Price, DATE_FORMAT(DateAdded, ?), DateAdded FROM UserSubs INNER JOIN Subscriptions ON UserSubs.SubID = Subscriptions.SubID INNER JOIN Users ON UserSubs.UserID = Users.UserID WHERE UserSubs.UserID > 1 AND DateRemoved IS NULL AND DATE_FORMAT(DateAdded, ?) BETWEEN ? AND ? ORDER By Email, DATE_FORMAT(DateAdded, ?), UserSubs.SubID;", SQLStringYearMonth, SQLStringYearMonth, currentDate, nextWeekDate, SQLStringYearMonth)
	if err != nil {
		return -502
	}

	if !sendReminders(rows, "Subscriptions to Renew", "Subscriptions to Renew "+stringDate+" (1 Week Left)") {
		return -401
	}

	fmt.Println("Sent Emails")
	return 1
}

// GET and POST Functions:
func DailyReminder(c *gin.Context) {
	_, err := c.Cookie("didReminder")

	if err != nil {
		errorCode := SendAllReminders()

		if errorCode == -502 {
			c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"Error": "Database Connection Issue"})
		} else if errorCode == -401 {
			c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"Error": "Emails Not Sent"})
		} else {
			c.SetCookie("didReminder", "yes", 60*60*24, "/", "localhost", false, true)
			c.JSON(http.StatusOK, gin.H{"Success": "Emails Were Sent!"})
		}
	} else {
		//c.SetCookie("didReminder", "yes", -1, "/", "localhost", false, true) //for testing
		fmt.Println("Emails Already Sent")

		c.JSON(http.StatusOK, gin.H{"Success": "Emails Already Sent!"})
	}

	//c.Redirect(http.StatusTemporaryRedirect, "/login")
}

func NewsLetter(c *gin.Context) {
	var newsMessage newsLetterInfo
	c.BindJSON(&newsMessage)

	message := newsMessage.Message
	fmt.Println(message)

	currentTime := time.Now()
	stringDate := strconv.Itoa(int(currentTime.Month())) + "/" + strconv.Itoa(int(currentTime.Day())) + "/" + strconv.Itoa(int(currentTime.Year()))

	SendEmailToAllUsers("UF-OneClick Newsletter "+stringDate, message)
	c.JSON(http.StatusOK, gin.H{"Success": "Newsletter Sent!"})
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

func startVerifyCheck(username string, userID int, email string) {
	var ID int = userID

	if ID == -1 {
		row, _ := currentDB.Query("SELECT UserID FROM Users WHERE Username = ?;", username)

		for row.Next() {
			row.Scan(&ID)
		}
	}

	codeGenerator := sha256.New()
	randomNumber, _ := rand.Int(rand.Reader, big.NewInt(900000))
	codeGenerator.Write([]byte(strconv.Itoa(int(randomNumber.Int64()) + 100000)))
	newCode := base64.URLEncoding.EncodeToString(codeGenerator.Sum(nil))

	randomNumber.Add(randomNumber, big.NewInt(100000))

	codeGenerator = sha256.New()
	codeGenerator.Write([]byte(newCode))
	newCodeEncrypted := base64.URLEncoding.EncodeToString(codeGenerator.Sum(nil))

	currentTime := time.Now()
	//expireDate := currentTime.Add(time.Second) //for testing
	expireDate := currentTime.Add(time.Hour * 24)

	currentDB.Exec("INSERT INTO Verification (UserID, Code, ExpireDate, Type) VALUES (?, ?, ?, \"vE\");", ID, newCodeEncrypted, expireDate)
	sendEmail(email, "Verify Identity", "http://localhost:4200/api/verify/"+newCode)

	//fmt.Println("http://localhost:4200/api/verify/" + newCode)
	newCode = ""
	username = ""
}

func VerifyEmail(c *gin.Context) {
	deleteUnverified()

	//Verify Current User
	currentTime := time.Now()
	possibleCode := c.Param("code")

	codeGenerator := sha256.New()
	codeGenerator.Write([]byte(possibleCode))
	possibleCodeEncrypted := base64.URLEncoding.EncodeToString(codeGenerator.Sum(nil))

	verifyUser, err := currentDB.Query("SELECT UserID FROM Verification WHERE Type = \"vE\" AND ExpireDate > ? AND Code = ?;", currentTime, possibleCodeEncrypted)

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"Error": "Database Connection Issue"})
	}

	for verifyUser.Next() {
		var ID int
		verifyUser.Scan(&ID)

		fmt.Println("Current User ID:", ID, "Verified!")
		currentDB.Exec("DELETE FROM Verification WHERE UserID = ? AND Type = \"vE\";", ID)
	}

	c.Redirect(http.StatusTemporaryRedirect, "/login")
}

func ChangeTimezone(c *gin.Context) {
	var newTimezone timezoneInfo
	c.BindJSON(&newTimezone)

	currentTimezone, _ = strconv.Atoi(newTimezone.TimezoneDifference)
}

/*func start2FA() {
	//Verify 2FA Code
	currentTime := time.Now()
	expireDate := currentTime.Add(time.Minute * 15)

	codeGenerator := sha256.New()
	randomNumber, _ := rand.Int(rand.Reader, big.NewInt(900000))
	codeGenerator.Write([]byte(strconv.Itoa(int(randomNumber.Int64()) + 100000)))
	newCode := base64.URLEncoding.EncodeToString(codeGenerator.Sum(nil))

	var userEmail string
	getEmail, _ := currentDB.Query("SELECT Email FROM Users WHERE UserID = ?;", currentID)

	for getEmail.Next() {
		getEmail.Scan(&userEmail)
	}

	randomNumber.Add(randomNumber, big.NewInt(100000))

	currentDB.Exec("INSERT INTO Verification (UserID, Code, ExpireDate, Type) VALUES (?, ?, ?, \"vL\");", currentID, newCode, expireDate)
	sendEmail(userEmail, "UF-OneClick 2FA Code", strconv.Itoa(int(randomNumber.Int64())))
}

func remove2FA() {
	currentTime := time.Now()
	currentDB.Exec("DELETE FROM Verification WHERE Type = \"vL\" AND ExpireDate < ?;", currentTime)
}

func do2FA(possibleCode string) bool {
	//Remove Old 2FA Codes
	remove2FA()

	//Verify 2FA Code
	currentTime := time.Now()
	verifyUser, err := currentDB.Query("SELECT UserID FROM Verification WHERE Type = \"vL\" AND ExpireDate > ? AND Code = ? AND UserID = ?;", currentTime, possibleCode, currentID)

	if err != nil {
		return false
	}

	for verifyUser.Next() {
		var ID int
		verifyUser.Scan(&ID)

		fmt.Println("Current User ID:", ID, "Verified!")
		currentDB.Exec("DELETE FROM Verification WHERE UserID = ? AND Type = \"vL\";", ID)

		return true
	}

	return false
}

func TwoFactorAuthentication(c *gin.Context) {
	var user2FA userData
	c.BindJSON(&user2FA)

	userCode := user2FA.Username
	//if c.Param("userCode") != "" {
	//	userCode = c.Param("userCode")
	//}

	codeGenerator := sha256.New()
	codeGenerator.Write([]byte(userCode))
	userCodeEncrypted := base64.URLEncoding.EncodeToString(codeGenerator.Sum(nil))

	didWork := do2FA(userCodeEncrypted)

	if didWork {
		//c.SetCookie("currentUserID", strconv.Itoa(currentID), 60*60, "/", "localhost", false, false)
		c.JSON(http.StatusOK, gin.H{"Success": "2FA"})
	} else {
		c.SetCookie("currentUserID", strconv.Itoa(currentID), -1, "/", "localhost", false, false)
		currentID = -1

		c.JSON(http.StatusOK, gin.H{"Error": "Invalid Code"})
		//c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": "No Username Entered"})
		return
	}

	password := login.Password
	if password == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": "No Password Entered"})
		return
	}
	login = userData{}

	//Encrypt Username and Password
	/*stringEncrypter := sha256.New()
	stringEncrypter.Write([]byte(username))
	username = base64.URLEncoding.EncodeToString(stringEncrypter.Sum(nil))*/

	stringEncrypter := sha256.New()
	stringEncrypter.Write([]byte(password))
	password = base64.URLEncoding.EncodeToString(stringEncrypter.Sum(nil))

	//Try Login
	currentID = MySQL.Login(currentDB, username, password)

	verifyUser, _ := currentDB.Query("SELECT * FROM Verification WHERE UserID = ? AND Type = \"vE\";", currentID)
	if verifyUser.Next() {
		currentID = -1
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"Error": "Unverified Username"})
		return
	}

	if currentID == -401 { //unauthorized
		currentID = -1
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"Error": "Incorrect Username or Password"})

	} else if currentID == -502 { //server error
		currentID = -1
		c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"Error": "Database Connection Issue"})

	} else {
		c.SetCookie("currentUserID", strconv.Itoa(currentID), 60*60, "/", "localhost", false, false)
		c.JSON(http.StatusOK, gin.H{"Success": "Logged In"})

		//start2FA()
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
	username = strings.Trim(username, " ")
	if username == "" {
		//c.SetCookie("signupOutput", "Error: No Username Entered!", 60, "/", "localhost", false, false)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": "No Username Entered"})
		return
	}

	password := login.Password
	password = strings.Trim(password, " ")
	if password == "" {
		//c.SetCookie("signupOutput", "Error: No Password Entered!", 60, "/", "localhost", false, false)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": "No Password Entered"})
		return
	}

	email := login.Email
	email = strings.Trim(email, " ")
	if email == "" {
		//c.SetCookie("signupOutput", "Error: No Email Entered!", 60, "/", "localhost", false, false)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": "No Email Entered"})
		return
	}
	login = userData{}

	//Encrypt Username and Password
	/*stringEncrypter := sha256.New()
	stringEncrypter.Write([]byte(username))
	encryptedusername := base64.URLEncoding.EncodeToString(stringEncrypter.Sum(nil))*/

	stringEncrypter := sha256.New()
	stringEncrypter.Write([]byte(password))
	password = base64.URLEncoding.EncodeToString(stringEncrypter.Sum(nil))

	//Try Create New User
	rowsAffected := MySQL.CreateNewUser(currentDB, username, password, email)

	if rowsAffected == (-223 - 0) { //already exists
		//c.SetCookie("signupOutput", "Error: Username Already Exists!", 60, "/", "localhost", false, false)
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"Error": "Username Already Exists"})

	} else if rowsAffected == (-223 - 2) {
		//c.SetCookie("signupOutput", "Error: Email "+email+" Already In Use!", 60, "/", "localhost", false, false)
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"Error": "Email " + email + " Already In Use"})

	} else if rowsAffected == -502 { //bad gateway
		//c.SetCookie("signupOutput", "Error: Database Connection Error!", 60, "/", "localhost", false, false)
		c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"Error": "Database Connection Issue"})

	} else if rowsAffected == -204 { //no content
		//c.SetCookie("signupOutput", "Enter Value Into All Columns!", 60, "/", "localhost", false, false)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": "Enter Value Into All Columns"})

	} else {
		//c.SetCookie("signupOutput", "New User "+username+" Has Been Created!", 60, "/", "localhost", false, false) //maybe add " Enter Username and Password!"
		c.JSON(http.StatusOK, gin.H{"Success": "New User " + username + " Has Been Created"}) //maybe add " Enter Username and Password!"

		//User Verification
		startVerifyCheck(username, -1, email)
		username = ""
	}
}

func GetAllSubscriptionServices() gin.HandlerFunc {
	return func(c *gin.Context) {
		var allSubs = []userData{}
		var subid int
		rows, err := currentDB.Query("SELECT SubID, Name, Price FROM Subscriptions;")

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"message": "Error"})
		}

		for rows.Next() {
			var newData userData
			rows.Scan(&subid, &newData.Name, &newData.Price)

			newData.SubID = strconv.Itoa(subid)

			allSubs = append(allSubs, newData)
		}

		c.JSON(http.StatusOK, allSubs)
	}
}

func GetAllCurrentUserSubscriptions(onlyActive bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("currentUserID")
		if err == nil {
			currentID, _ = strconv.Atoi(cookie)
		} else {
			currentID = -1
		}

		var usersubInfo = []userData{}

		if currentID != -1 {
			var rows *sql.Rows

			if onlyActive {
				rows, err = currentDB.Query("SELECT Name, Price, DateAdded, DateRemoved FROM UserSubs INNER JOIN Subscriptions ON UserSubs.SubID = Subscriptions.SubID WHERE UserID = ? AND DateRemoved is NULL ORDER BY DateAdded ASC", currentID)
			} else {
				rows, err = currentDB.Query("SELECT Name, Price, DateAdded, DateRemoved FROM UserSubs INNER JOIN Subscriptions ON UserSubs.SubID = Subscriptions.SubID WHERE UserID = ? ORDER BY DateAdded ASC", currentID)
			}
			//can order by anything

			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"Error": "Database Connection Issue"})
			}

			var index = 0
			for rows.Next() {
				var newUserSub userData
				rows.Scan(&newUserSub.Name, &newUserSub.Price, &newUserSub.DateAdded, &newUserSub.DateRemoved)

				newUserSub.DateAdded, _ = convert_timezone(newUserSub.DateAdded, false)
				if newUserSub.DateRemoved != "" {
					newUserSub.DateRemoved, _ = convert_timezone(newUserSub.DateRemoved, false)
				}

				usersubInfo = append(usersubInfo, newUserSub)

				//c.SetCookie("subscriptionsOutput"+strconv.Itoa(currentID)+"-"+strconv.Itoa(index), newUserSub.Name+" "+newUserSub.Price+" "+newUserSub.DateAdded+" "+newUserSub.DateRemoved, 60*5, "/", "localhost", false, false)
				index += 1
			}

			c.IndentedJSON(http.StatusOK, usersubInfo)
			//c.Redirect(http.StatusTemporaryRedirect, "/subscriptions") //change later

		} else {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid User ID"})
			//c.Redirect(http.StatusTemporaryRedirect, "/login")
		}
	}
}

func GetMostUsedUserSubscription(isContinuous bool, isActive bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if currentID != -1 {
			subName, usageTimeSeconds := MySQL.GetMostUsedSubscription(currentDB, currentID, isContinuous, isActive)
			stringYears, stringMonths, stringWeeks, stringDays, stringHours, stringMinutes, stringSeconds := "", "", "", "", "", "", ""
			seconds := usageTimeSeconds

			years := seconds / oneYearInSeconds
			if years == 1 {
				stringYears = strconv.Itoa(years) + " year "
			} else if years > 1 {
				stringYears = strconv.Itoa(years) + " years "
			}
			seconds -= years * oneYearInSeconds
			if seconds < 0 {
				seconds = 0
			}

			months := seconds / oneMonthInSeconds
			if months == 1 {
				stringMonths = strconv.Itoa(months) + " month "
			} else if months > 1 || stringYears != "" {
				stringMonths = strconv.Itoa(months) + " months "
			}
			seconds -= months * oneMonthInSeconds
			if seconds < 0 {
				seconds = 0
			}

			weeks := seconds / oneWeekInSeconds
			if weeks == 1 {
				stringWeeks = strconv.Itoa(weeks) + " week "
			} else if weeks > 1 || stringMonths != "" {
				stringWeeks = strconv.Itoa(weeks) + " weeks "
			}
			seconds -= weeks * oneWeekInSeconds
			if seconds < 0 {
				seconds = 0
			}

			days := seconds / oneDayInSeconds
			if days == 1 {
				stringDays = strconv.Itoa(days) + " day "
			} else if days > 1 || stringWeeks != "" {
				stringDays = strconv.Itoa(days) + " days "
			}
			seconds -= days * oneDayInSeconds
			if seconds < 0 {
				seconds = 0
			}

			hours := seconds / oneHourInSeconds
			if hours == 1 {
				stringHours = strconv.Itoa(hours) + " hour "
			} else if hours > 1 || stringDays != "" {
				stringHours = strconv.Itoa(hours) + " hours "
			}
			seconds -= hours * oneHourInSeconds
			if seconds < 0 {
				seconds = 0
			}

			minutes := seconds / oneMinuteInSeconds
			if minutes == 1 {
				stringMinutes = strconv.Itoa(minutes) + " minute "
			} else if minutes > 1 || stringHours != "" {
				stringMinutes = strconv.Itoa(minutes) + " minutes "
			}
			seconds -= minutes * oneMinuteInSeconds
			if seconds < 0 {
				seconds = 0
			}

			if seconds == 1 {
				stringSeconds = strconv.Itoa(seconds) + " second "
			} else {
				stringSeconds = strconv.Itoa(seconds) + " seconds "
			}

			message := "Active For: " + stringYears + stringMonths + stringWeeks + stringDays + stringHours + stringMinutes + stringSeconds
			c.JSON(http.StatusOK, gin.H{subName: message})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid User ID"})
		}
	}
}

func GetAvgPriceofAllCurrentUserSubscriptions(onlyActive bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if currentID != -1 {
			var rows *sql.Rows
			var err error
			var avgPrice string

			if onlyActive {
				rows, err = currentDB.Query("SELECT AVG(PRICE) FROM UserSubs INNER JOIN Subscriptions ON UserSubs.SubID = Subscriptions.SubID INNER JOIN Users ON UserSubs.UserID = Users.UserID WHERE UserSubs.UserID = ? AND DateRemoved IS NULL", currentID)
			} else {
				rows, err = currentDB.Query("SELECT AVG(PRICE) FROM UserSubs INNER JOIN Subscriptions ON UserSubs.SubID = Subscriptions.SubID INNER JOIN Users ON UserSubs.UserID = Users.UserID WHERE UserSubs.UserID = ?", currentID)
			}

			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"message": "Error"})
			}

			for rows.Next() {
				rows.Scan(&avgPrice)
			}

			c.JSON(http.StatusBadRequest, gin.H{"AVG Price: $": "Invalid User ID"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid User ID"})
		}
	}
}

/*func Logout(message string) gin.HandlerFunc {
	return func(c *gin.Context) {
		currentID = -1
		c.SetCookie("currentUserID", strconv.Itoa(currentID), -1, "/", "localhost", false, false)

		//c.SetCookie("logoutOutput", "Logged Out!"+message, 60, "/", "localhost", false, false)
		//c.JSON(http.StatusOK, gin.H{"Success": "Logged Out" + message})

		if c.Param("valid") != "" { //Test if GET Request or Not
			c.Redirect(http.StatusTemporaryRedirect, "/login")
		} else {
			c.JSON(http.StatusOK, gin.H{"Success": "Logged Out" + message})
		}
	}
}*/

func DeleteUser(c *gin.Context) {
	MySQL.DeleteUser(currentDB, currentID)
	c.SetCookie("currentUserID", strconv.Itoa(currentID), -1, "/", "localhost", false, false)
	currentID = -1
}

func DeleteUserSub(c *gin.Context) {
	if currentID != -1 {
		var userSubscriptionData userData
		c.BindJSON(&userSubscriptionData)

		if userSubscriptionData.DateRemoved == "" {
			userSubscriptionData.DateRemoved = "IS NULL"
		} else {
			userSubscriptionData.DateRemoved = "=" + userSubscriptionData.DateRemoved
		}

		_, err := currentDB.Exec("DELETE FROM UserSubs WHERE UserID = ? AND SubID = ? AND DateAdded = ? AND DateRemoved ? LIMIT 1;", currentID, userSubscriptionData.SubID, userSubscriptionData.DateAdded, userSubscriptionData.DateRemoved)
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"Error": "Database Connection Issue!"})
		} else {
			c.JSON(http.StatusAccepted, gin.H{"Success": "User Subscription Deleted!"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid User ID"})
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
			c.JSON(http.StatusConflict, gin.H{"Error": "Subscription to " + subscriptionName + " Already Active"})

		} else if rowsAffected == -404 {
			//c.SetCookie("newusersubOutput", "Subscription to "+subscriptionName+" Does Not Exist!", 60, "/", "localhost", false, false)
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Subscription to " + subscriptionName + " Does Not Exist"})

		} else if rowsAffected == -502 {
			//c.SetCookie("newusersubOutput", "Error: Database Connection Error", 60, "/", "localhost", false, false)
			c.JSON(http.StatusServiceUnavailable, gin.H{"Error": "Database Connection Issue"})

		} else if rowsAffected == -204 {
			//c.SetCookie("newusersubOutput", "Error: Enter Value Into All Columns!", 60, "/", "localhost", false, false)
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Enter Value Into All Columns"})

		} else if rowsAffected == 223 {
			//c.SetCookie("newusersubOutput", "Subscription to "+subscriptionName+" Renewed!", 60, "/", "localhost", false, false)
			c.JSON(http.StatusContinue, gin.H{"Success": "Subscription to " + subscriptionName + " Renewed"})

		} else {
			//c.SetCookie("newusersubOutput", "Subscription to "+subscriptionName+" Added!", 60, "/", "localhost", false, false)
			c.JSON(http.StatusCreated, gin.H{"Success": "Subscription to " + subscriptionName + " Added"})

		}

	} else {
		//c.SetCookie("newusersubOutput", "Error: Invalid UserID", 60, "/", "localhost", false, false)
		//c.Redirect(http.StatusTemporaryRedirect, "/login")
		c.JSON(http.StatusOK, gin.H{"Error": "Invalid User ID"})
	}
}

func NewPreviousUserSubscription(c *gin.Context) {
	if currentID != -1 {
		var userSubscriptionData userData
		c.BindJSON(&userSubscriptionData)

		subscriptionName := userSubscriptionData.Name
		subscriptionDateAdded := strings.Replace(userSubscriptionData.DateAdded, "-02-29", "-03-01", 1)
		subscriptionDateAdded, _ = convert_timezone(subscriptionDateAdded, true)

		subscriptionDateRemoved := userSubscriptionData.DateRemoved
		if subscriptionDateRemoved != "" {
			subscriptionDateRemoved = strings.Replace(subscriptionDateRemoved, "-02-29", "-03-01", 1)
			subscriptionDateRemoved, _ = convert_timezone(subscriptionDateRemoved, true)
		}
		//subscriptionID := userSubscriptionData.ID
		userSubscriptionData = userData{}

		//fmt.Println(subscriptionName, " ", subscriptionDateAdded, " ", subscriptionDateRemoved)

		rowsAffected := MySQL.AddOldUserSub(currentDB, currentID, subscriptionName, subscriptionDateAdded, subscriptionDateRemoved)

		if rowsAffected == -223 {
			c.JSON(http.StatusConflict, gin.H{"Error": "Subscription to " + subscriptionName + " Already Active"})

		} else if rowsAffected == -404 {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Subscription to " + subscriptionName + " Does Not Exist"})

		} else if rowsAffected == (-415 - 2) {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Date Added Not Formated Properly"})

		} else if rowsAffected == (-415 - 3) {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Date Canceled Not Formated Properly"})

		} else if rowsAffected == -502 {
			c.JSON(http.StatusServiceUnavailable, gin.H{"Error": "Database Connection Issue"})

		} else if rowsAffected == -204 {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Enter Value Into All Columns"})

		} else if rowsAffected == 223 {
			c.JSON(http.StatusContinue, gin.H{"Success": "Subscription to " + subscriptionName + " Renewed"})

		} else if rowsAffected == -401 {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Can't Cancel Subscription Before Adding It"})

		} else if rowsAffected == (-223 - 1) {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Can't Have An Active Subscription Before An Already Active Subscription"})

		} else if rowsAffected == (-223-2) || rowsAffected == (-223-4) {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Can't Have A Subscription Be Active In The Middle Of That Same Subscription"})

		} else if rowsAffected == (-223 - 3) {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Can't Have The Two of Same Subscription Be On the Same Exact Time"})

		} else if rowsAffected == (-223 - 5) {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Can't Add or Cancel Subscription In The Future"})

		} else {
			//c.SetCookie("newusersubOutput", "Subscription to "+subscriptionName+" Added!", 60, "/", "localhost", false, false)
			c.JSON(http.StatusOK, gin.H{"Success": "Subscription to " + subscriptionName + " Added"})

		}

	} else {
		//c.SetCookie("newusersubOutput", "Error: Invalid UserID", 60, "/", "localhost", false, false)
		//c.Redirect(http.StatusTemporaryRedirect, "/login")
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid User ID"})
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
			c.JSON(http.StatusConflict, gin.H{"Error": "Subscription Service of " + subscriptionName + " Already Exists"})

		} else if rowsAffected == -502 {
			//c.SetCookie("newsubOutput", "Error: Database Connection Error", 60, "/", "localhost", false, false)
			c.JSON(http.StatusServiceUnavailable, gin.H{"Error": "Database Connection Issue"})

		} else if rowsAffected == -204 {
			//c.SetCookie("newsubOutput", "Error: Enter Value Into All Columns!", 60, "/", "localhost", false, false)
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Enter Value Into All Columns"})

		} else {
			//c.SetCookie("newsubOutput", "Subscription to "+subscriptionName+" Created!", 60, "/", "localhost", false, false)
			c.JSON(http.StatusCreated, gin.H{"Success": "Subscription to " + subscriptionName + " Created"})
		}

	} else {
		//c.SetCookie("newsubOutput", "Error: Invalid UserID", 60, "/", "localhost", false, false)
		//c.Redirect(http.StatusTemporaryRedirect, "/login")

		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid User ID"})
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

		if rowsAffected == -404 || rowsAffected == 0 {
			//c.SetCookie("cancelsubOutput", "Subscription to "+subscriptionName+" Does Not Exist!", 60, "/", "localhost", false, false)
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Subscription to " + subscriptionName + " Does Not Exist"})

		} else if rowsAffected == -1 {
			//c.SetCookie("cancelsubOutput", "Error: Database Connection Error", 60, "/", "localhost", false, false)
			c.JSON(http.StatusServiceUnavailable, gin.H{"Error": " Database Connection Issue"})

		} else if rowsAffected == -204 {
			//c.SetCookie("cancelsubOutput", "Error: Enter Value Into All Columns!", 60, "/", "localhost", false, false)
			c.JSON(http.StatusBadGateway, gin.H{"Error": "Enter Value Into All Columns"})

		} else {
			//c.SetCookie("cancelsubOutput", "Subscription to "+subscriptionName+" Canceled!", 60, "/", "localhost", false, false)
			c.JSON(http.StatusAccepted, gin.H{"Success": "Subscription to " + subscriptionName + " Canceled"})
		}

	} else {
		//c.SetCookie("cancelsubOutput", "Error: Invalid UserID", 60, "/", "localhost", false, false)
		c.Redirect(http.StatusTemporaryRedirect, "/api/login")
	}
}

func ChangeUserPassword(c *gin.Context) {
	if currentID != -1 {
		var passwordInfo passwordChange
		c.BindJSON(&passwordInfo)

		oldPassword := passwordInfo.OldPassword
		newPassword := passwordInfo.NewPassword
		passwordInfo = passwordChange{}

		stringEncrypter := sha256.New()
		stringEncrypter.Write([]byte(oldPassword))
		oldPassword = base64.URLEncoding.EncodeToString(stringEncrypter.Sum(nil))

		stringEncrypter = sha256.New()
		stringEncrypter.Write([]byte(newPassword))
		newPassword = base64.URLEncoding.EncodeToString(stringEncrypter.Sum(nil))

		rowsAffected := MySQL.ChangePassword(currentDB, currentID, oldPassword, newPassword)

		if rowsAffected == 0 {
			c.JSON(http.StatusConflict, gin.H{"Error": "Incorrect Old Password"})
		} else if rowsAffected == -502 {
			c.JSON(http.StatusServiceUnavailable, gin.H{"Error": "Database Connection Issue"})
		} else if rowsAffected == -204 {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Enter Value Into All Columns"})
		} else {
			currentID = -1
			c.JSON(http.StatusOK, gin.H{"Success": "Password Changed"})
			//c.Redirect(http.StatusTemporaryRedirect, "/api/login")
		}
	} else {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
}

func ChangeUserUsername(c *gin.Context) {
	var changeInfo userData
	c.BindJSON(&changeInfo)

	rowsAffected := MySQL.ChangeUsername(currentDB, currentID, changeInfo.Username)

	if rowsAffected == (-223 - 0) { //already exists
		//c.SetCookie("signupOutput", "Error: Username Already Exists!", 60, "/", "localhost", false, false)
		c.JSON(http.StatusConflict, gin.H{"Error": "Username Already Exists"})
	}

	changeInfo = userData{}
}

func ChangeUserEmail(c *gin.Context) {
	var changeInfo userData
	c.BindJSON(&changeInfo)

	rowsAffected := MySQL.ChangeEmail(currentDB, currentID, changeInfo.Email)

	if rowsAffected == (-223 - 0) { //already exists
		//c.SetCookie("signupOutput", "Error: Username Already Exists!", 60, "/", "localhost", false, false)
		c.JSON(http.StatusConflict, gin.H{"Error": "Email Already Exists"})
	} else {
		startVerifyCheck("", currentID, changeInfo.Email)
	}

	changeInfo = userData{}
}

func GetAllCurrentUserInfo(c *gin.Context) {
	cookie, err := c.Cookie("currentUserID")
	if err == nil {
		currentID, _ = strconv.Atoi(cookie)
	} else {
		currentID = -1
	}

	if currentID != -1 {
		outputUserData := userData{}

		outputUserData.Username = MySQL.GetUsername(currentDB, currentID)
		outputUserData.Email = MySQL.GetEmail(currentDB, currentID)

		fmt.Println(currentTimezone)
		outputUserData.Timezone = strconv.Itoa(currentTimezone)

		if currentTimezone < 0 {
			outputUserData.Timezone = strings.ReplaceAll(outputUserData.Timezone, "-", "")
			for len(outputUserData.Timezone) < 4 {
				outputUserData.Timezone = "0" + outputUserData.Timezone
			}
			outputUserData.Timezone = "-" + outputUserData.Timezone
		} else {
			for len(outputUserData.Timezone) < 4 {
				outputUserData.Timezone = "0" + outputUserData.Timezone
			}

			outputUserData.Timezone = "+" + outputUserData.Timezone
		}

		c.JSON(http.StatusOK, outputUserData)
	}
}

func GetUserInfo(c *gin.Context) {
	cookie, err := c.Cookie("currentUserID")
	if err == nil {
		currentID, _ = strconv.Atoi(cookie)
	} else {
		currentID = -1
	}

	if currentID != -1 {
		username := MySQL.GetUsername(currentDB, currentID)
		email := MySQL.GetEmail(currentDB, currentID)
		c.JSON(http.StatusOK, gin.H{username: email})
	} else {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
} //remove if above function works

func GetTimezone(c *gin.Context) {
	//c.JSON(http.StatusOK, gin.H{"CurrentTimezone": currentTimezone})
	fmt.Println(currentTimezone)
	timezoneString := strconv.Itoa(currentTimezone)

	if currentTimezone < 0 {
		timezoneString = strings.ReplaceAll(timezoneString, "-", "")
		for len(timezoneString) < 4 {
			timezoneString = "0" + timezoneString
		}
		timezoneString = "-" + timezoneString
	} else {
		for len(timezoneString) < 4 {
			timezoneString = "0" + timezoneString
		}

		timezoneString = "+" + timezoneString
	}

	c.JSON(http.StatusOK, gin.H{"CurrentTimezone": timezoneString})

} //remove if above function worksf

func ResetALL(c *gin.Context) {
	if currentID == 1 {
		MySQL.ResetAllTables(currentDB)
		MySQL.SetUpTables(currentDB)
		MySQL.CreateAdminUser(currentDB)
		MySQL.CreateCommonSubscriptions(currentDB)
		c.SetCookie("didReminder", "yes", -1, "/", "localhost", false, true)
		c.SetCookie("currentUserID", strconv.Itoa(currentID), -1, "/", "localhost", false, false)

		c.JSON(http.StatusOK, gin.H{"Success": "Reset Successful"})
		//c.Redirect(http.StatusTemporaryRedirect, "/login")
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
				newData.DateAdded, _ = convert_timezone(newData.DateAdded, false)
				if newData.DateRemoved != "" {
					newData.DateRemoved, _ = convert_timezone(newData.DateRemoved, false)
				}

				allUserData = append(allUserData, newData)
			}

			c.IndentedJSON(http.StatusOK, allUserData)
		} else {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			//c.Redirect(http.StatusTemporaryRedirect, "/api/subscriptions")
		}
	}
}
