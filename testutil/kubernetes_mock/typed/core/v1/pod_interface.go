// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	corev1 "k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	mock "github.com/stretchr/testify/mock"

	policyv1 "k8s.io/api/policy/v1"

	rest "k8s.io/client-go/rest"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/client-go/applyconfigurations/core/v1"

	v1beta1 "k8s.io/api/policy/v1beta1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// PodInterface is an autogenerated mock type for the PodInterface type
type PodInterface struct {
	mock.Mock
}

type PodInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *PodInterface) EXPECT() *PodInterface_Expecter {
	return &PodInterface_Expecter{mock: &_m.Mock}
}

// Apply provides a mock function with given fields: ctx, pod, opts
func (_m *PodInterface) Apply(ctx context.Context, pod *v1.PodApplyConfiguration, opts metav1.ApplyOptions) (*corev1.Pod, error) {
	ret := _m.Called(ctx, pod, opts)

	var r0 *corev1.Pod
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.PodApplyConfiguration, metav1.ApplyOptions) (*corev1.Pod, error)); ok {
		return rf(ctx, pod, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.PodApplyConfiguration, metav1.ApplyOptions) *corev1.Pod); ok {
		r0 = rf(ctx, pod, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Pod)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.PodApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, pod, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PodInterface_Apply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Apply'
type PodInterface_Apply_Call struct {
	*mock.Call
}

// Apply is a helper method to define mock.On call
//   - ctx context.Context
//   - pod *v1.PodApplyConfiguration
//   - opts metav1.ApplyOptions
func (_e *PodInterface_Expecter) Apply(ctx interface{}, pod interface{}, opts interface{}) *PodInterface_Apply_Call {
	return &PodInterface_Apply_Call{Call: _e.mock.On("Apply", ctx, pod, opts)}
}

func (_c *PodInterface_Apply_Call) Run(run func(ctx context.Context, pod *v1.PodApplyConfiguration, opts metav1.ApplyOptions)) *PodInterface_Apply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1.PodApplyConfiguration), args[2].(metav1.ApplyOptions))
	})
	return _c
}

func (_c *PodInterface_Apply_Call) Return(result *corev1.Pod, err error) *PodInterface_Apply_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *PodInterface_Apply_Call) RunAndReturn(run func(context.Context, *v1.PodApplyConfiguration, metav1.ApplyOptions) (*corev1.Pod, error)) *PodInterface_Apply_Call {
	_c.Call.Return(run)
	return _c
}

// ApplyStatus provides a mock function with given fields: ctx, pod, opts
func (_m *PodInterface) ApplyStatus(ctx context.Context, pod *v1.PodApplyConfiguration, opts metav1.ApplyOptions) (*corev1.Pod, error) {
	ret := _m.Called(ctx, pod, opts)

	var r0 *corev1.Pod
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.PodApplyConfiguration, metav1.ApplyOptions) (*corev1.Pod, error)); ok {
		return rf(ctx, pod, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.PodApplyConfiguration, metav1.ApplyOptions) *corev1.Pod); ok {
		r0 = rf(ctx, pod, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Pod)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.PodApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, pod, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PodInterface_ApplyStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ApplyStatus'
type PodInterface_ApplyStatus_Call struct {
	*mock.Call
}

// ApplyStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - pod *v1.PodApplyConfiguration
//   - opts metav1.ApplyOptions
func (_e *PodInterface_Expecter) ApplyStatus(ctx interface{}, pod interface{}, opts interface{}) *PodInterface_ApplyStatus_Call {
	return &PodInterface_ApplyStatus_Call{Call: _e.mock.On("ApplyStatus", ctx, pod, opts)}
}

func (_c *PodInterface_ApplyStatus_Call) Run(run func(ctx context.Context, pod *v1.PodApplyConfiguration, opts metav1.ApplyOptions)) *PodInterface_ApplyStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1.PodApplyConfiguration), args[2].(metav1.ApplyOptions))
	})
	return _c
}

func (_c *PodInterface_ApplyStatus_Call) Return(result *corev1.Pod, err error) *PodInterface_ApplyStatus_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *PodInterface_ApplyStatus_Call) RunAndReturn(run func(context.Context, *v1.PodApplyConfiguration, metav1.ApplyOptions) (*corev1.Pod, error)) *PodInterface_ApplyStatus_Call {
	_c.Call.Return(run)
	return _c
}

