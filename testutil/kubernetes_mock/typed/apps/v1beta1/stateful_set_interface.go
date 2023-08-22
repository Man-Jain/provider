// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	appsv1beta1 "k8s.io/api/apps/v1beta1"

	mock "github.com/stretchr/testify/mock"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1beta1 "k8s.io/client-go/applyconfigurations/apps/v1beta1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// StatefulSetInterface is an autogenerated mock type for the StatefulSetInterface type
type StatefulSetInterface struct {
	mock.Mock
}

type StatefulSetInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *StatefulSetInterface) EXPECT() *StatefulSetInterface_Expecter {
	return &StatefulSetInterface_Expecter{mock: &_m.Mock}
}

// Apply provides a mock function with given fields: ctx, statefulSet, opts
func (_m *StatefulSetInterface) Apply(ctx context.Context, statefulSet *v1beta1.StatefulSetApplyConfiguration, opts v1.ApplyOptions) (*appsv1beta1.StatefulSet, error) {
	ret := _m.Called(ctx, statefulSet, opts)

	var r0 *appsv1beta1.StatefulSet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.StatefulSetApplyConfiguration, v1.ApplyOptions) (*appsv1beta1.StatefulSet, error)); ok {
		return rf(ctx, statefulSet, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.StatefulSetApplyConfiguration, v1.ApplyOptions) *appsv1beta1.StatefulSet); ok {
		r0 = rf(ctx, statefulSet, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1beta1.StatefulSet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1beta1.StatefulSetApplyConfiguration, v1.ApplyOptions) error); ok {
		r1 = rf(ctx, statefulSet, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StatefulSetInterface_Apply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Apply'
type StatefulSetInterface_Apply_Call struct {
	*mock.Call
}

// Apply is a helper method to define mock.On call
//   - ctx context.Context
//   - statefulSet *v1beta1.StatefulSetApplyConfiguration
//   - opts v1.ApplyOptions
func (_e *StatefulSetInterface_Expecter) Apply(ctx interface{}, statefulSet interface{}, opts interface{}) *StatefulSetInterface_Apply_Call {
	return &StatefulSetInterface_Apply_Call{Call: _e.mock.On("Apply", ctx, statefulSet, opts)}
}

func (_c *StatefulSetInterface_Apply_Call) Run(run func(ctx context.Context, statefulSet *v1beta1.StatefulSetApplyConfiguration, opts v1.ApplyOptions)) *StatefulSetInterface_Apply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1beta1.StatefulSetApplyConfiguration), args[2].(v1.ApplyOptions))
	})
	return _c
}

func (_c *StatefulSetInterface_Apply_Call) Return(result *appsv1beta1.StatefulSet, err error) *StatefulSetInterface_Apply_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *StatefulSetInterface_Apply_Call) RunAndReturn(run func(context.Context, *v1beta1.StatefulSetApplyConfiguration, v1.ApplyOptions) (*appsv1beta1.StatefulSet, error)) *StatefulSetInterface_Apply_Call {
	_c.Call.Return(run)
	return _c
}

// ApplyStatus provides a mock function with given fields: ctx, statefulSet, opts
func (_m *StatefulSetInterface) ApplyStatus(ctx context.Context, statefulSet *v1beta1.StatefulSetApplyConfiguration, opts v1.ApplyOptions) (*appsv1beta1.StatefulSet, error) {
	ret := _m.Called(ctx, statefulSet, opts)

	var r0 *appsv1beta1.StatefulSet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.StatefulSetApplyConfiguration, v1.ApplyOptions) (*appsv1beta1.StatefulSet, error)); ok {
		return rf(ctx, statefulSet, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.StatefulSetApplyConfiguration, v1.ApplyOptions) *appsv1beta1.StatefulSet); ok {
		r0 = rf(ctx, statefulSet, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1beta1.StatefulSet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1beta1.StatefulSetApplyConfiguration, v1.ApplyOptions) error); ok {
		r1 = rf(ctx, statefulSet, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StatefulSetInterface_ApplyStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ApplyStatus'
type StatefulSetInterface_ApplyStatus_Call struct {
	*mock.Call
}

// ApplyStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - statefulSet *v1beta1.StatefulSetApplyConfiguration
//   - opts v1.ApplyOptions
func (_e *StatefulSetInterface_Expecter) ApplyStatus(ctx interface{}, statefulSet interface{}, opts interface{}) *StatefulSetInterface_ApplyStatus_Call {
	return &StatefulSetInterface_ApplyStatus_Call{Call: _e.mock.On("ApplyStatus", ctx, statefulSet, opts)}
}

func (_c *StatefulSetInterface_ApplyStatus_Call) Run(run func(ctx context.Context, statefulSet *v1beta1.StatefulSetApplyConfiguration, opts v1.ApplyOptions)) *StatefulSetInterface_ApplyStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1beta1.StatefulSetApplyConfiguration), args[2].(v1.ApplyOptions))
	})
	return _c
}

