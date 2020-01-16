package mysql

import (
	"database/sql"
	"database/sql/driver"
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

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestProjectRepo(t *testing.T) {
	suite.Run(t, new(Suite))
}

func fixedSql(s string) string {
	return fmt.Sprintf("^%s$", regexp.QuoteMeta(s))
}

func mockExecSql(s *Suite, sql string, args ...driver.Value) {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(fixedSql(sql)).
		WithArgs(args...).
		WillReturnResult(
			sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
}

func (s *Suite) TestProjectRepo_CreateProject() {
	sql := "INSERT INTO `projects` (`name`,`version`,`created_at`,`updated_at`) VALUES (?,?,?,?)"
	args := []driver.Value{p1.Name, p1.Version, p1.CreatedAt, p1.UpdatedAt}
	mockExecSql(s, sql, args...)

	p2, err := s.repository.CreateProject(p1)
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(p1, p2))
}

func mockQuerySql(s *Suite, sql string, fields []string, projects []*models.Project) {
	rows := sqlmock.NewRows(fields)

	for _, p := range projects {
		rows.AddRow(p.ID, p.Name, p.Version, p.CreatedAt, p.UpdatedAt)
	}
	s.mock.ExpectQuery(fixedSql(sql)).WillReturnRows(rows)
}

func (s *Suite) TestProjectRepo_GetProjects() {
	sql := "SELECT * FROM `projects`"
	fields := []string{"id", "name", "version", "created_at", "updated_at"}
	projects := []*models.Project{p2}

	mockQuerySql(s, sql, fields, projects)

	projects, err := s.repository.GetProjects()
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(projects[0], p2))
}

func (s *Suite) TestProjectRepo_GetByID() {
	sql := "SELECT * FROM `projects`  WHERE (id = ?) ORDER BY `projects`.`id` ASC LIMIT 1"
	fields := []string{"id", "name", "version", "created_at", "updated_at"}
	projects := []*models.Project{p1}
	mockQuerySql(s, sql, fields, projects)
	project, err := s.repository.GetByID("1")

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(project, p1))

}

func (s *Suite) TestProjectRepo_UpdateProject() {
	sql := "UPDATE `projects` SET `created_at` = ?, `id` = ?, `name` = ?, `updated_at` = ?, `version` = ?  WHERE `projects`.`id` = ?"
	args := []driver.Value{sqlmock.AnyArg(), p2.ID, p2.Name, sqlmock.AnyArg(), p2.Version, p2.ID}
	mockExecSql(s, sql, args...)
	project, err := s.repository.UpdateProject(p2)

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(project, p2))

}

func (s *Suite) TestProjectRepo_DeleteProject() {
	sql := "DELETE FROM `projects` WHERE (id = ?)"
	args := []driver.Value{p2.ID}
	mockExecSql(s, sql, args...)
	result, err := s.repository.DeleteProject(p2.ID)

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(true, result))
}
