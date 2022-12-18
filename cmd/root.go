/*
Copyright © 2022 Tyler Garner garnertb@github.com
*/

package cmd

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

const gitignoreUrl = "https://api.github.com/repos/github/gitignore/contents"

type Files struct {
	Tree []File `json:"tree"`
	Sha  string `json:"sha"`
}

type File struct {
	Path string `json:"path"`
	Sha  string `json:"sha"`
	Url  string `json:"url"`
}

func (f File) isSupported() bool {
	return strings.Contains(f.Path, ".gitignore")
}

func (f File) getPayload() string {
	var res FileContent
	getJson(f.Url, &res)

	dec, _ := b64.StdEncoding.DecodeString(res.Content)
	res.Payload = string(dec)
	return res.Payload
}

type FileContent struct {
	Sha      string `json:"sha"`
	NodeId   string `json:"node_id"`
	Size     int    `json:"size"`
	Url      string `json:"url"`
	Content  string `json:"content"`
	Encoding string `json:"encoding"`
	Payload  string
}

type Commands map[string]File

var commands = make(Commands)

func getGitIgnore(filetype string) string {
	command := commands[filetype]
	return command.getPayload()
}

var rootCmd = &cobra.Command{
	Use:       "gh-gitignore",
	Short:     "Load gitignore files from GitHub into your project",
	Long:      `Load gitignore files from GitHub into your project`,
	ValidArgs: []string{"node", "go"},
	//ValidArgsFunction: ValidArgsFunction,
	Args:    cobra.ExactValidArgs(1),
	Example: "gh-gitignore node",
	Run: func(cmd *cobra.Command, args []string) {
		var filetype = args[0]
		fmt.Println(getGitIgnore(filetype))
	},
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		log.Fatal(err)
		return err
	}
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return json.Unmarshal(body, &target)
}

func getCommands() Commands {
	var files []File
	getJson(gitignoreUrl, &files)

	for i := 0; i < len(files); i++ {
		file := files[i]
		if file.isSupported() {
			command := strings.ToLower(strings.Split(file.Path, ".gitignore")[0])
			commands[command] = file
		}
	}

	return commands
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	getCommands()
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
