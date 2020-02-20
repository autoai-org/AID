package utility

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWebUIUtilitySpec(t *testing.T) {
	Convey("WebUI() should Install Dashboard App", t, func() {
		So(InstallWebUi(), ShouldEqual, true)
	})
}
