// Code generated by mockery v1.0.0
package mocks

import flamingo "flamingo.me/flamingo/framework/flamingo"
import mock "github.com/stretchr/testify/mock"
import web "flamingo.me/flamingo/framework/web"

// Logger is an autogenerated mock type for the Logger type
type Logger struct {
	mock.Mock
}

// Debug provides a mock function with given fields: args
func (_m *Logger) Debug(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Error provides a mock function with given fields: args
func (_m *Logger) Error(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Fatal provides a mock function with given fields: args
func (_m *Logger) Fatal(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Flush provides a mock function with given fields:
func (_m *Logger) Flush() {
	_m.Called()
}

// Info provides a mock function with given fields: args
func (_m *Logger) Info(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Panic provides a mock function with given fields: args
func (_m *Logger) Panic(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Warn provides a mock function with given fields: args
func (_m *Logger) Warn(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// WithContext provides a mock function with given fields: ctx
func (_m *Logger) WithContext(ctx web.Context) flamingo.Logger {
	ret := _m.Called(ctx)

	var r0 flamingo.Logger
	if rf, ok := ret.Get(0).(func(web.Context) flamingo.Logger); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flamingo.Logger)
		}
	}

	return r0
}

// WithField provides a mock function with given fields: key, value
func (_m *Logger) WithField(key flamingo.LogKey, value interface{}) flamingo.Logger {
	ret := _m.Called(key, value)

	var r0 flamingo.Logger
	if rf, ok := ret.Get(0).(func(flamingo.LogKey, interface{}) flamingo.Logger); ok {
		r0 = rf(key, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flamingo.Logger)
		}
	}

	return r0
}

// WithFields provides a mock function with given fields: fields
func (_m *Logger) WithFields(fields map[flamingo.LogKey]interface{}) flamingo.Logger {
	ret := _m.Called(fields)

	var r0 flamingo.Logger
	if rf, ok := ret.Get(0).(func(map[flamingo.LogKey]interface{}) flamingo.Logger); ok {
		r0 = rf(fields)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flamingo.Logger)
		}
	}

	return r0
}
