// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	schedulingv1alpha1 "k8s.io/api/scheduling/v1alpha1"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1alpha1 "k8s.io/client-go/applyconfigurations/scheduling/v1alpha1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// PriorityClassInterface is an autogenerated mock type for the PriorityClassInterface type
type PriorityClassInterface struct {
	mock.Mock
}

type PriorityClassInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *PriorityClassInterface) EXPECT() *PriorityClassInterface_Expecter {
	return &PriorityClassInterface_Expecter{mock: &_m.Mock}
}

// Apply provides a mock function with given fields: ctx, priorityClass, opts
func (_m *PriorityClassInterface) Apply(ctx context.Context, priorityClass *v1alpha1.PriorityClassApplyConfiguration, opts v1.ApplyOptions) (*schedulingv1alpha1.PriorityClass, error) {
	ret := _m.Called(ctx, priorityClass, opts)

	var r0 *schedulingv1alpha1.PriorityClass
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.PriorityClassApplyConfiguration, v1.ApplyOptions) (*schedulingv1alpha1.PriorityClass, error)); ok {
		return rf(ctx, priorityClass, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.PriorityClassApplyConfiguration, v1.ApplyOptions) *schedulingv1alpha1.PriorityClass); ok {
		r0 = rf(ctx, priorityClass, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*schedulingv1alpha1.PriorityClass)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1alpha1.PriorityClassApplyConfiguration, v1.ApplyOptions) error); ok {
		r1 = rf(ctx, priorityClass, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PriorityClassInterface_Apply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Apply'
type PriorityClassInterface_Apply_Call struct {
	*mock.Call
}

// Apply is a helper method to define mock.On call
//   - ctx context.Context
//   - priorityClass *v1alpha1.PriorityClassApplyConfiguration
//   - opts v1.ApplyOptions
func (_e *PriorityClassInterface_Expecter) Apply(ctx interface{}, priorityClass interface{}, opts interface{}) *PriorityClassInterface_Apply_Call {
	return &PriorityClassInterface_Apply_Call{Call: _e.mock.On("Apply", ctx, priorityClass, opts)}
}

func (_c *PriorityClassInterface_Apply_Call) Run(run func(ctx context.Context, priorityClass *v1alpha1.PriorityClassApplyConfiguration, opts v1.ApplyOptions)) *PriorityClassInterface_Apply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1alpha1.PriorityClassApplyConfiguration), args[2].(v1.ApplyOptions))
	})
	return _c
}

func (_c *PriorityClassInterface_Apply_Call) Return(result *schedulingv1alpha1.PriorityClass, err error) *PriorityClassInterface_Apply_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *PriorityClassInterface_Apply_Call) RunAndReturn(run func(context.Context, *v1alpha1.PriorityClassApplyConfiguration, v1.ApplyOptions) (*schedulingv1alpha1.PriorityClass, error)) *PriorityClassInterface_Apply_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, priorityClass, opts
func (_m *PriorityClassInterface) Create(ctx context.Context, priorityClass *schedulingv1alpha1.PriorityClass, opts v1.CreateOptions) (*schedulingv1alpha1.PriorityClass, error) {
	ret := _m.Called(ctx, priorityClass, opts)

	var r0 *schedulingv1alpha1.PriorityClass
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *schedulingv1alpha1.PriorityClass, v1.CreateOptions) (*schedulingv1alpha1.PriorityClass, error)); ok {
		return rf(ctx, priorityClass, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *schedulingv1alpha1.PriorityClass, v1.CreateOptions) *schedulingv1alpha1.PriorityClass); ok {
		r0 = rf(ctx, priorityClass, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*schedulingv1alpha1.PriorityClass)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *schedulingv1alpha1.PriorityClass, v1.CreateOptions) error); ok {
		r1 = rf(ctx, priorityClass, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PriorityClassInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type PriorityClassInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - priorityClass *schedulingv1alpha1.PriorityClass
//   - opts v1.CreateOptions
func (_e *PriorityClassInterface_Expecter) Create(ctx interface{}, priorityClass interface{}, opts interface{}) *PriorityClassInterface_Create_Call {
	return &PriorityClassInterface_Create_Call{Call: _e.mock.On("Create", ctx, priorityClass, opts)}
}

func (_c *PriorityClassInterface_Create_Call) Run(run func(ctx context.Context, priorityClass *schedulingv1alpha1.PriorityClass, opts v1.CreateOptions)) *PriorityClassInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*schedulingv1alpha1.PriorityClass), args[2].(v1.CreateOptions))
	})
	return _c
}

func (_c *PriorityClassInterface_Create_Call) Return(_a0 *schedulingv1alpha1.PriorityClass, _a1 error) *PriorityClassInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PriorityClassInterface_Create_Call) RunAndReturn(run func(context.Context, *schedulingv1alpha1.PriorityClass, v1.CreateOptions) (*schedulingv1alpha1.PriorityClass, error)) *PriorityClassInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *PriorityClassInterface) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PriorityClassInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type PriorityClassInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts v1.DeleteOptions
func (_e *PriorityClassInterface_Expecter) Delete(ctx interface{}, name interface{}, opts interface{}) *PriorityClassInterface_Delete_Call {
	return &PriorityClassInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, name, opts)}
}

