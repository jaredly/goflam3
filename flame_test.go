package flame

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRWMatrix(t *testing.T) {
	Convey("With a small matrix", t, func() {
		matrix := [][]Pixel{{
			{10, [20]int64{6, 7, 8, 9}},
			{20, [20]int64{3}},
		}, {
			{50, [20]int64{2, 3, 4}},
			{70, [20]int64{5, 6, 7}},
		}}
		Convey("It should write and read", func() {
			saveMatrix("test_mx.bin", &matrix)
			gotten := loadMatrix("test_mx.bin")
			So(*gotten, ShouldResemble, matrix)
		})
	})
}
