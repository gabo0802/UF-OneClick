package handler

import (
	"bytes"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/gabo0802/UF-OneClick/api/httpd/handler/MySQL"
	"github.com/gin-gonic/gin"
)

func ConnectResetAndSetUpDB() *sql.DB {
	//Establishes a connection to the remote MySQL server's database:
	db := MySQL.MySQLConnect()
	MySQL.ResetAllTables(db)
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

	//verifies that the currentDB variable has been updated
	if currentDB == nil {
		t.Fatalf("currentDB is nil, expected non-nil")
	}
}

func TestTryLogin(t *testing.T) {
	//establishes a connection to the database
	db := ConnectResetAndSetUpDB()
	SetDB(db)
	//sets up a test Gin context
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	//sets up a test request body
	login := map[string]string{
		// Uses admin credentials
		"username": "test",
		"password": "password",
	}
	//Marshal returns the JSON encoding of login
	//i.e. golang data turns into JSON
	body, err := json.Marshal(login)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	//creates a test request
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	//calls the TryLogin function
	TryLogin(c)

	//checks the response status code
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

func TestVerifyEmail(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	SetDB(db)

	//sets up a test Gin context
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	//sets up a test request with a valid verification code parameter
	req, _ := http.NewRequest("GET", "/verify", nil)
	q := req.URL.Query()
	var possibleCode string

	//adds temporary verification parameters to the Verification table to test deletion
	codeGenerator := sha256.New()
	codeGenerator.Write([]byte(possibleCode))
	possibleCodeEncrypted := base64.URLEncoding.EncodeToString(codeGenerator.Sum(nil))
	q.Add("code", possibleCodeEncrypted)
	expireDate := time.Now().Add(time.Hour * 24)
	currentDB.Exec("INSERT INTO Verification (UserID, Code, ExpireDate, Type) VALUES (?, ?, ?, \"vE\");", 2, possibleCodeEncrypted, expireDate)

	req.URL.RawQuery = q.Encode()
	c.Request = req

	//calls the VerifyEmail function
	VerifyEmail(c)

	//checks that the response status code is a temporary redirect (307)
	if w.Code != http.StatusTemporaryRedirect {
		t.Errorf("Expected status code %d, but got %d", http.StatusTemporaryRedirect, w.Code)
	}

	//checks that the user was verified and the verification code was deleted from the database
	var count int
	db.QueryRow("SELECT COUNT(*) FROM Verification WHERE Type = 'vE' AND Code = ?", possibleCodeEncrypted).Scan(&count)
	if count != 0 {
		t.Errorf("Expected verification code to be deleted from database, but it still exists")
	}
}

func TestNewUser(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	SetDB(db)
	//creates a new HTTP request with a JSON body
	jsonString := []byte(`{"username":"testuser", "password":"testpassword", "email":"testuser@example.com"}`)
	req, err := http.NewRequest("POST", "/newuser", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	//creates a new recorder to capture the response
	w := httptest.NewRecorder()

	//creates a new gin context with the request and recorder
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	//calls function
	NewUser(c)

	//checks the response status code
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	//checks the response body for success or error messages
	expected := `{"Success":"New User testuser Has Been Created"}`
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}

	//checks if user was actually created in the database
	var count int
	err = currentDB.QueryRow("SELECT COUNT(*) FROM Users WHERE Username = 'testuser'").Scan(&count)
	if err != nil {
		t.Fatalf("Failed to query users table: %v", err)
	}
	if count != 1 {
		t.Error("Expected 1 user to be created, but found", count)
	}
}

func TestGetAllCurrentUserInfo(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	SetDB(db)
	//sets up test Gin context
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	//sets up test cookie with user ID of 2 (test user)
	cookie := &http.Cookie{Name: "currentUserID", Value: "2"}
	c.Request = &http.Request{Header: http.Header{"Cookie": []string{cookie.String()}}}

	//call GetAllCurrentUserInfo function
	GetAllCurrentUserInfo(c)

	// Check response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	//checks response body (admin/root's information)
	expectedBody := `{"userid":"","username":"test","password":"","email":"sir.testmctestington.the.tester2@gmail.com","subid":"","name":"","price":"","usersubid":"","dateadded":"","dateremoved":"","timezone":"-0400"}`
	if w.Body.String() != expectedBody {
		t.Errorf("Expected response body to be %s, but got %s", expectedBody, w.Body.String())
	}
}

func TestChangeUserPassword(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	SetDB(db)
	//creates a new HTTP request with a test JSON payload to satisfy "c.BindJSON(&passwordInfo)" in ChangeUserPassword()
	req, err := http.NewRequest("PUT", "/changepassword", bytes.NewBuffer([]byte(`{
		"oldPassword": "password",
		"newPassword": "updatedPassword"
	}`)))
	if err != nil {
		t.Fatal(err)
	}

	//creates a new response writer
	w := httptest.NewRecorder()

	//sets up test gin context with the request and response writer
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	//sets the currentID value to a valid ID (test user)
	currentID = 2

	//calls the function
	ChangeUserPassword(c)

	//check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusOK)
	}

	//checks the response body
	expected := `{"Success":"Password Changed"}`
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}
}

