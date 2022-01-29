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
	"github.com/feigme/fmgr-go/app/models/strategy"
	"github.com/spf13/cobra"
)

// strategyCmd represents the strategy command
var strategyCmd = &cobra.Command{
	Use:   "st",
	Short: "期权策略管理",
	Long:  `a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		blue.Println("请选择策略Id: ")
		for _, v := range enum.OptionStEnumList() {
			blue.Printf("  %d: %-20v %s\n", v.Key, v.Code, v.Desc)
		}

		var stId int
		fmt.Scanln(&stId)
		OptionStrategyEnum, err := enum.GetOptionStEnumByKey(stId)
		if err != nil {
			red.Printf("%s! \n", err)
			os.Exit(1)
		}

		// 不同策略不同交互提示
		if OptionStrategyEnum == enum.OST_Naked_Short_Put {
			param := &strategy.NakedShortPut{}
			err = param.Hci()
			if err != nil {
				red.Printf("%s! \n", err)
				os.Exit(1)
			}
			
			blue.Println("操作成功! ")
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