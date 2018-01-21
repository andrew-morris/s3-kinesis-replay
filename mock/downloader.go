// Code generated by mockery v1.0.0
package mock

import io "io"
import mock "github.com/stretchr/testify/mock"
import s3 "github.com/aws/aws-sdk-go/service/s3"

import s3manager "github.com/aws/aws-sdk-go/service/s3/s3manager"

// Downloader is an autogenerated mock type for the Downloader type
type Downloader struct {
	mock.Mock
}

// Download provides a mock function with given fields: _a0, _a1, _a2
func (_m *Downloader) Download(_a0 io.WriterAt, _a1 *s3.GetObjectInput, _a2 ...func(*s3manager.Downloader)) (int64, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(io.WriterAt, *s3.GetObjectInput, ...func(*s3manager.Downloader)) int64); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(io.WriterAt, *s3.GetObjectInput, ...func(*s3manager.Downloader)) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
