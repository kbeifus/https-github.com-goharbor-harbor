// Code generated by mockery v2.43.2. DO NOT EDIT.

package artifact

import (
	context "context"

	artifact "github.com/goharbor/harbor/src/pkg/artifact"

	mock "github.com/stretchr/testify/mock"

	q "github.com/goharbor/harbor/src/lib/q"

	time "time"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx, query
func (_m *Manager) Count(ctx context.Context, query *q.Query) (int64, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for Count")
	}

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

// Create provides a mock function with given fields: ctx, _a1
func (_m *Manager) Create(ctx context.Context, _a1 *artifact.Artifact) (int64, error) {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *artifact.Artifact) (int64, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *artifact.Artifact) int64); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *artifact.Artifact) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Manager) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteReference provides a mock function with given fields: ctx, id
func (_m *Manager) DeleteReference(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteReference")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id
func (_m *Manager) Get(ctx context.Context, id int64) (*artifact.Artifact, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *artifact.Artifact
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*artifact.Artifact, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *artifact.Artifact); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artifact.Artifact)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByDigest provides a mock function with given fields: ctx, repository, digest
func (_m *Manager) GetByDigest(ctx context.Context, repository string, digest string) (*artifact.Artifact, error) {
	ret := _m.Called(ctx, repository, digest)

	if len(ret) == 0 {
		panic("no return value specified for GetByDigest")
	}

	var r0 *artifact.Artifact
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*artifact.Artifact, error)); ok {
		return rf(ctx, repository, digest)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *artifact.Artifact); ok {
		r0 = rf(ctx, repository, digest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artifact.Artifact)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, repository, digest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, query
func (_m *Manager) List(ctx context.Context, query *q.Query) ([]*artifact.Artifact, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []*artifact.Artifact
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) ([]*artifact.Artifact, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) []*artifact.Artifact); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*artifact.Artifact)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListReferences provides a mock function with given fields: ctx, query
func (_m *Manager) ListReferences(ctx context.Context, query *q.Query) ([]*artifact.Reference, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for ListReferences")
	}

	var r0 []*artifact.Reference
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) ([]*artifact.Reference, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) []*artifact.Reference); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*artifact.Reference)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, _a1, props
func (_m *Manager) Update(ctx context.Context, _a1 *artifact.Artifact, props ...string) error {
	_va := make([]interface{}, len(props))
	for _i := range props {
		_va[_i] = props[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *artifact.Artifact, ...string) error); ok {
		r0 = rf(ctx, _a1, props...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePullTime provides a mock function with given fields: ctx, id, pullTime
func (_m *Manager) UpdatePullTime(ctx context.Context, id int64, pullTime time.Time) error {
	ret := _m.Called(ctx, id, pullTime)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePullTime")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, time.Time) error); ok {
		r0 = rf(ctx, id, pullTime)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewManager creates a new instance of Manager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *Manager {
	mock := &Manager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
