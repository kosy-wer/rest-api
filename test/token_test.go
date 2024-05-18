package test

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "rest_api/internal/apps/token"
    "github.com/stretchr/testify/assert"
)

// TestLoginHandlerSuccess tests successful login
func TestLoginHandlerSuccess(t *testing.T) {
    user := token.User{Username: "Chek", Password: "123456"}
    jsonUser, _ := json.Marshal(user)
    req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonUser))
    if err != nil {
        t.Fatal(err)
    }
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(token.LoginHandler)
    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code, "Expected status OK")

    // Check if a token is returned
    token := rr.Body.String()
    assert.NotEmpty(t, token, "Expected a token to be returned")
}

// TestLoginHandlerFailure tests failed login
func TestLoginHandlerFailure(t *testing.T) {
    user :=token.User{Username: "Chek", Password: "wrongpassword"}
    jsonUser, _ := json.Marshal(user)
    req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonUser))
    if err != nil {
        t.Fatal(err)
    }
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(token.LoginHandler)
    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusUnauthorized, rr.Code, "Expected status Unauthorized")
}

// TestProtectedHandlerSuccess tests successful access to protected route
func TestProtectedHandlerSuccess(t *testing.T) {
    token_jwt, _ := token.CreateToken("Chek")
    req, err := http.NewRequest("GET", "/protected", nil)
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Authorization", "Bearer "+token_jwt)
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(token.ProtectedHandler)
    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code, "Expected status OK")
    assert.Equal(t, "Welcome to the the protected area", rr.Body.String(), "Expected welcome message")
}

// TestProtectedHandlerFailure tests access to protected route with invalid token
func TestProtectedHandlerFailure(t *testing.T) {
    req, err := http.NewRequest("GET", "/protected", nil)
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Authorization", "Bearer invalidtoken")
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(token.ProtectedHandler)
    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusUnauthorized, rr.Code, "Expected status Unauthorized")
    assert.Equal(t, "Invalid token", rr.Body.String(), "Expected invalid token message")
}
