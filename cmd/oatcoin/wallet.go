package oatcoin

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func NewWalletCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "newwallet",
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

func LoadWalletCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "loadwallet [address]",
		Short: "🔐 Load an existing wallet",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			a := args[0]
			walletData, err := walletApp.LoadWallet(a)
			if err != nil {
				color.Red("Failed to load wallet: %v", err)
				return
			}
			color.Cyan("🔓 Wallet Loaded Successfully")
			fmt.Println()
			fmt.Printf("📬 Address     : %s\n", color.GreenString(walletData.Adress))
			fmt.Printf("🔑 Public Key  : %s\n", color.GreenString(walletData.PubKey))
			fmt.Printf("🕵️  Private Key : %s\n", color.GreenString(walletData.PrivKey))

		},
	}
	return cmd
}

func ListWalletsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "listwallets",
		Short: "🔐 List all existing wallets",
		Run: func(cmd *cobra.Command, args []string) {
			wallets, err := walletApp.ListWallets()
			if err != nil {
				color.Red("Failed to list wallets: %v", err)
				return
			}
			color.Cyan("🔓 Wallets Feteched Successfully")
			for _, wallet := range wallets {
				fmt.Printf("📬 Address     : %s\n", color.GreenString(wallet.Adress))
			}
		},
	}
	return cmd
}
