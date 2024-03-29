package MySQL

import (
	"testing"
	"time"
)

//Tests all non-helper functions in MySQL package for basic functionality
//Except for "func ManuallyTestBackend()"

// Tests if server error is caught properly
// Change the MySQLPassword.txt name to test this
func TestMySQLConnect(t *testing.T) {
	defer func() {
		//checks for panic()
		//example would be when password txt file is not found
		if r := recover(); r != nil {
			t.Errorf("MySQLConnect() FAILED. An error occurred: %v", r)
		} else {
			t.Logf("MySQLConnect() PASSED. Connection is established")
		}
	}()

	// Function under test:
	MySQLConnect()
}

func TestGetDatabaseSize(t *testing.T) {
	db := MySQLConnect()
	defer db.Close()
	ResetAllTables(db)

	db.Exec("CREATE TABLE IF NOT EXISTS Users (UserID int NOT NULL AUTO_INCREMENT, Email varchar(255) NOT NULL, Username varchar(255) NOT NULL, Password varchar(255) NOT NULL, UNIQUE(Username), UNIQUE(Email), PRIMARY KEY(UserID));")
	expected := 1
	actual := GetDatabaseSize(db)
	if actual != expected {
		t.Errorf("GetDatabaseSize(db) FAILED. Expected %d table, got %d", expected, actual)
	} else {
		t.Logf("GetDatabaseSize(db) PASSED. Expected %d table, got %d", expected, actual)
	}
}

func TestSetUpTables(t *testing.T) {
	db := MySQLConnect()
	defer db.Close()
	ResetAllTables(db)
	SetUpTables(db)

	expected := 4
	actual := GetDatabaseSize(db)
	if actual != expected {
		t.Errorf("SetUpTables(db) FAILED. Expected %d tables, got %d", expected, actual)
	} else {
		t.Logf("SetUpTables(db) PASSED. Expected %d tables, got %d", expected, actual)
	}
}

func TestGetTableSize(t *testing.T) {
	db := MySQLConnect()
	defer db.Close()
	ResetAllTables(db)
	SetUpTables(db)

	// Tests if error calls work as intended
	tableName := "userss"
	// Call the function expecting to return -404 if there is an error
	errorCode := GetTableSize(db, tableName)

	// Checks if the error code is working as intended (shows up when the table does not exist)
	if errorCode != -404 {
		t.Errorf("GetTableSize(db, \"userss\") FAILED. Expected an error code -404, got %d", errorCode)
	} else {
		t.Logf("GetTableSize(db, \"userss\") PASSED. Expected an error code -404, got %d", errorCode)
	}

	CreateAdminUser(db)

	expected := 1
	actual := GetTableSize(db, "users")
	if actual != expected {
		t.Errorf("GetTableSize(db, \"users\") FAILED. Expected %d entry, got %d", expected, actual)
	} else {
		t.Logf("GetTableSize(db, \"users\") PASSED. Expected %d entry, got %d", expected, actual)
	}
}

func TestResetTable(t *testing.T) {
	db := MySQLConnect()
	defer db.Close()
	ResetAllTables(db)
	SetUpTables(db)

	ResetTable(db, "Verification")

	expected := 3
	actual := GetDatabaseSize(db)
	if actual != expected {
		t.Errorf("ResetTable(db, \"Verification\") FAILED. Expected %d tables, got %d", expected, actual)
	} else {
		t.Logf("ResetTable(db, \"Verification\") PASSED. Expected %d tables, got %d", expected, actual)
	}
}

func TestResetAllTables(t *testing.T) {
	db := MySQLConnect()
	defer db.Close()
	ResetAllTables(db)
	SetUpTables(db)

	ResetAllTables(db)

	expected := 0
	actual := GetDatabaseSize(db)
	if actual != expected {
		t.Errorf("ResetAllTables(db) FAILED. Expected %d tables, got %d", expected, actual)
	} else {
		t.Logf("ResetAllTables(db) PASSED. Expected %d tables, got %d", expected, actual)
	}
}

