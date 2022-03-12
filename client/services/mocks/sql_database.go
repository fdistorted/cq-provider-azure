// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cq-provider-azure/client/services (interfaces: SQLDatabaseClient,SQLDatabaseBlobAuditingPoliciesClient,SQLDatabaseThreatDetectionPoliciesClient,SQLDatabaseVulnerabilityAssessmentsClient,SQLDatabaseVulnerabilityAssessmentScansClient,TransparentDataEncryptionsClient,BackupLongTermRetentionPoliciesClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	sql "github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
	gomock "github.com/golang/mock/gomock"
)

// MockSQLDatabaseClient is a mock of SQLDatabaseClient interface.
type MockSQLDatabaseClient struct {
	ctrl     *gomock.Controller
	recorder *MockSQLDatabaseClientMockRecorder
}

// MockSQLDatabaseClientMockRecorder is the mock recorder for MockSQLDatabaseClient.
type MockSQLDatabaseClientMockRecorder struct {
	mock *MockSQLDatabaseClient
}

// NewMockSQLDatabaseClient creates a new mock instance.
func NewMockSQLDatabaseClient(ctrl *gomock.Controller) *MockSQLDatabaseClient {
	mock := &MockSQLDatabaseClient{ctrl: ctrl}
	mock.recorder = &MockSQLDatabaseClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSQLDatabaseClient) EXPECT() *MockSQLDatabaseClientMockRecorder {
	return m.recorder
}

// ListByServer mocks base method.
func (m *MockSQLDatabaseClient) ListByServer(arg0 context.Context, arg1, arg2 string) (sql.DatabaseListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByServer", arg0, arg1, arg2)
	ret0, _ := ret[0].(sql.DatabaseListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByServer indicates an expected call of ListByServer.
func (mr *MockSQLDatabaseClientMockRecorder) ListByServer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByServer", reflect.TypeOf((*MockSQLDatabaseClient)(nil).ListByServer), arg0, arg1, arg2)
}

// MockSQLDatabaseBlobAuditingPoliciesClient is a mock of SQLDatabaseBlobAuditingPoliciesClient interface.
type MockSQLDatabaseBlobAuditingPoliciesClient struct {
	ctrl     *gomock.Controller
	recorder *MockSQLDatabaseBlobAuditingPoliciesClientMockRecorder
}

// MockSQLDatabaseBlobAuditingPoliciesClientMockRecorder is the mock recorder for MockSQLDatabaseBlobAuditingPoliciesClient.
type MockSQLDatabaseBlobAuditingPoliciesClientMockRecorder struct {
	mock *MockSQLDatabaseBlobAuditingPoliciesClient
}

// NewMockSQLDatabaseBlobAuditingPoliciesClient creates a new mock instance.
func NewMockSQLDatabaseBlobAuditingPoliciesClient(ctrl *gomock.Controller) *MockSQLDatabaseBlobAuditingPoliciesClient {
	mock := &MockSQLDatabaseBlobAuditingPoliciesClient{ctrl: ctrl}
	mock.recorder = &MockSQLDatabaseBlobAuditingPoliciesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSQLDatabaseBlobAuditingPoliciesClient) EXPECT() *MockSQLDatabaseBlobAuditingPoliciesClientMockRecorder {
	return m.recorder
}

// ListByDatabase mocks base method.
func (m *MockSQLDatabaseBlobAuditingPoliciesClient) ListByDatabase(arg0 context.Context, arg1, arg2, arg3 string) (sql.DatabaseBlobAuditingPolicyListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByDatabase", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(sql.DatabaseBlobAuditingPolicyListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByDatabase indicates an expected call of ListByDatabase.
func (mr *MockSQLDatabaseBlobAuditingPoliciesClientMockRecorder) ListByDatabase(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByDatabase", reflect.TypeOf((*MockSQLDatabaseBlobAuditingPoliciesClient)(nil).ListByDatabase), arg0, arg1, arg2, arg3)
}

// MockSQLDatabaseThreatDetectionPoliciesClient is a mock of SQLDatabaseThreatDetectionPoliciesClient interface.
type MockSQLDatabaseThreatDetectionPoliciesClient struct {
	ctrl     *gomock.Controller
	recorder *MockSQLDatabaseThreatDetectionPoliciesClientMockRecorder
}

// MockSQLDatabaseThreatDetectionPoliciesClientMockRecorder is the mock recorder for MockSQLDatabaseThreatDetectionPoliciesClient.
type MockSQLDatabaseThreatDetectionPoliciesClientMockRecorder struct {
	mock *MockSQLDatabaseThreatDetectionPoliciesClient
}

// NewMockSQLDatabaseThreatDetectionPoliciesClient creates a new mock instance.
func NewMockSQLDatabaseThreatDetectionPoliciesClient(ctrl *gomock.Controller) *MockSQLDatabaseThreatDetectionPoliciesClient {
	mock := &MockSQLDatabaseThreatDetectionPoliciesClient{ctrl: ctrl}
	mock.recorder = &MockSQLDatabaseThreatDetectionPoliciesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSQLDatabaseThreatDetectionPoliciesClient) EXPECT() *MockSQLDatabaseThreatDetectionPoliciesClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockSQLDatabaseThreatDetectionPoliciesClient) Get(arg0 context.Context, arg1, arg2, arg3 string) (sql.DatabaseSecurityAlertPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(sql.DatabaseSecurityAlertPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSQLDatabaseThreatDetectionPoliciesClientMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSQLDatabaseThreatDetectionPoliciesClient)(nil).Get), arg0, arg1, arg2, arg3)
}

