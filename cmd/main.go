package main

import "github.com/isif00/oat-coin/cmd/oatcoin"

func main() {
	oatcoin.Register(
		oatcoin.NewWalletCmd(),
		oatcoin.LoadWalletCmd(),
		oatcoin.ListWalletsCmd(),

		oatcoin.InitChainCmd(),
		oatcoin.MineBlockCmd(),
		oatcoin.LatestBlockCmd(),
		oatcoin.ListBlocksCmd(),
	)
	oatcoin.Execute()
}
