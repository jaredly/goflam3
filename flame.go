package main

import (
	bin "encoding/binary"
	"fmt"
	"image"
	"math/rand"
	"os"
	"time"
)

type Point struct {
	X, Y float64
}

// Generate the Fractal image
func altFlame(config Config) *image.RGBA {
	start := time.Now()
	allfuncs := AllVariations()

	fmt.Fprintln(os.Stderr, "Flaming", config.Functions)
	var data *[][]Pixel
	if config.DataIn != "" {
		fmt.Fprintln(os.Stderr, "From data")
		data = loadMatrix(config.DataIn)
		if data == nil {
			fmt.Fprintln(os.Stderr, "Failed to open datain")
			return nil
		}
		fmt.Fprintln(os.Stderr, "Load time", time.Since(start).String())
	} else {
		fmt.Fprintln(os.Stderr, "Generating")
		data = genMatrix(config.Iterations, config.Width, config.Height, config.Functions, allfuncs)
		fmt.Fprintln(os.Stderr, "Generate time", time.Since(start).String())
		if config.DataOut != "" {
			start = time.Now()
			fmt.Fprintln(os.Stderr, "To data")
			saveMatrix(config.DataOut, data)
			fmt.Fprintln(os.Stderr, "Save time", time.Since(start).String())
		}
	}

	start = time.Now()
	if config.NoImage {
		return nil
	}
	image := renderMatrix(data)
	fmt.Fprintln(os.Stderr, "Render time", time.Since(start).String())
	return image
}

// Holds information about a pixel; how many times it was hit, and by which
// functions. We could actually probably drop the alpha...
type Pixel struct {
	Alpha int64
	Funcs [20]int64
}

// Run the algorithm, returning a matrix of Pixels
func genMatrix(iters, width, height int, usefuncs []FunConfig, variations []FullVar) *[][]Pixel {
	x := rand.Float64()*2 - 1
	y := rand.Float64()*2 - 1
	raw := make([]Pixel, height*width)
	mx := make([][]Pixel, height)
	for i := range mx {
		mx[i], raw = raw[:width], raw[width:]
	}
	// should these be passed in as an array?
	var a, b, c, d, e, f float64
	// these are our parameters
	a, b, c, d, e, f = 1, 2, 1, 1, 4, 5
	// and the F_i s that we'll be using
	for at := 0; at < iters; at++ {
		fi := rand.Intn(len(usefuncs))
		x, y = variations[usefuncs[fi].Num].Fn(x, y, a, b, c, d, e, f)
		if at < 20 {
			continue
		}
		if x > 1 || x < -1 || y > 1 || y < -1 {
			continue
		}
		tx := int((x + 1) * float64(width) / 2)
		if tx == width {
			tx = width - 1
		}
		ty := int((y + 1) * float64(height) / 2)
		if ty == height {
			ty = height - 1
		}
		if ty < 0 {
			continue
		}
		if tx < 0 {
			continue
		}
		mx[ty][tx].Alpha += 1
		mx[ty][tx].Funcs[fi] += 1
	}
	return &mx
}

// Save the matrix in raw form to a file
func saveMatrix(fname string, matrix *[][]Pixel) {
	fmt.Fprintln(os.Stderr, "Storing Data")
	file, err := os.Create(fname)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to open dataout file for writing")
		return
	}
	defer file.Close()
	// height
	bin.Write(file, bin.LittleEndian, int64(len(*matrix)))
	// width
	bin.Write(file, bin.LittleEndian, int64(len((*matrix)[0])))
	for y := range *matrix {
		for x := range (*matrix)[y] {
			bin.Write(file, bin.LittleEndian, (*matrix)[y][x])
		}
	}
}

// Load a raw matrix from a file
func loadMatrix(fname string) *[][]Pixel {
	file, err := os.Open(fname)
	if err != nil {
		return nil
	}
	defer file.Close()
	var height, width int64
	bin.Read(file, bin.LittleEndian, &height)
	bin.Read(file, bin.LittleEndian, &width)
	raw := make([]Pixel, height*width)
	matrix := make([][]Pixel, height)
	for i := range matrix {
		matrix[i], raw = raw[:width], raw[width:]
	}
	for y := range matrix {
		for x := range matrix[y] {
			bin.Read(file, bin.LittleEndian, &matrix[y][x])
		}
	}
	return &matrix
}
