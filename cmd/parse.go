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
	"crypto/md5"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "parsing",
	Long:  `Parsing JSON or XML`,
	Run: func(cmd *cobra.Command, args []string) {
		filterJSON, err := cmd.Flags().GetBool("json")
		filterXML, err := cmd.Flags().GetBool("xml")
		if err != nil {
			fmt.Println("Error with flag!", err)
		}

		getUUID()

		data := joinStr(args)
		hash := md5.Sum([]byte(data))
		hashStr := fmt.Sprintf("%x", hash)

		str, ok := fileArchiveHash[hashStr]
		if ok == false {
			if isJSON(data) && filterJSON {
				fmt.Println("Valid json!")
				//fmt.Printf("%x", md5.Sum([]byte(data)))
				fileName := getUUID()
				fileName1 := string(fileName) + ".json"

				f, err := os.Create(string(fileName1))
				defer f.Close()
				if err != nil {
					fmt.Println(err)
				} else {

					f.WriteString(data)
					fmt.Println("Done")
				}

				//err := os.WriteFile("/tmp/dat1", d1, 0644)
			} else if isXML(data) && filterXML {
				fmt.Println("Valid XML!")

				fileName := getUUID()
				fileName1 := string(fileName) + ".xml"
				f, err := os.Create(string(fileName1))
				defer f.Close()
				if err != nil {
					fmt.Println(err)
				} else {
					f.WriteString(data)
					fmt.Println("Done")
				}
			} else {
				fmt.Println("Choose valid data!")
			}
			fmt.Println("DONE!")
		} else {
			panic(str)
		}
	},
}

func getUUID() []byte {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	return out
}

func joinStr(s []string) string {
	var str string
	for _, v := range s {
		str += v
	}
	return str
}

func isJSON(data string) bool {
	var js interface{}
	return json.Unmarshal([]byte(data), &js) == nil
}

func isXML(data string) bool {
	var js interface{}
	return xml.Unmarshal([]byte(data), &js) == nil
}

var fileArchiveHash = map[string]bool{}

func init() {
	parseCmd.Flags().BoolP("json", "j", false, "is json")
	parseCmd.Flags().BoolP("xml", "x", false, "is xml")

	var files []string
	root := "/home/osoko/GolandProjects/cli-decoder/"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".json" || filepath.Ext(path) == ".xml" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file)

		f, err := os.Open(file)

		if err != nil {
			panic(err)
		}

		defer f.Close()

		hash := md5.New()
		_, err = io.Copy(hash, f)

		if err != nil {
			panic(err)
		}

		md5HashString := fmt.Sprintf("%x", hash.Sum(nil))
		fileArchiveHash[md5HashString] = true
	}
	//fmt.Println(fileArchiveHash)

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
