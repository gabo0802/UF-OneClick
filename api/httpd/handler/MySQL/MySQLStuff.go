package MySQL

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

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
	//defer db.Close()

	err = db.Ping()
	checkError(err)
	fmt.Println("Successfully created connection to database.")
	// Drop previous table of same name if one exists.
	_, err = db.Exec("DROP TABLE IF EXISTS inventory;")
	checkError(err)
	fmt.Println("Finished dropping table (if existed).")

	// Create table.
	_, err = db.Exec("CREATE TABLE inventory (id serial PRIMARY KEY, name VARCHAR(50), quantity INTEGER);")
	checkError(err)
	fmt.Println("Finished creating table.")

	// Insert some data into table.
	/*
		sqlStatement, err := db.Prepare("INSERT INTO inventory (name, quantity) VALUES (?, ?);")
		res, err := sqlStatement.Exec("banana", 150)
		checkError(err)
		rowCount, err := res.RowsAffected()
		fmt.Printf("Inserted %d row(s) of data.\n", rowCount)

		res, err = sqlStatement.Exec("orange", 154)
		checkError(err)
		rowCount, err = res.RowsAffected()
		fmt.Printf("Inserted %d row(s) of data.\n", rowCount)

		res, err = sqlStatement.Exec("apple", 100)
		checkError(err)
		rowCount, err = res.RowsAffected()
		fmt.Printf("Inserted %d row(s) of data.\n", rowCount)
		fmt.Println("Done.")
	*/

	return db
}

// func MySQLConnect() *sql.DB {
// 	var db *sql.DB

// 	//Try Connection
// 	//db, err := sql.Open("mysql", "remoteuser:Rem0teUser!@tcp(DESKTOP-KOPDURN:3306)/user")
// 	db, err := sql.Open("mysql", "root:MySQLP@ssw0rd@tcp(127.0.0.1:3306)/sys")

// 		//Test Connection
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		pingErr := db.Ping()
// 		if pingErr != nil {
// 			log.Fatal(pingErr)
// 		}

// 		//Confirmation
// 		fmt.Println("Connected")

// 	db.Exec("CREATE DATABASE IF NOT EXISTS user;")
// 	db.Exec("USE user;")

// 	/*pingErr = db.Ping()
// 	if pingErr != nil {
// 		log.Fatal(pingErr)
// 	}*/

// 	//fmt.Println("Connected")

// 	return db
// 	//Function Code Based From: https://go.dev/doc/tutorial/database-access
// }

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

	//Subscriptions
	db.Exec("CREATE TABLE IF NOT EXISTS Subscriptions(UserID int NOT NULL, FOREIGN KEY(UserID) REFERENCES Users(UserID));") //update for necessary parameters
}

func ResetTables(db *sql.DB) {
	db.Exec("DROP TABLE IF EXISTS Users;")
	db.Exec("DROP TABLE IF EXISTS Subscriptions;")
}

func CreateNewUser(db *sql.DB, username string, password string) {
	//Create New User
	result, err := db.Exec("INSERT INTO Users(Username, Password) VALUES (?,?);", username, password)

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			fmt.Println("Username Already Exists!")
			return
		} else {
			log.Fatal(err)

		}
	}

	numRows, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Rows Affected:", numRows)
	//Test If User Creation Worked
}

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
