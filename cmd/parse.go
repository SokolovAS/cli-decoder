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
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "parsing",
	Long:  `Parsing JSON or XML`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("parse called")
		fmt.Println(args)
		var str string
		for _, s := range args {
			str += s
		}
		if isJSON(str) {
			fmt.Println("Valid json!")
		} else {
			fmt.Println(str)
			fmt.Println("Invalid json!")
		}
	},
}

func isJSON(s string) bool {
	var js interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

func init() {
	parseCmd.Flags().BoolP("json", "j", true, "is json")
	rootCmd.AddCommand(parseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// parseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// parseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
