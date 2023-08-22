// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

// ControllerRevisionsGetter is an autogenerated mock type for the ControllerRevisionsGetter type
type ControllerRevisionsGetter struct {
	mock.Mock
}

type ControllerRevisionsGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *ControllerRevisionsGetter) EXPECT() *ControllerRevisionsGetter_Expecter {
	return &ControllerRevisionsGetter_Expecter{mock: &_m.Mock}
}

// ControllerRevisions provides a mock function with given fields: namespace
func (_m *ControllerRevisionsGetter) ControllerRevisions(namespace string) v1.ControllerRevisionInterface {
	ret := _m.Called(namespace)

	var r0 v1.ControllerRevisionInterface
	if rf, ok := ret.Get(0).(func(string) v1.ControllerRevisionInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.ControllerRevisionInterface)
		}
	}

	return r0
}

// ControllerRevisionsGetter_ControllerRevisions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ControllerRevisions'
type ControllerRevisionsGetter_ControllerRevisions_Call struct {
	*mock.Call
}

// ControllerRevisions is a helper method to define mock.On call
//   - namespace string
func (_e *ControllerRevisionsGetter_Expecter) ControllerRevisions(namespace interface{}) *ControllerRevisionsGetter_ControllerRevisions_Call {
	return &ControllerRevisionsGetter_ControllerRevisions_Call{Call: _e.mock.On("ControllerRevisions", namespace)}
}

func (_c *ControllerRevisionsGetter_ControllerRevisions_Call) Run(run func(namespace string)) *ControllerRevisionsGetter_ControllerRevisions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *ControllerRevisionsGetter_ControllerRevisions_Call) Return(_a0 v1.ControllerRevisionInterface) *ControllerRevisionsGetter_ControllerRevisions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ControllerRevisionsGetter_ControllerRevisions_Call) RunAndReturn(run func(string) v1.ControllerRevisionInterface) *ControllerRevisionsGetter_ControllerRevisions_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewControllerRevisionsGetter interface {
	mock.TestingT
	Cleanup(func())
}

// NewControllerRevisionsGetter creates a new instance of ControllerRevisionsGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewControllerRevisionsGetter(t mockConstructorTestingTNewControllerRevisionsGetter) *ControllerRevisionsGetter {
	mock := &ControllerRevisionsGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
