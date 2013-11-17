package main

import (
	"fmt"
        "os"
	"strings"
	"strconv"
	"image/png"
	"image"
	"bytes"
	"encoding/json"
	"encoding/base64"
	"github.com/codegangsta/cli"
	"github.com/hoisie/web"
	"time"
)

func imageToBase64(image *image.RGBA) string {
	var b bytes.Buffer
	png.Encode(&b, image)
	return fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(b.Bytes()))
}

type WebRender struct {
	Image string
	Time time.Duration
	Formulas []FunConfig
	Disabled bool
	Text string
}

type WebFunc struct {
	Num int
	Text string
}

type InitialResponse struct {
	MainImage string
	Formulas []WebFunc
	ChildImages []WebRender
	MainFormulas []FunConfig
}

func GetWebFuncs() []WebFunc {
	exp := Explanation()
	res := make([]WebFunc, len(exp))
	for i := range res {
		res[i] = WebFunc{
			Num: i,
			Text: exp[i],
		}
	}
	return res
}

func usingToFuns(using []bool) []FunConfig {
	num := 0
	for _, t := range using {
		if t {
			num += 1
		}
	}
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
	texts := Explanation()
	variations := len(texts)
	res := make([]WebRender, variations)
	using := make([]bool, variations)
	for _, f := range funs {
		using[f.Num] = true
	}
	for i := range res {
		using[i] = !using[i]
		funs := usingToFuns(using)
		if len(funs) == 0 {
			continue
		}
		start := time.Now()
		res[i] = WebRender{
			Image: imageToBase64(altFlame(Config{
				Width: width,
				Height: height,
				Iterations: iterations,
				Functions: funs,
			})),
			Disabled: using[i],
			Time: time.Since(start),
			Formulas: funs,
			Text: texts[i],
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

func cliWebserver(c *cli.Context) {
	w := web.NewServer()
	w.Get("/render", func(ctx *web.Context) string {
		var funs []FunConfig
		if "" != ctx.Params["funcs"] {
			funs = getFunsFromParam(ctx.Params["funcs"])
			if funs == nil {
				funs = []FunConfig{
					{5}, {7},
				}
			}
		} else {
			funs = []FunConfig{
				{5}, {7},
			}
		}
		res, _ := json.Marshal(InitialResponse{
			MainImage: imageToBase64(altFlame(Config{
				Width: 300,
				Height: 300,
				Iterations: 1000 * 1000,
				Functions: funs,
			})),
			MainFormulas: funs,
			Formulas: GetWebFuncs(),
			ChildImages: RenderChildren(150, 150, 100 * 1000, funs),
		})
		ctx.SetHeader("Content-Type", "application/json", true)
		return string(res)
	})
	w.Get("/high-def", func(ctx *web.Context) string {
		var funs []FunConfig
		if "" != ctx.Params["funcs"] {
			funs = getFunsFromParam(ctx.Params["funcs"])
			if funs == nil {
				funs = []FunConfig{
					{5}, {7},
				}
			}
		} else {
			funs = []FunConfig{
				{5}, {7},
			}
		}
		return imageToBase64(altFlame(Config{
			Width: 1000,
			Height: 1000,
			Iterations: 10 * 1000 * 1000,
			Functions: funs,
		}))
	})
	w.Config.StaticDir = "public"
	w.Run("localhost:" + os.Getenv("PORT"))
}
