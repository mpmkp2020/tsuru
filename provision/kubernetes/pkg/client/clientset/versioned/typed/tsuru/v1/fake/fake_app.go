// Copyright 2018 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	tsuruv1 "github.com/tsuru/tsuru/provision/kubernetes/pkg/apis/tsuru/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApps implements AppInterface
type FakeApps struct {
	Fake *FakeTsuruV1
	ns   string
}

var appsResource = schema.GroupVersionResource{Group: "tsuru.io", Version: "v1", Resource: "apps"}

var appsKind = schema.GroupVersionKind{Group: "tsuru.io", Version: "v1", Kind: "App"}

// Get takes name of the app, and returns the corresponding app object, and an error if there is any.
func (c *FakeApps) Get(ctx context.Context, name string, options v1.GetOptions) (result *tsuruv1.App, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(appsResource, c.ns, name), &tsuruv1.App{})

	if obj == nil {
		return nil, err
	}
	return obj.(*tsuruv1.App), err
}

// List takes label and field selectors, and returns the list of Apps that match those selectors.
func (c *FakeApps) List(ctx context.Context, opts v1.ListOptions) (result *tsuruv1.AppList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(appsResource, appsKind, c.ns, opts), &tsuruv1.AppList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &tsuruv1.AppList{ListMeta: obj.(*tsuruv1.AppList).ListMeta}
	for _, item := range obj.(*tsuruv1.AppList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apps.
func (c *FakeApps) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(appsResource, c.ns, opts))

}

// Create takes the representation of a app and creates it.  Returns the server's representation of the app, and an error, if there is any.
func (c *FakeApps) Create(ctx context.Context, app *tsuruv1.App, opts v1.CreateOptions) (result *tsuruv1.App, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(appsResource, c.ns, app), &tsuruv1.App{})

	if obj == nil {
		return nil, err
	}
	return obj.(*tsuruv1.App), err
}

// Update takes the representation of a app and updates it. Returns the server's representation of the app, and an error, if there is any.
func (c *FakeApps) Update(ctx context.Context, app *tsuruv1.App, opts v1.UpdateOptions) (result *tsuruv1.App, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(appsResource, c.ns, app), &tsuruv1.App{})

	if obj == nil {
		return nil, err
	}
	return obj.(*tsuruv1.App), err
}

// Delete takes name of the app and deletes it. Returns an error if one occurs.
func (c *FakeApps) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(appsResource, c.ns, name), &tsuruv1.App{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApps) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(appsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &tsuruv1.AppList{})
	return err
}

// Patch applies the patch and returns the patched app.
func (c *FakeApps) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *tsuruv1.App, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(appsResource, c.ns, name, pt, data, subresources...), &tsuruv1.App{})

	if obj == nil {
		return nil, err
	}
	return obj.(*tsuruv1.App), err
}
