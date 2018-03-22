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
	"github.com/prodatalab/components/websocket/pkg/client"
	"github.com/spf13/viper"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "A websocket client component",
	Long:  `A websocket client component`,
	Run: func(cmd *cobra.Command, args []string) {
		client.Run()
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
	cobra.OnInitialize(clientConfig)
	clientCmd.Flags().StringArrayVarP(&client.Val.Sockets, "sockets", "s", []string{}, "Use the form: <tcp|ipc|inproc>://localhost:5555?type=<req|rep|push|pull|pub|sub>")
	clientCmd.Flags().StringVarP(&client.Val.WSURL, "websocket", "w", "", "The websocket address to connect to")
	clientCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $PWD/client.yaml)")
	viper.BindPFlags(clientCmd.Flags)
	viper.SetDefault("websocket", "http://localhost:8080")
	viper.SetDefault("sockets", []string{"tcp://localhost:5555?type=push", "tcp://localhost:5556?type=pull"})
}

// clientConfig reads in config file and ENV variables if set.
func clientConfig() {
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

		// Search config in home directory with name ".blah" (without extension).
		viper.AddConfigPath(cwd)
		viper.SetConfigName("server")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
