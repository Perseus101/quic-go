// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/perseus101/quic-go/internal/ackhandler (interfaces: ReceivedPacketHandler)

// Package mockackhandler is a generated GoMock package.
package mockackhandler

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	protocol "github.com/perseus101/quic-go/internal/protocol"
	wire "github.com/perseus101/quic-go/internal/wire"
)

// MockReceivedPacketHandler is a mock of ReceivedPacketHandler interface
type MockReceivedPacketHandler struct {
	ctrl     *gomock.Controller
	recorder *MockReceivedPacketHandlerMockRecorder
}

// MockReceivedPacketHandlerMockRecorder is the mock recorder for MockReceivedPacketHandler
type MockReceivedPacketHandlerMockRecorder struct {
	mock *MockReceivedPacketHandler
}

// NewMockReceivedPacketHandler creates a new mock instance
func NewMockReceivedPacketHandler(ctrl *gomock.Controller) *MockReceivedPacketHandler {
	mock := &MockReceivedPacketHandler{ctrl: ctrl}
	mock.recorder = &MockReceivedPacketHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReceivedPacketHandler) EXPECT() *MockReceivedPacketHandlerMockRecorder {
	return m.recorder
}

// GetAckFrame mocks base method
func (m *MockReceivedPacketHandler) GetAckFrame(arg0 protocol.EncryptionLevel) *wire.AckFrame {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAckFrame", arg0)
	ret0, _ := ret[0].(*wire.AckFrame)
	return ret0
}

// GetAckFrame indicates an expected call of GetAckFrame
func (mr *MockReceivedPacketHandlerMockRecorder) GetAckFrame(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAckFrame", reflect.TypeOf((*MockReceivedPacketHandler)(nil).GetAckFrame), arg0)
}

// GetAlarmTimeout mocks base method
func (m *MockReceivedPacketHandler) GetAlarmTimeout() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAlarmTimeout")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetAlarmTimeout indicates an expected call of GetAlarmTimeout
func (mr *MockReceivedPacketHandlerMockRecorder) GetAlarmTimeout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlarmTimeout", reflect.TypeOf((*MockReceivedPacketHandler)(nil).GetAlarmTimeout))
}

// IgnoreBelow mocks base method
func (m *MockReceivedPacketHandler) IgnoreBelow(arg0 protocol.PacketNumber) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IgnoreBelow", arg0)
}

// IgnoreBelow indicates an expected call of IgnoreBelow
func (mr *MockReceivedPacketHandlerMockRecorder) IgnoreBelow(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IgnoreBelow", reflect.TypeOf((*MockReceivedPacketHandler)(nil).IgnoreBelow), arg0)
}

// ReceivedPacket mocks base method
func (m *MockReceivedPacketHandler) ReceivedPacket(arg0 protocol.PacketNumber, arg1 protocol.EncryptionLevel, arg2 time.Time, arg3 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReceivedPacket", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReceivedPacket indicates an expected call of ReceivedPacket
func (mr *MockReceivedPacketHandlerMockRecorder) ReceivedPacket(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceivedPacket", reflect.TypeOf((*MockReceivedPacketHandler)(nil).ReceivedPacket), arg0, arg1, arg2, arg3)
}
