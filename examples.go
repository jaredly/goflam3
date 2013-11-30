package flame

import (
	"image"
	"image/color"
	"image/draw"
  /*
	bin "encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"io"
	"time"
  */
)

func RenderPreview(width, height, xs, ys int, fn Variation) *image.RGBA {
  return lines(width, height, matrix(xs, ys, fn))
}

func matrix(width, height int, fn Variation) [][]Point {
  data := make([][]Point, height)
	// these are our parameters
	var a, b, c, d, e, f float64
	a, b, c, d, e, f = DefaultParams()
  for y := range data {
    data[y] = make([]Point, width)
    for x := range data[y] {
      a, b := fn(frow(x, width), frow(y, height), a, b, c, d, e, f)
      data[y][x] = Point{a,b}
    }
  }
  return data
}

func interp(i, n int, p1, p2 Point) (Point, Point) {
  dx := (p2.X-p1.X)/float64(n)
  dy := (p2.Y-p1.Y)/float64(n)
  z := float64(i)
  return Point{p1.X + dx*z, p1.Y + dy*z}, Point{p1.X + dx*(z+1), p1.Y + dy*(z+1)}
}

// width & height correspond to the output image
func lines(width, height int, data [][]Point) *image.RGBA {
  img := image.NewRGBA(image.Rect(0, 0, width, height))
  ys := len(data)
  xs := len(data[0])
  parts := 4
  for i := 0; i < parts + 1; i++ {
    for y := range data {
      for x := range data[y] {
        // draw down
        p1, p2 := interp(i, parts, Point{frow(x, xs), frow(y, ys)}, data[y][x])
        line(width, height, p1, p2, color.RGBA{0, 0, 0, uint8(255*i/parts)}, img)
      }
    }
  }
  return img
}

// width & height correspond to the output image
func oldlines(width, height int, data [][]Point) *image.RGBA {
  img := image.NewRGBA(image.Rect(0, 0, width, height))
  ys := len(data)
  xs := len(data[0])
  for y := range data {
    for x := range data[y] {
      // draw down
      o := Point{frow(x, xs), frow(y, ys)}
      line(width, height, o, data[y][x], color.RGBA{0, 0, 0, 255}, img)
    }
  }
  return img
}

func tow(x float64, w, margin int) int {
  return int((x + 1)/2 * float64(w-margin*2)) + margin
}

func frow(x int, w int) float64 {
  return float64(x*2)/float64(w) - 1
}

func line(width, height int, p1, p2 Point, c color.RGBA, img draw.Image) {
  // a := c.A
  x1 := tow(p1.X, width, width/10)
  y1 := tow(p1.Y, height, height/10)
  x2 := tow(p2.X, width, width/10)
  y2 := tow(p2.Y, height, height/10)
  bresneham(img, c, x1, y1, x2, y2)
}

func abs(i int) int {
  if i < 0 { return i*-1 }
  return i
}

// alg taken from wikipedia
func bresneham(image draw.Image, c color.Color, x0, y0, x1, y1 int) {
  dx := abs(x1-x0)
  dy := abs(y1-y0)
  var sx, sy int
  if x0 < x1 {
    sx = 1
  } else {
    sx = -1
  }
  if y0 < y1 {
    sy = 1
  } else{
    sy = -1
  }
  err := dx-dy

  for {
    image.Set(x0,y0,c)
    if x0 == x1 && y0 == y1 { return }
    e2 := 2*err
    if e2 > -dy {
      err = err - dy
      x0 = x0 + sx
    }
    if x0 == x1 && y0 == y1 {
      image.Set(x0,y0, c)
      return
    }
    if e2 <  dx {
      err = err + dx
      y0 = y0 + sy
    }
  }
}

