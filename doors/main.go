package main

import (
	"doors/door" // import the package relative to the module name
	"fmt"
)

func main() {
	var doors [100]door.Door
	for count := 1; count <= 100; count++ {
		doors[count-1] = door.GetNewDoor()
	}

	for pass := 1; pass <= 100; pass++ {
		for d := 1; d <= 100; d++ {
			if d%pass == 0 {
				// toggle the door
				door.Toggle(&(doors[d-1]))
			}
		}

		fmt.Println("Pass:", pass, "Doors:", doors)
	}
}
