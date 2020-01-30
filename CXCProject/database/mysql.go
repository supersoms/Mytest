package database

import (
	"CXCProject/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //加载mysql驱动
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

//用户表
type User struct {
	ID           int    `gorm:"primary_key:id;auto_increment:number"`
	Cid          string `gorm:"type:varchar(255);primary_key;not null;"`
	Address      string `gorm:"type:varchar(255);not null;unique"`         //64位的一个地址
	InviteCode   string `gorm:"type:varchar(255)"`                         //允许为空，邀请码6位
	Level        int    `gorm:"default:'1'"`                               //允许为空，会员等级,1-9级,默认是1级
	InvitePeople string `gorm:"type:varchar(255);default:'zhongsheng007'"` //允许为空，邀请人,默认官方指定的账户
}

//会员等级表(后台用)
type MemberLevel struct {
	Level       int     `gorm:"number;not null;default:'1'"`             //等级(1-9级数字)
	LevelName   string  `gorm:"type:varchar(255);not null;default:'一级'"` //等级名称(一级到九级中文)
	BetAmount   float64 `gorm:"not null"`                                //投注金额
	RewardRatio float64 `gorm:"null;default:'0'"`                        //返奖比例
}

//投注表
/*type Bet struct {
	Cid            string  `gorm:"type:varchar(255);not null;column:cid;unique"`     //设置列名为`column:cid`投注账号,唯一
	BetQuantity    int     `gorm:"not null;column:betQuantity"`                      //投注数量
	BettingPeriod  int64   `gorm:"not null;column:bettingPeriod"`                    //投注期号
	BettingContent string  `gorm:"type:varchar(255);not null;column:bettingContent"` //投注内容
	BettingDate    string  `gorm:"not null;column:bettingDate"`                      //投注日期
	WinningDisplay string  `gorm:"type:varchar(255);not null;column:winningDisplay"` //中奖显示，是否中奖
	BetAmount      float64 `gorm:"not null;column:betAmount"`                        //投注金额
}*/

//投注表
type Bet struct {
	ID             int     `gorm:"primary_key:id;auto_increment:number"`
	Cid            string  `gorm:"type:varchar(255);not null;"` //投注账号
	BigOrSmall     string  `gorm:"type:varchar(255);null"`      //大小，big为大, Small为小
	SingleOrDouble string  `gorm:"type:varchar(255);null"`      //单双，single为单，Double为双
	BetQuantity    int     `gorm:"not null"`                    //投注数量
	BettingPeriod  int64   `gorm:"not null"`                    //投注期号
	BettingContent string  `gorm:"not null"`                    //投注内容(0-9个数字，如果为空，)
	BettingDate    string  `gorm:"type:varchar(255);not null"`  //投注日期
	BetAmount      int     `gorm:"not null"`                    //投注金额
	WinningAmount  float64 `gorm:"null"`                        //中奖金额
}

//开奖表
type Lottery struct {
	LotteryDate        uint   `gorm:"not null"`                   //开奖日期
	LotteryNumber      string `gorm:"type:varchar(255);not null"` //开奖内容(5个数字)
	WinningNumberIndex string `gorm:"type:varchar(255);not null"` //中奖号码的索引(对应的key,value)使用map
	LotteryPeriod      int64  `gorm:"not null;unique"`            //开奖期号
	BlockHeight        uint64 `gorm:"not null"`                   //区块高度(用于前端显示)
}

//区块信息表
type Block struct {
	Height        uint64 `gorm:"pk;not null" json:"height"`              //区块高度
	Hash          string `gorm:"type:varchar(255);not null" json:"hash"` //区块hash
	Confirmations uint64 `gorm:"not null" json:"confirmations"`          //确认数量
}

// RawTx 表示交易表
type RawTransaction struct {
	Cid             string `gorm:"type:varchar(255);not null" json:"cid"`
	Txid            string `gorm:"type:varchar(255);primary_key;not null" json:"txid"` //交易ID，主键，唯一
	BlockHash       string `gorm:"type:varchar(255);not null" json:"blockHash"`        //区块Hash
	TransactionTime int    `gorm:"not null" json:"transactionTime"`                    //交易时间
}

//支付订单
type PayOrder struct {
	ID         int    `gorm:"primary_key:id;auto_increment:number"`
	Cid        string `gorm:"type:varchar(255);not null;"`       //关联投注账号
	OrderId    string `gorm:"type:varchar(255);not null;unique"` //订单单号
	CreateDate string `gorm:"type:varchar(255);not null"`        //订单创建日期
	IsPay      bool   `gorm:"default:false"`                     //是否支付
	PayDate    string `gorm:"not null"`                          //支付日期
	PayAddress string `gorm:"type:varchar(255);not null"`        //支付账号，就是一个地址
	PayAmount  int    `form:"not null"`                          //支付金额
}

func init() {
	config, err := utils.ReadDBConfig() //也可以通过os.arg或flag从命令行指定配置文件路径
	if err != nil {
		log.Fatal(err)
	}
	conn := config.DBUser + ":" + config.DBPwd + "@tcp(" + config.DBHost + ":" + config.DBPort + ")/" + config.DBName + "?charset=utf8&parseTime=True&loc=Local&timeout=10ms"
	//create database `cxcdb` character set utf8 collate utf8_general_ci;
	//DB, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/cxcdb?charset=utf8&parseTime=True&loc=Local&timeout=10ms")
	DB, err = gorm.Open("mysql", conn)
	if err != nil {
		fmt.Println("mysql connect err：", err)
		return
	}
	if DB.Error != nil {
		fmt.Println("database err：", DB.Error)
		return
	}

	// 全局禁用表名复数，如果设置为true,`User`的默认表名为`user`，为false，创建的表名会多加个s
	DB.SingularTable(true)

	if !DB.HasTable("User") { //如果没表User表，则创建
		//DB.AutoMigrate(&User{}) //自动生成表
		DB.Table("user").CreateTable(&User{}) //使用User结构体创建名为`user`的表
		log.Println("init create table user success!")
	}
	if !DB.HasTable("MemberLevel") { //如果没表User表，则创建
		//DB.AutoMigrate(&MemberLevel{}) //自动生成表
		DB.Table("memberlevel").CreateTable(&MemberLevel{}) //使用User结构体创建名为`user`的表
		log.Println("init create table memberlevel success!")
	}
	if !DB.HasTable("Bet") {
		//DB.AutoMigrate(&Bet{})
		DB.Table("bet").CreateTable(&Bet{})
		log.Println("init create table bet success!")
	}
	if !DB.HasTable("Lottery") {
		//DB.AutoMigrate(&Lottery{})
		DB.Table("lottery").CreateTable(&Lottery{})
		log.Println("init create table lottery success!")
	}
	if !DB.HasTable("Block") {
		//DB.AutoMigrate(&Block{})
		DB.Table("block").CreateTable(&Block{})
		log.Println("init create table block success!")
	}
	if !DB.HasTable("RawTransaction") {
		//DB.AutoMigrate(&RawTransaction{})
		DB.Table("rawTransaction").CreateTable(&RawTransaction{})
		log.Println("init create table rawtransaction success!")
	}
	if !DB.HasTable("pay_order") {
		//DB.AutoMigrate(&PayOrder{})
		DB.Table("pay_order").CreateTable(&PayOrder{})
		log.Println("init create table pay_order success!")
	}
}
