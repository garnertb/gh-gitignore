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

var validArgs = []string{"al", "actionscript", "ada", "agda", "android", "appengine", "appceleratortitanium",
	"archlinuxpackages", "autotools", "c++", "c", "cfwheels", "cmake", "cuda", "cakephp", "chefcookbook", "clojure",
	"codeigniter", "commonlisp", "composer", "concrete5", "coq", "craftcms", "d", "dm", "dart", "delphi", "drupal",
	"episerver", "eagle", "elisp", "elixir", "elm", "erlang", "expressionengine", "extjs", "fancy", "finale",
	"flaxengine", "forcedotcom", "fortran", "fuelphp", "gwt", "gcov", "gitbook", "go", "godot", "gradle", "grails",
	"haskell", "igorpro", "idris", "jboss", "jenkins_home", "java", "jekyll", "joomla", "julia", "kicad", "kohana",
	"kotlin", "labview", "laravel", "leiningen", "lemonstand", "lilypond", "lithium", "lua", "magento", "maven", "mercury",
	"metaprogrammingsystem", "nanoc", "nim", "node", "ocaml", "objective-c", "opa", "opencart", "oracleforms", "packer",
	"perl", "phalcon", "playframework", "plone", "prestashop", "processing", "purescript", "python", "qooxdoo", "qt", "r",
	"ros", "racket", "rails", "raku", "rhodesrhomobile", "ruby", "rust", "scons", "sass", "scala", "scheme", "scrivener",
	"sdcc", "seamgen", "sketchup", "smalltalk", "stella", "sugarcrm", "swift", "symfony", "symphonycms", "tex", "terraform",
	"textpattern", "turbogears2", "twincat3", "typo3", "unity", "unrealengine", "vvvv", "visualstudio", "waf", "wordpress",
	"xojo", "yeoman", "yii", "zendframework", "zephir"}

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
	return string(dec)
}

type FileContent struct {
	Sha      string `json:"sha"`
	NodeId   string `json:"node_id"`
	Size     int    `json:"size"`
	Url      string `json:"url"`
	Content  string `json:"content"`
	Encoding string `json:"encoding"`
}

type Commands struct {
	commands          map[string]command
	supportedCommands []string
}

func (c *Commands) runCommand(name string) string {
	return c.commands[name].getGitIgnore()
}

func (c *Commands) registerCommand(com command) {
	if c.commands == nil {
		c.commands = make(map[string]command)
	}
	c.commands[com.Arg] = com
	c.supportedCommands = append(c.supportedCommands, com.Arg)
}

func (c *Commands) loadCommands() {
	var files []File
	getJson(gitignoreUrl, &files)

	for i := 0; i < len(files); i++ {
		file := files[i]
		if file.isSupported() {
			arg := strings.ToLower(strings.Split(file.Path, ".gitignore")[0])
			c.registerCommand(command{Arg: arg, File: file})
		}
	}
}

type command struct {
	Arg  string
	File File
}

func (c command) getGitIgnore() string {
	return c.File.getPayload()
}

var commands = new(Commands)

var rootCmd = &cobra.Command{
	Use:   "gh-gitignore [arg]",
	Short: "Load gitignore files from GitHub into your project",
	Long: `Load gitignore files from GitHub into your project.

Supported gitignores are: ` + strings.Join(validArgs, ", ") + `
	`,
	ValidArgs: validArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		commands.loadCommands()
	},
	Args:    cobra.ExactValidArgs(1),
	Example: "gh-gitignore node: returns the .gitignore for node projects",
	Run: func(cmd *cobra.Command, args []string) {
		v := commands.runCommand(args[0])
		fmt.Println(v)
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
