package models

import (
	"testing"

	"github.com/feigme/fmgr-go/app/enum"
	"github.com/stretchr/testify/require"
)

func TestNewOptionTrade_卖call(t *testing.T) {
	option, err := NewOption("TCH211230C500000")
	require.NoError(t, err)
	trade, err := NewOptionTrade(option, enum.SHORT, "8.0")
	require.NoError(t, err)

	// 验证期权基本信息
	require.Equal(t, "TCH211230C500000", trade.Code)
	require.Equal(t, "TCH", trade.Stock)
	require.Equal(t, "211230", trade.ExerciseDate)
	require.Equal(t, "C", trade.Type)
	require.Equal(t, "500000", trade.StrikePrice)
	require.Equal(t, int64(100), trade.ContractSize)

	// 交易信息
	require.NotNil(t, trade.CreateTime)
	require.NotNil(t, trade.UpdateTime)
	require.Equal(t, "short", trade.Position)
	require.Equal(t, "8.00", trade.Price)
	require.Equal(t, int64(-1), trade.Count)
	require.Equal(t, int64(enum.OPTION_STATUS_HAVING), trade.Status)
	require.Equal(t, "800.00", trade.Premium)
	require.Empty(t, trade.ClosePrice)
	require.Empty(t, trade.Profit)
	require.Empty(t, trade.ProfitRate)
}

func TestNewOptionTrade_卖call_失效(t *testing.T) {
	option, err := NewOption("TCH211230C500000")
	require.NoError(t, err)
	trade, err := NewOptionTrade(option, enum.SHORT, "8.0")
	require.NoError(t, err)

	trade.Invalid()

	// 验证期权基本信息
	require.Equal(t, "TCH211230C500000", trade.Code)
	require.Equal(t, "TCH", trade.Stock)
	require.Equal(t, "211230", trade.ExerciseDate)
	require.Equal(t, "C", trade.Type)
	require.Equal(t, "500000", trade.StrikePrice)
	require.Equal(t, int64(100), trade.ContractSize)

	// 交易信息
	require.NotNil(t, trade.CreateTime)
	require.NotNil(t, trade.UpdateTime)
	require.Equal(t, "short", trade.Position)
	require.Equal(t, "8.00", trade.Price)
	require.Equal(t, int64(-1), trade.Count)
	require.Equal(t, int64(enum.OPTION_STATUS_INVALID), trade.Status)
	require.Equal(t, "800.00", trade.Premium)
	require.Equal(t, "0.00", trade.ClosePrice)
	require.Equal(t, "800.00", trade.Profit)
	require.Equal(t, "1.00", trade.ProfitRate)
}

func TestNewOptionTrade_卖call_平仓_盈(t *testing.T) {
	option, err := NewOption("TCH211230C500000")
	require.NoError(t, err)
	trade, err := NewOptionTrade(option, enum.SHORT, "8.0")
	require.NoError(t, err)

	trade.Close("1.12")

	// 验证期权基本信息
	require.Equal(t, "TCH211230C500000", trade.Code)
	require.Equal(t, "TCH", trade.Stock)
	require.Equal(t, "211230", trade.ExerciseDate)
	require.Equal(t, "C", trade.Type)
	require.Equal(t, "500000", trade.StrikePrice)
	require.Equal(t, int64(100), trade.ContractSize)

	// 交易信息
	require.NotNil(t, trade.CreateTime)
	require.NotNil(t, trade.UpdateTime)
	require.Equal(t, "short", trade.Position)
	require.Equal(t, "8.00", trade.Price)
	require.Equal(t, int64(-1), trade.Count)
	require.Equal(t, int64(enum.OPTION_STATUS_CLOSE), trade.Status)
	require.Equal(t, "800.00", trade.Premium)
	require.Equal(t, "1.12", trade.ClosePrice)
	require.Equal(t, "688.00", trade.Profit)
	require.Equal(t, "0.86", trade.ProfitRate)
}

