package demos

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)
type MockCar struct {
    mock.Mock
}

func (m *MockCar) IsRunning() bool {
    // args sent to method
    // Expect a bool in the first position
    args := m.Called()
    return args.Bool(0)
}


func (m *MockCar) SetAC(ac AirCondition) {
    // m.Called(ac)
}

func (m *MockCar) Start() {
    // m.Called()
}

func (m *MockCar) GetAC() *AirCondition {
    // args := m.Called()
    // return args.Get(0).(AirCondition)
    return nil
}


func TestARunningCar(t *testing.T) {
    mockCar := new(MockCar)
    mockCar.On("IsRunning").Return(true)

	driver := &Driver{mockCar}
    assert.False(t, driver.CanStartDriving())
}