// Bind provides a mock function with given fields: ctx, binding, opts
func (_m *PodInterface) Bind(ctx context.Context, binding *corev1.Binding, opts metav1.CreateOptions) error {
	ret := _m.Called(ctx, binding, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.Binding, metav1.CreateOptions) error); ok {
		r0 = rf(ctx, binding, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PodInterface_Bind_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Bind'
type PodInterface_Bind_Call struct {
	*mock.Call
}

// Bind is a helper method to define mock.On call
//   - ctx context.Context
//   - binding *corev1.Binding
//   - opts metav1.CreateOptions
func (_e *PodInterface_Expecter) Bind(ctx interface{}, binding interface{}, opts interface{}) *PodInterface_Bind_Call {
	return &PodInterface_Bind_Call{Call: _e.mock.On("Bind", ctx, binding, opts)}
}

func (_c *PodInterface_Bind_Call) Run(run func(ctx context.Context, binding *corev1.Binding, opts metav1.CreateOptions)) *PodInterface_Bind_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*corev1.Binding), args[2].(metav1.CreateOptions))
	})
	return _c
}

func (_c *PodInterface_Bind_Call) Return(_a0 error) *PodInterface_Bind_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PodInterface_Bind_Call) RunAndReturn(run func(context.Context, *corev1.Binding, metav1.CreateOptions) error) *PodInterface_Bind_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, pod, opts
func (_m *PodInterface) Create(ctx context.Context, pod *corev1.Pod, opts metav1.CreateOptions) (*corev1.Pod, error) {
	ret := _m.Called(ctx, pod, opts)

	var r0 *corev1.Pod
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.Pod, metav1.CreateOptions) (*corev1.Pod, error)); ok {
		return rf(ctx, pod, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.Pod, metav1.CreateOptions) *corev1.Pod); ok {
		r0 = rf(ctx, pod, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Pod)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *corev1.Pod, metav1.CreateOptions) error); ok {
		r1 = rf(ctx, pod, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PodInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type PodInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - pod *corev1.Pod
//   - opts metav1.CreateOptions
func (_e *PodInterface_Expecter) Create(ctx interface{}, pod interface{}, opts interface{}) *PodInterface_Create_Call {
	return &PodInterface_Create_Call{Call: _e.mock.On("Create", ctx, pod, opts)}
}

func (_c *PodInterface_Create_Call) Run(run func(ctx context.Context, pod *corev1.Pod, opts metav1.CreateOptions)) *PodInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*corev1.Pod), args[2].(metav1.CreateOptions))
	})
	return _c
}

func (_c *PodInterface_Create_Call) Return(_a0 *corev1.Pod, _a1 error) *PodInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PodInterface_Create_Call) RunAndReturn(run func(context.Context, *corev1.Pod, metav1.CreateOptions) (*corev1.Pod, error)) *PodInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *PodInterface) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PodInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type PodInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts metav1.DeleteOptions
func (_e *PodInterface_Expecter) Delete(ctx interface{}, name interface{}, opts interface{}) *PodInterface_Delete_Call {
	return &PodInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, name, opts)}
}

func (_c *PodInterface_Delete_Call) Run(run func(ctx context.Context, name string, opts metav1.DeleteOptions)) *PodInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(metav1.DeleteOptions))
	})
	return _c
}

func (_c *PodInterface_Delete_Call) Return(_a0 error) *PodInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PodInterface_Delete_Call) RunAndReturn(run func(context.Context, string, metav1.DeleteOptions) error) *PodInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteCollection provides a mock function with given fields: ctx, opts, listOpts
func (_m *PodInterface) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	ret := _m.Called(ctx, opts, listOpts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.DeleteOptions, metav1.ListOptions) error); ok {
		r0 = rf(ctx, opts, listOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PodInterface_DeleteCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCollection'
type PodInterface_DeleteCollection_Call struct {
	*mock.Call
}

// DeleteCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.DeleteOptions
//   - listOpts metav1.ListOptions
func (_e *PodInterface_Expecter) DeleteCollection(ctx interface{}, opts interface{}, listOpts interface{}) *PodInterface_DeleteCollection_Call {
	return &PodInterface_DeleteCollection_Call{Call: _e.mock.On("DeleteCollection", ctx, opts, listOpts)}
}

func (_c *PodInterface_DeleteCollection_Call) Run(run func(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions)) *PodInterface_DeleteCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.DeleteOptions), args[2].(metav1.ListOptions))
	})
	return _c
}

