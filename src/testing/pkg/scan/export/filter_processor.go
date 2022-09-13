// Code generated by mockery v2.14.0. DO NOT EDIT.

package export

import (
	context "context"

	artifact "github.com/goharbor/harbor/src/controller/artifact"

	mock "github.com/stretchr/testify/mock"
)

// FilterProcessor is an autogenerated mock type for the FilterProcessor type
type FilterProcessor struct {
	mock.Mock
}

// ProcessLabelFilter provides a mock function with given fields: ctx, labelIDs, arts
func (_m *FilterProcessor) ProcessLabelFilter(ctx context.Context, labelIDs []int64, arts []*artifact.Artifact) ([]*artifact.Artifact, error) {
	ret := _m.Called(ctx, labelIDs, arts)

	var r0 []*artifact.Artifact
	if rf, ok := ret.Get(0).(func(context.Context, []int64, []*artifact.Artifact) []*artifact.Artifact); ok {
		r0 = rf(ctx, labelIDs, arts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*artifact.Artifact)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []int64, []*artifact.Artifact) error); ok {
		r1 = rf(ctx, labelIDs, arts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessProjectFilter provides a mock function with given fields: ctx, userName, projectsToFilter
func (_m *FilterProcessor) ProcessProjectFilter(ctx context.Context, userName string, projectsToFilter []int64) ([]int64, error) {
	ret := _m.Called(ctx, userName, projectsToFilter)

	var r0 []int64
	if rf, ok := ret.Get(0).(func(context.Context, string, []int64) []int64); ok {
		r0 = rf(ctx, userName, projectsToFilter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, []int64) error); ok {
		r1 = rf(ctx, userName, projectsToFilter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessRepositoryFilter provides a mock function with given fields: ctx, filter, projectIds
func (_m *FilterProcessor) ProcessRepositoryFilter(ctx context.Context, filter string, projectIds []int64) ([]int64, error) {
	ret := _m.Called(ctx, filter, projectIds)

	var r0 []int64
	if rf, ok := ret.Get(0).(func(context.Context, string, []int64) []int64); ok {
		r0 = rf(ctx, filter, projectIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, []int64) error); ok {
		r1 = rf(ctx, filter, projectIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessTagFilter provides a mock function with given fields: ctx, filter, repositoryIds
func (_m *FilterProcessor) ProcessTagFilter(ctx context.Context, filter string, repositoryIds []int64) ([]*artifact.Artifact, error) {
	ret := _m.Called(ctx, filter, repositoryIds)

	var r0 []*artifact.Artifact
	if rf, ok := ret.Get(0).(func(context.Context, string, []int64) []*artifact.Artifact); ok {
		r0 = rf(ctx, filter, repositoryIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*artifact.Artifact)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, []int64) error); ok {
		r1 = rf(ctx, filter, repositoryIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFilterProcessor interface {
	mock.TestingT
	Cleanup(func())
}

// NewFilterProcessor creates a new instance of FilterProcessor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFilterProcessor(t mockConstructorTestingTNewFilterProcessor) *FilterProcessor {
	mock := &FilterProcessor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
