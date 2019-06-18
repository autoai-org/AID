package utility

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)


func TestRavenSpec(t *testing.T) {
	Convey("IsExists(path) should check if path exists", t, func() {
		InitRaven()
	})
}
