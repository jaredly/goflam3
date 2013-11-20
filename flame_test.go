package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestReadWrite(t *testing.T) {
	Convey("With a small list of points", t, func() {
		Convey("It should write and read just fine", func() {
			points := []Point{
				{0, 0},
				{1, 1},
				{2, 2},
			}
			storeData("test_save.bin", &points)
			gotten := loadData("test_save.bin")
			So(*gotten, ShouldResemble, points)
		})
	})
}

func TestRWMatrix(t *testing.T) {
	Convey("With a small matrix", t, func() {
		matrix := [][]Pixel{{
			{10, [20]int32{6, 7, 8, 9}},
			{20, [20]int32{3}},
		}, {
			{50, [20]int32{2, 3, 4}},
			{70, [20]int32{5, 6, 7}},
		}}
		Convey("It should write and read", func() {
			saveMatrix("test_mx.bin", &matrix)
			gotten := loadMatrix("test_mx.bin")
			So(*gotten, ShouldResemble, matrix)
		})
	})
}
