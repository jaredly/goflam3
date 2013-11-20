package main

import (
	"encoding/json"
	"github.com/codegangsta/cli"
	"github.com/kless/term"
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"os"
)

type FunConfig struct {
	Num int
	// Args []float64
}

type Config struct {
	Width      int
	Height     int
	Iterations int
	Functions  []FunConfig
	DataIn     string
	DataOut    string
	NoImage    bool
	// LogEqualize
	// GammaCorrect
}

func readConfig(fname string, config *Config) {
	bytes, err := ioutil.ReadFile(fname)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, config)
	if err != nil {
		println("Failed to parse config file")
	}
}

func cliRender(c *cli.Context) {
	var config Config
	config.Width = 400
	config.Height = 400
	config.Iterations = 10 * 1000 * 1000
	config.Functions = []FunConfig{
		{5},
		{7},
	}
	config.DataOut = c.String("dataout")
	config.DataIn = c.String("datain")
	config.NoImage = c.Bool("noimage")
	if c.String("config") != "" {
		readConfig(c.String("config"), &config)
	}
	if c.Int("width") != 0 {
		config.Width = c.Int("width")
	}
	if c.Int("height") != 0 {
		config.Height = c.Int("height")
	}
	if c.Int("iterations") != 0 {
		config.Iterations = c.Int("iterations")
	}
	if len(c.IntSlice("f")) != 0 {
		config.Functions = make([]FunConfig, len(c.IntSlice("f")))
		for i, v := range c.IntSlice("f") {
			config.Functions[i] = FunConfig{v}
		}
	}
	outfile := "-"
	if c.String("outfile") != "" {
		outfile = c.String("outfile")
	}
	image := altFlame(config)
	if image != nil {
		write(outfile, image)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "Goflam3'"
	app.Usage = "Render fractal flames using go"
	var nofuncs cli.IntSlice
	app.Commands = []cli.Command{
		{
			Name:      "serve",
			ShortName: "s",
			Usage:     "Webserver",
			Flags: []cli.Flag{
				cli.IntFlag{"port", 4000, "The port to run on (defaults to 3000)"},
			},
			Action: cliWebserver,
		},
		{
			Name:      "render",
			ShortName: "r",
			Usage:     "Render a file",
			Flags: []cli.Flag{
				cli.StringFlag{"dataout, d", "", "Output generated points to the given path"},
				cli.StringFlag{"datain", "", "Don't generate points; read them from the given path"},
				cli.StringFlag{"config, c", "", "File containing config. Defaults to stdin if no datain is given"},
				cli.StringFlag{"outfile, o", "", "The name of the output image. Defaults to stdout if -noimage is not set"},
				cli.BoolFlag{"noimage, n", "Don't output an image; only dataout"},
				// these can also be set by the config file
				cli.IntFlag{"width, w", 0, "Width of the image in px"},
				cli.IntFlag{"height", 0, "Height of the image in px"},
				cli.IntFlag{"iterations, i", 0, "Number of iterations to execute"},
				cli.IntSliceFlag{"function, f", &nofuncs, "Functions to use"},
			},
			Action: cliRender,
		},
	}
	app.Action = cliRender
	app.Run(os.Args)
}

func write(outfile string, image *image.RGBA) {
	var out io.Writer
	if outfile == "-" && term.IsTerminal(int(os.Stdout.Fd())) {
		println("You're on a terminal, and I assume you don't want a face full of PNG binary. " +
			"Specify -outfile if you want some name other than flame-image.png")
		outfile = "flame-image.png"
	}
	if outfile != "-" {
		outimg, err := os.Create(outfile)
		if err != nil {
			println("Failed to open file for writing", outfile)
			return
		}
		out = outimg
		defer outimg.Close()
	} else {
		out = os.Stdout
	}
	png.Encode(out, image)
}

// cliwebserver

/*
func main() {
  w := 800
  h := 800
  i := 10000000
  writeit(w, h, i, []int{7,3})
  // writeit(w, h, i, []int{3,5})
  // callThemAll(w, h, i, 3, 12)
  // allCombos(w, h, i, 0, 7)
}
*/
