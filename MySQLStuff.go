package main

/*
//Start of Back-End

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql" //error getting this to work
)

var db *sql.DB


func SetUpDatabase(){
	//see CREATE TABLES
}

func CreateNewUser(){
	//Insert INTO Users _______________
}

func Login(){
	//see Select ID FROM Users WHERE _________
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
}

func main() {
	Connect()
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
