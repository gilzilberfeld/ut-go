package demos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestCannotDriveARunningCar(t *testing.T) {
    mockCar := &MockRunningCar{}
    driver := &Driver{mockCar}
    assert.False(t, driver.CanStartDriving())
}


func TestCanDriveNotRunningCar(t *testing.T) {
    mockCar := &MockNotRunningCar{}
    driver := &Driver{mockCar}
    assert.True(t, driver.CanStartDriving())
}
