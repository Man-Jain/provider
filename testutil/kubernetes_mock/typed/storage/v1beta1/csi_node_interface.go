// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	storagev1beta1 "k8s.io/api/storage/v1beta1"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1beta1 "k8s.io/client-go/applyconfigurations/storage/v1beta1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// CSINodeInterface is an autogenerated mock type for the CSINodeInterface type
type CSINodeInterface struct {
	mock.Mock
}

type CSINodeInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *CSINodeInterface) EXPECT() *CSINodeInterface_Expecter {
	return &CSINodeInterface_Expecter{mock: &_m.Mock}
}

// Apply provides a mock function with given fields: ctx, cSINode, opts
func (_m *CSINodeInterface) Apply(ctx context.Context, cSINode *v1beta1.CSINodeApplyConfiguration, opts v1.ApplyOptions) (*storagev1beta1.CSINode, error) {
	ret := _m.Called(ctx, cSINode, opts)

	var r0 *storagev1beta1.CSINode
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.CSINodeApplyConfiguration, v1.ApplyOptions) (*storagev1beta1.CSINode, error)); ok {
		return rf(ctx, cSINode, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.CSINodeApplyConfiguration, v1.ApplyOptions) *storagev1beta1.CSINode); ok {
		r0 = rf(ctx, cSINode, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storagev1beta1.CSINode)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1beta1.CSINodeApplyConfiguration, v1.ApplyOptions) error); ok {
		r1 = rf(ctx, cSINode, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CSINodeInterface_Apply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Apply'
type CSINodeInterface_Apply_Call struct {
	*mock.Call
}

// Apply is a helper method to define mock.On call
//   - ctx context.Context
//   - cSINode *v1beta1.CSINodeApplyConfiguration
//   - opts v1.ApplyOptions
func (_e *CSINodeInterface_Expecter) Apply(ctx interface{}, cSINode interface{}, opts interface{}) *CSINodeInterface_Apply_Call {
	return &CSINodeInterface_Apply_Call{Call: _e.mock.On("Apply", ctx, cSINode, opts)}
}

func (_c *CSINodeInterface_Apply_Call) Run(run func(ctx context.Context, cSINode *v1beta1.CSINodeApplyConfiguration, opts v1.ApplyOptions)) *CSINodeInterface_Apply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1beta1.CSINodeApplyConfiguration), args[2].(v1.ApplyOptions))
	})
	return _c
}

func (_c *CSINodeInterface_Apply_Call) Return(result *storagev1beta1.CSINode, err error) *CSINodeInterface_Apply_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *CSINodeInterface_Apply_Call) RunAndReturn(run func(context.Context, *v1beta1.CSINodeApplyConfiguration, v1.ApplyOptions) (*storagev1beta1.CSINode, error)) *CSINodeInterface_Apply_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, cSINode, opts
func (_m *CSINodeInterface) Create(ctx context.Context, cSINode *storagev1beta1.CSINode, opts v1.CreateOptions) (*storagev1beta1.CSINode, error) {
	ret := _m.Called(ctx, cSINode, opts)

	var r0 *storagev1beta1.CSINode
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storagev1beta1.CSINode, v1.CreateOptions) (*storagev1beta1.CSINode, error)); ok {
		return rf(ctx, cSINode, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storagev1beta1.CSINode, v1.CreateOptions) *storagev1beta1.CSINode); ok {
		r0 = rf(ctx, cSINode, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storagev1beta1.CSINode)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storagev1beta1.CSINode, v1.CreateOptions) error); ok {
		r1 = rf(ctx, cSINode, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CSINodeInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type CSINodeInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - cSINode *storagev1beta1.CSINode
//   - opts v1.CreateOptions
func (_e *CSINodeInterface_Expecter) Create(ctx interface{}, cSINode interface{}, opts interface{}) *CSINodeInterface_Create_Call {
	return &CSINodeInterface_Create_Call{Call: _e.mock.On("Create", ctx, cSINode, opts)}
}

func (_c *CSINodeInterface_Create_Call) Run(run func(ctx context.Context, cSINode *storagev1beta1.CSINode, opts v1.CreateOptions)) *CSINodeInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*storagev1beta1.CSINode), args[2].(v1.CreateOptions))
	})
	return _c
}