func (_c *StatefulSetInterface_ApplyStatus_Call) Return(result *appsv1beta1.StatefulSet, err error) *StatefulSetInterface_ApplyStatus_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *StatefulSetInterface_ApplyStatus_Call) RunAndReturn(run func(context.Context, *v1beta1.StatefulSetApplyConfiguration, v1.ApplyOptions) (*appsv1beta1.StatefulSet, error)) *StatefulSetInterface_ApplyStatus_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, statefulSet, opts
func (_m *StatefulSetInterface) Create(ctx context.Context, statefulSet *appsv1beta1.StatefulSet, opts v1.CreateOptions) (*appsv1beta1.StatefulSet, error) {
	ret := _m.Called(ctx, statefulSet, opts)

	var r0 *appsv1beta1.StatefulSet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appsv1beta1.StatefulSet, v1.CreateOptions) (*appsv1beta1.StatefulSet, error)); ok {
		return rf(ctx, statefulSet, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appsv1beta1.StatefulSet, v1.CreateOptions) *appsv1beta1.StatefulSet); ok {
		r0 = rf(ctx, statefulSet, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1beta1.StatefulSet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appsv1beta1.StatefulSet, v1.CreateOptions) error); ok {
		r1 = rf(ctx, statefulSet, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StatefulSetInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type StatefulSetInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - statefulSet *appsv1beta1.StatefulSet
//   - opts v1.CreateOptions
func (_e *StatefulSetInterface_Expecter) Create(ctx interface{}, statefulSet interface{}, opts interface{}) *StatefulSetInterface_Create_Call {
	return &StatefulSetInterface_Create_Call{Call: _e.mock.On("Create", ctx, statefulSet, opts)}
}

func (_c *StatefulSetInterface_Create_Call) Run(run func(ctx context.Context, statefulSet *appsv1beta1.StatefulSet, opts v1.CreateOptions)) *StatefulSetInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*appsv1beta1.StatefulSet), args[2].(v1.CreateOptions))
	})
	return _c
}

func (_c *StatefulSetInterface_Create_Call) Return(_a0 *appsv1beta1.StatefulSet, _a1 error) *StatefulSetInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StatefulSetInterface_Create_Call) RunAndReturn(run func(context.Context, *appsv1beta1.StatefulSet, v1.CreateOptions) (*appsv1beta1.StatefulSet, error)) *StatefulSetInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *StatefulSetInterface) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StatefulSetInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type StatefulSetInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts v1.DeleteOptions
func (_e *StatefulSetInterface_Expecter) Delete(ctx interface{}, name interface{}, opts interface{}) *StatefulSetInterface_Delete_Call {
	return &StatefulSetInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, name, opts)}
}

func (_c *StatefulSetInterface_Delete_Call) Run(run func(ctx context.Context, name string, opts v1.DeleteOptions)) *StatefulSetInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(v1.DeleteOptions))
	})
	return _c
}

func (_c *StatefulSetInterface_Delete_Call) Return(_a0 error) *StatefulSetInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StatefulSetInterface_Delete_Call) RunAndReturn(run func(context.Context, string, v1.DeleteOptions) error) *StatefulSetInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteCollection provides a mock function with given fields: ctx, opts, listOpts
func (_m *StatefulSetInterface) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	ret := _m.Called(ctx, opts, listOpts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.DeleteOptions, v1.ListOptions) error); ok {
		r0 = rf(ctx, opts, listOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StatefulSetInterface_DeleteCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCollection'
type StatefulSetInterface_DeleteCollection_Call struct {
	*mock.Call
}

// DeleteCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - opts v1.DeleteOptions
//   - listOpts v1.ListOptions
func (_e *StatefulSetInterface_Expecter) DeleteCollection(ctx interface{}, opts interface{}, listOpts interface{}) *StatefulSetInterface_DeleteCollection_Call {
	return &StatefulSetInterface_DeleteCollection_Call{Call: _e.mock.On("DeleteCollection", ctx, opts, listOpts)}
}

func (_c *StatefulSetInterface_DeleteCollection_Call) Run(run func(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions)) *StatefulSetInterface_DeleteCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.DeleteOptions), args[2].(v1.ListOptions))
	})
	return _c
}

