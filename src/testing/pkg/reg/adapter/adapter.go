// Code generated by mockery v2.12.3. DO NOT EDIT.

package adapter

import (
	mock "github.com/stretchr/testify/mock"

	model "github.com/goharbor/harbor/src/pkg/reg/model"
)

// Adapter is an autogenerated mock type for the Adapter type
type Adapter struct {
	mock.Mock
}

// HealthCheck provides a mock function with given fields:
func (_m *Adapter) HealthCheck() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Info provides a mock function with given fields:
func (_m *Adapter) Info() (*model.RegistryInfo, error) {
	ret := _m.Called()

	var r0 *model.RegistryInfo
	if rf, ok := ret.Get(0).(func() *model.RegistryInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.RegistryInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PrepareForPush provides a mock function with given fields: _a0
func (_m *Adapter) PrepareForPush(_a0 []*model.Resource) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*model.Resource) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NewAdapterT interface {
	mock.TestingT
	Cleanup(func())
}

// NewAdapter creates a new instance of Adapter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAdapter(t NewAdapterT) *Adapter {
	mock := &Adapter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