func TestChangeUserUsername(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	SetDB(db)
	//creates a new HTTP request with a test JSON payload
	req, err := http.NewRequest("PUT", "/changeusername", bytes.NewBuffer([]byte(`{"userid":"","username":"newUser","password":"","email":"sir.testmctestington.the.tester2@gmail.com","subid":"","name":"","price":"","usersubid":"","dateadded":"","dateremoved":"","timezone":"-0400"}`)))
	if err != nil {
		t.Fatal(err)
	}

	//creates a new response writer
	w := httptest.NewRecorder()

	//sets up test gin context with the request and response writer
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	//sets the currentID value to a valid ID (test user)
	currentID = 2

	//calls the function
	ChangeUserUsername(c)

	//check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusOK)
	}

	//checks the response body (should not return anything if successful)
	expected := ""
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}
}

func TestChangeUserEmail(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	SetDB(db)
	//creates a new HTTP request with a test JSON payload
	req, err := http.NewRequest("PUT", "/changeemail", bytes.NewBuffer([]byte(`{"userid":"","username":"root","password":"","email":"sir.testmctestington.the.tester2@gmail.com","subid":"","name":"","price":"","usersubid":"","dateadded":"","dateremoved":"","timezone":"-0400"}`)))
	if err != nil {
		t.Fatal(err)
	}

	//creates a new response writer
	w := httptest.NewRecorder()

	//sets up test gin context with the request and response writer
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	//sets the currentID value to a valid ID (test user)
	currentID = 2

	//calls the function
	ChangeUserEmail(c)

	//check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusOK)
	}

	//checks the response body (should not return anything if successful)
	expected := ""
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}
}

func TestDeleteUser(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	SetDB(db)
	//sets up a test Gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	//sets the current user ID to 2 for testing
	currentID = 2

	//calls function
	DeleteUser(c)

	//checks that the current user ID was set to -1
	if currentID != -1 {
		t.Errorf("currentID should be -1 but got %d", currentID)
	}

	//checks the database to see if the user was deleted
	var name string
	err := currentDB.QueryRow("SELECT username FROM users WHERE userid=?", 2).Scan(&name)
	if err != sql.ErrNoRows {
		t.Errorf("Expected user to be deleted, but found name %q", name)
	}
}

func TestGetAllTimezones(t *testing.T) {
	// create a new mocked gin.Context object
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// call the GetAllTimezones function with the mocked gin.Context
	GetAllTimezones(c)

	// check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Unexpected status code: got %v, expected %v", w.Code, http.StatusOK)
	}

	// check the response body
	var testAllTimezones = []timezoneInfo{{"Eastern Standard Time (EST)", "-0500UTC"}, {"Eastern Daylight Time (EDT)", "-0400UTC"}, {"Central Standard Time (CST)", "-0600UTC"},
		{"Central Daylight Time (CDT)", "-0500UTC"}, {"Pacific Standard Time (PST)", "-0800UTC"}, {"Pacific Daylight Time (PDT)", "-0700UTC"}}
	expectedBody := gin.H{"All Timezones": testAllTimezones}
	actualBody := gin.H{"All Timezones": allTimezones}

	if !reflect.DeepEqual(expectedBody, actualBody) {
		t.Errorf("Unexpected response body:\n\texpected: %v\n\tactual: %v", expectedBody, actualBody)
	}
}

