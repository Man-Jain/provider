// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	v1 "k8s.io/api/policy/v1"
)

// EvictionExpansion is an autogenerated mock type for the EvictionExpansion type
type EvictionExpansion struct {
	mock.Mock
}

type EvictionExpansion_Expecter struct {
	mock *mock.Mock
}

func (_m *EvictionExpansion) EXPECT() *EvictionExpansion_Expecter {
	return &EvictionExpansion_Expecter{mock: &_m.Mock}
}

// Evict provides a mock function with given fields: ctx, eviction
func (_m *EvictionExpansion) Evict(ctx context.Context, eviction *v1.Eviction) error {
	ret := _m.Called(ctx, eviction)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.Eviction) error); ok {
		r0 = rf(ctx, eviction)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EvictionExpansion_Evict_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Evict'
type EvictionExpansion_Evict_Call struct {
	*mock.Call
}

// Evict is a helper method to define mock.On call
//   - ctx context.Context
//   - eviction *v1.Eviction
func (_e *EvictionExpansion_Expecter) Evict(ctx interface{}, eviction interface{}) *EvictionExpansion_Evict_Call {
	return &EvictionExpansion_Evict_Call{Call: _e.mock.On("Evict", ctx, eviction)}
}

func (_c *EvictionExpansion_Evict_Call) Run(run func(ctx context.Context, eviction *v1.Eviction)) *EvictionExpansion_Evict_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1.Eviction))
	})
	return _c
}

func (_c *EvictionExpansion_Evict_Call) Return(_a0 error) *EvictionExpansion_Evict_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EvictionExpansion_Evict_Call) RunAndReturn(run func(context.Context, *v1.Eviction) error) *EvictionExpansion_Evict_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewEvictionExpansion interface {
	mock.TestingT
	Cleanup(func())
}

// NewEvictionExpansion creates a new instance of EvictionExpansion. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEvictionExpansion(t mockConstructorTestingTNewEvictionExpansion) *EvictionExpansion {
	mock := &EvictionExpansion{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
