package utility

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)


func TestCVPMUtilitySpec(t *testing.T) {
	Convey("IsExists() should check if path exists", t, func() {
		So(IsExists("utils_test.go"), ShouldEqual, true)
		So(IsExists("utils_test_non_exist.go"), ShouldEqual, false)
	})
	Convey("GetHomedir() should return current users directory", t, func() {
		So(reflect.TypeOf(GetHomeDir()).String(), ShouldEqual, "string")
	})
	Convey("IsStringInSlice() should check if a string exists in a list", t, func() {
		baseStringList := []string{"a", "b"}
		So(IsStringInSlice("a", baseStringList), ShouldEqual, true)
		So(IsStringInSlice("c", baseStringList), ShouldEqual, false)
	})
}