func TestCreateNewUser(t *testing.T) {
	db := MySQLConnect()
	defer db.Close()
	ResetAllTables(db)
	SetUpTables(db)
	CreateAdminUser(db)

	username := ""
	password := ""
	email := ""

	errorCode := CreateNewUser(db, username, password, email)
	if errorCode != -204 {
		t.Errorf("CreateNewUser(db, \"\", \"\", \"\") FAILED. Expected an error code -204, got %d", errorCode)
	} else {
		t.Logf("CreateNewUser(db, \"\", \"\", \"\") PASSED. Expected an error code -204, got %d", errorCode)
	}

	//Will always have admin credentials, so tests using them
	//testing duplicate username
	username = "root"
	password = "test"
	email = "unique@gmail.com"

	errorCode = CreateNewUser(db, username, password, email)
	if errorCode != -223 {
		t.Errorf("CreateNewUser(db, \"root\", \"test\", \"unique@gmail.com\") FAILED. Expected an error code -223, got %d", errorCode)
	} else {
		t.Logf("CreateNewUser(db, \"root\", \"test\", \"unique@gmail.com\") PASSED. Expected an error code -223, got %d", errorCode)
	}

	//testing duplicate email
	username = "unique"
	email = "vanbestindustries@gmail.com"

	errorCode = CreateNewUser(db, username, password, email)
	if errorCode != -225 {
		t.Errorf("CreateNewUser(db, \"unique\", \"test\", \"vanbestindustries@gmail.com\") FAILED. Expected an error code -225, got %d", errorCode)
	} else {
		t.Logf("CreateNewUser(db, \"unique\", \"test\", \"vanbestindustries@gmail.com\") PASSED. Expected an error code -225, got %d", errorCode)
	}

	ResetAllTables(db)
	SetUpTables(db)
	CreateAdminUser(db)
	CreateNewUser(db, "testUser", "password", "example@gmail.com")

	expected := 2
	actual := GetTableSize(db, "users")
	if actual != expected {
		t.Errorf("CreateNewUser(db, \"testUser\", \"password\", \"example@gmail.com\") FAILED. Expected %d entries, got %d", expected, actual)
	} else {
		t.Logf("CreateNewUser(db, \"testUser\", \"password\", \"example@gmail.com\") PASSED. Expected %d entries, got %d", expected, actual)
	}
}

func TestCreateAdminUser(t *testing.T) {
	db := MySQLConnect()
	defer db.Close()
	ResetAllTables(db)
	SetUpTables(db)
	CreateAdminUser(db)

	expected := 1
	actual := GetTableSize(db, "users")
	if actual != expected {
		t.Errorf("CreateAdminUser(db) FAILED. Expected %d entry, got %d", expected, actual)
	} else {
		t.Logf("CreateAdminUser(db) PASSED. Expected %d entry, got %d", expected, actual)
	}
}

func TestCreateCommonSubscriptions(t *testing.T) {
	db := MySQLConnect()
	defer db.Close()
	ResetAllTables(db)
	SetUpTables(db)
	CreateCommonSubscriptions(db)

	expected := 46
	actual := GetTableSize(db, "subscriptions")
	if actual != expected {
		t.Errorf("CreateCommonSubscriptions(db) FAILED. Expected %d entries, got %d", expected, actual)
	} else {
		t.Logf("CreateCommonSubscriptions(db) PASSED. Expected %d entries, got %d", expected, actual)
	}
}

func TestGetPassword(t *testing.T) {
	db := MySQLConnect()
	defer db.Close()
	ResetAllTables(db)
	SetUpTables(db)
	CreateAdminUser(db)

	expected := "XohImNooBHFR0OVvjcYpJ3NgPQ1qq73WKhHvch0VQtg="
	actual := GetPassword(db, 1)
	if actual != expected {
		t.Errorf("GetPassword(db, 1) FAILED. Expected password: %s, got %s", expected, actual)
	} else {
		t.Logf("GetPassword(db, 1) PASSED. Expected password: %s, got %s", expected, actual)
	}
}

