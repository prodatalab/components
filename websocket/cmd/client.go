// Copyright © 2018 Peter Alexander <peter.alexander@prodatalab.com>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"fmt"
	"os"

	"github.com/prodatalab/cobra"
	"github.com/prodatalab/components/http/pkg/client"
	"github.com/spf13/viper"
)

var cfgFile string

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "An http client component",
	Long:  `An http client component`,
	Run: func(cmd *cobra.Command, args []string) {
		client.Run()
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
	cobra.OnInitialize(clientConfig)
	clientCmd.Flags().StringVarP(&cfgFile, "config", "c", "${PWD}/client.yaml", "config file (default is $PWD/client.yaml)")

	clientCmd.Flags().StringP("insocket", "i", "tcp://localhost:5555?type=pull")
	clientCmd.Flags().StringP("outsocket", "o", "tcp://localhost:5556?type=push")
	clientCmd.Flags().StringP("wsurl", "w", "ws://localhost:8080", "the websocket addr to connect")

	viper.BindPFlags(clientCmd.Flags)
	viper.SetDefault("wsurl", "ws://localhost:8080")
	viper.SetDefault("insocket", "tcp://localhost:5555?type=pull")
	viper.SetDefault("outsocket", "tcp://localhost:5556?type=push")
}

func clientConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		cwd, _ := os.Getwd()
		viper.AddConfigPath(cwd)
		viper.SetConfigName("client")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
