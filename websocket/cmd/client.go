// Copyright © 2018 Peter Alexander <peter.alexander@prodatalab.com>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"github.com/prodatalab/cobra"
	"github.com/prodatalab/components/websocket/pkg/client"
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
	clientCmd.Flags().StringArrayVarP(&client.Val.Sockets, "sockets", "s", []string{}, "Use the form: <tcp|ipc|inproc>://localhost:5555?type=<req|rep|push|pull|pub|sub>")
	clientCmd.Flags().StringVarP(&client.Val.WSURL, "websocket", "w", "", "The websocket address to connect to")

}
