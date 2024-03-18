// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	types "github.com/dydxprotocol/v4-chain/protocol/x/bridge/types"
)

// BridgeQueryClient is an autogenerated mock type for the BridgeQueryClient type
type BridgeQueryClient struct {
	mock.Mock
}

// AcknowledgedEventInfo provides a mock function with given fields: ctx, in, opts
func (_m *BridgeQueryClient) AcknowledgedEventInfo(ctx context.Context, in *types.QueryAcknowledgedEventInfoRequest, opts ...grpc.CallOption) (*types.QueryAcknowledgedEventInfoResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AcknowledgedEventInfo")
	}

	var r0 *types.QueryAcknowledgedEventInfoResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAcknowledgedEventInfoRequest, ...grpc.CallOption) (*types.QueryAcknowledgedEventInfoResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAcknowledgedEventInfoRequest, ...grpc.CallOption) *types.QueryAcknowledgedEventInfoResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryAcknowledgedEventInfoResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryAcknowledgedEventInfoRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DelayedCompleteBridgeMessages provides a mock function with given fields: ctx, in, opts
func (_m *BridgeQueryClient) DelayedCompleteBridgeMessages(ctx context.Context, in *types.QueryDelayedCompleteBridgeMessagesRequest, opts ...grpc.CallOption) (*types.QueryDelayedCompleteBridgeMessagesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DelayedCompleteBridgeMessages")
	}

	var r0 *types.QueryDelayedCompleteBridgeMessagesResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryDelayedCompleteBridgeMessagesRequest, ...grpc.CallOption) (*types.QueryDelayedCompleteBridgeMessagesResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryDelayedCompleteBridgeMessagesRequest, ...grpc.CallOption) *types.QueryDelayedCompleteBridgeMessagesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryDelayedCompleteBridgeMessagesResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryDelayedCompleteBridgeMessagesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventParams provides a mock function with given fields: ctx, in, opts
func (_m *BridgeQueryClient) EventParams(ctx context.Context, in *types.QueryEventParamsRequest, opts ...grpc.CallOption) (*types.QueryEventParamsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for EventParams")
	}

	var r0 *types.QueryEventParamsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryEventParamsRequest, ...grpc.CallOption) (*types.QueryEventParamsResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryEventParamsRequest, ...grpc.CallOption) *types.QueryEventParamsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryEventParamsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryEventParamsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProposeParams provides a mock function with given fields: ctx, in, opts
func (_m *BridgeQueryClient) ProposeParams(ctx context.Context, in *types.QueryProposeParamsRequest, opts ...grpc.CallOption) (*types.QueryProposeParamsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ProposeParams")
	}

	var r0 *types.QueryProposeParamsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryProposeParamsRequest, ...grpc.CallOption) (*types.QueryProposeParamsResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryProposeParamsRequest, ...grpc.CallOption) *types.QueryProposeParamsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryProposeParamsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryProposeParamsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RecognizedEventInfo provides a mock function with given fields: ctx, in, opts
func (_m *BridgeQueryClient) RecognizedEventInfo(ctx context.Context, in *types.QueryRecognizedEventInfoRequest, opts ...grpc.CallOption) (*types.QueryRecognizedEventInfoResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RecognizedEventInfo")
	}

	var r0 *types.QueryRecognizedEventInfoResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryRecognizedEventInfoRequest, ...grpc.CallOption) (*types.QueryRecognizedEventInfoResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryRecognizedEventInfoRequest, ...grpc.CallOption) *types.QueryRecognizedEventInfoResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryRecognizedEventInfoResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryRecognizedEventInfoRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SafetyParams provides a mock function with given fields: ctx, in, opts
func (_m *BridgeQueryClient) SafetyParams(ctx context.Context, in *types.QuerySafetyParamsRequest, opts ...grpc.CallOption) (*types.QuerySafetyParamsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SafetyParams")
	}

	var r0 *types.QuerySafetyParamsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QuerySafetyParamsRequest, ...grpc.CallOption) (*types.QuerySafetyParamsResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QuerySafetyParamsRequest, ...grpc.CallOption) *types.QuerySafetyParamsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QuerySafetyParamsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QuerySafetyParamsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBridgeQueryClient creates a new instance of BridgeQueryClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBridgeQueryClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *BridgeQueryClient {
	mock := &BridgeQueryClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
