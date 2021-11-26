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
	"encoding/xml"
	"fmt"
	"github.com/spf13/cobra"
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "parsing",
	Long:  `Parsing JSON or XML`,
	Run: func(cmd *cobra.Command, args []string) {
		filterJSON, err := cmd.Flags().GetBool("json")
		//filterXML, err := cmd.Flags().GetBool("xml")
		if err != nil {
			fmt.Println("Error with flag!", err)
		}
		//fmt.Println("Filter JSON", filterJSON)

		if isJSON(args) && filterJSON {
			fmt.Println("Valid json!")
		} else if isXML(args) {
			fmt.Println("Valid XML!")
		} else {
			fmt.Println("Choose valid data!")
		}
		fmt.Println("DONE!")
	},
}

func isJSON(strings []string) bool {
	var str string
	for _, s := range strings {
		str += s
	}
	var js interface{}
	return json.Unmarshal([]byte(str), &js) == nil
}

func isXML(strings []string) bool {
	var str string
	for _, s := range strings {
		str += s
	}
	var js interface{}
	err := xml.Unmarshal([]byte(str), &js)
	if err != nil {
		fmt.Println("Error", err)
		return false
	}
	return true
}

func init() {
	parseCmd.Flags().BoolP("json", "j", true, "is json")
	parseCmd.Flags().BoolP("xml", "x", true, "is xml")
	rootCmd.AddCommand(parseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// parseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// parseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

//'<?xml version="1.0" encoding="UTF-8"?><note><to>Tove</to><from>Jani</from><heading>Reminder</heading><body>Dont forget me this weekend!</body></note>'