func TestChangeTimezone(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	SetDB(db)
	//creates a mock context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// creates a test JSON payload
	//This request payload is important information in a data block that clients send to the server in the body
	//of an HTTP POST, PUT or PATCH message that contains important information about the request.
	jsonPayload := `{"timezoneDifference": "4"}`

	//sets the request body
	req := httptest.NewRequest("PUT", "/changetimezone", strings.NewReader(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	//calls function
	ChangeTimezone(c)

	//checks response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %v but got %v", http.StatusOK, w.Code)
	}

	//checks if timezone was updated
	if currentTimezone != 4 {
		t.Errorf("Expected timezone to be %v but got %v", 4, currentTimezone)
	}
}

func TestGetAllCurrentUserSubscriptions(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	SetDB(db)
	//adds usersubs to test with the test user
	MySQL.AddOldUserSub(db, 2, "Disney+ (Basic)", "2023-02-15 01:18:56", "2023-03-02 11:45:53")
	//MySQL.AddOldUserSub(db, 1, "Hulu (Student)", "2022-02-01 09:28:33", "2023-01-01 11:48:53")
	// Creates a test context and request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	//sets the user ID cookie to 2 (test user)
	cookie := &http.Cookie{Name: "currentUserID", Value: "2"}
	c.Request = httptest.NewRequest("GET", "/subscriptions", nil)
	c.Request.AddCookie(cookie)

	//calls the handler function with onlyActive set to false
	GetAllCurrentUserSubscriptions(false)(c)

	//checks that the response status code is 200
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, w.Code)
	}
}

func TestGetAllSubscriptionServices(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	SetDB(db)

	// Creates a test context and request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/subscriptions/services", nil)

	// Calls the handler function
	GetAllSubscriptionServices()(c)

	// Checks that the response status code is 200
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, w.Code)
	}
}

func TestNewSubscriptionService(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	SetDB(db)
	currentID = 2
	//creates a new HTTP request with a JSON body
	jsonString := []byte(`{"name":"AppleTV", "price":"4.99"}`)
	req, err := http.NewRequest("POST", "/subscriptions/createsubscription", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	//creates a new recorder to capture the response
	w := httptest.NewRecorder()

	//creates a new gin context with the request and recorder
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	//calls function
	NewSubscriptionService(c)

	//checks the response status code
	if status := w.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	//checks the response body for success or error messages
	expected := `{"Success":"Subscription to AppleTV Created"}`
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}

	//checks if subscription was actually created in the database
	var count int
	err = currentDB.QueryRow("SELECT COUNT(*) FROM Subscriptions WHERE Name = 'AppleTV'").Scan(&count)
	if err != nil {
		t.Fatalf("Failed to query subscriptions table: %v", err)
	}
	if count != 1 {
		t.Error("Expected 1 subscription to be created, but found", count)
	}
}

func TestNewUserSubscription(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	SetDB(db)
	currentID = 2
	//creates a new HTTP request with a JSON body
	jsonString := []byte(`{"name":"Amazon Prime"}`)
	req, err := http.NewRequest("POST", "/subscriptions/addsubscription", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	//creates a new recorder to capture the response
	w := httptest.NewRecorder()

	//creates a new gin context with the request and recorder
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	//calls function
	NewUserSubscription(c)

	//checks the response status code
	if status := w.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	//checks the response body for success or error messages
	expected := `{"Success":"Subscription to Amazon Prime Added"}`
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}

	//checks if usersub was actually created in the database
	var count int
	err = currentDB.QueryRow("SELECT COUNT(*) FROM Usersubs WHERE UserID = '2'").Scan(&count)
	if err != nil {
		t.Fatalf("Failed to query usersubs table: %v", err)
	}
	//there are already 5 usersubs, so if one more is created, the total count will be 6
	if count != 6 {
		t.Error("Expected 1 usersub to be created, but found", count)
	}
}

func TestNewPreviousUserSubscription(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	SetDB(db)
	currentID = 2
	//creates a new HTTP request with a JSON body
	jsonString := []byte(`{"name":"Amazon Prime", "dateadded":"2022-02-01 09:28:33", "dateremoved":"2023-01-01 11:48:53"}`)
	req, err := http.NewRequest("POST", "/subscriptions/addoldsubscription", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	//creates a new recorder to capture the response
	w := httptest.NewRecorder()

	//creates a new gin context with the request and recorder
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	//calls function
	NewPreviousUserSubscription(c)

	//checks the response status code
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	//checks the response body for success or error messages
	expected := `{"Success":"Subscription to Amazon Prime Added"}`
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}

	//checks if usersub was actually created in the database
	var count int
	err = currentDB.QueryRow("SELECT COUNT(*) FROM Usersubs WHERE UserID = '2'").Scan(&count)
	if err != nil {
		t.Fatalf("Failed to query usersubs table: %v", err)
	}
	if count != 6 {
		//there are already 5 usersubs, so if one more is created, the total count will be 6
		t.Error("Expected 1 usersub to be created, but found", count)
	}
}

