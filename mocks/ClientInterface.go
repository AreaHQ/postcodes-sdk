package mocks

import "github.com/areatech/idealpostcodes-sdk"
import "github.com/stretchr/testify/mock"

// ClientInterface is a mocked object implementing ServiceInterface
type ClientInterface struct {
	mock.Mock
}

// GetPostcode just records the activity, and returns what the Mock object tells it to
func (_m *ClientInterface) GetPostcode(postcode string) ([]*idealpostcodes.Address, error) {
	ret := _m.Called(postcode)

	var r0 []*idealpostcodes.Address
	if rf, ok := ret.Get(0).(func(string) []*idealpostcodes.Address); ok {
		r0 = rf(postcode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*idealpostcodes.Address)
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
