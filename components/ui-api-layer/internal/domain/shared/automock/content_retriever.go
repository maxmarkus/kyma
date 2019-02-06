// Code generated by mockery v1.0.0
package automock

import mock "github.com/stretchr/testify/mock"
import shared "github.com/kyma-project/kyma/components/ui-api-layer/internal/domain/shared"

// ContentRetriever is an autogenerated mock type for the ContentRetriever type
type ContentRetriever struct {
	mock.Mock
}

// ApiSpec provides a mock function with given fields:
func (_m *ContentRetriever) ApiSpec() shared.ApiSpecGetter {
	ret := _m.Called()

	var r0 shared.ApiSpecGetter
	if rf, ok := ret.Get(0).(func() shared.ApiSpecGetter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(shared.ApiSpecGetter)
		}
	}

	return r0
}

// AsyncApiSpec provides a mock function with given fields:
func (_m *ContentRetriever) AsyncApiSpec() shared.AsyncApiSpecGetter {
	ret := _m.Called()

	var r0 shared.AsyncApiSpecGetter
	if rf, ok := ret.Get(0).(func() shared.AsyncApiSpecGetter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(shared.AsyncApiSpecGetter)
		}
	}

	return r0
}

// Content provides a mock function with given fields:
func (_m *ContentRetriever) Content() shared.ContentGetter {
	ret := _m.Called()

	var r0 shared.ContentGetter
	if rf, ok := ret.Get(0).(func() shared.ContentGetter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(shared.ContentGetter)
		}
	}

	return r0
}
