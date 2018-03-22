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

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "client is an http client component",
	Long: `client is an http client component`,
	Run: func(cmd *cobra.Command, args []string) {
		client.Run()
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
	cobra.OnInitialize(clientConfig)
	clientCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $PWD/client.yaml)")
	clientCmd.Flags().StringP("url", "u", "http://localhost:8080", "the url for this client to connect to")
	clientCmd.Flags().StringP("insocket", "i", "tcp://localhost:5555?type=pull", "the addr and type of the in socket")
	clientCmd.Flags().StringP("outsocket", "o", "tcp://localhost:5556?type=push" "the addr and type of the out socket")
	viper.BindPFlags(clientCmd.Flags())
	viper.SetDefault("url", "http://localhost:8080")
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
