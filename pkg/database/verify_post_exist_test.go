package database

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

// TestVerifyPostExistHandlerSuccess testa o cenário de sucesso do handler.
func TestVerifyPostExistHandlerSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM posts WHERE post_id = \\$1\\)").
		WithArgs("ABCD").
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

	exists, err := VerifyIfPostExist(db, "ABCD")
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	if !exists {
		t.Errorf("expected post to exist, but it don't")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// TestVerifyPostDontExits testa o cenário de falha do handler.
func TestVerifyPostDontExits(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM posts WHERE post_id = \\$1\\)").
		WithArgs("ABCD").
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

	exists, err := VerifyIfPostExist(db, "ABCD")
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	if exists {
		t.Errorf("expected post to not exist, but it does")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
