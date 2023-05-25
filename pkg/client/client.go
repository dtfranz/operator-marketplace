package client

import (
	"context"

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ObjectKey identifies a Kubernetes Object.
type ObjectKey = types.NamespacedName

// Client is a wrapper around the raw kube client provided
// by operator-sdk. Using the wrapper facilitates mocking of client
// interactions with the cluster, while using fakeclient during unit testing.
// TODO(tflannag): Should this be removed entirely in favor of c-r's dynamic client?
type Client interface {
	Create(ctx context.Context, obj client.Object, opts ...client.CreateOption) error
	Get(ctx context.Context, key ObjectKey, objExisting client.Object, opts ...client.GetOption) error
	Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error
	Status() client.StatusWriter
	Delete(ctx context.Context, obj client.Object, opts ...client.DeleteOption) error
	List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error
}

// kubeClient is an implementation of the Client interface
type kubeClient struct {
	client client.Client
}

// NewClient returns a kubeClient that can perform
// create, get and update operations on a runtime object
func NewClient(client client.Client) Client {
	return &kubeClient{
		client: client,
	}
}

// Create creates a new runtime object in the cluster
func (h *kubeClient) Create(ctx context.Context, obj client.Object, opts ...client.CreateOption) error {
	return h.client.Create(ctx, obj, opts...)
}

// Get gets an existing runtime object from the cluster
func (h *kubeClient) Get(ctx context.Context, key ObjectKey, objExisting client.Object, opts ...client.GetOption) error {
	return h.client.Get(ctx, key, objExisting, opts...)
}

// Update updates an existing runtime object in the cluster
func (h *kubeClient) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	return h.client.Update(ctx, obj, opts...)
}

// Status updates an existing runtime object's status in the cluster
func (h *kubeClient) Status() client.StatusWriter {
	return h.client.Status()
}

// Delete deletes a new runtime object in the cluster
func (h *kubeClient) Delete(ctx context.Context, obj client.Object, opts ...client.DeleteOption) error {
	return h.client.Delete(ctx, obj, opts...)
}

// List lists runtime objects in the cluster
func (h *kubeClient) List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error {
	return h.client.List(ctx, list, opts...)
}
