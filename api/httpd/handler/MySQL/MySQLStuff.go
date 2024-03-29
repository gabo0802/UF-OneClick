package MySQL

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host         = "oneclickserver.mysql.database.azure.com"
	database     = "userdb"
	user         = "adminUser"
	companyemail = "vanbestindustries@gmail.com"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func openPasswordFile() string {
	code, missing := os.ReadFile("MySQLPassword.txt")
	checkError(missing)
	mySQLPass := string(code)
	mySQLPass = strings.ReplaceAll(mySQLPass, "\n", "")
	return mySQLPass
}

func MySQLConnect() *sql.DB {
	var db *sql.DB

	//Get Password From .txt file
	mySQLPass := openPasswordFile()

	//Connect to remote server using Microsoft Azure
	// Initialize connection string
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, mySQLPass, host, database)

	// Initialize connection object
	db, err := sql.Open("mysql", connectionString)
	checkError(err)

	err = db.Ping()
	checkError(err)
	//fmt.Println("Successfully created connection to database.")
	//Function code based on https://learn.microsoft.com/en-us/azure/mysql/single-server/connect-go

	return db
}

func GetDatabaseSize(db *sql.DB) int {
	var tableCount int
	err := db.QueryRow("SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = 'userdb'").Scan(&tableCount)
	checkError(err)

	return tableCount
}

func SetUpTables(db *sql.DB) {
	//Users
	db.Exec("CREATE TABLE IF NOT EXISTS Users (UserID int NOT NULL AUTO_INCREMENT, Email varchar(255) NOT NULL, Username varchar(255) NOT NULL, Password varchar(255) NOT NULL, UNIQUE(Username), UNIQUE(Email), PRIMARY KEY(UserID));")

	//All available subscription services
	db.Exec("CREATE TABLE IF NOT EXISTS Subscriptions (SubID int NOT NULL AUTO_INCREMENT, Name varchar(255) NOT NULL, Price float NOT NULL, UNIQUE(Name), PRIMARY KEY(SubID));")

	//Individual user subscriptions
	db.Exec("CREATE TABLE IF NOT EXISTS UserSubs (UserSubID int NOT NULL AUTO_INCREMENT, UserID int NOT NULL, SubID int NOT NULL, DateAdded DATETIME NOT NULL, DateRemoved DATETIME, PRIMARY KEY(UserSubID), FOREIGN KEY(UserID) REFERENCES Users(UserID), FOREIGN KEY(SubID) REFERENCES Subscriptions(SubID))")

	//Email Verification and Two-Factor Authentication
	db.Exec("CREATE TABLE IF NOT EXISTS Verification (UserID int NOT NULL, ExpireDate DATETIME NOT NULL, Code varchar(255) NOT NULL, Type varchar(255) NOT NULL, FOREIGN KEY(UserID) REFERENCES Users(UserID))")
}

func GetTableSize(db *sql.DB, tableName string) int {
	sqlCode := "SELECT * FROM " + tableName + ";"

	rows, err := db.Query(sqlCode)
	size := 0

	if err != nil {
		fmt.Printf("Error: Table \"%v\" Does Not Exist!\n", tableName)
		return -404
	}

	for rows.Next() {
		size += 1
	}

	return size
}

func ResetTable(db *sql.DB, tableName string) {
	db.Exec("DROP TABLE IF EXISTS " + tableName)
}

func ResetAllTables(db *sql.DB) {
	//Must drop UserSubs first since its foreign keys depend on the other tables for primary keys
	db.Exec("DROP TABLE IF EXISTS Verification;")
	db.Exec("DROP TABLE IF EXISTS UserSubs;")
	db.Exec("DROP TABLE IF EXISTS Users;")
	db.Exec("DROP TABLE IF EXISTS Subscriptions;")

	//db.Exec("DELETE FROM Users WHERE UserID > 1;")
	//db.Exec("DELETE FROM Subscriptions WHERE SubID > 1;")
}