func TestNewOptionTrade_卖call_平仓_亏(t *testing.T) {
	option, err := NewOption("TCH211230C500000")
	require.NoError(t, err)
	trade, err := NewOptionTrade(option, enum.SHORT, "8.0")
	require.NoError(t, err)

	trade.Close("11.2")

	// 验证期权基本信息
	require.Equal(t, "TCH211230C500000", trade.Code)
	require.Equal(t, "TCH", trade.Stock)
	require.Equal(t, "211230", trade.ExerciseDate)
	require.Equal(t, "C", trade.Type)
	require.Equal(t, "500000", trade.StrikePrice)
	require.Equal(t, int64(100), trade.ContractSize)

	// 交易信息
	require.NotNil(t, trade.CreateTime)
	require.NotNil(t, trade.UpdateTime)
	require.Equal(t, "short", trade.Position)
	require.Equal(t, "8.00", trade.Price)
	require.Equal(t, int64(-1), trade.Count)
	require.Equal(t, int64(enum.OPTION_STATUS_CLOSE), trade.Status)
	require.Equal(t, "800.00", trade.Premium)
	require.Equal(t, "11.20", trade.ClosePrice)
	require.Equal(t, "-320.00", trade.Profit)
	require.Equal(t, "-0.40", trade.ProfitRate)
}

func TestNewOptionTrade_卖put(t *testing.T) {
	option, err := NewOption("TCH211230P450000")
	require.NoError(t, err)
	trade, err := NewOptionTrade(option, enum.SHORT, "12.1")
	require.NoError(t, err)

	// 验证期权基本信息
	require.Equal(t, "TCH211230P450000", trade.Code)
	require.Equal(t, "TCH", trade.Stock)
	require.Equal(t, "211230", trade.ExerciseDate)
	require.Equal(t, "P", trade.Type)
	require.Equal(t, "450000", trade.StrikePrice)
	require.Equal(t, int64(100), trade.ContractSize)

	// 交易信息
	require.NotNil(t, trade.CreateTime)
	require.NotNil(t, trade.UpdateTime)
	require.Equal(t, "short", trade.Position)
	require.Equal(t, "12.10", trade.Price)
	require.Equal(t, int64(-1), trade.Count)
	require.Equal(t, int64(enum.OPTION_STATUS_HAVING), trade.Status)
	require.Equal(t, "1210.00", trade.Premium)
	require.Empty(t, trade.ClosePrice)
	require.Empty(t, trade.Profit)
	require.Empty(t, trade.ProfitRate)
}

func TestNewOptionTrade_卖put_失效(t *testing.T) {
	option, err := NewOption("TCH211230P450000")
	require.NoError(t, err)
	trade, err := NewOptionTrade(option, enum.SHORT, "12.1")
	require.NoError(t, err)

	trade.Invalid()

	// 验证期权基本信息
	require.Equal(t, "TCH211230P450000", trade.Code)
	require.Equal(t, "TCH", trade.Stock)
	require.Equal(t, "211230", trade.ExerciseDate)
	require.Equal(t, "P", trade.Type)
	require.Equal(t, "450000", trade.StrikePrice)
	require.Equal(t, int64(100), trade.ContractSize)

	// 交易信息
	require.NotNil(t, trade.CreateTime)
	require.NotNil(t, trade.UpdateTime)
	require.Equal(t, "short", trade.Position)
	require.Equal(t, "12.10", trade.Price)
	require.Equal(t, int64(-1), trade.Count)
	require.Equal(t, int64(enum.OPTION_STATUS_INVALID), trade.Status)
	require.Equal(t, "1210.00", trade.Premium)
	require.Equal(t, "0.00", trade.ClosePrice)
	require.Equal(t, "1210.00", trade.Profit)
	require.Equal(t, "1.00", trade.ProfitRate)
}

