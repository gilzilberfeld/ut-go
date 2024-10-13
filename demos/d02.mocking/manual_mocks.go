package demos

type MockRunningCar struct{}

func (m *MockRunningCar) IsRunning() bool {
	return true
}
func (m *MockRunningCar) SetAC(ac AirCondition) {
}

func (m *MockRunningCar) Start() {
}

type MockNotRunningCar struct{}

func (m *MockNotRunningCar) IsRunning() bool {
	return false
}

func (m *MockNotRunningCar) SetAC(ac AirCondition) {
}

func (m *MockNotRunningCar) Start() {
}
