package main

import (
	"fmt"
	"math"
	"time"
)

// Turtle das turtle
type Turtle struct {
	Speed            float64
	TimeTraveled     time.Duration
	DistanceTraveled float64
	DistanceInt      int
}

func main() {
	speedFirst := float64(550)
	speedSecond := float64(700)
	drugaKreceNakon := 300
	first := Turtle{
		Speed: speedFirst / 3600,
	}
	second := Turtle{
		Speed: speedSecond / 3600,
	}
	for {
		goTurtleGo(&first)
		if first.DistanceTraveled < float64(drugaKreceNakon) {
			continue
		}

		goTurtleGo(&second)
		if first.DistanceInt == second.DistanceInt {
			fmt.Printf("Bingo!\n")
			fmt.Printf("First: %+v\n", first)
			fmt.Printf("Second: %+v\n", second)

			break
		}
	}

}

func goTurtleGo(turtle *Turtle) {
	turtle.DistanceTraveled = turtle.DistanceTraveled + turtle.Speed
	turtle.TimeTraveled = turtle.TimeTraveled + (1 * time.Second)
	turtle.DistanceInt = int(math.RoundToEven(turtle.DistanceTraveled))
}