func TestNewOptionTrade_卖put_亏(t *testing.T) {
	option, err := NewOption("TCH211230P450000")
	require.NoError(t, err)
	trade, err := NewOptionTrade(option, enum.SHORT, "12.1")
	require.NoError(t, err)

	trade.Close("16.1")

	// 验证期权基本信息
	require.Equal(t, "TCH211230P450000", trade.Code)
	require.Equal(t, "TCH", trade.Stock)
	require.Equal(t, "211230", trade.ExerciseDate)
	require.Equal(t, "P", trade.Type)
	require.Equal(t, "450000", trade.StrikePrice)
	require.Equal(t, int64(100), trade.ContractSize)

	// 交易信息
	require.NotNil(t, trade.CreateTime)
	require.NotNil(t, trade.UpdateTime)
	require.Equal(t, "short", trade.Position)
	require.Equal(t, "12.10", trade.Price)
	require.Equal(t, int64(-1), trade.Count)
	require.Equal(t, int64(enum.OPTION_STATUS_CLOSE), trade.Status)
	require.Equal(t, "1210.00", trade.Premium)
	require.Equal(t, "16.10", trade.ClosePrice)
	require.Equal(t, "-400.00", trade.Profit)
	require.Equal(t, "-0.33", trade.ProfitRate)
}

func TestNewOptionTrade_卖put_盈(t *testing.T) {
	option, err := NewOption("TCH211230P450000")
	require.NoError(t, err)
	trade, err := NewOptionTrade(option, enum.SHORT, "12.1")
	require.NoError(t, err)

	trade.Close("0.3")

	// 验证期权基本信息
	require.Equal(t, "TCH211230P450000", trade.Code)
	require.Equal(t, "TCH", trade.Stock)
	require.Equal(t, "211230", trade.ExerciseDate)
	require.Equal(t, "P", trade.Type)
	require.Equal(t, "450000", trade.StrikePrice)
	require.Equal(t, int64(100), trade.ContractSize)

	// 交易信息
	require.NotNil(t, trade.CreateTime)
	require.NotNil(t, trade.UpdateTime)
	require.Equal(t, "short", trade.Position)
	require.Equal(t, "12.10", trade.Price)
	require.Equal(t, int64(-1), trade.Count)
	require.Equal(t, int64(enum.OPTION_STATUS_CLOSE), trade.Status)
	require.Equal(t, "1210.00", trade.Premium)
	require.Equal(t, "0.30", trade.ClosePrice)
	require.Equal(t, "1180.00", trade.Profit)
	require.Equal(t, "0.98", trade.ProfitRate)
}

func TestNewOptionTrade_买call(t *testing.T) {
	option, err := NewOption("bili211217C70000")
	require.NoError(t, err)
	trade, err := NewOptionTrade(option, enum.LONG, "0.34")
	require.NoError(t, err)

	// 验证期权基本信息
	require.Equal(t, "BILI211217C70000", trade.Code)
	require.Equal(t, "BILI", trade.Stock)
	require.Equal(t, "211217", trade.ExerciseDate)
	require.Equal(t, "C", trade.Type)
	require.Equal(t, "70000", trade.StrikePrice)
	require.Equal(t, int64(100), trade.ContractSize)

	// 交易信息
	require.NotNil(t, trade.CreateTime)
	require.NotNil(t, trade.UpdateTime)
	require.Equal(t, "long", trade.Position)
	require.Equal(t, "0.34", trade.Price)
	require.Equal(t, int64(1), trade.Count)
	require.Equal(t, int64(enum.OPTION_STATUS_HAVING), trade.Status)
	require.Equal(t, "-34.00", trade.Premium)
	require.Empty(t, trade.ClosePrice)
	require.Empty(t, trade.Profit)
	require.Empty(t, trade.ProfitRate)
}

func TestNewOptionTrade_买call_失效(t *testing.T) {
	option, err := NewOption("bili211217C70000")
	require.NoError(t, err)
	trade, err := NewOptionTrade(option, enum.LONG, "0.34")
	require.NoError(t, err)

	trade.Invalid()

	// 验证期权基本信息
	require.Equal(t, "BILI211217C70000", trade.Code)
	require.Equal(t, "BILI", trade.Stock)
	require.Equal(t, "211217", trade.ExerciseDate)
	require.Equal(t, "C", trade.Type)
	require.Equal(t, "70000", trade.StrikePrice)
	require.Equal(t, int64(100), trade.ContractSize)

	// 交易信息
	require.NotNil(t, trade.CreateTime)
	require.NotNil(t, trade.UpdateTime)
	require.Equal(t, "long", trade.Position)
	require.Equal(t, "0.34", trade.Price)
	require.Equal(t, int64(1), trade.Count)
	require.Equal(t, int64(enum.OPTION_STATUS_INVALID), trade.Status)
	require.Equal(t, "-34.00", trade.Premium)
	require.Equal(t, "0.00", trade.ClosePrice)
	require.Equal(t, "-34.00", trade.Profit)
	require.Equal(t, "-1.00", trade.ProfitRate)
}

