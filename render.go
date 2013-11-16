package main

import (
	"math"
  "image"
	"image/color"
)

// data the data and render it within certain dimentions
func render(width, height int, data *[]Point) *image.RGBA {
	// now render
	mx := make([][]int, height)
	for y := range mx {
		mx[y] = make([]int, width)
		for x := range mx[y] {
			mx[y][x] = 0
		}
	}
	for _, v := range *data {
		mx[int((v.Y+1)/2*float64(height-1))][int((v.X+1)/2*float64(width-1))] += 1
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