func TestChangePassword(t *testing.T) {
	db := MySQLConnect()
	defer db.Close()
	ResetAllTables(db)
	SetUpTables(db)
	CreateAdminUser(db)

	userID := 1
	oldPassword := ""
	newPassword := ""

	errorCode := ChangePassword(db, userID, oldPassword, newPassword)
	if errorCode != -204 {
		t.Errorf("ChangePassword(db, 1, \"\", \"\") FAILED. Expected an error code -204, got %d", errorCode)
	} else {
		t.Logf("ChangePassword(db, 1, \"\", \"\") PASSED. Expected an error code -204, got %d", errorCode)
	}

	//Checks if password changed
	oldPassword = GetPassword(db, 1)
	newPassword = "newPassw0rd"

	ChangePassword(db, 1, oldPassword, newPassword)

	expected := "newPassw0rd"
	actual := GetPassword(db, 1)

	if actual != expected {
		t.Errorf("ChangePassword(db, 1, \"XohImNooBHFR0OVvjcYpJ3NgPQ1qq73WKhHvch0VQtg=\", \"newPassw0rd\") FAILED. Expected new password: %s, got %s", expected, actual)
	} else {
		t.Logf("ChangePassword(db, 1, \"XohImNooBHFR0OVvjcYpJ3NgPQ1qq73WKhHvch0VQtg=\", \"newPassw0rd\") PASSED. Expected new password: %s, got %s", expected, actual)
	}
}

func TestGetEmail(t *testing.T) {
	db := MySQLConnect()
	defer db.Close()
	ResetAllTables(db)
	SetUpTables(db)
	CreateAdminUser(db)

	expected := "vanbestindustries@gmail.com"
	actual := GetEmail(db, 1)
	if actual != expected {
		t.Errorf("GetEmail(db, 1) FAILED. Expected email: %s, got %s", expected, actual)
	} else {
		t.Logf("GetEmail(db, 1) PASSED. Expected email: %s, got %s", expected, actual)
	}
}

func TestChangeEmail(t *testing.T) {
	db := MySQLConnect()
	defer db.Close()
	ResetAllTables(db)
	SetUpTables(db)
	CreateAdminUser(db)
	CreateNewUser(db, "test", "doesn't matter", "valekseev2003@gmail.com")

	userID := 1
	//Takes from a user that is not the admin's email, so to pass this test
	//this, the email must already be in the database
	newEmail := "valekseev2003@gmail.com"

	errorCode := ChangeEmail(db, userID, newEmail)
	if errorCode != -223 {
		t.Errorf("ChangeEmail(db, 1, \"valekseev2003@gmail.com\") FAILED. Expected an error code -223 got %d", errorCode)
	} else {
		t.Logf("ChangeEmail(db, 1, \"valekseev2003@gmail.com\") PASSED. Expected an error code -223 got %d", errorCode)
	}

	newEmail = "test@gmail.com"
	ChangeEmail(db, 1, newEmail)

	//Checks if email changed
	expected := newEmail
	actual := GetEmail(db, userID)

	if actual != expected {
		t.Errorf("ChangeEmail(db, 1, \"test@gmail.com\") FAILED. Expected new email: %s, got %s", expected, actual)
	} else {
		t.Logf("ChangeEmail(db, 1, \"test@gmail.com\") PASSED. Expected new email: %s, got %s", expected, actual)
	}
}

func TestGetUsername(t *testing.T) {
	db := MySQLConnect()
	defer db.Close()
	ResetAllTables(db)
	SetUpTables(db)
	CreateAdminUser(db)

	expected := "root"
	actual := GetUsername(db, 1)
	if actual != expected {
		t.Errorf("GetUsername(db, 1) FAILED. Expected username: %s, got %s", expected, actual)
	} else {
		t.Logf("GetUsername(db, 1) PASSED. Expected username: %s, got %s", expected, actual)
	}
}

