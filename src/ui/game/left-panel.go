package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
	"main/config"
	"main/consts"
	"main/ui"
)

type LeftPanel struct {
	ui.DrawInterface
	window *ui.Window
}

func (m *LeftPanel) GetName() string {
	return consts.LeftPanel
}

func (m *LeftPanel) Draw() {
	rl.DrawRectangleV(
		rl.Vector2{X: float32(config.GlobalConfig.LeftPanelDim.OffsetX), Y: float32(config.GlobalConfig.LeftPanelDim.OffsetY)},
		rl.Vector2{X: float32(config.GlobalConfig.LeftPanelDim.Width), Y: float32(config.GlobalConfig.LeftPanelDim.Height)},
		color.RGBA{
			R: 0,
			G: 90,
			B: 0,
			A: 255,
		},
	)

	rl.DrawLine(
		int32(config.GlobalConfig.LeftPanelDim.Width),
		int32(config.GlobalConfig.LeftPanelDim.OffsetY),
		int32(config.GlobalConfig.LeftPanelDim.Width),
		int32(config.GlobalConfig.LeftPanelDim.Height),
		color.RGBA{
			R: 250,
			G: 197,
			B: 40,
			A: 255,
		},
	)
}

func (m *LeftPanel) SetWindow(window *ui.Window) {
	m.window = window
}

func (m *LeftPanel) RegisterListeners() {
}

func (m *LeftPanel) RemoveListeners() {
}
