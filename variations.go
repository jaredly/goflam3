package flame

import (
	"math"
)

type Variation func(float64, float64, float64, float64, float64, float64, float64, float64) (float64, float64)

type FullVar struct {
	Fn   Variation
	Text string
}

func DefaultParams() (float64, float64, float64, float64, float64, float64) {
  return 1, 2, 1, 1, 4, 5
}

func AllVariations() []FullVar {
	return []FullVar{
		{
			Fn: func(x, y, a, b, c, d, e, f float64) (float64, float64) {
				return x / 2, y / 2
			},
			Text: "x/2, y/2",
		},
		{
			Fn: func(x, y, a, b, c, d, e, f float64) (float64, float64) {
				return (x + 1) / 2, (y + 1) / 2
			},
			Text: "(x+1)/2, (y+1)/2",
		},
		{
			Fn: func(x, y, a, b, c, d, e, f float64) (float64, float64) {
				return (x + 1) / 2, y / 2
			},
			Text: "(x+1)/2, y/2",
		},
		{
			Fn: func(x, y, a, b, c, d, e, f float64) (float64, float64) {
				return x / 2, (y + 1) / 2
			},
			Text: "x/2, (y+1)/2",
		},
		{
			Fn: func(x, y, a, b, c, d, e, f float64) (float64, float64) {
				return (x - 1) / 2, y / 2
			},
			Text: "(x-1)/2, y/2",
		},
		{
			Fn: func(x, y, a, b, c, d, e, f float64) (float64, float64) {
				return x / 2, (y - 1) / 2
			},
			Text: "x/2, (y-1)/2",
		},
		{
			Fn: func(x, y, a, b, c, d, e, f float64) (float64, float64) {
				return x / 2, (y + 3) / 4
			},
			Text: "x/2, (y+3)/4",
		},
		{
			Fn: func(x, y, a, b, c, d, e, f float64) (float64, float64) {
				return math.Sin(x), math.Sin(y)
			},
			Text: "sin(x), sin(y)",
		},
		{
			Fn: func(x, y, a, b, c, d, e, f float64) (float64, float64) {
				return math.Cos(x), math.Cos(y)
			},
			Text: "cos(x), cos(y)",
		},
		{
			Fn: func(x, y, a, b, c, d, e, f float64) (float64, float64) {
				return math.Cos(a*x + b*y + c), math.Cos(d*x + e*y + f)
			},
			Text: "cos(ax+by+c), cos(dx+ey+f)",
		},
		{
			Fn: func(x, y, a, b, c, d, e, f float64) (float64, float64) {
				return math.Sin(a*x + b*y + c), math.Sin(d*x + e*y + f)
			},
			Text: "sin(ax+by+c), sin(dx+ey+f)",
		},
		{
			Fn: func(x, y, a, b, c, d, e, f float64) (float64, float64) {
				return math.Sin(a*x + b*y + c), math.Cos(d*x + e*y + f)
			},
			Text: "sin(ax+by+c), cos(dx+ey+f)",
		},
		{
			Fn: func(x, y, a, b, c, d, e, f float64) (float64, float64) {
				one := a*x + b*y + c
				two := d*x + e*y + f
				return math.Sin(one) * math.Sin(one), math.Sin(two) * math.Sin(two)
			},
			Text: "sin2(ax+by+c), sin2(dx+ey+f)",
		},
		{
			Fn: func(x, y, a, b, c, d, e, f float64) (float64, float64) {
				one := a*x + b*y + c
				two := d*x + e*y + f
				return math.Cos(one) * math.Cos(one), math.Cos(two) * math.Cos(two)
			},
			Text: "cos2(ax+by+c), cos(dx+ey+f)",
		},
		{
			Fn: func(x, y, a, b, c, d, e, f float64) (float64, float64) {
				one := a*x + b*y + c
				two := d*x + e*y + f
				return math.Sin(one) * math.Cos(one), math.Sin(two) * math.Cos(two)
			},
			Text: "sin*cos(ax+by+c), sin*cos(dx+ey+f)",
		},
		{
			Fn: func(x, y, a, b, c, d, e, f float64) (float64, float64) {
				a = -1.8
				b = -2.0
				c = -0.5
				d = -0.9
				return math.Sin(a*y) + c*math.Cos(a*x), math.Sin(b*x) + d*math.Cos(b*y)
			},
			Text: "sin(ay)+ccos(ax), sin(bx)+dcos(by)",
		},
		{
			Fn: func(x, y, a, b, c, d, e, f float64) (float64, float64) {
				  if x == 0 { x = 1 }
				  if y == 0 {y=1}
				return math.Sin(1/x), math.Sin(1/y)
			},
			Text: "sin(1/x), sin(1/y)",
		},
		{
			Fn: func(x, y, a, b, c, d, e, f float64) (float64, float64) {
				  if x == 0 { x = 1 }
				  if y == 0 { y = 1 }
				return math.Cos(1/x), math.Cos(1/y)
			},
			Text: "cos(1/x), cos(1/y)",
		},
		{
			Fn: func(x, y, a, b, c, d, e, f float64) (float64, float64) {
				  if x == 0 { x = 1 }
				  if y == 0 { y = 1 }
				return math.Sin(1/x)*math.Cos(1/x), math.Sin(1/x)*math.Cos(1/y)
			},
			Text: "sin.cos(1/x), sin.cos(1/y)",
		},
		/*
		{
			Fn: func(x, y, a, b, c, d, e, f float64) (float64, float64) {
				  if x == 0 { x = 1 }
				  a = math.Cos(x)
				  if a == 0 { a = 1 }
				  if y == 0 { y = 1 }
				  b = math.Cos(y)
				  if b == 0 { b = 1 }
				return math.Sin(x)/a, math.Sin(y)/b
			},
			Text: "tan(x), tan(y)",
		},
		*/
	}
}
