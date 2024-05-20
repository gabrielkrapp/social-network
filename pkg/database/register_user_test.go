package database

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

// TestRegisterHandlerSuccess testa o cenário de sucesso do handler Register.
func TestRegisterHandlerSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM users WHERE username = \\$1\\)").
		WithArgs("testuser").
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

	mock.ExpectExec("INSERT INTO users").
		WithArgs("testuser", sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = RegisterUser(db, "testuser", "hashedpassword123")
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