func CreateNewUser(db *sql.DB, username string, password string, email string) int {
	//Create New User

	if username == "" || password == "" || email == "" {
		return -204 //no content
	}

	result, err := db.Exec("INSERT INTO Users(Username, Password, Email) VALUES (?,?,?);", username, password, email)

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			if strings.Contains(err.Error(), "mail") {
				fmt.Println("Email Already Exists!")
				return (-223 - 2) //already exists (third variable)
			} else {
				fmt.Println("Username Already Exists!")
				fmt.Println(err.Error())
				return (-223 - 0) //already exists (first variable)
			}
		} else {
			log.Fatal(err)
			return -502 //server error
		}
	}

	numRows, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("Rows Affected:", numRows)
	//Test If User Creation Worked

	return int(numRows)
}

func CreateAdminUser(db *sql.DB) {
	//result, err := db.Exec("INSERT INTO Users(UserID, Username, Password, Email) VALUES (1, \"SBNJTRN-FjG7owHVrKtue7eqdM4RhdRWVl71HXN2d7I=\", \"XohImNooBHFR0OVvjcYpJ3NgPQ1qq73WKhHvch0VQtg=\", ?);", companyemail)
	_, err := db.Exec("INSERT INTO Users(UserID, Username, Password, Email) VALUES (1, \"root\", \"XohImNooBHFR0OVvjcYpJ3NgPQ1qq73WKhHvch0VQtg=\", ?);", companyemail)
	//maybe change password to something more secure?

	if err != nil {
		log.Fatal(err)
	}

	//commented out so they don't show on unit tests
	//uncomment to test and show rows affected
	/*numRows, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Rows Affected:", numRows)*/
}

func CreateTestUser(db *sql.DB) {
	_, err := db.Exec("INSERT INTO Users(UserID, Username, Password, Email) VALUES (2, \"test\", \"XohImNooBHFR0OVvjcYpJ3NgPQ1qq73WKhHvch0VQtg=\", ?);", "sir.testmctestington.the.tester2@gmail.com")
	if err != nil {
		log.Fatal(err)
	}

	AddOldUserSub(db, 2, "Disney+ (Basic)", "2023-02-15 01:18:56", "")
	AddOldUserSub(db, 2, "Playstation Plus (Essential) (3 Months)", "2020-12-05 04:13:44", "")
	AddOldUserSub(db, 2, "AMC+ (Yearly)", "2020-03-01 11:42:33", "")
	AddOldUserSub(db, 2, "Hulu (Student)", "2022-02-01 09:28:33", "2023-01-01 11:48:53")
	AddOldUserSub(db, 2, "Amazon Prime (Student) (Yearly)", "2022-04-01 11:53:33", "")
}

func CreateCommonSubscriptions(db *sql.DB) {
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Netflix (Basic with ads)\", 6.99);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Netflix (Basic)\", 9.99);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Netflix (Standard)\", 15.49);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Netflix (Premium)\", 19.99);")

	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Amazon Prime\", 14.99);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Amazon Prime (Yearly)\", 139);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Amazon Prime (Student)\", 7.49);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Amazon Prime (Student) (Yearly)\", 69);")

	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Prime Video\", 8.99);")

	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Disney+ (Basic)\", 7.99);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Disney+ (Premium)\", 10.99);")

	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Hulu\", 7.99);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Hulu (Student)\", 1.99);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Hulu (No Ads)\", 14.99);")

	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"ESPN+\", 9.99);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"ESPN+ (Yearly)\", 99.99);")

	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Disney Bundle Duo Basic\", 9.99);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Disney Bundle Trio Basic\", 12.99);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Disney Bundle Trio Premium\", 19.99);")

	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"HBO Max (With ADS)\", 9.99);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"HBO Max (AD-Free)\", 15.99);")

	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Playstation Plus (Essential)\", 9.99);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Playstation Plus (Essential) (3 Months)\", 24.99);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Playstation Plus (Essential) (Yearly)\", 59.99);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Playstation Plus (Extra)\", 14.99);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Playstation Plus (Extra) (3 Months)\", 39.99);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Playstation Plus (Extra) (Yearly)\", 99.99);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Playstation Plus (Premium)\", 17.99);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Playstation Plus (Premium) (3 Months)\", 49.99);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Playstation Plus (Premium) (Yearly)\", 119.99);")

	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"XBOX Live Gold\", 9.99);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"XBOX Live Gold (3 Months)\", 24.99);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"XBOX Live Gold (Yearly)\", 59.99);")

	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"XBOX Game Pass (PC)\", 9.99);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"XBOX Game Pass (PC)\", 9.99);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"XBOX Game Pass (Console)\", 9.99);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"XBOX Game Pass (Ultimate)\", 14.99);")

	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Spotify Premium (Individual)\", 9.99);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Spotify Premium (Duo)\", 12.99);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Spotify Premium (Family)\", 15.99);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Spotify Premium (Student)\", 4.99);")

	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Apple Music (Voice)\", 4.99);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Apple Music (Student)\", 5.99);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Apple Music (Individual)\", 10.99);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"Apple Music (Family)\", 16.99);")

	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"AMC+\", 8.99);")
	db.Exec("INSERT INTO Subscriptions(name, price) VALUES (\"AMC+ (Yearly)\", 83.88);")
}

