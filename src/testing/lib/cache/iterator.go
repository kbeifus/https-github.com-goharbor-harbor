// Code generated by mockery v2.35.4. DO NOT EDIT.

package cache

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Iterator is an autogenerated mock type for the Iterator type
type Iterator struct {
	mock.Mock
}

// Next provides a mock function with given fields: ctx
func (_m *Iterator) Next(ctx context.Context) bool {
	ret := _m.Called(ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Val provides a mock function with given fields:
func (_m *Iterator) Val() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewIterator creates a new instance of Iterator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIterator(t interface {
	mock.TestingT
	Cleanup(func())
}) *Iterator {
	mock := &Iterator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
