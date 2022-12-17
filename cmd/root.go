/*
Copyright © 2022 Tyler Garner garnertb@github.com
*/

package cmd

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
	"encoding/json"
	"log"
	"net/http"
	"io/ioutil"
	"time"
	b64 "encoding/base64"
	"strings"
	_ "embed"
)

//go:embed gitignore.json
var fileContent string

type Files struct {
	Tree []File `json:"tree"`
	Sha string `json:"sha"`
}

type File struct {
	Path string `json:"path"`
	Mode string `json:"mode"`
	Filetype string `json:"type"`
	Sha string `json:"sha"`
	Url string `json:"url"`
}

type FileContent struct {
	Sha string `json:"sha"`
	NodeId string `json:"node_id"`
	Size int `json:"size"`
	Url string `json:"url"`
	Content string `json:"content"`
	Encoding string `json:"encoding"`
}

type Commands map[string][]string

var rootCmd = &cobra.Command{
	Use:   "gh-gitignore",
	Short: "Load gitignore files from GitHub into your project",
	Long: `Load gitignore files from GitHub into your project`,
	ValidArgs: []string{"node", "go"},
	Args:  cobra.ExactValidArgs(1),
	Example: "gh-gitignore node",
	Run: func(cmd *cobra.Command, args []string) {
		var filetype = args[0]

		var files = ReadFile()

		//get length of files
		log.Println(files[filetype])

		var res FileContent
		getJson(files[filetype][0], &res)

		var Dec, _ = b64.StdEncoding.DecodeString(res.Content)
		fmt.Println(string(Dec))
	 },
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
    r, err := myClient.Get(url)
    if err != nil {
				log.Fatal(err)
        return err
    }
    var body, error = ioutil.ReadAll(r.Body)
		
		if error != nil {
			log.Fatal(error)
			return error
		}

    return json.Unmarshal(body, &target)
}

func ReadFile()(Commands) {
	 var files Files
   json.Unmarshal([]byte(fileContent), &files)

	 filteredCommands := make(map[string][]string)

	 for i := 0; i < len(files.Tree); i++ {
		var file = files.Tree[i]

		if (strings.Contains(file.Path, ".gitignore")) {
			command := strings.ToLower(strings.Split(file.Path, ".gitignore")[0])
			filteredCommands[command] = append(filteredCommands[command], file.Url)
		}
	 }

	 return filteredCommands
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "There was an error while executing gitignore '%s'", err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gh-gitignore.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


