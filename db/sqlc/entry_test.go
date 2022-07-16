package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateEntry(t *testing.T) {
	params := makeCreateEntryParams(t)
	entry := createEntryForTest(t, params)
	require.Equal(t, params.AccountID, entry.AccountID)
	require.Equal(t, params.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
}

func TestGetEntry(t *testing.T) {
	params := makeCreateEntryParams(t)
	entry1 := createEntryForTest(t, params)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
}

func TestListEntries(t *testing.T) {
	amt := 5
	for i := 0; i < amt; i++ {
		createEntryForTest(t, makeCreateEntryParams(t))
	}

	entries, err := testQueries.ListEntries(context.Background(), makeListEntriesParams(5, 0))
	require.NoError(t, err)
	require.NotEmpty(t, entries)

	require.Equal(t, amt, len(entries))

	for _, acc := range entries {
		require.NotEmpty(t, acc)
	}
}

func TestUpdateEntry(t *testing.T) {
	createParams := makeCreateEntryParams(t)
	entry := createEntryForTest(t, createParams)

	newAmount := RandomMoney()
	updateParams := makeUpdateEntryParams(entry.ID, newAmount)
	updatedEntry, err := testQueries.UpdateEntry(context.Background(), updateParams)
	require.NoError(t, err)
	require.NotEmpty(t, updatedEntry)

	require.Equal(t, entry.ID, updatedEntry.ID)
	require.Equal(t, entry.AccountID, updatedEntry.AccountID)
	require.Equal(t, newAmount, updatedEntry.Amount)
}

func TestDeleteEntry(t *testing.T) {
	original := createEntryForTest(t, makeCreateEntryParams(t))
	err := testQueries.DeleteEntry(context.Background(), original.ID)
	require.NoError(t, err)

	deleted, err := testQueries.GetEntry(context.Background(), original.ID)
	require.Error(t, err)
	require.Empty(t, deleted)
}

func makeCreateEntryParams(t *testing.T) CreateEntryParams {
	account := createAccountForTest(t, makeCreateAccountParams())
	return CreateEntryParams{
		AccountID: account.ID,
		Amount:    RandomMoney(),
	}
}

func makeUpdateEntryParams(id, amount int64) UpdateEntryParams {
	return UpdateEntryParams{
		ID:     id,
		Amount: amount,
	}
}

func makeListEntriesParams(limit, offset int32) ListEntriesParams {
	return ListEntriesParams{
		Limit:  limit,
		Offset: offset,
	}
}

func createEntryForTest(t *testing.T, params CreateEntryParams) Entry {
	entry, err := testQueries.CreateEntry(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	return entry
}
