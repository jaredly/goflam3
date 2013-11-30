package flame

import (
	bin "encoding/binary"
	"fmt"
	"image"
	"math/rand"
	"os"
	"io"
	"time"
)

const LOG = false

func log(i io.Writer, a... interface{}) {
    if (LOG) {
	fmt.Fprintln(i, a...)
    }
}


type Point struct {
	X, Y float64
}

func FromConfigFile(fname string) (*image.RGBA, error) {
    var c Config
    err := ReadConfig(fname, &c)
    if err != nil {
	return nil, err
    }
    return Flame(c)
}

// Flame generates the Fractal image
func Flame(config Config) (*image.RGBA, error) {
	start := time.Now()
	allfuncs := AllVariations()

	log(os.Stderr, "Flaming", config.Functions, config.Dims)
	var data *[][]Pixel
	if config.DataIn != "" {
		log(os.Stderr, "From data")
		data = loadMatrix(config.DataIn)
		if data == nil {
			log(os.Stderr, "Failed to open datain")
			return nil, nil
		}
		log(os.Stderr, "Load time", time.Since(start).String())
	} else {
		log(os.Stderr, "Generating")
		data = genMatrix(config.Iterations, config.Dims, config.Functions, allfuncs)
		log(os.Stderr, "Generate time", time.Since(start).String())
		if config.DataOut != "" {
			start = time.Now()
			log(os.Stderr, "To data")
			saveMatrix(config.DataOut, data)
			log(os.Stderr, "Save time", time.Since(start).String())
		}
	}

	start = time.Now()
	if config.NoImage {
		return nil, nil
	}
	image := renderMatrix(data, config.LogEqualize)
	log(os.Stderr, "Render time", time.Since(start).String())
	return image, nil
}

// Holds information about a pixel; how many times it was hit, and by which
// functions. We could actually probably drop the alpha...
type Pixel struct {
	Alpha int64
	Funcs [20]int64
}

/*
func goMatrix(concurrency, iters, width, height int, usefuncs []FunConfig, variations []FullVar) *[][]Pixel {
    // return genMatrix(iters, width, height, usefuncs, variations)
    cha := make(chan bool)
    raw := make([]Pixel, height*width)
    mx := make([][]Pixel, height)
    beg := time.Now()
    for i := range mx {
	mx[i], raw = raw[:width], raw[width:]
    }
    worker := func(i int){
	start := time.Now()
	log(os.Stderr, "Delay", i, time.Since(beg).String())
	populateMatrix(&mx, iters/concurrency, width, height, usefuncs, variations)
	log(os.Stderr, "Gen time", i, time.Since(start).String())
	cha<-true
    }
    for i:=0; i<concurrency; i++ {
	go worker(i)
    }
    for i:=0; i<concurrency; i++ {
	<-cha
    }
    return &mx
}
*/

// Run the algorithm, returning a matrix of Pixels
func genMatrix(iters int, dims Dims, usefuncs []FunConfig, variations []FullVar) *[][]Pixel {
	raw := make([]Pixel, dims.Height*dims.Width)
	mx := make([][]Pixel, dims.Height)
	for i := range mx {
		mx[i], raw = raw[:dims.Width], raw[dims.Width:]
	}
	/*
	    populateMatrix(&mx, iters, dims, usefuncs, variations)
	    }
	    func populateMatrix(mx *[][]Pixel, iters, width, height int, usefuncs []FunConfig, variations []FullVar) {
	*/
	// should these be passed in as an array?
	x := rand.Float64()*2 - 1
	y := rand.Float64()*2 - 1
	var a, b, c, d, e, f float64
	// these are our parameters
	a, b, c, d, e, f = DefaultParams()
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
		ty := int(((x - dims.X) * dims.Xscale + 1) * float64(dims.Width) / 2)
		if ty == dims.Width {
			ty = dims.Width - 1
		}
		tx := int(((-y + dims.Y) * dims.Yscale + 1) * float64(dims.Height) / 2)
		if tx == dims.Height {
			tx = dims.Height - 1
		}
		if ty < 0 {
			continue
		}
		if tx < 0 {
			continue
		}
		if tx > dims.Width {
		    continue
		}
		if ty > dims.Height {
		    continue
		}
		mx[ty][tx].Alpha += 1
		mx[ty][tx].Funcs[fi] += 1
	}
	return &mx
}

// Save the matrix in raw form to a file
func saveMatrix(fname string, matrix *[][]Pixel) {
	log(os.Stderr, "Storing Data")
	file, err := os.Create(fname)
	if err != nil {
		log(os.Stderr, "Failed to open dataout file for writing")
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