func GetPassword(db *sql.DB, userID int) string {
	var password string
	err := db.QueryRow("SELECT Password FROM Users WHERE UserId = ?", userID).Scan(&password)
	checkError(err)

	return password
}

func ChangePassword(db *sql.DB, userID int, oldPassword string, newPassword string) int {
	if oldPassword == "" || newPassword == "" {
		return -204
	}

	result, err := db.Exec("UPDATE Users SET Password = ? WHERE userID = ? AND Password = ?;", newPassword, userID, oldPassword)
	if err != nil {
		return -502
	}

	numRows, err := result.RowsAffected()

	if err != nil {
		return -502
	}

	return int(numRows)
}

func GetEmail(db *sql.DB, userID int) string {
	rows, err := db.Query("SELECT Email FROM Users WHERE userID = ?;", userID)
	var email string = "none"

	if err != nil {
		return "Error Code: -502"
	}

	for rows.Next() {
		rows.Scan(&email)
	}

	if err != nil {
		return "Error Code: -502"
	}

	return email
}

func ChangeEmail(db *sql.DB, userID int, newEmail string) int {
	result, err := db.Exec("UPDATE Users SET Email = ? WHERE userID = ?;", newEmail, userID)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return -223
		} else {
			return -502
		}
	}

	numRows, err := result.RowsAffected()

	if err != nil {
		return -502
	}

	return int(numRows)
}

func GetUsername(db *sql.DB, userID int) string {
	rows, err := db.Query("SELECT Username FROM Users WHERE userID = ?;", userID)
	var username string = "none"

	if err != nil {
		return "Error Code: -502"
	}

	for rows.Next() {
		rows.Scan(&username)
	}

	if err != nil {
		return "Error Code: -502"
	}

	return username
}

func ChangeUsername(db *sql.DB, userID int, newUsername string) int {
	result, err := db.Exec("UPDATE Users SET Username = ? WHERE userID = ?;", newUsername, userID)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return -223
		} else {
			return -502
		}
	}

	numRows, err := result.RowsAffected()

	if err != nil {
		return -502
	}

	return int(numRows)
}

func CreateNewSub(db *sql.DB, name string, price string) int {
	//Create New Subscription
	name = strings.Trim(name, " ")
	if name == "" || price == "" {
		return -204
	}

	realPrice, _ := strconv.ParseFloat(price, 32)
	result, err := db.Exec("INSERT INTO Subscriptions(name, price) VALUES (?,?);", name, realPrice)

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			fmt.Println("Subscription Name Already Exists!")
			return -223
		} else {
			log.Fatal(err)
			return -502
		}
	}

	numRows, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
		return -502
	}

	//fmt.Println("Rows Affected:", numRows)
	return int(numRows)
}

