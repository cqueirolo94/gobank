package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateTransfer(t *testing.T) {
	params := makeCreateTransferParams(t)
	transfer := createTransferForTest(t, params)
	require.Equal(t, params.FromAccountID, transfer.FromAccountID)
	require.Equal(t, params.ToAccountID, transfer.ToAccountID)
	require.Equal(t, params.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
}

func TestGetTransfer(t *testing.T) {
	params := makeCreateTransferParams(t)
	transfer1 := createTransferForTest(t, params)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
}

func TestListTransfers(t *testing.T) {
	amt := 5
	for i := 0; i < amt; i++ {
		createTransferForTest(t, makeCreateTransferParams(t))
	}

	transfers, err := testQueries.ListTransfers(context.Background(), makeListTransfersParams(5, 0))
	require.NoError(t, err)
	require.NotEmpty(t, transfers)

	require.Equal(t, amt, len(transfers))

	for _, acc := range transfers {
		require.NotEmpty(t, acc)
	}
}

func TestUpdateTransfer(t *testing.T) {
	createParams := makeCreateTransferParams(t)
	original := createTransferForTest(t, createParams)

	newAmount := RandomMoney()
	updateParams := makeUpdateTransferParams(original.ID, newAmount)
	updated, err := testQueries.UpdateTransfer(context.Background(), updateParams)
	require.NoError(t, err)
	require.NotEmpty(t, updated)

	require.Equal(t, original.ID, updated.ID)
	require.Equal(t, original.FromAccountID, updated.FromAccountID)
	require.Equal(t, original.ToAccountID, updated.ToAccountID)
	require.Equal(t, newAmount, updated.Amount)
}

func TestDeleteTransfer(t *testing.T) {
	original := createTransferForTest(t, makeCreateTransferParams(t))
	err := testQueries.DeleteTransfer(context.Background(), original.ID)
	require.NoError(t, err)

	deleted, err := testQueries.GetTransfer(context.Background(), original.ID)
	require.Error(t, err)
	require.Empty(t, deleted)
}

func makeCreateTransferParams(t *testing.T) CreateTransferParams {
	account1 := createAccountForTest(t, makeCreateAccountParams())
	account2 := createAccountForTest(t, makeCreateAccountParams())
	return CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        RandomMoney(),
	}
}

func makeUpdateTransferParams(id, amount int64) UpdateTransferParams {
	return UpdateTransferParams{
		ID:     id,
		Amount: amount,
	}
}

func makeListTransfersParams(limit, offset int32) ListTransfersParams {
	return ListTransfersParams{
		Limit:  limit,
		Offset: offset,
	}
}

func createTransferForTest(t *testing.T, params CreateTransferParams) Transfer {
	transfer, err := testQueries.CreateTransfer(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	return transfer
}
