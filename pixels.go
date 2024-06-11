package main

var pixels = make([]byte, windowWidth*windowHeight*4)

func setPixel(x, y int, c color) {
	index := (y*windowWidth + x) * 4
	if index < len(pixels)-4 && index >= 0 {
		pixels[index] = c.r
		pixels[index+1] = c.g
		pixels[index+2] = c.b
		pixels[index+3] = 255
	}
}

func clear() {
	for i := range pixels {
		pixels[i] = 0
	}
}
