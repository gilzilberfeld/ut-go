package demos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockRunningCar struct{}

func (m *MockRunningCar) IsRunning() bool {
    return true
}

func (m *MockRunningCar) SetAC(ac AirCondition) {
}

func (m *MockRunningCar) Start() {
}

func TestCannotDriveARunningCar(t *testing.T) {
    mockCar := &MockRunningCar{}
    driver := &Driver{mockCar}
    assert.False(t, driver.CanStartDriving())
}

type MockNotRunningCar struct{}

func (m *MockNotRunningCar) IsRunning() bool {
    return false
}

func (m *MockNotRunningCar) SetAC(ac AirCondition) {
}

func (m *MockNotRunningCar) Start() {
}

func TestCanDriveNotRunningCar(t *testing.T) {
    mockCar := &MockNotRunningCar{}
    driver := &Driver{mockCar}
    assert.True(t, driver.CanStartDriving())
}
