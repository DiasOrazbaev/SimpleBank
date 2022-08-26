package db

import (
	"context"
	"testing"
	"time"

	"github.com/DiasOrazbaev/SimpleBank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntrie(t *testing.T, account *Account) *Entry {
	arg := CreateEntrieParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	in, err := testQueries.CreateEntrie(context.Background(), arg)
	require.NoError(t, err, "Error on creating accout")
	require.NotNil(t, account, "Empty response from db")

	require.Equal(t, arg.AccountID, in.AccountID)
	require.Equal(t, arg.Amount, in.Amount)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return &in
}

func TestCreateEntry(t *testing.T) {
	acc := createRandomAccount(t)
	createRandomEntrie(t, &acc)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	entry1 := createRandomEntrie(t, &account)
	entry2, err := testQueries.GetEntrie(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt.Time, entry2.CreatedAt.Time, time.Second)
}

func TestListEntries(t *testing.T) {
	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntrie(t, &account)
	}

	arg := ListEntrieParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntrie(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
