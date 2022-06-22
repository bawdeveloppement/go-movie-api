package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateCountry(t *testing.T) {
	arg := CreateCountryParams{
		Code:          1,
		CountryName:   "France",
		ContinentName: "Europe",
	}

	country, err := testQueries.CreateCountry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, country)

	require.Equal(t, arg.Code, country.Code)
	require.Equal(t, arg.CountryName, country.CountryName)
	require.Equal(t, arg.ContinentName, country.ContinentName)
}
