package main

type World [][]bool

func visualizeWorld(world World) string {
	output := ""
	for i := 0; i < len(world); i++ {
		for j := 0; j < len(world[i]); j++ {
			if world[i][j] {
				output += "*"
			} else {
				output += " "
			}
		}
		output += "\n"
	}
	return output
}

func randomWorld(size int) World {
	var world World = make([][]bool, size)
	return world
}
