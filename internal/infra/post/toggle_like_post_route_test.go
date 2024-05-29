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

func TestToggleLikePostRoute_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM users WHERE username = \\$1\\)").
		WithArgs("testuser").
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM posts WHERE post_id = \\$1\\)").
		WithArgs("ABCD").
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM likes WHERE username = \\$1 AND post_id = \\$2\\)").
		WithArgs("testuser", "ABCD").
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

	mock.ExpectExec("INSERT INTO likes \\(username, post_id\\) VALUES \\(\\$1, \\$2\\)").
		WithArgs("testuser", "ABCD").
		WillReturnResult(sqlmock.NewResult(1, 1))

	reqBody, _ := json.Marshal(likePostInfoRequest{
		PostId: "ABCD",
	})
	req, err := http.NewRequest("POST", "/togglelike", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer testtoken")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ToggleLikePostRoute(db, mockVerifyJwt, mockExtractUsernameFromToken))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Post Liked successfully"
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
