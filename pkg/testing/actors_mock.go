// Code generated by mockery v1.0.0.

package testing

import actors "github.com/dapr/dapr/pkg/actors"
import mock "github.com/stretchr/testify/mock"

// MockActors is an autogenerated mock type for the MockActors type
type MockActors struct {
	mock.Mock
}

// Call provides a mock function with given fields: req
func (_m *MockActors) Call(req *actors.CallRequest) (*actors.CallResponse, error) {
	ret := _m.Called(req)

	var r0 *actors.CallResponse
	if rf, ok := ret.Get(0).(func(*actors.CallRequest) *actors.CallResponse); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*actors.CallResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*actors.CallRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateReminder provides a mock function with given fields: req
func (_m *MockActors) CreateReminder(req *actors.CreateReminderRequest) error {
	ret := _m.Called(req)

	var r0 error
	if rf, ok := ret.Get(0).(func(*actors.CreateReminderRequest) error); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsActorHosted provides a mock function with given fields: req
func (_m *MockActors) IsActorHosted(req *actors.ActorHostedRequest) bool {
	ret := _m.Called(req)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*actors.ActorHostedRequest) bool); ok {
		r0 = rf(req)
	} else {
		r0 = true
	}

	return r0
}

// CreateTimer provides a mock function with given fields: req
func (_m *MockActors) CreateTimer(req *actors.CreateTimerRequest) error {
	ret := _m.Called(req)

	var r0 error
	if rf, ok := ret.Get(0).(func(*actors.CreateTimerRequest) error); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteReminder provides a mock function with given fields: req
func (_m *MockActors) DeleteReminder(req *actors.DeleteReminderRequest) error {
	ret := _m.Called(req)

	var r0 error
	if rf, ok := ret.Get(0).(func(*actors.DeleteReminderRequest) error); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTimer provides a mock function with given fields: req
func (_m *MockActors) DeleteTimer(req *actors.DeleteTimerRequest) error {
	ret := _m.Called(req)

	var r0 error
	if rf, ok := ret.Get(0).(func(*actors.DeleteTimerRequest) error); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetState provides a mock function with given fields: req
func (_m *MockActors) GetState(req *actors.GetStateRequest) (*actors.StateResponse, error) {
	ret := _m.Called(req)

	var r0 *actors.StateResponse
	if rf, ok := ret.Get(0).(func(*actors.GetStateRequest) *actors.StateResponse); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*actors.StateResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*actors.GetStateRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Init provides a mock function with given fields:
func (_m *MockActors) Init() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveState provides a mock function with given fields: req
func (_m *MockActors) SaveState(req *actors.SaveStateRequest) error {
	ret := _m.Called(req)

	var r0 error
	if rf, ok := ret.Get(0).(func(*actors.SaveStateRequest) error); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteState provides a mock function with given fields: req
func (_m *MockActors) DeleteState(req *actors.DeleteStateRequest) error {
	ret := _m.Called(req)

	var r0 error
	if rf, ok := ret.Get(0).(func(*actors.DeleteStateRequest) error); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TransactionalStateOperation provides a mock function with given fields: req
func (_m *MockActors) TransactionalStateOperation(req *actors.TransactionalRequest) error {
	ret := _m.Called(req)

	var r0 error
	if rf, ok := ret.Get(0).(func(*actors.TransactionalRequest) error); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetReminder provides a mock function with given fields: req
func (_m *MockActors) GetReminder(req *actors.GetReminderRequest) (*actors.Reminder, error) {
	ret := _m.Called(req)

	var r0 error
	if rf, ok := ret.Get(0).(func(*actors.GetReminderRequest) error); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Error(0)
	}

	return nil, r0
}
