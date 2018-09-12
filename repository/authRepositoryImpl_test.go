package repository_test

import (
	"rest-api/repository"
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestAuthRepositoryImpl_FindUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"userid", "username", "displayname", "password"}).
		AddRow(1, "User1", "User-1", "pwdforuser1")
	mock.ExpectQuery("SELECT userid, username, displayname, password FROM users where username").WithArgs("user1").
		WillReturnRows(rows)

	impl := repository.NewAuthRepositoryImpl(db)
	got, err := impl.FindUser("user1")
	if got.Displayname != "User-1" {
		t.Errorf("Expeting Display Name : %v but got : %v", "User-1", got.Displayname)
	}
}

func TestAuthDaoImpl_FindUser_Fail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"userid", "username", "displayname", "password"})

	mock.ExpectQuery("SELECT userid, username, displayname, password FROM users where username").WithArgs("user1").
		WillReturnRows(rows)

	impl := repository.NewAuthRepositoryImpl(db)
	_, err = impl.FindUser("user1")
	t.Logf("Error is %v ", err)
	if err == nil || err.Error() != "User not found" {
		t.Errorf("Expeting Error : %v but got : %v", "User not found", err)
	}
}
