package db

import (
	"context"
	"simple-bank/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:        util.RandomAccountOwner(),
		Balance:      util.RandomAmount(),
		CurrencyCode: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.CurrencyCode, account.CurrencyCode)
	require.True(t, arg.Balance.Equal(account.Balance))

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}
