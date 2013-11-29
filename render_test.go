package flame

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFindMax(t *testing.T) {
	Convey("Given a 2x float matrix", t, func() {
		Convey("It should find the max", func() {
			in := [][]float64{
				{0, 0, 1},
				{2, 3, 4},
			}
			So(find_max(&in), ShouldEqual, 4.0)
		})
	})
}
