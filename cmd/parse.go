package cmd

import (
	"crypto/md5"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

const jsonFlag string = "json"
const xmlFlag string = "xml"

const projectPath = "/GolandProjects/cli-decoder/"

func getFlagBool(cmd *cobra.Command, flag string) bool {
	filter, err := cmd.Flags().GetBool(flag)
	checkError("Error while getting a flag", err)
	return filter
}

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "parsing",
	Long:  `Parsing JSON or XML`,
	Run: func(cmd *cobra.Command, args []string) {
		filterJSON := getFlagBool(cmd, jsonFlag)
		filterXML := getFlagBool(cmd, xmlFlag)

		data := joinStr(args)
		hash := md5.Sum([]byte(data))
		hashStr := fmt.Sprintf("%x", hash)

		_, ok := fileArchiveHash[hashStr]
		if ok == false {
			if isJSON(data) && filterJSON {
				fmt.Println("Valid json!")
				fileName := getUUID()
				fileName1 := string(fileName) + ".json"

				f, err := os.Create(string(fileName1))
				defer f.Close()

				checkError("Error while creating a file", err)

				_, err = f.WriteString(data)
				checkError("Error while writing data to file", err)

			} else if isXML(data) && filterXML {
				fmt.Println("Valid XML!")

				fileName := getUUID()
				fileName1 := string(fileName) + ".xml"
				f, err := os.Create(string(fileName1))
				defer f.Close()
				checkError("Error while creating the filename", err)
				_, err = f.WriteString(data)
				checkError("Error while writing data to file", err)

			} else {
				fmt.Println("Choose a valid data!")
			}
			fmt.Println("DONE!")
		} else {
			panic("the data is already existed")
		}
	},
}

func getUUID() []byte {
	out, err := exec.Command("uuidgen").Output()
	checkError("Error while creating uuid", err)
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

	setFileHashes()
	rootCmd.AddCommand(parseCmd)
}

func getFiles() []string {
	var files []string
	root, _ := homedir.Dir()
	root = root + projectPath

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".json" || filepath.Ext(path) == ".xml" {
			files = append(files, path)
		}
		return nil
	})
	checkError("Error while getting existed files", err)
	return files
}

func setFileHashes() {
	for _, file := range getFiles() {
		f, err := os.Open(file)
		checkError("Error while opening a file", err)

		defer f.Close()
		hash := md5.New()
		_, err = io.Copy(hash, f)
		checkError("Error while creating a hash", err)

		md5HashString := fmt.Sprintf("%x", hash.Sum(nil))
		fileArchiveHash[md5HashString] = true
	}
}

func checkError(s string, err error) {
	if err != nil {
		panic(s)
	}
}

//'<?xml version="1.0" encoding="UTF-8"?><note><to>Tove</to><from>Jani</from><heading>Reminder</heading><body>Dont forget me this weekend!</body></note>'
