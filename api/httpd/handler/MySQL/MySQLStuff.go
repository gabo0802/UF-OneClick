package MySQL

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func MySQLConnect() *sql.DB {
	var db *sql.DB

	//Try Connection
	db, err := sql.Open("mysql", "root:MySQLP@ssw0rd@tcp(127.0.0.1:3306)/sys")
	//db, err := sql.Open("mysql", "root:MySQLP@ssw0rd@tcp(127.0.0.1:3306)/user") //unnecessary

	//Test Connection
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	//Confirmation
	fmt.Println("Connected")

	db.Exec("CREATE DATABASE IF NOT EXISTS user;")
	db.Exec("USE user;")

	pingErr = db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected")

	return db
	//Function Code Based From: https://go.dev/doc/tutorial/database-access
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

	//Subscriptions
	//db.Exec("CREATE TABLE IF NOT EXISTS Subscriptions(UserID int NOT NULL, FOREIGN KEY(UserID) REFERENCES Users(UserID));") //update for necessary parameters
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
		fmt.Println(CurrentUserID)

		return CurrentUserID
		//Login Behavior

	} else {
		fmt.Println("Incorrect Username or Password!")
		return -1
	}
}
