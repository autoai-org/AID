package cvpm_core_test

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/unarxiv/cvpm/pkg/utility"
)

func TestSystemUtilitySpec(t *testing.T) {
	systemInfo := utility.GetSystemInfo()
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

func TestCVPMUtilitySpec(t *testing.T) {
	Convey("IsExists should check if path exists", t, func() {
		So(utility.IsExists("utility_test.go"), ShouldEqual, true)
		So(utility.IsExists("utility_test_non_exist.go"), ShouldEqual, false)
	})
	Convey("Get Homedir should return current users directory", t, func() {
		So(reflect.TypeOf(utility.GetHomeDir()).String(), ShouldEqual, "string")
	})
	Convey("IsStringInSlice should check if a string exists in a list", t, func() {
		baseStringList := []string{"a", "b"}
		So(utility.IsStringInSlice("a", baseStringList), ShouldEqual, true)
		So(utility.IsStringInSlice("c", baseStringList), ShouldEqual, false)
	})
}
