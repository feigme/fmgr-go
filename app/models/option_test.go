package models

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewOption(t *testing.T) {
	Convey("期权code匹配", t, func() {
		Convey("美股期权", func() {
			option, err := NewOption("NIO211126P40000")
			So(err, ShouldBeNil)
			So(option.Code, ShouldEqual, "NIO211126P40000")
			So(option.Stock, ShouldEqual, "NIO")
			So(option.ExerciseDate, ShouldEqual, "211126")
			So(option.Type, ShouldEqual, "P")
			So(option.StrikePrice, ShouldEqual, "40000")
			So(option.ContractSize, ShouldEqual, int64(100))
		})

		Convey("港股期权", func() {
			option, err := NewOption("TCH211230P440000")
			So(err, ShouldBeNil)
			So(option.Code, ShouldEqual, "TCH211230P440000")
			So(option.Stock, ShouldEqual, "TCH")
			So(option.ExerciseDate, ShouldEqual, "211230")
			So(option.Type, ShouldEqual, "P")
			So(option.StrikePrice, ShouldEqual, "440000")
			So(option.ContractSize, ShouldEqual, int64(100))
		})

		Convey("大小写匹配", func() {
			option, err := NewOption("nio211126p40000")
			So(err, ShouldBeNil)
			So(option.Code, ShouldEqual, "NIO211126P40000")
			So(option.Stock, ShouldEqual, "NIO")
			So(option.ExerciseDate, ShouldEqual, "211126")
			So(option.Type, ShouldEqual, "P")
			So(option.StrikePrice, ShouldEqual, "40000")
			So(option.ContractSize, ShouldEqual, int64(100))
		})

		Convey("标的code过长", func() {
			_, err := NewOption("neeio20211126T40000")
			So(err, ShouldNotBeNil)
		})

		Convey("行权日格式错误", func() {
			_, err := NewOption("nio20211126T40000")
			So(err, ShouldNotBeNil)
		})

		Convey("期权类型错误", func() {
			_, err := NewOption("nio211126T40000")
			So(err, ShouldNotBeNil)
		})
	})

}
