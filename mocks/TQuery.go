// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	transactions "TokoGadget/internal/features/transactions"

	mock "github.com/stretchr/testify/mock"
)

// TQuery is an autogenerated mock type for the TQuery type
type TQuery struct {
	mock.Mock
}

// CheckPendingTransaction provides a mock function with given fields: _a0
func (_m *TQuery) CheckPendingTransaction(_a0 uint) (transactions.Transaction, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for CheckPendingTransaction")
	}

	var r0 transactions.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (transactions.Transaction, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(uint) transactions.Transaction); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(transactions.Transaction)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckStock provides a mock function with given fields: _a0
func (_m *TQuery) CheckStock(_a0 uint) ([]transactions.CheckStock, bool) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for CheckStock")
	}

	var r0 []transactions.CheckStock
	var r1 bool
	if rf, ok := ret.Get(0).(func(uint) ([]transactions.CheckStock, bool)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(uint) []transactions.CheckStock); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]transactions.CheckStock)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) bool); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Checkout provides a mock function with given fields: _a0
func (_m *TQuery) Checkout(_a0 uint) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Checkout")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateTransaction provides a mock function with given fields: _a0
func (_m *TQuery) CreateTransaction(_a0 uint) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for CreateTransaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTransaction provides a mock function with given fields: _a0
func (_m *TQuery) DeleteTransaction(_a0 uint) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTransaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllTransactions provides a mock function with given fields: _a0
func (_m *TQuery) GetAllTransactions(_a0 uint) ([]transactions.Transaction, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetAllTransactions")
	}

	var r0 []transactions.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) ([]transactions.Transaction, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(uint) []transactions.Transaction); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]transactions.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPaymentDetails provides a mock function with given fields: _a0
func (_m *TQuery) GetPaymentDetails(_a0 uint) transactions.PaymentDetails {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetPaymentDetails")
	}

	var r0 transactions.PaymentDetails
	if rf, ok := ret.Get(0).(func(uint) transactions.PaymentDetails); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(transactions.PaymentDetails)
	}

	return r0
}

// GetTransaction provides a mock function with given fields: _a0
func (_m *TQuery) GetTransaction(_a0 uint) (transactions.Transaction, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetTransaction")
	}

	var r0 transactions.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (transactions.Transaction, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(uint) transactions.Transaction); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(transactions.Transaction)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RevertStock provides a mock function with given fields: _a0
func (_m *TQuery) RevertStock(_a0 []transactions.CheckStock) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for RevertStock")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]transactions.CheckStock) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateStock provides a mock function with given fields: _a0
func (_m *TQuery) UpdateStock(_a0 []transactions.CheckStock) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for UpdateStock")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]transactions.CheckStock) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTQuery creates a new instance of TQuery. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTQuery(t interface {
	mock.TestingT
	Cleanup(func())
}) *TQuery {
	mock := &TQuery{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
