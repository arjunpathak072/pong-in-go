package main

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

var winner = nobody
var state = start
var keyState = sdl.GetKeyboardState()

const (
	windowWidth  = 800
	windowHeight = 600
	maxScore     = 1
	lineWidth    = 2
)

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println("initializing SDL: ", err)
		return
	}
	defer sdl.Quit()

	err = mix.Init(0)
	if err != nil {
		fmt.Println("initializing Mix: ", err)
		return
	}

	mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 1024)
	defer mix.CloseAudio()

	defeatSound, err := mix.LoadWAV("assets/game-over.mp3")
	if err != nil {
		fmt.Println("reading game over sound effect: ", err)
		return
	}
	defer defeatSound.Free()

	ballHitSound, err := mix.LoadWAV("assets/ball-hit.wav")
	if err != nil {
		fmt.Println("reading ball hit sound: ", err)
	}
	defer ballHitSound.Free()

	victorySound, err := mix.LoadWAV("assets/victory.mp3")
	if err != nil {
		fmt.Println("reading victory sound effect: ", err)
	}
	defer victorySound.Free()

	backgroundMusic, err := mix.LoadMUS("assets/background-music.mp3")
	if err != nil {
		fmt.Println("reading background music: ", err)
	}

	window, err := sdl.CreateWindow(
		"pong in Go I guess",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		windowWidth,
		windowHeight,
		sdl.WINDOW_SHOWN,
	)

	if err != nil {
		fmt.Println("initializing window: ", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing the renderer: ", err)
		return
	}
	defer renderer.Destroy()

	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, windowWidth, windowHeight)
	if err != nil {
		fmt.Println("creating texture: ", err)
		return
	}
	defer texture.Destroy()

	playerOne := paddle{
		pos:    pos{50, getCenter().y},
		width:  20,
		height: 100,
		speed:  300,
		score:  0,
		color:  color{255, 255, 255},
	}

	playerTwo := paddle{
		pos:    pos{windowWidth - 50, getCenter().y},
		width:  20,
		height: 100,
		speed:  300,
		score:  0,
		color:  color{255, 255, 255},
	}

	ball := ball{
		pos:       getCenter(),
		radius:    25,
		xVelocity: 400,
		yVelocity: 400,
		color:     color{255, 255, 255},
	}

	var frameStart time.Time
	var elapsedTime float32

	for {
		frameStart = time.Now()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		if state == play {
			playerOne.update(elapsedTime)
			playerTwo.aiUpdate(&ball)
			ball.update(&playerOne, &playerTwo, elapsedTime, ballHitSound)
		} else if state == start {
			if playerTwo.score == maxScore {
				winner = computer
			} else if playerOne.score == maxScore {
				winner = player
			}
			if keyState[sdl.SCANCODE_SPACE] != 0 {
				backgroundMusic.Play(10)
				state = play
			}
		}

		clear()
		if winner != nobody {
			switch winner {
			case player:
				mix.HaltMusic()
				defeatSound.Play(1, 0)
				drawMessage(getCenter(), color{0, 255, 0}, 10, &winnerMessage)
				fmt.Println("you won!")
			case computer:
				mix.HaltMusic()
				victorySound.Play(1, 0)
				drawMessage(getCenter(), color{255, 0, 0}, 10, &loserMessage)
				fmt.Println("compter won.")
			}

			playerOne.score = 0
			playerTwo.score = 0
			winner = nobody
			state = start

			texture.Update(nil,
				unsafe.Pointer(&pixels[0]),
				windowWidth*4,
			)
			renderer.Copy(texture, nil, nil)
			renderer.Present()
			sdl.WaitEventTimeout(10000)
		} else {
			drawLine()
			playerOne.draw()
			playerTwo.draw()
			ball.draw()
		}

		texture.Update(nil,
			unsafe.Pointer(&pixels[0]),
			windowWidth*4,
		)
		renderer.Copy(texture, nil, nil)
		renderer.Present()

		elapsedTime = float32(time.Since(frameStart).Seconds())
		if elapsedTime < 0.005 {
			sdl.Delay(5 - uint32(elapsedTime*1000.0))
			elapsedTime = float32(time.Since(frameStart).Seconds())
		}
	}
}
