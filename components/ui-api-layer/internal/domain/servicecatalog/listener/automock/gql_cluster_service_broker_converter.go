// Code generated by mockery v1.0.0
package automock

import gqlschema "github.com/kyma-project/kyma/components/ui-api-layer/internal/gqlschema"

import mock "github.com/stretchr/testify/mock"
import v1beta1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"

// gqlClusterServiceBrokerConverter is an autogenerated mock type for the gqlClusterServiceBrokerConverter type
type gqlClusterServiceBrokerConverter struct {
	mock.Mock
}

// ToGQL provides a mock function with given fields: in
func (_m *gqlClusterServiceBrokerConverter) ToGQL(in *v1beta1.ClusterServiceBroker) (*gqlschema.ClusterServiceBroker, error) {
	ret := _m.Called(in)

	var r0 *gqlschema.ClusterServiceBroker
	if rf, ok := ret.Get(0).(func(*v1beta1.ClusterServiceBroker) *gqlschema.ClusterServiceBroker); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlschema.ClusterServiceBroker)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1beta1.ClusterServiceBroker) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
