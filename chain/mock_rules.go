// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ava-labs/hypersdk/chain (interfaces: Rules)
//
// Generated by this command:
//
//	mockgen -package=chain -destination=chain/mock_rules.go github.com/ava-labs/hypersdk/chain Rules
//

// Package chain is a generated GoMock package.
package chain

import (
	reflect "reflect"

	ids "github.com/ava-labs/avalanchego/ids"
	fees "github.com/ava-labs/hypersdk/fees"
	gomock "go.uber.org/mock/gomock"
)

// MockRules is a mock of Rules interface.
type MockRules struct {
	ctrl     *gomock.Controller
	recorder *MockRulesMockRecorder
}

// MockRulesMockRecorder is the mock recorder for MockRules.
type MockRulesMockRecorder struct {
	mock *MockRules
}

// NewMockRules creates a new mock instance.
func NewMockRules(ctrl *gomock.Controller) *MockRules {
	mock := &MockRules{ctrl: ctrl}
	mock.recorder = &MockRulesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRules) EXPECT() *MockRulesMockRecorder {
	return m.recorder
}

// ChainID mocks base method.
func (m *MockRules) ChainID() ids.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainID")
	ret0, _ := ret[0].(ids.ID)
	return ret0
}

// ChainID indicates an expected call of ChainID.
func (mr *MockRulesMockRecorder) ChainID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainID", reflect.TypeOf((*MockRules)(nil).ChainID))
}

// Fees mocks base method.
func (m *MockRules) Fees() fees.Metadata {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fees")
	ret0, _ := ret[0].(fees.Metadata)
	return ret0
}

// Fees indicates an expected call of Fees.
func (mr *MockRulesMockRecorder) Fees() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fees", reflect.TypeOf((*MockRules)(nil).Fees))
}

// FetchCustom mocks base method.
func (m *MockRules) FetchCustom(arg0 string) (any, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchCustom", arg0)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// FetchCustom indicates an expected call of FetchCustom.
func (mr *MockRulesMockRecorder) FetchCustom(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchCustom", reflect.TypeOf((*MockRules)(nil).FetchCustom), arg0)
}

// GetBaseComputeUnits mocks base method.
func (m *MockRules) GetBaseComputeUnits() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBaseComputeUnits")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetBaseComputeUnits indicates an expected call of GetBaseComputeUnits.
func (mr *MockRulesMockRecorder) GetBaseComputeUnits() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBaseComputeUnits", reflect.TypeOf((*MockRules)(nil).GetBaseComputeUnits))
}

// GetBaseWarpComputeUnits mocks base method.
func (m *MockRules) GetBaseWarpComputeUnits() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBaseWarpComputeUnits")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetBaseWarpComputeUnits indicates an expected call of GetBaseWarpComputeUnits.
func (mr *MockRulesMockRecorder) GetBaseWarpComputeUnits() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBaseWarpComputeUnits", reflect.TypeOf((*MockRules)(nil).GetBaseWarpComputeUnits))
}

// GetMinBlockGap mocks base method.
func (m *MockRules) GetMinBlockGap() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMinBlockGap")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetMinBlockGap indicates an expected call of GetMinBlockGap.
func (mr *MockRulesMockRecorder) GetMinBlockGap() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMinBlockGap", reflect.TypeOf((*MockRules)(nil).GetMinBlockGap))
}

// GetMinEmptyBlockGap mocks base method.
func (m *MockRules) GetMinEmptyBlockGap() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMinEmptyBlockGap")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetMinEmptyBlockGap indicates an expected call of GetMinEmptyBlockGap.
func (mr *MockRulesMockRecorder) GetMinEmptyBlockGap() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMinEmptyBlockGap", reflect.TypeOf((*MockRules)(nil).GetMinEmptyBlockGap))
}

// GetOutgoingWarpComputeUnits mocks base method.
func (m *MockRules) GetOutgoingWarpComputeUnits() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOutgoingWarpComputeUnits")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetOutgoingWarpComputeUnits indicates an expected call of GetOutgoingWarpComputeUnits.
func (mr *MockRulesMockRecorder) GetOutgoingWarpComputeUnits() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOutgoingWarpComputeUnits", reflect.TypeOf((*MockRules)(nil).GetOutgoingWarpComputeUnits))
}

// GetSponsorStateKeysMaxChunks mocks base method.
func (m *MockRules) GetSponsorStateKeysMaxChunks() []uint16 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSponsorStateKeysMaxChunks")
	ret0, _ := ret[0].([]uint16)
	return ret0
}