// MockSQLDatabaseVulnerabilityAssessmentsClient is a mock of SQLDatabaseVulnerabilityAssessmentsClient interface.
type MockSQLDatabaseVulnerabilityAssessmentsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSQLDatabaseVulnerabilityAssessmentsClientMockRecorder
}

// MockSQLDatabaseVulnerabilityAssessmentsClientMockRecorder is the mock recorder for MockSQLDatabaseVulnerabilityAssessmentsClient.
type MockSQLDatabaseVulnerabilityAssessmentsClientMockRecorder struct {
	mock *MockSQLDatabaseVulnerabilityAssessmentsClient
}

// NewMockSQLDatabaseVulnerabilityAssessmentsClient creates a new mock instance.
func NewMockSQLDatabaseVulnerabilityAssessmentsClient(ctrl *gomock.Controller) *MockSQLDatabaseVulnerabilityAssessmentsClient {
	mock := &MockSQLDatabaseVulnerabilityAssessmentsClient{ctrl: ctrl}
	mock.recorder = &MockSQLDatabaseVulnerabilityAssessmentsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSQLDatabaseVulnerabilityAssessmentsClient) EXPECT() *MockSQLDatabaseVulnerabilityAssessmentsClientMockRecorder {
	return m.recorder
}

// ListByDatabase mocks base method.
func (m *MockSQLDatabaseVulnerabilityAssessmentsClient) ListByDatabase(arg0 context.Context, arg1, arg2, arg3 string) (sql.DatabaseVulnerabilityAssessmentListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByDatabase", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(sql.DatabaseVulnerabilityAssessmentListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByDatabase indicates an expected call of ListByDatabase.
func (mr *MockSQLDatabaseVulnerabilityAssessmentsClientMockRecorder) ListByDatabase(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByDatabase", reflect.TypeOf((*MockSQLDatabaseVulnerabilityAssessmentsClient)(nil).ListByDatabase), arg0, arg1, arg2, arg3)
}

// MockSQLDatabaseVulnerabilityAssessmentScansClient is a mock of SQLDatabaseVulnerabilityAssessmentScansClient interface.
type MockSQLDatabaseVulnerabilityAssessmentScansClient struct {
	ctrl     *gomock.Controller
	recorder *MockSQLDatabaseVulnerabilityAssessmentScansClientMockRecorder
}

// MockSQLDatabaseVulnerabilityAssessmentScansClientMockRecorder is the mock recorder for MockSQLDatabaseVulnerabilityAssessmentScansClient.
type MockSQLDatabaseVulnerabilityAssessmentScansClientMockRecorder struct {
	mock *MockSQLDatabaseVulnerabilityAssessmentScansClient
}

// NewMockSQLDatabaseVulnerabilityAssessmentScansClient creates a new mock instance.
func NewMockSQLDatabaseVulnerabilityAssessmentScansClient(ctrl *gomock.Controller) *MockSQLDatabaseVulnerabilityAssessmentScansClient {
	mock := &MockSQLDatabaseVulnerabilityAssessmentScansClient{ctrl: ctrl}
	mock.recorder = &MockSQLDatabaseVulnerabilityAssessmentScansClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSQLDatabaseVulnerabilityAssessmentScansClient) EXPECT() *MockSQLDatabaseVulnerabilityAssessmentScansClientMockRecorder {
	return m.recorder
}

// ListByDatabase mocks base method.
func (m *MockSQLDatabaseVulnerabilityAssessmentScansClient) ListByDatabase(arg0 context.Context, arg1, arg2, arg3 string) (sql.VulnerabilityAssessmentScanRecordListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByDatabase", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(sql.VulnerabilityAssessmentScanRecordListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByDatabase indicates an expected call of ListByDatabase.
func (mr *MockSQLDatabaseVulnerabilityAssessmentScansClientMockRecorder) ListByDatabase(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByDatabase", reflect.TypeOf((*MockSQLDatabaseVulnerabilityAssessmentScansClient)(nil).ListByDatabase), arg0, arg1, arg2, arg3)
}

// MockTransparentDataEncryptionsClient is a mock of TransparentDataEncryptionsClient interface.
type MockTransparentDataEncryptionsClient struct {
	ctrl     *gomock.Controller
	recorder *MockTransparentDataEncryptionsClientMockRecorder
}

// MockTransparentDataEncryptionsClientMockRecorder is the mock recorder for MockTransparentDataEncryptionsClient.
type MockTransparentDataEncryptionsClientMockRecorder struct {
	mock *MockTransparentDataEncryptionsClient
}

// NewMockTransparentDataEncryptionsClient creates a new mock instance.
func NewMockTransparentDataEncryptionsClient(ctrl *gomock.Controller) *MockTransparentDataEncryptionsClient {
	mock := &MockTransparentDataEncryptionsClient{ctrl: ctrl}
	mock.recorder = &MockTransparentDataEncryptionsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransparentDataEncryptionsClient) EXPECT() *MockTransparentDataEncryptionsClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockTransparentDataEncryptionsClient) Get(arg0 context.Context, arg1, arg2, arg3 string) (sql.TransparentDataEncryption, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(sql.TransparentDataEncryption)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockTransparentDataEncryptionsClientMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTransparentDataEncryptionsClient)(nil).Get), arg0, arg1, arg2, arg3)
}

