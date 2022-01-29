package enum

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestOptionPositionEnum(t *testing.T) {
	convey.Convey("枚举类测试", t, func() {
		convey.So(Option_Position_Buyer.Name(), convey.ShouldEqual, "buyer")
		convey.So(Option_Position_Buyer.Desc(), convey.ShouldEqual, "买方")

		convey.So(Option_Position_Seller.Name(), convey.ShouldEqual, "seller")
		convey.So(Option_Position_Seller.Desc(), convey.ShouldEqual, "卖方")

		convey.So(len(OptionPositionEnumList()), convey.ShouldEqual, 2)

		p, err := GetOptionPositionEnumByName("seller")
		convey.So(err, convey.ShouldBeNil)
		convey.So(p, convey.ShouldEqual, Option_Position_Seller)
	})
}
