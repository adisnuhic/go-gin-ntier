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

type AuthProviderRepositorySuite struct {
	suite.Suite
	mock       sqlmock.Sqlmock
	store      db.Store
	repository IAuthProviderRepository
}

func (s *AuthProviderRepositorySuite) SetupTest() {
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

	s.repository = NewAuthProviderRepository(s.store)

}

func TestAuthProviderRepositorySuite(t *testing.T) {
	suite.Run(t, new(AuthProviderRepositorySuite))
}

func (s *AuthProviderRepositorySuite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *AuthProviderRepositorySuite) Test_GetByUserID() {
	// test data
	authProvider := &models.AuthProvider{
		Provider:  "local",
		UserID:    1,
		UID:       "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// mock query
	query := "SELECT * FROM `auth_providers` WHERE (user_id = ?)"
	s.mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(authProvider.UserID).
		WillReturnRows(sqlmock.NewRows(
			[]string{"provider", "user_id", "uid", "created_at", "updated_at"}).
			AddRow(authProvider.Provider, authProvider.UserID, authProvider.UID, authProvider.CreatedAt, authProvider.UpdatedAt),
		)

	// actual call
	ap, err := s.repository.GetByUserID(authProvider.UserID)

	assert.Nil(s.T(), err, "err should be <nil>, got %v", err)
	assert.Equal(s.T(), authProvider, ap, "authProvider should be the same as model")

}
