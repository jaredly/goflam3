package main

import (
  "os"
  "fmt"
  "image"
	bin "encoding/binary"
  "math/rand"
  "time"
)

type Point struct {
  X, Y float64
}

/*
func loadData(fname string) *[]Point {
	file, err := os.Open(fname)
	if err != nil {
		return nil
	}
	defer file.Close()
	var num, i int32
	bin.Read(file, bin.LittleEndian, &num)
	data := make([]Point, num)
	for i=0; i<num; i++ {
		bin.Read(file, bin.LittleEndian, &data[i])
	}
	return &data
}

func storeData(fname string, data *[]Point) {
	// fmt.Fprintln(os.Stderr, "Storing Data")
	file, err := os.Create(fname)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to open dataout file for writing")
		return
	}
	defer file.Close()
	num := len(*data)
	// println("Number to save :", num)
	bin.Write(file, bin.LittleEndian, int32(num))
	for _, v := range *data {
		bin.Write(file, bin.LittleEndian, &v)
	}
}

func flame(config Config) *image.RGBA {
	start := time.Now()
	allfuncs := AllVariations()

	fmt.Fprintln(os.Stderr, "Flaming")
	var data *[]Point
	if config.DataIn != "" {
		fmt.Fprintln(os.Stderr, "From data")
		data = loadData(config.DataIn)
		if data == nil {
			fmt.Fprintln(os.Stderr, "Failed to open datain")
			return nil
		}
		fmt.Fprintln(os.Stderr, "Load time", time.Since(start).String())
	} else {
		fmt.Fprintln(os.Stderr, "Generating")
		data = generate(config.Iterations, config.Functions, allfuncs)
		fmt.Fprintln(os.Stderr, "Generate time", time.Since(start).String())
		if config.DataOut != "" {
			start = time.Now()
			fmt.Fprintln(os.Stderr, "To data")
			storeData(config.DataOut, data)
			fmt.Fprintln(os.Stderr, "Save time", time.Since(start).String())
		}
	}

	start = time.Now()
	if config.NoImage {
		return nil
	}
	image := render(config.Width, config.Height, data)
	fmt.Fprintln(os.Stderr, "Render time", time.Since(start).String())
	return image
}

func generate(iters int, usefuncs []FunConfig, variations []Variation) *[]Point {
	x := rand.Float64()*2 - 1
	y := rand.Float64()*2 - 1
	data := make([]Point, iters)
	// should these be passed in as an array?
	var a, b, c, d, e, f float64
	// these are our parameters
	a, b, c, d, e, f = 1, 2, 1, 1, 4, 5
	// and the F_i s that we'll be using
	for at := 0; at < iters; at++ {
		fi := rand.Intn(len(usefuncs))
		x, y = variations[usefuncs[fi].Num](x, y, a, b, c, d, e, f)
		if at < 20 {
			continue
		}
		data[at] = Point{x, y}
	}
	return &data
}

*/

func altFlame(config Config) *image.RGBA {
	start := time.Now()
	allfuncs := AllVariations()

	fmt.Fprintln(os.Stderr, "Flaming")
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

type Pixel struct {
	Alpha int32
	Funcs [20]int32
}

func genMatrix(iters, width, height int, usefuncs []FunConfig, variations []FullVar) *[][]Pixel {
	x := rand.Float64()*2 - 1
	y := rand.Float64()*2 - 1
	raw := make([]Pixel, height * width)
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
			tx = width -1
		}
		ty := int((y + 1) * float64(height) / 2)
		if ty == height {
			ty = height - 1
		}
		mx[ty][tx].Alpha += 1
		mx[ty][tx].Funcs[fi] += 1
	}
	return &mx
}

func saveMatrix(fname string, matrix *[][]Pixel) {
	fmt.Fprintln(os.Stderr, "Storing Data")
	file, err := os.Create(fname)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to open dataout file for writing")
		return
	}
	defer file.Close()
	// height
	bin.Write(file, bin.LittleEndian, int32(len(*matrix)))
	// width
	bin.Write(file, bin.LittleEndian, int32(len((*matrix)[0])))
	for y := range *matrix {
		for x := range (*matrix)[y] {
			bin.Write(file, bin.LittleEndian, (*matrix)[y][x])
		}
	}
}

func loadMatrix(fname string) *[][]Pixel {
	file, err := os.Open(fname)
	if err != nil {
		return nil
	}
	defer file.Close()
	var height, width int32
	bin.Read(file, bin.LittleEndian, &height)
	bin.Read(file, bin.LittleEndian, &width)
	raw := make([]Pixel, height * width)
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
