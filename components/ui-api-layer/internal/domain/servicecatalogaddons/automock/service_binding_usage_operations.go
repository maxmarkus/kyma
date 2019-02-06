// Code generated by mockery v1.0.0
package automock

import mock "github.com/stretchr/testify/mock"
import resource "github.com/kyma-project/kyma/components/ui-api-layer/pkg/resource"

import v1alpha1 "github.com/kyma-project/kyma/components/binding-usage-controller/pkg/apis/servicecatalog/v1alpha1"

// serviceBindingUsageOperations is an autogenerated mock type for the serviceBindingUsageOperations type
type serviceBindingUsageOperations struct {
	mock.Mock
}

// Create provides a mock function with given fields: namespace, sb
func (_m *serviceBindingUsageOperations) Create(namespace string, sb *v1alpha1.ServiceBindingUsage) (*v1alpha1.ServiceBindingUsage, error) {
	ret := _m.Called(namespace, sb)

	var r0 *v1alpha1.ServiceBindingUsage
	if rf, ok := ret.Get(0).(func(string, *v1alpha1.ServiceBindingUsage) *v1alpha1.ServiceBindingUsage); ok {
		r0 = rf(namespace, sb)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.ServiceBindingUsage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *v1alpha1.ServiceBindingUsage) error); ok {
		r1 = rf(namespace, sb)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: namespace, name
func (_m *serviceBindingUsageOperations) Delete(namespace string, name string) error {
	ret := _m.Called(namespace, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(namespace, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: namespace, name
func (_m *serviceBindingUsageOperations) Find(namespace string, name string) (*v1alpha1.ServiceBindingUsage, error) {
	ret := _m.Called(namespace, name)

	var r0 *v1alpha1.ServiceBindingUsage
	if rf, ok := ret.Get(0).(func(string, string) *v1alpha1.ServiceBindingUsage); ok {
		r0 = rf(namespace, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.ServiceBindingUsage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(namespace, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListForServiceInstance provides a mock function with given fields: namespace, instanceName
func (_m *serviceBindingUsageOperations) ListForServiceInstance(namespace string, instanceName string) ([]*v1alpha1.ServiceBindingUsage, error) {
	ret := _m.Called(namespace, instanceName)

	var r0 []*v1alpha1.ServiceBindingUsage
	if rf, ok := ret.Get(0).(func(string, string) []*v1alpha1.ServiceBindingUsage); ok {
		r0 = rf(namespace, instanceName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1alpha1.ServiceBindingUsage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(namespace, instanceName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Subscribe provides a mock function with given fields: listener
func (_m *serviceBindingUsageOperations) Subscribe(listener resource.Listener) {
	_m.Called(listener)
}

// Unsubscribe provides a mock function with given fields: listener
func (_m *serviceBindingUsageOperations) Unsubscribe(listener resource.Listener) {
	_m.Called(listener)
}
