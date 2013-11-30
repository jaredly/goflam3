package flame

import (
  "io/ioutil"
  "encoding/json"
)

// FunConfig contains configuration for the various functions to be used in
// generating a flame
type FunConfig struct {
	Num int
	// Args []float64
}

type Dims struct {
  Width int
  Height int
  X float64
  Y float64
  Xscale float64
  Yscale float64
}

// Config holds all of the parameters necessary to generate a flame
type Config struct {
	Dims Dims
	Iterations int
	Functions  []FunConfig
	DataIn     string
	DataOut    string
	NoImage    bool
	LogEqualize bool
	// GammaCorrect
}

func ReadConfig(fname string, config *Config) (err error) {
	bytes, err := ioutil.ReadFile(fname)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, config)
	if err != nil {
		println("Failed to parse config file")
	}
	return err
}

