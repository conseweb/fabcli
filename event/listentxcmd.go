/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package event

import (
	"fmt"

	"github.com/conseweb/fabcli/common"
	"github.com/hyperledger/fabric-sdk-go/api/apitxn"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var listenTxCmd = &cobra.Command{
	Use:   "listentx",
	Short: "Listen to transaction events.",
	Long:  "Listen to transaction events",
	Run: func(cmd *cobra.Command, args []string) {
		if common.Config().TxID() == "" {
			fmt.Printf("\nMust specify the transaction ID\n\n")
			cmd.HelpFunc()(cmd, args)
			return
		}
		action, err := newListenTXAction(cmd.Flags())
		if err != nil {
			common.Config().Logger().Criticalf("Error while initializing listenTxAction: %v", err)
			return
		}

		defer action.Terminate()

		err = action.invoke()
		if err != nil {
			common.Config().Logger().Criticalf("Error while running listenTxAction: %v", err)
		}
	},
}

func getListenTXCmd() *cobra.Command {
	flags := listenTxCmd.Flags()
	common.Config().InitTxID(flags)
	common.Config().InitPeerURL(flags, "", "The URL of the peer on which to listen for events, e.g. localhost:7051")
	return listenTxCmd
}

type listentxAction struct {
	common.Action
	inputEvent
}

func newListenTXAction(flags *pflag.FlagSet) (*listentxAction, error) {
	action := &listentxAction{inputEvent: inputEvent{done: make(chan bool)}}
	err := action.Initialize(flags)
	return action, err
}

func (action *listentxAction) invoke() error {
	done := make(chan bool)

	eventHub, err := action.EventHub()
	if err != nil {
		return err
	}

	fmt.Printf("Registering TX event for TxID [%s]\n", common.Config().TxID())

	txnID := apitxn.TransactionID{ID: common.Config().TxID()}
	eventHub.RegisterTxEvent(txnID, func(txID string, code pb.TxValidationCode, err error) {
		fmt.Printf("Received TX event. TxID: %s, Code: %s, Error: %s\n", txID, code, err)
		done <- true
	})

	action.WaitForEnter()

	fmt.Printf("Unregistering TX event for TxID [%s]\n", common.Config().TxID())
	eventHub.UnregisterTxEvent(txnID)

	return nil
}
