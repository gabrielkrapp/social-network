package post

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

// TestCreatePostSuccess testa o cenário de sucesso do handler Register.
func TestCreatePostSuccess(t *testing.T) {
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

	err = CreatePost(db, "testuser", "Test Title", "Test Description")
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreatePost_UserNotExists(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM users WHERE username = \\$1\\)").
		WithArgs("testuser").
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

	err = CreatePost(db, "testuser", "testepost", "testepostdescription")
	if err == nil {
		t.Fatalf("expected error, but got nil")
	}

	expectedError := "user does not exist"
	if err.Error() != expectedError {
		t.Fatalf("expected error %v, but got %v", expectedError, err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