func canAddUserSub(db *sql.DB, userID int, subID int) int {
	rows, err := db.Query("SELECT DateRemoved FROM UserSubs WHERE UserID = ? AND SubID = ? ORDER BY DateRemoved;", userID, subID)

	if err != nil {
		log.Fatal(err)
	}

	if rows.Next() {
		var currentDateRemoved string
		rows.Scan(&currentDateRemoved)

		if currentDateRemoved == "" { //tests if the subscription has been canceled (DateRemoved = nil)
			return -401 //if not then subscription still exists
		} else {
			return (223 - 1) //subscription can be renewed (added again), 1 will be added to it in CreateNewUserSubFunction
		}
	}

	return 0 //first time adding new subscription
}

// Adds based on the current time
func CreateNewUserSub(db *sql.DB, userID int, subscriptionName string) int {
	subscriptionName = strings.Trim(subscriptionName, " ")
	if subscriptionName == "" {
		return -204
	}

	//Gets the current time and formats it into YYYY-MM-DD hh:mm:ss
	currentTime := time.Now()
	if int(currentTime.Month()) == 2 && currentTime.Day() == 29 { //No Leap Days
		currentTime = time.Date(currentTime.Year(), time.March, 1, currentTime.Hour(), currentTime.Minute(), currentTime.Second(), 0, time.Local)
	}
	currentTime.Format("2006-01-02 15:04:05")

	var CurrentSubID int

	//Gets the SubID from Subscriptions table
	sub_name, err := db.Query("SELECT SubID FROM Subscriptions WHERE Name = ?;", subscriptionName)

	if err != nil {
		log.Fatal(err)
		return -502
	}

	//Checks If Query Returns Empty Set or if the Subscription Name exists
	if sub_name.Next() {
		//Gets the SubID
		sub_name.Scan(&CurrentSubID)
		//fmt.Println("Sub ID:", CurrentSubID)

	} else {
		fmt.Println("Subscription Name is Invalid")
		return -404
	}

	//Checks to see if sub was already added to user before creating a new table value
	var isRenewed int = canAddUserSub(db, userID, CurrentSubID)
	if isRenewed < 0 {
		fmt.Println("Subscription already added to User's Profile!")
		return -223
	}

	//Create New UserSub Data
	result, _ := db.Exec("INSERT INTO UserSubs(UserID, SubID, DateAdded) VALUES (?,?,?);", userID, CurrentSubID, currentTime)

	//Tests to see if function worked
	numRows, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
		return -402
	}

	//fmt.Println("Rows Affected:", numRows)
	return int(numRows) + isRenewed
}

func canAddOldUserSub(db *sql.DB, userID int, subID int, oldDate string, oldCanceledDate string) int {
	rows, err := db.Query("SELECT DateAdded, DateRemoved FROM UserSubs WHERE UserID = ? AND SubID = ? AND DateAdded > ?;", userID, subID, oldDate) //sub in future already exists (need to have old sub be canceled)
	if err != nil {
		log.Fatal(err)
	}
	if rows.Next() && (oldCanceledDate == "0001-01-01 00:00:00 +0000 UTC" || oldCanceledDate == "") {
		//fmt.Println("Error: Can't Have Two of the Same Subscription At The Same Time")
		return (-223 - 1)
	}

	rows, err = db.Query("SELECT DateAdded, DateRemoved FROM UserSubs WHERE UserID = ? AND SubID = ? AND DateAdded < ? AND DateRemoved > ?;", userID, subID, oldDate, oldDate)
	if err != nil {
		log.Fatal(err)
	}
	if rows.Next() {
		//fmt.Println("Error: Can't Have The Same Subscription In The Middle Of A Subscription")
		return (-223 - 2)
	}

	rows, err = db.Query("SELECT DateAdded, DateRemoved FROM UserSubs WHERE UserID = ? AND SubID = ? AND DateAdded = ? ;", userID, subID, oldDate)
	if err != nil {
		log.Fatal(err)
	}
	if rows.Next() {
		//fmt.Println("Error: Can't Have The Same Subscription Be On the Same Exact Time")
		return (-223 - 3)
	}

	if !(oldCanceledDate == "0001-01-01 00:00:00 +0000 UTC" || oldCanceledDate == "") {
		rows, err = db.Query("SELECT DateAdded FROM UserSubs WHERE UserID = ? AND SubID = ? AND DateAdded < ? AND DateAdded > ?;", userID, subID, oldCanceledDate, oldDate)

		if err != nil {
			log.Fatal(err)
		}

		if rows.Next() {
			//fmt.Println("Error: Can't Have The Same Subscription Be In The Middle Of That Same Subscription")
			return (-223 - 4)
		}
	}

	rows, err = db.Query("SELECT DateRemoved FROM UserSubs WHERE UserID = ? AND SubID = ? AND DateAdded < ?;", userID, subID, oldDate)

	if err != nil {
		log.Fatal(err)
	}

	if rows.Next() {
		var currentDateRemoved string
		rows.Scan(&currentDateRemoved)

		if currentDateRemoved == "" { //tests if the subscription has been canceled (DateRemoved = nil)
			fmt.Println("Error: Old Un-Canceled Subscription Still Exists!")
			return (-223) //if not then subscription still exists
		} else {
			return 1 //can add new subscription
		}
	}

	return 0 //doesn't exist yet
}

