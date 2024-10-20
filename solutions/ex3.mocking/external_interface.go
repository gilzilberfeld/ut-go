package solutions_ex3

type IExternalDisplay interface {
	Show(text string)
	IsOn() bool
}
