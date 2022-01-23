package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
	"main/config"
	"main/game/message"
	"main/ui"
)

type BottomPanel struct {
	ui.DrawInterface
	window     *ui.Window
	MessageLog *message.MessageLog
}

func (m *BottomPanel) Draw() {
	rl.DrawRectangleV(
		rl.Vector2{X: float32(config.GlobalConfig.BottomPanelDim.OffsetX), Y: float32(config.GlobalConfig.BottomPanelDim.OffsetY)},
		rl.Vector2{X: float32(config.GlobalConfig.BottomPanelDim.Width), Y: float32(config.GlobalConfig.BottomPanelDim.Height)},
		color.RGBA{
			R: 90,
			G: 0,
			B: 0,
			A: 255,
		},
	)

	i := 0
	for _, msg := range m.MessageLog.GetMessages() {
		rl.DrawText(
			msg.Time+": "+msg.Content,
			int32(config.GlobalConfig.BottomPanelDim.OffsetX+10),
			int32(config.GlobalConfig.BottomPanelDim.OffsetY+10+(config.GlobalConfig.TileSize*i)),
			int32(config.GlobalConfig.TileFontSize),
			black,
		)
		i++
	}

	rl.DrawLine(
		int32(config.GlobalConfig.BottomPanelDim.OffsetX),
		int32(config.GlobalConfig.BottomPanelDim.OffsetY),
		int32(config.GlobalConfig.BottomPanelDim.Width+config.GlobalConfig.BottomPanelDim.OffsetX),
		int32(config.GlobalConfig.BottomPanelDim.OffsetY),
		color.RGBA{
			R: 250,
			G: 197,
			B: 40,
			A: 255,
		},
	)
}

func (m *BottomPanel) SetWindow(window *ui.Window) {
	m.window = window
}