// Allows pre-existing subscriptions with dates other than the current time to be added
func AddOldUserSub(db *sql.DB, userID int, subscriptionName string, dateAdded string, dateCanceled string) int {
	const reference = "2006-01-02 15:04:05"
	var dateCanceledTime time.Time

	dateAddedTime, err := time.Parse(reference, dateAdded)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error: Date Added Not Formatted Properly")
		return (-415 - 2)
	}

	if dateCanceled != "" {
		dateCanceledTime, err = time.Parse(reference, dateCanceled)
		if err != nil {
			log.Fatal(err)
			fmt.Println("Error: Date Canceled Not Formatted Properly")
			return (-415 - 3)
		}
	}

	currentTime := time.Now().UTC()
	if int(currentTime.Month()) == 2 && currentTime.Day() == 29 {
		currentTime = time.Date(currentTime.Year(), time.March, 1, currentTime.Hour(), currentTime.Minute(), currentTime.Second(), 0, time.Local) //No Leap Days
	}

	//println(currentTime.String())
	//println(dateAddedTime.String())
	//println(dateCanceledTime.String())

	if dateCanceledTime.After(currentTime) || dateAddedTime.After(currentTime) {
		//fmt.Println("Error: Can't Add or Cancel Subscription in Future (After Now)!")
		return (-223 - 5)
	}

	if !dateCanceledTime.IsZero() && (dateCanceledTime.Before(dateAddedTime) || dateCanceledTime.Equal(dateAddedTime)) {
		//fmt.Println("Error: Can't Cancel Subscription Before Adding It")
		return -401
	}

	if subscriptionName == "" || dateAdded == "" {
		return -204
	}

	var CurrentSubID int

	//Gets the SubID from Subscriptions table
	sub_name, err := db.Query("SELECT SubID FROM Subscriptions WHERE Name = ?;", subscriptionName)

	if err != nil {
		log.Fatal(err)
		return -502
	}

	//Checks If Query Returns Empty Set or if the Subscription Name exists
	if sub_name.Next() {
		//Gets the SubID
		sub_name.Scan(&CurrentSubID)
		//fmt.Println("Sub ID:", CurrentSubID)

	} else {
		fmt.Println("Subscription Name is Invalid")
		return -404
	}

	//Checks to see if sub was already added to user before creating a new table value (night be irrelevent, future issue)
	var isRenewed int = canAddOldUserSub(db, userID, CurrentSubID, dateAddedTime.String(), dateCanceledTime.String())
	if isRenewed < 0 {
		//fmt.Println("Subscription already added to User's Profile!")
		//fmt.Println(dateAdded + ", " + dateCanceled)
		return isRenewed
	}

	//Create New UserSub Data
	var result sql.Result
	if !dateCanceledTime.IsZero() {
		result, _ = db.Exec("INSERT INTO UserSubs(UserID, SubID, DateAdded, DateRemoved) VALUES (?,?,?,?);", userID, CurrentSubID, dateAddedTime, dateCanceledTime)
	} else {
		result, _ = db.Exec("INSERT INTO UserSubs(UserID, SubID, DateAdded) VALUES (?,?,?);", userID, CurrentSubID, dateAddedTime)
	}

	//Tests to see if function worked
	numRows, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
		return -502
	}

	//fmt.Println("Rows Affected:", numRows)
	return int(numRows) + isRenewed
}