func (_c *PriorityClassInterface_Delete_Call) Run(run func(ctx context.Context, name string, opts v1.DeleteOptions)) *PriorityClassInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(v1.DeleteOptions))
	})
	return _c
}

func (_c *PriorityClassInterface_Delete_Call) Return(_a0 error) *PriorityClassInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PriorityClassInterface_Delete_Call) RunAndReturn(run func(context.Context, string, v1.DeleteOptions) error) *PriorityClassInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteCollection provides a mock function with given fields: ctx, opts, listOpts
func (_m *PriorityClassInterface) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	ret := _m.Called(ctx, opts, listOpts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.DeleteOptions, v1.ListOptions) error); ok {
		r0 = rf(ctx, opts, listOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PriorityClassInterface_DeleteCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCollection'
type PriorityClassInterface_DeleteCollection_Call struct {
	*mock.Call
}

// DeleteCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - opts v1.DeleteOptions
//   - listOpts v1.ListOptions
func (_e *PriorityClassInterface_Expecter) DeleteCollection(ctx interface{}, opts interface{}, listOpts interface{}) *PriorityClassInterface_DeleteCollection_Call {
	return &PriorityClassInterface_DeleteCollection_Call{Call: _e.mock.On("DeleteCollection", ctx, opts, listOpts)}
}

func (_c *PriorityClassInterface_DeleteCollection_Call) Run(run func(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions)) *PriorityClassInterface_DeleteCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.DeleteOptions), args[2].(v1.ListOptions))
	})
	return _c
}

func (_c *PriorityClassInterface_DeleteCollection_Call) Return(_a0 error) *PriorityClassInterface_DeleteCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PriorityClassInterface_DeleteCollection_Call) RunAndReturn(run func(context.Context, v1.DeleteOptions, v1.ListOptions) error) *PriorityClassInterface_DeleteCollection_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *PriorityClassInterface) Get(ctx context.Context, name string, opts v1.GetOptions) (*schedulingv1alpha1.PriorityClass, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *schedulingv1alpha1.PriorityClass
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.GetOptions) (*schedulingv1alpha1.PriorityClass, error)); ok {
		return rf(ctx, name, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.GetOptions) *schedulingv1alpha1.PriorityClass); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*schedulingv1alpha1.PriorityClass)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, v1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PriorityClassInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type PriorityClassInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts v1.GetOptions
func (_e *PriorityClassInterface_Expecter) Get(ctx interface{}, name interface{}, opts interface{}) *PriorityClassInterface_Get_Call {
	return &PriorityClassInterface_Get_Call{Call: _e.mock.On("Get", ctx, name, opts)}
}

func (_c *PriorityClassInterface_Get_Call) Run(run func(ctx context.Context, name string, opts v1.GetOptions)) *PriorityClassInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(v1.GetOptions))
	})
	return _c
}

func (_c *PriorityClassInterface_Get_Call) Return(_a0 *schedulingv1alpha1.PriorityClass, _a1 error) *PriorityClassInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PriorityClassInterface_Get_Call) RunAndReturn(run func(context.Context, string, v1.GetOptions) (*schedulingv1alpha1.PriorityClass, error)) *PriorityClassInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, opts
func (_m *PriorityClassInterface) List(ctx context.Context, opts v1.ListOptions) (*schedulingv1alpha1.PriorityClassList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *schedulingv1alpha1.PriorityClassList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) (*schedulingv1alpha1.PriorityClassList, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) *schedulingv1alpha1.PriorityClassList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*schedulingv1alpha1.PriorityClassList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, v1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PriorityClassInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type PriorityClassInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - opts v1.ListOptions
func (_e *PriorityClassInterface_Expecter) List(ctx interface{}, opts interface{}) *PriorityClassInterface_List_Call {
	return &PriorityClassInterface_List_Call{Call: _e.mock.On("List", ctx, opts)}
}

func (_c *PriorityClassInterface_List_Call) Run(run func(ctx context.Context, opts v1.ListOptions)) *PriorityClassInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.ListOptions))
	})
	return _c
}

func (_c *PriorityClassInterface_List_Call) Return(_a0 *schedulingv1alpha1.PriorityClassList, _a1 error) *PriorityClassInterface_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PriorityClassInterface_List_Call) RunAndReturn(run func(context.Context, v1.ListOptions) (*schedulingv1alpha1.PriorityClassList, error)) *PriorityClassInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// Patch provides a mock function with given fields: ctx, name, pt, data, opts, subresources
func (_m *PriorityClassInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (*schedulingv1alpha1.PriorityClass, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *schedulingv1alpha1.PriorityClass
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) (*schedulingv1alpha1.PriorityClass, error)); ok {
		return rf(ctx, name, pt, data, opts, subresources...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) *schedulingv1alpha1.PriorityClass); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*schedulingv1alpha1.PriorityClass)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PriorityClassInterface_Patch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Patch'
