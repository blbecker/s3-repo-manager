package repomd

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMarshalRepomd(t *testing.T) {
	Convey("repomd should parse", t, func() {

		repomd, err := FromFile("../test_files/repomd.xml")

		So(err, ShouldBeNil)
		So(len(repomd.Data), ShouldEqual, 3)
	})
}
