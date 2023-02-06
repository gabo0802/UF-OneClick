package MySQL

import (
	"database/sql"
	"fmt"
	"log"
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
		fmt.Println("Error: Database Does Not Exist!")
		return -1
	}

	for rows.Next() {
		size += 1
	}

	return size
}

func SetUpTables(db *sql.DB) {
	//Users
	db.Exec("CREATE TABLE IF NOT EXISTS Users (UserID int NOT NULL AUTO_INCREMENT, Username varchar(255) NOT NULL, Password varchar(255) NOT NULL, UNIQUE(Username), PRIMARY KEY(UserID));")

	//All available subscriptions
	db.Exec("CREATE TABLE IF NOT EXISTS Subscriptions (SubID int NOT NULL AUTO_INCREMENT, Name varchar(255) NOT NULL, Price varchar(255) NOT NULL, UNIQUE(Name), PRIMARY KEY(SubID));")

	//Individual user subscriptions
	db.Exec("CREATE TABLE IF NOT EXISTS UserSubs (UserID int NOT NULL, SubID int NOT NULL, DateAdded varchar(255) NOT NULL, DateRemoved varchar(255), FOREIGN KEY(UserID) REFERENCES Users(UserID), FOREIGN KEY(SubID) REFERENCES Subscriptions(SubID))")
}

func ResetTable(db *sql.DB, tableName string) {
	db.Exec("DROP TABLE IF EXISTS " + tableName)
}

func ResetAllTables(db *sql.DB) {
	db.Exec("DROP TABLE IF EXISTS Users;")         //won't work due to foreign key constraints
	db.Exec("DROP TABLE IF EXISTS Subscriptions;") //won't work due to foreign key constraints
	db.Exec("DROP TABLE IF EXISTS UserSubs;")

	db.Exec("DELETE Users WHERE UserID > 1;")
	db.Exec("DELETE Subscriptions WHERE SubID > 1;")
}

func CreateNewUser(db *sql.DB, username string, password string) int {
	//Create New User
	result, err := db.Exec("INSERT INTO Users(Username, Password) VALUES (?,?);", username, password)

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			fmt.Println("Username Already Exists!")
			return 0
		} else {
			log.Fatal(err)
			return -1
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
	db.Exec("INSERT INTO Users(UserID, Username, Password) VALUES (1,root,password);")
}

func ChangePassword(db *sql.DB, userID int, oldPassword string, newPassword string) int {
	result, err := db.Exec("UPDATE Users SET Password = ? WHERE userID = ? AND Password = ?;", newPassword, userID, oldPassword)
	if err != nil {
		return -1
	}

	numRows, err := result.RowsAffected()

	if err != nil {
		return -1
	}

	return int(numRows)
}

func CreateNewSub(db *sql.DB, name string, price string) int {
	//Create New Subscription
	result, err := db.Exec("INSERT INTO Subscriptions(name, price) VALUES (?,?);", name, price)

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			fmt.Println("Subscription Name Already Exists!")
			return 0
		} else {
			log.Fatal(err)
			return -1
		}
	}

	//Tests to see if function worked (can remove later)
	numRows, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
		return -1
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
			return -1 //if not then subscription still exists
		} else {
			return 1 //can add new subscription
		}
	}

	return 0 //doesn't exist yet
}

func CreateNewUserSub(db *sql.DB, userID int, subscriptionName string) int {
	//Gets the current time and formats it into YYYY-MM-DD hh:mm:ss
	currentTime := time.Now()
	currentTime.Format("2006-01-02 15:04:05")

	var CurrentSubID int

	//Gets the SubID from Subscriptions table
	sub_name, err := db.Query("SELECT SubID FROM Subscriptions WHERE Name = ?;", subscriptionName)

	if err != nil {
		log.Fatal(err)
		return -2
	}

	//Checks If Query Returns Empty Set or if the Subscription Name exists
	if sub_name.Next() {
		//Gets the SubID
		sub_name.Scan(&CurrentSubID)
		//fmt.Println("Sub ID:", CurrentSubID)

	} else {
		fmt.Println("Subscription Name is Invalid")
		return -1
	}

	//Checks to see if sub was already added to user before creating a new table value
	var isRenewed int = CanAddUserSub(db, userID, CurrentSubID)
	if isRenewed < 0 {
		fmt.Println("Subscription already added to User's Profile!")
		return 0
	}

	//Create New UserSub Data
	result, _ := db.Exec("INSERT INTO UserSubs(UserID, SubID, DateAdded) VALUES (?,?,?);", userID, CurrentSubID, currentTime.Format("2006-01-02 15:04:05"))

	//Tests to see if function worked (can remove later)
	numRows, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
		return -1
	}

	fmt.Println("Rows Affected:", numRows)
	return int(numRows) + isRenewed
}

// Sets DateRemoved Value to current time based on userID and subscriptionName
func CancelUserSub(db *sql.DB, userID int, subscriptionName string) int {
	//Gets the current time and formats it into YYYY-MM-DD hh:mm:ss
	currentTime := time.Now()
	currentTime.Format("2006-01-02 15:04:05")

	var CurrentSubID int

	//Gets the SubID from Subscriptions table
	sub_name, err := db.Query("SELECT SubID FROM Subscriptions WHERE Name = ?;", subscriptionName)

	if err != nil {
		log.Fatal(err)
	}

	//Checks If Query Returns Empty Set or if the Subscription Name exists
	if sub_name.Next() {
		//Gets the SubID
		sub_name.Scan(&CurrentSubID)
		//fmt.Println("Sub ID:", CurrentSubID)

	} else {
		fmt.Println("Subscription Name is Invalid")
		return 0
	}

	//Update UserSub Data
	result, _ := db.Exec("UPDATE UserSubs SET DateRemoved = ? WHERE UserID = ? AND SubID = ? AND DateRemoved IS NULL;", currentTime.Format("2006-01-02 15:04:05"), userID, CurrentSubID)

	//Tests to see if function worked (can remove later)
	numRows, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
		return -1
	}

	fmt.Println("Rows Affected:", numRows)
	return int(numRows)
}

//TODO: Add a way to remove a subscription linked to a user and add it to the DateRemoved column <- No We Keep This For Future Data Use

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
		return -2
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
		return -1
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
