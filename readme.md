# 🥣 Oatcoin (v0.1)

> A minimalist reimplementation of the original Bitcoin (v0.1) using Go.
> Inspired by Satoshi Nakamoto’s whitepaper. No smart contracts. No bloat. Just pure UTXO, Proof-of-Work, and P2P vibes.

---

## ✨ Features

* 🔐 ECDSA-based keypair wallet
* 💸 UTXO-based transaction model
* ⛓ Block & blockchain validation logic
* ⚒ Proof-of-Work mining (SHA-256)
* 📡 P2P network for block & transaction propagation
* 🧠 Blockchain persistence using flat files
* 🧾 Coinbase transactions (block rewards)
* ⚙️ CLI wallet & node control

---

## 🔧 How It Works

### ⛓ Blockchain

Each block contains:

* A list of transactions
* A timestamp
* The previous block hash
* A nonce (for mining)
* A hash (block ID)

Blocks are stored to disk using flat files, just like the original Bitcoin client.

### 💸 UTXO Model

Each transaction references outputs of previous transactions (unspent outputs), creating a chain of value ownership.

### ⚒ Mining

* Simple SHA-256 double hashing
* Adjustable difficulty via leading-zero target
* Coinbase transactions create new coins

### 🔐 Wallets

* ECDSA (secp256k1)
* Public key hash = wallet address
* Supports signing and verifying transactions

### 📡 P2P

* Nodes communicate over TCP
* Blocks and transactions are serialized and propagated
* New peers can sync the chain from others

---

## 💻 CLI Commands

These are the commands supported by the CLI:

## 🔐 Wallet
```
oatcoin newaddress                # Generate a new wallet address
oatcoin getbalance                # Show wallet balance
oatcoin listutxos                 # List all unspent outputs
oatcoin dumpwallet                # Dump private keys (dev/debug use)
oatcoin importkey [privkey]       # Import a private key
```
## 💸 Transactions
```
oatcoin send [to] [amount]                # Send coins to an address
oatcoin sendfrom [from] [to] [amount]     # Send coins from a specific address
oatcoin gettx [txid]                      # Get transaction details
oatcoin listtxs                           # List recent transactions
```
## ⚒ Mining
```
oatcoin mine start               # Start CPU mining
oatcoin mine stop                # Stop mining
oatcoin mine status              # Show mining info (hashrate, difficulty)
```
## 📦 Blocks
```
oatcoin getblock [height|hash]   # View a block's data
oatcoin getblockhash [height]    # Get block hash by height
oatcoin getlatestblock           # Get tip of the chain
```
---

## 🛠️ To Do (Future Ideas)

* [ ] Mempool support
* [ ] Difficulty adjustment
* [ ] Node discovery & handshake
* [ ] Web UI / JSON-RPC server
* [ ] P2P encryption (TLS/libp2p)
* [ ] Lightweight explorer
* [ ] Enhanced wallet backup and restore

---

## 🧠 Based On

* [Bitcoin Whitepaper](https://bitcoin.org/bitcoin.pdf)
* Satoshi’s [original source code](https://github.com/benjiqq/bitcoinArchive)
* [btcsuite/btcd](https://github.com/btcsuite/btcd)

---

## 🕳 Philosophy

> *“If you don’t believe me or don’t get it, I don’t have time to try to convince you, sorry.”*
> — Satoshi Nakamoto

---

## 🪪 License

MIT — free as in freedom.

---

## 🧱 Want to Contribute?

This project is for educational purposes but built to be extensible. If you’re interested in hacking on the protocol, reach out or fork it!
