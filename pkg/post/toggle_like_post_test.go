package post

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

// TestToggleLike_Success_AddLike testa o cenário de sucesso do handler.
func TestToggleLike_Success_AddLike(t *testing.T) {
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

	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM likes WHERE username = \\$1\\ AND post_id = \\$2\\)").
		WithArgs("testuser", "ABCD").
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

	mock.ExpectExec("INSERT INTO likes \\(username, post_id\\) VALUES \\(\\$1, \\$2\\)").
		WithArgs("testuser", "ABCD").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = ToggleLike(db, "testuser", "ABCD")
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// TestToggleLike_Success_RemoveLike testa o cenário de sucesso do handler.
func TestToggleLike_Success_RemoveLike(t *testing.T) {
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

	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM likes WHERE username = \\$1\\ AND post_id = \\$2\\)").
		WithArgs("testuser", "ABCD").
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

	mock.ExpectExec("DELETE FROM likes WHERE username = \\$1 AND post_id = \\$2").
		WithArgs("testuser", "ABCD").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = ToggleLike(db, "testuser", "ABCD")
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
