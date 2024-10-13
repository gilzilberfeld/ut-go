package demos

type ACMode int

const (
	On ACMode = iota
	Off
)

type AirCondition struct {
	mode ACMode
}

func NewAirCondition(mode ACMode) *AirCondition {
	return &AirCondition{mode: mode}
}

func (ac *AirCondition) GetMode() ACMode {
	return ac.mode
}