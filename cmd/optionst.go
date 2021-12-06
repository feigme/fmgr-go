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

	"github.com/fatih/color"
	"github.com/feigme/fmgr-go/app/enum"
	"github.com/feigme/fmgr-go/app/models"
	"github.com/feigme/fmgr-go/app/service"
	"github.com/spf13/cobra"
)

// optionstCmd represents the optionst command
var optionstCmd = &cobra.Command{
	Use:   "optionst",
	Short: "期权策略",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("optionst called")
	},
}

var optionstCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "创建期权策略",
	Long:  `创建期权策略命令.`,
	Run: func(cmd *cobra.Command, args []string) {
		red := color.New(color.FgRed)
		blue := color.New(color.FgHiBlue)

		// 期权code
		blue.Println("请输入期权code：")
		var code string
		fmt.Scanln(&code)
		option, err := models.NewOption(code)
		if err != nil {
			red.Printf("输入期权code错误：%s\n", err.Error())
			os.Exit(1)
		}
		blue.Printf("期权code：%s\n", option.Code)

		// 期权操作
		blue.Println("请现在操作序号：")
		for _, v := range enum.OptionOperateList() {
			blue.Printf("  %d: %s\n", v.Val, v.Key)
		}
		var ops int
		fmt.Scanln(&ops)
		optionOperate, err := enum.GetOptionOperateByKey(ops)
		if err != nil {
			red.Printf("输入操作错误：%s\n", err.Error())
			os.Exit(1)
		}
		blue.Printf("期权code：%s，操作：%s\n", option.Code, optionOperate.Desc())

		// 操作价格
		blue.Println("请输入价格：")
		var price string
		fmt.Scanln(&price)

		trade, err := models.NewOptionTrade(option, optionOperate, price)
		if err != nil {
			red.Printf("输入价格错误：%s\n", err.Error())
			os.Exit(1)
		}
		blue.Printf("期权code：%s，操作：%s，价格：%s\n", option.Code, optionOperate.Desc(), price)

		err = service.OptionTradeSvc.Save(trade)
		if err != nil {
			red.Printf("创建错误：%s\n", err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(optionstCmd)
	optionstCmd.AddCommand(optionstCreateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// optionstCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// optionstCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
