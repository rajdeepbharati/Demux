// Code generated by mockery v2.4.0-beta. DO NOT EDIT.

package mocks

import (
	"fmt"

	model "github.com/buidl-labs/Demux/model"
	mock "github.com/stretchr/testify/mock"
)

// AssetDatabase is an autogenerated mock type for the AssetDatabase type
type AssetDatabase struct {
	mock.Mock
}

// GetAsset provides a mock function with given fields: _a0
func (_m *AssetDatabase) GetAsset(_a0 string) (model.Asset, error) {
	asset := model.Asset{
		AssetID:         "1b2e976a-983d-4845-967a-f60b33c82869",
		AssetReady:      false,
		AssetStatusCode: 0,
		AssetStatus:     "asset created",
		AssetError:      false,
		CreatedAt:       1605030069,
		Thumbnail:       "",
	}

	if _a0 == asset.AssetID {
		return asset, nil
	}

	return model.Asset{}, fmt.Errorf("couldn't find asset")
}

// IfAssetExists provides a mock function with given fields: _a0
func (_m *AssetDatabase) IfAssetExists(_a0 string) bool {
	asset := model.Asset{
		AssetID:         "1b2e976a-983d-4845-967a-f60b33c82869",
		AssetReady:      false,
		AssetStatusCode: 0,
		AssetStatus:     "asset created",
		AssetError:      false,
		CreatedAt:       1605030069,
		Thumbnail:       "",
	}

	if _a0 == asset.AssetID {
		return true
	}

	return false
}

// InsertAsset provides a mock function with given fields: _a0
func (_m *AssetDatabase) InsertAsset(_a0 model.Asset) error {
	return nil
}

// UpdateAssetReady provides a mock function with given fields: _a0, _a1
func (_m *AssetDatabase) UpdateAssetReady(_a0 string, _a1 bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateAssetStatus provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *AssetDatabase) UpdateAssetStatus(_a0 string, _a1 int32, _a2 string, _a3 bool) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int32, string, bool) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateStreamURL provides a mock function with given fields: _a0, _a1
func (_m *AssetDatabase) UpdateStreamURL(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateThumbnail provides a mock function with given fields: _a0, _a1
func (_m *AssetDatabase) UpdateThumbnail(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
