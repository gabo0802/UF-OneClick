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
		"username": "root",
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
	currentDB.Exec("INSERT INTO Verification (UserID, Code, ExpireDate, Type) VALUES (?, ?, ?, \"vE\");", 1, possibleCodeEncrypted, expireDate)

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

	//sets up test cookie with user ID of 1 (admin)
	cookie := &http.Cookie{Name: "currentUserID", Value: "1"}
	c.Request = &http.Request{Header: http.Header{"Cookie": []string{cookie.String()}}}

	//call GetAllCurrentUserInfo function
	GetAllCurrentUserInfo(c)

	// Check response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	//checks response body (admin/root's information)
	expectedBody := `{"userid":"","username":"root","password":"","email":"vanbestindustries@gmail.com","subid":"","name":"","price":"","usersubid":"","dateadded":"","dateremoved":"","timezone":"-0400"}`
	if w.Body.String() != expectedBody {
		t.Errorf("Expected response body to be %s, but got %s", expectedBody, w.Body.String())
	}
}

func TestChangeUserPassword(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	SetDB(db)
	//creates a new HTTP request
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

	//sets the currentID value to a valid ID
	currentID = 1

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

func TestSendEmailToAllUsers(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	// sets the current database for the function to use
	SetDB(db)

	// calls the function with test data
	emailSubject := "Test Subject"
	emailMessage := "Test Message"
	result := SendEmailToAllUsers(emailSubject, emailMessage)

	// checks the result
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

func TestNewsLetter(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	SetDB(db)
	// creates a new gin context with a JSON request body
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	requestBody := gin.H{
		"message": "Newsletter message",
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
	req := httptest.NewRequest("POST", "/changetimezone", strings.NewReader(jsonPayload))
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

/*func TestGetAllSubscriptionServices(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	SetDB(db)
	//creates a new recorder to capture the response
	w := httptest.NewRecorder()

	//creates a new gin context with the recorder
	c, _ := gin.CreateTestContext(w)

	//calls the function being tested
	GetAllSubscriptionServices()(c)

	//checks response status code
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	//checks the response body for success or error messages
	expected := `[{"userid":"","username":"","password":"","email":"","subid":"1","name":"Netflix (Basic with ads)","price":"6.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"2","name":"Netflix (Basic)","price":"9.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"3","name":"Netflix (Standard)","price":"15.49","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"4","name":"Netflix (Premium)","price":"19.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"5","name":"Amazon Prime","price":"14.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"6","name":"Amazon Prime (Yearly)","price":"139","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"7","name":"Amazon Prime (Student)","price":"7.49","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"8","name":"Amazon Prime (Student) (Yearly)","price":"69","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"9","name":"Prime Video","price":"8.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"10","name":"Disney+ (Basic)","price":"7.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"11","name":"Disney+ (Premium)","price":"10.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"12","name":"Hulu","price":"7.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"13","name":"Hulu (Student)","price":"1.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"14","name":"Hulu (No Ads)","price":"14.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"15","name":"ESPN+","price":"9.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"16","name":"ESPN+ (Yearly)","price":"99.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"17","name":"Disney Bundle Duo Basic","price":"9.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"18","name":"Disney Bundle Trio Basic","price":"12.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"19","name":"Disney Bundle Trio Premium","price":"19.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"20","name":"HBO Max (With ADS)","price":"9.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"21","name":"HBO Max (AD-Free)","price":"15.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"22","name":"Playstation Plus (Essential)","price":"9.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"23","name":"Playstation Plus (Essential) (3 Months)","price":"24.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"24","name":"Playstation Plus (Essential) (Yearly)","price":"59.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"25","name":"Playstation Plus (Extra)","price":"14.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"26","name":"Playstation Plus (Extra) (3 Months)","price":"39.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"27","name":"Playstation Plus (Extra) (Yearly)","price":"99.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"28","name":"Playstation Plus (Premium)","price":"17.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"29","name":"Playstation Plus (Premium) (3 Months)","price":"49.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"30","name":"Playstation Plus (Premium) (Yearly)","price":"119.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"31","name":"XBOX Live Gold","price":"9.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"32","name":"XBOX Live Gold (3 Months)","price":"24.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"33","name":"XBOX Live Gold (Yearly)","price":"59.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"34","name":"XBOX Game Pass (PC)","price":"9.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"36","name":"XBOX Game Pass (Console)","price":"9.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"37","name":"XBOX Game Pass (Ultimate)","price":"14.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"38","name":"Spotify Premium (Individual)","price":"9.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"39","name":"Spotify Premium (Duo)","price":"12.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"40","name":"Spotify Premium (Family)","price":"15.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"41","name":"Spotify Premium (Student)","price":"4.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"42","name":"Apple Music (Voice)","price":"4.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"43","name":"Apple Music (Student)","price":"5.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"44","name":"Apple Music (Individual)","price":"10.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"45","name":"Apple Music (Family)","price":"16.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"46","name":"AMC+","price":"8.99","dateadded":"","dateremoved":"","timezone":""},{"userid":"","username":"","password":"","email":"","subid":"47","name":"AMC+ (Yearly)","price":"83.88","dateadded":"","dateremoved":"","timezone":""}]`
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}
}*/

func TestGetAllCurrentUserSubscriptions(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	SetDB(db)
	//adds usersubs to test with the admin user
	MySQL.AddOldUserSub(db, 1, "Disney+ (Basic)", "2023-02-15 01:18:56", "2023-03-02 11:45:53")
	//MySQL.AddOldUserSub(db, 1, "Hulu (Student)", "2022-02-01 09:28:33", "2023-01-01 11:48:53")
	// Creates a test context and request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	//sets the user ID cookie to 1 (admin)
	cookie := &http.Cookie{Name: "currentUserID", Value: "1"}
	c.Request = httptest.NewRequest("GET", "/subscriptions", nil)
	c.Request.AddCookie(cookie)

	//calls the handler function with onlyActive set to false
	GetAllCurrentUserSubscriptions(false)(c)

	//checks that the response status code is 200
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, w.Code)
	}

	//TODO: the function outputs the entirety of the userData struct and is an indented json,
	//so need to separate and convert it into only the necessary sub info
	// Checks that the response body contains the expected JSON
	/*expected := `[{"Name":"Disney+ (Basic)","Price":"$12.99","DateAdded":"2023-02-15 01:18:56","DateRemoved":"2023-03-02 11:45:53"}]`
	if w.Body.String() != expected {
		t.Errorf("expected body %q but got %q", expected, w.Body.String())
	}*/
}

func TestGetMostUsedUserSubscription(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	SetDB(db)
	MySQL.AddOldUserSub(db, 1, "Disney+ (Basic)", "2022-01-15 01:20:00", "2023-02-22 12:50:07")
	//creates a new gin context for testing
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	currentID = 1
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
		expectedBody := "{\"Disney+ (Basic)\":\"Active For: 1 year 1 month 1 week 1 day 1 hour 1 minute 1 second \"}"
		if !reflect.DeepEqual(w.Body.String(), expectedBody) {
			t.Errorf("expected body %v but got %v", expectedBody, w.Body.String())
		}
	}
}

func TestDeleteUser(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	SetDB(db)
	//sets up a test Gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	//sets the current user ID to 1 for testing
	currentID = 1

	//calls function
	DeleteUser(c)

	//checks that the current user ID was set to -1
	if currentID != -1 {
		t.Errorf("currentID should be -1 but got %d", currentID)
	}

	//checks the database to see if the user was deleted
	var name string
	err := currentDB.QueryRow("SELECT username FROM users WHERE userid=?", 1).Scan(&name)
	if err != sql.ErrNoRows {
		t.Errorf("Expected user to be deleted, but found name %q", name)
	}
}
