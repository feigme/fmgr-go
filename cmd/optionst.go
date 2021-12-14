/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/feigme/fmgr-go/app/enum"
	"github.com/feigme/fmgr-go/app/models"
	"github.com/feigme/fmgr-go/app/query"
	"github.com/feigme/fmgr-go/app/service"
	"github.com/spf13/cobra"
)

var (
	red  = color.New(color.FgRed)
	blue = color.New(color.FgHiBlue)
)

// optionstCmd represents the optionst command
var optionstCmd = &cobra.Command{
	Use:   "opst",
	Short: "期权策略",
	Long:  `a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("optionst called")
	},
}

var optionstCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "创建期权",
	Long:  `a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		// 期权code
		blue.Println("请输入期权code：")
		var code string
		fmt.Scanln(&code)
		option, err := models.NewOption(code)
		if err != nil {
			red.Printf("输入期权code错误：%s\n", err.Error())
			os.Exit(1)
		}

		// 期权操作
		blue.Println("请选择操作序号：")
		km := make([]enum.KeyMap, 0)
		km = append(km, enum.KeyMap{Key: fmt.Sprintf("%v", enum.LONG.Desc()), Val: int(enum.LONG)})
		km = append(km, enum.KeyMap{Key: fmt.Sprintf("%v", enum.SHORT.Desc()), Val: int(enum.SHORT)})
		for _, v := range km {
			blue.Printf("  %d: %s\n", v.Val, v.Key)
		}
		var ops int
		fmt.Scanln(&ops)
		operate := enum.OptionCreateEnum(ops)
		if operate != enum.LONG && operate != enum.SHORT {
			red.Printf("输入操作错误！\n")
			os.Exit(1)
		}

		// 操作价格
		blue.Println("请输入价格：")
		var price string
		fmt.Scanln(&price)

		// 数量
		blue.Println("请输入数量：")
		var count int
		fmt.Scanln(&count)

		for i := 0; i < count; i++ {
			trade, err := models.NewOptionTrade(option, operate, price)
			if err != nil {
				red.Printf("输入价格错误：%s\n", err.Error())
				os.Exit(1)
			}
			err = service.OptionTradeSvc.Save(trade)
			if err != nil {
				red.Printf("创建错误：%s\n", err.Error())
				os.Exit(1)
			}
		}
		blue.Printf("期权code：%s，操作：%s，价格：%s，数量：%d \n", option.Code, operate.Desc(), price, count)
	},
}

var optionstListCmd = &cobra.Command{
	Use:   "list",
	Short: "期权列表",
	Long: `查询期权列表，可根据code模糊查询.
  例子：
    fmgr-go opst list baba // 查询期权带有baba关键字的
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var code string
		if len(args) > 0 {
			code = args[0]
		}

		optionTradeQuery.Code = code
		// 默认值, cobra命令默认值没生效
		if len(optionTradeQuery.StatusList) == 0 {
			optionTradeQuery.StatusList = append(optionTradeQuery.StatusList, 1)
		}

		tradeList := service.OptionTradeSvc.List(&optionTradeQuery)
		if len(tradeList) == 0 {
			red.Println("没有数据！")
			os.Exit(1)
		}

		blue.Println("期权code         \t 操作 \t 价格 \t 数量 \t 权利金 \t 状态 \t 收益")
		for _, trade := range tradeList {
			blue.Printf("%s \t %s \t %s \t %d \t %s         \t %s \t %s \n", trade.Code, trade.Position, trade.Price, trade.Count, trade.Premium, enum.OptionStatusEnum(trade.Status).Desc(), trade.Profit)
		}
	},
}

var optionstCloseCmd = &cobra.Command{
	Use:   "close",
	Short: "期权平仓",
	Long:  `a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			red.Println("选择要操作的期权code！")
			os.Exit(1)
		}

		optionTradeQuery.Code = args[0]
		tradeList := service.OptionTradeSvc.List(&optionTradeQuery)
		if len(tradeList) == 0 {
			red.Println("没找到对应的期权！")
			os.Exit(1)
		} else if len(tradeList) > 1 {
			red.Println("找到多个期权，请增加筛选条件！")
			os.Exit(1)
		}
		trade := &tradeList[0]

		blue.Println("期权code         \t 操作 \t 价格 \t 数量 \t 权利金 \t 状态 \t 收益")
		blue.Printf("%s \t %s \t %s \t %d \t %s         \t %s \t %s \n", trade.Code, trade.Position, trade.Price, trade.Count, trade.Premium, enum.OptionStatusEnum(trade.Status).Desc(), trade.Profit)
		fmt.Println()

		// 选择操作
		blue.Println("请选择操作序号：")
		km := make([]enum.KeyMap, 0)
		km = append(km, enum.KeyMap{Key: fmt.Sprintf("%v", enum.CLOSE.Desc()), Val: int(enum.CLOSE)})
		km = append(km, enum.KeyMap{Key: fmt.Sprintf("%v", enum.INVALID.Desc()), Val: int(enum.INVALID)})
		for _, v := range km {
			blue.Printf("  %d: %s\n", v.Val, v.Key)
		}
		var ops int
		fmt.Scanln(&ops)
		operate := enum.OptionCloseEnum(ops)
		if operate != enum.CLOSE && operate != enum.INVALID {
			red.Printf("输入操作错误！%s\n")
			os.Exit(1)
		}

		// 平仓
		if operate == enum.CLOSE {
			// 操作价格
			blue.Println("请输入价格：")
			var price string
			fmt.Scanln(&price)

			trade.Close(operate, price)
			service.OptionTradeSvc.Close(trade)
			blue.Printf("期权code：%s，操作：%s\n", trade.Code, operate.Desc())
		} else if operate == enum.INVALID {
			trade.Invalid(operate)
			service.OptionTradeSvc.Close(trade)
			blue.Printf("期权code：%s，操作：%s，价格：%s\n", trade.Code, operate.Desc(), trade.ClosePrice)
		}

	},
}

var (
	optionTradeQuery query.OptionTradeQuery
)

func init() {
	rootCmd.AddCommand(optionstCmd)
	optionstCmd.AddCommand(optionstCreateCmd)

	// 初始化slice没有生效
	optionstListCmd.Flags().IntSliceVarP(&optionTradeQuery.StatusList, "status", "s", []int{1}, "选择状态，1：持仓，2：平仓，-1：失效")
	optionstListCmd.Flags().StringVar(&optionTradeQuery.StartExerciseDate, "from", fmt.Sprint(time.Now().Format("060102")), "查询行权日范围，开始时间")
	optionstListCmd.Flags().StringVar(&optionTradeQuery.EndExerciseDate, "to", "", "查询行权日范围，结束时间")
	optionstListCmd.Flags().IntVar(&optionTradeQuery.PageSize, "pageSize", 20, "分页查询，查询数据量")
	optionstListCmd.Flags().IntVar(&optionTradeQuery.PageNo, "pageNo", 1, "分页查询，页数")
	optionstCmd.AddCommand(optionstListCmd)

	optionstCloseCmd.Flags().IntSliceVarP(&optionTradeQuery.StatusList, "status", "s", []int{}, "选择状态，1：持仓，2：平仓，-1：失效")
	optionstCmd.AddCommand(optionstCloseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// optionstCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// optionstCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