func (_c *PodInterface_DeleteCollection_Call) Return(_a0 error) *PodInterface_DeleteCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PodInterface_DeleteCollection_Call) RunAndReturn(run func(context.Context, metav1.DeleteOptions, metav1.ListOptions) error) *PodInterface_DeleteCollection_Call {
	_c.Call.Return(run)
	return _c
}

// Evict provides a mock function with given fields: ctx, eviction
func (_m *PodInterface) Evict(ctx context.Context, eviction *v1beta1.Eviction) error {
	ret := _m.Called(ctx, eviction)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.Eviction) error); ok {
		r0 = rf(ctx, eviction)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PodInterface_Evict_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Evict'
type PodInterface_Evict_Call struct {
	*mock.Call
}

// Evict is a helper method to define mock.On call
//   - ctx context.Context
//   - eviction *v1beta1.Eviction
func (_e *PodInterface_Expecter) Evict(ctx interface{}, eviction interface{}) *PodInterface_Evict_Call {
	return &PodInterface_Evict_Call{Call: _e.mock.On("Evict", ctx, eviction)}
}

func (_c *PodInterface_Evict_Call) Run(run func(ctx context.Context, eviction *v1beta1.Eviction)) *PodInterface_Evict_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1beta1.Eviction))
	})
	return _c
}

func (_c *PodInterface_Evict_Call) Return(_a0 error) *PodInterface_Evict_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PodInterface_Evict_Call) RunAndReturn(run func(context.Context, *v1beta1.Eviction) error) *PodInterface_Evict_Call {
	_c.Call.Return(run)
	return _c
}

// EvictV1 provides a mock function with given fields: ctx, eviction
func (_m *PodInterface) EvictV1(ctx context.Context, eviction *policyv1.Eviction) error {
	ret := _m.Called(ctx, eviction)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *policyv1.Eviction) error); ok {
		r0 = rf(ctx, eviction)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PodInterface_EvictV1_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EvictV1'
type PodInterface_EvictV1_Call struct {
	*mock.Call
}

// EvictV1 is a helper method to define mock.On call
//   - ctx context.Context
//   - eviction *policyv1.Eviction
func (_e *PodInterface_Expecter) EvictV1(ctx interface{}, eviction interface{}) *PodInterface_EvictV1_Call {
	return &PodInterface_EvictV1_Call{Call: _e.mock.On("EvictV1", ctx, eviction)}
}

func (_c *PodInterface_EvictV1_Call) Run(run func(ctx context.Context, eviction *policyv1.Eviction)) *PodInterface_EvictV1_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*policyv1.Eviction))
	})
	return _c
}

func (_c *PodInterface_EvictV1_Call) Return(_a0 error) *PodInterface_EvictV1_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PodInterface_EvictV1_Call) RunAndReturn(run func(context.Context, *policyv1.Eviction) error) *PodInterface_EvictV1_Call {
	_c.Call.Return(run)
	return _c
}

// EvictV1beta1 provides a mock function with given fields: ctx, eviction
func (_m *PodInterface) EvictV1beta1(ctx context.Context, eviction *v1beta1.Eviction) error {
	ret := _m.Called(ctx, eviction)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.Eviction) error); ok {
		r0 = rf(ctx, eviction)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PodInterface_EvictV1beta1_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EvictV1beta1'
type PodInterface_EvictV1beta1_Call struct {
	*mock.Call
}

// EvictV1beta1 is a helper method to define mock.On call
//   - ctx context.Context
//   - eviction *v1beta1.Eviction
func (_e *PodInterface_Expecter) EvictV1beta1(ctx interface{}, eviction interface{}) *PodInterface_EvictV1beta1_Call {
	return &PodInterface_EvictV1beta1_Call{Call: _e.mock.On("EvictV1beta1", ctx, eviction)}
}

func (_c *PodInterface_EvictV1beta1_Call) Run(run func(ctx context.Context, eviction *v1beta1.Eviction)) *PodInterface_EvictV1beta1_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1beta1.Eviction))
	})
	return _c
}

