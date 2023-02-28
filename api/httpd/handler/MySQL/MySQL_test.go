package MySQL

import (
	"testing"
)

//TODO: Check if server connection errors work as intended (-502)
//Clear tables and add tables as needed to check for duplicates

// Tests if server error is caught properly
// Change the MySQLPassword.txt name to test this
func TestMySQLConnect(t *testing.T) {
	defer func() {
		//checks for panic()
		//example would be when password txt file is not found
		if r := recover(); r != nil {
			t.Errorf("An error occurred: %v", r)
		}
	}()

	// Function under test:
	MySQLConnect()
}

func TestGetDatabaseSize(t *testing.T) {
	db := MySQLConnect()
	ResetAllTables(db)

	db.Exec("CREATE TABLE IF NOT EXISTS Users (UserID int NOT NULL AUTO_INCREMENT, Email varchar(255) NOT NULL, Username varchar(255) NOT NULL, Password varchar(255) NOT NULL, UNIQUE(Username), UNIQUE(Email), PRIMARY KEY(UserID));")
	expected := 1
	actual := GetDatabaseSize(db)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSetUpTables(t *testing.T) {
	db := MySQLConnect()
	ResetAllTables(db)
	SetUpTables(db)

	expected := 4
	actual := GetDatabaseSize(db)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestGetTableSize(t *testing.T) {
	db := MySQLConnect()
	ResetAllTables(db)
	SetUpTables(db)

	// Tests if error calls work as intended
	tableName := "userss"
	// Call the function expecting to return -404 if there is an error
	errorCode := GetTableSize(db, tableName)

	// Checks if the error code is working as intended (shows up when the table does not exist)
	if errorCode != -404 {
		t.Errorf("Expected an error code -404, but got %d", errorCode)
	}

	CreateAdminUser(db)

	expected := 1
	actual := GetTableSize(db, "users")
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestResetTable(t *testing.T) {
	db := MySQLConnect()
	ResetAllTables(db)
	SetUpTables(db)

	ResetTable(db, "Verification")

	expected := 3
	actual := GetDatabaseSize(db)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestResetAllTables(t *testing.T) {
	db := MySQLConnect()
	ResetAllTables(db)
	SetUpTables(db)

	ResetAllTables(db)

	expected := 0
	actual := GetDatabaseSize(db)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestCreateNewUser(t *testing.T) {
	db := MySQLConnect()
	ResetAllTables(db)
	SetUpTables(db)
	CreateAdminUser(db)

	username := ""
	password := ""
	email := ""

	errorCode := CreateNewUser(db, username, password, email)
	if errorCode != -204 {
		t.Errorf("Expected an error code -204, but got %d", errorCode)
	}

	//Will always have admin credentials, so tests using them
	//testing duplicate username
	username = "root"
	password = "test"
	email = "unique@gmail.com"

	errorCode = CreateNewUser(db, username, password, email)
	if errorCode != -223 {
		t.Errorf("Expected an error code -223, but got %d", errorCode)
	}

	//testing duplicate email
	username = "unique"
	email = "vanbestindustries@gmail.com"

	errorCode = CreateNewUser(db, username, password, email)
	if errorCode != -225 {
		t.Errorf("Expected an error code -225, but got %d", errorCode)
	}

	ResetAllTables(db)
	SetUpTables(db)
	CreateAdminUser(db)
	CreateNewUser(db, "testUser", "password", "example@gmail.com")
	expected := 2
	actual := GetTableSize(db, "users")
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestCreateAdminUser(t *testing.T) {
	db := MySQLConnect()
	ResetAllTables(db)
	SetUpTables(db)
	CreateAdminUser(db)

	expected := 1
	actual := GetTableSize(db, "users")
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestCreateCommonSubscriptions(t *testing.T) {
	db := MySQLConnect()
	ResetAllTables(db)
	SetUpTables(db)
	CreateCommonSubscriptions(db)

	expected := 46
	actual := GetTableSize(db, "subscriptions")
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestGetPassword(t *testing.T) {
	db := MySQLConnect()
	ResetAllTables(db)
	SetUpTables(db)
	CreateAdminUser(db)

	expected := "XohImNooBHFR0OVvjcYpJ3NgPQ1qq73WKhHvch0VQtg="
	actual := GetPassword(db, 1)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestChangePassword(t *testing.T) {
	db := MySQLConnect()
	ResetAllTables(db)
	SetUpTables(db)
	CreateAdminUser(db)

	userID := 1
	oldPassword := ""
	newPassword := ""

	errorCode := ChangePassword(db, userID, oldPassword, newPassword)
	if errorCode != -204 {
		t.Errorf("Expected an error code -204, but got %d", errorCode)
	}

	//Checks if password changed
	//var oldPassword string

	//expected := db.
}

func TestChangeEmail(t *testing.T) {
	db := MySQLConnect()

	//ResetAllTables(db)
	//SetUpTables(db)
	//CreateAdminUser(db)
	//CreateNewUser(db, "test", "doesn't matter", "valekseev2003@gmail.com")

	userID := 1
	//Takes from a user that is not the admin's email, so to pass this test
	//this, the email must already be in the database
	newEmail := "valekseev2003@gmail.com"

	errorCode := ChangeEmail(db, userID, newEmail)
	if errorCode != -223 {
		t.Errorf("Expected an error code -223, but got %d", errorCode)
	}
}

func TestChangeUsername(t *testing.T) {
	db := MySQLConnect()
	userID := 1
	//Takes from a user that is not the admin's username, so to pass this test
	//this, the username must already be in the database
	newUsername := "valek"

	errorCode := ChangeUsername(db, userID, newUsername)
	if errorCode != -223 {
		t.Errorf("Expected an error code -223, but got %d", errorCode)
	}
}

func TestCreateNewSub(t *testing.T) {
	db := MySQLConnect()
	name := ""
	price := ""

	errorCode := CreateNewSub(db, name, price)
	if errorCode != -204 {
		t.Errorf("Expected an error code -204, but got %d", errorCode)
	}

	//The name of "test" must already be in the database
	name = "test"
	price = "0"
	errorCode = CreateNewSub(db, name, price)
	if errorCode != -223 {
		t.Errorf("Expected an error code -223, but got %d", errorCode)
	}
}

func TestDeleteUser(t *testing.T) {
	db := MySQLConnect()
	//ResetAllTables(db)
	//SetUpTables(db)
	//CreateAdminUser(db)

	defer func() {
		//checks for panic()
		//example would be when password txt file is not found
		if r := recover(); r != nil {
			t.Errorf("An error occurred: %v", r)
		}
	}()

	// Function under test:
	originalTableSize := GetTableSize(db, "Users")
	DeleteUser(db, 3)
	finalTableSize := GetTableSize(db, "Users")

	if originalTableSize != finalTableSize {
		t.Errorf("Expected to not delete user, but user was deleted")
	}

	originalTableSize = GetTableSize(db, "Users")
	DeleteUser(db, 1)
	finalTableSize = GetTableSize(db, "Users")

	if originalTableSize == finalTableSize {
		t.Errorf("Expected to delete user, but user was not deleted")
	}
}
