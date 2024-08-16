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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "sigs.k8s.io/scheduler-plugins/pkg/apis/scheduling/v1alpha1"
)

// NetworkTopologyLister helps list NetworkTopologies.
// All objects returned here must be treated as read-only.
type NetworkTopologyLister interface {
	// List lists all NetworkTopologies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NetworkTopology, err error)
	// NetworkTopologies returns an object that can list and get NetworkTopologies.
	NetworkTopologies(namespace string) NetworkTopologyNamespaceLister
	NetworkTopologyListerExpansion
}

// networkTopologyLister implements the NetworkTopologyLister interface.
type networkTopologyLister struct {
	indexer cache.Indexer
}

// NewNetworkTopologyLister returns a new NetworkTopologyLister.
func NewNetworkTopologyLister(indexer cache.Indexer) NetworkTopologyLister {
	return &networkTopologyLister{indexer: indexer}
}

// List lists all NetworkTopologies in the indexer.
func (s *networkTopologyLister) List(selector labels.Selector) (ret []*v1alpha1.NetworkTopology, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetworkTopology))
	})
	return ret, err
}

// NetworkTopologies returns an object that can list and get NetworkTopologies.
func (s *networkTopologyLister) NetworkTopologies(namespace string) NetworkTopologyNamespaceLister {
	return networkTopologyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NetworkTopologyNamespaceLister helps list and get NetworkTopologies.
// All objects returned here must be treated as read-only.
type NetworkTopologyNamespaceLister interface {
	// List lists all NetworkTopologies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NetworkTopology, err error)
	// Get retrieves the NetworkTopology from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.NetworkTopology, error)
	NetworkTopologyNamespaceListerExpansion
}

// networkTopologyNamespaceLister implements the NetworkTopologyNamespaceLister
// interface.
type networkTopologyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NetworkTopologies in the indexer for a given namespace.
func (s networkTopologyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.NetworkTopology, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetworkTopology))
	})
	return ret, err
}

// Get retrieves the NetworkTopology from the indexer for a given namespace and name.
func (s networkTopologyNamespaceLister) Get(name string) (*v1alpha1.NetworkTopology, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("networktopology"), name)
	}
	return obj.(*v1alpha1.NetworkTopology), nil
}
