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

	// verifies that the currentDB variable has been updated
	if currentDB == nil {
		t.Fatalf("currentDB is nil, expected non-nil")
	}
}

/*func TestSendEmailToAllUsers(t *testing.T) {
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
}*/

func TestDailyReminder(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	SetDB(db)
	// creates a new gin context
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// sets a mock cookie to return a response indicating that emails have already been sent
	c.Request, _ = http.NewRequest("GET", "/", nil)
	cookie := &http.Cookie{Name: "didReminder", Value: "yes"}
	c.Request.AddCookie(cookie)

	// calls the function
	DailyReminder(c)

	// checks the response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %v but got %v", http.StatusOK, w.Code)
	}
	if w.Body.String() != "{\"Success\":\"Emails Already Sent!\"}" {
		t.Errorf("Expected response body %v but got %v", "{\"Success\":\"Emails Already Sent!\"}", w.Body.String())
	}

	// resets the recorder and context
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/", nil)

	// calls the function again without the cookie to return a response that emails have just been sent
	DailyReminder(c)

	// checks the response
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

	// calls the function
	NewsLetter(c)

	// checks the response
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

/*func TestVerifyEmail(t *testing.T) {
	db := ConnectResetAndSetUpDB()
	SetDB(db)

	// Inserts a verification code for a user
	currentTime := time.Now()
	codeGenerator := sha256.New()
	codeGenerator.Write([]byte("test-code"))
	code := base64.URLEncoding.EncodeToString(codeGenerator.Sum(nil))

	_, err := currentDB.Exec("INSERT INTO Verification (UserID, ExpireDate, Code, Type) VALUES (?, ?, ?, 'vE')", 1, currentTime.Add(time.Minute), code)
	if err != nil {
		t.Fatalf("Failed to insert verification code: %v", err)
	}

	// Creates a test context and request with the verification code
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{Key: "code", Value: code})

	// Calls the handler function
	VerifyEmail(c)

	// Checks that the user was verified and the verification code was deleted
	var count int
	err = currentDB.QueryRow("SELECT COUNT(*) FROM Verification WHERE UserID = 1 AND Type = 'vE'").Scan(&count)
	if err != nil {
		t.Fatalf("Failed to query verification codes: %v", err)
	}
	if count != 0 {
		t.Error("Expected verification code to be deleted, but it still exists in the database")
	}
}*/

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
