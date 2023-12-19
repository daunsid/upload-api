package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/daunsid/upload-api/internal/db"
	"github.com/daunsid/upload-api/pkg/controller"

	"github.com/stretchr/testify/assert"
)

// MockDB is a mock implementation of the DB interface for testing purposes.
// MockDB is a mock implementation of the DB interface for testing purposes.
type MockDB struct {
	db.DBTX
}

func (m *MockDB) CreateUser(ctx context.Context, params db.CreateUserParams) (db.User, error) {
	// Implement your mock CreateUser logic here
	return db.User{}, nil
}

func (m *MockDB) GetUser(ctx context.Context, username string) (db.User, error) {
	// Implement your mock GetUser logic here
	return db.User{}, nil
}

func TestHandlerCreateUser(t *testing.T) {
	apiConfig := &controller.ApiConfig{
		DB: &MockDB{}, // Use your mock database implementation
	}

	t.Run("CreateUser_Success", func(t *testing.T) {
		// Prepare a request body with valid parameters
		reqBody := map[string]string{
			"username": "newuser",
			"password": "password123",
		}
		reqBodyBytes, err := json.Marshal(reqBody)
		assert.NoError(t, err)

		// Create a request with the prepared body
		req, err := http.NewRequest("POST", "/createuser", bytes.NewReader(reqBodyBytes))
		assert.NoError(t, err)

		// Create a ResponseRecorder to record the response
		rr := httptest.NewRecorder()

		// Serve the HTTP request to the ResponseRecorder
		http.HandlerFunc(apiConfig.HandlerCreateUser).ServeHTTP(rr, req)

		// Check the status code
		assert.Equal(t, http.StatusOK, rr.Code)

		// Check the response body
		var response map[string]interface{}
		err = json.Unmarshal(rr.Body.Bytes(), &response)
		assert.NoError(t, err)

		// Add additional assertions based on your expected response
		assert.Contains(t, response, "user_id")
		assert.Contains(t, response, "username")
	})

	t.Run("CreateUser_UsernameExists", func(t *testing.T) {
		// Prepare a request body with an existing username
		reqBody := map[string]string{
			"username": "existinguser",
			"password": "password123",
		}
		reqBodyBytes, err := json.Marshal(reqBody)
		assert.NoError(t, err)

		// Create a request with the prepared body
		req, err := http.NewRequest("POST", "/createuser", bytes.NewReader(reqBodyBytes))
		assert.NoError(t, err)

		// Create a ResponseRecorder to record the response
		rr := httptest.NewRecorder()

		// Serve the HTTP request to the ResponseRecorder
		http.HandlerFunc(apiConfig.HandlerCreateUser).ServeHTTP(rr, req)

		// Check the status code
		assert.Equal(t, http.StatusBadRequest, rr.Code)

		// Check the response body
		var response map[string]interface{}
		err = json.Unmarshal(rr.Body.Bytes(), &response)
		assert.NoError(t, err)

		// Add additional assertions based on your expected response
		assert.Contains(t, response, "error")
		assert.Equal(t, "Username already exists", response["error"])
	})
}
