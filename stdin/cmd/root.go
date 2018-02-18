// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/prodatalab/components/stdin/pkg/stdin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile              string
	configSocketURL      string
	reportSocketURL      string
	dstreamSocketURL     string
	dstreamTransportType string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "stdin",
	Short: "A component of CodeDepot's component-based programming platform",
	Long: `    
	stdin is a "source" component and accepts input from the terminal shell. 
        It then continually sends the input to downstream components.`,
	Run: func(cmd *cobra.Command, args []string) {
		stdin.Init(configSocketURL, reportSocketURL, dstreamSocketURL, dstreamTransportType)
		stdin.Run()
		// stdin.Stop()
	},
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
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "configFile", "f", "$HOME/.config/cbp/stdin.yaml", "config file location")
	rootCmd.PersistentFlags().StringVarP(&configSocketURL, "configURL", "c", "127.0.0.1:5555", "URL with portfor the upstream configuration service")
	rootCmd.PersistentFlags().StringVarP(&reportSocketURL, "reportURL", "r", "127.0.0.1:6666", "URL for the downstream reporting service")
	rootCmd.PersistentFlags().StringVarP(&dstreamSocketURL, "dstreamURL", "d", "127.0.0.1:7777", "URL for the downstream component service")
	rootCmd.PersistentFlags().StringVarP(&dstreamTransportType, "dstreamTransportType", "t", "tcp", "transport type")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Search config in home directory with name ".stdin" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".stdin")
	}
	viper.AutomaticEnv() // read in environment variables that match
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
