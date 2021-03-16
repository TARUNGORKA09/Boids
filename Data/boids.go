package Data

import (
	"math/rand"
	"time"
)

type Boid struct {
	position Vector2D
	velocity Vector2D
	id       int
}

func (b *Boid) moveone() {
	b.position = b.position.Add(b.velocity)
	next := b.position.Add(b.velocity)
	if next.x > ScreenWidth || next.x < 0 {
		b.velocity = Vector2D{x: -b.velocity.x, y: b.velocity.y}
	}
	if next.y > Screenheight || next.y < 0 {
		b.velocity = Vector2D{x: b.velocity.x, y: -b.velocity.y}
	}

}
func (b *Boid) Start() {
	for {
		b.moveone()
		time.Sleep(5 * time.Millisecond)
	}
}

func CreateBoid(bid int) {
	b := Boid{
		position: Vector2D{x: rand.Float64() * ScreenWidth, y: rand.Float64() * Screenheight},
		velocity: Vector2D{x: (rand.Float64() * 2) - 1.0, y: (rand.Float64() * 2) - 1.0},
		id:       bid,
	}
	boids[bid] = &b
	go b.Start()
}
