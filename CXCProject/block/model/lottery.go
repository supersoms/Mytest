package model

import (
	"github.com/astaxie/beego/orm"
	"log"
)

// RawTx 表示交易
//id,periods,blockhash,blockheight,placeone,placetow,placethree,placefour,placefive,betamount,profit,createtime
type Lottery struct {
	ID          int64   `orm:"column(id);pk,auto" description:"ID"  json:"id,omitempty"`
	Periods     int64   `orm:"column(periods);"  description:"期数" json:"periods"`                              //交易ID
	BlockHash   string  `orm:"column(blockhash);size(100)"  description:"区块Hash" json:"blockhash,omitempty"`   //区块Hash
	BlockHeight uint64  `orm:"column(blockheight);size(100)"  description:"区块高度" json:"blockheight,omitempty"` //区块Hash
	PlaceOne    string  `orm:"column(placeone);size(100)"  json:"placeone"`                                    //
	PlaceTow    string  `orm:"column(placetow);size(100)" json:"placetow"`                                     //
	PlaceThree  string  `orm:"column(placethree);size(100)" json:"placethree"`                                 //
	PlaceFour   string  `orm:"column(placefour);size(100)" json:"placefour"`                                   //
	PlaceFive   string  `orm:"column(placefive);size(100)" json:"placefive"`                                   //
	BetAmount   float64 `orm:"column(betamount);" json:"betamount"`
	Profit      float64 `orm:"column(confirmations);"  description:"确认次数"  json:"confirmations,omitempty"` //
	CreateTime  int64   `orm:"column(createtime);"  description:"时间" json:"blocktime,omitempty"`           //创建时间
}

//获取表名
func GetLotteryTableName() string {
	return getTable("lottery")
}

func NewLottery() *Lottery {
	return &Lottery{}
}

//添加彩票记录
func AddLottery(this Lottery) int64 {
	row, err := orm.NewOrm().Insert(&this)
	if err != nil {
		log.Panicln("AddLotery ERROR:", err.Error())
	}
	return row
}

//获取彩票期数
func GetLotteryPeriods() (peridos int64) {
	qs := orm.NewOrm().QueryTable(GetLotteryTableName())
	peridos, _ = qs.Count()
	return
}
