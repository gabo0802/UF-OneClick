package MySQL

import "testing"

// Test by looking for specific errors returned
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

func TestGetTableSize(t *testing.T) {
	db := MySQLConnect()
	tableName := "userss"
	// Call the function that is expected to not return -404
	errorCode := GetTableSize(db, tableName)

	// Check the error code and error message
	if errorCode == -404 {
		t.Errorf("Error: Table \"%v\" Does Not Exist!\n", tableName)
	}
}
