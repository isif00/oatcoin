package tx

type TxInput struct {
	TxID      string
	OutputIdx int
	Signature []byte
	PubKey    []byte
}

type TxOutput struct {
	Amount     int
	PubKeyHash []byte
}

type Transaction struct {
	ID      string
	Inputs  []TxInput
	Outputs []TxOutput
}
