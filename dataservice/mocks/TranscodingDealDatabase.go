// Code generated by mockery v2.4.0-beta. DO NOT EDIT.

package mocks

import (
	model "github.com/buidl-labs/Demux/model"
	mock "github.com/stretchr/testify/mock"
)

// TranscodingDealDatabase is an autogenerated mock type for the TranscodingDealDatabase type
type TranscodingDealDatabase struct {
	mock.Mock
}

// GetTranscodingDeal provides a mock function with given fields: _a0
func (_m *TranscodingDealDatabase) GetTranscodingDeal(_a0 string) (model.TranscodingDeal, error) {
	ret := _m.Called(_a0)

	var r0 model.TranscodingDeal
	if rf, ok := ret.Get(0).(func(string) model.TranscodingDeal); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(model.TranscodingDeal)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertTranscodingDeal provides a mock function with given fields: _a0
func (_m *TranscodingDealDatabase) InsertTranscodingDeal(_a0 model.TranscodingDeal) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.TranscodingDeal) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
