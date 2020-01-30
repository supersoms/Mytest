/***************************************************
 ** @Desc :钱包操作
 ** @Time : 2020/1/9
 ** @Author : Administrator
 ** @File : walletRPC.go
 ** @Last Modified by : Administrator
 ** @Last Modified time: 2020-01-09-11:40
 ** @Software: GoLand
****************************************************/
package block

import (
	"CXCProject/block/model"
	"encoding/json"
	"errors"
	"log"
)

//动态创建CXCChain创建新的钱包地址
func (b *CXCcoind) GetNewAddress(account ...string) (addr string, err error) {
	// 0 or 1 account
	if len(account) > 1 {
		err = errors.New("创建失败：必须设置一个或不设置")
		return
	}
	r, err := b.client.call("addnewaddr", []interface{}{account})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &addr)
	return
}

//查询节点钱包中的地址列表地址详细信息
//方法	                  参数
//addr(es)	    （显示所有地址）或addr(单个地址)或地址数组，默认为
//count	         显示地址的数量,默认为INT_MAX
//start	         从第几条开始，默认为-count
func (b *CXCcoind) ListAddressGroupings(count int, start int) (addr []model.Address, err error) {
	r, err := b.client.call("showaddrs", nil) //[]interface{}{count,start}
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &addr)
	return
}

//返回此账户节点钱包中所有资产余额的列表
//使用showallbals或者showaddrbals可查询用户余额,资产最小单位为0.000001。,
//参数				描述
//address		 查询地址
//minconf		 最小区块确认值  （选填）
//false/true 默认为false,若为true，则包含只读账户  （选填）
func (b *CXCcoind) GetBalance(address string) (asset []model.Asset, err error) {
	r, err := b.client.call("showaddrbals", []interface{}{address, 1, false})
	if err = handleError(err, &r); err != nil {
		log.Println("ShowAllBalls Error:", err.Error())
		return
	}
	err = json.Unmarshal(r.Result, &asset)
	return
}

//发送一定数量的资产到指定地址
//参数	           		描述
//address	       	地址
//asset-id	       	资产id或资产名称
//count				资产数量
//native-amount		原生资产数量（选填）
//comment			附加信息，不在区块链上（选填）
//comment-to		附加信息归属，不在区块链上（选填）
//如果支付成功，将返回交易ID<txid>。需要未锁定钱包。
func (b *CXCcoind) SendAsset(address, asset string, count float64) (txid string, err error) {
	r, err := b.client.call("sendasset", []interface{}{address, asset, count})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &txid)
	return
}

//指定地址发起转账
//	参数					  描述
//from-address				发起地址
//to-address				目标地址
//quantities-object			若直接填写数量，则为发送原生资产，或者填写{"asset-name" : asset-count,"": nativecurrency-count Native currency use "",...}
//comment					附加信息，不在区块链上
//comment-to				附加信息归属，不在区块链上
//如果支付成功，将返回交易ID<txid>。需要未锁定钱包。
func (b *CXCcoind) Send(formAddress, toAddress string, count float64) (txid string, err error) {
	r, err := b.client.call("sendfrom", []interface{}{formAddress, toAddress, count})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &txid)
	return
}

//walletpass   解锁钱包-时间
//password 解锁钱包的密码
//时间 1表示1分钟
func (b *CXCcoind) WalletPass(password string, time int) (rperr rpcError) {
	r, err := b.client.call("walletpass", []interface{}{password, time * 60})
	if err = handleError(err, &r); err != nil {
		log.Println("ShowAllBalls Error:", err.Error())
		return
	}
	err = json.Unmarshal(r.Result, &rperr)
	return
}

//锁定钱包
func (b *CXCcoind) WalletLock() error {
	r, err := b.client.call("walletlock", nil)
	return handleError(err, &r)
}

// 不可鎖定輸出
type UnspendableOutput struct {
	TxId string `json:"txid"`
	Vout uint64 `json:"vout"`
}

// 列出锁定的未动用输出
func (b *CXCcoind) ListLockUnspent() (unspendableOutputs []UnspendableOutput, err error) {
	r, err := b.client.call("listlockunspent", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &unspendableOutputs)
	return
}

