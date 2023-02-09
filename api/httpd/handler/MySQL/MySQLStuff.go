package MySQL

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "oneclickserver.mysql.database.azure.com"
	database = "userdb"
	user     = "adminUser"
	password = "MySQLP@ssw0rd"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func MySQLConnect() *sql.DB {
	var db *sql.DB

	//Connect to remote server using Microsoft Azure
	// Initialize connection string
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)

	// Initialize connection object
	db, err := sql.Open("mysql", connectionString)
	checkError(err)

	err = db.Ping()
	checkError(err)
	fmt.Println("Successfully created connection to database.")
	//Function code based on https://learn.microsoft.com/en-us/azure/mysql/single-server/connect-go

	return db
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

func SetUpTables(db *sql.DB) {
	//Users
	db.Exec("CREATE TABLE IF NOT EXISTS Users (UserID int NOT NULL AUTO_INCREMENT, Email varchar(255) NOT NULL, Username varchar(255) NOT NULL, Password varchar(255) NOT NULL, UNIQUE(Username), UNIQUE(Email), PRIMARY KEY(UserID));")

	//All available subscriptions
	db.Exec("CREATE TABLE IF NOT EXISTS Subscriptions (SubID int NOT NULL AUTO_INCREMENT, Name varchar(255) NOT NULL, Price varchar(255) NOT NULL, UNIQUE(Name), PRIMARY KEY(SubID));")

	//Individual user subscriptions
	db.Exec("CREATE TABLE IF NOT EXISTS UserSubs (UserID int NOT NULL, SubID int NOT NULL, DateAdded varchar(255) NOT NULL, DateRemoved varchar(255), FOREIGN KEY(UserID) REFERENCES Users(UserID), FOREIGN KEY(SubID) REFERENCES Subscriptions(SubID))")
}

func ResetTable(db *sql.DB, tableName string) {
	db.Exec("DROP TABLE IF EXISTS " + tableName)
}

func ResetAllTables(db *sql.DB) {
	//Must drop UserSubs first since its foreign keys depend on the other tables for primary keys
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
				return (-223 + 2) //already exists (third variable)
			} else {
				fmt.Println("Username Already Exists!")
				fmt.Println(err.Error())
				return (-223 + 0) //already exists (first variable)
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

	fmt.Println("Rows Affected:", numRows)
	//Test If User Creation Worked (can remove later)

	return int(numRows)
}

func CreateAdminUser(db *sql.DB) {
	result, err := db.Exec("INSERT INTO Users(UserID, Username, Password, Email) VALUES (1, \"SBNJTRN-FjG7owHVrKtue7eqdM4RhdRWVl71HXN2d7I=\", \"XohImNooBHFR0OVvjcYpJ3NgPQ1qq73WKhHvch0VQtg=\", \"companyemail@gmail.com\");")
	//maybe change password to something more secure?

	if err != nil {
		log.Fatal(err)
	}

	numRows, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Rows Affected:", numRows)
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

func CreateNewSub(db *sql.DB, name string, price string) int {
	//Create New Subscription
	if name == "" || price == "" {
		return -204
	}

	result, err := db.Exec("INSERT INTO Subscriptions(name, price) VALUES (?,?);", name, price)

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			fmt.Println("Subscription Name Already Exists!")
			return -223
		} else {
			log.Fatal(err)
			return -502
		}
	}

	//Tests to see if function worked (can remove later)
	numRows, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
		return -502
	}

	fmt.Println("Rows Affected:", numRows)
	return int(numRows)
}

func CanAddUserSub(db *sql.DB, userID int, subID int) int {
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
			return 1 //can add new subscription
		}
	}

	return 0 //doesn't exist yet
}

// Adds based on the current time
func CreateNewUserSub(db *sql.DB, userID int, subscriptionName string) int {
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

	//Checks to see if sub was already added to user before creating a new table value
	var isRenewed int = CanAddUserSub(db, userID, CurrentSubID)
	if isRenewed < 0 {
		fmt.Println("Subscription already added to User's Profile!")
		return -223
	}

	//Create New UserSub Data
	result, _ := db.Exec("INSERT INTO UserSubs(UserID, SubID, DateAdded) VALUES (?,?,?);", userID, CurrentSubID, currentTime.Format("2006-01-02 15:04:05"))

	//Tests to see if function worked (can remove later)
	numRows, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
		return -402
	}

	fmt.Println("Rows Affected:", numRows)
	return int(numRows) + isRenewed
}

