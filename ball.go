package main

import (
	"github.com/veandco/go-sdl2/mix"
)

type ball struct {
	pos
	radius    float32
	xVelocity float32
	yVelocity float32
	color     color
}

func (ball *ball) draw() {
	for y := -ball.radius; y < ball.radius; y++ {
		for x := -ball.radius; x < ball.radius; x++ {
			if x*x+y*y < ball.radius*ball.radius {
				setPixel(int(ball.x+x), int(ball.y+y), color{255, 255, 255})
			}
		}
	}
}

func getCenter() pos {
	return pos{
		x: float32(windowWidth / 2),
		y: float32(windowHeight / 2),
	}
}

func (ball *ball) update(left *paddle, right *paddle, elapsedTime float32, ballHitSound *mix.Chunk) {
	ball.x += ball.xVelocity * elapsedTime
	ball.y += ball.yVelocity * elapsedTime

	if ball.y-ball.radius < 0 || ball.y+ball.radius > float32(windowHeight) {
		ballHitSound.Play(1, 0)
		ball.yVelocity = -ball.yVelocity
	}

	if ball.x < 0 {
		right.score++
		ball.pos = getCenter()
		state = start
	} else if int(ball.x) > windowWidth {
		left.score++
		ball.pos = getCenter()
		state = start
	}

	if ball.x-ball.radius < left.x+left.width/2 {
		if ball.y > left.y-left.height/2 && ball.y < left.y+left.height/2 {
			ball.xVelocity = -ball.xVelocity
			ball.x = left.x + left.width/2 + ball.radius
			ballHitSound.Play(1, 0)
		}
	}

	if ball.x+ball.radius > right.x-right.width/2 {
		if ball.y > right.y-right.height/2 && ball.y < right.y+right.height/2 {
			ball.xVelocity = -ball.xVelocity
			ball.x = right.x - right.width/2 - ball.radius
			ballHitSound.Play(1, 0)
		}
	}
}