func (_c *StatefulSetInterface_DeleteCollection_Call) Return(_a0 error) *StatefulSetInterface_DeleteCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StatefulSetInterface_DeleteCollection_Call) RunAndReturn(run func(context.Context, v1.DeleteOptions, v1.ListOptions) error) *StatefulSetInterface_DeleteCollection_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *StatefulSetInterface) Get(ctx context.Context, name string, opts v1.GetOptions) (*appsv1beta1.StatefulSet, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *appsv1beta1.StatefulSet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.GetOptions) (*appsv1beta1.StatefulSet, error)); ok {
		return rf(ctx, name, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.GetOptions) *appsv1beta1.StatefulSet); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1beta1.StatefulSet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, v1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StatefulSetInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type StatefulSetInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts v1.GetOptions
func (_e *StatefulSetInterface_Expecter) Get(ctx interface{}, name interface{}, opts interface{}) *StatefulSetInterface_Get_Call {
	return &StatefulSetInterface_Get_Call{Call: _e.mock.On("Get", ctx, name, opts)}
}

func (_c *StatefulSetInterface_Get_Call) Run(run func(ctx context.Context, name string, opts v1.GetOptions)) *StatefulSetInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(v1.GetOptions))
	})
	return _c
}

func (_c *StatefulSetInterface_Get_Call) Return(_a0 *appsv1beta1.StatefulSet, _a1 error) *StatefulSetInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StatefulSetInterface_Get_Call) RunAndReturn(run func(context.Context, string, v1.GetOptions) (*appsv1beta1.StatefulSet, error)) *StatefulSetInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, opts
func (_m *StatefulSetInterface) List(ctx context.Context, opts v1.ListOptions) (*appsv1beta1.StatefulSetList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *appsv1beta1.StatefulSetList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) (*appsv1beta1.StatefulSetList, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) *appsv1beta1.StatefulSetList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1beta1.StatefulSetList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, v1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StatefulSetInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type StatefulSetInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - opts v1.ListOptions
func (_e *StatefulSetInterface_Expecter) List(ctx interface{}, opts interface{}) *StatefulSetInterface_List_Call {
	return &StatefulSetInterface_List_Call{Call: _e.mock.On("List", ctx, opts)}
}

func (_c *StatefulSetInterface_List_Call) Run(run func(ctx context.Context, opts v1.ListOptions)) *StatefulSetInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.ListOptions))
	})
	return _c
}

func (_c *StatefulSetInterface_List_Call) Return(_a0 *appsv1beta1.StatefulSetList, _a1 error) *StatefulSetInterface_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StatefulSetInterface_List_Call) RunAndReturn(run func(context.Context, v1.ListOptions) (*appsv1beta1.StatefulSetList, error)) *StatefulSetInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// Patch provides a mock function with given fields: ctx, name, pt, data, opts, subresources
func (_m *StatefulSetInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (*appsv1beta1.StatefulSet, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *appsv1beta1.StatefulSet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) (*appsv1beta1.StatefulSet, error)); ok {
		return rf(ctx, name, pt, data, opts, subresources...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) *appsv1beta1.StatefulSet); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1beta1.StatefulSet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StatefulSetInterface_Patch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Patch'
type StatefulSetInterface_Patch_Call struct {
	*mock.Call
}

// Patch is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - pt types.PatchType
//   - data []byte
//   - opts v1.PatchOptions
//   - subresources ...string
func (_e *StatefulSetInterface_Expecter) Patch(ctx interface{}, name interface{}, pt interface{}, data interface{}, opts interface{}, subresources ...interface{}) *StatefulSetInterface_Patch_Call {
	return &StatefulSetInterface_Patch_Call{Call: _e.mock.On("Patch",
		append([]interface{}{ctx, name, pt, data, opts}, subresources...)...)}
}

func (_c *StatefulSetInterface_Patch_Call) Run(run func(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string)) *StatefulSetInterface_Patch_Call {
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

func (_c *StatefulSetInterface_Patch_Call) Return(result *appsv1beta1.StatefulSet, err error) *StatefulSetInterface_Patch_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *StatefulSetInterface_Patch_Call) RunAndReturn(run func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) (*appsv1beta1.StatefulSet, error)) *StatefulSetInterface_Patch_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, statefulSet, opts
func (_m *StatefulSetInterface) Update(ctx context.Context, statefulSet *appsv1beta1.StatefulSet, opts v1.UpdateOptions) (*appsv1beta1.StatefulSet, error) {
	ret := _m.Called(ctx, statefulSet, opts)

	var r0 *appsv1beta1.StatefulSet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appsv1beta1.StatefulSet, v1.UpdateOptions) (*appsv1beta1.StatefulSet, error)); ok {
		return rf(ctx, statefulSet, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appsv1beta1.StatefulSet, v1.UpdateOptions) *appsv1beta1.StatefulSet); ok {
		r0 = rf(ctx, statefulSet, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1beta1.StatefulSet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appsv1beta1.StatefulSet, v1.UpdateOptions) error); ok {
		r1 = rf(ctx, statefulSet, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StatefulSetInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type StatefulSetInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - statefulSet *appsv1beta1.StatefulSet
//   - opts v1.UpdateOptions
func (_e *StatefulSetInterface_Expecter) Update(ctx interface{}, statefulSet interface{}, opts interface{}) *StatefulSetInterface_Update_Call {
	return &StatefulSetInterface_Update_Call{Call: _e.mock.On("Update", ctx, statefulSet, opts)}
}

func (_c *StatefulSetInterface_Update_Call) Run(run func(ctx context.Context, statefulSet *appsv1beta1.StatefulSet, opts v1.UpdateOptions)) *StatefulSetInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*appsv1beta1.StatefulSet), args[2].(v1.UpdateOptions))
	})
	return _c
}

