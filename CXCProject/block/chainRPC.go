package block

import (
	"CXCProject/block/model"
	"encoding/json"
	"log"
	"strconv"
)

//钱包的JSON-RPC链接
const (
	// 钱包版本
	VERSION = 0.1
	// 客户端默认超时时间
	RPCCLIENT_TIMEOUT = 30
)

//链接客户端
type CXCcoind struct {
	client *RpcClient
}

//创建一个新的客户端连接
func New(host string, port int, user, passwd string, useSSL bool) (bit *CXCcoind, err error) {
	rpcC, err := newClient(host, port, user, passwd, useSSL)
	if err != nil {
		return nil, err
	}
	return &CXCcoind{rpcC}, nil
}

//查看区块同步情况
func (b *CXCcoind) ShowChain() (block model.BlockChain, err error) {
	r, err := b.client.call("showchain", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &block)
	return
}

//获取区块数量
func (b *CXCcoind) GetBlockCount() (count uint64, err error) {
	r, err := b.client.call("showblockcount", nil)
	if err = handleError(err, &r); err != nil {
		log.Println("ShowBlockCount Error:", err.Error())
		return
	}

	count, err = strconv.ParseUint(string(r.Result), 10, 64)
	b.client.httpClient.CloseIdleConnections()
	return
}

//获取最新的区块Hash
func (b *CXCcoind) GetBestBlock() (hash string, err error) {

	r, err := b.client.call("showbestblockhash", nil)
	if err = handleError(err, &r); err != nil {
		log.Println("ShowBlock Error:", err.Error())
		return
	}
	err = json.Unmarshal(r.Result, &hash)
	return
}

//查看链中已发行资产基本信息
func (b *CXCcoind) GetAssets(name ...string) (assets []model.AssetsInfo, err error) {

	r := rpcResponse{}
	if len(name) != 0 {
		r, err = b.client.call("showassets", []interface{}{name})
	} else {
		r, err = b.client.call("showassets", nil)
	}

	if err = handleError(err, &r); err != nil {
		log.Println("ShowAssets Error:", err.Error())
		return
	}

	err = json.Unmarshal(r.Result, &assets)
	return
}

//根据区块hash获取区块信息
//参数：hash string
func (b *CXCcoind) GetBlock(hash string) (block model.Block, err error) {

	r, err := b.client.call("showblock", []interface{}{hash, 4})

	//log.Printf("DATA:%s",r.Result)
	if err = handleError(err, &r); err != nil {
		log.Println("ShowBlock Error:", err.Error())
		return
	}
	err = json.Unmarshal(r.Result, &block)
	return
}

//根据区块ID获取区块信息 ShowBlocks
//参数 height uint64
func (b *CXCcoind) GetBlocks(height ...uint64) (block []model.Block, err error) {
	r, err := b.client.call("showblocks", []interface{}{height})
	if err = handleError(err, &r); err != nil {
		log.Println("ShowBlocks Error:", err.Error())
		return
	}
	err = json.Unmarshal(r.Result, &block)
	return
}

//根据区块ID获取的区块hash
//参数 height uint64
func (b *CXCcoind) GetBlockHash(height uint64) (hash string, err error) {
	r, err := b.client.call("showblockhash", []interface{}{height})
	if err = handleError(err, &r); err != nil {
		log.Println("ShowBlockHash Error:", err.Error())
		return
	}
	err = json.Unmarshal(r.Result, &hash)
	return
}

//获取区块参数
type getBlockTemplateParams struct {
	Mode         string   `json:"mode,omitempty"`
	Capabilities []string `json:"capabilities,omitempty"`
}

//获取区块参数
/*func (b *CXCcoind) GetBlockTemplate(capabilities []string, mode string) (template string, err error) {
	params := getBlockTemplateParams{
		Mode:         mode,
		Capabilities: capabilities,
	}
	// TODO []interface{}{mode, capa}
	r, err := b.client.call("getblocktemplate", []getBlockTemplateParams{params})
	if err = handleError(err, &r); err != nil {
		return
	}
	fmt.Println(json.Unmarshal(r.Result, &template))
	return
}
*/

//=============================================以下是BTC操作方法===============================================

// "验证签名后信息<signature>是否与指定地址<bitcoinaddress>的私钥签名的信息<message>一致。
//Verify a signed message. "
/*func (b *CXCcoind) VerifyMessage(address, sign, message string) (success bool, err error) {
	r, err := b.client.call("verifymessage", []interface{}{address, sign, message})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &success)
	return
}
*/

//通过账户查找地址
/*func (b *CXCcoind) GetAddressesByAccount(account string) (addresses []string, err error) {
	r, err := b.client.call("getaddressesbyaccount", []string{account})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &addresses)
	return
}
*/

