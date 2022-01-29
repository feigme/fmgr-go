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
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

// optionCmd represents the option command
var optionCmd = &cobra.Command{
	Use:   "option",
	Short: "期权操作命令",
	Long:  `a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("option called")
	},
}

var optionListCmd = &cobra.Command{
	Use:   "list",
	Short: "期权列表",
	Long: `查询期权列表，可根据code模糊查询.
  例子: 
    fmgr-go opst list baba // 查询期权带有baba关键字的
	`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
			optionListCmdQuery.Code = args[0]
		}

		tradeList := service.OptionTradeSvc.List(&optionListCmdQuery)
		if len(tradeList) == 0 {
			red.Println("没有数据! ")
			os.Exit(1)
		}

		printOptionTableHead()
		for _, trade := range tradeList {
			printOptionTableRow(&trade)
		}
	},
}

var (
	optionListCmdQuery   query.OptionTradeQuery
	optionCloseCmdQuery  query.OptionTradeQuery
	optionDeleteCmdQuery query.OptionTradeQuery
)

func init() {
	rootCmd.AddCommand(optionCmd)

	optionListCmd.Flags().IntSliceVarP(&optionListCmdQuery.StatusList, "status", "s", []int{1}, "选择状态，1: 持仓，2: 平仓，-1: 失效")
	optionListCmd.Flags().StringVar(&optionListCmdQuery.StartExerciseDate, "from", fmt.Sprint(time.Now().Format("060102")), "查询行权日范围，开始时间")
	optionListCmd.Flags().StringVar(&optionListCmdQuery.EndExerciseDate, "to", "", "查询行权日范围，结束时间")
	optionListCmd.Flags().IntVar(&optionListCmdQuery.PageSize, "pageSize", 20, "分页查询，查询数据量")
	optionListCmd.Flags().IntVar(&optionListCmdQuery.PageNo, "pageNo", 1, "分页查询，页数")

	optionCmd.AddCommand(optionListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// optionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// optionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func printOptionTableHead() {
	blue.Printf("%-6v %-20v %-7v %-8v %-6v %-10v %-8v %-11v %-10v \n", "ID", "期权code", "操作", "数量", "买价", "卖价", "状态", "收益", "收益率")
}

func printOptionTableRow(trade *models.OptionTrade) {
	colorFunc := color.New(color.FgHiBlue).SprintFunc()
	if trade.Profit != "" && cast.ToFloat64(trade.Profit) < 0 {
		colorFunc = color.New(color.FgGreen).SprintFunc()
	}
	if trade.Profit != "" && cast.ToFloat64(trade.Profit) > 0 {
		colorFunc = color.New(color.FgRed).SprintFunc()
	}

	rate := ""
	if trade.ProfitRate != "" {
		rate = fmt.Sprintf("%.2f%%", cast.ToFloat64(trade.ProfitRate)*100)
	}

	blue.Printf("%-6v %-22v %-9v %-10v %-8v %-14v %-8v %-22v %-21v \n", trade.Id, trade.Code, trade.Position, trade.Count, trade.BuyPrice, trade.SellPrice, enum.OptionStatusEnum(trade.Status).Desc(), colorFunc(trade.Profit), colorFunc(rate))
}