func TestCancelSubscriptionService(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	SetDB(db)
	//ID #2 is for test user
	currentID = 2
	//creates a new HTTP request with a JSON body
	jsonString := []byte(`{"name":"Disney+ (Basic)"}`)
	req, err := http.NewRequest("POST", "/subscriptions/cancelsubscription", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	//creates a new recorder to capture the response
	w := httptest.NewRecorder()

	//creates a new gin context with the request and recorder
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	//calls function
	CancelSubscriptionService(c)

	//checks the response status code
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	//checks the response body for success or error messages
	//will now have a DateRemoved value in that column
	expected := `{"Success":"Subscription to Disney+ (Basic) Canceled"}`
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}
}

func TestDeleteUserSubID(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	SetDB(db)
	//creates a new recorder to capture the response
	w := httptest.NewRecorder()

	//creates a new gin context with the request and recorder
	c, _ := gin.CreateTestContext(w)
	userSubID := "1"
	c.Params = gin.Params{{Key: "id", Value: userSubID}}

	//calls function
	DeleteUserSubID(c)

	//checks the response status code
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	//checks the response body for success or error messages
	expected := `{"Success":"User Subscription Deleted!"}`
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}
}

func TestGetMostUsedUserSubscription(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	SetDB(db)
	MySQL.AddOldUserSub(db, 2, "Netflix (Basic)", "2019-01-16 01:20:00", "2023-02-22 12:50:07")
	//creates a new gin context for testing
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	currentID = 2
	//calls the GetMostUsedUserSubscription handler
	GetMostUsedUserSubscriptionHandler := GetMostUsedUserSubscription(true, false)

	//checks the response status code
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, w.Code)
	}

	//checks the response body
	expectedBody := gin.H{"Error": "Invalid User ID"}
	if currentID == -1 {
		if !reflect.DeepEqual(w.Body.String(), expectedBody) {
			t.Errorf("expected body %v but got %v", expectedBody, w.Body.String())
		}
	} else {
		//calls the GetMostUsedUserSubscription handler again
		GetMostUsedUserSubscriptionHandler(c)

		//checks the response status code again
		if w.Code != http.StatusOK {
			t.Errorf("expected status code %d but got %d", http.StatusOK, w.Code)
		}

		//checks the response body again
		expectedBody := `{"Netflix (Basic)":"Active For: 4 years 1 month 1 week 1 days 1 hour 1 minute 1 second "}`
		if !reflect.DeepEqual(w.Body.String(), expectedBody) {
			t.Errorf("expected body %v but got %v", expectedBody, w.Body.String())
		}
	}
}

func TestGetAvgPriceofAllCurrentUserSubscriptions(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	SetDB(db)
	//creates a new test context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	//adds up already existing usersubs from test user
	currentID = 2
	onlyActive := true

	//calls function for active subscriptions
	GetAvgPriceofAllCurrentUserSubscriptions(onlyActive)(c)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}

	// Check the response body
	expected := gin.H{"AVG Price: ": "$46.46"}
	var response gin.H
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error unmarshaling response: %v", err)
	}
	if !reflect.DeepEqual(expected, response) {
		t.Errorf("Expected response body %+v but got %+v", expected, response)
	}
}