//获取链接数量
func (b *CXCcoind) GetConnectionCount() (count uint64, err error) {
	r, err := b.client.call("getconnectioncount", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	count, err = strconv.ParseUint(string(r.Result), 10, 64)
	return
}

//获取挖矿难度
func (b *CXCcoind) GetDifficulty() (difficulty float64, err error) {
	r, err := b.client.call("getdifficulty", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	difficulty, err = strconv.ParseFloat(string(r.Result), 64)
	return
}

//获取挖矿状态
func (b *CXCcoind) GetGenerate() (generate bool, err error) {
	r, err := b.client.call("getgenerate", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &generate)
	return
}

//最近每秒的hash值
func (b *CXCcoind) GetHashesPerSec() (hashpersec float64, err error) {
	r, err := b.client.call("gethashespersec", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &hashpersec)
	return
}

//获取钱包矿池信息
func (b *CXCcoind) GetMiningInfo() (miningInfo model.MiningInfo, err error) {
	r, err := b.client.call("getmininginfo", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &miningInfo)
	return
}

//获取原始更改地址 在交易的时候使用.
/*func (b *CXCcoind) GetRawChangeAddress(account ...string) (rawAddress string, err error) {
	// 0 or 1 account
	if len(account) > 1 {
		err = errors.New("Bad parameters for GetRawChangeAddress: you can set 0 or 1 account")
		return
	}
	r, err := b.client.call("getrawchangeaddress", account)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &rawAddress)
	return
}*/

// 获得内存池信息  在版本0.7中被替换成getblocktemplate, submitblock, getrawmempool
func (b *CXCcoind) GetRawMempool() (txId []string, err error) {
	r, err := b.client.call("getrawmempool", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &txId)
	return
}

//通过TXID 获得原始交易信息.
/*func (b *CXCcoind) GetRawTransaction(txId string, verbose bool) (rawTx interface{}, err error) {
	intVerbose := 0
	if verbose {
		intVerbose = 1
	}
	r, err := b.client.call("getrawtransaction", []interface{}{txId, intVerbose})
	if err = handleError(err, &r); err != nil {
		return
	}
	if !verbose {
		err = json.Unmarshal(r.Result, &rawTx)
	} else {
		var t model.RawTransaction
		err = json.Unmarshal(r.Result, &t)
		rawTx = t
	}
	return
}*/

//接收帐户<帐户> [minconf=1]
//返回在至少与[minconf]确认的事务中使用<account>的地址接收的总金额。
/*func (b *CXCcoind) GetReceivedByAccount(account string, minconf uint32) (amount float64, err error) {
	if account == "all" {
		account = ""
	}
	r, err := b.client.call("getreceivedbyaccount", []interface{}{account, minconf})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &amount)
	return
}*/

//获取地址收款金额
//在本地钱包中使用，外部地址总是显示0。
/*func (b *CXCcoind) GetReceivedByAddress(address string, minconf uint32) (amount float64, err error) {
	r, err := b.client.call("getreceivedbyaddress", []interface{}{address, minconf})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &amount)
	return
}*/

//通过TXID 获得交易信息
/*func (b *CXCcoind) GetTransaction(txid string) (transaction model.Transaction, err error) {
	r, err := b.client.call("gettransaction", []interface{}{txid})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &transaction)
	return
}*/

// gettxout 取得交易输出 用于交易	参数 <txid> <n>
/*func (b *CXCcoind) GetTxOut(txid string, n uint32, includeMempool bool) (transactionOut model.UTransactionOut, err error) {
	r, err := b.client.call("gettxout", []interface{}{txid, n, includeMempool})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &transactionOut)
	return
}*/

// 取得交易输出设定信息
func (b *CXCcoind) GetTxOutsetInfo() (txOutSet model.TransactionOutSet, err error) {
	r, err := b.client.call("gettxoutsetinfo", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &txOutSet)
	return
}

// GetWork 计算挖矿
// If [data] is not specified, returns formatted hash data to work on
// If [data] is specified, tries to solve the block and returns true if it was successful.
/*func (b *CXCcoind) GetWork(data ...string) (response interface{}, err error) {
	if len(data) > 1 {
		err = errors.New("查找失败，设置一个或不设置参数")
		return
	}
	var r rpcResponse

	if len(data) == 0 {
		r, err = b.client.call("getwork", nil)
		if err = handleError(err, &r); err != nil {
			return
		}
		var work model.Work
		err = json.Unmarshal(r.Result, &work)
		response = work
	} else {
		r, err = b.client.call("getwork", data)
		if err = handleError(err, &r); err != nil {
			return
		}
		var t bool
		err = json.Unmarshal(r.Result, &t)
		response = t
	}
	return
}*/

//导入私钥
/*func (b *CXCcoind) ImportPrivKey(privKey, label string, rescan bool) error {
	r, err := b.client.call("importprivkey", []interface{}{privKey, label, rescan})
	return handleError(err, &r)
}*/

//填满矿池
func (b *CXCcoind) KeyPoolRefill() error {
	r, err := b.client.call("keypoolrefill", nil)
	return handleError(err, &r)
}

/*//账户信息，账户名称 --账户余额
func (b *CXCcoind) ListAccounts(minconf int32) (accounts map[string]float64, err error) {
	r, err := b.client.call("listaccounts", []int32{minconf})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &accounts)
	return
}
*/

// 接收账户：一个账户接收多少币
type ReceivedByAccount struct {
	Account       string  //接收地址
	Amount        float64 //接收数量
	Confirmations uint32  //确认次数
}

//列出账户的收款信息
/*func (b *CXCcoind) ListReceivedByAccount(minConf uint32, includeEmpty bool) (list []ReceivedByAccount, err error) {
	r, err := b.client.call("listreceivedbyaccount", []interface{}{minConf, includeEmpty})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &list)
	return
}
*/
// 地址收款
type ReceivedByAddress struct {
	Address       string   // 收款地址
	Account       string   //收款賬戶
	Amount        float64  //收款數量
	Confirmations uint32   //確認數量
	TxIds         []string //TXID
}

// 列出指定块之后的交易
/*func (b *CXCcoind) ListSinceBlock(blockHash string, targetConfirmations uint32) (transaction []model.Transaction, err error) {
	r, err := b.client.call("listsinceblock", []interface{}{blockHash, targetConfirmations})
	if err = handleError(err, &r); err != nil {
		return
	}
	type ts struct {
		Transactions []model.Transaction
	}
	var result ts
	if err = json.Unmarshal(r.Result, &result); err != nil {
		return
	}
	transaction = result.Transactions
	return
}*/

// 列出交易
/*func (b *CXCcoind) ListTransactions(account string, count, from uint32) (transaction []model.Transaction, err error) {
	r, err := b.client.call("listtransactions", []interface{}{account, count, from})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &transaction)
	return
}*/

// 在钱包中列出未动用输出返回未使用事务输入的数组。
/*func (b *CXCcoind) ListUnspent(minconf, maxconf uint32) (transactions []model.Transaction, err error) {
	if maxconf > 999999 {
		maxconf = 999999
	}
	r, err := b.client.call("listunspent", []interface{}{minconf, maxconf})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &transactions)
	return
}*/

//更新暂时未动用的交易输出列表
/*func (b *CXCcoind) LockUnspent(lock bool, outputs []UnspendableOutput) (success bool, err error) {
	r, err := b.client.call("lockunspent", []interface{}{lock, outputs})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &success)
	return
}*/

//"根据<generate>参数的值（true/false）决定客户端是否采矿。
//[genproclimit]参数指定采矿最大线程数，-1代表无限制。
//<generate> is true or false to turn generation on or off.
//Generation is limited to [genproclimit] processors, -1 is unlimited."
/*func (b *CXCcoind) SetGenerate(generate bool, genProcLimit int32) error {
	r, err := b.client.call("setgenerate", []interface{}{generate, genProcLimit})
	return handleError(err, &r)
}*/

//开启挖矿
/*func (b *CXCcoind) Generate(nblocks int32) error {
	r, err := b.client.call("generate", []interface{}{nblocks})
	return handleError(err, &r)
}*/

// "交易费<amount>是一个四舍五入至小数点后8位的实数。
//<amount> is a real and is rounded to the nearest 0.00000001"
/*func (b *CXCcoind) SetTxFee(amount float64) error {
	r, err := b.client.call("settxfee", []interface{}{amount})
	return handleError(err, &r)
}*/

// 暫停錢包服務.
func (b *CXCcoind) Stop() error {
	r, err := b.client.call("stop", nil)
	return handleError(err, &r)
}

//"用地址<bitcoinaddress>的私钥对信息<message>进行数字签名。需要未锁定钱包。
//Sign a message with the private key of an address."
/*func (b *CXCcoind) SignMessage(address, message string) (sig string, err error) {
	r, err := b.client.call("signmessage", []interface{}{address, message})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &sig)
	return
}*/

// "把钱包的密码<passphrase>存储在内存中持续<timeout>秒。
//Stores the wallet decryption key in memory for <timeout> seconds."
/*func (b *CXCcoind) WalletPassphrase(passPhrase string, timeout uint64) error {
	r, err := b.client.call("walletpassphrase", []interface{}{passPhrase, timeout})
	return handleError(err, &r)
}*/
//"将钱包的原密码<oldpassphrase>修改为新密码<newpassphrase>
//Changes the wallet passphrase from <oldpassphrase> to <newpassphrase>. "
/*func (b *CXCcoind) WalletPassphraseChange(oldPassphrase, newPassprhase string) error {
	r, err := b.client.call("walletpassphrasechange", []interface{}{oldPassphrase, newPassprhase})
	return handleError(err, &r)
}*/
