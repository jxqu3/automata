package main

import (

	// Lib to replace err != nil

	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/checkm4ted/automata/internal/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const Width = 800
const Height = 600
const CellSize = 10

// Iterations Per Second
var Speed_IPSecond = uint(10)
var Paused = false

func main() {
	// Initial grid
	grid := utils.InitGrid(Width, Height, CellSize)

	for x := range grid {
		for y := range grid[x] {
			// Make 25% cells alive
			grid[x][y].Alive = rand.Intn(4) == 1
		}
	}

	// Initial next grid: this is a temporary grid to store the next state
	nGrid := utils.InitGrid(Width, Height, CellSize)

	// Init game
	game := utils.Game{
		Width:        Width,
		Height:       Height,
		CellSize:     CellSize,
		InitCellSize: CellSize,
		Grid: utils.Grid{
			Width:  Width / CellSize,
			Height: Height / CellSize,
			Cells:  grid,
		},
		NextGrid: utils.Grid{
			Width:  Width / CellSize,
			Height: Height / CellSize,
			Cells:  nGrid,
		},
	}

	rl.InitWindow(Width, Height, "CheckM4te Automata Template")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	go func() {
		for !rl.WindowShouldClose() {

			if !Paused {
				game.Update()
				time.Sleep(time.Duration(1000/Speed_IPSecond) * time.Millisecond)
			}
		}
	}()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		game.Draw()
		rl.DrawText(fmt.Sprint("Iterations Per Second: ", Speed_IPSecond), 10, 10, 20, color.RGBA{255, 0, 255, 255})

		// Pause with Space
		if rl.IsKeyPressed(rl.KeySpace) {
			Paused = !Paused
		}

		// IPS & Zoom controls
		if rl.GetMouseWheelMove() != 0 {
			if rl.IsKeyDown(rl.KeyLeftControl) {
				if game.CellSize+int(rl.GetMouseWheelMove()) > 0 {
					game.CellSize += int(rl.GetMouseWheelMove())
				}
			} else if Speed_IPSecond+uint(rl.GetMouseWheelMove()) > 0 {
				Speed_IPSecond += uint(rl.GetMouseWheelMove())
			}
		}

		rl.EndDrawing()
	}
}
