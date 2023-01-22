package main

/*
//Start of Back-End
//Might Work (haven't tested yet since can't get driver to work yet)

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql" //error getting this to work
)

var db *sql.DB


function getDatabaseSize(string database)(int64){
	rows, err := db.Query("SELECT * FROM ?", database)
	size := 0

	for rows.Next(){
		size := size + 1
	}

	return size
}

func SetUpDatabase(){
	//Users
	db.Exec("CREATE TABLE IF NOT EXISTS Users (UserID int NOT NULL AUTO_INCREMENT, Username varchar(255) NOT NULL, Password varchar(255) NOT NULL, UNIQUE(Username), PRIMARY KEY(UserID));")

	//Subscriptions
	db.Exec("CREATE TABLE IF NOT EXISTS Subscriptions(UserID int NOT NULL, FOREIGN KEY(UserID) REFERENCES Users(UserID));")
}

func CreateNewUser(username string, password string){
	//Create New User
	db.Exec("INSERT INTO Users(Username, Password) VALUES (?,?)", username, password);

	//Test If User Creation Worked
}

func Login(username string, password string){
	//Try To Login
	db.Query("SELECT ID FROM Users WHERE Username = ? AND Password = ?", username, password);

	//Login Behavior
}

func Connect() {
	//Config
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	//Try Connection
	db, err := sql.Open("mysql", cfg.FormatDSN())

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

	//Function Code From: https://go.dev/doc/tutorial/database-access
}

func main() {
	Connect()
	SetUpDatabase()
}
*/


/*
//SQL Commands

CREATE TABLE IF NOT EXISTS Users (
    UserID int NOT NULL AUTO_INCREMENT,
	Username varchar(255) NOT NULL,
	Password varchar(255) NOT NULL,

	UNIQUE(Username),
	PRIMARY KEY(UserID),
);

CREATE TABLE IF NOT EXISTS Subscriptions(
	UserID int NOT NULL,
	FOREIGN KEY(UserID) REFERENCES Users(UserID)
)

SELECT ID FROM Users WHERE Username = 'root' AND Password = 'password' //login, if empty username or password wrong, use ID to see subscriptions
//SELECT MAX(UserID) AS LastID FROM Users //probably unnecessary since UserID has AUTO_INCREMENT

*/
