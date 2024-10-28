// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	tlb "github.com/tonkeeper/tongo/tlb"

	ton "github.com/tonkeeper/tongo/ton"
)

// LiteClient is an autogenerated mock type for the LiteClient type
type LiteClient struct {
	mock.Mock
}

// GetBlockHeader provides a mock function with given fields: ctx, blockID, mode
func (_m *LiteClient) GetBlockHeader(ctx context.Context, blockID ton.BlockIDExt, mode uint32) (tlb.BlockInfo, error) {
	ret := _m.Called(ctx, blockID, mode)

	if len(ret) == 0 {
		panic("no return value specified for GetBlockHeader")
	}

	var r0 tlb.BlockInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ton.BlockIDExt, uint32) (tlb.BlockInfo, error)); ok {
		return rf(ctx, blockID, mode)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ton.BlockIDExt, uint32) tlb.BlockInfo); ok {
		r0 = rf(ctx, blockID, mode)
	} else {
		r0 = ret.Get(0).(tlb.BlockInfo)
	}

	if rf, ok := ret.Get(1).(func(context.Context, ton.BlockIDExt, uint32) error); ok {
		r1 = rf(ctx, blockID, mode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFirstTransaction provides a mock function with given fields: ctx, id
func (_m *LiteClient) GetFirstTransaction(ctx context.Context, id ton.AccountID) (*ton.Transaction, int, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetFirstTransaction")
	}

	var r0 *ton.Transaction
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, ton.AccountID) (*ton.Transaction, int, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ton.AccountID) *ton.Transaction); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ton.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ton.AccountID) int); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, ton.AccountID) error); ok {
		r2 = rf(ctx, id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetTransactionsSince provides a mock function with given fields: ctx, acc, lt, bits
func (_m *LiteClient) GetTransactionsSince(ctx context.Context, acc ton.AccountID, lt uint64, bits ton.Bits256) ([]ton.Transaction, error) {
	ret := _m.Called(ctx, acc, lt, bits)

	if len(ret) == 0 {
		panic("no return value specified for GetTransactionsSince")
	}

	var r0 []ton.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ton.AccountID, uint64, ton.Bits256) ([]ton.Transaction, error)); ok {
		return rf(ctx, acc, lt, bits)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ton.AccountID, uint64, ton.Bits256) []ton.Transaction); ok {
		r0 = rf(ctx, acc, lt, bits)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ton.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ton.AccountID, uint64, ton.Bits256) error); ok {
		r1 = rf(ctx, acc, lt, bits)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewLiteClient creates a new instance of LiteClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLiteClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *LiteClient {
	mock := &LiteClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}