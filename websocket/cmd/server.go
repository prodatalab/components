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

// serverCmd represents the server command
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

	serverCmd.Flags().StringArrayVarP(&server.Val.Sockets, "sockets", "s", []string{},
		"Use the form: <tcp|ipc|inproc>://localhost:5555?type=<req|rep|push|pull|pub|sub>")
	serverCmd.Flags().StringVarP(&server.Val.WSURL, "websocket", "w", "ws://localhost:8080", "The websocket address to connect to")
	serverCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $PWD/server.yaml)")
	viper.BindPFlags(clientCmd.Flags)
	viper.SetDefault("websocket", "http://localhost:8080")
	viper.SetDefault("sockets", []string{"tcp://localhost:5555?type=push", "tcp://localhost:5556?type=pull"})
}

// Config reads in config file and ENV variables if set.
func serverConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// // Find home directory.
		// home, err := homedir.Dir()
		// if err != nil {
		// 	fmt.Println(err)
		// 	os.Exit(1)
		cwd, _ := os.Getwd()

		// Search config in home directory with name "server" (without extension).
		viper.AddConfigPath(cwd)
		viper.SetConfigName("server")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
