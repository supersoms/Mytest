package router

import (
	. "CXCProject/apis"
	"CXCProject/middleware"
	"CXCProject/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	//go func() {}()
	//utils.SyncBlock() //程序一启动时，实时同步区块信息
	utils.WriterGinLogToFile()
	router := gin.Default()
	router.Use(middleware.IPAuthMiddleWare())
	router.POST("/game/userOauthLogin", UserOauthLogin) //用户APP授权登陆
	router.POST("/game/userLogin", UserLogin)           //用户登陆
	router.POST("/game/liveBetting", LiveBetting)       //获取实时投注记录（WebSocket长连接）所有人投注记录都可以看到
	router.GET("/game/myBetting", MyBetting)            //我的投注
	router.GET("/game/lotteryRecords", LotteryRecords)  //开奖记录
	router.GET("/game/getBlock", GetBlock)              //获取区块信息 (Websocket长连接)
	router.POST("/game/selectBet", SelectBet)           //选号下注(核心)
	return router
}
