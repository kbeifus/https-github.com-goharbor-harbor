// Code generated by mockery v2.14.0. DO NOT EDIT.

package dao

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/goharbor/harbor/src/pkg/label/model"

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

// Create provides a mock function with given fields: ctx, label
func (_m *DAO) Create(ctx context.Context, label *model.Label) (int64, error) {
	ret := _m.Called(ctx, label)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *model.Label) int64); ok {
		r0 = rf(ctx, label)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.Label) error); ok {
		r1 = rf(ctx, label)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateReference provides a mock function with given fields: ctx, reference
func (_m *DAO) CreateReference(ctx context.Context, reference *model.Reference) (int64, error) {
	ret := _m.Called(ctx, reference)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *model.Reference) int64); ok {
		r0 = rf(ctx, reference)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.Reference) error); ok {
		r1 = rf(ctx, reference)
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

// DeleteReference provides a mock function with given fields: ctx, id
func (_m *DAO) DeleteReference(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteReferences provides a mock function with given fields: ctx, query
func (_m *DAO) DeleteReferences(ctx context.Context, query *q.Query) (int64, error) {
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

// Get provides a mock function with given fields: ctx, id
func (_m *DAO) Get(ctx context.Context, id int64) (*model.Label, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.Label
	if rf, ok := ret.Get(0).(func(context.Context, int64) *model.Label); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Label)
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

// List provides a mock function with given fields: ctx, query
func (_m *DAO) List(ctx context.Context, query *q.Query) ([]*model.Label, error) {
	ret := _m.Called(ctx, query)

	var r0 []*model.Label
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) []*model.Label); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Label)
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

// ListByArtifact provides a mock function with given fields: ctx, artifactID
func (_m *DAO) ListByArtifact(ctx context.Context, artifactID int64) ([]*model.Label, error) {
	ret := _m.Called(ctx, artifactID)

	var r0 []*model.Label
	if rf, ok := ret.Get(0).(func(context.Context, int64) []*model.Label); ok {
		r0 = rf(ctx, artifactID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Label)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, artifactID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, label
func (_m *DAO) Update(ctx context.Context, label *model.Label) error {
	ret := _m.Called(ctx, label)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Label) error); ok {
		r0 = rf(ctx, label)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewDAO interface {
	mock.TestingT
	Cleanup(func())
}

// NewDAO creates a new instance of DAO. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDAO(t mockConstructorTestingTNewDAO) *DAO {
	mock := &DAO{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
