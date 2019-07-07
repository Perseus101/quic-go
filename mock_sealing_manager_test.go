// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/perseus101/quic-go (interfaces: SealingManager)

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	handshake "github.com/perseus101/quic-go/internal/handshake"
	protocol "github.com/perseus101/quic-go/internal/protocol"
)

// MockSealingManager is a mock of SealingManager interface
type MockSealingManager struct {
	ctrl     *gomock.Controller
	recorder *MockSealingManagerMockRecorder
}

// MockSealingManagerMockRecorder is the mock recorder for MockSealingManager
type MockSealingManagerMockRecorder struct {
	mock *MockSealingManager
}

// NewMockSealingManager creates a new mock instance
func NewMockSealingManager(ctrl *gomock.Controller) *MockSealingManager {
	mock := &MockSealingManager{ctrl: ctrl}
	mock.recorder = &MockSealingManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSealingManager) EXPECT() *MockSealingManagerMockRecorder {
	return m.recorder
}

// GetSealer mocks base method
func (m *MockSealingManager) GetSealer() (protocol.EncryptionLevel, handshake.Sealer) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSealer")
	ret0, _ := ret[0].(protocol.EncryptionLevel)
	ret1, _ := ret[1].(handshake.Sealer)
	return ret0, ret1
}

// GetSealer indicates an expected call of GetSealer
func (mr *MockSealingManagerMockRecorder) GetSealer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSealer", reflect.TypeOf((*MockSealingManager)(nil).GetSealer))
}

// GetSealerWithEncryptionLevel mocks base method
func (m *MockSealingManager) GetSealerWithEncryptionLevel(arg0 protocol.EncryptionLevel) (handshake.Sealer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSealerWithEncryptionLevel", arg0)
	ret0, _ := ret[0].(handshake.Sealer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSealerWithEncryptionLevel indicates an expected call of GetSealerWithEncryptionLevel
func (mr *MockSealingManagerMockRecorder) GetSealerWithEncryptionLevel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSealerWithEncryptionLevel", reflect.TypeOf((*MockSealingManager)(nil).GetSealerWithEncryptionLevel), arg0)
}
