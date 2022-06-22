package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		FullName:    "tom",
		CountryCode: sql.NullInt32{Int32: 1, Valid: true},
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.CountryCode, user.CountryCode)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
}
