package utility

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSystemUtilitySpec(t *testing.T) {
	systemInfo := GetSystemInfo()
	Convey("Type of CPU name is 'string'", t, func() {
		So(reflect.TypeOf(systemInfo.CpuName).String(), ShouldEqual, "string")
	})
	Convey("Type of Memory is uint64", t, func() {
		So(reflect.TypeOf(systemInfo.Memory).String(), ShouldEqual, "uint64")
	})
	Convey("Type of Platform is 'string'", t, func() {
		So(reflect.TypeOf(systemInfo.Platform).String(), ShouldEqual, "string")
	})
	Convey("Type of Os is 'string'", t, func() {
		So(reflect.TypeOf(systemInfo.Os).String(), ShouldEqual, "string")
	})
	Convey("Type of PlatformVersion is 'string'", t, func() {
		So(reflect.TypeOf(systemInfo.PlatformVersion).String(), ShouldEqual, "string")
	})
}
