// Code generated by mockery v2.4.0-beta. DO NOT EDIT.

package mocks

import (
	model "github.com/buidl-labs/Demux/model"
	mock "github.com/stretchr/testify/mock"
)

// StorageDealDatabase is an autogenerated mock type for the StorageDealDatabase type
type StorageDealDatabase struct {
	mock.Mock
}

// GetPendingDeals provides a mock function with given fields:
func (_m *StorageDealDatabase) GetPendingDeals() ([]model.StorageDeal, error) {
	ret := _m.Called()

	var r0 []model.StorageDeal
	if rf, ok := ret.Get(0).(func() []model.StorageDeal); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.StorageDeal)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStorageDeal provides a mock function with given fields: _a0
func (_m *StorageDealDatabase) GetStorageDeal(_a0 string) (model.StorageDeal, error) {
	ret := _m.Called(_a0)

	var r0 model.StorageDeal
	if rf, ok := ret.Get(0).(func(string) model.StorageDeal); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(model.StorageDeal)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStorageDealByCID provides a mock function with given fields: _a0
func (_m *StorageDealDatabase) GetStorageDealByCID(_a0 string) (model.StorageDeal, error) {
	ret := _m.Called(_a0)

	var r0 model.StorageDeal
	if rf, ok := ret.Get(0).(func(string) model.StorageDeal); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(model.StorageDeal)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertStorageDeal provides a mock function with given fields: _a0
func (_m *StorageDealDatabase) InsertStorageDeal(_a0 model.StorageDeal) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.StorageDeal) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateStorageDeal provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5
func (_m *StorageDealDatabase) UpdateStorageDeal(_a0 string, _a1 uint32, _a2 string, _a3 string, _a4 string, _a5 int64) error {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, uint32, string, string, string, int64) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
