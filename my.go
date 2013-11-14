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

// get the third-largest value in a matrix. I can probably do this better
func equalize(values int, arr [][]int) [][]int {
  mx := 0
  total := len(arr) * len(arr[0])
  for _, row := range arr {
    for x, v := range row {
      if v > 0 {
        row[x] = int(math.Log(float64(row[x])))
      }
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
  top := int(float64(total) * .9995)
  bottom := 0//int(float64(total) * .0005)
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

/*
  (x, y) = a random point in the biunit square
  iterate {
  i = a random integer from 0 to n  1 inclusive
  (x, y) = Fi(x, y)
  plot (x, y) except during the first 20 iterations
  }
*/

func flame(width, height, iters int, usefuncs []int) *image.RGBA {
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
  allfuncs := [11]func(float64, float64, float64, float64, float64, float64, float64, float64) (float64, float64){f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10}
  // and the F_i s that we'll be using
  funcs := make([]func(float64, float64, float64, float64, float64, float64, float64, float64) (float64, float64), len(usefuncs))
  for i, v := range(usefuncs) {
    funcs[i] = allfuncs[v]
  }
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
  mx = equalize(255, mx)
  m := image.NewRGBA(image.Rect(0, 0, width, height))
  // now write the values to an image, equalized by the 3rd-brightest point
  for x, row := range mx {
    for y, v := range row {
      val := uint8(v)
      m.Set(x, y, color.RGBA{val, val * 100 / 255, val, 255})
    }
  }
  return m
}

func writeit(w, h, i int, use []int) {
  m := flame(w, h, i, use)
  name := "smallish-"
  num := ""
  for _, z := range use {
    num = strconv.Itoa(z) + num
  }
  toimg, _ := os.Create(name + num + ".png")
  defer toimg.Close()

  png.Encode(toimg, m)
}

func allCombos(w, h, i, upto int) {
  all := int(math.Pow(2, float64(upto)))
  for z := 7; z < all; z++ {
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
  w := 400
  h := 400
  i := 10000000
  allCombos(w, h, i, 7)
  /// writeit(w, h, i, []int{3,5})
}
