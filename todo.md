### 1. **Wallets**

* Keypair generation (ECDSA, secp256k1) ✅
* Address generation (compressed pubkey + hashing + Base58Check) ✅
* Signing transactions 
* Verifying signatures
* Local storage ✅
* Wallet management (list, load, save, backup)
* Recovery mechanism (mnemonics / seeds)
* Unit and integration tests 
---

### 2. **Transactions & UTXO**

* Create and validate transactions
* Implement UTXO set tracking
* Transaction inputs referencing UTXOs
* Outputs creating new UTXOs
* Double-spend prevention
* Fee calculation
* Signature verification on tx inputs
* Mempool for unconfirmed txs
* Tests for tx validation and UTXO correctness

---

### 3. **Blockchain Core**

* Block structure (header + txs)
* Chain linking via prev block hash
* Block validation (PoW, timestamp, tx validation)
* Maintain chain state (tip, height)
* Persistence (flat file or DB)
* Fork handling (longest chain rule)
* Difficulty adjustment (future)
* Genesis block creation
* Tests for block validation and chain integrity

---

### 4. **Proof-of-Work Mining**

* Nonce iteration to find valid hash
* Adjustable difficulty target
* Coinbase transaction creation (block reward)
* Mining loop and interrupt support
* Validation of PoW on incoming blocks
* Tests for mining and difficulty checks

---

### 5. **Networking / P2P**

* Peer discovery and connection management
* Message serialization and deserialization
* Propagation of blocks and transactions
* Syncing chain on node startup
* Handling forks and conflicting blocks
* Security and DoS prevention (basic rate limiting)
* Future: encryption, handshakes, node reputation
* Tests for message handling and sync correctness

---

### 6. **Storage**

* Flat file storage for blocks (like blk\*.dat)
* Indexing for fast lookup (block height, hashes)
* UTXO set persistence (if stored separately)
* Data corruption detection and recovery
* Backup and restore functionality
* Tests for data consistency and I/O correctness

---

### 7. **CLI & User Interface**

* Wallet creation and management commands
* Sending transactions and querying balance
* Viewing blockchain info (printchain)
* Node control commands (startnode, stop)
* Clear error reporting and user feedback
* Future: REST API / Web UI
* Tests for CLI command correctness and edge cases

---

### 8. **Security & Testing**

* Comprehensive unit tests for core logic
* Integration tests covering node sync and transactions
* Handling malformed / malicious input gracefully
* Private key safety and secure random number generation
* Code audits and reviews
* Continuous Integration setup (optional but recommended)