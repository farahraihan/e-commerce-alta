// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	products "TokoGadget/internal/features/products"

	mock "github.com/stretchr/testify/mock"

	users "TokoGadget/internal/features/users"
)

// PQuery is an autogenerated mock type for the PQuery type
type PQuery struct {
	mock.Mock
}

// AddProduct provides a mock function with given fields: newProduct
func (_m *PQuery) AddProduct(newProduct products.Product) error {
	ret := _m.Called(newProduct)

	if len(ret) == 0 {
		panic("no return value specified for AddProduct")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(products.Product) error); ok {
		r0 = rf(newProduct)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CountAllProducts provides a mock function with given fields: term
func (_m *PQuery) CountAllProducts(term string) (int64, error) {
	ret := _m.Called(term)

	if len(ret) == 0 {
		panic("no return value specified for CountAllProducts")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (int64, error)); ok {
		return rf(term)
	}
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(term)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(term)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountProductsByUserID provides a mock function with given fields: userID, term
func (_m *PQuery) CountProductsByUserID(userID uint, term string) (int64, error) {
	ret := _m.Called(userID, term)

	if len(ret) == 0 {
		panic("no return value specified for CountProductsByUserID")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, string) (int64, error)); ok {
		return rf(userID, term)
	}
	if rf, ok := ret.Get(0).(func(uint, string) int64); ok {
		r0 = rf(userID, term)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(uint, string) error); ok {
		r1 = rf(userID, term)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteProduct provides a mock function with given fields: id
func (_m *PQuery) DeleteProduct(id uint) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteProduct")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllProducts provides a mock function with given fields: term, limit, offset
func (_m *PQuery) GetAllProducts(term string, limit int, offset int) ([]products.Product, error) {
	ret := _m.Called(term, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for GetAllProducts")
	}

	var r0 []products.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int, int) ([]products.Product, error)); ok {
		return rf(term, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(string, int, int) []products.Product); ok {
		r0 = rf(term, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]products.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int, int) error); ok {
		r1 = rf(term, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProductByID provides a mock function with given fields: id
func (_m *PQuery) GetProductByID(id uint) (products.Product, users.User, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetProductByID")
	}

	var r0 products.Product
	var r1 users.User
	var r2 error
	if rf, ok := ret.Get(0).(func(uint) (products.Product, users.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) products.Product); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(products.Product)
	}

	if rf, ok := ret.Get(1).(func(uint) users.User); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Get(1).(users.User)
	}

	if rf, ok := ret.Get(2).(func(uint) error); ok {
		r2 = rf(id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetProductsByUserID provides a mock function with given fields: user_id, term, limit, offset
func (_m *PQuery) GetProductsByUserID(user_id uint, term string, limit int, offset int) ([]products.Product, error) {
	ret := _m.Called(user_id, term, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for GetProductsByUserID")
	}

	var r0 []products.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, string, int, int) ([]products.Product, error)); ok {
		return rf(user_id, term, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(uint, string, int, int) []products.Product); ok {
		r0 = rf(user_id, term, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]products.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, string, int, int) error); ok {
		r1 = rf(user_id, term, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProductByID provides a mock function with given fields: id, updatedProduct
func (_m *PQuery) UpdateProductByID(id uint, updatedProduct products.Product) error {
	ret := _m.Called(id, updatedProduct)

	if len(ret) == 0 {
		panic("no return value specified for UpdateProductByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, products.Product) error); ok {
		r0 = rf(id, updatedProduct)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewPQuery creates a new instance of PQuery. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPQuery(t interface {
	mock.TestingT
	Cleanup(func())
}) *PQuery {
	mock := &PQuery{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
