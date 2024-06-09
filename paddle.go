package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type paddle struct {
	pos
	width  float32
	height float32
	speed  float32
	score  int
	color  color
}

func (paddle *paddle) draw() {
	startX := int(paddle.x - paddle.width/2)
	startY := int(paddle.y - paddle.height/2)

	for y := 0; y < int(paddle.height); y++ {
		for x := 0; x < int(paddle.width); x++ {
			setPixel(startX+x, startY+y, color{255, 255, 255})
		}
	}

	numX := lerp(paddle.x, getCenter().x, 0.2)
	drawNumber(pos{numX, 35}, paddle.color, 10, paddle.score)
}

func (paddle *paddle) update(elapsedTime float32) {
	if keyState[sdl.SCANCODE_UP] != 0 || keyState[sdl.SCANCODE_W] != 0 {
		paddle.y -= paddle.speed * elapsedTime
	}
	if keyState[sdl.SCANCODE_DOWN] != 0 || keyState[sdl.SCANCODE_S] != 0 {
		paddle.y += paddle.speed * elapsedTime
	}
}

func (paddle *paddle) aiUpdate(ball *ball) {
	paddle.y = ball.y
}