func TestChangeUsername(t *testing.T) {
	db := MySQLConnect()
	defer db.Close()
	ResetAllTables(db)
	SetUpTables(db)
	CreateAdminUser(db)
	CreateNewUser(db, "valek", "doesn't matter", "valekseev2003@gmail.com")

	userID := 1
	//Takes from a user that is not the admin's username, so to pass this test
	//this, the username must already be in the database
	newUsername := "valek"

	errorCode := ChangeUsername(db, userID, newUsername)
	if errorCode != -223 {
		t.Errorf("ChangeUsername(db, 1, \"valek\") FAILED. Expected an error code -223, got %d", errorCode)
	} else {
		t.Logf("ChangeUsername(db, 1, \"valek\") PASSED. Expected an error code -223, got %d", errorCode)
	}

	newUsername = "userExample"
	ChangeUsername(db, 1, newUsername)

	//Checks if email changed
	expected := newUsername
	actual := GetUsername(db, 1)

	if actual != expected {
		t.Errorf("ChangeUsername(db, 1, \"userExample\") FAILED. Expected new username: %s, got %s", expected, actual)
	} else {
		t.Logf("ChangeUsername(db, 1, \"userExample\") PASSED. Expected new username: %s, got %s", expected, actual)
	}
}

func TestCreateNewSub(t *testing.T) {
	db := MySQLConnect()
	defer db.Close()
	ResetAllTables(db)
	SetUpTables(db)
	CreateNewSub(db, "HBO Max", "9.99")

	expected := 1
	actual := GetTableSize(db, "subscriptions")
	if actual != expected {
		t.Errorf("CreateNewSub(db, \"HBO Max\", \"9.99\") FAILED. Expected %d entry, got %d", expected, actual)
	} else {
		t.Logf("CreateNewSub(db, \"HBO Max\", \"9.99\") PASSED. Expected %d entry, got %d", expected, actual)
	}
}

func TestCreateNewUserSub(t *testing.T) {
	db := MySQLConnect()
	defer db.Close()
	ResetAllTables(db)
	SetUpTables(db)
	CreateAdminUser(db)
	CreateNewSub(db, "HBO Max", "9.99")
	CreateNewUserSub(db, 1, "HBO Max")

	expected := 1
	actual := GetTableSize(db, "usersubs")
	if actual != expected {
		t.Errorf("CreateNewUserSub(db, 1, \"HBO Max\") FAILED. Expected %d entry, got %d", expected, actual)
	} else {
		t.Logf("CreateNewUserSub(db, 1, \"HBO Max\") PASSED. Expected %d entry, got %d", expected, actual)
	}
}

func TestAddOldUserSub(t *testing.T) {
	db := MySQLConnect()
	defer db.Close()
	ResetAllTables(db)
	SetUpTables(db)
	CreateAdminUser(db)
	CreateNewSub(db, "HBO Max", "9.99")
	AddOldUserSub(db, 1, "HBO Max", "2022-01-01 02:30:20", "2022-04-05 04:30:20")

	expected := 1
	actual := GetTableSize(db, "usersubs")
	if actual != expected {
		t.Errorf("AddOldUserSub(db, 1, \"HBO Max\", \"2022-01-01 02:30:20\", \"2022-04-05 04:30:20\") FAILED. Expected %d entry, got %d", expected, actual)
	} else {
		t.Logf("AddOldUserSub(db, 1, \"HBO Max\", \"2022-01-01 02:30:20\", \"2022-04-05 04:30:20\") PASSED. Expected %d entry, got %d", expected, actual)
	}
}

func TestCancelUserSub(t *testing.T) {
	db := MySQLConnect()
	defer db.Close()
	ResetAllTables(db)
	SetUpTables(db)
	CreateAdminUser(db)
	CreateNewSub(db, "HBO Max", "9.99")
	CreateNewUserSub(db, 1, "HBO Max")
	CancelUserSub(db, 1, "HBO Max")

	var cancelledDate string
	db.QueryRow("SELECT DateRemoved FROM UserSubs WHERE UserId = ?", 1).Scan(&cancelledDate)

	//Times are stored in Universal Time, so current time must be converted to test it
	loc, _ := time.LoadLocation("UTC")
	currentTime := time.Now().In(loc)

	expected := currentTime.Format("2006-01-02 15:04:05")
	actual := cancelledDate
	if actual != expected {
		t.Errorf("CancelUserSub(db, 1, \"HBO Max\") FAILED. Expected cancelled date: %s, got %s", expected, actual)
	} else {
		t.Logf("CancelUserSub(db, 1, \"HBO Max\") PASSED. Expected cancelled date: %s, got %s", expected, actual)
	}

}

