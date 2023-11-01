package utils

type Game struct {
	Width    int
	Height   int
	CellSize int
	Grid     Grid
}

type Grid struct {
	Width  int
	Height int
	Cells  [][]*Cell
}

func (g *Game) GetNeighbor(c *Cell, x int, y int) *Cell {
	cX := c.Position.X
	cY := c.Position.Y

	return g.GetCell(cX+x, cY+y)
}

func (g *Game) GetCell(x, y int) *Cell {
	y = (g.Grid.Height + y) % g.Grid.Height
	x = (g.Grid.Width + x) % g.Grid.Width
	return g.Grid.Cells[x][y]
}
func (g *Game) GetNumberAliveNeighbors(c *Cell) int {
	var neighbors int

	y := c.Position.Y
	x := c.Position.X

	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if i == y && j == x {
				continue
			}
			if g.GetCell(j, i).Alive {
				neighbors++
			}
		}
	}
	return neighbors
}
