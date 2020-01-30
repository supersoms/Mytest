package model

//Vin 表示 输入值
type Vin struct {
	ID       int    `orm:"column(id);pk,auto" description:"ID"  json:"id,omitempty"`
	Txid     string `orm:"column(txid);size(100)" description:"ID" json:"txid"`         //交易ID
	Coinbase string `orm:"column(coinbase);size(100)" description:"ID" json:"coinbase"` //交易基础
	Sequence uint32 `orm:"column(sequence);" description:"交易序列" json:"sequence"`        //交易序列
}
