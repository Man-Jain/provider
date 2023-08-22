// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1beta1 "k8s.io/api/authorization/v1beta1"
)

// SelfSubjectAccessReviewInterface is an autogenerated mock type for the SelfSubjectAccessReviewInterface type
type SelfSubjectAccessReviewInterface struct {
	mock.Mock
}

type SelfSubjectAccessReviewInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *SelfSubjectAccessReviewInterface) EXPECT() *SelfSubjectAccessReviewInterface_Expecter {
	return &SelfSubjectAccessReviewInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, selfSubjectAccessReview, opts
func (_m *SelfSubjectAccessReviewInterface) Create(ctx context.Context, selfSubjectAccessReview *v1beta1.SelfSubjectAccessReview, opts v1.CreateOptions) (*v1beta1.SelfSubjectAccessReview, error) {
	ret := _m.Called(ctx, selfSubjectAccessReview, opts)

	var r0 *v1beta1.SelfSubjectAccessReview
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.SelfSubjectAccessReview, v1.CreateOptions) (*v1beta1.SelfSubjectAccessReview, error)); ok {
		return rf(ctx, selfSubjectAccessReview, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.SelfSubjectAccessReview, v1.CreateOptions) *v1beta1.SelfSubjectAccessReview); ok {
		r0 = rf(ctx, selfSubjectAccessReview, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.SelfSubjectAccessReview)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1beta1.SelfSubjectAccessReview, v1.CreateOptions) error); ok {
		r1 = rf(ctx, selfSubjectAccessReview, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelfSubjectAccessReviewInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type SelfSubjectAccessReviewInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - selfSubjectAccessReview *v1beta1.SelfSubjectAccessReview
//   - opts v1.CreateOptions
func (_e *SelfSubjectAccessReviewInterface_Expecter) Create(ctx interface{}, selfSubjectAccessReview interface{}, opts interface{}) *SelfSubjectAccessReviewInterface_Create_Call {
	return &SelfSubjectAccessReviewInterface_Create_Call{Call: _e.mock.On("Create", ctx, selfSubjectAccessReview, opts)}
}

func (_c *SelfSubjectAccessReviewInterface_Create_Call) Run(run func(ctx context.Context, selfSubjectAccessReview *v1beta1.SelfSubjectAccessReview, opts v1.CreateOptions)) *SelfSubjectAccessReviewInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1beta1.SelfSubjectAccessReview), args[2].(v1.CreateOptions))
	})
	return _c
}

func (_c *SelfSubjectAccessReviewInterface_Create_Call) Return(_a0 *v1beta1.SelfSubjectAccessReview, _a1 error) *SelfSubjectAccessReviewInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SelfSubjectAccessReviewInterface_Create_Call) RunAndReturn(run func(context.Context, *v1beta1.SelfSubjectAccessReview, v1.CreateOptions) (*v1beta1.SelfSubjectAccessReview, error)) *SelfSubjectAccessReviewInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewSelfSubjectAccessReviewInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewSelfSubjectAccessReviewInterface creates a new instance of SelfSubjectAccessReviewInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSelfSubjectAccessReviewInterface(t mockConstructorTestingTNewSelfSubjectAccessReviewInterface) *SelfSubjectAccessReviewInterface {
	mock := &SelfSubjectAccessReviewInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
