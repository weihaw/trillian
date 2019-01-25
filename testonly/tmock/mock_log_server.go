// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/google/trillian (interfaces: TrillianLogServer)

// Package tmock is a generated GoMock package.
package tmock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	trillian "github.com/google/trillian"
	reflect "reflect"
)

// MockTrillianLogServer is a mock of TrillianLogServer interface
type MockTrillianLogServer struct {
	ctrl     *gomock.Controller
	recorder *MockTrillianLogServerMockRecorder
}

// MockTrillianLogServerMockRecorder is the mock recorder for MockTrillianLogServer
type MockTrillianLogServerMockRecorder struct {
	mock *MockTrillianLogServer
}

// NewMockTrillianLogServer creates a new mock instance
func NewMockTrillianLogServer(ctrl *gomock.Controller) *MockTrillianLogServer {
	mock := &MockTrillianLogServer{ctrl: ctrl}
	mock.recorder = &MockTrillianLogServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTrillianLogServer) EXPECT() *MockTrillianLogServerMockRecorder {
	return m.recorder
}

// AddSequencedLeaf mocks base method
func (m *MockTrillianLogServer) AddSequencedLeaf(arg0 context.Context, arg1 *trillian.AddSequencedLeafRequest) (*trillian.AddSequencedLeafResponse, error) {
	ret := m.ctrl.Call(m, "AddSequencedLeaf", arg0, arg1)
	ret0, _ := ret[0].(*trillian.AddSequencedLeafResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSequencedLeaf indicates an expected call of AddSequencedLeaf
func (mr *MockTrillianLogServerMockRecorder) AddSequencedLeaf(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSequencedLeaf", reflect.TypeOf((*MockTrillianLogServer)(nil).AddSequencedLeaf), arg0, arg1)
}

// AddSequencedLeaves mocks base method
func (m *MockTrillianLogServer) AddSequencedLeaves(arg0 context.Context, arg1 *trillian.AddSequencedLeavesRequest) (*trillian.AddSequencedLeavesResponse, error) {
	ret := m.ctrl.Call(m, "AddSequencedLeaves", arg0, arg1)
	ret0, _ := ret[0].(*trillian.AddSequencedLeavesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSequencedLeaves indicates an expected call of AddSequencedLeaves
func (mr *MockTrillianLogServerMockRecorder) AddSequencedLeaves(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSequencedLeaves", reflect.TypeOf((*MockTrillianLogServer)(nil).AddSequencedLeaves), arg0, arg1)
}

// GetCertHistory mocks base method
func (m *MockTrillianLogServer) GetCertHistory(arg0 context.Context, arg1 *trillian.GetCertHistoryRequest) (*trillian.GetCertHistoryResponse, error) {
	ret := m.ctrl.Call(m, "GetCertHistory", arg0, arg1)
	ret0, _ := ret[0].(*trillian.GetCertHistoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCertHistory indicates an expected call of GetCertHistory
func (mr *MockTrillianLogServerMockRecorder) GetCertHistory(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertHistory", reflect.TypeOf((*MockTrillianLogServer)(nil).GetCertHistory), arg0, arg1)
}

// GetConsistencyProof mocks base method
func (m *MockTrillianLogServer) GetConsistencyProof(arg0 context.Context, arg1 *trillian.GetConsistencyProofRequest) (*trillian.GetConsistencyProofResponse, error) {
	ret := m.ctrl.Call(m, "GetConsistencyProof", arg0, arg1)
	ret0, _ := ret[0].(*trillian.GetConsistencyProofResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConsistencyProof indicates an expected call of GetConsistencyProof
func (mr *MockTrillianLogServerMockRecorder) GetConsistencyProof(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConsistencyProof", reflect.TypeOf((*MockTrillianLogServer)(nil).GetConsistencyProof), arg0, arg1)
}

// GetEntryAndProof mocks base method
func (m *MockTrillianLogServer) GetEntryAndProof(arg0 context.Context, arg1 *trillian.GetEntryAndProofRequest) (*trillian.GetEntryAndProofResponse, error) {
	ret := m.ctrl.Call(m, "GetEntryAndProof", arg0, arg1)
	ret0, _ := ret[0].(*trillian.GetEntryAndProofResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntryAndProof indicates an expected call of GetEntryAndProof
func (mr *MockTrillianLogServerMockRecorder) GetEntryAndProof(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntryAndProof", reflect.TypeOf((*MockTrillianLogServer)(nil).GetEntryAndProof), arg0, arg1)
}

// GetInclusionProof mocks base method
func (m *MockTrillianLogServer) GetInclusionProof(arg0 context.Context, arg1 *trillian.GetInclusionProofRequest) (*trillian.GetInclusionProofResponse, error) {
	ret := m.ctrl.Call(m, "GetInclusionProof", arg0, arg1)
	ret0, _ := ret[0].(*trillian.GetInclusionProofResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInclusionProof indicates an expected call of GetInclusionProof
func (mr *MockTrillianLogServerMockRecorder) GetInclusionProof(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInclusionProof", reflect.TypeOf((*MockTrillianLogServer)(nil).GetInclusionProof), arg0, arg1)
}

// GetInclusionProofByHash mocks base method
func (m *MockTrillianLogServer) GetInclusionProofByHash(arg0 context.Context, arg1 *trillian.GetInclusionProofByHashRequest) (*trillian.GetInclusionProofByHashResponse, error) {
	ret := m.ctrl.Call(m, "GetInclusionProofByHash", arg0, arg1)
	ret0, _ := ret[0].(*trillian.GetInclusionProofByHashResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInclusionProofByHash indicates an expected call of GetInclusionProofByHash
func (mr *MockTrillianLogServerMockRecorder) GetInclusionProofByHash(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInclusionProofByHash", reflect.TypeOf((*MockTrillianLogServer)(nil).GetInclusionProofByHash), arg0, arg1)
}

// GetLatestSignedLogRoot mocks base method
func (m *MockTrillianLogServer) GetLatestSignedLogRoot(arg0 context.Context, arg1 *trillian.GetLatestSignedLogRootRequest) (*trillian.GetLatestSignedLogRootResponse, error) {
	ret := m.ctrl.Call(m, "GetLatestSignedLogRoot", arg0, arg1)
	ret0, _ := ret[0].(*trillian.GetLatestSignedLogRootResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestSignedLogRoot indicates an expected call of GetLatestSignedLogRoot
func (mr *MockTrillianLogServerMockRecorder) GetLatestSignedLogRoot(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestSignedLogRoot", reflect.TypeOf((*MockTrillianLogServer)(nil).GetLatestSignedLogRoot), arg0, arg1)
}

// GetLeavesByHash mocks base method
func (m *MockTrillianLogServer) GetLeavesByHash(arg0 context.Context, arg1 *trillian.GetLeavesByHashRequest) (*trillian.GetLeavesByHashResponse, error) {
	ret := m.ctrl.Call(m, "GetLeavesByHash", arg0, arg1)
	ret0, _ := ret[0].(*trillian.GetLeavesByHashResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLeavesByHash indicates an expected call of GetLeavesByHash
func (mr *MockTrillianLogServerMockRecorder) GetLeavesByHash(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLeavesByHash", reflect.TypeOf((*MockTrillianLogServer)(nil).GetLeavesByHash), arg0, arg1)
}

// GetLeavesByIndex mocks base method
func (m *MockTrillianLogServer) GetLeavesByIndex(arg0 context.Context, arg1 *trillian.GetLeavesByIndexRequest) (*trillian.GetLeavesByIndexResponse, error) {
	ret := m.ctrl.Call(m, "GetLeavesByIndex", arg0, arg1)
	ret0, _ := ret[0].(*trillian.GetLeavesByIndexResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLeavesByIndex indicates an expected call of GetLeavesByIndex
func (mr *MockTrillianLogServerMockRecorder) GetLeavesByIndex(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLeavesByIndex", reflect.TypeOf((*MockTrillianLogServer)(nil).GetLeavesByIndex), arg0, arg1)
}

// GetLeavesByRange mocks base method
func (m *MockTrillianLogServer) GetLeavesByRange(arg0 context.Context, arg1 *trillian.GetLeavesByRangeRequest) (*trillian.GetLeavesByRangeResponse, error) {
	ret := m.ctrl.Call(m, "GetLeavesByRange", arg0, arg1)
	ret0, _ := ret[0].(*trillian.GetLeavesByRangeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLeavesByRange indicates an expected call of GetLeavesByRange
func (mr *MockTrillianLogServerMockRecorder) GetLeavesByRange(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLeavesByRange", reflect.TypeOf((*MockTrillianLogServer)(nil).GetLeavesByRange), arg0, arg1)
}

// GetSequencedLeafCount mocks base method
func (m *MockTrillianLogServer) GetSequencedLeafCount(arg0 context.Context, arg1 *trillian.GetSequencedLeafCountRequest) (*trillian.GetSequencedLeafCountResponse, error) {
	ret := m.ctrl.Call(m, "GetSequencedLeafCount", arg0, arg1)
	ret0, _ := ret[0].(*trillian.GetSequencedLeafCountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSequencedLeafCount indicates an expected call of GetSequencedLeafCount
func (mr *MockTrillianLogServerMockRecorder) GetSequencedLeafCount(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSequencedLeafCount", reflect.TypeOf((*MockTrillianLogServer)(nil).GetSequencedLeafCount), arg0, arg1)
}

// InitLog mocks base method
func (m *MockTrillianLogServer) InitLog(arg0 context.Context, arg1 *trillian.InitLogRequest) (*trillian.InitLogResponse, error) {
	ret := m.ctrl.Call(m, "InitLog", arg0, arg1)
	ret0, _ := ret[0].(*trillian.InitLogResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InitLog indicates an expected call of InitLog
func (mr *MockTrillianLogServerMockRecorder) InitLog(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitLog", reflect.TypeOf((*MockTrillianLogServer)(nil).InitLog), arg0, arg1)
}

// QueueLeaf mocks base method
func (m *MockTrillianLogServer) QueueLeaf(arg0 context.Context, arg1 *trillian.QueueLeafRequest) (*trillian.QueueLeafResponse, error) {
	ret := m.ctrl.Call(m, "QueueLeaf", arg0, arg1)
	ret0, _ := ret[0].(*trillian.QueueLeafResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueueLeaf indicates an expected call of QueueLeaf
func (mr *MockTrillianLogServerMockRecorder) QueueLeaf(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueLeaf", reflect.TypeOf((*MockTrillianLogServer)(nil).QueueLeaf), arg0, arg1)
}

// QueueLeaves mocks base method
func (m *MockTrillianLogServer) QueueLeaves(arg0 context.Context, arg1 *trillian.QueueLeavesRequest) (*trillian.QueueLeavesResponse, error) {
	ret := m.ctrl.Call(m, "QueueLeaves", arg0, arg1)
	ret0, _ := ret[0].(*trillian.QueueLeavesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueueLeaves indicates an expected call of QueueLeaves
func (mr *MockTrillianLogServerMockRecorder) QueueLeaves(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueLeaves", reflect.TypeOf((*MockTrillianLogServer)(nil).QueueLeaves), arg0, arg1)
}
