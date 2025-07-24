package oatcoin

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/isif00/oat-coin/internal/app"
	"github.com/isif00/oat-coin/internal/domain/tx"
)

func InitChainCmd(blockApp *app.BlockApp) *cobra.Command {
	return &cobra.Command{
		Use:   "initchain",
		Short: "🚀 Initialize the blockchain with a genesis block",
		Run: func(cmd *cobra.Command, args []string) {
			blocks, err := blockApp.InitializeBlockchain()
			if err != nil {
				color.Red("❌ Failed to initialize blockchain: %v", err)
				return
			}
			color.Green("✅ Blockchain initialized with %d block(s)", len(blocks))
		},
	}
}

func MineBlockCmd(blockApp *app.BlockApp) *cobra.Command {
	return &cobra.Command{
		Use:   "mineblock",
		Short: "⛏️  Mine a new block with dummy transactions",
		Run: func(cmd *cobra.Command, args []string) {
			txs := []*tx.Transaction{
				{
					ID: []byte("0x1234"),
					Inputs: []tx.TxInput{
						{
							TxID:      []byte("0x1234"),
							OutputIdx: 0,
							Signature: []byte("signature"),
							PubKey:    []byte("public key"),
						},
					},
					Outputs: []tx.TxOutput{
						{
							Amount:     50,
							PubKeyHash: []byte("public key hash"),
						},
					},
				},
			}
			block, err := blockApp.MineBlock(txs)
			if err != nil {
				color.Red("❌ Failed to mine block: %v", err)
				return
			}
			color.Green("✅ Block mined successfully!")
			fmt.Printf("🔗 Hash: %s\n", color.YellowString(block.HashHex()))
		},
	}
}

func LatestBlockCmd(blockApp *app.BlockApp) *cobra.Command {
	return &cobra.Command{
		Use:   "latestblock",
		Short: "📦 Fetch the latest block",
		Run: func(cmd *cobra.Command, args []string) {
			block, err := blockApp.GetLatestBlock()
			if err != nil {
				color.Red("❌ Failed to fetch latest block: %v", err)
				return
			}
			color.Cyan("🧱 Latest Block:")
			fmt.Printf("🔗 Hash: %s\n", block.Hash)
			fmt.Printf("⏱️ Time: %s\n", time.Unix(block.Timestamp, 0).Format("2006-01-02 15:04:05 MST"))
			fmt.Printf("📦 Tx Count: %d\n", len(block.Transactions))
		},
	}
}

func ListBlocksCmd(blockApp *app.BlockApp) *cobra.Command {
	return &cobra.Command{
		Use:   "listblocks",
		Short: "📚 List all blocks in the blockchain",
		Run: func(cmd *cobra.Command, args []string) {
			blocks, err := blockApp.GetAllBlocks()
			if err != nil {
				color.Red("❌ Failed to list blocks: %v", err)
				return
			}
			color.Cyan("🧱 All Blocks:")
			for i, blk := range blocks {
				fmt.Printf("[%d] 🔗 %s | ⏱️  %d | 📦 %d txs\n", i, blk.Hash, blk.Timestamp, len(blk.Transactions))
			}
		},
	}
}
