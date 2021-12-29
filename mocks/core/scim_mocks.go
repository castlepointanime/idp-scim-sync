// Code generated by MockGen. DO NOT EDIT.
// Source: scim.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/slashdevops/idp-scim-sync/internal/model"
)

// MockSCIMService is a mock of SCIMService interface.
type MockSCIMService struct {
	ctrl     *gomock.Controller
	recorder *MockSCIMServiceMockRecorder
}

// MockSCIMServiceMockRecorder is the mock recorder for MockSCIMService.
type MockSCIMServiceMockRecorder struct {
	mock *MockSCIMService
}

// NewMockSCIMService creates a new mock instance.
func NewMockSCIMService(ctrl *gomock.Controller) *MockSCIMService {
	mock := &MockSCIMService{ctrl: ctrl}
	mock.recorder = &MockSCIMServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSCIMService) EXPECT() *MockSCIMServiceMockRecorder {
	return m.recorder
}

// CreateGroups mocks base method.
func (m *MockSCIMService) CreateGroups(ctx context.Context, gr *model.GroupsResult) (*model.GroupsResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGroups", ctx, gr)
	ret0, _ := ret[0].(*model.GroupsResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGroups indicates an expected call of CreateGroups.
func (mr *MockSCIMServiceMockRecorder) CreateGroups(ctx, gr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGroups", reflect.TypeOf((*MockSCIMService)(nil).CreateGroups), ctx, gr)
}

// CreateGroupsMembers mocks base method.
func (m *MockSCIMService) CreateGroupsMembers(ctx context.Context, gmr *model.GroupsMembersResult) (*model.GroupsMembersResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGroupsMembers", ctx, gmr)
	ret0, _ := ret[0].(*model.GroupsMembersResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGroupsMembers indicates an expected call of CreateGroupsMembers.
func (mr *MockSCIMServiceMockRecorder) CreateGroupsMembers(ctx, gmr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGroupsMembers", reflect.TypeOf((*MockSCIMService)(nil).CreateGroupsMembers), ctx, gmr)
}

// CreateUsers mocks base method.
func (m *MockSCIMService) CreateUsers(ctx context.Context, ur *model.UsersResult) (*model.UsersResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUsers", ctx, ur)
	ret0, _ := ret[0].(*model.UsersResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUsers indicates an expected call of CreateUsers.
func (mr *MockSCIMServiceMockRecorder) CreateUsers(ctx, ur interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUsers", reflect.TypeOf((*MockSCIMService)(nil).CreateUsers), ctx, ur)
}

// DeleteGroups mocks base method.
func (m *MockSCIMService) DeleteGroups(ctx context.Context, gr *model.GroupsResult) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGroups", ctx, gr)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGroups indicates an expected call of DeleteGroups.
func (mr *MockSCIMServiceMockRecorder) DeleteGroups(ctx, gr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGroups", reflect.TypeOf((*MockSCIMService)(nil).DeleteGroups), ctx, gr)
}

// DeleteGroupsMembers mocks base method.
func (m *MockSCIMService) DeleteGroupsMembers(ctx context.Context, gmr *model.GroupsMembersResult) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGroupsMembers", ctx, gmr)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGroupsMembers indicates an expected call of DeleteGroupsMembers.
func (mr *MockSCIMServiceMockRecorder) DeleteGroupsMembers(ctx, gmr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGroupsMembers", reflect.TypeOf((*MockSCIMService)(nil).DeleteGroupsMembers), ctx, gmr)
}

// DeleteUsers mocks base method.
func (m *MockSCIMService) DeleteUsers(ctx context.Context, ur *model.UsersResult) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUsers", ctx, ur)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUsers indicates an expected call of DeleteUsers.
func (mr *MockSCIMServiceMockRecorder) DeleteUsers(ctx, ur interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUsers", reflect.TypeOf((*MockSCIMService)(nil).DeleteUsers), ctx, ur)
}

// GetGroups mocks base method.
func (m *MockSCIMService) GetGroups(ctx context.Context) (*model.GroupsResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroups", ctx)
	ret0, _ := ret[0].(*model.GroupsResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroups indicates an expected call of GetGroups.
func (mr *MockSCIMServiceMockRecorder) GetGroups(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroups", reflect.TypeOf((*MockSCIMService)(nil).GetGroups), ctx)
}

// GetGroupsMembers mocks base method.
func (m *MockSCIMService) GetGroupsMembers(ctx context.Context, gr *model.GroupsResult) (*model.GroupsMembersResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupsMembers", ctx, gr)
	ret0, _ := ret[0].(*model.GroupsMembersResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupsMembers indicates an expected call of GetGroupsMembers.
func (mr *MockSCIMServiceMockRecorder) GetGroupsMembers(ctx, gr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupsMembers", reflect.TypeOf((*MockSCIMService)(nil).GetGroupsMembers), ctx, gr)
}

// GetGroupsMembersBruteForce mocks base method.
func (m *MockSCIMService) GetGroupsMembersBruteForce(ctx context.Context, gr *model.GroupsResult, ur *model.UsersResult) (*model.GroupsMembersResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupsMembersBruteForce", ctx, gr, ur)
	ret0, _ := ret[0].(*model.GroupsMembersResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupsMembersBruteForce indicates an expected call of GetGroupsMembersBruteForce.
func (mr *MockSCIMServiceMockRecorder) GetGroupsMembersBruteForce(ctx, gr, ur interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupsMembersBruteForce", reflect.TypeOf((*MockSCIMService)(nil).GetGroupsMembersBruteForce), ctx, gr, ur)
}

// GetUsers mocks base method.
func (m *MockSCIMService) GetUsers(ctx context.Context) (*model.UsersResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", ctx)
	ret0, _ := ret[0].(*model.UsersResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockSCIMServiceMockRecorder) GetUsers(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockSCIMService)(nil).GetUsers), ctx)
}

// UpdateGroups mocks base method.
func (m *MockSCIMService) UpdateGroups(ctx context.Context, gr *model.GroupsResult) (*model.GroupsResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGroups", ctx, gr)
	ret0, _ := ret[0].(*model.GroupsResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateGroups indicates an expected call of UpdateGroups.
func (mr *MockSCIMServiceMockRecorder) UpdateGroups(ctx, gr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGroups", reflect.TypeOf((*MockSCIMService)(nil).UpdateGroups), ctx, gr)
}

// UpdateUsers mocks base method.
func (m *MockSCIMService) UpdateUsers(ctx context.Context, ur *model.UsersResult) (*model.UsersResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUsers", ctx, ur)
	ret0, _ := ret[0].(*model.UsersResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUsers indicates an expected call of UpdateUsers.
func (mr *MockSCIMServiceMockRecorder) UpdateUsers(ctx, ur interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUsers", reflect.TypeOf((*MockSCIMService)(nil).UpdateUsers), ctx, ur)
}
