package main

import (
  "math"
)

type Variation func(float64, float64, float64, float64, float64, float64, float64, float64) (float64, float64)

func AllVariations() []Variation {
	return []Variation{f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11}
}

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