func (_c *PodInterface_EvictV1beta1_Call) Return(_a0 error) *PodInterface_EvictV1beta1_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PodInterface_EvictV1beta1_Call) RunAndReturn(run func(context.Context, *v1beta1.Eviction) error) *PodInterface_EvictV1beta1_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *PodInterface) Get(ctx context.Context, name string, opts metav1.GetOptions) (*corev1.Pod, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *corev1.Pod
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) (*corev1.Pod, error)); ok {
		return rf(ctx, name, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) *corev1.Pod); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Pod)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, metav1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PodInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type PodInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts metav1.GetOptions
func (_e *PodInterface_Expecter) Get(ctx interface{}, name interface{}, opts interface{}) *PodInterface_Get_Call {
	return &PodInterface_Get_Call{Call: _e.mock.On("Get", ctx, name, opts)}
}

func (_c *PodInterface_Get_Call) Run(run func(ctx context.Context, name string, opts metav1.GetOptions)) *PodInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(metav1.GetOptions))
	})
	return _c
}

func (_c *PodInterface_Get_Call) Return(_a0 *corev1.Pod, _a1 error) *PodInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PodInterface_Get_Call) RunAndReturn(run func(context.Context, string, metav1.GetOptions) (*corev1.Pod, error)) *PodInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetLogs provides a mock function with given fields: name, opts
func (_m *PodInterface) GetLogs(name string, opts *corev1.PodLogOptions) *rest.Request {
	ret := _m.Called(name, opts)

	var r0 *rest.Request
	if rf, ok := ret.Get(0).(func(string, *corev1.PodLogOptions) *rest.Request); ok {
		r0 = rf(name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rest.Request)
		}
	}

	return r0
}

// PodInterface_GetLogs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLogs'
type PodInterface_GetLogs_Call struct {
	*mock.Call
}

// GetLogs is a helper method to define mock.On call
//   - name string
//   - opts *corev1.PodLogOptions
func (_e *PodInterface_Expecter) GetLogs(name interface{}, opts interface{}) *PodInterface_GetLogs_Call {
	return &PodInterface_GetLogs_Call{Call: _e.mock.On("GetLogs", name, opts)}
}

func (_c *PodInterface_GetLogs_Call) Run(run func(name string, opts *corev1.PodLogOptions)) *PodInterface_GetLogs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*corev1.PodLogOptions))
	})
	return _c
}

func (_c *PodInterface_GetLogs_Call) Return(_a0 *rest.Request) *PodInterface_GetLogs_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PodInterface_GetLogs_Call) RunAndReturn(run func(string, *corev1.PodLogOptions) *rest.Request) *PodInterface_GetLogs_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, opts
func (_m *PodInterface) List(ctx context.Context, opts metav1.ListOptions) (*corev1.PodList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *corev1.PodList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) (*corev1.PodList, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) *corev1.PodList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.PodList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PodInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type PodInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.ListOptions
func (_e *PodInterface_Expecter) List(ctx interface{}, opts interface{}) *PodInterface_List_Call {
	return &PodInterface_List_Call{Call: _e.mock.On("List", ctx, opts)}
}

func (_c *PodInterface_List_Call) Run(run func(ctx context.Context, opts metav1.ListOptions)) *PodInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.ListOptions))
	})
	return _c
}

func (_c *PodInterface_List_Call) Return(_a0 *corev1.PodList, _a1 error) *PodInterface_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PodInterface_List_Call) RunAndReturn(run func(context.Context, metav1.ListOptions) (*corev1.PodList, error)) *PodInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// Patch provides a mock function with given fields: ctx, name, pt, data, opts, subresources
func (_m *PodInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*corev1.Pod, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *corev1.Pod
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) (*corev1.Pod, error)); ok {
		return rf(ctx, name, pt, data, opts, subresources...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) *corev1.Pod); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Pod)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PodInterface_Patch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Patch'
type PodInterface_Patch_Call struct {
	*mock.Call
}

// Patch is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - pt types.PatchType
//   - data []byte
//   - opts metav1.PatchOptions
//   - subresources ...string
func (_e *PodInterface_Expecter) Patch(ctx interface{}, name interface{}, pt interface{}, data interface{}, opts interface{}, subresources ...interface{}) *PodInterface_Patch_Call {
	return &PodInterface_Patch_Call{Call: _e.mock.On("Patch",
		append([]interface{}{ctx, name, pt, data, opts}, subresources...)...)}
}