func TestGetAvgAgeofAllCurrentUserSubscriptionsHandler(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	SetDB(db)
	//removes usersubs based on current time to test since it is not possible to predict ahead of time
	db.Exec("DELETE FROM UserSubs WHERE UserSubID = ?;", 1)
	db.Exec("DELETE FROM UserSubs WHERE UserSubID = ?;", 2)
	db.Exec("DELETE FROM UserSubs WHERE UserSubID = ?;", 3)
	db.Exec("DELETE FROM UserSubs WHERE UserSubID = ?;", 5)
	//adds another usersub that has already been canceled
	db.Exec("INSERT INTO UserSubs(UserID, SubID, DateAdded, DateRemoved) VALUES (?,?,?,?);", "2", "1", "2022-05-01 09:28:33", "2023-01-01 11:48:53")
	//create a test HTTP request
	req, err := http.NewRequest("GET", "/avgageallsubs", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set up the Gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	currentID = 2

	//calls the handler function
	GetAvgAgeofAllCurrentUserSubscriptions(true, false)(c)

	//checks the response status code
	if w.Code != http.StatusOK {
		t.Errorf("expected status OK, but got %d", w.Code)
	}

	//checks the response body
	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	expected := "9 months 2 weeks 1 day 15 hours 58 minutes 26 seconds "
	if response["AVG Age: "] != expected {
		t.Errorf("expected response body '%s', but got '%s'", expected, response["AVG Age: "])
	}
}

func TestGetPriceForMonth(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	SetDB(db)
	//adds another usersub in the same month so it can be tested properly
	//7.99 + 6.99
	db.Exec("INSERT INTO UserSubs(UserID, SubID, DateAdded, DateRemoved) VALUES (?,?,?,?);", "2", "1", "2023-02-05 09:28:33", "2023-03-01 11:48:53")
	//creates a new HTTP request with JSON payload
	reqBody := `{"month": 2, "year": 2023}`
	req, err := http.NewRequest("POST", "/getprice", strings.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	//uses test user ID
	currentID = 2

	//creates a new HTTP response writer
	w := httptest.NewRecorder()

	//creates test Gin context using the response writer and request
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// calls function
	GetPriceForMonth()(c)

	//checks the response status code
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, w.Code)
	}

	//checks the response body
	expectedBody := `{
    "Total Cost": "$14.98"
}`
	if w.Body.String() != expectedBody {
		t.Errorf("expected body %q but got %q", expectedBody, w.Body.String())
	}
}

func TestGetAllPricesInRange(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	SetDB(db)
	//adds another usersub in the same month so it can be tested properly
	db.Exec("INSERT INTO UserSubs(UserID, SubID, DateAdded, DateRemoved) VALUES (?,?,?,?);", "2", "1", "2023-02-05 09:28:33", "2023-03-01 11:48:53")
	//creates a new HTTP request with JSON payload
	reqBody := `{"startmonth": 2, "startyear": 2023, "endmonth": 3, "endyear": 2023}`
	req, err := http.NewRequest("POST", "/getallprices", strings.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	//uses test user ID
	currentID = 2

	//creates a new HTTP response writer
	w := httptest.NewRecorder()

	//creates test Gin context using the response writer and request
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	//calls function
	GetAllPricesInRange()(c)

	//checks the response status code
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, w.Code)
	}
}

func TestNewsLetter(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	SetDB(db)
	//creates a new gin context with a JSON request body
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	requestBody := gin.H{
		"message": "Insert Newsletter message",
	}
	requestBytes, _ := json.Marshal(requestBody)
	requestReader := bytes.NewReader(requestBytes)
	c.Request, _ = http.NewRequest("POST", "/news", requestReader)
	c.Request.Header.Set("Content-Type", "application/json")

	//calls the function
	NewsLetter(c)

	//checks the response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %v but got %v", http.StatusOK, w.Code)
	}
	expectedResponse := gin.H{"Success": "Newsletter Sent!"}
	var actualResponse gin.H
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
	if err != nil {
		t.Errorf("Error parsing response JSON: %v", err)
	}
	if !reflect.DeepEqual(expectedResponse, actualResponse) {
		t.Errorf("Expected response %v but got %v", expectedResponse, actualResponse)
	}
}

func TestSendEmailToAllUsers(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	//sets the current database for the function to use
	SetDB(db)

	//calls the function with test data
	emailSubject := "Test Subject"
	emailMessage := "Test Message"
	result := SendEmailToAllUsers(emailSubject, emailMessage)

	//checks the result
	if !result {
		t.Errorf("SendEmailToAllUsers returned false")
	}
}

func TestDailyReminder(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	SetDB(db)
	//creates a new gin context
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	//sets a mock cookie to return a response indicating that emails have already been sent
	c.Request, _ = http.NewRequest("GET", "/", nil)
	cookie := &http.Cookie{Name: "didReminder", Value: "yes"}
	c.Request.AddCookie(cookie)

	//calls the function
	DailyReminder(c)

	//checks the response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %v but got %v", http.StatusOK, w.Code)
	}
	if w.Body.String() != "{\"Success\":\"Emails Already Sent!\"}" {
		t.Errorf("Expected response body %v but got %v", "{\"Success\":\"Emails Already Sent!\"}", w.Body.String())
	}

	//resets the recorder and context
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/", nil)

	//calls the function again without the cookie to return a response that emails have just been sent
	DailyReminder(c)

	//checks the response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %v but got %v", http.StatusOK, w.Code)
	}
	if w.Body.String() != "{\"Success\":\"Emails Were Sent!\"}" {
		t.Errorf("Expected response body %v but got %v", "{\"Success\":\"Emails Were Sent!\"}", w.Body.String())
	}
}
