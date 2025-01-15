// Code generated by MockGen. DO NOT EDIT.
// Source: ./syncer.go
//
// Generated by this command:
//
//	mockgen -package=metadata -destination=./mocks.go -source=./syncer.go
//

// Package metadata is a generated GoMock package.
package metadata

import (
	reflect "reflect"

	types "github.com/ssvlabs/ssv-spec/types"
	beacon "github.com/ssvlabs/ssv/protocol/v2/blockchain/beacon"
	types0 "github.com/ssvlabs/ssv/protocol/v2/types"
	storage "github.com/ssvlabs/ssv/registry/storage"
	basedb "github.com/ssvlabs/ssv/storage/basedb"
	gomock "go.uber.org/mock/gomock"
)

// MockshareStorage is a mock of shareStorage interface.
type MockshareStorage struct {
	ctrl     *gomock.Controller
	recorder *MockshareStorageMockRecorder
	isgomock struct{}
}

// MockshareStorageMockRecorder is the mock recorder for MockshareStorage.
type MockshareStorageMockRecorder struct {
	mock *MockshareStorage
}

// NewMockshareStorage creates a new mock instance.
func NewMockshareStorage(ctrl *gomock.Controller) *MockshareStorage {
	mock := &MockshareStorage{ctrl: ctrl}
	mock.recorder = &MockshareStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockshareStorage) EXPECT() *MockshareStorageMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockshareStorage) List(txn basedb.Reader, filters ...storage.SharesFilter) []*types0.SSVShare {
	m.ctrl.T.Helper()
	varargs := []any{txn}
	for _, a := range filters {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*types0.SSVShare)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockshareStorageMockRecorder) List(txn any, filters ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{txn}, filters...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockshareStorage)(nil).List), varargs...)
}

// Range mocks base method.
func (m *MockshareStorage) Range(txn basedb.Reader, fn func(*types0.SSVShare) bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Range", txn, fn)
}

// Range indicates an expected call of Range.
func (mr *MockshareStorageMockRecorder) Range(txn, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Range", reflect.TypeOf((*MockshareStorage)(nil).Range), txn, fn)
}

// UpdateValidatorsMetadata mocks base method.
func (m *MockshareStorage) UpdateValidatorsMetadata(arg0 map[types.ValidatorPK]*beacon.ValidatorMetadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateValidatorsMetadata", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateValidatorsMetadata indicates an expected call of UpdateValidatorsMetadata.
func (mr *MockshareStorageMockRecorder) UpdateValidatorsMetadata(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateValidatorsMetadata", reflect.TypeOf((*MockshareStorage)(nil).UpdateValidatorsMetadata), arg0)
}

// MockselfValidatorStore is a mock of selfValidatorStore interface.
type MockselfValidatorStore struct {
	ctrl     *gomock.Controller
	recorder *MockselfValidatorStoreMockRecorder
	isgomock struct{}
}

// MockselfValidatorStoreMockRecorder is the mock recorder for MockselfValidatorStore.
type MockselfValidatorStoreMockRecorder struct {
	mock *MockselfValidatorStore
}

// NewMockselfValidatorStore creates a new mock instance.
func NewMockselfValidatorStore(ctrl *gomock.Controller) *MockselfValidatorStore {
	mock := &MockselfValidatorStore{ctrl: ctrl}
	mock.recorder = &MockselfValidatorStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockselfValidatorStore) EXPECT() *MockselfValidatorStoreMockRecorder {
	return m.recorder
}

// SelfValidators mocks base method.
func (m *MockselfValidatorStore) SelfValidators() []*types0.SSVShare {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelfValidators")
	ret0, _ := ret[0].([]*types0.SSVShare)
	return ret0
}

// SelfValidators indicates an expected call of SelfValidators.
func (mr *MockselfValidatorStoreMockRecorder) SelfValidators() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelfValidators", reflect.TypeOf((*MockselfValidatorStore)(nil).SelfValidators))
}
