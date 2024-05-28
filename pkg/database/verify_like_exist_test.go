package database

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

// TestVerifyLikeExistHandlerSuccess testa o cenário de sucesso do handler.
func TestVerifyLikeExistHandlerSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM likes WHERE username = \\$1\\ AND post_id = \\$2\\)").
		WithArgs("TesteUser", "ABCD").
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

	exists, err := VerifyIfLikeExist(db, "TesteUser", "ABCD")
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

// TestVerifyLikeDontExits testa o cenário de falha do handler.
func TestVerifyLikeDontExits(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM likes WHERE username = \\$1\\ AND post_id = \\$2\\)").
		WithArgs("TesteUser", "ABCD").
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

	exists, err := VerifyIfLikeExist(db, "TesteUser", "ABCD")
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
