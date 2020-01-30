package models

import (
	orm "CXCProject/database"
	"CXCProject/message"
	"errors"
	"log"
)

//下注实体
type Bet struct {
	Cid            string  `form:"cid" json:"cid"`                                                   //投注账号,唯一
	BigOrSmall     string  `form:"bigOrSmall" json:"bigOrSmall"`                                     //大小，传big为大, small为小
	SingleOrDouble string  `form:"singleOrDouble" json:"singleOrDouble"`                             //单双，传single为单，double为双
	BetQuantity    int     `form:"betQuantity" binding:"required,gt=0" json:"betQuantity"`           //投注数量
	BettingPeriod  int64   `form:"bettingPeriod" json:"bettingPeriod"`                               //投注期号
	BettingContent string  `form:"bettingContent" json:"bettingContent"`                             //投注内容(传数字0-9)用map来存
	BettingDate    string  `form:"bettingDate" time_format:"2006-01-01 15:04:05" json:"bettingDate"` //投注日期
	BetAmount      int     `form:"betAmount" binding:"required,gt=0" json:"betAmount"`               //投注金额
	WinningAmount  float64 `form:"winningAmount" json:"winningAmount"`                               //中奖金额
}

//实时投注
func (bet *Bet) LiveBetting() (err error) {
	result := orm.DB.Create(&bet)
	if result.Error != nil {
		log.Println(result.Error)
		err = errors.New(message.LIVEBETTING_FAILED_MSG)
	}
	return
}

//获取上一期投注期号
func (bet *Bet) GetLastTermBettingPeriod() (betInfo Bet, err error) {
	result := orm.DB.Order("betting_period desc").Limit(1).Find(&betInfo)
	if result.Error != nil {
		log.Println(result.Error)
		err = errors.New(message.GET_USER_BETTING_RECORDING_FAILED_MSG)
	}
	return
}

/*func (user *User) Users() (users []User, err error) {
	if err = orm.DB.Find(&users).Error; err != nil {
		return
	}
	return
}
*/






//我的投注记录
func (bet *Bet) MyBetting(cid string) (bets []Bet, err error) {
	//先检查用户是否已注册
	result := orm.DB.Where("cid = ?", cid).Find(&bets)
	if result.Error != nil {
		log.Println(result.Error)
		err = errors.New(message.GET_USER_BETTING_RECORDING_FAILED_MSG)
	}
	return
}

//开奖记录
func (bet *Bet) LotteryRecords(cid string) (bets []Bet, err error) {
	//先检查用户是否已注册
	result := orm.DB.Where("cid = ?", cid).Find(&bets)
	if result.Error != nil {
		log.Println(result.Error)
		err = errors.New(message.GET_USER_BETTING_RECORDING_FAILED_MSG)
	}
	return
}

//实时投注
func (bet *Bet) CreateOrder() (err error) {
	result := orm.DB.Create(&bet)
	if result.Error != nil {
		log.Println(result.Error)
		err = errors.New(message.LIVEBETTING_FAILED_MSG)
	}
	return
}