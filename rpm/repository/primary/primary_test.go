package primary

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMarshalPrimary(t *testing.T) {
	Convey("primary metadata should parse", t, func() {
		metadata, err := FromFile("../test_files/primary.xml.gz")
		So(err, ShouldBeNil)
		So(metadata, ShouldNotBeNil)
	})
}