func CanAddOldUserSub(db *sql.DB, userID int, subID int, oldDate string, oldCanceledDate string) int {
	rows, err := db.Query("SELECT DateAdded, DateRemoved FROM UserSubs WHERE UserID = ? AND SubID = ? AND DateAdded > ?;", userID, subID, oldDate) //sub in future already exists (need to have old sub be canceled)
	if err != nil {
		log.Fatal(err)
	}
	if rows.Next() && oldCanceledDate == "" {
		fmt.Println("Error: Can't Have Old Subscription and New Subscription At The Same Time")
		return -401
	}

	rows, err = db.Query("SELECT DateAdded, DateRemoved FROM UserSubs WHERE UserID = ? AND SubID = ? AND DateAdded <= ? AND DateRemoved >= ?;", userID, subID, oldDate, oldDate)
	if err != nil {
		log.Fatal(err)
	}
	if rows.Next() {
		fmt.Println("Error: Can Have The Same Subscription In The Middle Of A Subscription")
		return -401
	}

	rows, err = db.Query("SELECT DateAdded, DateRemoved FROM UserSubs WHERE UserID = ? AND SubID = ? AND DateAdded = ? ;", userID, subID, oldDate)
	if err != nil {
		log.Fatal(err)
	}
	if rows.Next() {
		fmt.Println("Error: Can't Have The Same Subscription Be On the Same Exact Time")
		return -401
	}

	if oldCanceledDate != "" {
		rows, err = db.Query("SELECT DateAdded FROM UserSubs WHERE UserID = ? AND SubID = ? AND DateAdded < ? AND DateAdded > ?;", userID, subID, oldCanceledDate, oldDate)

		if err != nil {
			log.Fatal(err)
		}

		if rows.Next() {
			fmt.Println("Error: Can Have The Same Subscription In The Middle Of A Subscription")
			return -401
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
			return -401 //if not then subscription still exists
		} else {
			return 1 //can add new subscription
		}
	}

	return 0 //doesn't exist yet
}

// Allows pre-existing subscriptions with dates other than the current time to be added
func AddOldUserSub(db *sql.DB, userID int, subscriptionName string, dateAdded string, dateCanceled string) int {
	_, err := time.Parse("2006-01-02 15:04:01", dateAdded)
	if err != nil {
		fmt.Println("Error: Date Added Not Formatted Properly")
		log.Fatal(err)
		return -415
	}

	if dateCanceled != "" {
		_, err = time.Parse("2006-01-02 15:04:01", dateCanceled)
		if err != nil {
			log.Fatal(err)
			fmt.Println("Error: Date Canceled Not Formatted Properly")
			return -415
		}
	}

	if dateCanceled != "" && dateCanceled < dateAdded {
		fmt.Println("Error: Can't Cancel Subscription Before Adding It")
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
	var isRenewed int = CanAddOldUserSub(db, userID, CurrentSubID, dateAdded, dateCanceled)
	if isRenewed < 0 {
		fmt.Println("Subscription already added to User's Profile!")
		fmt.Println(dateAdded + ", " + dateCanceled)
		return -223
	}

	//Create New UserSub Data
	result, _ := db.Exec("")

	if dateCanceled != "" {
		result, _ = db.Exec("INSERT INTO UserSubs(UserID, SubID, DateAdded, DateRemoved) VALUES (?,?,?,?);", userID, CurrentSubID, dateAdded, dateCanceled)
	} else {
		result, _ = db.Exec("INSERT INTO UserSubs(UserID, SubID, DateAdded) VALUES (?,?,?);", userID, CurrentSubID, dateAdded)
	}

	//Tests to see if function worked (can remove later)
	numRows, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
		return -502
	}

	fmt.Println("Rows Affected:", numRows)
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
	result, _ := db.Exec("UPDATE UserSubs SET DateRemoved = ? WHERE UserID = ? AND SubID = ? AND DateRemoved IS NULL;", currentTime.Format("2006-01-02 15:04:05"), userID, CurrentSubID)

	//Tests to see if function worked (can remove later)
	numRows, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
		return -502
	}

	fmt.Println("Rows Affected:", numRows)
	return int(numRows)
}

// Deletes entry based on username and password from MySQL table called "Users"
/*func DeleteUser(db *sql.DB, username string, password string) {
    result, err := db.Exec("DELETE FROM Users WHERE Username = ? AND Password = ?;", username, password)

    if err != nil {
        log.Fatal(err)
    }

    numRows, err := result.RowsAffected()

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Rows Affected:", numRows)
}*/

// Deletes entry based on UserID from MySQL table called "Users"
func DeleteUser(db *sql.DB, ID int) {
	db.Exec("DELETE FROM UserSubs WHERE UserID = ?;", ID)
	result, err := db.Exec("DELETE FROM Users WHERE UserID = ?;", ID)

	if err != nil {
		log.Fatal(err)
	}

	numRows, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Rows Affected:", numRows)
}

// Selects entry from database "Users" based on username and password
// Returns UserID or -1 when current user does not exist
// Returns -2 if there is an error with database connection
func Login(db *sql.DB, username string, password string) int {
	//Try To Login
	rows, err := db.Query("SELECT UserID FROM Users WHERE Username = ? AND Password = ?;", username, password)

	if err != nil {
		log.Fatal(err)
		return -502
	}

	//Tests If Query Returns Empty Set or Not (Valid Username and Password or Not)
	if rows.Next() {
		fmt.Println("Login Successful!")

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

// UserID and dateAdded will normally be taken automatically from the database
// instead of specifying it directly, unlike in this test
func TestBackend(db *sql.DB) {
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

// Can use for unit testing later on
// Outputs database data onto the terminal
func ShowDatabaseTables(db *sql.DB, databaseName string) {
	db.Exec("USE " + databaseName + ";")
	res, _ := db.Query("SHOW TABLES;")

	var table string

	for res.Next() {
		res.Scan(&table)
		fmt.Println(table)
	}
}

func GetColumnData(db *sql.DB, databaseName string, tableName string, columnName string) {
	db.Exec("USE " + databaseName + ";")
	sqlCode := "SELECT " + columnName + " FROM " + tableName + ";"

	rows, _ := db.Query(sqlCode)

	var singleRow string

	for rows.Next() {
		rows.Scan(&singleRow)
		fmt.Println(singleRow)
	}
}
