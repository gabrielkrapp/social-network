package post

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func mockVerifyJwt(tokenString string) (bool, error) {
	return true, nil
}

func mockExtractUsernameFromToken(tokenString string) (string, error) {
	return "testuser", nil
}

func TestCreatePostRoute_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM users WHERE username = \\$1\\)").
		WithArgs("testuser").
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

	mock.ExpectExec("INSERT INTO posts \\(username, title, description\\) VALUES \\(\\$1, \\$2, \\$3\\)").
		WithArgs("testuser", "Test Title", "Test Description").
		WillReturnResult(sqlmock.NewResult(1, 1))

	reqBody, _ := json.Marshal(postInfosRequest{
		Title:       "Test Title",
		Description: "Test Description",
	})
	req, err := http.NewRequest("POST", "/posts", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer testtoken")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreatePostRoute(db, mockVerifyJwt, mockExtractUsernameFromToken))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Post created successfully"
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreatePostRoute_InvalidToken(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	reqBody, _ := json.Marshal(postInfosRequest{
		Title:       "Test Title",
		Description: "Test Description",
	})
	req, err := http.NewRequest("POST", "/posts", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer invalidtoken")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreatePostRoute(db, func(tokenString string) (bool, error) {
		return false, nil
	}, mockExtractUsernameFromToken))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}

	expected := "Invalid token"
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestCreatePostRoute_MissingAuthHeader(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	reqBody, _ := json.Marshal(postInfosRequest{
		Title:       "Test Title",
		Description: "Test Description",
	})
	req, err := http.NewRequest("POST", "/posts", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreatePostRoute(db, mockVerifyJwt, mockExtractUsernameFromToken))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}

	expected := "Authorization header missing"
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestCreatePostRoute_InvalidRequestBody(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	req, err := http.NewRequest("POST", "/posts", bytes.NewBuffer([]byte("invalid body")))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer testtoken")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreatePostRoute(db, mockVerifyJwt, mockExtractUsernameFromToken))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	expected := "Invalid request body"
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
