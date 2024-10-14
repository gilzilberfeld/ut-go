package exercises

type IExternalDisplay interface {
	Show(text string)
	IsOn() bool
}
