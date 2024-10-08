// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Logger is an autogenerated mock type for the Logger type
type Logger struct {
	mock.Mock
}

type Logger_Expecter struct {
	mock *mock.Mock
}

func (_m *Logger) EXPECT() *Logger_Expecter {
	return &Logger_Expecter{mock: &_m.Mock}
}

// Debugf provides a mock function with given fields: format, a
func (_m *Logger) Debugf(format string, a ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, a...)
	_m.Called(_ca...)
}

// Logger_Debugf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Debugf'
type Logger_Debugf_Call struct {
	*mock.Call
}

// Debugf is a helper method to define mock.On call
//   - format string
//   - a ...interface{}
func (_e *Logger_Expecter) Debugf(format interface{}, a ...interface{}) *Logger_Debugf_Call {
	return &Logger_Debugf_Call{Call: _e.mock.On("Debugf",
		append([]interface{}{format}, a...)...)}
}

func (_c *Logger_Debugf_Call) Run(run func(format string, a ...interface{})) *Logger_Debugf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Logger_Debugf_Call) Return() *Logger_Debugf_Call {
	_c.Call.Return()
	return _c
}

func (_c *Logger_Debugf_Call) RunAndReturn(run func(string, ...interface{})) *Logger_Debugf_Call {
	_c.Call.Return(run)
	return _c
}

// Errorf provides a mock function with given fields: format, a
func (_m *Logger) Errorf(format string, a ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, a...)
	_m.Called(_ca...)
}

// Logger_Errorf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Errorf'
type Logger_Errorf_Call struct {
	*mock.Call
}

// Errorf is a helper method to define mock.On call
//   - format string
//   - a ...interface{}
func (_e *Logger_Expecter) Errorf(format interface{}, a ...interface{}) *Logger_Errorf_Call {
	return &Logger_Errorf_Call{Call: _e.mock.On("Errorf",
		append([]interface{}{format}, a...)...)}
}

func (_c *Logger_Errorf_Call) Run(run func(format string, a ...interface{})) *Logger_Errorf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Logger_Errorf_Call) Return() *Logger_Errorf_Call {
	_c.Call.Return()
	return _c
}

func (_c *Logger_Errorf_Call) RunAndReturn(run func(string, ...interface{})) *Logger_Errorf_Call {
	_c.Call.Return(run)
	return _c
}

// Fatalf provides a mock function with given fields: format, a
func (_m *Logger) Fatalf(format string, a ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, a...)
	_m.Called(_ca...)
}

// Logger_Fatalf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Fatalf'
type Logger_Fatalf_Call struct {
	*mock.Call
}

// Fatalf is a helper method to define mock.On call
//   - format string
//   - a ...interface{}
func (_e *Logger_Expecter) Fatalf(format interface{}, a ...interface{}) *Logger_Fatalf_Call {
	return &Logger_Fatalf_Call{Call: _e.mock.On("Fatalf",
		append([]interface{}{format}, a...)...)}
}

func (_c *Logger_Fatalf_Call) Run(run func(format string, a ...interface{})) *Logger_Fatalf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Logger_Fatalf_Call) Return() *Logger_Fatalf_Call {
	_c.Call.Return()
	return _c
}

func (_c *Logger_Fatalf_Call) RunAndReturn(run func(string, ...interface{})) *Logger_Fatalf_Call {
	_c.Call.Return(run)
	return _c
}

// Infof provides a mock function with given fields: format, a
func (_m *Logger) Infof(format string, a ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, a...)
	_m.Called(_ca...)
}

// Logger_Infof_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Infof'
type Logger_Infof_Call struct {
	*mock.Call
}

// Infof is a helper method to define mock.On call
//   - format string
//   - a ...interface{}
func (_e *Logger_Expecter) Infof(format interface{}, a ...interface{}) *Logger_Infof_Call {
	return &Logger_Infof_Call{Call: _e.mock.On("Infof",
		append([]interface{}{format}, a...)...)}
}

func (_c *Logger_Infof_Call) Run(run func(format string, a ...interface{})) *Logger_Infof_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Logger_Infof_Call) Return() *Logger_Infof_Call {
	_c.Call.Return()
	return _c
}

func (_c *Logger_Infof_Call) RunAndReturn(run func(string, ...interface{})) *Logger_Infof_Call {
	_c.Call.Return(run)
	return _c
}

// Warningf provides a mock function with given fields: format, a
func (_m *Logger) Warningf(format string, a ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, a...)
	_m.Called(_ca...)
}

// Logger_Warningf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Warningf'
type Logger_Warningf_Call struct {
	*mock.Call
}

// Warningf is a helper method to define mock.On call
//   - format string
//   - a ...interface{}
func (_e *Logger_Expecter) Warningf(format interface{}, a ...interface{}) *Logger_Warningf_Call {
	return &Logger_Warningf_Call{Call: _e.mock.On("Warningf",
		append([]interface{}{format}, a...)...)}
}

func (_c *Logger_Warningf_Call) Run(run func(format string, a ...interface{})) *Logger_Warningf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Logger_Warningf_Call) Return() *Logger_Warningf_Call {
	_c.Call.Return()
	return _c
}

func (_c *Logger_Warningf_Call) RunAndReturn(run func(string, ...interface{})) *Logger_Warningf_Call {
	_c.Call.Return(run)
	return _c
}

// NewLogger creates a new instance of Logger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLogger(t interface {
	mock.TestingT
	Cleanup(func())
}) *Logger {
	mock := &Logger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
