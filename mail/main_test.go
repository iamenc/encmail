package mail

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

// TestSum 测试单元
func TestSum(t *testing.T) {
	convey.Convey("oktesttest", t, func() {
		convey.Convey("a", func() {
			convey.So(sum(1, 2, 3, 4), convey.ShouldAlmostEqual, 10)
		})
		convey.Convey("b", func() {
			convey.So(sum(2, 2, 3, 4), convey.ShouldAlmostEqual, 11)
		})
		convey.Convey("c", func() {
			convey.So(sum(1, 2, 4), convey.ShouldAlmostEqual, 7)
		})
	})
}
