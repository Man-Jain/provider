/*
Copyright The Akash Network Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v2beta2

import (
	"context"
	"time"

	v2beta2 "github.com/akash-network/provider/pkg/apis/akash.network/v2beta2"
	scheme "github.com/akash-network/provider/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ProviderLeasedIPsGetter has a method to return a ProviderLeasedIPInterface.
// A group's client should implement this interface.
type ProviderLeasedIPsGetter interface {
	ProviderLeasedIPs(namespace string) ProviderLeasedIPInterface
}

// ProviderLeasedIPInterface has methods to work with ProviderLeasedIP resources.
type ProviderLeasedIPInterface interface {
	Create(ctx context.Context, providerLeasedIP *v2beta2.ProviderLeasedIP, opts v1.CreateOptions) (*v2beta2.ProviderLeasedIP, error)
	Update(ctx context.Context, providerLeasedIP *v2beta2.ProviderLeasedIP, opts v1.UpdateOptions) (*v2beta2.ProviderLeasedIP, error)
	UpdateStatus(ctx context.Context, providerLeasedIP *v2beta2.ProviderLeasedIP, opts v1.UpdateOptions) (*v2beta2.ProviderLeasedIP, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v2beta2.ProviderLeasedIP, error)
	List(ctx context.Context, opts v1.ListOptions) (*v2beta2.ProviderLeasedIPList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2beta2.ProviderLeasedIP, err error)
	ProviderLeasedIPExpansion
}

// providerLeasedIPs implements ProviderLeasedIPInterface
type providerLeasedIPs struct {
	client rest.Interface
	ns     string
}

// newProviderLeasedIPs returns a ProviderLeasedIPs
func newProviderLeasedIPs(c *AkashV2beta2Client, namespace string) *providerLeasedIPs {
	return &providerLeasedIPs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the providerLeasedIP, and returns the corresponding providerLeasedIP object, and an error if there is any.
func (c *providerLeasedIPs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2beta2.ProviderLeasedIP, err error) {
	result = &v2beta2.ProviderLeasedIP{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("providerleasedips").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ProviderLeasedIPs that match those selectors.
func (c *providerLeasedIPs) List(ctx context.Context, opts v1.ListOptions) (result *v2beta2.ProviderLeasedIPList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2beta2.ProviderLeasedIPList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("providerleasedips").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested providerLeasedIPs.
func (c *providerLeasedIPs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("providerleasedips").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a providerLeasedIP and creates it.  Returns the server's representation of the providerLeasedIP, and an error, if there is any.
func (c *providerLeasedIPs) Create(ctx context.Context, providerLeasedIP *v2beta2.ProviderLeasedIP, opts v1.CreateOptions) (result *v2beta2.ProviderLeasedIP, err error) {
	result = &v2beta2.ProviderLeasedIP{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("providerleasedips").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(providerLeasedIP).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a providerLeasedIP and updates it. Returns the server's representation of the providerLeasedIP, and an error, if there is any.
func (c *providerLeasedIPs) Update(ctx context.Context, providerLeasedIP *v2beta2.ProviderLeasedIP, opts v1.UpdateOptions) (result *v2beta2.ProviderLeasedIP, err error) {
	result = &v2beta2.ProviderLeasedIP{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("providerleasedips").
		Name(providerLeasedIP.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(providerLeasedIP).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *providerLeasedIPs) UpdateStatus(ctx context.Context, providerLeasedIP *v2beta2.ProviderLeasedIP, opts v1.UpdateOptions) (result *v2beta2.ProviderLeasedIP, err error) {
	result = &v2beta2.ProviderLeasedIP{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("providerleasedips").
		Name(providerLeasedIP.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(providerLeasedIP).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the providerLeasedIP and deletes it. Returns an error if one occurs.
func (c *providerLeasedIPs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("providerleasedips").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *providerLeasedIPs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("providerleasedips").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched providerLeasedIP.
func (c *providerLeasedIPs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2beta2.ProviderLeasedIP, err error) {
	result = &v2beta2.ProviderLeasedIP{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("providerleasedips").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
