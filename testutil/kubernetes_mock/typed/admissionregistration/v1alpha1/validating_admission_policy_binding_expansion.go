// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import mock "github.com/stretchr/testify/mock"

// ValidatingAdmissionPolicyBindingExpansion is an autogenerated mock type for the ValidatingAdmissionPolicyBindingExpansion type
type ValidatingAdmissionPolicyBindingExpansion struct {
	mock.Mock
}

type ValidatingAdmissionPolicyBindingExpansion_Expecter struct {
	mock *mock.Mock
}

func (_m *ValidatingAdmissionPolicyBindingExpansion) EXPECT() *ValidatingAdmissionPolicyBindingExpansion_Expecter {
	return &ValidatingAdmissionPolicyBindingExpansion_Expecter{mock: &_m.Mock}
}

type mockConstructorTestingTNewValidatingAdmissionPolicyBindingExpansion interface {
	mock.TestingT
	Cleanup(func())
}

// NewValidatingAdmissionPolicyBindingExpansion creates a new instance of ValidatingAdmissionPolicyBindingExpansion. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewValidatingAdmissionPolicyBindingExpansion(t mockConstructorTestingTNewValidatingAdmissionPolicyBindingExpansion) *ValidatingAdmissionPolicyBindingExpansion {
	mock := &ValidatingAdmissionPolicyBindingExpansion{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}