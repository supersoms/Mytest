package model

import (
	"github.com/astaxie/beego/orm"
	"log"
)

//区块信息
type Block struct {
	Height        uint64           `orm:"column(height);pk" description:"区块高度" json:"height"`                                // 区块高度
	Hash          string           `orm:"column(hash);size(100)"  description:"区块hash" json:"hash"`                          // 区块hash
	Miner         string           `orm:"column(miner);size(100)"  description:"挖矿者" json:"miner"`                           // 挖矿者
	Confirmations uint64           `orm:"column(confirmations)"  description:"确认数量" json:"confirmations"`                    // 确认数量
	Size          uint64           `orm:"column(size)"  description:"区块大小" json:"size"`                                      // 区块大小
	Version       uint32           `orm:"column(version)"  description:"区块版本" json:"version"`                                // 区块版本
	Merkleroot    string           `orm:"column(merkleroot);size(100)"  description:"Merkle根" json:"merkleroot"`             // Merkle根
	TX            []RawTransaction `orm:"-" json:"tx,omitempty"`                                                             // 区块里面的交易
	Time          int64            `orm:"column(time)"  description:"区块hash" json:"time"`                                    // 区块产生时间
	Nonce         uint64           `orm:"column(nonce)"  description:"确认数量" json:"nonce"`                                    // 确认数量
	Bits          string           `orm:"column(bits);size(100)"  description:"单位 bits" json:"bits"`                         // 单位 bits
	Difficulty    float64          `orm:"column(difficulty);digits(40);decimals(16)"  description:"难度" json:"difficulty"`    // 难度
	Chainwork     string           `orm:"column(chainwork);size(100)"  description:"区块验证信息，十六进制" json:"chainwork,omitempty"` //区块验证信息，十六进制
	PrevBlockHash string           `orm:"column(prevblockhash);size(100)"  description:"上一个区块块的哈希" json:"prevblockhash"`     // 上一个区块块的哈希
	NextBlockHash string           `orm:"column(nextblockhash);size(100)"  description:"下一个区块哈希" json:"nextblockhash"`       // 下一个区块哈希
}

//获取表名
func GetBlockTableName() string {
	return getTable("block")
}

func NewBlokc() *Block {
	return &Block{}
}

//获取区块列表
func GetBlokList() (list []Block, row int64) {
	qs := orm.NewOrm().QueryTable(GetBlockTableName())
	row, err := qs.All(&list)
	if row == 0 {
		return nil, 0
	}
	if err != nil {
		log.Panic("GetBlockList ERROR:", err.Error())
	}
	return
}

//获取区块总数
func GetBlockCount() (count int64) {
	qs := orm.NewOrm().QueryTable(GetBlockTableName())
	count, _ = qs.Count()
	return
}

//根据区块高度获取区块信息
func GetBlok(height uint64) (block Block, err error) {
	block.Height = height
	err = orm.NewOrm().QueryTable(GetBlockTableName()).One(&block)
	if err != nil {
		log.Println("GetBlock ERROR:", err.Error())
	}
	return
}

//根据区块高度获取区块信息
func GetBlokLimt(i int) (blocks []Block, err error) {
	qs := orm.NewOrm().QueryTable(GetBlockTableName()).Limit(i)
	_, err = qs.All(&blocks)
	if err != nil {
		log.Println("GetBlock ERROR:", err.Error())
	}
	log.Println("GetBlock DATA:", blocks)
	return
}

//添加区块新
//返回受影响行数
func AddBlock(this Block) int64 {
	row, err := orm.NewOrm().Insert(&this)
	if err != nil {
		log.Panicln("ERROR:", err.Error())
	}
	return row
}
