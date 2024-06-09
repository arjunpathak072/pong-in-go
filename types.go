package main

type gameWinner int
const (
	nobody   gameWinner = iota
	player   gameWinner = iota
	computer gameWinner = iota
)

type gameState int
const (
	start gameState = iota
	play  gameState = iota
)

type color struct {
	r, g, b byte
}

type pos struct {
	x, y float32
}