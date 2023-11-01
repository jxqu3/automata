package utils

// What to do in the next iteration with cell x, y?
func (g *Game) Next(x, y int) bool {

	// Put your iteratin logic here
	return (x%2 == 0 && g.GetNeighbor(g.GetCell(x, y), -1, 0).Alive) || (y%3 == 0 && !g.GetNeighbor(g.GetCell(x, y), 0, -3).Alive)
}

// Next iteration of the world
func (g *Game) Update() {
	for x := range g.NextGrid.Cells {
		for y := range g.NextGrid.Cells[x] {
			g.NextGrid.Cells[x][y].Alive = g.Next(x, y)
		}
	}

	// Turn the grid into the new one and the clear the next
	g.Grid.Cells, g.NextGrid.Cells = g.NextGrid.Cells, InitGrid(g.Width, g.Height, g.InitCellSize)
}
