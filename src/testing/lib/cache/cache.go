// Code generated by mockery v2.14.0. DO NOT EDIT.

package cache

import (
	context "context"

	cache "github.com/goharbor/harbor/src/lib/cache"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Cache is an autogenerated mock type for the Cache type
type Cache struct {
	mock.Mock
}

// Contains provides a mock function with given fields: ctx, key
func (_m *Cache) Contains(ctx context.Context, key string) bool {
	ret := _m.Called(ctx, key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, key
func (_m *Cache) Delete(ctx context.Context, key string) error {
	ret := _m.Called(ctx, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: ctx, key, value
func (_m *Cache) Fetch(ctx context.Context, key string, value interface{}) error {
	ret := _m.Called(ctx, key, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) error); ok {
		r0 = rf(ctx, key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Ping provides a mock function with given fields: ctx
func (_m *Cache) Ping(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: ctx, key, value, expiration
func (_m *Cache) Save(ctx context.Context, key string, value interface{}, expiration ...time.Duration) error {
	_va := make([]interface{}, len(expiration))
	for _i := range expiration {
		_va[_i] = expiration[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, key, value)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, ...time.Duration) error); ok {
		r0 = rf(ctx, key, value, expiration...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Scan provides a mock function with given fields: ctx, match
func (_m *Cache) Scan(ctx context.Context, match string) (cache.Iterator, error) {
	ret := _m.Called(ctx, match)

	var r0 cache.Iterator
	if rf, ok := ret.Get(0).(func(context.Context, string) cache.Iterator); ok {
		r0 = rf(ctx, match)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cache.Iterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, match)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCache interface {
	mock.TestingT
	Cleanup(func())
}

// NewCache creates a new instance of Cache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCache(t mockConstructorTestingTNewCache) *Cache {
	mock := &Cache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
