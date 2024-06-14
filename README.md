# Pong game in Go
A classic pong game clone written in Go, using SDL2

## Getting Started

### Prerequisites
* Go 1.17 or later installed on your system
* SDL 2.0 or later installed on your system
* A compatible graphics driver (e.g. OpenGL, DirectX, etc.)

### Building and Running
1. Clone this repository: `git clone https://github.com/arjunpathak072/pong-in-go.git`
2. Change into the project directory: `cd pong-in-go`
3. Build the project: `go build .`
4. Run the game: `./pong-in-go`

## Gameplay in action
[![Watch the video](https://raw.githubusercontent.com/arjunpathak072/pong-in-go/main/assets/gameplay.gif)](https://raw.githubusercontent.com/arjunpathak072/pong-in-go/main/assets/gameplay.mp4)

## Technical Jargon
* Uses Go bindings of the SDL2 library to handle rendering
* Uses the sdl/mix module in go to implement sounds in the game
* Pixel array manipulation for rendering content on screen
* Custom Pixelated Font made using byte arrays

## Features of the game
* Intuitive arrow key controls to move the paddle
* A literally undefeatable AI as your opponent
* Retro arcade sound effects all through the game!
* Random crashes and bugs for extra unpredictability

## Acknowledgement
* The SDL library for providing a cross-platform graphics and input handling API.
* The Go community for providing a great language and ecosystem.