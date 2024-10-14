// Code generated by mockery v2.46.1. DO NOT EDIT.

package mockport

import (
	context "context"
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// MockObjectStorage is an autogenerated mock type for the ObjectStorage type
type MockObjectStorage struct {
	mock.Mock
}

type MockObjectStorage_Expecter struct {
	mock *mock.Mock
}

func (_m *MockObjectStorage) EXPECT() *MockObjectStorage_Expecter {
	return &MockObjectStorage_Expecter{mock: &_m.Mock}
}

// DeleteObjects provides a mock function with given fields: ctx, keys
func (_m *MockObjectStorage) DeleteObjects(ctx context.Context, keys []string) error {
	ret := _m.Called(ctx, keys)

	if len(ret) == 0 {
		panic("no return value specified for DeleteObjects")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) error); ok {
		r0 = rf(ctx, keys)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockObjectStorage_DeleteObjects_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteObjects'
type MockObjectStorage_DeleteObjects_Call struct {
	*mock.Call
}

// DeleteObjects is a helper method to define mock.On call
//   - ctx context.Context
//   - keys []string
func (_e *MockObjectStorage_Expecter) DeleteObjects(ctx interface{}, keys interface{}) *MockObjectStorage_DeleteObjects_Call {
	return &MockObjectStorage_DeleteObjects_Call{Call: _e.mock.On("DeleteObjects", ctx, keys)}
}

func (_c *MockObjectStorage_DeleteObjects_Call) Run(run func(ctx context.Context, keys []string)) *MockObjectStorage_DeleteObjects_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string))
	})
	return _c
}

func (_c *MockObjectStorage_DeleteObjects_Call) Return(_a0 error) *MockObjectStorage_DeleteObjects_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockObjectStorage_DeleteObjects_Call) RunAndReturn(run func(context.Context, []string) error) *MockObjectStorage_DeleteObjects_Call {
	_c.Call.Return(run)
	return _c
}

// ListObjectKeysPrefix provides a mock function with given fields: ctx, prefix
func (_m *MockObjectStorage) ListObjectKeysPrefix(ctx context.Context, prefix string) ([]string, error) {
	ret := _m.Called(ctx, prefix)

	if len(ret) == 0 {
		panic("no return value specified for ListObjectKeysPrefix")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]string, error)); ok {
		return rf(ctx, prefix)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []string); ok {
		r0 = rf(ctx, prefix)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, prefix)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockObjectStorage_ListObjectKeysPrefix_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListObjectKeysPrefix'
type MockObjectStorage_ListObjectKeysPrefix_Call struct {
	*mock.Call
}

// ListObjectKeysPrefix is a helper method to define mock.On call
//   - ctx context.Context
//   - prefix string
func (_e *MockObjectStorage_Expecter) ListObjectKeysPrefix(ctx interface{}, prefix interface{}) *MockObjectStorage_ListObjectKeysPrefix_Call {
	return &MockObjectStorage_ListObjectKeysPrefix_Call{Call: _e.mock.On("ListObjectKeysPrefix", ctx, prefix)}
}

func (_c *MockObjectStorage_ListObjectKeysPrefix_Call) Run(run func(ctx context.Context, prefix string)) *MockObjectStorage_ListObjectKeysPrefix_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockObjectStorage_ListObjectKeysPrefix_Call) Return(_a0 []string, _a1 error) *MockObjectStorage_ListObjectKeysPrefix_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockObjectStorage_ListObjectKeysPrefix_Call) RunAndReturn(run func(context.Context, string) ([]string, error)) *MockObjectStorage_ListObjectKeysPrefix_Call {
	_c.Call.Return(run)
	return _c
}

// UploadObject provides a mock function with given fields: ctx, key, body
func (_m *MockObjectStorage) UploadObject(ctx context.Context, key string, body io.Reader) error {
	ret := _m.Called(ctx, key, body)

	if len(ret) == 0 {
		panic("no return value specified for UploadObject")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, io.Reader) error); ok {
		r0 = rf(ctx, key, body)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockObjectStorage_UploadObject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UploadObject'
type MockObjectStorage_UploadObject_Call struct {
	*mock.Call
}

// UploadObject is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - body io.Reader
func (_e *MockObjectStorage_Expecter) UploadObject(ctx interface{}, key interface{}, body interface{}) *MockObjectStorage_UploadObject_Call {
	return &MockObjectStorage_UploadObject_Call{Call: _e.mock.On("UploadObject", ctx, key, body)}
}

func (_c *MockObjectStorage_UploadObject_Call) Run(run func(ctx context.Context, key string, body io.Reader)) *MockObjectStorage_UploadObject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(io.Reader))
	})
	return _c
}

func (_c *MockObjectStorage_UploadObject_Call) Return(_a0 error) *MockObjectStorage_UploadObject_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockObjectStorage_UploadObject_Call) RunAndReturn(run func(context.Context, string, io.Reader) error) *MockObjectStorage_UploadObject_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockObjectStorage creates a new instance of MockObjectStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockObjectStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockObjectStorage {
	mock := &MockObjectStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
