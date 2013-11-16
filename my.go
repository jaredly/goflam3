package main

import (
  "os"
  "fmt"
  "strconv"
  "image"
  "image/png"
  "image/color"
  "math"
  "math/rand"
  "time"
)

type Variation func(float64, float64, float64, float64, float64, float64, float64, float64) (float64, float64)

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

func f6(x, y, a, b, c, d, e, f float64) (float64, float64) {
  return math.Sin(a*x + b*y + c), math.Sin(d*x + e*y + f)
}

func f7(x, y, a, b, c, d, e, f float64) (float64, float64) {
  return math.Sin(a*x + b*y + c), math.Cos(d*x + e*y + f)
}

func f8(x, y, a, b, c, d, e, f float64) (float64, float64) {
  one := a*x+b*y+c
  two := d*x + e*y + f
  return math.Sin(one) * math.Sin(one), math.Sin(two) * math.Sin(two)
}

func f9(x, y, a, b, c, d, e, f float64) (float64, float64) {
  one := a*x+b*y+c
  two := d*x + e*y + f
  return math.Cos(one) * math.Cos(one), math.Cos(two) * math.Cos(two)
}

func f10(x, y, a, b, c, d, e, f float64) (float64, float64) {
  one := a*x+b*y+c
  two := d*x + e*y + f
  return math.Sin(one) * math.Cos(one), math.Sin(two) * math.Cos(two)
}

func f11(x, y, a, b, c, d, e, f float64) (float64, float64) {
  a = -1.8
  b = -2.0
  c = -0.5
  d = -0.9
  return math.Sin(a * y) + c*math.Cos(a*x), math.Sin(b*x)+d*math.Cos(b*y)
}

type Point struct {
  x, y float64
}

func flame(width, height, iters int, usefuncs []int) *image.RGBA {
  start := time.Now()
  allfuncs := []Variation{f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10}

  fmt.Println("Flaming")
  data := generate(iters, usefuncs, allfuncs)
  fmt.Println("Generate time", time.Since(start).String())

  other := time.Now()
  image := render(width, height, data)
  fmt.Println("Render time", time.Since(other).String())
  return image
}

func generate(iters int, usefuncs []int, variations []Variation) *[]Point {
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
    x, y = variations[usefuncs[fi]](x, y, a, b, c, d, e, f)
    if at < 20 {
      continue
    }
    data[at] = Point{x, y}
  }
  return &data
}

// get the third-largest value in a matrix. I can probably do this better
func equalize_copy(values int, arr [][]int, maxp, minp float64) [][]int {
  mx := 0
  total := len(arr) * len(arr[0])
  for _, row := range arr {
    for _, v := range row {
      if mx < v {
        mx = v
      }
    }
  }
  hist := make([]int, mx + 1)
  for _, row := range arr {
    for _, v := range row {
      hist[v] += 1
    }
  }
  min := 0
  max := 0
  top := int(float64(total) * maxp)
  bottom := int(float64(total) * minp)
  count := 0
  // fmt.Printf("Eq", mx, total, top, bottom)
  // fmt.Printf("Hist", hist)
  for i, n := range hist {
    count += n
    if count > bottom && min == 0 {
      min = i
    }
    if count > top {
      max = i
      break
    }
  }
  // fmt.Printf("Eq", mx, total, min, max)
  for i := range arr {
    for j := range arr[i] {
      if max == min {
        arr[i][j] = 0
        continue
      }
      if arr[i][j] > max {
        arr[i][j] = max
      }
      if arr[i][j] < min {
        arr[i][j] = min
      }
      arr[i][j] = values * (arr[i][j] - min) / (max - min)
    }
  }
  return arr
}

