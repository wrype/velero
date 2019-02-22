/*
Copyright the Heptio Ark contributors.

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
	arkv1 "github.com/heptio/velero/pkg/apis/ark/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServerStatusRequests implements ServerStatusRequestInterface
type FakeServerStatusRequests struct {
	Fake *FakeArkV1
	ns   string
}

var serverstatusrequestsResource = schema.GroupVersionResource{Group: "ark.heptio.com", Version: "v1", Resource: "serverstatusrequests"}

var serverstatusrequestsKind = schema.GroupVersionKind{Group: "ark.heptio.com", Version: "v1", Kind: "ServerStatusRequest"}

// Get takes name of the serverStatusRequest, and returns the corresponding serverStatusRequest object, and an error if there is any.
func (c *FakeServerStatusRequests) Get(name string, options v1.GetOptions) (result *arkv1.ServerStatusRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(serverstatusrequestsResource, c.ns, name), &arkv1.ServerStatusRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*arkv1.ServerStatusRequest), err
}

// List takes label and field selectors, and returns the list of ServerStatusRequests that match those selectors.
func (c *FakeServerStatusRequests) List(opts v1.ListOptions) (result *arkv1.ServerStatusRequestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(serverstatusrequestsResource, serverstatusrequestsKind, c.ns, opts), &arkv1.ServerStatusRequestList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &arkv1.ServerStatusRequestList{ListMeta: obj.(*arkv1.ServerStatusRequestList).ListMeta}
	for _, item := range obj.(*arkv1.ServerStatusRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serverStatusRequests.
func (c *FakeServerStatusRequests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(serverstatusrequestsResource, c.ns, opts))

}

// Create takes the representation of a serverStatusRequest and creates it.  Returns the server's representation of the serverStatusRequest, and an error, if there is any.
func (c *FakeServerStatusRequests) Create(serverStatusRequest *arkv1.ServerStatusRequest) (result *arkv1.ServerStatusRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(serverstatusrequestsResource, c.ns, serverStatusRequest), &arkv1.ServerStatusRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*arkv1.ServerStatusRequest), err
}

// Update takes the representation of a serverStatusRequest and updates it. Returns the server's representation of the serverStatusRequest, and an error, if there is any.
func (c *FakeServerStatusRequests) Update(serverStatusRequest *arkv1.ServerStatusRequest) (result *arkv1.ServerStatusRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(serverstatusrequestsResource, c.ns, serverStatusRequest), &arkv1.ServerStatusRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*arkv1.ServerStatusRequest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServerStatusRequests) UpdateStatus(serverStatusRequest *arkv1.ServerStatusRequest) (*arkv1.ServerStatusRequest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(serverstatusrequestsResource, "status", c.ns, serverStatusRequest), &arkv1.ServerStatusRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*arkv1.ServerStatusRequest), err
}

// Delete takes name of the serverStatusRequest and deletes it. Returns an error if one occurs.
func (c *FakeServerStatusRequests) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(serverstatusrequestsResource, c.ns, name), &arkv1.ServerStatusRequest{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServerStatusRequests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(serverstatusrequestsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &arkv1.ServerStatusRequestList{})
	return err
}

// Patch applies the patch and returns the patched serverStatusRequest.
func (c *FakeServerStatusRequests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *arkv1.ServerStatusRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(serverstatusrequestsResource, c.ns, name, data, subresources...), &arkv1.ServerStatusRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*arkv1.ServerStatusRequest), err
}
