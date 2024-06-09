package main

func drawLine() {
	pos := getCenter()
	startX := int(pos.x - float32(lineWidth))
	startY := 0

	for y := startY; y < windowWidth; y++ {
		for x := startX; x < int(pos.x+float32(lineWidth)); x++ {
			setPixel(x, y, color{128, 128, 128})
		}
	}
}

func drawMessage(pos pos, color color, size int, message *[][]byte) {
	offset := size * 5 * 2
	var startX int
	if startX%2 == 0 {
		startX = int(pos.x) - size*5*((len(*message)/2)*2)
	} else {
		startX = int(pos.x) - (size*5)/2 - size*5*((len(*message)/2)*2)
	}
	startY := int(pos.y) - (size*6)/2

	for _, letter := range *message {
		for i, v := range letter {
			if v == 1 {
				for y := startY; y < startY+size; y++ {
					for x := startX; x < startX+size; x++ {
						setPixel(x, y, color)
					}
				}
			}
			startX += size
			if (i+1)%5 == 0 {
				startY += size
				startX -= size * 5
			}
		}
		startY = int(pos.y) - (size*6)/2
		startX += offset
	}
}

func drawNumber(pos pos, color color, size int, number int) {
	startX := int(pos.x) - (size*3)/2
	startY := int(pos.y) - (size*5)/2

	for i, v := range nums[number] {
		if v == 1 {
			for y := startY; y < startY+size; y++ {
				for x := startX; x < startX+size; x++ {
					setPixel(x, y, color)
				}
			}
		}
		startX += size
		if (i+1)%3 == 0 {
			startY += size
			startX -= size * 3
		}
	}
}

func lerp(a, b, pct float32) float32 {
	return a + pct*(b-a)
}
