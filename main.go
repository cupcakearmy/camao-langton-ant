package main

import (
	"fmt"
	"math/rand"
	"time"

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
	x         int
	y         int
	direction Direction
}

// 0: Black, 1: White
type Field = [][]bool

type World struct {
	field Field
	ant   Position
}

var (
	Black = color.New(color.BgBlack, color.FgWhite)
	White = color.New(color.BgWhite, color.FgBlack)
)

func visualizeWorld(world *World) string {
	output := ""
	for y := 0; y < len(world.field); y++ {
		for x := 0; x < len(world.field[y]); x++ {
			var c *color.Color
			if world.field[y][x] {
				c = White
			} else {
				c = Black
			}

			s := "  "
			if x == world.ant.x && y == world.ant.y {
				if world.ant.direction == Top {
					s = "️↑ "
				} else if world.ant.direction == Right {
					s = "→ "
				}
				if world.ant.direction == Bottom {
					s = "↓ "
				}
				if world.ant.direction == Left {
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
		x:         rand.Intn(size),
		y:         rand.Intn(size),
		direction: randomDirection(),
	}
}
func randomWorld(size int) World {
	return World{
		field: randomField(size),
		ant:   randomAnt(size),
	}
}

func negMod(x, mod int) int {
	// Go module does not support negative modulo
	return (x + mod) % mod
}

func nextState(world *World) {
	fmt.Println(world.ant.x, world.ant.y, world.ant.direction)
	isOnWhite := world.field[world.ant.y][world.ant.x]

	// Turn left/right
	fmt.Printf("Ant %d\n", world.ant.direction)
	if isOnWhite {
		world.ant.direction += 1
	} else {
		world.ant.direction -= 1
	}
	world.ant.direction = Direction(negMod(int(world.ant.direction), 4))

	// Flip
	world.field[world.ant.y][world.ant.x] = !isOnWhite

	// Move
	fmt.Printf("Ant %d\n", world.ant.direction)
	if world.ant.direction == Top {
		world.ant.y -= 1
	} else if world.ant.direction == Right {
		world.ant.x += 1
	} else if world.ant.direction == Bottom {
		world.ant.y += 1
	} else if world.ant.direction == Left {
		world.ant.x -= 1
	}

	// Wrap
	world.ant.x = negMod(world.ant.x, len(world.field))
	world.ant.y = negMod(world.ant.y, len(world.field[0]))

	fmt.Println(world.ant.x, world.ant.y, world.ant.direction)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	world := randomWorld(5)
	it := 0
	for {
		fmt.Printf("Iteration: %d\n", it)
		fmt.Println(visualizeWorld((&world)))
		time.Sleep(time.Second)
		nextState(&world)
		it += 1
	}
}
