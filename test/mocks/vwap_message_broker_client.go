// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	aggregates "github.com/westcottian/vwap-calculation-engine/internal/aggregates"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// VWAPMessageBrokerClient is an autogenerated mock type for the VWAPMessageBrokerClient type
type VWAPMessageBrokerClient struct {
	mock.Mock
}

// Send provides a mock function with given fields: aggregate
func (_m *VWAPMessageBrokerClient) Send(aggregate aggregates.VWAPAggregate) {
	_m.Called(aggregate)
}

// NewVWAPMessageBrokerClient creates a new instance of VWAPMessageBrokerClient. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewVWAPMessageBrokerClient(t testing.TB) *VWAPMessageBrokerClient {
	mock := &VWAPMessageBrokerClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
