// Code generated by MockGen. DO NOT EDIT.
// Source: secretsmanager.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	secretsmanager "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	gomock "github.com/golang/mock/gomock"
)

// MockSecretsManagerClientAPI is a mock of SecretsManagerClientAPI interface.
type MockSecretsManagerClientAPI struct {
	ctrl     *gomock.Controller
	recorder *MockSecretsManagerClientAPIMockRecorder
}

// MockSecretsManagerClientAPIMockRecorder is the mock recorder for MockSecretsManagerClientAPI.
type MockSecretsManagerClientAPIMockRecorder struct {
	mock *MockSecretsManagerClientAPI
}

// NewMockSecretsManagerClientAPI creates a new mock instance.
func NewMockSecretsManagerClientAPI(ctrl *gomock.Controller) *MockSecretsManagerClientAPI {
	mock := &MockSecretsManagerClientAPI{ctrl: ctrl}
	mock.recorder = &MockSecretsManagerClientAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretsManagerClientAPI) EXPECT() *MockSecretsManagerClientAPIMockRecorder {
	return m.recorder
}

// GetSecretValue mocks base method.
func (m *MockSecretsManagerClientAPI) GetSecretValue(ctx context.Context, params *secretsmanager.GetSecretValueInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.GetSecretValueOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSecretValue", varargs...)
	ret0, _ := ret[0].(*secretsmanager.GetSecretValueOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretValue indicates an expected call of GetSecretValue.
func (mr *MockSecretsManagerClientAPIMockRecorder) GetSecretValue(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretValue", reflect.TypeOf((*MockSecretsManagerClientAPI)(nil).GetSecretValue), varargs...)
}
