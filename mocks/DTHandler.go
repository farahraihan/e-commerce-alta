// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// DTHandler is an autogenerated mock type for the DTHandler type
type DTHandler struct {
	mock.Mock
}

// AddToCart provides a mock function with given fields: _a0
func (_m *DTHandler) AddToCart(_a0 echo.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for AddToCart")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCart provides a mock function with given fields: _a0
func (_m *DTHandler) DeleteCart(_a0 echo.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCart")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllCart provides a mock function with given fields: _a0
func (_m *DTHandler) GetAllCart(_a0 echo.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetAllCart")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateCart provides a mock function with given fields: _a0
func (_m *DTHandler) UpdateCart(_a0 echo.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCart")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewDTHandler creates a new instance of DTHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDTHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *DTHandler {
	mock := &DTHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
