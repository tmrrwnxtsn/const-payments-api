package teststore_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/tmrrwnxtsn/const-payments-api/internal/model"
	"github.com/tmrrwnxtsn/const-payments-api/internal/store/teststore"
	"testing"
)

func TestTransactionRepository_Create(t *testing.T) {
	st := teststore.NewStore()

	testCases := []struct {
		name        string
		transaction func() model.Transaction
		isValid     bool
	}{
		{
			name: "valid",
			transaction: func() model.Transaction {
				return *model.TestTransaction(t)
			},
			isValid: true,
		},
		{
			name: "invalid",
			transaction: func() model.Transaction {
				return model.Transaction{
					UserEmail:    "tmrrwnxtsn",
					Amount:       -125125.12,
					CurrencyCode: "ruble",
					Status:       model.StatusNew,
				}
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := st.Transactions().Create(tc.transaction())
			if tc.isValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestTransactionRepository_GetAllByUserID(t *testing.T) {
	st := teststore.NewStore()

	transactions, err := st.Transactions().GetAllByUserID(10)
	assert.NoError(t, err)
	assert.EqualValues(t, 0, len(transactions))

	transaction := model.TestTransaction(t)
	_, _ = st.Transactions().Create(*transaction)
	_, _ = st.Transactions().Create(*transaction)

	transactions, err = st.Transactions().GetAllByUserID(transaction.UserID)
	assert.NoError(t, err)
	assert.EqualValues(t, 2, len(transactions))
}

func TestTransactionRepository_GetAllByUserEmail(t *testing.T) {
	st := teststore.NewStore()

	transactions, err := st.Transactions().GetAllByUserEmail("tmrrwnxtsn@gmail.com")
	assert.NoError(t, err)
	assert.EqualValues(t, 0, len(transactions))

	transaction := model.TestTransaction(t)
	_, _ = st.Transactions().Create(*transaction)
	_, _ = st.Transactions().Create(*transaction)

	transactions, err = st.Transactions().GetAllByUserEmail(transaction.UserEmail)
	assert.NoError(t, err)
	assert.EqualValues(t, 2, len(transactions))
}

func TestTransactionRepository_GetByID(t *testing.T) {
	st := teststore.NewStore()

	transactionFound, err := st.Transactions().GetByID(12)
	assert.Error(t, err)
	assert.EqualValues(t, transactionFound, model.Transaction{})

	transaction := model.TestTransaction(t)
	transactionID, _ := st.Transactions().Create(*transaction)

	transactionFound, err = st.Transactions().GetByID(transactionID)
	assert.NoError(t, err)
	assert.NotNil(t, transactionFound)
	assert.EqualValues(t, transaction.Amount, transactionFound.Amount)
}

func TestTransactionRepository_ChangeStatus(t *testing.T) {
	st := teststore.NewStore()

	err := st.Transactions().ChangeStatus(10, model.StatusSuccess)
	assert.Error(t, err)

	transaction := model.TestTransaction(t)
	transactionID, _ := st.Transactions().Create(*transaction)

	err = st.Transactions().ChangeStatus(transactionID, model.StatusSuccess)
	assert.NoError(t, err)

	transactionFound, err := st.Transactions().GetByID(transactionID)
	assert.NoError(t, err)
	assert.EqualValues(t, model.StatusSuccess, transactionFound.Status)
}

func TestTransactionRepository_Delete(t *testing.T) {
	st := teststore.NewStore()
	transaction := model.TestTransaction(t)
	transactionID, _ := st.Transactions().Create(*transaction)

	err := st.Transactions().Delete(transactionID)
	assert.NoError(t, err)

	transactionFound, err := st.Transactions().GetByID(transactionID)
	assert.EqualValues(t, model.Transaction{}, transactionFound)
}
