/*
Copyright The Kubernetes Authors.

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

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "sigs.k8s.io/scheduler-plugins/pkg/apis/scheduling/v1alpha1"
)

// FakeNetworkTopologies implements NetworkTopologyInterface
type FakeNetworkTopologies struct {
	Fake *FakeSchedulingV1alpha1
	ns   string
}

var networktopologiesResource = schema.GroupVersionResource{Group: "scheduling.sigs.k8s.io", Version: "v1alpha1", Resource: "networktopologies"}

var networktopologiesKind = schema.GroupVersionKind{Group: "scheduling.sigs.k8s.io", Version: "v1alpha1", Kind: "NetworkTopology"}

// Get takes name of the networkTopology, and returns the corresponding networkTopology object, and an error if there is any.
func (c *FakeNetworkTopologies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NetworkTopology, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networktopologiesResource, c.ns, name), &v1alpha1.NetworkTopology{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkTopology), err
}

// List takes label and field selectors, and returns the list of NetworkTopologies that match those selectors.
func (c *FakeNetworkTopologies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NetworkTopologyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networktopologiesResource, networktopologiesKind, c.ns, opts), &v1alpha1.NetworkTopologyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NetworkTopologyList{ListMeta: obj.(*v1alpha1.NetworkTopologyList).ListMeta}
	for _, item := range obj.(*v1alpha1.NetworkTopologyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkTopologies.
func (c *FakeNetworkTopologies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networktopologiesResource, c.ns, opts))

}

// Create takes the representation of a networkTopology and creates it.  Returns the server's representation of the networkTopology, and an error, if there is any.
func (c *FakeNetworkTopologies) Create(ctx context.Context, networkTopology *v1alpha1.NetworkTopology, opts v1.CreateOptions) (result *v1alpha1.NetworkTopology, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networktopologiesResource, c.ns, networkTopology), &v1alpha1.NetworkTopology{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkTopology), err
}

// Update takes the representation of a networkTopology and updates it. Returns the server's representation of the networkTopology, and an error, if there is any.
func (c *FakeNetworkTopologies) Update(ctx context.Context, networkTopology *v1alpha1.NetworkTopology, opts v1.UpdateOptions) (result *v1alpha1.NetworkTopology, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networktopologiesResource, c.ns, networkTopology), &v1alpha1.NetworkTopology{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkTopology), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkTopologies) UpdateStatus(ctx context.Context, networkTopology *v1alpha1.NetworkTopology, opts v1.UpdateOptions) (*v1alpha1.NetworkTopology, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networktopologiesResource, "status", c.ns, networkTopology), &v1alpha1.NetworkTopology{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkTopology), err
}

// Delete takes name of the networkTopology and deletes it. Returns an error if one occurs.
func (c *FakeNetworkTopologies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(networktopologiesResource, c.ns, name), &v1alpha1.NetworkTopology{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkTopologies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networktopologiesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NetworkTopologyList{})
	return err
}

// Patch applies the patch and returns the patched networkTopology.
func (c *FakeNetworkTopologies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NetworkTopology, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networktopologiesResource, c.ns, name, pt, data, subresources...), &v1alpha1.NetworkTopology{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkTopology), err
}
