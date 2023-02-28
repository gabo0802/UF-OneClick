package MySQL

import "testing"

//TODO: Check if server connection errors work as intended (-502)

// Tests if server error is caught properly once here and omitted in the rest of tests
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

// Tests if error calls work as intended
func TestGetTableSize(t *testing.T) {
	db := MySQLConnect()
	tableName := "userss"
	// Call the function expecting to return -404 if there is an error
	errorCode := GetTableSize(db, tableName)

	// Checks if the error code is working as intended (shows up when the table does not exist)
	if errorCode != -404 {
		t.Errorf("Expected an error code -404, but got %d", errorCode)
	}
}

func TestCreateNewUser(t *testing.T) {
	db := MySQLConnect()

	//ResetAllTables(db)
	//SetUpTables(db)
	//CreateAdminUser(db)

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
}

func TestChangePassword(t *testing.T) {
	db := MySQLConnect()

	//ResetAllTables(db)
	//SetUpTables(db)
	//CreateAdminUser(db)

	userID := 1
	oldPassword := ""
	newPassword := ""

	errorCode := ChangePassword(db, userID, oldPassword, newPassword)
	if errorCode != -204 {
		t.Errorf("Expected an error code -204, but got %d", errorCode)
	}
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