// Sets DateRemoved Value to current time based on userID and subscriptionName
func CancelUserSub(db *sql.DB, userID int, subscriptionName string) int {
	if subscriptionName == "" {
		return -204
	}

	//Gets the current time and formats it into YYYY-MM-DD hh:mm:ss
	currentTime := time.Now()
	currentTime.Format("2006-01-02 15:04:05")

	var CurrentSubID int

	//Gets the SubID from Subscriptions table
	sub_name, err := db.Query("SELECT SubID FROM Subscriptions WHERE Name = ?;", subscriptionName)

	if err != nil {
		log.Fatal(err)
		return -502
	}

	//Checks If Query Returns Empty Set or if the Subscription Name exists
	if sub_name.Next() {
		//Gets the SubID
		sub_name.Scan(&CurrentSubID)
		//fmt.Println("Sub ID:", CurrentSubID)

	} else {
		fmt.Println("Subscription Name is Invalid")
		return -404
	}

	//Update UserSub Data
	result, _ := db.Exec("UPDATE UserSubs SET DateRemoved = ? WHERE UserID = ? AND SubID = ? AND DateRemoved IS NULL;", currentTime, userID, CurrentSubID)

	//Tests to see if function worked
	numRows, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
		return -502
	}

	//fmt.Println("Rows Affected:", numRows)
	return int(numRows)
}

func CancelUserSubID(db *sql.DB, userSubID int) int {
	//Gets the current time and formats it into YYYY-MM-DD hh:mm:ss
	currentTime := time.Now()
	currentTime.Format("2006-01-02 15:04:05")

	//Update UserSub Data
	result, _ := db.Exec("UPDATE UserSubs SET DateRemoved = ? WHERE UserSubID = ?;", currentTime, userSubID)

	//Tests to see if function worked
	numRows, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
		return -502
	}

	//fmt.Println("Rows Affected:", numRows)
	return int(numRows)
}

// Deletes entry based on UserID from MySQL table called "Users"
func DeleteUser(db *sql.DB, ID int) {
	db.Exec("DELETE FROM Verification WHERE UserID = ?;", ID)
	db.Exec("DELETE FROM UserSubs WHERE UserID = ?;", ID)
	//originally result, err := db.Exec("DELETE FROM Users WHERE UserID = ?;", ID)
	_, err := db.Exec("DELETE FROM Users WHERE UserID = ?;", ID)

	if err != nil {
		log.Fatal(err)
	}

	//commented out to not show on unit test
	/*numRows, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Rows Affected:", numRows)*/
}

// Selects entry from database "Users" based on username and password
// Returns UserID or -401 when current user does not exist
// Returns -502 if there is an error with database connection
func Login(db *sql.DB, username string, password string) int {
	//Try To Login
	rows, err := db.Query("SELECT UserID FROM Users WHERE Username = ? AND Password = ?;", username, password)

	if err != nil {
		log.Fatal(err)
		return -502
	}

	//Tests If Query Returns Empty Set or Not (Valid Username and Password or Not)
	if rows.Next() {
		fmt.Println("Got Valid User ID!")

		var CurrentUserID int
		rows.Scan(&CurrentUserID)
		fmt.Println("Current User ID:", CurrentUserID)

		return CurrentUserID
		//Login Behavior

	} else {
		fmt.Println("Incorrect Username or Password!")
		return -401
	}
}

