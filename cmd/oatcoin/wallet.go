package oatcoin

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func NewWalletCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "newaddress",
		Short: "🔐 Generate a new wallet address",
		Run: func(cmd *cobra.Command, args []string) {
			color.Green("New address: %s", "oat1qxyz...")
		},
	}
	return cmd
}
