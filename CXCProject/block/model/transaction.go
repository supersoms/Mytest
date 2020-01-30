package model

import (
	"github.com/astaxie/beego/orm"
	"log"
)

// RawTx 表示交易
type RawTransaction struct {
	ID            int    `orm:"column(id);pk,auto" description:"ID"  json:"id,omitempty"`
	Txid          string `orm:"column(txid);size(100)"  description:"TXID" json:"txid"`                       //交易ID
	Version       uint32 `orm:"column(version);"  description:"版本" json:"version"`                            //版本
	LockTime      uint32 `orm:"column(locktime);"  description:"锁定时间" json:"locktime"`                        //锁定时间秒
	Vin           []Vin  `orm:"-" json:"vin"`                                                                 //交易输入
	Vout          []Vout `orm:"-" json:"vout"`                                                                //交易输出
	BlockHash     string `orm:"column(blockhash);size(100)"  description:"区块Hash" json:"blockhash,omitempty"` //区块Hash
	Confirmations uint64 `orm:"column(confirmations);"  description:"确认次数"  json:"confirmations,omitempty"`   //区块确认数量
	Time          int64  `orm:"column(time);"  description:"交易时间" json:"time,omitempty"`                      //时间
	Blocktime     int64  `orm:"column(blocktime);"  description:"区块时间" json:"blocktime,omitempty"`            //区块时间
	Hex           string `orm:"column(hex);size(100)"  description:"Hex原始数据" json:"hex"`                      //原始数据
}

//获取表名
func GetRawTransactionTableName() string {
	return getTable("raw_transaction")
}

func NewRawTransaction() *RawTransaction {
	return &RawTransaction{}
}

//添加交易记录
func AddRawTransaction(this RawTransaction) int64 {

	row, err := orm.NewOrm().Insert(&this)
	if err != nil {
		log.Panicln("ERROR:", err.Error())
	}
	return row
}

//获取区块列表
func GetRawTransactionList() (list []RawTransaction, row int64) {
	qs := orm.NewOrm().QueryTable(GetRawTransactionTableName())
	row, err := qs.All(&list)
	if row == 0 {
		return nil, 0
	}
	if err != nil {
		log.Panic("GetRawTransactionList ERROR:", err.Error())
	}
	return
}

//获交易总数
func GetRawTransactionCount() (count int64) {
	qs := orm.NewOrm().QueryTable(GetRawTransactionTableName())
	count, _ = qs.Count()
	return
}

//根据TXID 获取交易信息
func GetRawTransaction(txid string) (rt RawTransaction, err error) {
	rt.Txid = txid
	err = orm.NewOrm().QueryTable(GetRawTransactionTableName()).One(&rt)

	if err != nil {
		log.Println("GetRawTransaction ERROR:", err.Error())
	}
	return
}

//获取最新的交易
func GeRawTransactionLimt(i int) (blocks []RawTransaction, err error) {

	qs := orm.NewOrm().QueryTable(GetRawTransactionTableName()).Limit(i)
	_, err = qs.All(&blocks)

	if err != nil {
		log.Println("GeRawTransactionLimt ERROR:", err.Error())
	}

	return
}

//======================================================================================================================
//  公钥脚本
type ScriptSig struct {
	Asm string `json:"asm"` //
	Hex string `json:"hex"` //禁止
}

//交易详情
type TransactionDetails struct {
	Account  string  `json:"account"`
	Address  string  `json:"address,omitempty"`
	Category string  `json:"category"`
	Amount   float64 `json:"amount"`
	Fee      float64 `json:"fee,omitempty"`
}

// 交易
type Transaction struct {
	TxID            string               `json:"txid"`               //交易TXID
	Amount          float64              `json:"amount"`             //交易总量
	Account         string               `json:"account,omitempty"`  //交易账户
	Address         string               `json:"address,omitempty"`  //地址
	Category        string               `json:"category,omitempty"` //交易类型
	Fee             float64              `json:"fee,omitempty"`      //	交易手续费
	Confirmations   int64                `json:"confirmations"`      //确认数量
	BlockHash       string               `json:"blockhash"`          //区块hash
	BlockIndex      int64                `json:"blockindex"`         //区块索引
	BlockTime       int64                `json:"blocktime"`          //区块时间
	WalletConflicts []string             `json:"walletconflicts"`    //
	Time            int64                `json:"time"`               //交易时间
	TimeReceived    int64                `json:"timereceived"`       //确认时间
	Details         []TransactionDetails `json:"details,omitempty"`  //事务数据
	Hex             string               `json:"hex,omitempty"`      //带有签名的原始处理(十六进制编码的字符串)
}

// 交易输出 (UTXO)
type UTransactionOut struct {
	Bestblock     string       `json:"bestblock"`
	Confirmations uint32       `json:"confirmations"`
	Value         float64      `json:"value"`
	ScriptPubKey  ScriptPubKey `json:"scriptPubKey"`
	Version       uint32       `json:"version"`
	Coinbase      bool         `json:"coinbase"`
}

// TransactionOutSet 关于未使用的事务输出数据库的统计信息
type TransactionOutSet struct {
	Height          uint32  `json:"height"`
	Bestblock       string  `json:"bestblock"`
	Transactions    float64 `json:"transactions"`
	TxOuts          float64 `json:"txouts"`
	BytesSerialized float64 `json:"bytes_serialized"`
	HashSerialized  string  `json:"hash_serialized"`
	TotalAmount     float64 `json:"total_amount"`
}
