// Copyright © 2018 Peter Alexander <peter.alexander@prodatalab.com>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"github.com/prodatalab/cobra"
	"github.com/prodatalab/components/websocket/pkg/server"
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
	serverCmd.Flags().StringArrayVarP(&server.Val.Sockets, "sockets", "s", []string{},
		"Use the form: <tcp|ipc|inproc>://localhost:5555?type=<req|rep|push|pull|pub|sub>")
	serverCmd.Flags().StringVarP(&server.Val.WSURL, "websocket", "w", "ws://localhost:8080", "The websocket address to connect to")

}