func TestDeleteUser(t *testing.T) {
	db := MySQLConnect()
	defer db.Close()
	ResetAllTables(db)
	SetUpTables(db)
	CreateAdminUser(db)

	DeleteUser(db, 1)

	expected := 0
	actual := GetTableSize(db, "users")
	if actual != expected {
		t.Errorf("DeleteUser(db, 1) FAILED. Expected %d entries, got %d", expected, actual)
	} else {
		t.Logf("DeleteUser(db, 1) PASSED. Expected %d entries, got %d", expected, actual)
	}
}

func TestLogin(t *testing.T) {
	db := MySQLConnect()
	defer db.Close()
	ResetAllTables(db)
	SetUpTables(db)
	CreateAdminUser(db)

	//Returns UserID if login is successful
	expected := 1
	actual := Login(db, "root", "XohImNooBHFR0OVvjcYpJ3NgPQ1qq73WKhHvch0VQtg=")
	if actual != expected {
		t.Errorf("Login(db, \"root\", \"XohImNooBHFR0OVvjcYpJ3NgPQ1qq73WKhHvch0VQtg=\") FAILED. Expected userID: %d, got %d", expected, actual)
	} else {
		t.Logf("Login(db, \"root\", \"XohImNooBHFR0OVvjcYpJ3NgPQ1qq73WKhHvch0VQtg=\") PASSED. Expected userID: %d, got %d", expected, actual)
	}
}

func TestGetMostUsedSubscription(t *testing.T) {
	db := MySQLConnect()
	defer db.Close()
	ResetAllTables(db)
	SetUpTables(db)
	CreateAdminUser(db)

	CreateNewSub(db, "HBO Max", "9.99")
	AddOldUserSub(db, 1, "HBO Max", "2022-01-01 02:30:20", "2022-01-01 02:30:50")

	expectedName, expectedSubTime := "HBO Max", 30
	actualName, actualSubTime := GetMostUsedSubscription(db, 1, true, false)
	if expectedName != actualName || expectedSubTime != actualSubTime {
		t.Errorf("GetMostUsedSubscription(db, 1, true, false) FAILED. Expected subName: %s, subTimeUsedInSeconds: %d, got %s and %d", expectedName, expectedSubTime, actualName, actualSubTime)
	} else {
		t.Logf("GetMostUsedSubscription(db, 1, true, false) PASSED. Expected subName: %s, subTimeUsedInSeconds: %d, got %s and %d", expectedName, expectedSubTime, actualName, actualSubTime)
	}
}

func TestGetPriceForMonth(t *testing.T) {
	db := MySQLConnect()
	defer db.Close()
	ResetAllTables(db)
	SetUpTables(db)
	CreateAdminUser(db)

	CreateCommonSubscriptions(db)
	AddOldUserSub(db, 1, "HBO Max (With ADS)", "2022-01-01 02:30:20", "2022-01-01 02:30:50")
	AddOldUserSub(db, 1, "Hulu", "2022-01-01 02:30:20", "2022-01-01 02:30:50")
	AddOldUserSub(db, 1, "Amazon Prime", "2022-01-01 02:30:20", "2022-01-01 02:30:50")

	//9.99 + 7.99 + 14.99 = 32.97
	expected := "32.97"
	actual := GetPriceForMonth(db, 1, 1, 2022)
	if actual != expected {
		t.Errorf("GetPriceForMonth(db, 1, 1, 2022) FAILED. Expected cost: %s, got %s", expected, actual)
	} else {
		t.Logf("GetPriceForMonth(db, 1, 1, 2022) PASSED. Expected cost: %s, got %s", expected, actual)
	}
}
