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
	"github.com/prodatalab/components/websocket/pkg/server"
	"github.com/spf13/viper"
)

var cfgFile string

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A websocket server component",
	Long:  `A websocket server component`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Run()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	cobra.OnInitialize(serverConfig)
	serverCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $PWD/server.yaml)")

	serverCmd.Flags().StringP("insocket", "i", "tcp://localhost:5555?type=pull", "the addr/type of the in socket")
	serverCmd.Flags().StringP("outsocket", "o", "tcp://localhost:5556?type=push", "the addr/type of the out socket")
	serverCmd.Flags().StringP("wsurl", "w", "ws://localhost:8080", "The websocket address to bind")

	viper.BindPFlags(clientCmd.Flags)
	viper.SetDefault("wsurl", "http://localhost:8080")
	viper.SetDefault("insocket", "tcp://localhost:5555?type=pull")
	viper.SetDefault("outsocket", "tcp://localhost:5556?type=push")
}

func serverConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		cwd, _ := os.Getwd()
		viper.AddConfigPath(cwd)
		viper.SetConfigName("server")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
