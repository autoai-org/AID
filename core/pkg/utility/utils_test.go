package utility

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCVPMUtilitySpec(t *testing.T) {
	Convey("IsExists(path) should check if path exists", t, func() {
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
	Convey("IsRoot() should return false if current user is not administrator", t, func() {
		So(IsRoot(), ShouldEqual, false)
	})
	Convey("GetDirSizeMB(path string) should return the size of a folder", t, func() {
		So(reflect.TypeOf(GetDirSizeMB(".")).String(), ShouldEqual, "float64")
	})
	Convey("IsPortOpen(port string, timeout time.Duration) should return if the port is available", t, func() {
		So(reflect.TypeOf(IsPortOpen("8080", defaultTimeout)).String(), ShouldEqual, "bool")
		So(IsPortOpen("10500", defaultTimeout), ShouldEqual, true)
	})
	Convey("FindNextOpenPort(port string) should return the next open port", t, func() {
		So(reflect.TypeOf(FindNextOpenPort(8080)).String(), ShouldEqual, "string")
		So(FindNextOpenPort(10500), ShouldEqual, "10500")
	})
	Convey("ReadFileContent(filename string) should return the content of the requested file", t, func() {
		So(reflect.TypeOf(ReadFileContent("non_exist.go")).String(), ShouldEqual, "string")
		So(ReadFileContent("non_exist.go"), ShouldEqual, "Read non_exist.go Failed!")
		So(reflect.TypeOf(ReadFileContent("utils.go")).String(), ShouldEqual, "string")
	})
	Convey("IsPathExists(path string) should return true if path exists", t, func() {
		exists, _ := IsPathExists("./")
		So(reflect.TypeOf(exists).String(), ShouldEqual, "bool")
		So(exists, ShouldEqual, true)
		exists, _ = IsPathExists("./non")
		So(reflect.TypeOf(exists).String(), ShouldEqual, "bool")
		So(exists, ShouldEqual, false)
	})
	Convey("CreateFolderIfNotExist(folderPath string) should create the folder if not exists", t, func() {
		exists, _ := IsPathExists("./non")
		So(reflect.TypeOf(exists).String(), ShouldEqual, "bool")
		So(exists, ShouldEqual, false)
		CreateFolderIfNotExist("./non")
		exists, _ = IsPathExists("./non")
		So(reflect.TypeOf(exists).String(), ShouldEqual, "bool")
		So(exists, ShouldEqual, true)
	})
}
