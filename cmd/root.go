/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
	"github.com/roccowong95/airtable-cli/common/conf"
	"github.com/roccowong95/airtable-cli/core"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	cfg     conf.Conf
	name    string
	clicore core.CliCore
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "airtable-cli",
	Short: "A brief description of your application",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.airtable-cli.yaml)")
	rootCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "app to use (as defined in config)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	f, err := os.Open(cfgFile)
	if nil != err {
		panic(err)
	}
	bts, err := ioutil.ReadAll(f)
	if nil != err {
		panic(err)
	}
	err = yaml.Unmarshal(bts, &cfg)
	if nil != err {
		panic(err)
	}
	cfg.Fixup()
	appconf := cfg.GetApp(name)
	if nil == appconf {
		panic(fmt.Sprintf("no app named %s found", name))
	}
	clicore, err = core.NewAirtableCore(*appconf)
	if nil != err {
		panic(err)
	}
}
