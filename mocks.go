package idealpostcodes

import (
	"github.com/stretchr/testify/mock"
)

// ClientMock is a mocked object implementing ServiceInterface
type ClientMock struct {
	mock.Mock
}

// GetPostcode just records the activity, and returns what the Mock object tells it to
func (_m *ClientMock) GetPostcode(postcode string) ([]*Address, error) {
	ret := _m.Called(postcode)

	var r0 []*Address
	if rf, ok := ret.Get(0).(func(string) []*Address); ok {
		r0 = rf(postcode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(postcode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
