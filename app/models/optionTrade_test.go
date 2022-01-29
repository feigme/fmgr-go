package models

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/feigme/fmgr-go/app/enum"
)

func TestNewOptionTrade(t *testing.T) {
	Convey("单腿策略", t, func() {
		Convey("卖call", func() {
			trade, err := NewOptionTrade("TCH211230C500000", enum.Option_Position_Seller, "8.0")
			So(err, ShouldBeNil)

			// 交易信息
			So(trade.CreateTime, ShouldNotBeNil)
			So(trade.UpdateTime, ShouldNotBeNil)
			So(trade.Position, ShouldEqual, "seller")
			So(trade.SellPrice, ShouldEqual, "8.00")
			So(trade.Count, ShouldEqual, int64(-1))
			So(trade.Status, ShouldEqual, enum.OPTION_STATUS_HAVING.Name())
			So(trade.Premium, ShouldEqual, "800.00")
			So(trade.BuyPrice, ShouldBeEmpty)
			So(trade.Profit, ShouldBeEmpty)
			So(trade.ProfitRate, ShouldBeEmpty)
		})

		Convey("卖call后，操作", func() {
			Convey("失效", func() {
				trade, err := NewOptionTrade("TCH211230C500000", enum.Option_Position_Seller, "8.0")
				So(err, ShouldBeNil)

				trade.Invalid()

				// 交易信息
				So(trade.CreateTime, ShouldNotBeNil)
				So(trade.UpdateTime, ShouldNotBeNil)
				So(trade.Position, ShouldEqual, "seller")
				So(trade.SellPrice, ShouldEqual, "8.00")
				So(trade.Count, ShouldEqual, int64(-1))
				So(trade.Status, ShouldEqual, enum.OPTION_STATUS_INVALID.Name())
				So(trade.Premium, ShouldEqual, "800.00")
				So(trade.BuyPrice, ShouldEqual, "0.00")
				So(trade.Profit, ShouldEqual, "800.00")
				So(trade.ProfitRate, ShouldEqual, "1.00")
			})

			Convey("行权", func() {
				trade, err := NewOptionTrade("TCH211230C500000", enum.Option_Position_Seller, "8.0")
				So(err, ShouldBeNil)

				trade.Exercise()

				// 交易信息
				So(trade.CreateTime, ShouldNotBeNil)
				So(trade.UpdateTime, ShouldNotBeNil)
				So(trade.Position, ShouldEqual, "seller")
				So(trade.SellPrice, ShouldEqual, "8.00")
				So(trade.Count, ShouldEqual, int64(-1))
				So(trade.Status, ShouldEqual, enum.OPTION_STATUS_EXERCISE.Name())
				So(trade.Premium, ShouldEqual, "800.00")
				So(trade.BuyPrice, ShouldEqual, "0.00")
				So(trade.Profit, ShouldEqual, "800.00")
				So(trade.ProfitRate, ShouldEqual, "1.00")
			})

			Convey("平仓", func() {
				Convey("盈利", func() {
					trade, err := NewOptionTrade("TCH211230C500000", enum.Option_Position_Seller, "8.0")
					So(err, ShouldBeNil)

					trade.Close("1.12")

					// 交易信息
					So(trade.CreateTime, ShouldNotBeNil)
					So(trade.UpdateTime, ShouldNotBeNil)
					So(trade.Position, ShouldEqual, "seller")
					So(trade.SellPrice, ShouldEqual, "8.00")
					So(trade.Count, ShouldEqual, int64(-1))
					So(trade.Status, ShouldEqual, enum.OPTION_STATUS_CLOSE.Name())
					So(trade.Premium, ShouldEqual, "800.00")
					So(trade.BuyPrice, ShouldEqual, "1.12")
					So(trade.Profit, ShouldEqual, "688.00")
					So(trade.ProfitRate, ShouldEqual, "0.86")
				})

				Convey("亏损", func() {
					trade, err := NewOptionTrade("TCH211230C500000", enum.Option_Position_Seller, "8.0")
					So(err, ShouldBeNil)

					trade.Close("11.2")

					// 交易信息
					So(trade.CreateTime, ShouldNotBeNil)
					So(trade.UpdateTime, ShouldNotBeNil)
					So(trade.Position, ShouldEqual, "seller")
					So(trade.SellPrice, ShouldEqual, "8.00")
					So(trade.Count, ShouldEqual, int64(-1))
					So(trade.Status, ShouldEqual, enum.OPTION_STATUS_CLOSE.Name())
					So(trade.Premium, ShouldEqual, "800.00")
					So(trade.BuyPrice, ShouldEqual, "11.20")
					So(trade.Profit, ShouldEqual, "-320.00")
					So(trade.ProfitRate, ShouldEqual, "-0.40")
				})
			})
		})

		Convey("卖put", func() {
			trade, err := NewOptionTrade("TCH211230P450000", enum.Option_Position_Seller, "12.1")
			So(err, ShouldBeNil)

			// 交易信息
			So(trade.CreateTime, ShouldNotBeNil)
			So(trade.UpdateTime, ShouldNotBeNil)
			So(trade.Position, ShouldEqual, "seller")
			So(trade.SellPrice, ShouldEqual, "12.10")
			So(trade.Count, ShouldEqual, int64(-1))
			So(trade.Status, ShouldEqual, enum.OPTION_STATUS_HAVING.Name())
			So(trade.Premium, ShouldEqual, "1210.00")
			So(trade.BuyPrice, ShouldBeEmpty)
			So(trade.Profit, ShouldBeEmpty)
			So(trade.ProfitRate, ShouldBeEmpty)
		})

		Convey("卖put后，操作", func() {
			Convey("失效", func() {
				trade, err := NewOptionTrade("TCH211230P450000", enum.Option_Position_Seller, "12.1")
				So(err, ShouldBeNil)

				trade.Invalid()

				// 交易信息
				So(trade.CreateTime, ShouldNotBeNil)
				So(trade.UpdateTime, ShouldNotBeNil)
				So(trade.Position, ShouldEqual, "seller")
				So(trade.SellPrice, ShouldEqual, "12.10")
				So(trade.Count, ShouldEqual, int64(-1))
				So(trade.Status, ShouldEqual, enum.OPTION_STATUS_INVALID.Name())
				So(trade.Premium, ShouldEqual, "1210.00")
				So(trade.BuyPrice, ShouldEqual, "0.00")
				So(trade.Profit, ShouldEqual, "1210.00")
				So(trade.ProfitRate, ShouldEqual, "1.00")
			})

			Convey("行权", func() {
				trade, err := NewOptionTrade("TCH211230P450000", enum.Option_Position_Seller, "12.1")
				So(err, ShouldBeNil)

				trade.Exercise()

				// 交易信息
				So(trade.CreateTime, ShouldNotBeNil)
				So(trade.UpdateTime, ShouldNotBeNil)
				So(trade.Position, ShouldEqual, "seller")
				So(trade.SellPrice, ShouldEqual, "12.10")
				So(trade.Count, ShouldEqual, int64(-1))
				So(trade.Status, ShouldEqual, enum.OPTION_STATUS_EXERCISE.Name())
				So(trade.Premium, ShouldEqual, "1210.00")
				So(trade.BuyPrice, ShouldEqual, "0.00")
				So(trade.Profit, ShouldEqual, "1210.00")
				So(trade.ProfitRate, ShouldEqual, "1.00")
			})

			Convey("平仓", func() {
				Convey("盈利", func() {
					trade, err := NewOptionTrade("TCH211230P450000", enum.Option_Position_Seller, "12.1")
					So(err, ShouldBeNil)

					trade.Close("0.3")

					// 交易信息
					So(trade.CreateTime, ShouldNotBeNil)
					So(trade.UpdateTime, ShouldNotBeNil)
					So(trade.Position, ShouldEqual, "seller")
					So(trade.SellPrice, ShouldEqual, "12.10")
					So(trade.Count, ShouldEqual, int64(-1))
					So(trade.Status, ShouldEqual, enum.OPTION_STATUS_CLOSE.Name())
					So(trade.Premium, ShouldEqual, "1210.00")
					So(trade.BuyPrice, ShouldEqual, "0.30")
					So(trade.Profit, ShouldEqual, "1180.00")
					So(trade.ProfitRate, ShouldEqual, "0.98")
				})
				Convey("亏损", func() {
					trade, err := NewOptionTrade("TCH211230P450000", enum.Option_Position_Seller, "12.1")
					So(err, ShouldBeNil)

					trade.Close("16.1")

					// 交易信息
					So(trade.CreateTime, ShouldNotBeNil)
					So(trade.UpdateTime, ShouldNotBeNil)
					So(trade.Position, ShouldEqual, "seller")
					So(trade.SellPrice, ShouldEqual, "12.10")
					So(trade.Count, ShouldEqual, int64(-1))
					So(trade.Status, ShouldEqual, enum.OPTION_STATUS_CLOSE.Name())
					So(trade.Premium, ShouldEqual, "1210.00")
					So(trade.BuyPrice, ShouldEqual, "16.10")
					So(trade.Profit, ShouldEqual, "-400.00")
					So(trade.ProfitRate, ShouldEqual, "-0.33")
				})
			})
		})

		Convey("买call", func() {
			trade, err := NewOptionTrade("bili211217C70000", enum.Option_Position_Buyer, "0.34")
			So(err, ShouldBeNil)

			// 交易信息
			So(trade.CreateTime, ShouldNotBeNil)
			So(trade.UpdateTime, ShouldNotBeNil)
			So(trade.Position, ShouldEqual, "buyer")
			So(trade.BuyPrice, ShouldEqual, "0.34")
			So(trade.Count, ShouldEqual, int64(1))
			So(trade.Status, ShouldEqual, enum.OPTION_STATUS_HAVING.Name())
			So(trade.Premium, ShouldEqual, "-34.00")
			So(trade.SellPrice, ShouldBeEmpty)
			So(trade.Profit, ShouldBeEmpty)
			So(trade.ProfitRate, ShouldBeEmpty)
		})

		Convey("买call，操作", func() {
			Convey("失效", func() {
				trade, err := NewOptionTrade("bili211217C70000", enum.Option_Position_Buyer, "0.34")
				So(err, ShouldBeNil)

				trade.Invalid()

				// 验证期权基本信息
				So(trade.Code, ShouldEqual, "BILI211217C70000")
				So(trade.Stock, ShouldEqual, "BILI")
				So(trade.ExerciseDate, ShouldEqual, "211217")
				So(trade.Type, ShouldEqual, "C")
				So(trade.StrikePrice, ShouldEqual, "70000")
				So(trade.ContractSize, ShouldEqual, int64(100))

				// 交易信息
				So(trade.CreateTime, ShouldNotBeNil)
				So(trade.UpdateTime, ShouldNotBeNil)
				So(trade.Position, ShouldEqual, "buyer")
				So(trade.BuyPrice, ShouldEqual, "0.34")
				So(trade.Count, ShouldEqual, int64(1))
				So(trade.Status, ShouldEqual, enum.OPTION_STATUS_INVALID.Name())
				So(trade.Premium, ShouldEqual, "-34.00")
				So(trade.SellPrice, ShouldEqual, "0.00")
				So(trade.Profit, ShouldEqual, "-34.00")
				So(trade.ProfitRate, ShouldEqual, "-1.00")
			})

			Convey("行权", func() {
				trade, err := NewOptionTrade("bili211217C70000", enum.Option_Position_Buyer, "0.34")
				So(err, ShouldBeNil)

				trade.Exercise()

				// 验证期权基本信息
				So(trade.Code, ShouldEqual, "BILI211217C70000")
				So(trade.Stock, ShouldEqual, "BILI")
				So(trade.ExerciseDate, ShouldEqual, "211217")
				So(trade.Type, ShouldEqual, "C")
				So(trade.StrikePrice, ShouldEqual, "70000")
				So(trade.ContractSize, ShouldEqual, int64(100))

				// 交易信息
				So(trade.CreateTime, ShouldNotBeNil)
				So(trade.UpdateTime, ShouldNotBeNil)
				So(trade.Position, ShouldEqual, "buyer")
				So(trade.BuyPrice, ShouldEqual, "0.34")
				So(trade.Count, ShouldEqual, int64(1))
				So(trade.Status, ShouldEqual, enum.OPTION_STATUS_EXERCISE.Name())
				So(trade.Premium, ShouldEqual, "-34.00")
				So(trade.SellPrice, ShouldEqual, "0.00")
				So(trade.Profit, ShouldEqual, "-34.00")
				So(trade.ProfitRate, ShouldEqual, "-1.00")
			})

			Convey("平仓", func() {
				Convey("盈利", func() {
					trade, err := NewOptionTrade("bili211217C70000", enum.Option_Position_Buyer, "0.34")
					So(err, ShouldBeNil)

					trade.Close("5.1")

					// 验证期权基本信息
					So(trade.Code, ShouldEqual, "BILI211217C70000")
					So(trade.Stock, ShouldEqual, "BILI")
					So(trade.ExerciseDate, ShouldEqual, "211217")
					So(trade.Type, ShouldEqual, "C")
					So(trade.StrikePrice, ShouldEqual, "70000")
					So(trade.ContractSize, ShouldEqual, int64(100))

					// 交易信息
					So(trade.CreateTime, ShouldNotBeNil)
					So(trade.UpdateTime, ShouldNotBeNil)
					So(trade.Position, ShouldEqual, "buyer")
					So(trade.BuyPrice, ShouldEqual, "0.34")
					So(trade.Count, ShouldEqual, int64(1))
					So(trade.Status, ShouldEqual, enum.OPTION_STATUS_CLOSE.Name())
					So(trade.Premium, ShouldEqual, "-34.00")
					So(trade.SellPrice, ShouldEqual, "5.10")
					So(trade.Profit, ShouldEqual, "476.00")
					So(trade.ProfitRate, ShouldEqual, "14.00")
				})

				Convey("亏损", func() {
					trade, err := NewOptionTrade("bili211217C70000", enum.Option_Position_Buyer, "0.34")
					So(err, ShouldBeNil)

					trade.Close("0.1")

					// 验证期权基本信息
					So(trade.Code, ShouldEqual, "BILI211217C70000")
					So(trade.Stock, ShouldEqual, "BILI")
					So(trade.ExerciseDate, ShouldEqual, "211217")
					So(trade.Type, ShouldEqual, "C")
					So(trade.StrikePrice, ShouldEqual, "70000")
					So(trade.ContractSize, ShouldEqual, int64(100))

					// 交易信息
					So(trade.CreateTime, ShouldNotBeNil)
					So(trade.UpdateTime, ShouldNotBeNil)
					So(trade.Position, ShouldEqual, "buyer")
					So(trade.BuyPrice, ShouldEqual, "0.34")
					So(trade.Count, ShouldEqual, int64(1))
					So(trade.Status, ShouldEqual, enum.OPTION_STATUS_CLOSE.Name())
					So(trade.Premium, ShouldEqual, "-34.00")
					So(trade.SellPrice, ShouldEqual, "0.10")
					So(trade.Profit, ShouldEqual, "-24.00")
					So(trade.ProfitRate, ShouldEqual, "-0.71")
				})
			})
		})

		Convey("买put", func() {
			trade, err := NewOptionTrade("futu211210P47000", enum.Option_Position_Buyer, "1.6")
			So(err, ShouldBeNil)

			// 交易信息
			So(trade.CreateTime, ShouldNotBeNil)
			So(trade.UpdateTime, ShouldNotBeNil)
			So(trade.Position, ShouldEqual, "buyer")
			So(trade.BuyPrice, ShouldEqual, "1.60")
			So(trade.Count, ShouldEqual, int64(1))
			So(trade.Status, ShouldEqual, enum.OPTION_STATUS_HAVING.Name())
			So(trade.Premium, ShouldEqual, "-160.00")
			So(trade.SellPrice, ShouldBeEmpty)
			So(trade.Profit, ShouldBeEmpty)
			So(trade.ProfitRate, ShouldBeEmpty)
		})

		Convey("买put后，操作", func() {
			Convey("失效", func() {
				trade, err := NewOptionTrade("futu211210P47000", enum.Option_Position_Buyer, "1.6")
				So(err, ShouldBeNil)

				trade.Invalid()

				// 交易信息
				So(trade.CreateTime, ShouldNotBeNil)
				So(trade.UpdateTime, ShouldNotBeNil)
				So(trade.Position, ShouldEqual, "buyer")
				So(trade.BuyPrice, ShouldEqual, "1.60")
				So(trade.Count, ShouldEqual, int64(1))
				So(trade.Status, ShouldEqual, enum.OPTION_STATUS_INVALID.Name())
				So(trade.Premium, ShouldEqual, "-160.00")
				So(trade.SellPrice, ShouldEqual, "0.00")
				So(trade.Profit, ShouldEqual, "-160.00")
				So(trade.ProfitRate, ShouldEqual, "-1.00")
			})

			Convey("行权", func() {
				trade, err := NewOptionTrade("futu211210P47000", enum.Option_Position_Buyer, "1.6")
				So(err, ShouldBeNil)

				trade.Exercise()

				// 交易信息
				So(trade.CreateTime, ShouldNotBeNil)
				So(trade.UpdateTime, ShouldNotBeNil)
				So(trade.Position, ShouldEqual, "buyer")
				So(trade.BuyPrice, ShouldEqual, "1.60")
				So(trade.Count, ShouldEqual, int64(1))
				So(trade.Status, ShouldEqual, enum.OPTION_STATUS_EXERCISE.Name())
				So(trade.Premium, ShouldEqual, "-160.00")
				So(trade.SellPrice, ShouldEqual, "0.00")
				So(trade.Profit, ShouldEqual, "-160.00")
				So(trade.ProfitRate, ShouldEqual, "-1.00")
			})

			Convey("平仓", func() {
				Convey("盈利", func() {
					trade, err := NewOptionTrade("futu211210P47000", enum.Option_Position_Buyer, "1.6")
					So(err, ShouldBeNil)

					trade.Close("3")

					// 交易信息
					So(trade.CreateTime, ShouldNotBeNil)
					So(trade.UpdateTime, ShouldNotBeNil)
					So(trade.Position, ShouldEqual, "buyer")
					So(trade.BuyPrice, ShouldEqual, "1.60")
					So(trade.Count, ShouldEqual, int64(1))
					So(trade.Status, ShouldEqual, enum.OPTION_STATUS_CLOSE.Name())
					So(trade.Premium, ShouldEqual, "-160.00")
					So(trade.SellPrice, ShouldEqual, "3.00")
					So(trade.Profit, ShouldEqual, "140.00")
					So(trade.ProfitRate, ShouldEqual, "0.87")
				})

				Convey("亏损", func() {
					trade, err := NewOptionTrade("futu211210P47000", enum.Option_Position_Buyer, "1.6")
					So(err, ShouldBeNil)

					trade.Close("0.6")

					// 交易信息
					So(trade.CreateTime, ShouldNotBeNil)
					So(trade.UpdateTime, ShouldNotBeNil)
					So(trade.Position, ShouldEqual, "buyer")
					So(trade.BuyPrice, ShouldEqual, "1.60")
					So(trade.Count, ShouldEqual, int64(1))
					So(trade.Status, ShouldEqual, enum.OPTION_STATUS_CLOSE.Name())
					So(trade.Premium, ShouldEqual, "-160.00")
					So(trade.SellPrice, ShouldEqual, "0.60")
					So(trade.Profit, ShouldEqual, "-100.00")
					So(trade.ProfitRate, ShouldEqual, "-0.62")
				})
			})
		})
	})

}
