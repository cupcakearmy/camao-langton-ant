package main

import (
	"fmt"
	"langton/logic"
	"math/rand"
	"time"
)

func main() {
	// go server.Init()
	rand.Seed(time.Now().UTC().UnixNano())
	world := logic.RandomWorld(10)
	it := 0
	for {
		fmt.Printf("Iteration: %d\n", it)
		fmt.Println(logic.VisualizeWorld((&world)))
		time.Sleep(time.Second)
		logic.NextWorldState(&world)
		it += 1
	}
}
