package ui

type DrawInterface interface {
	Draw()
	SetWindow(window *Window)
	RegisterListeners()
	RemoveListeners()
	GetName() string
}
