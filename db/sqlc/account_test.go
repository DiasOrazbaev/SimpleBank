package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DiasOrazbaev/SimpleBank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err, "Error on creating accout")
	require.NotNil(t, account, "Empty response from db")

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	input := createRandomAccount(t)
	check, err := testQueries.GetAccount(context.Background(), input.ID)

	require.NoError(t, err, "Error on creating accout")
	require.NotNil(t, check, "Empty response from db")

	require.Equal(t, input.ID, check.ID)
	require.Equal(t, input.Owner, check.Owner)
	require.Equal(t, input.Balance, check.Balance)
	require.Equal(t, input.Currency, check.Currency)
	require.Equal(t, input.CreatedAt, check.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	input := createRandomAccount(t)

	arg := UpdateAccountParams{
		input.ID, util.RandomMoney(),
	}

	err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)

	check, err := testQueries.GetAccount(context.Background(), input.ID)
	require.NoError(t, err)

	require.Equal(t, input.ID, check.ID)
	require.Equal(t, input.Owner, check.Owner)
	require.Equal(t, arg.Balance, check.Balance)
	require.Equal(t, input.Currency, check.Currency)
	require.Equal(t, input.CreatedAt, check.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	input := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), input.ID)

	require.NoError(t, err)

	out, err := testQueries.GetAccount(context.Background(), input.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, out)
}

func TestListAccountss(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccount(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
