package sqlstore_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/tmrrwnxtsn/const-payments-api/internal/model"
	"github.com/tmrrwnxtsn/const-payments-api/internal/store/sqlstore"
	"github.com/tmrrwnxtsn/const-payments-api/pkg/log"
	"testing"
)

const transactionTable = "transactions"

func TestTransactionRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, dsn)
	defer teardown(transactionTable)

	logger := log.New()

	st := sqlstore.NewStore(db, logger)

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
	db, teardown := sqlstore.TestDB(t, dsn)

	logger := log.New()

	st := sqlstore.NewStore(db, logger)

	transactions, err := st.Transactions().GetAllByUserID(10)
	assert.NoError(t, err)
	assert.EqualValues(t, 0, len(transactions))

	transaction := model.TestTransaction(t)
	_, _ = st.Transactions().Create(*transaction)
	_, _ = st.Transactions().Create(*transaction)

	transactions, err = st.Transactions().GetAllByUserID(transaction.UserID)
	assert.NoError(t, err)
	assert.EqualValues(t, 2, len(transactions))

	teardown(transactionTable)
	transactions, err = st.Transactions().GetAllByUserID(transaction.UserID)
	assert.Error(t, err)
}

func TestTransactionRepository_GetAllByUserEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, dsn)

	logger := log.New()

	st := sqlstore.NewStore(db, logger)

	transactions, err := st.Transactions().GetAllByUserEmail("tmrrwnxtsn@gmail.com")
	assert.NoError(t, err)
	assert.EqualValues(t, 0, len(transactions))

	transaction := model.TestTransaction(t)
	_, _ = st.Transactions().Create(*transaction)
	_, _ = st.Transactions().Create(*transaction)

	transactions, err = st.Transactions().GetAllByUserEmail(transaction.UserEmail)
	assert.NoError(t, err)
	assert.EqualValues(t, 2, len(transactions))

	teardown(transactionTable)
	transactions, err = st.Transactions().GetAllByUserEmail(transaction.UserEmail)
	assert.Error(t, err)
}

func TestTransactionRepository_GetByID(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, dsn)

	logger := log.New()

	st := sqlstore.NewStore(db, logger)

	transactionFound, err := st.Transactions().GetByID(12)
	assert.Error(t, err)
	assert.EqualValues(t, transactionFound, model.Transaction{})

	transaction := model.TestTransaction(t)
	transactionID, _ := st.Transactions().Create(*transaction)

	transactionFound, err = st.Transactions().GetByID(transactionID)
	assert.NoError(t, err)
	assert.NotNil(t, transactionFound)
	assert.EqualValues(t, transaction.Amount, transactionFound.Amount)

	teardown(transactionTable)
	transactionFound, err = st.Transactions().GetByID(transaction.ID)
	assert.Error(t, err)
}

func TestTransactionRepository_ChangeStatus(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, dsn)

	logger := log.New()

	st := sqlstore.NewStore(db, logger)

	transaction := model.TestTransaction(t)
	transactionID, _ := st.Transactions().Create(*transaction)

	err := st.Transactions().ChangeStatus(transactionID, model.StatusSuccess)
	assert.NoError(t, err)

	teardown(transactionTable)
	err = st.Transactions().ChangeStatus(transactionID, model.StatusSuccess)
	assert.Error(t, err)
}

func TestTransactionRepository_Delete(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, dsn)

	logger := log.New()

	st := sqlstore.NewStore(db, logger)
	transaction := model.TestTransaction(t)
	transactionID, _ := st.Transactions().Create(*transaction)

	err := st.Transactions().Delete(transactionID)
	assert.NoError(t, err)

	teardown(transactionTable)
	err = st.Transactions().Delete(transactionID)
	assert.Error(t, err)
}
