// Code generated by mockery v2.22.1. DO NOT EDIT.

package securityhub

import (
	context "context"

	model "github.com/goharbor/harbor/src/pkg/securityhub/model"
	mock "github.com/stretchr/testify/mock"

	q "github.com/goharbor/harbor/src/lib/q"

	scan "github.com/goharbor/harbor/src/pkg/scan/dao/scan"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// DangerousArtifacts provides a mock function with given fields: ctx, scannerUUID, projectID, query
func (_m *Manager) DangerousArtifacts(ctx context.Context, scannerUUID string, projectID int64, query *q.Query) ([]*model.DangerousArtifact, error) {
	ret := _m.Called(ctx, scannerUUID, projectID, query)

	var r0 []*model.DangerousArtifact
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, *q.Query) ([]*model.DangerousArtifact, error)); ok {
		return rf(ctx, scannerUUID, projectID, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, *q.Query) []*model.DangerousArtifact); ok {
		r0 = rf(ctx, scannerUUID, projectID, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.DangerousArtifact)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int64, *q.Query) error); ok {
		r1 = rf(ctx, scannerUUID, projectID, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DangerousCVEs provides a mock function with given fields: ctx, scannerUUID, projectID, query
func (_m *Manager) DangerousCVEs(ctx context.Context, scannerUUID string, projectID int64, query *q.Query) ([]*scan.VulnerabilityRecord, error) {
	ret := _m.Called(ctx, scannerUUID, projectID, query)

	var r0 []*scan.VulnerabilityRecord
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, *q.Query) ([]*scan.VulnerabilityRecord, error)); ok {
		return rf(ctx, scannerUUID, projectID, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, *q.Query) []*scan.VulnerabilityRecord); ok {
		r0 = rf(ctx, scannerUUID, projectID, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*scan.VulnerabilityRecord)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int64, *q.Query) error); ok {
		r1 = rf(ctx, scannerUUID, projectID, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListVuls provides a mock function with given fields: ctx, scannerUUID, projectID, query
func (_m *Manager) ListVuls(ctx context.Context, scannerUUID string, projectID int64, query *q.Query) ([]*model.VulnerabilityItem, error) {
	ret := _m.Called(ctx, scannerUUID, projectID, query)

	var r0 []*model.VulnerabilityItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, *q.Query) ([]*model.VulnerabilityItem, error)); ok {
		return rf(ctx, scannerUUID, projectID, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, *q.Query) []*model.VulnerabilityItem); ok {
		r0 = rf(ctx, scannerUUID, projectID, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.VulnerabilityItem)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int64, *q.Query) error); ok {
		r1 = rf(ctx, scannerUUID, projectID, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ScannedArtifactsCount provides a mock function with given fields: ctx, scannerUUID, projectID, query
func (_m *Manager) ScannedArtifactsCount(ctx context.Context, scannerUUID string, projectID int64, query *q.Query) (int64, error) {
	ret := _m.Called(ctx, scannerUUID, projectID, query)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, *q.Query) (int64, error)); ok {
		return rf(ctx, scannerUUID, projectID, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, *q.Query) int64); ok {
		r0 = rf(ctx, scannerUUID, projectID, query)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int64, *q.Query) error); ok {
		r1 = rf(ctx, scannerUUID, projectID, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Summary provides a mock function with given fields: ctx, scannerUUID, projectID, query
func (_m *Manager) Summary(ctx context.Context, scannerUUID string, projectID int64, query *q.Query) (*model.Summary, error) {
	ret := _m.Called(ctx, scannerUUID, projectID, query)

	var r0 *model.Summary
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, *q.Query) (*model.Summary, error)); ok {
		return rf(ctx, scannerUUID, projectID, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, *q.Query) *model.Summary); ok {
		r0 = rf(ctx, scannerUUID, projectID, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Summary)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int64, *q.Query) error); ok {
		r1 = rf(ctx, scannerUUID, projectID, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TotalVuls provides a mock function with given fields: ctx, scannerUUID, projectID, tuneCount, query
func (_m *Manager) TotalVuls(ctx context.Context, scannerUUID string, projectID int64, tuneCount bool, query *q.Query) (int64, error) {
	ret := _m.Called(ctx, scannerUUID, projectID, tuneCount, query)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, bool, *q.Query) (int64, error)); ok {
		return rf(ctx, scannerUUID, projectID, tuneCount, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, bool, *q.Query) int64); ok {
		r0 = rf(ctx, scannerUUID, projectID, tuneCount, query)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int64, bool, *q.Query) error); ok {
		r1 = rf(ctx, scannerUUID, projectID, tuneCount, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewManager creates a new instance of Manager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewManager(t mockConstructorTestingTNewManager) *Manager {
	mock := &Manager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
