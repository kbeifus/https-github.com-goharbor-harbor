// Code generated by mockery v2.35.4. DO NOT EDIT.

package dao

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/goharbor/harbor/src/pkg/replication/model"

	q "github.com/goharbor/harbor/src/lib/q"
)

// DAO is an autogenerated mock type for the DAO type
type DAO struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx, query
func (_m *DAO) Count(ctx context.Context, query *q.Query) (int64, error) {
	ret := _m.Called(ctx, query)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) (int64, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) int64); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, policy
func (_m *DAO) Create(ctx context.Context, policy *model.Policy) (int64, error) {
	ret := _m.Called(ctx, policy)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Policy) (int64, error)); ok {
		return rf(ctx, policy)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.Policy) int64); ok {
		r0 = rf(ctx, policy)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.Policy) error); ok {
		r1 = rf(ctx, policy)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *DAO) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id
func (_m *DAO) Get(ctx context.Context, id int64) (*model.Policy, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.Policy
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*model.Policy, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *model.Policy); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Policy)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, query
func (_m *DAO) List(ctx context.Context, query *q.Query) ([]*model.Policy, error) {
	ret := _m.Called(ctx, query)

	var r0 []*model.Policy
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) ([]*model.Policy, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) []*model.Policy); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Policy)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, policy, props
func (_m *DAO) Update(ctx context.Context, policy *model.Policy, props ...string) error {
	_va := make([]interface{}, len(props))
	for _i := range props {
		_va[_i] = props[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, policy)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Policy, ...string) error); ok {
		r0 = rf(ctx, policy, props...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewDAO creates a new instance of DAO. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDAO(t interface {
	mock.TestingT
	Cleanup(func())
}) *DAO {
	mock := &DAO{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
