/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package event

import (
	"github.com/conseweb/fabcli/common"
	"github.com/spf13/cobra"
)

var eventCmd = &cobra.Command{
	Use:   "event",
	Short: "Event commands",
	Long:  "Event commands",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

// Cmd returns the events command
func Cmd() *cobra.Command {
	common.Config().InitChannelID(eventCmd.Flags())

	eventCmd.AddCommand(getListenCCCmd())
	eventCmd.AddCommand(getListenTXCmd())
	eventCmd.AddCommand(getListenBlockCmd())

	return eventCmd
}
