package logic

import (
	"math/rand"

	"github.com/fatih/color"
)

type Direction int

const (
	Top Direction = iota
	Right
	Bottom
	Left
)

type Position struct {
	X         int       `json:"x"`
	Y         int       `json:"y"`
	Direction Direction `json:"direction"`
}

// 0: Black, 1: White
type Field = [][]bool

type World struct {
	Field Field    `json:"field"`
	Ant   Position `json:"ant"`
}

var (
	Black = color.New(color.BgBlack, color.FgWhite)
	White = color.New(color.BgWhite, color.FgBlack)
)

func negMod(x, mod int) int {
	// Go module does not support negative modulo
	return (x + mod) % mod
}

func VisualizeWorld(world *World) string {
	output := ""
	for y := 0; y < len(world.Field); y++ {
		for x := 0; x < len(world.Field[y]); x++ {
			var c *color.Color
			if world.Field[y][x] {
				c = White
			} else {
				c = Black
			}

			s := "  "
			if x == world.Ant.X && y == world.Ant.Y {
				if world.Ant.Direction == Top {
					s = "️↑ "
				} else if world.Ant.Direction == Right {
					s = "→ "
				}
				if world.Ant.Direction == Bottom {
					s = "↓ "
				}
				if world.Ant.Direction == Left {
					s = "← "
				}
			}
			output += c.Sprint(s)
		}

		output += "\n"
	}
	return output
}

func randomField(size int) Field {
	var field Field = make([][]bool, size)
	for y := 0; y < size; y++ {
		field[y] = make([]bool, size)
		for x := 0; x < size; x++ {
			field[y][x] = rand.Intn(2) == 1
		}
	}
	return field
}

func randomDirection() Direction {
	return Direction(rand.Intn(int(Left) - int(Top) + 1))
}

func randomAnt(size int) Position {
	return Position{
		X:         rand.Intn(size),
		Y:         rand.Intn(size),
		Direction: randomDirection(),
	}
}
func RandomWorld(size int) World {
	return World{
		Field: randomField(size),
		Ant:   randomAnt(size),
	}
}

func NextWorldState(world *World) {
	isOnWhite := world.Field[world.Ant.Y][world.Ant.X]

	// Turn left/right
	if isOnWhite {
		world.Ant.Direction += 1
	} else {
		world.Ant.Direction -= 1
	}
	// Wrap
	world.Ant.Direction = Direction(negMod(int(world.Ant.Direction), 4))

	// Flip
	world.Field[world.Ant.Y][world.Ant.X] = !isOnWhite

	// Move
	if world.Ant.Direction == Top {
		world.Ant.Y -= 1
	} else if world.Ant.Direction == Right {
		world.Ant.X += 1
	} else if world.Ant.Direction == Bottom {
		world.Ant.Y += 1
	} else if world.Ant.Direction == Left {
		world.Ant.X -= 1
	}
	// Wrap
	world.Ant.X = negMod(world.Ant.X, len(world.Field))
	world.Ant.Y = negMod(world.Ant.Y, len(world.Field[0]))
}