func GetMostUsedSubscription(db *sql.DB, currentID int, isContinuous bool, isCurrentlyActive bool) (string, int) {
	var rows *sql.Rows
	var err error

	if !isCurrentlyActive {
		if !isContinuous {
			rows, err = db.Query("SELECT Name, SUM(TotalUsageTimeSeconds) AS FinalTotalUsageTimeSeconds FROM (SELECT Name, TIMESTAMPDIFF(SECOND, DateAdded, DateRemoved) AS TotalUsageTimeSeconds FROM UserSubs INNER JOIN subscriptions ON UserSubs.SubID = Subscriptions.SubID WHERE UserID = ? AND DateRemoved IS NOT NULL UNION ALL SELECT Name, TIMESTAMPDIFF(SECOND, DateAdded, NOW()) AS TotalUsageTimeSeconds FROM UserSubs INNER JOIN subscriptions ON UserSubs.SubID = Subscriptions.SubID WHERE UserID = ? AND DateRemoved IS NULL) AS t1 GROUP BY NAME ORDER BY SUM(TotalUsageTimeSeconds) DESC LIMIT 1;", currentID, currentID)
		} else {
			rows, err = db.Query("SELECT * FROM (SELECT Name, TIMESTAMPDIFF(SECOND, DateAdded, DateRemoved) AS TotalUsageTimeSeconds FROM UserSubs INNER JOIN subscriptions ON UserSubs.SubID = Subscriptions.SubID WHERE UserID = ? AND DateRemoved IS NOT NULL UNION ALL SELECT Name, TIMESTAMPDIFF(SECOND, DateAdded, NOW()) AS TotalUsageTimeSeconds FROM UserSubs INNER JOIN subscriptions ON UserSubs.SubID = Subscriptions.SubID WHERE UserID = ? AND DateRemoved IS NULL) AS t1 GROUP BY TotalUsageTimeSeconds, NAME ORDER BY TotalUsageTimeSeconds DESC LIMIT 1;", currentID, currentID)
		}
	} else {
		rows, err = db.Query("SELECT Name, TIMESTAMPDIFF(SECOND, DateAdded, NOW()) AS TotalUsageTimeSeconds FROM UserSubs INNER JOIN subscriptions ON UserSubs.SubID = Subscriptions.SubID WHERE UserID = ? AND DateRemoved IS NULL ORDER BY TotalUsageTimeSeconds DESC LIMIT 1", currentID)
	}

	if err != nil {
		panic(err)
	}

	var subName string = "none"
	var subTimeUsedInSeconds int = 0

	for rows.Next() {
		rows.Scan(&subName, &subTimeUsedInSeconds)
	}

	return subName, subTimeUsedInSeconds
}

func GetPriceForMonth(db *sql.DB, currentID int, monthNumber int, yearNumber int) string {
	var stringPrice string = "0.00"
	compareDate := time.Date(yearNumber, time.Month(monthNumber+1), 1, 0, 0, 1, 0, time.UTC)

	var currentMonth string = strconv.Itoa(monthNumber)
	if len(currentMonth) < 2 {
		currentMonth = "0" + currentMonth
	}

	SQLStringYearMonth := strconv.Itoa(yearNumber) + "-" + currentMonth + "-%d 00:00:00"

	rows, err := db.Query("SELECT ROUND(SUM(TotalPrice),2) AS FinalTotalPrice FROM(SELECT ROUND(SUM(Price),2) AS TotalPrice FROM UserSubs INNER JOIN Subscriptions ON UserSubs.SubID = Subscriptions.SubID INNER JOIN Users ON UserSubs.UserID = Users.UserID WHERE UserSubs.UserID = ? AND Name NOT LIKE \"%3 Month%\" AND Name NOT LIKE \"%Yearly%\" AND EXTRACT(Year FROM DateAdded) <= ? AND DateAdded < ? AND (DATE_FORMAT(DateAdded, ?) < DateRemoved OR DateRemoved IS NULL) UNION ALL SELECT ROUND(SUM(Price),2) AS TotalPrice FROM UserSubs INNER JOIN Subscriptions ON UserSubs.SubID = Subscriptions.SubID INNER JOIN Users ON UserSubs.UserID = Users.UserID WHERE UserSubs.UserID = ? AND Name LIKE \"%3 Month%\" AND MOD(Extract(Month FROM DateAdded) - ?, 3) = 0 AND EXTRACT(Year FROM DateAdded) <= ?  AND DateAdded < ? AND (DATE_FORMAT(DateAdded, ?) < DateRemoved OR DateRemoved IS NULL) UNION ALL SELECT ROUND(SUM(Price),2) AS TotalPrice FROM UserSubs INNER JOIN Subscriptions ON UserSubs.SubID = Subscriptions.SubID INNER JOIN Users ON UserSubs.UserID = Users.UserID WHERE UserSubs.UserID = ? AND Name LIKE \"%Yearly%\" AND Extract(Month FROM DateAdded) = ? AND EXTRACT(Year FROM DateAdded) <= ? AND DateAdded < ? AND (DATE_FORMAT(DateAdded, ?) < DateRemoved OR DateRemoved IS NULL)) t1;", currentID, yearNumber, compareDate, SQLStringYearMonth, currentID, monthNumber, yearNumber, compareDate, SQLStringYearMonth, currentID, monthNumber, yearNumber, compareDate, SQLStringYearMonth)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		rows.Scan(&stringPrice)
	}

	return stringPrice
}

