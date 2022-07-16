package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	params := makeCreateAccountParams()
	account := createAccountForTest(t, params)
	require.Equal(t, params.Owner, account.Owner)
	require.Equal(t, params.Balance, account.Balance)
	require.Equal(t, params.Currency, account.Currency)

	require.NotZero(t, account.ID)
}

func TestGetAccount(t *testing.T) {
	params := makeCreateAccountParams()
	account1 := createAccountForTest(t, params)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
}

func TestUpdateAccount(t *testing.T) {
	createParams := makeCreateAccountParams()
	account := createAccountForTest(t, createParams)

	newBalance := RandomMoney()
	updateParams := makeUpdateAccountParams(account.ID, newBalance)
	updatedAccount, err := testQueries.UpdateAccount(context.Background(), updateParams)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	require.Equal(t, account.ID, updatedAccount.ID)
	require.Equal(t, account.Owner, updatedAccount.Owner)
	require.Equal(t, newBalance, updatedAccount.Balance)
	require.Equal(t, account.Currency, updatedAccount.Currency)
}

func TestListAccounts(t *testing.T) {
	amt := 5
	for i := 0; i < amt; i++ {
		createAccountForTest(t, makeCreateAccountParams())
	}

	accounts, err := testQueries.ListAccounts(context.Background(), makeListAccountsParams(5, 0))
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	require.Equal(t, amt, len(accounts))

	for _, acc := range accounts {
		require.NotEmpty(t, acc)
	}
}

func TestDeleteAccount(t *testing.T) {
	original := createAccountForTest(t, makeCreateAccountParams())
	err := testQueries.DeleteAccount(context.Background(), original.ID)
	require.NoError(t, err)

	deleted, err := testQueries.GetAccount(context.Background(), original.ID)
	require.Error(t, err)
	require.Empty(t, deleted)
}

func makeCreateAccountParams() CreateAccountParams {
	return CreateAccountParams{
		Owner:    RandomOwner(),
		Balance:  RandomMoney(),
		Currency: RandomCurrency(),
	}
}

func makeUpdateAccountParams(id, balance int64) UpdateAccountParams {
	return UpdateAccountParams{
		ID:      id,
		Balance: balance,
	}
}

func makeListAccountsParams(limit, offset int32) ListAccountsParams {
	return ListAccountsParams{
		Limit:  limit,
		Offset: offset,
	}
}

func createAccountForTest(t *testing.T, params CreateAccountParams) Account {
	account, err := testQueries.CreateAccount(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	return account
}
