// Copyright (c) 2022 Alibaba Group Holding Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1alpha1 "github.com/alibaba/higress/client/pkg/apis/extensions/v1alpha1"
	extensionsv1alpha1 "github.com/alibaba/higress/client/pkg/applyconfiguration/extensions/v1alpha1"
	scheme "github.com/alibaba/higress/client/pkg/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// WasmPluginsGetter has a method to return a WasmPluginInterface.
// A group's client should implement this interface.
type WasmPluginsGetter interface {
	WasmPlugins(namespace string) WasmPluginInterface
}

// WasmPluginInterface has methods to work with WasmPlugin resources.
type WasmPluginInterface interface {
	Create(ctx context.Context, wasmPlugin *v1alpha1.WasmPlugin, opts v1.CreateOptions) (*v1alpha1.WasmPlugin, error)
	Update(ctx context.Context, wasmPlugin *v1alpha1.WasmPlugin, opts v1.UpdateOptions) (*v1alpha1.WasmPlugin, error)
	UpdateStatus(ctx context.Context, wasmPlugin *v1alpha1.WasmPlugin, opts v1.UpdateOptions) (*v1alpha1.WasmPlugin, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.WasmPlugin, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.WasmPluginList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.WasmPlugin, err error)
	Apply(ctx context.Context, wasmPlugin *extensionsv1alpha1.WasmPluginApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.WasmPlugin, err error)
	ApplyStatus(ctx context.Context, wasmPlugin *extensionsv1alpha1.WasmPluginApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.WasmPlugin, err error)
	WasmPluginExpansion
}

// wasmPlugins implements WasmPluginInterface
type wasmPlugins struct {
	client rest.Interface
	ns     string
}

// newWasmPlugins returns a WasmPlugins
func newWasmPlugins(c *ExtensionsV1alpha1Client, namespace string) *wasmPlugins {
	return &wasmPlugins{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the wasmPlugin, and returns the corresponding wasmPlugin object, and an error if there is any.
func (c *wasmPlugins) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.WasmPlugin, err error) {
	result = &v1alpha1.WasmPlugin{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("wasmplugins").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of WasmPlugins that match those selectors.
func (c *wasmPlugins) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.WasmPluginList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.WasmPluginList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("wasmplugins").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested wasmPlugins.
func (c *wasmPlugins) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("wasmplugins").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a wasmPlugin and creates it.  Returns the server's representation of the wasmPlugin, and an error, if there is any.
func (c *wasmPlugins) Create(ctx context.Context, wasmPlugin *v1alpha1.WasmPlugin, opts v1.CreateOptions) (result *v1alpha1.WasmPlugin, err error) {
	result = &v1alpha1.WasmPlugin{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("wasmplugins").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(wasmPlugin).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a wasmPlugin and updates it. Returns the server's representation of the wasmPlugin, and an error, if there is any.
func (c *wasmPlugins) Update(ctx context.Context, wasmPlugin *v1alpha1.WasmPlugin, opts v1.UpdateOptions) (result *v1alpha1.WasmPlugin, err error) {
	result = &v1alpha1.WasmPlugin{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("wasmplugins").
		Name(wasmPlugin.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(wasmPlugin).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *wasmPlugins) UpdateStatus(ctx context.Context, wasmPlugin *v1alpha1.WasmPlugin, opts v1.UpdateOptions) (result *v1alpha1.WasmPlugin, err error) {
	result = &v1alpha1.WasmPlugin{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("wasmplugins").
		Name(wasmPlugin.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(wasmPlugin).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the wasmPlugin and deletes it. Returns an error if one occurs.
func (c *wasmPlugins) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("wasmplugins").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *wasmPlugins) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("wasmplugins").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched wasmPlugin.
func (c *wasmPlugins) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.WasmPlugin, err error) {
	result = &v1alpha1.WasmPlugin{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("wasmplugins").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied wasmPlugin.
func (c *wasmPlugins) Apply(ctx context.Context, wasmPlugin *extensionsv1alpha1.WasmPluginApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.WasmPlugin, err error) {
	if wasmPlugin == nil {
		return nil, fmt.Errorf("wasmPlugin provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(wasmPlugin)
	if err != nil {
		return nil, err
	}
	name := wasmPlugin.Name
	if name == nil {
		return nil, fmt.Errorf("wasmPlugin.Name must be provided to Apply")
	}
	result = &v1alpha1.WasmPlugin{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("wasmplugins").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *wasmPlugins) ApplyStatus(ctx context.Context, wasmPlugin *extensionsv1alpha1.WasmPluginApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.WasmPlugin, err error) {
	if wasmPlugin == nil {
		return nil, fmt.Errorf("wasmPlugin provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(wasmPlugin)
	if err != nil {
		return nil, err
	}

	name := wasmPlugin.Name
	if name == nil {
		return nil, fmt.Errorf("wasmPlugin.Name must be provided to Apply")
	}

	result = &v1alpha1.WasmPlugin{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("wasmplugins").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
