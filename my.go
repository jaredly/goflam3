package main

import (
	"os"
	"image"
	"image/png"
	"image/color"
	"math"
	"math/rand"
)

// Our functions 
func f0(x, y, a, b, c, d, e, f float64) (float64, float64) {
	return x / 2, y / 2
}

func f1(x, y, a, b, c, d, e, f float64) (float64, float64) {
	return (x + 1) / 2, y / 2
}

func f2(x, y, a, b, c, d, e, f float64) (float64, float64) {
	return x / 2, (y + 1) / 2
}

func f3(x, y, a, b, c, d, e, f float64) (float64, float64) {
	return math.Sin(x), math.Sin(y)
}

func f4(x, y, a, b, c, d, e, f float64) (float64, float64) {
	return math.Cos(x), math.Cos(y)
}

func f5(x, y, a, b, c, d, e, f float64) (float64, float64) {
	return math.Cos(a*x + b*y + c), math.Cos(d*x + e*y + f)
}

// get the third-largest value in a matrix. I can probably do this better
func maxmx(arr [][]int) int {
	mx := 0
	snd := 0
	thrd := 0
	for _, row := range arr {
		for _, v := range row {
			if mx < v {
				thrd = snd
				snd = mx
				mx = v
			}
		}
	}
	return thrd
}

/*
(x, y) = a random point in the biunit square
iterate {
   i = a random integer from 0 to n  1 inclusive
   (x, y) = Fi(x, y)
   plot (x, y) except during the first 20 iterations
}
*/

func flame(width, height, iters int) *image.RGBA {
	x := rand.Float64()*2 - 1
	y := rand.Float64()*2 - 1
	mx := make([][]int, height)
	for y := range mx {
		mx[y] = make([]int, width)
		for x := range mx[y] {
			mx[y][x] = 0
		}
	}
	var a, b, c, d, e, f float64
	// these are our parameters
	a, b, c, d, e, f = 1, 2, 1, 1, 4, 5
	// and the F_i s that we'll be using
	funcs := []func(float64, float64, float64, float64, float64, float64, float64, float64) (float64, float64){f1, f2, f3, f5}
	for at := 0; at < iters; at++ {
		x, y = funcs[rand.Intn(len(funcs))](x, y, a, b, c, d, e, f)
		// I should probably refactor this
		if x < -1 {
			x = -1
			continue
		}
		if x > 1 {
			x = 1
			continue
		}
		if y < -1 {
			y = -1
			continue
		}
		if y > 1 {
			y = 1
			continue
		}
		//fmt.Println("after", x,y)
		if at < 20 {
			continue
		}
		// refactor, make more readable
		mx[int((y+1)/2*float64(height-1))][int((x+1)/2*float64(width-1))] += 1
	}
	max := maxmx(mx)
	m := image.NewRGBA(image.Rect(0, 0, width, height))
	// now write the values to an image, equalized by the 3rd-brightest point
	for x, row := range mx {
		for y, v := range row {
			val := uint8(255 * v / max)
			if val > 255 {
				val = 255
			}
			m.Set(x, y, color.RGBA{val, val * 100 / 255, val, 255})
		}
	}
	return m
}

func main() {
	m := flame(800, 800, 10000000)
	toimg, _ := os.Create("new1235.png")
	defer toimg.Close()

	png.Encode(toimg, m)
}
