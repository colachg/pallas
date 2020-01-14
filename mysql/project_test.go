package mysql

import (
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func TestGetProjects(t *testing.T) {
}

func TestGetByID(t *testing.T) {}

func TestCreateProject(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO projects").
		WithArgs("A", "v20200113").
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = db.Exec("INSERT INTO projects(name, version) VALUES (?, ?)", "A", "v20200113")
	if err != nil {
		t.Errorf("error '%s' was not expected, while inserting a row", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestUpdateProject(t *testing.T) {}

func TestDeleteProject(t *testing.T) {}
