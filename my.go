package main

import (
  "os"
  "fmt"
  // "strconv"
  "image"
	bin "encoding/binary"
  // "image/png"
  // "math"
  "math/rand"
  "time"
)

type Point struct {
  X, Y float64
}

func loadData(fname string) *[]Point {
	file, err := os.Open(fname)
	if err != nil {
		return nil
	}
	defer file.Close()
	var num int
	bin.Read(file, bin.LittleEndian, &num)
	println("Read number", num)
	num = 10 * 1000 * 1000
	data := make([]Point, num)
	for i:=0; i<num; i++ {
		bin.Read(file, bin.LittleEndian, &data[i])
	}
	println("Done reading")
	return &data
}

func storeData(fname string, data *[]Point) {
	fmt.Fprintln(os.Stderr, "Storing Data")
	file, err := os.Create(fname)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to open dataout file for writing")
		return
	}
	defer file.Close()
	num := len(*data)
	println("Number to save :", num)
	bin.Write(file, bin.LittleEndian, &num)
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
	} else {
		fmt.Fprintln(os.Stderr, "Generating")
		data = generate(config.Iterations, config.Functions, allfuncs)
		if config.DataOut != "" {
			fmt.Fprintln(os.Stderr, "To data")
			storeData(config.DataOut, data)
		}
	}
	fmt.Fprintln(os.Stderr, "Generate time", time.Since(start).String())

	other := time.Now()
	if config.NoImage {
		return nil
	}
	image := render(config.Width, config.Height, data)
	fmt.Fprintln(os.Stderr, "Render time", time.Since(other).String())
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

/*
func writeit(w, h, i int, use []int) {
	m := flame(w, h, i, use)
	name := "nolimit-"
	num := ""
	for _, z := range use {
		num = strconv.Itoa(z) + num
	}
	toimg, _ := os.Create(name + num + ".png")
	defer toimg.Close()

	png.Encode(toimg, m)
}

func allCombos(w, h, i, from, to int) {
	from = int(math.Pow(2, float64(from))) - 1
	to = int(math.Pow(2, float64(to)))
	for z := from; z < to; z++ {
		n := 0
		for s := z; s > 0; s >>= 1 {
			if s % 2 == 1 {
				n += 1
			}
		}
		use := make([]int, n)
		n = 0
		a := 0
		for s := z; s > 0; s >>= 1 {
			if s % 2 == 1 {
				use[n] = a
				n += 1
			}
			a += 1
		}
		fmt.Println("Yeah", use)
		writeit(w, h, i, use)
	}
}
*/
