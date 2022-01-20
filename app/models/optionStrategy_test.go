package models

import (
	"testing"

	"github.com/feigme/fmgr-go/app/enum"
	"github.com/smartystreets/goconvey/convey"
)

func TestRollPut(t *testing.T) {
	convey.Convey("测试Roll put策略", t, func() {
		convey.Convey("主逻辑", func() {
			st := NewOptionStrategy(enum.OST_ROLLING_PUT)

			trade, err := NewOptionTrade("futu211210P47000", enum.SHORT, "1.6")
			convey.So(err, convey.ShouldBeNil)

			st.RollPut(trade, "3.0", "211217", "4.1")
		})
	})
}
