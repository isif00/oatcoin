package main

import (
	"log"

	"github.com/isif00/oat-coin/cmd/oatcoin"
)

func main() {
	app, err := oatcoin.NewOatCoin(".oatcoin")
	if err != nil {
		log.Fatalf("Failed to init oatcoin: %v", err)
	}

	oatcoin.Register(
		oatcoin.NewWalletCmd(app.WalletApp),
		oatcoin.LoadWalletCmd(app.WalletApp),
		oatcoin.ListWalletsCmd(app.WalletApp),

		oatcoin.InitChainCmd(app.BlockApp),
		oatcoin.MineBlockCmd(app.BlockApp),
		oatcoin.LatestBlockCmd(app.BlockApp),
		oatcoin.ListBlocksCmd(app.BlockApp),
	)
	oatcoin.Execute()
}
