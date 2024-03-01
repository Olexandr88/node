// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"
)

// EmissionAccountKeeper is an autogenerated mock type for the EmissionAccountKeeper type
type EmissionAccountKeeper struct {
	mock.Mock
}

// GetModuleAccount provides a mock function with given fields: ctx, moduleName
func (_m *EmissionAccountKeeper) GetModuleAccount(ctx types.Context, moduleName string) authtypes.ModuleAccountI {
	ret := _m.Called(ctx, moduleName)

	if len(ret) == 0 {
		panic("no return value specified for GetModuleAccount")
	}

	var r0 authtypes.ModuleAccountI
	if rf, ok := ret.Get(0).(func(types.Context, string) authtypes.ModuleAccountI); ok {
		r0 = rf(ctx, moduleName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(authtypes.ModuleAccountI)
		}
	}

	return r0
}

// NewEmissionAccountKeeper creates a new instance of EmissionAccountKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEmissionAccountKeeper(t interface {
	mock.TestingT
	Cleanup(func())
}) *EmissionAccountKeeper {
	mock := &EmissionAccountKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}