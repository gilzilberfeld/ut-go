package demos

type Factorial struct {
}

func (f *Factorial) Calculate(number int) int {
	if number <= 1 {
		return number
	} else {
		return f.Calculate(number-1) * number
	}
}
