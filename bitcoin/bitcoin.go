package bitcoin

import (
	"coinspaid.com/crypto/rpc"
	"encoding/json"
	"errors"
	"log"
	"strconv"
)

const (
	// DEFAULT_RPCCLIENT_TIMEOUT represent http timeout for rcp client
	RPCCLIENT_TIMEOUT = 30
)

// A Bitcoind represents a Bitcoind client
type Bitcoin struct {
	client *rpc.RpcClient
}

// handleError handle error returned by client.call
func handleError(err error, r *rpc.RpcResponse) error {
	if err != nil {
		return err
	}
	if r.Err != nil {
		return r.Err
	}

	return nil
}

// New return a new Bitcoin
func New(host string, port int, user, passwd string, useSSL bool, timeoutParam ...int) (*Bitcoin, error) {
	var timeout int = RPCCLIENT_TIMEOUT
	// If the timeout is specified in timeoutParam, allow it.
	if len(timeoutParam) != 0 {
		timeout = timeoutParam[0]
	}

	rpcClient, err := rpc.NewClient(host, port, user, passwd, useSSL, timeout)
	if err != nil {
		return nil, err
	}
	return &Bitcoin{rpcClient}, nil
}

// GetBlockCount returns the number of blocks in the longest block chain.
// https://bitcoincore.org/en/doc/0.20.0/rpc/blockchain/getblockcount/
func (b *Bitcoin) GetBlockCount() (count uint64, err error) {
	r, err := b.client.Call("getblockcount", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	count, err = strconv.ParseUint(string(r.Result), 10, 64)
	return
}

// GetNewAddress return a new address for account [account].
// https://bitcoincore.org/en/doc/0.20.0/rpc/wallet/getnewaddress/
func (b *Bitcoin) GetNewAddress(account ...string) (addr string, err error) {
	// 0 or 1 account
	if len(account) > 1 {
		err = errors.New("Bad parameters for GetNewAddress: you can set 0 or 1 account")
		return
	}
	r, err := b.client.Call("getnewaddress", account)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &addr)
	return
}

// GetBlock returns information about the block with the given hash.
func (b *Bitcoin) GetBlock(blockHash string) (block Block, err error) {
	r, err := b.client.Call("getblock", []string{blockHash})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &block)
	return
}

// GetBalance return the balance of the server or of a specific account
//If [account] is "", returns the server's total available balance.
//If [account] is specified, returns the balance in the account
func (b *Bitcoin) GetBalance(account string, minconf uint64) (balance float64, err error) {
	r, err := b.client.Call("getbalance", []interface{}{account, minconf})
	if err = handleError(err, &r); err != nil {
		return
	}
	balance, err = strconv.ParseFloat(string(r.Result), 64)
	return
}

// EstimateSmartFee stimates the approximate fee per kilobyte needed for a transaction..
// https://bitcoincore.org/en/doc/0.20.0/rpc/util/estimatesmartfee/
func (b *Bitcoin) EstimateSmartFee(minconf int, estimateMode string) (ret SmartFee, err error) {

	valid := map[string]bool{"UNSET": true, "ECONOMICAL": true, "CONSERVATIVE": true}
	if !valid[estimateMode] {
		err = errors.New("Not valid estimateMode: " + estimateMode)
		return
	}

	r, err := b.client.Call("estimatesmartfee", []interface{}{minconf, estimateMode})
	if err = handleError(err, &r); err != nil {
		return
	}

	err = json.Unmarshal(r.Result, &ret)
	return
}

func (b *Bitcoin) GetTransaction(txid string) (transaction Transaction, err error) {
	r, err := b.client.Call("gettransaction", []interface{}{txid})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &transaction)
	return
}

func (b *Bitcoin) GetRawTransaction(txId string) (transaction RawTransaction, err error) {
	r, err := b.client.Call("getrawtransaction", []interface{}{txId, 1})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &transaction)
	return
}

func (b *Bitcoin) DecodeRawTransaction(hex string) (transaction RawTransaction, err error) {
	r, err := b.client.Call("decoderawtransaction", []interface{}{hex})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &transaction)
	return
}

func (b *Bitcoin) SignRawTransactionWithWallet(hex string) (transaction SignedRawTransaction, err error) {
	r, err := b.client.Call("signrawtransactionwithwallet", []interface{}{hex})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &transaction)
	return
}

func (b *Bitcoin) CreateRawTransaction(inputs []Input, outputs map[string]float64) (hex string, err error) {
	r, err := b.client.Call("createrawtransaction", []interface{}{inputs, outputs})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &hex)
	return
}

func (b *Bitcoin) SendRawTransaction(signedHex string) (txhash string, err error) {
	r, err := b.client.Call("sendrawtransaction", []string{signedHex})
	log.Println(r.Result)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &txhash)
	return
}

func (b *Bitcoin) GetBlockHash(index uint64) (hash string, err error) {
	r, err := b.client.Call("getblockhash", []uint64{index})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &hash)
	return
}

func (b *Bitcoin) ListUnspent(minconf, maxconf uint64) (transactions []Transaction, err error) {

	args := []interface{}{minconf}

	if maxconf > 0 {
		if maxconf < minconf {
			maxconf = minconf
		}
		args = append(args, maxconf)
	}

	r, err := b.client.Call("listunspent", args)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &transactions)
	return
}

func (b *Bitcoin) SendToAddress(toAddress string, amount float64, comment, commentTo string) (txID string, err error) {
	r, err := b.client.Call("sendtoaddress", []interface{}{toAddress, amount, comment, commentTo})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &txID)
	return
}