// GetSponsorStateKeysMaxChunks indicates an expected call of GetSponsorStateKeysMaxChunks.
func (mr *MockRulesMockRecorder) GetSponsorStateKeysMaxChunks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSponsorStateKeysMaxChunks", reflect.TypeOf((*MockRules)(nil).GetSponsorStateKeysMaxChunks))
}

// GetStorageKeyAllocateUnits mocks base method.
func (m *MockRules) GetStorageKeyAllocateUnits() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorageKeyAllocateUnits")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetStorageKeyAllocateUnits indicates an expected call of GetStorageKeyAllocateUnits.
func (mr *MockRulesMockRecorder) GetStorageKeyAllocateUnits() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageKeyAllocateUnits", reflect.TypeOf((*MockRules)(nil).GetStorageKeyAllocateUnits))
}

// GetStorageKeyReadUnits mocks base method.
func (m *MockRules) GetStorageKeyReadUnits() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorageKeyReadUnits")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetStorageKeyReadUnits indicates an expected call of GetStorageKeyReadUnits.
func (mr *MockRulesMockRecorder) GetStorageKeyReadUnits() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageKeyReadUnits", reflect.TypeOf((*MockRules)(nil).GetStorageKeyReadUnits))
}

// GetStorageKeyWriteUnits mocks base method.
func (m *MockRules) GetStorageKeyWriteUnits() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorageKeyWriteUnits")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetStorageKeyWriteUnits indicates an expected call of GetStorageKeyWriteUnits.
func (mr *MockRulesMockRecorder) GetStorageKeyWriteUnits() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageKeyWriteUnits", reflect.TypeOf((*MockRules)(nil).GetStorageKeyWriteUnits))
}

// GetStorageValueAllocateUnits mocks base method.
func (m *MockRules) GetStorageValueAllocateUnits() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorageValueAllocateUnits")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetStorageValueAllocateUnits indicates an expected call of GetStorageValueAllocateUnits.
func (mr *MockRulesMockRecorder) GetStorageValueAllocateUnits() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageValueAllocateUnits", reflect.TypeOf((*MockRules)(nil).GetStorageValueAllocateUnits))
}

// GetStorageValueReadUnits mocks base method.
func (m *MockRules) GetStorageValueReadUnits() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorageValueReadUnits")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetStorageValueReadUnits indicates an expected call of GetStorageValueReadUnits.
func (mr *MockRulesMockRecorder) GetStorageValueReadUnits() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageValueReadUnits", reflect.TypeOf((*MockRules)(nil).GetStorageValueReadUnits))
}

// GetStorageValueWriteUnits mocks base method.
func (m *MockRules) GetStorageValueWriteUnits() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorageValueWriteUnits")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetStorageValueWriteUnits indicates an expected call of GetStorageValueWriteUnits.
func (mr *MockRulesMockRecorder) GetStorageValueWriteUnits() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageValueWriteUnits", reflect.TypeOf((*MockRules)(nil).GetStorageValueWriteUnits))
}

// GetValidityWindow mocks base method.
func (m *MockRules) GetValidityWindow() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidityWindow")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetValidityWindow indicates an expected call of GetValidityWindow.
func (mr *MockRulesMockRecorder) GetValidityWindow() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidityWindow", reflect.TypeOf((*MockRules)(nil).GetValidityWindow))
}

// GetWarpComputeUnitsPerSigner mocks base method.
func (m *MockRules) GetWarpComputeUnitsPerSigner() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWarpComputeUnitsPerSigner")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetWarpComputeUnitsPerSigner indicates an expected call of GetWarpComputeUnitsPerSigner.
func (mr *MockRulesMockRecorder) GetWarpComputeUnitsPerSigner() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWarpComputeUnitsPerSigner", reflect.TypeOf((*MockRules)(nil).GetWarpComputeUnitsPerSigner))
}

// GetWarpConfig mocks base method.
func (m *MockRules) GetWarpConfig(arg0 ids.ID) (bool, uint64, uint64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWarpConfig", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(uint64)
	return ret0, ret1, ret2
}

// GetWarpConfig indicates an expected call of GetWarpConfig.
func (mr *MockRulesMockRecorder) GetWarpConfig(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWarpConfig", reflect.TypeOf((*MockRules)(nil).GetWarpConfig), arg0)
}

// NetworkID mocks base method.
func (m *MockRules) NetworkID() uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetworkID")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// NetworkID indicates an expected call of NetworkID.
func (mr *MockRulesMockRecorder) NetworkID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkID", reflect.TypeOf((*MockRules)(nil).NetworkID))
}
