package ui

type Window struct {
	panels       []*panel
	currentLayer uint8
}

type panel struct {
	item  *DrawInterface
	layer uint8
}

func NewWindow() Window {
	return Window{
		panels:       []*panel{},
		currentLayer: uint8(0),
	}
}

func (w *Window) AddPanel(element DrawInterface) {
	element.SetWindow(w)

	w.panels = append(w.panels, &panel{
		item:  &element,
		layer: w.currentLayer,
	})

	w.refreshListeners()
}

func (w *Window) RemovePanel(name string) {
	index := -1

	for i, p := range w.panels {
		if (*p.item).GetName() == name {
			index = i
			break
		}
	}

	if index == -1 {
		return
	}

	(*w.panels[index].item).RemoveListeners()
	w.panels = append(w.panels[:index], w.panels[index+1:]...)
}

func (w *Window) GetPanels() (res []*DrawInterface) {
	for _, p := range w.panels {
		res = append(res, p.item)
	}

	return res
}

func (w *Window) LayerUp() {
	w.currentLayer++

	w.refreshListeners()
}

func (w *Window) LayerDown() {
	w.currentLayer--

	w.refreshListeners()
}

func (w *Window) refreshListeners() {
	for _, p := range w.panels {
		(*p.item).RemoveListeners()
	}

	for _, p := range w.panels {
		if p.layer == w.currentLayer {
			(*p.item).RegisterListeners()
		}
	}
}
