// Copyright © 2016 Jeremy 'TheHipbot' Chambers jeremy@thehipbot.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	appFs   afero.Fs
	cfgFile string
	profile string
	recur   bool
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "massnomer",
	Short: "A multifile renaming tool",
	Long: `A tool which can be used to rename files based on search
and replace patterns. Has some base configurations but can be extended
with your own patterns.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		if !viper.IsSet(profile) {
			return errors.New("profile not found")
		}
		moveFile()
		return nil
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	appFs = afero.NewOsFs()
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.massnomer.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Rename profile to use
	RootCmd.PersistentFlags().StringVar(&profile, "profile", "p", "profile to use from defaults or defined in config")

	// TODO recursive option
	// RootCmd.PersistentFlags().StringVar(&recur, "recursive", "r", "search and rename files recursively")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetFs(appFs)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".massnomer") // name of config file (without extension)
	viper.AddConfigPath("$HOME")      // adding home directory as first search path
	viper.AutomaticEnv()              // read in environment variables that match

	// set default profiles
	shows := map[string]interface{}{
		"exts":     []string{"mkv", "avi"},
		"patterns": []string{"/[sS]([0-9]+)[eE]([0-9]+).+(720p|1080p)?/"},
		"result":   "S$1E$2 $3",
	}

	viper.SetDefault("shows", shows)

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func moveFile() {
	appFs.Rename("testingblah", "testingnot")
}

func getFileMappings(cmd *cobra.Command, args []string) {

}
