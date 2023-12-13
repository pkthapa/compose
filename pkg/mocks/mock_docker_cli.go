// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/docker/cli/cli/command (interfaces: Cli)
//
// Generated by this command:
//
//	mockgen -destination pkg/mocks/mock_docker_cli.go -package mocks github.com/docker/cli/cli/command Cli
//
// Package mocks is a generated GoMock package.
package mocks

import (
	io "io"
	reflect "reflect"

	command "github.com/docker/cli/cli/command"
	configfile "github.com/docker/cli/cli/config/configfile"
	docker "github.com/docker/cli/cli/context/docker"
	store "github.com/docker/cli/cli/context/store"
	store0 "github.com/docker/cli/cli/manifest/store"
	client "github.com/docker/cli/cli/registry/client"
	streams "github.com/docker/cli/cli/streams"
	trust "github.com/docker/cli/cli/trust"
	client0 "github.com/docker/docker/client"
	client1 "github.com/theupdateframework/notary/client"
	gomock "go.uber.org/mock/gomock"
)

// MockCli is a mock of Cli interface.
type MockCli struct {
	ctrl     *gomock.Controller
	recorder *MockCliMockRecorder
}

// MockCliMockRecorder is the mock recorder for MockCli.
type MockCliMockRecorder struct {
	mock *MockCli
}

// NewMockCli creates a new mock instance.
func NewMockCli(ctrl *gomock.Controller) *MockCli {
	mock := &MockCli{ctrl: ctrl}
	mock.recorder = &MockCliMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCli) EXPECT() *MockCliMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockCli) Apply(arg0 ...command.DockerCliOption) error {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Apply", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Apply indicates an expected call of Apply.
func (mr *MockCliMockRecorder) Apply(arg0 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockCli)(nil).Apply), arg0...)
}

// BuildKitEnabled mocks base method.
func (m *MockCli) BuildKitEnabled() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildKitEnabled")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildKitEnabled indicates an expected call of BuildKitEnabled.
func (mr *MockCliMockRecorder) BuildKitEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildKitEnabled", reflect.TypeOf((*MockCli)(nil).BuildKitEnabled))
}

// Client mocks base method.
func (m *MockCli) Client() client0.APIClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(client0.APIClient)
	return ret0
}

// Client indicates an expected call of Client.
func (mr *MockCliMockRecorder) Client() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockCli)(nil).Client))
}

// ConfigFile mocks base method.
func (m *MockCli) ConfigFile() *configfile.ConfigFile {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigFile")
	ret0, _ := ret[0].(*configfile.ConfigFile)
	return ret0
}

// ConfigFile indicates an expected call of ConfigFile.
func (mr *MockCliMockRecorder) ConfigFile() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigFile", reflect.TypeOf((*MockCli)(nil).ConfigFile))
}

// ContentTrustEnabled mocks base method.
func (m *MockCli) ContentTrustEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContentTrustEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// ContentTrustEnabled indicates an expected call of ContentTrustEnabled.
func (mr *MockCliMockRecorder) ContentTrustEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContentTrustEnabled", reflect.TypeOf((*MockCli)(nil).ContentTrustEnabled))
}

// ContextStore mocks base method.
func (m *MockCli) ContextStore() store.Store {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContextStore")
	ret0, _ := ret[0].(store.Store)
	return ret0
}

// ContextStore indicates an expected call of ContextStore.
func (mr *MockCliMockRecorder) ContextStore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContextStore", reflect.TypeOf((*MockCli)(nil).ContextStore))
}

// CurrentContext mocks base method.
func (m *MockCli) CurrentContext() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentContext")
	ret0, _ := ret[0].(string)
	return ret0
}

// CurrentContext indicates an expected call of CurrentContext.
func (mr *MockCliMockRecorder) CurrentContext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentContext", reflect.TypeOf((*MockCli)(nil).CurrentContext))
}

// CurrentVersion mocks base method.
func (m *MockCli) CurrentVersion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

// CurrentVersion indicates an expected call of CurrentVersion.
func (mr *MockCliMockRecorder) CurrentVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentVersion", reflect.TypeOf((*MockCli)(nil).CurrentVersion))
}

// DefaultVersion mocks base method.
func (m *MockCli) DefaultVersion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

// DefaultVersion indicates an expected call of DefaultVersion.
func (mr *MockCliMockRecorder) DefaultVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultVersion", reflect.TypeOf((*MockCli)(nil).DefaultVersion))
}

// DockerEndpoint mocks base method.
func (m *MockCli) DockerEndpoint() docker.Endpoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DockerEndpoint")
	ret0, _ := ret[0].(docker.Endpoint)
	return ret0
}

// DockerEndpoint indicates an expected call of DockerEndpoint.
func (mr *MockCliMockRecorder) DockerEndpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DockerEndpoint", reflect.TypeOf((*MockCli)(nil).DockerEndpoint))
}

// Err mocks base method.
func (m *MockCli) Err() io.Writer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(io.Writer)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockCliMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockCli)(nil).Err))
}

// In mocks base method.
func (m *MockCli) In() *streams.In {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "In")
	ret0, _ := ret[0].(*streams.In)
	return ret0
}

// In indicates an expected call of In.
func (mr *MockCliMockRecorder) In() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "In", reflect.TypeOf((*MockCli)(nil).In))
}

// ManifestStore mocks base method.
func (m *MockCli) ManifestStore() store0.Store {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ManifestStore")
	ret0, _ := ret[0].(store0.Store)
	return ret0
}

// ManifestStore indicates an expected call of ManifestStore.
func (mr *MockCliMockRecorder) ManifestStore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ManifestStore", reflect.TypeOf((*MockCli)(nil).ManifestStore))
}

// NotaryClient mocks base method.
func (m *MockCli) NotaryClient(arg0 trust.ImageRefAndAuth, arg1 []string) (client1.Repository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotaryClient", arg0, arg1)
	ret0, _ := ret[0].(client1.Repository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NotaryClient indicates an expected call of NotaryClient.
func (mr *MockCliMockRecorder) NotaryClient(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotaryClient", reflect.TypeOf((*MockCli)(nil).NotaryClient), arg0, arg1)
}

// Out mocks base method.
func (m *MockCli) Out() *streams.Out {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Out")
	ret0, _ := ret[0].(*streams.Out)
	return ret0
}

// Out indicates an expected call of Out.
func (mr *MockCliMockRecorder) Out() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Out", reflect.TypeOf((*MockCli)(nil).Out))
}

// RegistryClient mocks base method.
func (m *MockCli) RegistryClient(arg0 bool) client.RegistryClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegistryClient", arg0)
	ret0, _ := ret[0].(client.RegistryClient)
	return ret0
}

// RegistryClient indicates an expected call of RegistryClient.
func (mr *MockCliMockRecorder) RegistryClient(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegistryClient", reflect.TypeOf((*MockCli)(nil).RegistryClient), arg0)
}

// ServerInfo mocks base method.
func (m *MockCli) ServerInfo() command.ServerInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerInfo")
	ret0, _ := ret[0].(command.ServerInfo)
	return ret0
}

// ServerInfo indicates an expected call of ServerInfo.
func (mr *MockCliMockRecorder) ServerInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerInfo", reflect.TypeOf((*MockCli)(nil).ServerInfo))
}

// SetIn mocks base method.
func (m *MockCli) SetIn(arg0 *streams.In) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetIn", arg0)
}

// SetIn indicates an expected call of SetIn.
func (mr *MockCliMockRecorder) SetIn(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetIn", reflect.TypeOf((*MockCli)(nil).SetIn), arg0)
}
