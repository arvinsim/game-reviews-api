package repository

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "username", "email", "password_hash"}).
			AddRow(1, "john.doe", "john.doe@gmail.com", "hashedpassword").
			AddRow(2, "jane.doe", "jane.doe@gmail.com", "hashedpassword")

		mock.ExpectQuery("SELECT id, username, email, password_hash FROM users").WillReturnRows(rows)

		users, err := repo.GetAllUsers(context.Background())
		assert.NoError(t, err)
		assert.Len(t, users, 2)
		assert.Equal(t, int64(1), users[0].ID)
		assert.Equal(t, "john.doe", users[0].Username)
		assert.Equal(t, "john.doe@gmail.com", users[0].Email)
		assert.Equal(t, "hashedpassword", users[0].PasswordHash)
		assert.Equal(t, int64(2), users[1].ID)
		assert.Equal(t, "jane.doe", users[1].Username)
		assert.Equal(t, "jane.doe@gmail.com", users[1].Email)
		assert.Equal(t, "hashedpassword", users[1].PasswordHash)
	})

	t.Run("query error", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, username, email, password_hash FROM users").WillReturnError(sql.ErrConnDone)

		users, err := repo.GetAllUsers(context.Background())
		assert.Error(t, err)
		assert.Nil(t, users)
	})

	t.Run("scan error", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "username", "email", "password_hash"}).
			AddRow("invalid_id", "john.doe", "john.doe@gmail.com", "hashedpassword")

		mock.ExpectQuery("SELECT id, username, email, password_hash FROM users").WillReturnRows(rows)

		users, err := repo.GetAllUsers(context.Background())
		assert.Error(t, err)
		assert.Nil(t, users)
	})
}
