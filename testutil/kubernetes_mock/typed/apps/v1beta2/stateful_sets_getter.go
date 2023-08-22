// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1beta2 "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
)

// StatefulSetsGetter is an autogenerated mock type for the StatefulSetsGetter type
type StatefulSetsGetter struct {
	mock.Mock
}

type StatefulSetsGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *StatefulSetsGetter) EXPECT() *StatefulSetsGetter_Expecter {
	return &StatefulSetsGetter_Expecter{mock: &_m.Mock}
}

// StatefulSets provides a mock function with given fields: namespace
func (_m *StatefulSetsGetter) StatefulSets(namespace string) v1beta2.StatefulSetInterface {
	ret := _m.Called(namespace)

	var r0 v1beta2.StatefulSetInterface
	if rf, ok := ret.Get(0).(func(string) v1beta2.StatefulSetInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta2.StatefulSetInterface)
		}
	}

	return r0
}

// StatefulSetsGetter_StatefulSets_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StatefulSets'
type StatefulSetsGetter_StatefulSets_Call struct {
	*mock.Call
}

// StatefulSets is a helper method to define mock.On call
//   - namespace string
func (_e *StatefulSetsGetter_Expecter) StatefulSets(namespace interface{}) *StatefulSetsGetter_StatefulSets_Call {
	return &StatefulSetsGetter_StatefulSets_Call{Call: _e.mock.On("StatefulSets", namespace)}
}

func (_c *StatefulSetsGetter_StatefulSets_Call) Run(run func(namespace string)) *StatefulSetsGetter_StatefulSets_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *StatefulSetsGetter_StatefulSets_Call) Return(_a0 v1beta2.StatefulSetInterface) *StatefulSetsGetter_StatefulSets_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StatefulSetsGetter_StatefulSets_Call) RunAndReturn(run func(string) v1beta2.StatefulSetInterface) *StatefulSetsGetter_StatefulSets_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewStatefulSetsGetter interface {
	mock.TestingT
	Cleanup(func())
}

// NewStatefulSetsGetter creates a new instance of StatefulSetsGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStatefulSetsGetter(t mockConstructorTestingTNewStatefulSetsGetter) *StatefulSetsGetter {
	mock := &StatefulSetsGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