func (_c *CSINodeInterface_Create_Call) Return(_a0 *storagev1beta1.CSINode, _a1 error) *CSINodeInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CSINodeInterface_Create_Call) RunAndReturn(run func(context.Context, *storagev1beta1.CSINode, v1.CreateOptions) (*storagev1beta1.CSINode, error)) *CSINodeInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *CSINodeInterface) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CSINodeInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type CSINodeInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts v1.DeleteOptions
func (_e *CSINodeInterface_Expecter) Delete(ctx interface{}, name interface{}, opts interface{}) *CSINodeInterface_Delete_Call {
	return &CSINodeInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, name, opts)}
}

func (_c *CSINodeInterface_Delete_Call) Run(run func(ctx context.Context, name string, opts v1.DeleteOptions)) *CSINodeInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(v1.DeleteOptions))
	})
	return _c
}

func (_c *CSINodeInterface_Delete_Call) Return(_a0 error) *CSINodeInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CSINodeInterface_Delete_Call) RunAndReturn(run func(context.Context, string, v1.DeleteOptions) error) *CSINodeInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteCollection provides a mock function with given fields: ctx, opts, listOpts
func (_m *CSINodeInterface) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	ret := _m.Called(ctx, opts, listOpts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.DeleteOptions, v1.ListOptions) error); ok {
		r0 = rf(ctx, opts, listOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CSINodeInterface_DeleteCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCollection'
type CSINodeInterface_DeleteCollection_Call struct {
	*mock.Call
}

// DeleteCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - opts v1.DeleteOptions
//   - listOpts v1.ListOptions
func (_e *CSINodeInterface_Expecter) DeleteCollection(ctx interface{}, opts interface{}, listOpts interface{}) *CSINodeInterface_DeleteCollection_Call {
	return &CSINodeInterface_DeleteCollection_Call{Call: _e.mock.On("DeleteCollection", ctx, opts, listOpts)}
}

func (_c *CSINodeInterface_DeleteCollection_Call) Run(run func(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions)) *CSINodeInterface_DeleteCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.DeleteOptions), args[2].(v1.ListOptions))
	})
	return _c
}

func (_c *CSINodeInterface_DeleteCollection_Call) Return(_a0 error) *CSINodeInterface_DeleteCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CSINodeInterface_DeleteCollection_Call) RunAndReturn(run func(context.Context, v1.DeleteOptions, v1.ListOptions) error) *CSINodeInterface_DeleteCollection_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *CSINodeInterface) Get(ctx context.Context, name string, opts v1.GetOptions) (*storagev1beta1.CSINode, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *storagev1beta1.CSINode
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.GetOptions) (*storagev1beta1.CSINode, error)); ok {
		return rf(ctx, name, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.GetOptions) *storagev1beta1.CSINode); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storagev1beta1.CSINode)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, v1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CSINodeInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type CSINodeInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts v1.GetOptions
func (_e *CSINodeInterface_Expecter) Get(ctx interface{}, name interface{}, opts interface{}) *CSINodeInterface_Get_Call {
	return &CSINodeInterface_Get_Call{Call: _e.mock.On("Get", ctx, name, opts)}
}

func (_c *CSINodeInterface_Get_Call) Run(run func(ctx context.Context, name string, opts v1.GetOptions)) *CSINodeInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(v1.GetOptions))
	})
	return _c
}

func (_c *CSINodeInterface_Get_Call) Return(_a0 *storagev1beta1.CSINode, _a1 error) *CSINodeInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CSINodeInterface_Get_Call) RunAndReturn(run func(context.Context, string, v1.GetOptions) (*storagev1beta1.CSINode, error)) *CSINodeInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, opts
func (_m *CSINodeInterface) List(ctx context.Context, opts v1.ListOptions) (*storagev1beta1.CSINodeList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *storagev1beta1.CSINodeList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) (*storagev1beta1.CSINodeList, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) *storagev1beta1.CSINodeList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storagev1beta1.CSINodeList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, v1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CSINodeInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type CSINodeInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - opts v1.ListOptions
func (_e *CSINodeInterface_Expecter) List(ctx interface{}, opts interface{}) *CSINodeInterface_List_Call {
	return &CSINodeInterface_List_Call{Call: _e.mock.On("List", ctx, opts)}
}

func (_c *CSINodeInterface_List_Call) Run(run func(ctx context.Context, opts v1.ListOptions)) *CSINodeInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.ListOptions))
	})
	return _c
}

