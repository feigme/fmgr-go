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

	"github.com/feigme/fmgr-go/app/enum"
	"github.com/feigme/fmgr-go/app/query"
	"github.com/feigme/fmgr-go/app/service"
	"github.com/spf13/cobra"
)

// strategyCmd represents the strategy command
var strategyCmd = &cobra.Command{
	Use:   "st",
	Short: "期权策略管理",
	Long:  `a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		blue.Println("请选择策略Id：")
		for _, v := range enum.OptionStEnumList() {
			blue.Printf("  %d: %-20v %s\n", v.Key, v.Code, v.Desc)
		}

		var stId int
		fmt.Scanln(&stId)
		OptionStrategyEnum, err := enum.GetOptionStEnumByKey(stId)
		if err != nil {
			red.Printf("%s！\n", err)
			os.Exit(1)
		}

		// 不同策略不同交互提示
		if OptionStrategyEnum == enum.OST_ROLLING_PUT {
			// 查出当前已存在的sell put
			list := service.OptionTradeSvc.List(&query.OptionTradeQuery{
				Position:   "short",
				StatusList: []int{int(enum.OPTION_STATUS_HAVING)},
			})
			if len(list) == 0 {
				red.Println("没找到持仓的put！")
				os.Exit(1)
			}
			printOptionTableHead()
			for _, v := range list {
				printOptionTableRow(&v)
			}
			blue.Println("")
			blue.Println("请选择需要Id：")

			var opitonId int
			fmt.Scanln(&opitonId)
			blue.Println("请输入平仓价格：")
			var closePrice string
			fmt.Scanln(&closePrice)
			blue.Println("请输入要roll到的行权日：（格式YYMMDD）")
			var exerciseDate string
			fmt.Scanln(&exerciseDate)
			blue.Println("请输入买入价格：")
			var sellPrice string
			fmt.Scanln(&sellPrice)

			err = service.OptionTradeSvc.RollPut(opitonId, closePrice, exerciseDate, sellPrice)
			if err != nil {
				red.Printf("%s！\n", err)
				os.Exit(1)
			}
			blue.Println("操作成功！")
		}
	},
}

func init() {
	rootCmd.AddCommand(strategyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// strategyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// strategyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
