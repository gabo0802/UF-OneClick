package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gabo0802/UF-OneClick/api/httpd/handler/MySQL"
	"github.com/gin-gonic/gin"
)

/*func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}*/

func TestTryLogin(t *testing.T) {
	// Establishes a connection to the database
	db := MySQL.MySQLConnect()
	SetDB(db)
	// Set up a test Gin context
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set up a test request body
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

	// Create a test request
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	// Call the TryLogin function
	TryLogin(c)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	// Check the response body
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