//向地址发行一定数量的资产
//			参数		描述
//address	地址
//asset-configs	资产名称或资产参数{"name":"asset1","open":true}，若open为true,则可增发
//count	发行数量
//unit	最小单位，默认为1
//native-amount	原生资产数量
//extend-data	附加数据
/*func (b *CXCcoind) Sell(address string,asset model.AssetsInfo, count int,unit float32) (txid string, err error) {
	r, err := b.client.call("sell", []interface{}{address,asset,count,unit})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &txid)
	return
}
*/

//创建新的账户和地址，如果账户存在创建新的地址
/*func (b *CXCcoind) GetAccountAddress(account string) (address string, err error) {
	r, err := b.client.call("getaccountaddress", []interface{account})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &address)
	return
}
*/

//备份钱包 wallet.dat 到新的文件夹或者目录中
/*func (b *CXCcoind) BackupWallet(destination string) error {
	r, err := b.client.call("backupwallet", []string{destination})
	return handleError(err, &r)
}*/

//根据地址导出私钥
/*func (b *CXCcoind) DumPrivkey(address string) (privKey string, err error) {
	r, err := b.client.call("dumpprivkey", []string{address})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &privKey) //json 解析
	return
}*/

//钱包加密
/*func (b *CXCcoind) EncryptWallet(passphrass string) error {
	r, err := b.client.call("encryptwallet", []string{passphrass})
	return handleError(err, &r)
}*/

//获取和地址关联的账户
/*func (b *CXCcoind) GetAccount(address string) (account string, err error) {
	r, err := b.client.call("getaccount", []string{address})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &account)
	return
}*/

//"从指定账户<fromaccount>向指定地址<tobitcoinaddress>
//发送指定金额<amount>的BTC，确保帐户拥有得到[minconf]个确认的有效余额。
//支付金额是一个四舍五入至小数点后8位的实数。
//如果支付成功，将返回交易ID<txid>（而不是一个JSON对象）。需要未锁定钱包
//<amount> is a real and is rounded to 8 decimal places.
//Will send the given amount to the given address,
//ensuring the account has a valid balance using [minconf] confirmations.
//Returns the transaction ID if successful (not in JSON object)."
/*func (b *CXCcoind) SendFrom(fromAccount, toAddress string, amount float64, minconf uint32, comment, commentTo string) (txID string, err error) {
	r, err := b.client.call("sendfrom", []interface{}{fromAccount, toAddress, amount, minconf, comment, commentTo})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &txID)
	return
}*/
//"从指定账户<fromaccount>向多个地址发送指定金额{address:amount,...}。
//金额是双精度浮点数。需要未锁定钱包
//amounts are double-precision floating point numbers"
/*func (b *CXCcoind) SendMany(fromAccount string, amounts map[string]float64, minconf uint32, comment string) (txID string, err error) {
	r, err := b.client.call("sendmany", []interface{}{fromAccount, amounts, minconf, comment})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &txID)
	return
}*/

// "向地址<bitcoinaddress>发送指定金额<amount>的BTC。
//支付金额是一个四舍五入至小数点后8位的实数。
//如果支付成功，将返回交易ID<txid>。需要未锁定钱包。
//<amount> is a real and is rounded to 8 decimal places.
//Returns the transaction ID <txid> if successful. "
/*func (b *CXCcoind) SendToAddress(toAddress string, amount float64, comment, commentTo string) (txID string, err error) {
	r, err := b.client.call("sendtoaddress", []interface{}{toAddress, amount, comment, commentTo})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &txID)
	return
}
*/
// "将地址<bitcoinaddress>关联到指定帐户<account>。
//如果该地址已经被关联到指定帐户，将创建一个新的地址与该帐户关联。
//Sets the account associated with the given address.
// Assigning address that is already assigned to the same account
//will create a new address associated with that account. "
/*func (b *CXCcoind) SetAccount(address, account string) error {
	r, err := b.client.call("setaccount", []interface{}{address, account})
	return handleError(err, &r)
}*/
