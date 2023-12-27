package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:        "Tom",
		Balance:      100,
		CurrencyCode: "GBP",
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.CurrencyCode, account.CurrencyCode)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}