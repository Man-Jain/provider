// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/authorization/v1beta1"
)

// LocalSubjectAccessReviewsGetter is an autogenerated mock type for the LocalSubjectAccessReviewsGetter type
type LocalSubjectAccessReviewsGetter struct {
	mock.Mock
}

type LocalSubjectAccessReviewsGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *LocalSubjectAccessReviewsGetter) EXPECT() *LocalSubjectAccessReviewsGetter_Expecter {
	return &LocalSubjectAccessReviewsGetter_Expecter{mock: &_m.Mock}
}

// LocalSubjectAccessReviews provides a mock function with given fields: namespace
func (_m *LocalSubjectAccessReviewsGetter) LocalSubjectAccessReviews(namespace string) v1beta1.LocalSubjectAccessReviewInterface {
	ret := _m.Called(namespace)

	var r0 v1beta1.LocalSubjectAccessReviewInterface
	if rf, ok := ret.Get(0).(func(string) v1beta1.LocalSubjectAccessReviewInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.LocalSubjectAccessReviewInterface)
		}
	}

	return r0
}

// LocalSubjectAccessReviewsGetter_LocalSubjectAccessReviews_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LocalSubjectAccessReviews'
type LocalSubjectAccessReviewsGetter_LocalSubjectAccessReviews_Call struct {
	*mock.Call
}

// LocalSubjectAccessReviews is a helper method to define mock.On call
//   - namespace string
func (_e *LocalSubjectAccessReviewsGetter_Expecter) LocalSubjectAccessReviews(namespace interface{}) *LocalSubjectAccessReviewsGetter_LocalSubjectAccessReviews_Call {
	return &LocalSubjectAccessReviewsGetter_LocalSubjectAccessReviews_Call{Call: _e.mock.On("LocalSubjectAccessReviews", namespace)}
}

func (_c *LocalSubjectAccessReviewsGetter_LocalSubjectAccessReviews_Call) Run(run func(namespace string)) *LocalSubjectAccessReviewsGetter_LocalSubjectAccessReviews_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *LocalSubjectAccessReviewsGetter_LocalSubjectAccessReviews_Call) Return(_a0 v1beta1.LocalSubjectAccessReviewInterface) *LocalSubjectAccessReviewsGetter_LocalSubjectAccessReviews_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *LocalSubjectAccessReviewsGetter_LocalSubjectAccessReviews_Call) RunAndReturn(run func(string) v1beta1.LocalSubjectAccessReviewInterface) *LocalSubjectAccessReviewsGetter_LocalSubjectAccessReviews_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewLocalSubjectAccessReviewsGetter interface {
	mock.TestingT
	Cleanup(func())
}

// NewLocalSubjectAccessReviewsGetter creates a new instance of LocalSubjectAccessReviewsGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLocalSubjectAccessReviewsGetter(t mockConstructorTestingTNewLocalSubjectAccessReviewsGetter) *LocalSubjectAccessReviewsGetter {
	mock := &LocalSubjectAccessReviewsGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
