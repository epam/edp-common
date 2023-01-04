package client

import (
	"context"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/stretchr/testify/mock"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Client struct {
	client.Reader
	client.Writer
	client.StatusClient
	client.SubResourceClientConstructor
	mock.Mock
}

func (c *Client) Scheme() *runtime.Scheme {
	panic("not implemented yet")
}

func (c *Client) RESTMapper() meta.RESTMapper {
	panic("not implemented yet")
}

func (c *Client) SubResource(subResource string) client.SubResourceClient {
	panic("not implemented yet")
}

func (c *Client) Get(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
	called := c.Called(key, obj)
	parent, ok := called.Get(0).(client.Client)
	if ok {
		return parent.Get(ctx, key, obj, opts...)
	}

	return called.Error(0)
}

func (c *Client) List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error {
	called := c.Called(list, opts)
	parent, ok := called.Get(0).(client.Client)
	if ok {
		return parent.List(ctx, list, opts...)
	}

	return called.Error(0)
}

func (c *Client) Create(ctx context.Context, obj client.Object, opts ...client.CreateOption) error {
	called := c.Called(obj, opts)
	parent, ok := called.Get(0).(client.Client)
	if ok {
		return parent.Create(ctx, obj, opts...)
	}

	return called.Error(0)
}

func (c *Client) Delete(ctx context.Context, obj client.Object, opts ...client.DeleteOption) error {
	called := c.Called(obj, opts)
	parent, ok := called.Get(0).(client.Client)
	if ok {
		return parent.Delete(ctx, obj, opts...)
	}

	return called.Error(0)
}

func (c *Client) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	called := c.Called(obj, opts)
	parent, ok := called.Get(0).(client.Client)
	if ok {
		return parent.Update(ctx, obj, opts...)
	}

	return called.Error(0)
}

func (c *Client) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.PatchOption) error {
	called := c.Called(obj, opts)
	parent, ok := called.Get(0).(client.Client)
	if ok {
		return parent.Patch(ctx, obj, patch, opts...)
	}

	return called.Error(0)
}

func (c *Client) Status() client.SubResourceWriter {
	return c.Called().Get(0).(client.SubResourceWriter)
}