func (_c *StatefulSetInterface_Update_Call) Return(_a0 *appsv1beta1.StatefulSet, _a1 error) *StatefulSetInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StatefulSetInterface_Update_Call) RunAndReturn(run func(context.Context, *appsv1beta1.StatefulSet, v1.UpdateOptions) (*appsv1beta1.StatefulSet, error)) *StatefulSetInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateStatus provides a mock function with given fields: ctx, statefulSet, opts
func (_m *StatefulSetInterface) UpdateStatus(ctx context.Context, statefulSet *appsv1beta1.StatefulSet, opts v1.UpdateOptions) (*appsv1beta1.StatefulSet, error) {
	ret := _m.Called(ctx, statefulSet, opts)

	var r0 *appsv1beta1.StatefulSet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appsv1beta1.StatefulSet, v1.UpdateOptions) (*appsv1beta1.StatefulSet, error)); ok {
		return rf(ctx, statefulSet, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appsv1beta1.StatefulSet, v1.UpdateOptions) *appsv1beta1.StatefulSet); ok {
		r0 = rf(ctx, statefulSet, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1beta1.StatefulSet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appsv1beta1.StatefulSet, v1.UpdateOptions) error); ok {
		r1 = rf(ctx, statefulSet, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StatefulSetInterface_UpdateStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateStatus'
type StatefulSetInterface_UpdateStatus_Call struct {
	*mock.Call
}

// UpdateStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - statefulSet *appsv1beta1.StatefulSet
//   - opts v1.UpdateOptions
func (_e *StatefulSetInterface_Expecter) UpdateStatus(ctx interface{}, statefulSet interface{}, opts interface{}) *StatefulSetInterface_UpdateStatus_Call {
	return &StatefulSetInterface_UpdateStatus_Call{Call: _e.mock.On("UpdateStatus", ctx, statefulSet, opts)}
}

func (_c *StatefulSetInterface_UpdateStatus_Call) Run(run func(ctx context.Context, statefulSet *appsv1beta1.StatefulSet, opts v1.UpdateOptions)) *StatefulSetInterface_UpdateStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*appsv1beta1.StatefulSet), args[2].(v1.UpdateOptions))
	})
	return _c
}

func (_c *StatefulSetInterface_UpdateStatus_Call) Return(_a0 *appsv1beta1.StatefulSet, _a1 error) *StatefulSetInterface_UpdateStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StatefulSetInterface_UpdateStatus_Call) RunAndReturn(run func(context.Context, *appsv1beta1.StatefulSet, v1.UpdateOptions) (*appsv1beta1.StatefulSet, error)) *StatefulSetInterface_UpdateStatus_Call {
	_c.Call.Return(run)
	return _c
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *StatefulSetInterface) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
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

// StatefulSetInterface_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type StatefulSetInterface_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - ctx context.Context
//   - opts v1.ListOptions
func (_e *StatefulSetInterface_Expecter) Watch(ctx interface{}, opts interface{}) *StatefulSetInterface_Watch_Call {
	return &StatefulSetInterface_Watch_Call{Call: _e.mock.On("Watch", ctx, opts)}
}

func (_c *StatefulSetInterface_Watch_Call) Run(run func(ctx context.Context, opts v1.ListOptions)) *StatefulSetInterface_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.ListOptions))
	})
	return _c
}

func (_c *StatefulSetInterface_Watch_Call) Return(_a0 watch.Interface, _a1 error) *StatefulSetInterface_Watch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StatefulSetInterface_Watch_Call) RunAndReturn(run func(context.Context, v1.ListOptions) (watch.Interface, error)) *StatefulSetInterface_Watch_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewStatefulSetInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewStatefulSetInterface creates a new instance of StatefulSetInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStatefulSetInterface(t mockConstructorTestingTNewStatefulSetInterface) *StatefulSetInterface {
	mock := &StatefulSetInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
