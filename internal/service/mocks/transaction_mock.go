// Code generated by MockGen. DO NOT EDIT.
// Source: transaction.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/tmrrwnxtsn/const-payments-api/internal/model"
)

// MockTransactionService is a mock of TransactionService interface.
type MockTransactionService struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionServiceMockRecorder
}

// MockTransactionServiceMockRecorder is the mock recorder for MockTransactionService.
type MockTransactionServiceMockRecorder struct {
	mock *MockTransactionService
}

// NewMockTransactionService creates a new mock instance.
func NewMockTransactionService(ctrl *gomock.Controller) *MockTransactionService {
	mock := &MockTransactionService{ctrl: ctrl}
	mock.recorder = &MockTransactionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionService) EXPECT() *MockTransactionServiceMockRecorder {
	return m.recorder
}

// Cancel mocks base method.
func (m *MockTransactionService) Cancel(transactionID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cancel", transactionID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Cancel indicates an expected call of Cancel.
func (mr *MockTransactionServiceMockRecorder) Cancel(transactionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockTransactionService)(nil).Cancel), transactionID)
}

// ChangeStatus mocks base method.
func (m *MockTransactionService) ChangeStatus(transactionID uint64, status model.Status) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeStatus", transactionID, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeStatus indicates an expected call of ChangeStatus.
func (mr *MockTransactionServiceMockRecorder) ChangeStatus(transactionID, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeStatus", reflect.TypeOf((*MockTransactionService)(nil).ChangeStatus), transactionID, status)
}

// Create mocks base method.
func (m *MockTransactionService) Create(transaction model.Transaction) (uint64, model.Status, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", transaction)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(model.Status)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Create indicates an expected call of Create.
func (mr *MockTransactionServiceMockRecorder) Create(transaction interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTransactionService)(nil).Create), transaction)
}

// GetAllByUserEmail mocks base method.
func (m *MockTransactionService) GetAllByUserEmail(userEmail string) ([]model.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByUserEmail", userEmail)
	ret0, _ := ret[0].([]model.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByUserEmail indicates an expected call of GetAllByUserEmail.
func (mr *MockTransactionServiceMockRecorder) GetAllByUserEmail(userEmail interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByUserEmail", reflect.TypeOf((*MockTransactionService)(nil).GetAllByUserEmail), userEmail)
}

// GetAllByUserID mocks base method.
func (m *MockTransactionService) GetAllByUserID(userID uint64) ([]model.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByUserID", userID)
	ret0, _ := ret[0].([]model.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByUserID indicates an expected call of GetAllByUserID.
func (mr *MockTransactionServiceMockRecorder) GetAllByUserID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByUserID", reflect.TypeOf((*MockTransactionService)(nil).GetAllByUserID), userID)
}

// GetStatus mocks base method.
func (m *MockTransactionService) GetStatus(transactionID uint64) (model.Status, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatus", transactionID)
	ret0, _ := ret[0].(model.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStatus indicates an expected call of GetStatus.
func (mr *MockTransactionServiceMockRecorder) GetStatus(transactionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatus", reflect.TypeOf((*MockTransactionService)(nil).GetStatus), transactionID)
}
