package database

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

// TestVerifyUserExistHandlerSuccess testa o cenário de sucesso do handler Register.
func TestVerifyUserExistHandlerSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM users WHERE username = \\$1\\)").
		WithArgs("testuser").
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

	exists, err := VerifyIfUserExist(db, "testuser")
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	if exists {
		t.Errorf("expected user to not exist, but it does")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