type PriorityClassInterface_Patch_Call struct {
	*mock.Call
}

// Patch is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - pt types.PatchType
//   - data []byte
//   - opts v1.PatchOptions
//   - subresources ...string
func (_e *PriorityClassInterface_Expecter) Patch(ctx interface{}, name interface{}, pt interface{}, data interface{}, opts interface{}, subresources ...interface{}) *PriorityClassInterface_Patch_Call {
	return &PriorityClassInterface_Patch_Call{Call: _e.mock.On("Patch",
		append([]interface{}{ctx, name, pt, data, opts}, subresources...)...)}
}

func (_c *PriorityClassInterface_Patch_Call) Run(run func(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string)) *PriorityClassInterface_Patch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-5)
		for i, a := range args[5:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(args[0].(context.Context), args[1].(string), args[2].(types.PatchType), args[3].([]byte), args[4].(v1.PatchOptions), variadicArgs...)
	})
	return _c
}

func (_c *PriorityClassInterface_Patch_Call) Return(result *schedulingv1alpha1.PriorityClass, err error) *PriorityClassInterface_Patch_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *PriorityClassInterface_Patch_Call) RunAndReturn(run func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) (*schedulingv1alpha1.PriorityClass, error)) *PriorityClassInterface_Patch_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, priorityClass, opts
func (_m *PriorityClassInterface) Update(ctx context.Context, priorityClass *schedulingv1alpha1.PriorityClass, opts v1.UpdateOptions) (*schedulingv1alpha1.PriorityClass, error) {
	ret := _m.Called(ctx, priorityClass, opts)

	var r0 *schedulingv1alpha1.PriorityClass
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *schedulingv1alpha1.PriorityClass, v1.UpdateOptions) (*schedulingv1alpha1.PriorityClass, error)); ok {
		return rf(ctx, priorityClass, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *schedulingv1alpha1.PriorityClass, v1.UpdateOptions) *schedulingv1alpha1.PriorityClass); ok {
		r0 = rf(ctx, priorityClass, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*schedulingv1alpha1.PriorityClass)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *schedulingv1alpha1.PriorityClass, v1.UpdateOptions) error); ok {
		r1 = rf(ctx, priorityClass, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PriorityClassInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type PriorityClassInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - priorityClass *schedulingv1alpha1.PriorityClass
//   - opts v1.UpdateOptions
func (_e *PriorityClassInterface_Expecter) Update(ctx interface{}, priorityClass interface{}, opts interface{}) *PriorityClassInterface_Update_Call {
	return &PriorityClassInterface_Update_Call{Call: _e.mock.On("Update", ctx, priorityClass, opts)}
}

func (_c *PriorityClassInterface_Update_Call) Run(run func(ctx context.Context, priorityClass *schedulingv1alpha1.PriorityClass, opts v1.UpdateOptions)) *PriorityClassInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*schedulingv1alpha1.PriorityClass), args[2].(v1.UpdateOptions))
	})
	return _c
}

func (_c *PriorityClassInterface_Update_Call) Return(_a0 *schedulingv1alpha1.PriorityClass, _a1 error) *PriorityClassInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PriorityClassInterface_Update_Call) RunAndReturn(run func(context.Context, *schedulingv1alpha1.PriorityClass, v1.UpdateOptions) (*schedulingv1alpha1.PriorityClass, error)) *PriorityClassInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *PriorityClassInterface) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	ret := _m.Called(ctx, opts)

	var r0 watch.Interface
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) (watch.Interface, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) watch.Interface); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(watch.Interface)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, v1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PriorityClassInterface_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type PriorityClassInterface_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - ctx context.Context
//   - opts v1.ListOptions
func (_e *PriorityClassInterface_Expecter) Watch(ctx interface{}, opts interface{}) *PriorityClassInterface_Watch_Call {
	return &PriorityClassInterface_Watch_Call{Call: _e.mock.On("Watch", ctx, opts)}
}

func (_c *PriorityClassInterface_Watch_Call) Run(run func(ctx context.Context, opts v1.ListOptions)) *PriorityClassInterface_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.ListOptions))
	})
	return _c
}

func (_c *PriorityClassInterface_Watch_Call) Return(_a0 watch.Interface, _a1 error) *PriorityClassInterface_Watch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PriorityClassInterface_Watch_Call) RunAndReturn(run func(context.Context, v1.ListOptions) (watch.Interface, error)) *PriorityClassInterface_Watch_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewPriorityClassInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewPriorityClassInterface creates a new instance of PriorityClassInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPriorityClassInterface(t mockConstructorTestingTNewPriorityClassInterface) *PriorityClassInterface {
	mock := &PriorityClassInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
