package demos

import "github.com/stretchr/testify/mock"

type TestifyMockCar struct {
	mock.Mock
}

func (m *TestifyMockCar) IsRunning() bool {
	// args sent to method
	// Expect a bool in the first position
	args := m.Called()
	return args.Bool(0)
}

func (m *TestifyMockCar) SetAC(ac AirCondition) {
	m.Called(ac)
}

func (m *TestifyMockCar) Start() {
	m.Called()
}
