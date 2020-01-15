package mysql

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/colachg/pallas/models"
	"github.com/go-test/deep"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"regexp"
	"testing"
	"time"
)

var (
	t1 = time.Now()
	p1 = &models.Project{
		Name:      "A",
		Version:   "v20200115",
		CreatedAt: t1,
		UpdatedAt: t1,
	}

	p2 = &models.Project{
		ID:        1,
		Name:      "B",
		Version:   "v20200115",
		CreatedAt: t1,
		UpdatedAt: t1,
	}
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository ProjectRepo
	project    models.Project
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("mysql", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)
	s.repository = ProjectRepo{DB: s.DB}
}

func fixedSql(s string) string {
	return fmt.Sprintf("^%s$", regexp.QuoteMeta(s))
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestProjectRepo_CreateProject() {
	sql := "INSERT INTO `projects`"

	s.mock.ExpectBegin()
	s.mock.ExpectExec(sql).
		WithArgs(p1.Name, p1.Version, p1.CreatedAt, p1.UpdatedAt).
		WillReturnResult(
			sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	p2, err := s.repository.CreateProject(p1)

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(p1, p2))
}

func (s *Suite) TestProjectRepo_GetProjects() {
	sql := "SELECT * FROM `projects`"

	rows := sqlmock.NewRows([]string{"id", "name", "version", "created_at", "updated_at"}).
		AddRow(1, p1.Name, p1.Version, p1.CreatedAt, p1.UpdatedAt).
		AddRow(2, "B", "v20200115", t1, t1).
		AddRow(3, "C", "v20200115", t1, t1)

	s.mock.ExpectQuery(fixedSql(sql)).WillReturnRows(rows)
	projects, err := s.repository.GetProjects()

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(projects[0], p1))
}

func (s *Suite) TestProjectRepo_GetByID() {
	sql := "SELECT * FROM `projects`  WHERE (id = ?) ORDER BY `projects`.`id` ASC LIMIT 1"
	rows := sqlmock.NewRows([]string{"id", "name", "version", "created_at", "updated_at"}).
		AddRow(1, p1.Name, p1.Version, p1.CreatedAt, p1.UpdatedAt)

	s.mock.ExpectQuery(fixedSql(sql)).WithArgs("1").WillReturnRows(rows)

	project, err := s.repository.GetByID("1")
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(project, p1))

}

func (s *Suite) TestProjectRepo_UpdateProject() {
	sql := "UPDATE `projects` SET `created_at` = ?, `id` = ?, `name` = ?, `updated_at` = ?, `version` = ?  WHERE `projects`.`id` = ?"
	s.mock.ExpectBegin()
	s.mock.ExpectExec(fixedSql(sql)).
		WithArgs(sqlmock.AnyArg(), p2.ID, p2.Name, sqlmock.AnyArg(), p2.Version, p2.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	project, err := s.repository.UpdateProject(p2)
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(project, p2))

}