// get the third-largest value in a matrix. I can probably do this better
func equalize(arr *[][]int, values int, maxp, minp float64) {
  mx := 0
  total := len(*arr) * len((*arr)[0])
  for _, row := range *arr {
    for _, v := range row {
      if mx < v {
        mx = v
      }
    }
  }
  hist := make([]int, mx + 1)
  for _, row := range *arr {
    for _, v := range row {
      hist[v] += 1
    }
  }
  min := 0
  max := 0
  top := int(float64(total) * maxp)
  bottom := int(float64(total) * minp)
  count := 0
  // fmt.Printf("Eq", mx, total, top, bottom)
  // fmt.Printf("Hist", hist)
  for i, n := range hist {
    count += n
    if count > bottom && min == 0 {
      min = i
    }
    if count > top {
      max = i
      break
    }
  }
  // fmt.Printf("Eq", mx, total, min, max)
  for i := range *arr {
    for j := range (*arr)[i] {
      if max == min {
        (*arr)[i][j] = 0
        continue
      }
      if (*arr)[i][j] > max {
        (*arr)[i][j] = max
      }
      if (*arr)[i][j] < min {
        (*arr)[i][j] = min
      }
      (*arr)[i][j] = values * ((*arr)[i][j] - min) / (max - min)
    }
  }
}

func find_max(mx *[][]float64) float64 {
  max := 0.0
  for y := range *mx {
    for x, v := range (*mx)[y] {
      if v > 0 {
        (*mx)[y][x] = math.Log(v)/v
      }
      if max < v {
	max = v
      }
    }
  }
  return max
}

func make_histogram(mx *[][]float64, by float64, max float64) *[]int {
  hist := make([]int, int(by) + 1)
  for y := range *mx {
    for x := range (*mx)[y] {
      hist[int((*mx)[y][x] * by / max)] += 1
    }
  }
  return &hist
}

func get_max_min(hist *[]int, total int, maxp, minp float64) (int, int) {
  min := 0
  max := 0
  top := int(float64(total) * maxp)
  bottom := int(float64(total) * minp)
  count := 0
  // fmt.Printf("Eq", mx, total, top, bottom)
  // fmt.Printf("Hist", hist)
  for i, n := range *hist {
    count += n
    if count > bottom && min == 0 {
      min = i
    }
    if count > top {
      max = i
      break
    }
  }
  return max, min
}

// get the third-largest value in a matrix. I can probably do this better
func equalize_log(arr *[][]float64, values int, maxp, minp float64) {
  mx := find_max(arr)

  by := 255 * 10.0
  total := len(*arr) * len((*arr)[0])
  hist := make_histogram(arr, by, mx)
  max, min := get_max_min(hist, total, maxp, minp)
  // fmt.Printf("Eq", mx, total, min, max)
  for i := range *arr {
    for j := range (*arr)[i] {
      if max == min {
	(*arr)[i][j] = 0
	continue
      }
      v := int((*arr)[i][j] * by / mx)
      if v > max {
	(*arr)[i][j] = float64(max)
      }
      if v < min {
	(*arr)[i][j] = float64(min)
      }
      (*arr)[i][j] = float64(values * (v - min) / (max - min))
    }
  }
}

// data the data and render it within certain dimentions
func render(width, height int, data *[]Point) *image.RGBA {
  fmt.Println("Render")
  // now render
  mx := make([][]int, height)
  for y := range mx {
    mx[y] = make([]int, width)
    for x := range mx[y] {
      mx[y][x] = 0
    }
  }
  for _, v := range *data {
    mx[int((v.y+1)/2*float64(height-1))][int((v.x+1)/2*float64(width-1))] += 1
  }
  equalize(&mx, 255, .995, .0005)
  image := image.NewRGBA(image.Rect(0, 0, width, height))
  // now write the values to an image, equalized by the 3rd-brightest point
  for x, row := range mx {
    for y, v := range row {
      val := uint8(v)
      image.Set(x, y, color.RGBA{val, val * 100 / 255, val, 255})
    }
  }
  return image
}

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

func main() {
  w := 800
  h := 800
  i := 10000000
  writeit(w, h, i, []int{7,3})
  // writeit(w, h, i, []int{3,5})
  // callThemAll(w, h, i, 3, 12)
  // allCombos(w, h, i, 0, 7)
}
