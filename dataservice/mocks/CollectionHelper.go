// Code generated by mockery v2.4.0-beta. DO NOT EDIT.

package mocks

import (
	context "context"

	dataservice "github.com/buidl-labs/Demux/dataservice"
	mock "github.com/stretchr/testify/mock"

	mongo "go.mongodb.org/mongo-driver/mongo"
)

// CollectionHelper is an autogenerated mock type for the CollectionHelper type
type CollectionHelper struct {
	mock.Mock
}

// DeleteOne provides a mock function with given fields: ctx, filter
func (_m *CollectionHelper) DeleteOne(ctx context.Context, filter interface{}) (int64, error) {
	ret := _m.Called(ctx, filter)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) int64); ok {
		r0 = rf(ctx, filter)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Find provides a mock function with given fields: ctx, filter
func (_m *CollectionHelper) Find(ctx context.Context, filter interface{}) (*mongo.Cursor, error) {
	ret := _m.Called(ctx, filter)

	var r0 *mongo.Cursor
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) *mongo.Cursor); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.Cursor)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOne provides a mock function with given fields: _a0, _a1
func (_m *CollectionHelper) FindOne(_a0 context.Context, _a1 interface{}) dataservice.SingleResultHelper {
	ret := _m.Called(_a0, _a1)

	var r0 dataservice.SingleResultHelper
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) dataservice.SingleResultHelper); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dataservice.SingleResultHelper)
		}
	}

	return r0
}

// InsertOne provides a mock function with given fields: _a0, _a1
func (_m *CollectionHelper) InsertOne(_a0 context.Context, _a1 interface{}) (interface{}, error) {
	ret := _m.Called(_a0, _a1)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) interface{}); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateOne provides a mock function with given fields: _a0, _a1, _a2
func (_m *CollectionHelper) UpdateOne(_a0 context.Context, _a1 interface{}, _a2 interface{}) (int64, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}) int64); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, interface{}) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
