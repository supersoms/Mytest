package models

import (
	orm "CXCProject/database"
	"CXCProject/message"
	"errors"
	"log"
)

//支付订单结构体
type PayOrder struct {
	Cid        string `form:"cid" json:"cid"`                                                 //关联投注账号
	OrderId    string `form:"orderId" json:"orderId"`                                         //订单单号
	CreateDate string `form:"createDate" time_format:"2006-01-01 15:04:05" json:"createDate"` //订单创建日期
	IsPay      bool   `form:"default:false"`                                                  //是否支付
	PayDate    string `form:"payDate" time_format:"2006-01-01 15:04:05" json:"payDate"`       //支付日期
	PayAddress string `json:"payAddress" form:"payAddress" binding:"required`                 //支付账号，就是一个地址
	PayAmount  int    `form:"payAmount" json:"payAmount"`                                     //支付金额
}

//实时投注
func (payOrder *PayOrder) CreatePayOrder() (err error) {
	result := orm.DB.Create(&payOrder)
	if result.Error != nil {
		log.Println(result.Error)
		err = errors.New(message.USER_SELECT_BET_FAIL_MSG)
	}
	return
}
