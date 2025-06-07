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
			walletAddress, err := walletApp.CreateWallet()
			if err != nil {
				color.Red("Failed to generate wallet: %v", err)
				return
			}
			color.Yellow("Wallet Created Successfully: %s", color.GreenString(walletAddress))
		},
	}
	return cmd
}
