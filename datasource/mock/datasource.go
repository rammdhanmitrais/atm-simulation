// Code generated by MockGen. DO NOT EDIT.
// Source: datasource/datasource.go
//
// Generated by this command:
//
//	mockgen -source=datasource/datasource.go -destination=datasource/mock/datasource.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	datasource "atm-simulation/datasource"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockDatasources is a mock of Datasources interface.
type MockDatasources struct {
	ctrl     *gomock.Controller
	recorder *MockDatasourcesMockRecorder
}

// MockDatasourcesMockRecorder is the mock recorder for MockDatasources.
type MockDatasourcesMockRecorder struct {
	mock *MockDatasources
}

// NewMockDatasources creates a new mock instance.
func NewMockDatasources(ctrl *gomock.Controller) *MockDatasources {
	mock := &MockDatasources{ctrl: ctrl}
	mock.recorder = &MockDatasourcesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatasources) EXPECT() *MockDatasourcesMockRecorder {
	return m.recorder
}

// GetLoggedUser mocks base method.
func (m *MockDatasources) GetLoggedUser() (datasource.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLoggedUser")
	ret0, _ := ret[0].(datasource.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLoggedUser indicates an expected call of GetLoggedUser.
func (mr *MockDatasourcesMockRecorder) GetLoggedUser() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoggedUser", reflect.TypeOf((*MockDatasources)(nil).GetLoggedUser))
}

// GetUserByAccountNumber mocks base method.
func (m *MockDatasources) GetUserByAccountNumber(accountNumber string) (datasource.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByAccountNumber", accountNumber)
	ret0, _ := ret[0].(datasource.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByAccountNumber indicates an expected call of GetUserByAccountNumber.
func (mr *MockDatasourcesMockRecorder) GetUserByAccountNumber(accountNumber any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByAccountNumber", reflect.TypeOf((*MockDatasources)(nil).GetUserByAccountNumber), accountNumber)
}

// Login mocks base method.
func (m *MockDatasources) Login(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Login indicates an expected call of Login.
func (mr *MockDatasourcesMockRecorder) Login(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockDatasources)(nil).Login), id)
}

// Logout mocks base method.
func (m *MockDatasources) Logout() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout")
	ret0, _ := ret[0].(error)
	return ret0
}

// Logout indicates an expected call of Logout.
func (mr *MockDatasourcesMockRecorder) Logout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockDatasources)(nil).Logout))
}

// UpdateUserBalance mocks base method.
func (m *MockDatasources) UpdateUserBalance(id int, balance int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserBalance", id, balance)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserBalance indicates an expected call of UpdateUserBalance.
func (mr *MockDatasourcesMockRecorder) UpdateUserBalance(id, balance any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserBalance", reflect.TypeOf((*MockDatasources)(nil).UpdateUserBalance), id, balance)
}
