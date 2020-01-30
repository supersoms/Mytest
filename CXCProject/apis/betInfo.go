package apis

import (
	"CXCProject/message"
	model "CXCProject/models"
	"CXCProject/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

//获取实时投注记录(UI前端滚动显示)
func LiveBetting(c *gin.Context) {
	var bet model.Bet
	//if err := c.ShouldBind(&betInfo); err != nil {
	//	c.String(http.StatusOK, "betInfo binding err：%s", err.Error())
	//	c.Abort()
	//	return
	//}
	data := c.Request.FormValue("betData")
	if err := json.Unmarshal([]byte(data), &bet); err != nil {
		c.String(http.StatusInternalServerError, "JSON数据解析错误：", err.Error())
		c.Abort()
		log.Println(err.Error())
		return
	}
	if len(bet.Cid) == 0 {
		c.String(http.StatusBadRequest, "param cid is not null")
		c.Abort()
		return
	} else if bet.BetQuantity <= 0 {
		c.String(http.StatusBadRequest, "param betQuantity is not null")
		c.Abort()
		return
	} else if bet.BettingPeriod <= 0 {
		c.String(http.StatusBadRequest, "param bettingPeriod is not null")
		c.Abort()
		return
		/*else if bet.BettingContent ==nil {
			c.String(http.StatusBadRequest, "param bettingContent is not null")
			c.Abort()
			return
		} */
	} else if bet.BetAmount <= 0 {
		c.String(http.StatusBadRequest, "param betAmount is not null")
		c.Abort()
		return
	}
	err := bet.LiveBetting()
	if err != nil {
		log.Println("live betting fail!")
		c.JSON(http.StatusOK, gin.H{
			"code":    message.USER_BETTING_FAIL_CODE,
			"message": err.Error(),
		})
		return
	} else {
		log.Println("user live betting success!")
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": message.LIVEBETTING_SUCCESS_MSG,
			"data":    bet,
		})
	}
}

//我的投注
func MyBetting(c *gin.Context) {
	cid := c.Query("cid")
	if len(cid) == 0 {
		c.String(http.StatusBadRequest, "param cid is not null")
		c.Abort()
		return
	}
	var bet model.Bet
	result, err := bet.MyBetting(cid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		log.Println("get my betting fail!")
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"data":    result,
			"total":   len(result),
			"message": "获取当前登录用户投注记录成功!",
		})
	}
}

//开奖记录
func LotteryRecords(c *gin.Context) {
	var bet model.Bet
	result, err := bet.LotteryRecords("1000")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		log.Println("get my betting fail!")
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"data":    result,
			"message": message.GET_USER_BETTING_RECORDING_SUCCESS_MSG,
		})
	}
}

//获取最新的开奖号码
func GetLotteryPeriods(c *gin.Context) int64 {
	return 0
}

//选号下注(核心) 要用多线程处理下注逻辑
func SelectBet(c *gin.Context) {
	//TODO 首先判断用户是否登录，还未实现
	betInfo := c.Request.FormValue("betInfo")
	if len(betInfo) == 0 {
		c.String(http.StatusBadRequest, "param betInfo is not null")
		c.Abort()
		return
	}
	bet := model.Bet{}
	if err := json.Unmarshal([]byte(betInfo), &bet); err != nil {
		c.String(http.StatusInternalServerError, "JSON数据解析错误：", err.Error())
		c.Abort()
		return
	} else {
		//TODO 校验用户传的数据
		if bet.BetQuantity <= 0 {
			c.String(http.StatusInternalServerError, "param betQuantity is not null")
			c.Abort()
			return
		}
		if bet.BettingPeriod <= 0 {
			c.String(http.StatusInternalServerError, "param bettingPeriod is not null")
			c.Abort()
			return
		} else {
			betData, _ := bet.GetLastTermBettingPeriod() //获取数据库中最后投注期数
			if bet.BettingPeriod < betData.BettingPeriod {
				c.String(http.StatusInternalServerError, "投注期号(bettingPeriod)不能小于上一期投注期号!")
				c.Abort()
				return
			}
		}
		if bet.BettingContent == "" {
			c.String(http.StatusInternalServerError, "投注内容(bettingContent)不能为空!")
			c.Abort()
			return
		} else {
			if strings.ToLower(bet.BigOrSmall) == "small" && bet.BettingContent != "0,1,2,3,4" {
				c.String(http.StatusInternalServerError, "投注内容(bettingContent)与所选号码不符!")
				c.Abort()
				return
			} else if strings.ToLower(bet.BigOrSmall) == "big" && bet.BettingContent != "5,6,7,8,9" {
				c.String(http.StatusInternalServerError, "投注内容(bettingContent)与所选号码不符!")
				c.Abort()
				return
			} else if strings.ToLower(bet.SingleOrDouble) == "single" && bet.BettingContent != "1,3,5,7,9" {
				c.String(http.StatusInternalServerError, "投注内容(bettingContent)与所选号码不符!")
				c.Abort()
				return
			} else if strings.ToLower(bet.SingleOrDouble) == "double" && bet.BettingContent != "0,2,4,6,8" {
				c.String(http.StatusInternalServerError, "投注内容(bettingContent)与所选号码不符!")
				c.Abort()
				return
			}
		}
		if bet.BettingDate == "" {
			c.String(http.StatusInternalServerError, "投注日期(bettingDate)不能为空!")
			c.Abort()
			return
		}
		if bet.BetAmount <= 0 {
			c.String(http.StatusInternalServerError, "投注金额(betAmount)不能为空!")
			c.Abort()
			return
		} else if bet.BetAmount != bet.BetQuantity {
			c.String(http.StatusInternalServerError, "投注金额(betAmount)计算错误!")
			c.Abort()
			return
		}
		//TODO 1：先将选号下注数据插入到数据库中，同时向订单表插入数据(下单时间，支付时间，是否支付，单号，cid，)，用一个状态来标识是否支付，然后App调用CXC钱包发起支付，
		//TODO 2：如果支付成功，重新调用一个接口将订单表的状态修改，同时更新支付时间等信息
		intChan := make(chan int, 10)
		err = bet.LiveBetting()
		intChan <- 1

		//Cid        string `form:"cid" json:"cid"`                                                 //关联投注账号
		//OrderId    string `form:"orderId" json:"orderId"`                                         //订单单号
		//CreateDate string `form:"createDate" time_format:"2006-01-01 15:04:05" json:"createDate"` //订单创建日期
		//IsPay      bool   `form:"default:false"`                                                  //是否支付
		//PayDate    string `form:"payDate" time_format:"2006-01-01 15:04:05" json:"payDate"`       //支付日期
		//PayAccount string `json:"payAccount" form:"cid" binding:"required`                        //支付账号，就是一个地址
		//PayAmount  int    `form:"payAmount" json:"payAmount"`                                     //支付金额
		payOrder := model.PayOrder{Cid: bet.Cid, OrderId: "CXCBET" + utils.Time2StrLink() + utils.GetNanoTimestamp(), CreateDate: utils.Time2Str()}
		payOrder.CreatePayOrder()
		<-intChan
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
			log.Println("选号下注失败!")
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": message.USER_SELECT_BET_SUCCESS_MSG,
			})
			log.Println("选号下注成功!")
		}
	}
}
