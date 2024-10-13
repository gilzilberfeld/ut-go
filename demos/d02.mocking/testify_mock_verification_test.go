package demos

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


func TestACIsSet_BeforeWeDrive(t *testing.T) {
    mockCar := new(TestifyMockCar)
    driver := &Driver{mockCar}

	// Expectations must be set for all calls
	// Otherwise, test will fail on unexpected call
	mockCar.On("SetAC", mock.AnythingOfType("AirCondition")).Return()
	mockCar.On("Start").Return()
	
    driver.Drive()
	
    mockCar.AssertCalled(t, "SetAC", mock.AnythingOfType("AirCondition"))
}

func TestACIsSetToOn_BeforeWeDrive(t *testing.T) {
	mockCar := new(TestifyMockCar)
	driver := &Driver{mockCar}
    sentACMode := Off
	
	// Collect the first argument of the SetAC call
	// Note, that the real method does not get called
    mockCar.On("SetAC", mock.AnythingOfType("AirCondition")).Run(func(args mock.Arguments) {
		ac := args.Get(0).(AirCondition)
        sentACMode = ac.mode
    })
	mockCar.On("Start").Return()
	
    driver.Drive()
	
    assert.Equal(t, On, sentACMode)
}
