// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	archive "github.com/wal-g/wal-g/internal/databases/mongo/archive"
)

// MongoMetaProvider is an autogenerated mock type for the MongoMetaProvider type
type MongoMetaProvider struct {
	mock.Mock
}

// Finalize provides a mock function with given fields:
func (_m *MongoMetaProvider) Finalize() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Init provides a mock function with given fields:
func (_m *MongoMetaProvider) Init() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Meta provides a mock function with given fields:
func (_m *MongoMetaProvider) Meta() archive.MongoMeta {
	ret := _m.Called()

	var r0 archive.MongoMeta
	if rf, ok := ret.Get(0).(func() archive.MongoMeta); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(archive.MongoMeta)
	}

	return r0
}