func TestNewOptionTrade_买call_赢(t *testing.T) {
	option, err := NewOption("bili211217C70000")
	require.NoError(t, err)
	trade, err := NewOptionTrade(option, enum.LONG, "0.34")
	require.NoError(t, err)

	trade.Close("5.1")

	// 验证期权基本信息
	require.Equal(t, "BILI211217C70000", trade.Code)
	require.Equal(t, "BILI", trade.Stock)
	require.Equal(t, "211217", trade.ExerciseDate)
	require.Equal(t, "C", trade.Type)
	require.Equal(t, "70000", trade.StrikePrice)
	require.Equal(t, int64(100), trade.ContractSize)

	// 交易信息
	require.NotNil(t, trade.CreateTime)
	require.NotNil(t, trade.UpdateTime)
	require.Equal(t, "long", trade.Position)
	require.Equal(t, "0.34", trade.Price)
	require.Equal(t, int64(1), trade.Count)
	require.Equal(t, int64(enum.OPTION_STATUS_CLOSE), trade.Status)
	require.Equal(t, "-34.00", trade.Premium)
	require.Equal(t, "5.10", trade.ClosePrice)
	require.Equal(t, "476.00", trade.Profit)
	require.Equal(t, "14.00", trade.ProfitRate)
}

func TestNewOptionTrade_买call_亏(t *testing.T) {
	option, err := NewOption("bili211217C70000")
	require.NoError(t, err)
	trade, err := NewOptionTrade(option, enum.LONG, "0.34")
	require.NoError(t, err)

	trade.Close("0.1")

	// 验证期权基本信息
	require.Equal(t, "BILI211217C70000", trade.Code)
	require.Equal(t, "BILI", trade.Stock)
	require.Equal(t, "211217", trade.ExerciseDate)
	require.Equal(t, "C", trade.Type)
	require.Equal(t, "70000", trade.StrikePrice)
	require.Equal(t, int64(100), trade.ContractSize)

	// 交易信息
	require.NotNil(t, trade.CreateTime)
	require.NotNil(t, trade.UpdateTime)
	require.Equal(t, "long", trade.Position)
	require.Equal(t, "0.34", trade.Price)
	require.Equal(t, int64(1), trade.Count)
	require.Equal(t, int64(enum.OPTION_STATUS_CLOSE), trade.Status)
	require.Equal(t, "-34.00", trade.Premium)
	require.Equal(t, "0.10", trade.ClosePrice)
	require.Equal(t, "-24.00", trade.Profit)
	require.Equal(t, "-0.71", trade.ProfitRate)
}

func TestNewOptionTrade_买put(t *testing.T) {
	option, err := NewOption("futu211210P47000")
	require.NoError(t, err)
	trade, err := NewOptionTrade(option, enum.LONG, "1.6")
	require.NoError(t, err)

	// 验证期权基本信息
	require.Equal(t, "FUTU211210P47000", trade.Code)
	require.Equal(t, "FUTU", trade.Stock)
	require.Equal(t, "211210", trade.ExerciseDate)
	require.Equal(t, "P", trade.Type)
	require.Equal(t, "47000", trade.StrikePrice)
	require.Equal(t, int64(100), trade.ContractSize)

	// 交易信息
	require.NotNil(t, trade.CreateTime)
	require.NotNil(t, trade.UpdateTime)
	require.Equal(t, "long", trade.Position)
	require.Equal(t, "1.60", trade.Price)
	require.Equal(t, int64(1), trade.Count)
	require.Equal(t, int64(enum.OPTION_STATUS_HAVING), trade.Status)
	require.Equal(t, "-160.00", trade.Premium)
	require.Empty(t, trade.ClosePrice)
	require.Empty(t, trade.Profit)
	require.Empty(t, trade.ProfitRate)
}

