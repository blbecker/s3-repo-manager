package filelists

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMarshalPrimary(t *testing.T) {
	Convey("filelists should parse", t, func() {
		filelists, err := FromFile("../test_files/filelists.xml.gz")
		So(err, ShouldBeNil)
		So(filelists, ShouldNotBeNil)
	})
}