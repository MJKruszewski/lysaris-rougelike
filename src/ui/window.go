package ui

type Window struct {
	panels       []*DrawInterface
	currentLayer uint8
}

func NewWindow() Window {
	return Window{
		panels:       []*DrawInterface{},
		currentLayer: 0,
	}
}

func (w *Window) AddPanel(panel DrawInterface) {
	panel.SetWindow(w)
	w.panels = append(w.panels, &panel)
}

func (w *Window) GetPanels() []*DrawInterface {
	return w.panels
}

func (w *Window) LayerUp() {
	w.currentLayer++
}

func (w *Window) LayerDown() {
	w.currentLayer--
}
