// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1beta1 "k8s.io/api/extensions/v1beta1"
)

// DeploymentExpansion is an autogenerated mock type for the DeploymentExpansion type
type DeploymentExpansion struct {
	mock.Mock
}

type DeploymentExpansion_Expecter struct {
	mock *mock.Mock
}

func (_m *DeploymentExpansion) EXPECT() *DeploymentExpansion_Expecter {
	return &DeploymentExpansion_Expecter{mock: &_m.Mock}
}

// Rollback provides a mock function with given fields: _a0, _a1, _a2
func (_m *DeploymentExpansion) Rollback(_a0 context.Context, _a1 *v1beta1.DeploymentRollback, _a2 v1.CreateOptions) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.DeploymentRollback, v1.CreateOptions) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeploymentExpansion_Rollback_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Rollback'
type DeploymentExpansion_Rollback_Call struct {
	*mock.Call
}

// Rollback is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *v1beta1.DeploymentRollback
//   - _a2 v1.CreateOptions
func (_e *DeploymentExpansion_Expecter) Rollback(_a0 interface{}, _a1 interface{}, _a2 interface{}) *DeploymentExpansion_Rollback_Call {
	return &DeploymentExpansion_Rollback_Call{Call: _e.mock.On("Rollback", _a0, _a1, _a2)}
}

func (_c *DeploymentExpansion_Rollback_Call) Run(run func(_a0 context.Context, _a1 *v1beta1.DeploymentRollback, _a2 v1.CreateOptions)) *DeploymentExpansion_Rollback_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1beta1.DeploymentRollback), args[2].(v1.CreateOptions))
	})
	return _c
}

func (_c *DeploymentExpansion_Rollback_Call) Return(_a0 error) *DeploymentExpansion_Rollback_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DeploymentExpansion_Rollback_Call) RunAndReturn(run func(context.Context, *v1beta1.DeploymentRollback, v1.CreateOptions) error) *DeploymentExpansion_Rollback_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewDeploymentExpansion interface {
	mock.TestingT
	Cleanup(func())
}

// NewDeploymentExpansion creates a new instance of DeploymentExpansion. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDeploymentExpansion(t mockConstructorTestingTNewDeploymentExpansion) *DeploymentExpansion {
	mock := &DeploymentExpansion{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
