package oatcoin

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "oatcoin",
	Short: "🥣 Oatcoin - A Bitcoin clone in Go",
	Long: banner() + `
Oatcoin is a v0.1 Bitcoin reimplementation.

Available commands include:
  💳 Wallet      🔗 Transactions     ⛏ Mining
  📦 Blocks      📡 P2P Network      🔧 Dev Tools
	`,
	Run: func(cmd *cobra.Command, args []string) {
		color.Cyan("Try running `oatcoin --help` to get started 🛠")
	},
}

func banner() string {
	c := color.New(color.FgHiGreen).Add(color.Bold)
	return c.Sprint(`
	  ___          ___                 ___          ___                   ___     
     /  /\        /  /\        ___    /  /\        /  /\      ___        /__/\    
    /  /::\      /  /::\      /  /\  /  /:/       /  /::\    /  /\       \  \:\   
   /  /:/\:\    /  /:/\:\    /  /:/ /  /:/       /  /:/\:\  /  /:/        \  \:\  
  /  /:/  \:\  /  /:/~/::\  /  /:/ /  /:/  ___  /  /:/  \:\/__/::\    _____\__\:\ 
 /__/:/ \__\:\/__/:/ /:/\:\/  /::\/__/:/  /  /\/__/:/ \__\:\__\/\:\__/__/::::::::\
 \  \:\ /  /:/\  \:\/:/__\/__/:/\:\  \:\ /  /:/\  \:\ /  /:/  \  \:\/\  \:\~~\~~\/
  \  \:\  /:/  \  \::/    \__\/  \:\  \:\  /:/  \  \:\  /:/    \__\::/\  \:\  ~~~ 
   \  \:\/:/    \  \:\         \  \:\  \:\/:/    \  \:\/:/     /__/:/  \  \:\     
    \  \::/      \  \:\         \__\/\  \::/      \  \::/      \__\/    \  \:\    
     \__\/        \__\/               \__\/        \__\/                 \__\/ 
	`)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func Register(cmd ...*cobra.Command) {
	for _, c := range cmd {
		rootCmd.AddCommand(c)
	}
}
