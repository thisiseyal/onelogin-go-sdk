// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	http "net/http"

	models "github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

type IClient_Expecter struct {
	mock *mock.Mock
}

func (_m *IClient) EXPECT() *IClient_Expecter {
	return &IClient_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: path
func (_m *IClient) Delete(path *string) (*http.Response, error) {
	ret := _m.Called(path)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(*string) (*http.Response, error)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(*string) *http.Response); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(*string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IClient_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type IClient_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - path *string
func (_e *IClient_Expecter) Delete(path interface{}) *IClient_Delete_Call {
	return &IClient_Delete_Call{Call: _e.mock.On("Delete", path)}
}

func (_c *IClient_Delete_Call) Run(run func(path *string)) *IClient_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*string))
	})
	return _c
}

func (_c *IClient_Delete_Call) Return(_a0 *http.Response, _a1 error) *IClient_Delete_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IClient_Delete_Call) RunAndReturn(run func(*string) (*http.Response, error)) *IClient_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteWithBody provides a mock function with given fields: path, body
func (_m *IClient) DeleteWithBody(path *string, body interface{}) (*http.Response, error) {
	ret := _m.Called(path, body)

	if len(ret) == 0 {
		panic("no return value specified for DeleteWithBody")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(*string, interface{}) (*http.Response, error)); ok {
		return rf(path, body)
	}
	if rf, ok := ret.Get(0).(func(*string, interface{}) *http.Response); ok {
		r0 = rf(path, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(*string, interface{}) error); ok {
		r1 = rf(path, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IClient_DeleteWithBody_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteWithBody'
type IClient_DeleteWithBody_Call struct {
	*mock.Call
}

// DeleteWithBody is a helper method to define mock.On call
//   - path *string
//   - body interface{}
func (_e *IClient_Expecter) DeleteWithBody(path interface{}, body interface{}) *IClient_DeleteWithBody_Call {
	return &IClient_DeleteWithBody_Call{Call: _e.mock.On("DeleteWithBody", path, body)}
}

func (_c *IClient_DeleteWithBody_Call) Run(run func(path *string, body interface{})) *IClient_DeleteWithBody_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*string), args[1].(interface{}))
	})
	return _c
}

func (_c *IClient_DeleteWithBody_Call) Return(_a0 *http.Response, _a1 error) *IClient_DeleteWithBody_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IClient_DeleteWithBody_Call) RunAndReturn(run func(*string, interface{}) (*http.Response, error)) *IClient_DeleteWithBody_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: path, queryParams
func (_m *IClient) Get(path *string, queryParams models.Queryable) (*http.Response, error) {
	ret := _m.Called(path, queryParams)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(*string, models.Queryable) (*http.Response, error)); ok {
		return rf(path, queryParams)
	}
	if rf, ok := ret.Get(0).(func(*string, models.Queryable) *http.Response); ok {
		r0 = rf(path, queryParams)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(*string, models.Queryable) error); ok {
		r1 = rf(path, queryParams)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IClient_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type IClient_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - path *string
//   - queryParams models.Queryable
func (_e *IClient_Expecter) Get(path interface{}, queryParams interface{}) *IClient_Get_Call {
	return &IClient_Get_Call{Call: _e.mock.On("Get", path, queryParams)}
}

func (_c *IClient_Get_Call) Run(run func(path *string, queryParams models.Queryable)) *IClient_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*string), args[1].(models.Queryable))
	})
	return _c
}

func (_c *IClient_Get_Call) Return(_a0 *http.Response, _a1 error) *IClient_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IClient_Get_Call) RunAndReturn(run func(*string, models.Queryable) (*http.Response, error)) *IClient_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetAccountId provides a mock function with given fields:
func (_m *IClient) GetAccountId() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAccountId")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IClient_GetAccountId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAccountId'
type IClient_GetAccountId_Call struct {
	*mock.Call
}

