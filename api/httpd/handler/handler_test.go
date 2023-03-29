package handler

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gabo0802/UF-OneClick/api/httpd/handler/MySQL"
	"github.com/gin-gonic/gin"
)

func ConnectResetAndSetUpDB() *sql.DB {
	//Establishes a connection to the remote MySQL server's database:
	db := MySQL.MySQLConnect()
	//MySQL.ResetAllTables(db)
	MySQL.SetUpTables(db)

	if MySQL.GetTableSize(db, "Subscriptions") == 0 {
		MySQL.CreateCommonSubscriptions(db)
	}

	if MySQL.GetTableSize(db, "Users") == 0 {
		MySQL.CreateAdminUser(db)
		MySQL.CreateTestUser(db) //for testing
	}

	return db
}

func TestSetDB(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	SetDB(db)

	// verifies that the currentDB variable has been updated
	if currentDB == nil {
		t.Fatalf("currentDB is nil, expected non-nil")
	}
}

func TestSendEmailToAllUsers(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	// set the current database for the function to use
	SetDB(db)

	// call the function with test data
	emailSubject := "Test Subject"
	emailMessage := "Test Message"
	result := SendEmailToAllUsers(emailSubject, emailMessage)

	// check the result
	if !result {
		t.Errorf("SendEmailToAllUsers returned false")
	}
}

func TestTryLogin(t *testing.T) {
	// Establishes a connection to the database
	db := ConnectResetAndSetUpDB()
	SetDB(db)
	// Sets up a test Gin context
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Sets up a test request body
	login := map[string]string{
		// Uses admin credentials
		"username": "root",
		"password": "password",
	}
	//Marshal returns the JSON encoding of login
	//i.e. golang data turns into JSON
	body, err := json.Marshal(login)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	// Creates a test request
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	// Calls the TryLogin function
	TryLogin(c)

	// Checks the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	// Checks the response body
	var responseBody gin.H
	//Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by &responseBody
	//i.e. turns JSON into golang data
	err = json.Unmarshal(w.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	expectedBody := gin.H{"Success": "Logged In"}
	if !reflect.DeepEqual(expectedBody, responseBody) {
		t.Errorf("Expected response body to be %v, but got %v", expectedBody, responseBody)
	}
}

/*func TestGetAllCurrentUserInfo(t *testing.T) {
	router := gin.Default()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/alluserinfo", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	//assert.Equal(t, "pong", w.Body.String())
}*/
