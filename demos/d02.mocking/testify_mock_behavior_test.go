package demos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestARunningCar(t *testing.T) {
    mockCar := new(TestifyMockCar)
    mockCar.On("IsRunning").Return(true)

	driver := &Driver{mockCar}
    assert.False(t, driver.CanStartDriving())
}
