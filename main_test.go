package main

import (
	"encoding/json"
	"go-gin-student/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupTestRouter() *gin.Engine {
	r := gin.Default()
	routes.SetUpRoutes(r) // Using the SetUpRoutes function from your code to set up routes
	return r
}

func TestGetAllStudents(t *testing.T) {
	// Initialize the router for testing
	r := SetupTestRouter()

	// Create a fake HTTP request to access the "/students" endpoint
	req, _ := http.NewRequest("GET", "/students", nil)

	// Record the HTTP response
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Verify the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the JSON response
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)

	// Verify that the response contains the message "success fetch all students"
	message, exists := response["message"].(string)
	assert.True(t, exists)
	assert.Equal(t, "success fetch all students", message)

	// Verify that the response data is an array
	data, exists := response["data"].([]interface{})
	assert.True(t, exists)
	assert.GreaterOrEqual(t, len(data), 0) // Ensure that the data has at least 0 elements (none or more)

	// You can perform further testing as needed
}
