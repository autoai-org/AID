package cvpm_core_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRepositorySpec(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given some integer with a starting value", t, func() {
		x := 1
		So(x, ShouldEqual, 1)
		Convey("When the integer is incremented", func() {
			x++

			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 2)
			})
		})
	})
}
