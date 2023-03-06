// Code generated by mockery v2.14.0. DO NOT EDIT.

package task

import (
	context "context"

	q "github.com/goharbor/harbor/src/lib/q"
	mock "github.com/stretchr/testify/mock"
)

// mockTaskManager is an autogenerated mock type for the Manager type
type mockTaskManager struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx, query
func (_m *mockTaskManager) Count(ctx context.Context, query *q.Query) (int64, error) {
	ret := _m.Called(ctx, query)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) int64); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, executionID, job, extraAttrs
func (_m *mockTaskManager) Create(ctx context.Context, executionID int64, job *Job, extraAttrs ...map[string]interface{}) (int64, error) {
	_va := make([]interface{}, len(extraAttrs))
	for _i := range extraAttrs {
		_va[_i] = extraAttrs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, executionID, job)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, int64, *Job, ...map[string]interface{}) int64); ok {
		r0 = rf(ctx, executionID, job, extraAttrs...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, *Job, ...map[string]interface{}) error); ok {
		r1 = rf(ctx, executionID, job, extraAttrs...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecutionIDsByVendorAndStatus provides a mock function with given fields: ctx, vendorType, status
func (_m *mockTaskManager) ExecutionIDsByVendorAndStatus(ctx context.Context, vendorType string, status string) ([]int64, error) {
	ret := _m.Called(ctx, vendorType, status)

	var r0 []int64
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []int64); ok {
		r0 = rf(ctx, vendorType, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, vendorType, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, id
func (_m *mockTaskManager) Get(ctx context.Context, id int64) (*Task, error) {
	ret := _m.Called(ctx, id)

	var r0 *Task
	if rf, ok := ret.Get(0).(func(context.Context, int64) *Task); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLog provides a mock function with given fields: ctx, id
func (_m *mockTaskManager) GetLog(ctx context.Context, id int64) ([]byte, error) {
	ret := _m.Called(ctx, id)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, int64) []byte); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLogByJobID provides a mock function with given fields: ctx, jobID
func (_m *mockTaskManager) GetLogByJobID(ctx context.Context, jobID string) ([]byte, error) {
	ret := _m.Called(ctx, jobID)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, string) []byte); ok {
		r0 = rf(ctx, jobID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, jobID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, query
func (_m *mockTaskManager) List(ctx context.Context, query *q.Query) ([]*Task, error) {
	ret := _m.Called(ctx, query)

	var r0 []*Task
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) []*Task); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Stop provides a mock function with given fields: ctx, id
func (_m *mockTaskManager) Stop(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, task, props
func (_m *mockTaskManager) Update(ctx context.Context, task *Task, props ...string) error {
	_va := make([]interface{}, len(props))
	for _i := range props {
		_va[_i] = props[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, task)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *Task, ...string) error); ok {
		r0 = rf(ctx, task, props...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateExtraAttrs provides a mock function with given fields: ctx, id, extraAttrs
func (_m *mockTaskManager) UpdateExtraAttrs(ctx context.Context, id int64, extraAttrs map[string]interface{}) error {
	ret := _m.Called(ctx, id, extraAttrs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, map[string]interface{}) error); ok {
		r0 = rf(ctx, id, extraAttrs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateStatusInBatch provides a mock function with given fields: ctx, jobIDs, status, batchSize
func (_m *mockTaskManager) UpdateStatusInBatch(ctx context.Context, jobIDs []string, status string, batchSize int) error {
	ret := _m.Called(ctx, jobIDs, status, batchSize)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []string, string, int) error); ok {
		r0 = rf(ctx, jobIDs, status, batchSize)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTnewMockTaskManager interface {
	mock.TestingT
	Cleanup(func())
}

// newMockTaskManager creates a new instance of mockTaskManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockTaskManager(t mockConstructorTestingTnewMockTaskManager) *mockTaskManager {
	mock := &mockTaskManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
