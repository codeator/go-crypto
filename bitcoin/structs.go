package bitcoin

type Block struct {
	Hash              string   `json:"hash"`
	Confirmations     uint64   `json:"confirmations"`
	Size              uint64   `json:"size"`
	Height            uint64   `json:"height"`
	Version           uint32   `json:"version"`
	Merkleroot        string   `json:"merkleroot"`
	Tx                []string `json:"tx"`
	Time              int64    `json:"time"`
	Nonce             uint64   `json:"nonce"`
	Bits              string   `json:"bits"`
	Difficulty        float64  `json:"difficulty"`
	Chainwork         string   `json:"chainwork,omitempty"`
	Previousblockhash string   `json:"previousblockhash"`
	Nextblockhash     string   `json:"nextblockhash"`
}

type SmartFee struct {
	FeeRate float64  `json:"feerate"`
	Errors  []string `json:"errors"`
	Blocks  int      `json:"blocks"`
}

type ScriptSig struct {
	Asm string `json:"asm"`
	Hex string `json:"hex"`
}

type Vin struct {
	Coinbase  string    `json:"coinbase"`
	Txid      string    `json:"txid"`
	Vout      int       `json:"vout"`
	ScriptSig ScriptSig `json:"scriptSig"`
	Sequence  uint32    `json:"sequence"`
}

type ScriptPubKey struct {
	Asm       string   `json:"asm"`
	Hex       string   `json:"hex"`
	ReqSigs   int      `json:"reqSigs,omitempty"`
	Type      string   `json:"type"`
	Addresses []string `json:"addresses,omitempty"`
}

type Vout struct {
	Value        float64      `json:"value"`
	N            int          `json:"n"`
	ScriptPubKey ScriptPubKey `json:"scriptPubKey"`
}

type RawTransaction struct {
	Hex           string `json:"hex"`
	Txid          string `json:"txid"`
	Version       uint32 `json:"version"`
	LockTime      uint32 `json:"locktime"`
	Vin           []Vin  `json:"vin"`
	Vout          []Vout `json:"vout"`
	BlockHash     string `json:"blockhash,omitempty"`
	Confirmations uint64 `json:"confirmations,omitempty"`
	Time          int64  `json:"time,omitempty"`
	Blocktime     int64  `json:"blocktime,omitempty"`
}

type SignedRawTransactionError struct {
	Txid      string `json:"txid"`
	Vout      int64  `json:"vout"`
	ScriptSig string `json:"scriptSig"`
	Sequence  int64  `json:"sequence"`
	Error     string `json:"error"`
}

type Input struct {
	Txid string `json:"txid"`
	Vout int64 `json:"vout"`
	Sequence int64 `json:"sequence"`
}

type SignedRawTransaction struct {
	Hex      string                      `json:"hex"`
	Complete bool                        `json:"complete"`
	Errors   []SignedRawTransactionError `json:"errors"`
}

type TransactionDetails struct {
	Account  string  `json:"account"`
	Address  string  `json:"address,omitempty"`
	Category string  `json:"category"`
	Amount   float64 `json:"amount"`
	Fee      float64 `json:"fee,omitempty"`
}

type Transaction struct {
	Amount          float64              `json:"amount"`
	Account         string               `json:"account,omitempty"`
	Address         string               `json:"address,omitempty"`
	Category        string               `json:"category,omitempty"`
	Fee             float64              `json:"fee,omitempty"`
	Confirmations   int64                `json:"confirmations"`
	BlockHash       string               `json:"blockhash"`
	BlockIndex      int64                `json:"blockindex"`
	BlockTime       int64                `json:"blocktime"`
	Txid            string               `json:"txid"`
	WalletConflicts []string             `json:"walletconflicts"`
	Time            int64                `json:"time"`
	TimeReceived    int64                `json:"timereceived"`
	Details         []TransactionDetails `json:"details,omitempty"`
	Hex             string               `json:"hex,omitempty"`
}