// GetAccountId is a helper method to define mock.On call
func (_e *IClient_Expecter) GetAccountId() *IClient_GetAccountId_Call {
	return &IClient_GetAccountId_Call{Call: _e.mock.On("GetAccountId")}
}

func (_c *IClient_GetAccountId_Call) Run(run func()) *IClient_GetAccountId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *IClient_GetAccountId_Call) Return(_a0 string) *IClient_GetAccountId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IClient_GetAccountId_Call) RunAndReturn(run func() string) *IClient_GetAccountId_Call {
	_c.Call.Return(run)
	return _c
}

// GetToken provides a mock function with given fields:
func (_m *IClient) GetToken() (string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IClient_GetToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetToken'
type IClient_GetToken_Call struct {
	*mock.Call
}

// GetToken is a helper method to define mock.On call
func (_e *IClient_Expecter) GetToken() *IClient_GetToken_Call {
	return &IClient_GetToken_Call{Call: _e.mock.On("GetToken")}
}

func (_c *IClient_GetToken_Call) Run(run func()) *IClient_GetToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *IClient_GetToken_Call) Return(_a0 string, _a1 error) *IClient_GetToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IClient_GetToken_Call) RunAndReturn(run func() (string, error)) *IClient_GetToken_Call {
	_c.Call.Return(run)
	return _c
}

// Post provides a mock function with given fields: path, body
func (_m *IClient) Post(path *string, body interface{}) (*http.Response, error) {
	ret := _m.Called(path, body)

	if len(ret) == 0 {
		panic("no return value specified for Post")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(*string, interface{}) (*http.Response, error)); ok {
		return rf(path, body)
	}
	if rf, ok := ret.Get(0).(func(*string, interface{}) *http.Response); ok {
		r0 = rf(path, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(*string, interface{}) error); ok {
		r1 = rf(path, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IClient_Post_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Post'
type IClient_Post_Call struct {
	*mock.Call
}

// Post is a helper method to define mock.On call
//   - path *string
//   - body interface{}
func (_e *IClient_Expecter) Post(path interface{}, body interface{}) *IClient_Post_Call {
	return &IClient_Post_Call{Call: _e.mock.On("Post", path, body)}
}

func (_c *IClient_Post_Call) Run(run func(path *string, body interface{})) *IClient_Post_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*string), args[1].(interface{}))
	})
	return _c
}

func (_c *IClient_Post_Call) Return(_a0 *http.Response, _a1 error) *IClient_Post_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IClient_Post_Call) RunAndReturn(run func(*string, interface{}) (*http.Response, error)) *IClient_Post_Call {
	_c.Call.Return(run)
	return _c
}

// Put provides a mock function with given fields: path, body
func (_m *IClient) Put(path *string, body interface{}) (*http.Response, error) {
	ret := _m.Called(path, body)

	if len(ret) == 0 {
		panic("no return value specified for Put")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(*string, interface{}) (*http.Response, error)); ok {
		return rf(path, body)
	}
	if rf, ok := ret.Get(0).(func(*string, interface{}) *http.Response); ok {
		r0 = rf(path, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(*string, interface{}) error); ok {
		r1 = rf(path, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IClient_Put_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Put'
type IClient_Put_Call struct {
	*mock.Call
}

// Put is a helper method to define mock.On call
//   - path *string
//   - body interface{}
func (_e *IClient_Expecter) Put(path interface{}, body interface{}) *IClient_Put_Call {
	return &IClient_Put_Call{Call: _e.mock.On("Put", path, body)}
}

func (_c *IClient_Put_Call) Run(run func(path *string, body interface{})) *IClient_Put_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*string), args[1].(interface{}))
	})
	return _c
}

func (_c *IClient_Put_Call) Return(_a0 *http.Response, _a1 error) *IClient_Put_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IClient_Put_Call) RunAndReturn(run func(*string, interface{}) (*http.Response, error)) *IClient_Put_Call {
	_c.Call.Return(run)
	return _c
}

// NewIClient creates a new instance of IClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *IClient {
	mock := &IClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
