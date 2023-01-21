// Code generated by MockGen. DO NOT EDIT.
// Source: domain/user.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/loukaspe/auth/mongo-handler/domain"
	bson "go.mongodb.org/mongo-driver/bson"
)

// MockUserServiceInterface is a mock of UserServiceInterface interface.
type MockUserServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceInterfaceMockRecorder
}

// MockUserServiceInterfaceMockRecorder is the mock recorder for MockUserServiceInterface.
type MockUserServiceInterfaceMockRecorder struct {
	mock *MockUserServiceInterface
}

// NewMockUserServiceInterface creates a new mock instance.
func NewMockUserServiceInterface(ctrl *gomock.Controller) *MockUserServiceInterface {
	mock := &MockUserServiceInterface{ctrl: ctrl}
	mock.recorder = &MockUserServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserServiceInterface) EXPECT() *MockUserServiceInterfaceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserServiceInterface) CreateUser(user *domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserServiceInterfaceMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserServiceInterface)(nil).CreateUser), user)
}

// DeleteUser mocks base method.
func (m *MockUserServiceInterface) DeleteUser(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserServiceInterfaceMockRecorder) DeleteUser(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserServiceInterface)(nil).DeleteUser), id)
}

// GetUser mocks base method.
func (m *MockUserServiceInterface) GetUser(id string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", id)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserServiceInterfaceMockRecorder) GetUser(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserServiceInterface)(nil).GetUser), id)
}

// Login mocks base method.
func (m *MockUserServiceInterface) Login(username, password string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", username, password)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUserServiceInterfaceMockRecorder) Login(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserServiceInterface)(nil).Login), username, password)
}

// UpdateUser mocks base method.
func (m *MockUserServiceInterface) UpdateUser(id string, user *domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", id, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserServiceInterfaceMockRecorder) UpdateUser(id, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserServiceInterface)(nil).UpdateUser), id, user)
}

// MockUserDBInterface is a mock of UserDBInterface interface.
type MockUserDBInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserDBInterfaceMockRecorder
}

// MockUserDBInterfaceMockRecorder is the mock recorder for MockUserDBInterface.
type MockUserDBInterfaceMockRecorder struct {
	mock *MockUserDBInterface
}

// NewMockUserDBInterface creates a new mock instance.
func NewMockUserDBInterface(ctrl *gomock.Controller) *MockUserDBInterface {
	mock := &MockUserDBInterface{ctrl: ctrl}
	mock.recorder = &MockUserDBInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserDBInterface) EXPECT() *MockUserDBInterfaceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserDBInterface) CreateUser(arg0 *domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserDBInterfaceMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserDBInterface)(nil).CreateUser), arg0)
}

// DeleteUser mocks base method.
func (m *MockUserDBInterface) DeleteUser(arg0 bson.M) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserDBInterfaceMockRecorder) DeleteUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserDBInterface)(nil).DeleteUser), arg0)
}

// GetUser mocks base method.
func (m *MockUserDBInterface) GetUser(arg0 bson.M) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserDBInterfaceMockRecorder) GetUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserDBInterface)(nil).GetUser), arg0)
}

// UpdateUser mocks base method.
func (m *MockUserDBInterface) UpdateUser(updater bson.D, filter bson.M) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", updater, filter)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserDBInterfaceMockRecorder) UpdateUser(updater, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserDBInterface)(nil).UpdateUser), updater, filter)
}