func (_c *CSINodeInterface_List_Call) Return(_a0 *storagev1beta1.CSINodeList, _a1 error) *CSINodeInterface_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CSINodeInterface_List_Call) RunAndReturn(run func(context.Context, v1.ListOptions) (*storagev1beta1.CSINodeList, error)) *CSINodeInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// Patch provides a mock function with given fields: ctx, name, pt, data, opts, subresources
func (_m *CSINodeInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (*storagev1beta1.CSINode, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *storagev1beta1.CSINode
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) (*storagev1beta1.CSINode, error)); ok {
		return rf(ctx, name, pt, data, opts, subresources...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) *storagev1beta1.CSINode); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storagev1beta1.CSINode)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CSINodeInterface_Patch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Patch'
type CSINodeInterface_Patch_Call struct {
	*mock.Call
}

// Patch is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - pt types.PatchType
//   - data []byte
//   - opts v1.PatchOptions
//   - subresources ...string
func (_e *CSINodeInterface_Expecter) Patch(ctx interface{}, name interface{}, pt interface{}, data interface{}, opts interface{}, subresources ...interface{}) *CSINodeInterface_Patch_Call {
	return &CSINodeInterface_Patch_Call{Call: _e.mock.On("Patch",
		append([]interface{}{ctx, name, pt, data, opts}, subresources...)...)}
}

func (_c *CSINodeInterface_Patch_Call) Run(run func(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string)) *CSINodeInterface_Patch_Call {
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

func (_c *CSINodeInterface_Patch_Call) Return(result *storagev1beta1.CSINode, err error) *CSINodeInterface_Patch_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *CSINodeInterface_Patch_Call) RunAndReturn(run func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) (*storagev1beta1.CSINode, error)) *CSINodeInterface_Patch_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, cSINode, opts
func (_m *CSINodeInterface) Update(ctx context.Context, cSINode *storagev1beta1.CSINode, opts v1.UpdateOptions) (*storagev1beta1.CSINode, error) {
	ret := _m.Called(ctx, cSINode, opts)

	var r0 *storagev1beta1.CSINode
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storagev1beta1.CSINode, v1.UpdateOptions) (*storagev1beta1.CSINode, error)); ok {
		return rf(ctx, cSINode, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storagev1beta1.CSINode, v1.UpdateOptions) *storagev1beta1.CSINode); ok {
		r0 = rf(ctx, cSINode, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storagev1beta1.CSINode)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storagev1beta1.CSINode, v1.UpdateOptions) error); ok {
		r1 = rf(ctx, cSINode, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CSINodeInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type CSINodeInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - cSINode *storagev1beta1.CSINode
//   - opts v1.UpdateOptions
func (_e *CSINodeInterface_Expecter) Update(ctx interface{}, cSINode interface{}, opts interface{}) *CSINodeInterface_Update_Call {
	return &CSINodeInterface_Update_Call{Call: _e.mock.On("Update", ctx, cSINode, opts)}
}

func (_c *CSINodeInterface_Update_Call) Run(run func(ctx context.Context, cSINode *storagev1beta1.CSINode, opts v1.UpdateOptions)) *CSINodeInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*storagev1beta1.CSINode), args[2].(v1.UpdateOptions))
	})
	return _c
}

func (_c *CSINodeInterface_Update_Call) Return(_a0 *storagev1beta1.CSINode, _a1 error) *CSINodeInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CSINodeInterface_Update_Call) RunAndReturn(run func(context.Context, *storagev1beta1.CSINode, v1.UpdateOptions) (*storagev1beta1.CSINode, error)) *CSINodeInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *CSINodeInterface) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
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

// CSINodeInterface_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type CSINodeInterface_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - ctx context.Context
//   - opts v1.ListOptions
func (_e *CSINodeInterface_Expecter) Watch(ctx interface{}, opts interface{}) *CSINodeInterface_Watch_Call {
	return &CSINodeInterface_Watch_Call{Call: _e.mock.On("Watch", ctx, opts)}
}

func (_c *CSINodeInterface_Watch_Call) Run(run func(ctx context.Context, opts v1.ListOptions)) *CSINodeInterface_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.ListOptions))
	})
	return _c
}

func (_c *CSINodeInterface_Watch_Call) Return(_a0 watch.Interface, _a1 error) *CSINodeInterface_Watch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CSINodeInterface_Watch_Call) RunAndReturn(run func(context.Context, v1.ListOptions) (watch.Interface, error)) *CSINodeInterface_Watch_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewCSINodeInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewCSINodeInterface creates a new instance of CSINodeInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCSINodeInterface(t mockConstructorTestingTNewCSINodeInterface) *CSINodeInterface {
	mock := &CSINodeInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
