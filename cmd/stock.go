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

	"github.com/feigme/fmgr-go/app/models"
	"github.com/feigme/fmgr-go/app/service"
	"github.com/spf13/cobra"
)

// stockCmd represents the stock command
var stockCmd = &cobra.Command{
	Use:   "stock",
	Short: "股票操作命令",
	Long:  `a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stock called")
	},
}

var stockCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "记录购买股票",
	Long:  `a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		blue.Println("选择股票市场: (HK/US)")
		var market string
		fmt.Scanln(&market)

		blue.Println("股票代码: (港股用中文)")
		var code string
		fmt.Scanln(&code)

		optionCode := service.StockTradeSvc.GetOptionCode(code)
		if optionCode == "" {
			blue.Println("输入股票对应的期权code: ")
			fmt.Scanln(&optionCode)
		}

		blue.Println("股票购买价格: ")
		var price string
		fmt.Scanln(&price)

		blue.Println("股票数量: ")
		var count int
		fmt.Scanln(&count)

		stock, err := models.NewStockTrade(market, code, optionCode, price, count)
		if err != nil {
			red.Println(err.Error())
			os.Exit(1)
		}

		err = service.StockTradeSvc.Save(stock)
		if err != nil {
			red.Println(err.Error())
			os.Exit(1)
		}

		blue.Println("操作成功! ")
	},
}

func init() {
	rootCmd.AddCommand(stockCmd)

	stockCmd.AddCommand(stockCreateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stockCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stockCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
