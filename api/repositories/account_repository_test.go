package repositories

import (
	"database/sql"
	"regexp"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/adisnuhic/hearken/db"
	"github.com/adisnuhic/hearken/models"

	"github.com/jinzhu/gorm"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type AccountRepositorySuite struct {
	suite.Suite
	mock       sqlmock.Sqlmock
	store      db.Store
	repository IAccountRepository
}

func (s *AccountRepositorySuite) SetupTest() {
	var (
		db   *sql.DB
		mock sqlmock.Sqlmock
	)
	db, mock, err := sqlmock.New()
	require.NoError(s.T(), err, "testing db connection error")
	s.mock = mock

	st, err := gorm.Open("mysql", db)
	require.NoError(s.T(), err)
	s.store = st

	s.repository = NewAccountRepository(s.store)

}

func TestAccountRepositorySuite(t *testing.T) {
	suite.Run(t, new(AccountRepositorySuite))
}

func (s *AccountRepositorySuite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *AccountRepositorySuite) Test_Register() {
	// test model
	user := &models.User{
		FirstName:   "Test",
		LastName:    "Test",
		Email:       "test@example.com",
		IsConfirmed: false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// mock query
	query := "INSERT INTO `users` (`first_name`,`last_name`,`email`,`is_confirmed`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?)"
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(user.FirstName, user.LastName, user.Email, user.IsConfirmed, user.CreatedAt, user.UpdatedAt).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	// actual call
	u, err := s.repository.Register(user)

	assert.Nil(s.T(), err, "err should be <nil>, got %v", err)
	assert.Equal(s.T(), user, u, "user model and user should be equal")
}
