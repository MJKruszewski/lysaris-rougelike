package config

import (
	"fmt"
)

type MapConfig struct {
	AppName string

	TileSize     int
	TileFontSize int

	WindowWidth  int
	WindowHeight int

	MapPanelDim    Dimension
	LeftPanelDim   Dimension
	BottomPanelDim Dimension
}

type Dimension struct {
	Width   int
	Height  int
	OffsetX int
	OffsetY int
}

var GlobalConfig = NewMapConfig()

func NewMapConfig() MapConfig {
	fmt.Println("Event | Creating config")

	width := 800
	height := 600

	return MapConfig{
		AppName:      "Lysaris - roguelike",
		TileSize:     20,
		TileFontSize: 10,
		WindowWidth:  width,
		WindowHeight: height,

		MapPanelDim: Dimension{
			Width:   int(float32(width) * 0.8),
			Height:  int(float32(height) * 0.7),
			OffsetX: int(float32(width) * 0.2),
			OffsetY: 0,
		},
		LeftPanelDim: Dimension{
			Width:   int(float32(width) * 0.2),
			Height:  height,
			OffsetX: 0,
			OffsetY: 0,
		},
		BottomPanelDim: Dimension{
			Width:   int(float32(width) * 0.8),
			Height:  int(float32(height) * 0.3),
			OffsetX: int(float32(width) * 0.2),
			OffsetY: int(float32(height) * 0.7),
		},
	}
}

func (c *MapConfig) RecalculateScreen(width int, height int) {
	fmt.Println("Event | Recalculate window sizes")

	c.WindowWidth = width
	c.WindowHeight = height

	c.MapPanelDim.Width = int(float32(width) * 0.8)
	c.MapPanelDim.Height = int(float32(height) * 0.7)
	c.MapPanelDim.OffsetX = int(float32(width) * 0.2)
	c.MapPanelDim.OffsetY = 0

	c.LeftPanelDim.Width = int(float32(width) * 0.2)
	c.LeftPanelDim.Height = height
	c.LeftPanelDim.OffsetX = 0
	c.LeftPanelDim.OffsetY = 0

	c.BottomPanelDim.Width = int(float32(width) * 0.8)
	c.BottomPanelDim.Height = int(float32(height) * 0.3)
	c.BottomPanelDim.OffsetX = int(float32(width) * 0.2)
	c.BottomPanelDim.OffsetY = int(float32(height) * 0.7)

	if c.MapPanelDim.OffsetX%c.TileSize != 0 {
		c.MapPanelDim.OffsetX = c.MapPanelDim.OffsetX - (c.MapPanelDim.OffsetX % c.TileSize) + c.TileSize
	}
	if c.MapPanelDim.Width%c.TileSize != 0 {
		c.MapPanelDim.Width = c.MapPanelDim.Width - (c.MapPanelDim.Width % c.TileSize) + c.TileSize
	}
	if c.MapPanelDim.Height%c.TileSize != 0 {
		c.MapPanelDim.Height = c.MapPanelDim.Height - (c.MapPanelDim.Height % c.TileSize) + c.TileSize
	}

	if c.LeftPanelDim.Width%c.TileSize != 0 {
		c.LeftPanelDim.Width = c.LeftPanelDim.Width - (c.LeftPanelDim.Width % c.TileSize) + c.TileSize
	}

	if c.BottomPanelDim.Width%c.TileSize != 0 {
		c.BottomPanelDim.Width = c.BottomPanelDim.Width - (c.BottomPanelDim.Width % c.TileSize) + c.TileSize
	}
	if c.BottomPanelDim.Height%c.TileSize != 0 {
		c.BottomPanelDim.Height = c.BottomPanelDim.Height - (c.BottomPanelDim.Height % c.TileSize) + c.TileSize
	}
	if c.BottomPanelDim.OffsetX%c.TileSize != 0 {
		c.BottomPanelDim.OffsetX = c.BottomPanelDim.OffsetX - (c.BottomPanelDim.OffsetX % c.TileSize) + c.TileSize
	}
	if c.BottomPanelDim.OffsetY%c.TileSize != 0 {
		c.BottomPanelDim.OffsetY = c.BottomPanelDim.OffsetY - (c.BottomPanelDim.OffsetY % c.TileSize) + c.TileSize
	}
}
