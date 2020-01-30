package utils

import (
	"CXCProject/block"
	"CXCProject/block/model"
	"bytes"
	"log"
	"math/rand"
	"time"
)

//一分钟以后获取一个区块，彩票服务节点不能停机
func LotteryCheck(hash string) []string {
	if hash == "" {
		log.Println("Hash参数不能为空！")
		return nil
	}
	//随机获取5个数字作为开奖结果
	var lucky []string
	hashluck := bytes.NewBufferString(hash).Bytes()
	c := 0
	for i := 0; i < len(hash); i++ {
		l := rand.Intn(len(hash))
		if IsNumeric(string(hashluck[l])) {
			if c >= 5 {
				break
			}
			lucky = append(lucky, string(hashluck[l]))
			log.Println("随机数哈希位置：", l, " 幸运号码：", string(hashluck[l]))
			c++
		}
	}
	log.Println("开奖Hash：", hash)
	return lucky
}

//定时同步区块任务
func SyncBlock() {
	bc, err := block.New(SERVER_HOST, SERVER_PORT, USER, PASSWD, USESSL)
	if err != nil {
		//TODO 代办：如果钱包服务器链接失败，此时要启动钱包，startWattle()
		log.Panicln("钱包服务器链接失败！", err.Error())
	}
	//periods := 1 + apis.GetLotteryPeriods() //获取数据库开奖期号
	periods:=int64(1)
	blockTime := time.Now().Unix()
	log.Println("彩票开始时间:", blockTime)
	for {
		time.Sleep(time.Millisecond * 10)
		count, _ := bc.GetBlockCount() //获取链上区块总数
		i := model.GetBlockCount()     //数据读取区块总数
		if uint64(i) <= count {
			blocks, _ := bc.GetBlocks(uint64(i)) //获取区块的最新哈希
			//获取完整的区块信息
			block, _ := bc.GetBlock(blocks[0].Hash)
			//区块信息添加到数据库
			model.AddBlock(block)
			log.Println("最新区块：", block.Height, "同步时间：", block.Time)
			//交易记录添加到数据库
			for i := 0; i < len(block.TX); i++ {
				block.TX[i].BlockHash = block.Hash //交易区块
				block.TX[i].Time = block.Time
				block.TX[i].Confirmations = block.Confirmations
				model.AddRawTransaction(block.TX[i])
				//log.Println("最新交易TXID ：",block.TX[i].Txid,"同步成功！")
			}
			//获取每一分钟开始的第一个开奖区块
			//t1 := time.Now().Unix()//当前时间戳
			if block.Time > blockTime { //block.Time >=blockTime+60  &&
				lucky := LotteryCheck(block.Hash)
				log.Println("第", periods, "期，开奖时间：", block.Time, "区块高度：", block.Height, "开奖号码分别是：", lucky)
				//将彩票添加到数据库
				lottery := model.Lottery{
					Periods:     periods,
					BlockHash:   block.Hash,
					BlockHeight: block.Height,
					PlaceOne:    lucky[0],
					PlaceTow:    lucky[1],
					PlaceThree:  lucky[2],
					PlaceFour:   lucky[3],
					PlaceFive:   lucky[4],
					BetAmount:   0,
					Profit:      0,
					CreateTime:  block.Time,
				}
				//TODO 代办：将获取的区块信息写入数据中，还需将区块数据写一份到redis中
				if model.AddLottery(lottery) < 1 {
					log.Println("入彩票失败！")
				}
				periods = periods + 1
				blockTime = block.Time + 20
			}
			if uint64(i) == count {
				log.Println("没有最新数据，休息一会！")
				time.Sleep(time.Second * 5)
			}
		}
	}
}