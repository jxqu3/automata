package utils

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) Draw() {
	cs := g.CellSize
	rl.ClearBackground(color.RGBA{0, 0, 0, 255})
	for x := range g.Grid.Cells {
		for y := range g.Grid.Cells[x] {
			c := g.GetCell(x, y) // get cell

			// Set cell color according to it's state
			c.Color = color.RGBA{0, 0, 0, 255}
			if c.Alive {
				c.Color = color.RGBA{50, uint8(c.Position.X), 0, 255}
			}

			// Draw cell
			rl.DrawRectangle(int32(x*cs), int32(y*cs), int32(cs), int32(cs), c.Color)
		}
	}
}
