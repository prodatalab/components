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
)

var rootCmd = &cobra.Command{
	Use:   "websocket",
	Short: "websocket provides commands to lauch ws server or client",
	Long:  `websocket provides commands to launch ws server or client`,
}

// Execute blah
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {}
