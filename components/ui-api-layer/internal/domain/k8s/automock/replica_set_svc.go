// Code generated by mockery v1.0.0. DO NOT EDIT.
package automock

import mock "github.com/stretchr/testify/mock"
import pager "github.com/kyma-project/kyma/components/ui-api-layer/internal/pager"
import v1 "k8s.io/api/apps/v1"

// replicaSetSvc is an autogenerated mock type for the replicaSetSvc type
type replicaSetSvc struct {
	mock.Mock
}

// Delete provides a mock function with given fields: name, namespace
func (_m *replicaSetSvc) Delete(name string, namespace string) error {
	ret := _m.Called(name, namespace)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(name, namespace)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: name, namespace
func (_m *replicaSetSvc) Find(name string, namespace string) (*v1.ReplicaSet, error) {
	ret := _m.Called(name, namespace)

	var r0 *v1.ReplicaSet
	if rf, ok := ret.Get(0).(func(string, string) *v1.ReplicaSet); ok {
		r0 = rf(name, namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ReplicaSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(name, namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: namespace, pagingParams
func (_m *replicaSetSvc) List(namespace string, pagingParams pager.PagingParams) ([]*v1.ReplicaSet, error) {
	ret := _m.Called(namespace, pagingParams)

	var r0 []*v1.ReplicaSet
	if rf, ok := ret.Get(0).(func(string, pager.PagingParams) []*v1.ReplicaSet); ok {
		r0 = rf(namespace, pagingParams)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.ReplicaSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, pager.PagingParams) error); ok {
		r1 = rf(namespace, pagingParams)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: name, namespace, update
func (_m *replicaSetSvc) Update(name string, namespace string, update v1.ReplicaSet) (*v1.ReplicaSet, error) {
	ret := _m.Called(name, namespace, update)

	var r0 *v1.ReplicaSet
	if rf, ok := ret.Get(0).(func(string, string, v1.ReplicaSet) *v1.ReplicaSet); ok {
		r0 = rf(name, namespace, update)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ReplicaSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, v1.ReplicaSet) error); ok {
		r1 = rf(name, namespace, update)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