func (_c *PodInterface_Patch_Call) Run(run func(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string)) *PodInterface_Patch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-5)
		for i, a := range args[5:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(args[0].(context.Context), args[1].(string), args[2].(types.PatchType), args[3].([]byte), args[4].(metav1.PatchOptions), variadicArgs...)
	})
	return _c
}

func (_c *PodInterface_Patch_Call) Return(result *corev1.Pod, err error) *PodInterface_Patch_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *PodInterface_Patch_Call) RunAndReturn(run func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) (*corev1.Pod, error)) *PodInterface_Patch_Call {
	_c.Call.Return(run)
	return _c
}

// ProxyGet provides a mock function with given fields: scheme, name, port, path, params
func (_m *PodInterface) ProxyGet(scheme string, name string, port string, path string, params map[string]string) rest.ResponseWrapper {
	ret := _m.Called(scheme, name, port, path, params)

	var r0 rest.ResponseWrapper
	if rf, ok := ret.Get(0).(func(string, string, string, string, map[string]string) rest.ResponseWrapper); ok {
		r0 = rf(scheme, name, port, path, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(rest.ResponseWrapper)
		}
	}

	return r0
}

// PodInterface_ProxyGet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProxyGet'
type PodInterface_ProxyGet_Call struct {
	*mock.Call
}

// ProxyGet is a helper method to define mock.On call
//   - scheme string
//   - name string
//   - port string
//   - path string
//   - params map[string]string
func (_e *PodInterface_Expecter) ProxyGet(scheme interface{}, name interface{}, port interface{}, path interface{}, params interface{}) *PodInterface_ProxyGet_Call {
	return &PodInterface_ProxyGet_Call{Call: _e.mock.On("ProxyGet", scheme, name, port, path, params)}
}

func (_c *PodInterface_ProxyGet_Call) Run(run func(scheme string, name string, port string, path string, params map[string]string)) *PodInterface_ProxyGet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string), args[3].(string), args[4].(map[string]string))
	})
	return _c
}

func (_c *PodInterface_ProxyGet_Call) Return(_a0 rest.ResponseWrapper) *PodInterface_ProxyGet_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PodInterface_ProxyGet_Call) RunAndReturn(run func(string, string, string, string, map[string]string) rest.ResponseWrapper) *PodInterface_ProxyGet_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, pod, opts
func (_m *PodInterface) Update(ctx context.Context, pod *corev1.Pod, opts metav1.UpdateOptions) (*corev1.Pod, error) {
	ret := _m.Called(ctx, pod, opts)

	var r0 *corev1.Pod
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.Pod, metav1.UpdateOptions) (*corev1.Pod, error)); ok {
		return rf(ctx, pod, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.Pod, metav1.UpdateOptions) *corev1.Pod); ok {
		r0 = rf(ctx, pod, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Pod)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *corev1.Pod, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, pod, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PodInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type PodInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - pod *corev1.Pod
//   - opts metav1.UpdateOptions
func (_e *PodInterface_Expecter) Update(ctx interface{}, pod interface{}, opts interface{}) *PodInterface_Update_Call {
	return &PodInterface_Update_Call{Call: _e.mock.On("Update", ctx, pod, opts)}
}

func (_c *PodInterface_Update_Call) Run(run func(ctx context.Context, pod *corev1.Pod, opts metav1.UpdateOptions)) *PodInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*corev1.Pod), args[2].(metav1.UpdateOptions))
	})
	return _c
}

