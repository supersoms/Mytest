package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//自定义IP白名单中间件
func IPAuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		ipList := []string{ //ip白名单列表
			"127.0.0.1",
		}
		flag := false
		clientIP := c.ClientIP() //拿到客户端请求的IP地址
		for _, host := range ipList {
			if clientIP == host {
				flag = true
				break
			}
		}
		if !flag { //客户端请求的ip不在白名单名表中，flag就一直为false
			c.String(http.StatusUnauthorized, "IP地址：%v no in ip list", clientIP)
			c.Abort() //中止
		}
	}
}
