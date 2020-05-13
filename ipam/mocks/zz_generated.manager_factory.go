// /*
// Copyright The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */
//
//

// Code generated by MockGen. DO NOT EDIT.
// Source: ./ipam/manager_factory.go

// Package ipam_mocks is a generated GoMock package.
package ipam_mocks

import (
	logr "github.com/go-logr/logr"
	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/metal3-io/ip-address-manager/api/v1alpha1"
	ipam "github.com/metal3-io/ip-address-manager/ipam"
	reflect "reflect"
)

// MockManagerFactoryInterface is a mock of ManagerFactoryInterface interface
type MockManagerFactoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockManagerFactoryInterfaceMockRecorder
}

// MockManagerFactoryInterfaceMockRecorder is the mock recorder for MockManagerFactoryInterface
type MockManagerFactoryInterfaceMockRecorder struct {
	mock *MockManagerFactoryInterface
}

// NewMockManagerFactoryInterface creates a new mock instance
func NewMockManagerFactoryInterface(ctrl *gomock.Controller) *MockManagerFactoryInterface {
	mock := &MockManagerFactoryInterface{ctrl: ctrl}
	mock.recorder = &MockManagerFactoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManagerFactoryInterface) EXPECT() *MockManagerFactoryInterfaceMockRecorder {
	return m.recorder
}

// NewIPPoolManager mocks base method
func (m *MockManagerFactoryInterface) NewIPPoolManager(arg0 *v1alpha1.IPPool, arg1 logr.Logger) (ipam.IPPoolManagerInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewIPPoolManager", arg0, arg1)
	ret0, _ := ret[0].(ipam.IPPoolManagerInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewIPPoolManager indicates an expected call of NewIPPoolManager
func (mr *MockManagerFactoryInterfaceMockRecorder) NewIPPoolManager(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewIPPoolManager", reflect.TypeOf((*MockManagerFactoryInterface)(nil).NewIPPoolManager), arg0, arg1)
}