func (_c *PodInterface_Update_Call) Return(_a0 *corev1.Pod, _a1 error) *PodInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PodInterface_Update_Call) RunAndReturn(run func(context.Context, *corev1.Pod, metav1.UpdateOptions) (*corev1.Pod, error)) *PodInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateEphemeralContainers provides a mock function with given fields: ctx, podName, pod, opts
func (_m *PodInterface) UpdateEphemeralContainers(ctx context.Context, podName string, pod *corev1.Pod, opts metav1.UpdateOptions) (*corev1.Pod, error) {
	ret := _m.Called(ctx, podName, pod, opts)

	var r0 *corev1.Pod
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *corev1.Pod, metav1.UpdateOptions) (*corev1.Pod, error)); ok {
		return rf(ctx, podName, pod, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *corev1.Pod, metav1.UpdateOptions) *corev1.Pod); ok {
		r0 = rf(ctx, podName, pod, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Pod)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *corev1.Pod, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, podName, pod, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PodInterface_UpdateEphemeralContainers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateEphemeralContainers'
type PodInterface_UpdateEphemeralContainers_Call struct {
	*mock.Call
}

// UpdateEphemeralContainers is a helper method to define mock.On call
//   - ctx context.Context
//   - podName string
//   - pod *corev1.Pod
//   - opts metav1.UpdateOptions
func (_e *PodInterface_Expecter) UpdateEphemeralContainers(ctx interface{}, podName interface{}, pod interface{}, opts interface{}) *PodInterface_UpdateEphemeralContainers_Call {
	return &PodInterface_UpdateEphemeralContainers_Call{Call: _e.mock.On("UpdateEphemeralContainers", ctx, podName, pod, opts)}
}

func (_c *PodInterface_UpdateEphemeralContainers_Call) Run(run func(ctx context.Context, podName string, pod *corev1.Pod, opts metav1.UpdateOptions)) *PodInterface_UpdateEphemeralContainers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*corev1.Pod), args[3].(metav1.UpdateOptions))
	})
	return _c
}

func (_c *PodInterface_UpdateEphemeralContainers_Call) Return(_a0 *corev1.Pod, _a1 error) *PodInterface_UpdateEphemeralContainers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PodInterface_UpdateEphemeralContainers_Call) RunAndReturn(run func(context.Context, string, *corev1.Pod, metav1.UpdateOptions) (*corev1.Pod, error)) *PodInterface_UpdateEphemeralContainers_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateStatus provides a mock function with given fields: ctx, pod, opts
func (_m *PodInterface) UpdateStatus(ctx context.Context, pod *corev1.Pod, opts metav1.UpdateOptions) (*corev1.Pod, error) {
	ret := _m.Called(ctx, pod, opts)

	var r0 *corev1.Pod
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.Pod, metav1.UpdateOptions) (*corev1.Pod, error)); ok {
		return rf(ctx, pod, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.Pod, metav1.UpdateOptions) *corev1.Pod); ok {
		r0 = rf(ctx, pod, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Pod)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *corev1.Pod, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, pod, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PodInterface_UpdateStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateStatus'
type PodInterface_UpdateStatus_Call struct {
	*mock.Call
}

// UpdateStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - pod *corev1.Pod
//   - opts metav1.UpdateOptions
func (_e *PodInterface_Expecter) UpdateStatus(ctx interface{}, pod interface{}, opts interface{}) *PodInterface_UpdateStatus_Call {
	return &PodInterface_UpdateStatus_Call{Call: _e.mock.On("UpdateStatus", ctx, pod, opts)}
}

func (_c *PodInterface_UpdateStatus_Call) Run(run func(ctx context.Context, pod *corev1.Pod, opts metav1.UpdateOptions)) *PodInterface_UpdateStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*corev1.Pod), args[2].(metav1.UpdateOptions))
	})
	return _c
}

func (_c *PodInterface_UpdateStatus_Call) Return(_a0 *corev1.Pod, _a1 error) *PodInterface_UpdateStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PodInterface_UpdateStatus_Call) RunAndReturn(run func(context.Context, *corev1.Pod, metav1.UpdateOptions) (*corev1.Pod, error)) *PodInterface_UpdateStatus_Call {
	_c.Call.Return(run)
	return _c
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *PodInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	ret := _m.Called(ctx, opts)

	var r0 watch.Interface
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) (watch.Interface, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) watch.Interface); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(watch.Interface)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PodInterface_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type PodInterface_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.ListOptions
func (_e *PodInterface_Expecter) Watch(ctx interface{}, opts interface{}) *PodInterface_Watch_Call {
	return &PodInterface_Watch_Call{Call: _e.mock.On("Watch", ctx, opts)}
}

func (_c *PodInterface_Watch_Call) Run(run func(ctx context.Context, opts metav1.ListOptions)) *PodInterface_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.ListOptions))
	})
	return _c
}

func (_c *PodInterface_Watch_Call) Return(_a0 watch.Interface, _a1 error) *PodInterface_Watch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PodInterface_Watch_Call) RunAndReturn(run func(context.Context, metav1.ListOptions) (watch.Interface, error)) *PodInterface_Watch_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewPodInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewPodInterface creates a new instance of PodInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPodInterface(t mockConstructorTestingTNewPodInterface) *PodInterface {
	mock := &PodInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
