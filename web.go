package main

import (
	"fmt"
	"bytes"
	"image"
	"encoding/json"
	"encoding/base64"
	"github.com/codegangsta/cli"
	"github.com/codegangsta/martini"
	"image/png"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func imageToBase64(image *image.RGBA) string {
	var b bytes.Buffer
	png.Encode(&b, image)
	return fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(b.Bytes()))
}

type WebRender struct {
	Image    string
	Time     time.Duration
	Formulas []FunConfig
	Disabled bool
	Text     string
}

type WebFunc struct {
	Num  int
	Text string
}

type InitialResponse struct {
	MainImage    string
	Formulas     []WebFunc
	ChildImages  []WebRender
	MainFormulas []FunConfig
}

func GetWebFuncs() []WebFunc {
	exp := AllVariations()
	res := make([]WebFunc, len(exp))
	for i := range res {
		res[i] = WebFunc{
			Num:  i,
			Text: exp[i].Text,
		}
	}
	return res
}

func howTrue(bools []bool) int {
	num := 0
	for _, t := range bools {
		if t {
			num += 1
		}
	}
  return num
}

func usingToFuns(using []bool) []FunConfig {
  num := howTrue(using)
	res := make([]FunConfig, num)
	z := 0
	for i, t := range using {
		if t {
			res[z] = FunConfig{i}
			z += 1
		}
	}
	return res
}

func RenderChildren(width, height, iterations int, funs []FunConfig) []WebRender {
	texts := AllVariations()
	variations := len(texts)
	res := make([]WebRender, variations)
	using := make([]bool, variations)
	for _, f := range funs {
		using[f.Num] = true
	}
	for i := range res {
		using[i] = !using[i]
		funs := usingToFuns(using)
    var im *image.RGBA
		start := time.Now()
		if len(funs) == 0 {
      im = blankImage(width, height)
		} else {
      im = altFlame(Config{
				Width:      width,
				Height:     height,
				Iterations: iterations,
				Functions:  funs,
			})
    }
		res[i] = WebRender{
			Image: imageToBase64(im),
			Disabled: using[i],
			Time:     time.Since(start),
			Formulas: funs,
			Text:     texts[i].Text,
		}
		using[i] = !using[i]
	}
	return res
}

func getFunsFromParam(param string) []FunConfig {
	nums := strings.Split(param, ":")
	ret := make([]FunConfig, len(nums))
	for i, x := range nums {
		n, err := strconv.Atoi(x)
		if err != nil {
			return nil
		}
		ret[i] = FunConfig{n}
	}
	return ret
}

type Cachier map[string]InitialResponse

func cliWebserver(c *cli.Context) {
	m := martini.Classic()
	// cachier := &Cachier{}
	// m.Map(cachier)

	m.Get("/render", func(resrw http.ResponseWriter, req *http.Request) string {
		var funs []FunConfig
		resrw.Header().Set("Content-Type", "application/json")
		req.ParseForm()
		if 1 == len(req.Form["funcs"]) {
			funs = getFunsFromParam(req.Form["funcs"][0])
		}
    var im *image.RGBA
    if funs == nil || len(funs) == 0 {
      im = blankImage(300, 300)
    } else {
      im = altFlame(Config{
			  Width:      300,
			  Height:     300,
			  Iterations: 1000 * 1000,
			  Functions:  funs,
		  })
    }
		response := InitialResponse{
			MainImage: imageToBase64(im),
			MainFormulas: funs,
			Formulas:     GetWebFuncs(),
			ChildImages:  RenderChildren(150, 150, 100*1000, funs),
		}
		res, _ := json.Marshal(response)
		return string(res)
	})
	m.Get("/high-def", func(req *http.Request) string {
		var funs []FunConfig
		req.ParseForm()
		if 1 == len(req.Form["funcs"]) {
			funs = getFunsFromParam(req.Form["funcs"][0])
		}
    var im *image.RGBA
    if funs == nil || len(funs) == 0 {
      im = blankImage(1000, 1000)
    } else {
      im = altFlame(Config{
			  Width:      1000,
			  Height:     1000,
			  Iterations: 10 * 1000 * 1000,
			  Functions:  funs,
		  })
    }
		return imageToBase64(im)
	})
	m.Run()
}