// MockBackupLongTermRetentionPoliciesClient is a mock of BackupLongTermRetentionPoliciesClient interface.
type MockBackupLongTermRetentionPoliciesClient struct {
	ctrl     *gomock.Controller
	recorder *MockBackupLongTermRetentionPoliciesClientMockRecorder
}

// MockBackupLongTermRetentionPoliciesClientMockRecorder is the mock recorder for MockBackupLongTermRetentionPoliciesClient.
type MockBackupLongTermRetentionPoliciesClientMockRecorder struct {
	mock *MockBackupLongTermRetentionPoliciesClient
}

// NewMockBackupLongTermRetentionPoliciesClient creates a new mock instance.
func NewMockBackupLongTermRetentionPoliciesClient(ctrl *gomock.Controller) *MockBackupLongTermRetentionPoliciesClient {
	mock := &MockBackupLongTermRetentionPoliciesClient{ctrl: ctrl}
	mock.recorder = &MockBackupLongTermRetentionPoliciesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackupLongTermRetentionPoliciesClient) EXPECT() *MockBackupLongTermRetentionPoliciesClientMockRecorder {
	return m.recorder
}

// ListByDatabase mocks base method.
func (m *MockBackupLongTermRetentionPoliciesClient) ListByDatabase(arg0 context.Context, arg1, arg2, arg3 string) (sql.BackupLongTermRetentionPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByDatabase", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(sql.BackupLongTermRetentionPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByDatabase indicates an expected call of ListByDatabase.
func (mr *MockBackupLongTermRetentionPoliciesClientMockRecorder) ListByDatabase(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByDatabase", reflect.TypeOf((*MockBackupLongTermRetentionPoliciesClient)(nil).ListByDatabase), arg0, arg1, arg2, arg3)
}