func TestNewOptionTrade_买put_失效(t *testing.T) {
	option, err := NewOption("futu211210P47000")
	require.NoError(t, err)
	trade, err := NewOptionTrade(option, enum.LONG, "1.6")
	require.NoError(t, err)

	trade.Invalid()

	// 验证期权基本信息
	require.Equal(t, "FUTU211210P47000", trade.Code)
	require.Equal(t, "FUTU", trade.Stock)
	require.Equal(t, "211210", trade.ExerciseDate)
	require.Equal(t, "P", trade.Type)
	require.Equal(t, "47000", trade.StrikePrice)
	require.Equal(t, int64(100), trade.ContractSize)

	// 交易信息
	require.NotNil(t, trade.CreateTime)
	require.NotNil(t, trade.UpdateTime)
	require.Equal(t, "long", trade.Position)
	require.Equal(t, "1.60", trade.Price)
	require.Equal(t, int64(1), trade.Count)
	require.Equal(t, int64(enum.OPTION_STATUS_INVALID), trade.Status)
	require.Equal(t, "-160.00", trade.Premium)
	require.Equal(t, "0.00", trade.ClosePrice)
	require.Equal(t, "-160.00", trade.Profit)
	require.Equal(t, "-1.00", trade.ProfitRate)
}

func TestNewOptionTrade_买put_盈(t *testing.T) {
	option, err := NewOption("futu211210P47000")
	require.NoError(t, err)
	trade, err := NewOptionTrade(option, enum.LONG, "1.6")
	require.NoError(t, err)

	trade.Close("3")

	// 验证期权基本信息
	require.Equal(t, "FUTU211210P47000", trade.Code)
	require.Equal(t, "FUTU", trade.Stock)
	require.Equal(t, "211210", trade.ExerciseDate)
	require.Equal(t, "P", trade.Type)
	require.Equal(t, "47000", trade.StrikePrice)
	require.Equal(t, int64(100), trade.ContractSize)

	// 交易信息
	require.NotNil(t, trade.CreateTime)
	require.NotNil(t, trade.UpdateTime)
	require.Equal(t, "long", trade.Position)
	require.Equal(t, "1.60", trade.Price)
	require.Equal(t, int64(1), trade.Count)
	require.Equal(t, int64(enum.OPTION_STATUS_CLOSE), trade.Status)
	require.Equal(t, "-160.00", trade.Premium)
	require.Equal(t, "3.00", trade.ClosePrice)
	require.Equal(t, "140.00", trade.Profit)
	require.Equal(t, "0.87", trade.ProfitRate)
}

func TestNewOptionTrade_买put_亏(t *testing.T) {
	option, err := NewOption("futu211210P47000")
	require.NoError(t, err)
	trade, err := NewOptionTrade(option, enum.LONG, "1.6")
	require.NoError(t, err)

	trade.Close("0.6")

	// 验证期权基本信息
	require.Equal(t, "FUTU211210P47000", trade.Code)
	require.Equal(t, "FUTU", trade.Stock)
	require.Equal(t, "211210", trade.ExerciseDate)
	require.Equal(t, "P", trade.Type)
	require.Equal(t, "47000", trade.StrikePrice)
	require.Equal(t, int64(100), trade.ContractSize)

	// 交易信息
	require.NotNil(t, trade.CreateTime)
	require.NotNil(t, trade.UpdateTime)
	require.Equal(t, "long", trade.Position)
	require.Equal(t, "1.60", trade.Price)
	require.Equal(t, int64(1), trade.Count)
	require.Equal(t, int64(enum.OPTION_STATUS_CLOSE), trade.Status)
	require.Equal(t, "-160.00", trade.Premium)
	require.Equal(t, "0.60", trade.ClosePrice)
	require.Equal(t, "-100.00", trade.Profit)
	require.Equal(t, "-0.62", trade.ProfitRate)
}
