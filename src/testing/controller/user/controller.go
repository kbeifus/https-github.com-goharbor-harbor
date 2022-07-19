// Code generated by mockery v2.12.3. DO NOT EDIT.

package user

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/goharbor/harbor/src/common/models"
	user "github.com/goharbor/harbor/src/controller/user"
	q "github.com/goharbor/harbor/src/lib/q"
	usermodels "github.com/goharbor/harbor/src/pkg/user/models"
)

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx, query
func (_m *Controller) Count(ctx context.Context, query *q.Query) (int64, error) {
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

// Create provides a mock function with given fields: ctx, u
func (_m *Controller) Create(ctx context.Context, u *models.User) (int, error) {
	ret := _m.Called(ctx, u)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, *models.User) int); ok {
		r0 = rf(ctx, u)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.User) error); ok {
		r1 = rf(ctx, u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Controller) Delete(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id, opt
func (_m *Controller) Get(ctx context.Context, id int, opt *user.Option) (*models.User, error) {
	ret := _m.Called(ctx, id, opt)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context, int, *user.Option) *models.User); ok {
		r0 = rf(ctx, id, opt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, *user.Option) error); ok {
		r1 = rf(ctx, id, opt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: ctx, username
func (_m *Controller) GetByName(ctx context.Context, username string) (*models.User, error) {
	ret := _m.Called(ctx, username)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.User); ok {
		r0 = rf(ctx, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBySubIss provides a mock function with given fields: ctx, sub, iss
func (_m *Controller) GetBySubIss(ctx context.Context, sub string, iss string) (*models.User, error) {
	ret := _m.Called(ctx, sub, iss)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *models.User); ok {
		r0 = rf(ctx, sub, iss)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, sub, iss)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, query, options
func (_m *Controller) List(ctx context.Context, query *q.Query, options ...usermodels.Option) ([]*models.User, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*models.User
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query, ...usermodels.Option) []*models.User); ok {
		r0 = rf(ctx, query, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *q.Query, ...usermodels.Option) error); ok {
		r1 = rf(ctx, query, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OnboardOIDCUser provides a mock function with given fields: ctx, u
func (_m *Controller) OnboardOIDCUser(ctx context.Context, u *models.User) error {
	ret := _m.Called(ctx, u)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.User) error); ok {
		r0 = rf(ctx, u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetCliSecret provides a mock function with given fields: ctx, id, secret
func (_m *Controller) SetCliSecret(ctx context.Context, id int, secret string) error {
	ret := _m.Called(ctx, id, secret)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string) error); ok {
		r0 = rf(ctx, id, secret)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetSysAdmin provides a mock function with given fields: ctx, id, adminFlag
func (_m *Controller) SetSysAdmin(ctx context.Context, id int, adminFlag bool) error {
	ret := _m.Called(ctx, id, adminFlag)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, bool) error); ok {
		r0 = rf(ctx, id, adminFlag)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateOIDCMeta provides a mock function with given fields: ctx, ou, cols
func (_m *Controller) UpdateOIDCMeta(ctx context.Context, ou *models.OIDCUser, cols ...string) error {
	_va := make([]interface{}, len(cols))
	for _i := range cols {
		_va[_i] = cols[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, ou)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.OIDCUser, ...string) error); ok {
		r0 = rf(ctx, ou, cols...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePassword provides a mock function with given fields: ctx, id, password
func (_m *Controller) UpdatePassword(ctx context.Context, id int, password string) error {
	ret := _m.Called(ctx, id, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string) error); ok {
		r0 = rf(ctx, id, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProfile provides a mock function with given fields: ctx, u, cols
func (_m *Controller) UpdateProfile(ctx context.Context, u *models.User, cols ...string) error {
	_va := make([]interface{}, len(cols))
	for _i := range cols {
		_va[_i] = cols[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, u)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.User, ...string) error); ok {
		r0 = rf(ctx, u, cols...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifyPassword provides a mock function with given fields: ctx, usernameOrEmail, password
func (_m *Controller) VerifyPassword(ctx context.Context, usernameOrEmail string, password string) (bool, error) {
	ret := _m.Called(ctx, usernameOrEmail, password)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, usernameOrEmail, password)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, usernameOrEmail, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewControllerT interface {
	mock.TestingT
	Cleanup(func())
}

// NewController creates a new instance of Controller. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewController(t NewControllerT) *Controller {
	mock := &Controller{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