// UserID and dateAdded will normally be taken automatically from the database
// instead of specifying it directly, unlike in this test
func ManuallyTestBackend(db *sql.DB) {
	fmt.Println("Type -1 to quit the test.")
	var choice int
	for choice != -1 {
		fmt.Print("Enter a number from 1 - 10: ")
		fmt.Scanln(&choice)
		if choice == 1 {
			ResetAllTables(db)
			fmt.Println("Choice 1: Clearing tables from database \"userdb.\"")
		} else if choice == 2 {
			SetUpTables(db)
			fmt.Println("Choice 2: Setting up all tables for database \"userdb.\"")
		} else if choice == 3 {
			fmt.Println("Choice 3: Getting table sizes.")
			fmt.Println("Subscriptions table size: " + strconv.Itoa(GetTableSize(db, "subscriptions")))
			fmt.Println("Users table size: " + strconv.Itoa(GetTableSize(db, "users")))
			fmt.Println("UserSubs table size: " + strconv.Itoa(GetTableSize(db, "usersubs")))
		} else if choice == 4 {
			fmt.Println("Choice 4: Creates new user.")
			fmt.Println("Enter a username, password, and email: ")
			var a, b, c string
			fmt.Scanln(&a, &b, &c)
			CreateNewUser(db, a, b, c)
		} else if choice == 5 {
			fmt.Println("Choice 5: Creates new subscription.")
			fmt.Println("Enter a name and price: ")
			var a, b string
			fmt.Scanln(&a, &b)
			CreateNewSub(db, a, b)
		} else if choice == 6 {
			fmt.Println("Choice 6: Subscribes to subscription service.")
			fmt.Println("Enter a UserID and Subscription Name: ")
			var a int
			var b string
			fmt.Scanln(&a, &b)
			CreateNewUserSub(db, a, b)
		} else if choice == 7 {
			fmt.Println("Choice 7: Cancels subscription service.")
			fmt.Println("Enter a UserID and Subscription Name: ")
			var a int
			var b string
			fmt.Scanln(&a, &b)
			CancelUserSub(db, a, b)
		} else if choice == 8 {
			fmt.Println("Choice 8: Adds pre-existing subscription service.")
			fmt.Println("Enter a UserID, Subscription Name, Date, Time Added, and Time Canceled (Possibly Optional): ")
			var a int
			var b, c, d, e, f string
			fmt.Scanln(&a, &b, &c, &d, &e, &f)
			dateAndTime := c + " " + d
			dateAndTimeCanceled := e + " " + f
			AddOldUserSub(db, a, b, dateAndTime, dateAndTimeCanceled)
		} else if choice == 9 {
			fmt.Println("Choice 9: Deletes user that is specified.")
			fmt.Println("Enter a UserID: ")
			var a int
			fmt.Scanln(&a)
			DeleteUser(db, a)
		} else if choice == 10 {
			fmt.Println("Choice 10: Changes the password of specified user.")
			fmt.Println("Enter a UserID, the Old Password, and the New Password: ")
			var a int
			var b, c string
			fmt.Scanln(&a, &b, &c)
			ChangePassword(db, a, b, c)
		}
	}
}
