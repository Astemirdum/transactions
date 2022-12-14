// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_handler is a generated GoMock package.
package mock_handler

import (
	context "context"
	reflect "reflect"

	service "github.com/Astemirdum/transactions/tx-user/internal/service"
	session "github.com/Astemirdum/transactions/tx-user/internal/session"
	gomock "github.com/golang/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// CheckSessionID mocks base method.
func (m *MockUserService) CheckSessionID(ctx context.Context, id *session.ID) (*session.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckSessionID", ctx, id)
	ret0, _ := ret[0].(*session.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckSessionID indicates an expected call of CheckSessionID.
func (mr *MockUserServiceMockRecorder) CheckSessionID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSessionID", reflect.TypeOf((*MockUserService)(nil).CheckSessionID), ctx, id)
}

// CreateAccount mocks base method.
func (m *MockUserService) CreateAccount(ctx context.Context, user *service.User) (*service.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", ctx, user)
	ret0, _ := ret[0].(*service.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockUserServiceMockRecorder) CreateAccount(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockUserService)(nil).CreateAccount), ctx, user)
}

// DeleteSessionID mocks base method.
func (m *MockUserService) DeleteSessionID(ctx context.Context, id *session.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSessionID", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSessionID indicates an expected call of DeleteSessionID.
func (mr *MockUserServiceMockRecorder) DeleteSessionID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSessionID", reflect.TypeOf((*MockUserService)(nil).DeleteSessionID), ctx, id)
}

// GenerateSessionID mocks base method.
func (m *MockUserService) GenerateSessionID(ctx context.Context, login, password string) (*session.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateSessionID", ctx, login, password)
	ret0, _ := ret[0].(*session.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateSessionID indicates an expected call of GenerateSessionID.
func (mr *MockUserServiceMockRecorder) GenerateSessionID(ctx, login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateSessionID", reflect.TypeOf((*MockUserService)(nil).GenerateSessionID), ctx, login, password)
}

// MockSessionService is a mock of SessionService interface.
type MockSessionService struct {
	ctrl     *gomock.Controller
	recorder *MockSessionServiceMockRecorder
}

// MockSessionServiceMockRecorder is the mock recorder for MockSessionService.
type MockSessionServiceMockRecorder struct {
	mock *MockSessionService
}

// NewMockSessionService creates a new mock instance.
func NewMockSessionService(ctrl *gomock.Controller) *MockSessionService {
	mock := &MockSessionService{ctrl: ctrl}
	mock.recorder = &MockSessionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionService) EXPECT() *MockSessionServiceMockRecorder {
	return m.recorder
}

// CheckSessionID mocks base method.
func (m *MockSessionService) CheckSessionID(ctx context.Context, id *session.ID) (*session.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckSessionID", ctx, id)
	ret0, _ := ret[0].(*session.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckSessionID indicates an expected call of CheckSessionID.
func (mr *MockSessionServiceMockRecorder) CheckSessionID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSessionID", reflect.TypeOf((*MockSessionService)(nil).CheckSessionID), ctx, id)
}

// DeleteSessionID mocks base method.
func (m *MockSessionService) DeleteSessionID(ctx context.Context, id *session.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSessionID", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSessionID indicates an expected call of DeleteSessionID.
func (mr *MockSessionServiceMockRecorder) DeleteSessionID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSessionID", reflect.TypeOf((*MockSessionService)(nil).DeleteSessionID), ctx, id)
}

// GenerateSessionID mocks base method.
func (m *MockSessionService) GenerateSessionID(ctx context.Context, login, password string) (*session.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateSessionID", ctx, login, password)
	ret0, _ := ret[0].(*session.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateSessionID indicates an expected call of GenerateSessionID.
func (mr *MockSessionServiceMockRecorder) GenerateSessionID(ctx, login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateSessionID", reflect.TypeOf((*MockSessionService)(nil).GenerateSessionID), ctx, login, password)
}
