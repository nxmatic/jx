// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/jenkins-x/jx/v2/pkg/apis/jenkins.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SourceRepositoryLister helps list SourceRepositories.
type SourceRepositoryLister interface {
	// List lists all SourceRepositories in the indexer.
	List(selector labels.Selector) (ret []*v1.SourceRepository, err error)
	// SourceRepositories returns an object that can list and get SourceRepositories.
	SourceRepositories(namespace string) SourceRepositoryNamespaceLister
	SourceRepositoryListerExpansion
}

// sourceRepositoryLister implements the SourceRepositoryLister interface.
type sourceRepositoryLister struct {
	indexer cache.Indexer
}

// NewSourceRepositoryLister returns a new SourceRepositoryLister.
func NewSourceRepositoryLister(indexer cache.Indexer) SourceRepositoryLister {
	return &sourceRepositoryLister{indexer: indexer}
}

// List lists all SourceRepositories in the indexer.
func (s *sourceRepositoryLister) List(selector labels.Selector) (ret []*v1.SourceRepository, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.SourceRepository))
	})
	return ret, err
}

// SourceRepositories returns an object that can list and get SourceRepositories.
func (s *sourceRepositoryLister) SourceRepositories(namespace string) SourceRepositoryNamespaceLister {
	return sourceRepositoryNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SourceRepositoryNamespaceLister helps list and get SourceRepositories.
type SourceRepositoryNamespaceLister interface {
	// List lists all SourceRepositories in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.SourceRepository, err error)
	// Get retrieves the SourceRepository from the indexer for a given namespace and name.
	Get(name string) (*v1.SourceRepository, error)
	SourceRepositoryNamespaceListerExpansion
}

// sourceRepositoryNamespaceLister implements the SourceRepositoryNamespaceLister
// interface.
type sourceRepositoryNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SourceRepositories in the indexer for a given namespace.
func (s sourceRepositoryNamespaceLister) List(selector labels.Selector) (ret []*v1.SourceRepository, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.SourceRepository))
	})
	return ret, err
}

// Get retrieves the SourceRepository from the indexer for a given namespace and name.
func (s sourceRepositoryNamespaceLister) Get(name string) (*v1.SourceRepository, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("sourcerepository"), name)
	}
	return obj.(*v1.SourceRepository), nil
}