/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"os"

	"github.com/feigme/fmgr-go/app/models"
	"github.com/feigme/fmgr-go/app/opst"
	"github.com/feigme/fmgr-go/app/query"
	"github.com/feigme/fmgr-go/app/service"
	"github.com/spf13/cobra"
)

// opstCmd represents the opst command
var opstCmd = &cobra.Command{
	Use:   "opst",
	Short: "期权策略",
	Long:  `a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var opstApplyCmd = &cobra.Command{
	Use:   "apply",
	Short: "期权策略apply",
	Long:  `a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		opst, err := opst.ParseYaml(opstApplyFile)
		if err != nil {
			red.Printf("操作失败: %s\n", err.Error())
			os.Exit(1)
		}

		err = service.OpstSvc.Apply(opst)
		if err != nil {
			red.Printf("操作失败: %s\n", err.Error())
			os.Exit(1)
		}

		blue.Println("操作成功! ")
	},
}

var opstListCmd = &cobra.Command{
	Use:   "list",
	Short: "期权策略列表",
	Long:  `a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		list := service.OpstSvc.List(&opstQuery)

		printOpstTableHead()
		printOpstTableRow(list)

	},
}

func printOpstTableHead() {
	blue.Printf("%-6v %-20v \n", "ID", "名称")
}

func printOpstTableRow(list []models.OptionStrategy) {
	for _, trade := range list {
		blue.Printf("%-6v %-22v \n", trade.Id, trade.Name)
	}
}

var (
	opstApplyFile string
	opstQuery     query.OpstQuery
)

func init() {
	rootCmd.AddCommand(opstCmd)

	opstApplyCmd.Flags().StringVarP(&opstApplyFile, "file", "f", "", "yaml文件路径")

	opstCmd.AddCommand(opstApplyCmd, opstListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// opstCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// opstCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
