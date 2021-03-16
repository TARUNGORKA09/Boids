package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

const (
	ScreenWidth, Screenheight = 640, 360
	boidCount                 = 500
)

var boids [boidCount]*Boid

var (
	green = color.RGBA{10, 255, 50, 255}
)

func Update(screen *ebiten.Image) error {
	if !ebiten.IsDrawingSkipped() {
		for _, boid := range boids {
			screen.Set(int(boid.position.x+1), int(boid.position.y), green)
			screen.Set(int(boid.position.x-1), int(boid.position.y), green)
			screen.Set(int(boid.position.x+1), int(boid.position.y-1), green)
			screen.Set(int(boid.position.x+1), int(boid.position.y+1), green)

		}
	}
	return nil
}

func main() {
	for i := 0; i < boidCount; i++ {
		CreateBoid(i)
	}
	if err := ebiten.Run(Update, ScreenWidth, Screenheight, 2, "Boids in a box"); err != nil {
		log.Fatal(err)
	}
}
