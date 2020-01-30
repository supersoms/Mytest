package apis

import (
	"CXCProject/client"
	"CXCProject/message"
	model "CXCProject/models"
	"CXCProject/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

//用户授权登陆
func UserOauthLogin(c *gin.Context) {
	code := c.Request.FormValue("code")             //786976478
	securityId := c.Request.FormValue("securityId") //8b945a3f43abbc6b8357
	if len(code) == 0 {
		c.String(http.StatusBadRequest, "param code is not null")
		c.Abort()
		return
	} else if len(securityId) == 0 {
		c.String(http.StatusBadRequest, "param securityId is not null")
		c.Abort()
		return
	}
	params := make(map[string]string)
	params["code"] = code
	params["securityId"] = securityId
	response, err := client.Get(utils.USER_OAUTH_LOGIN, params, nil)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	} else if response != nil && response.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(response.Body)
		var oauth model.Oauth
		if err = json.Unmarshal(body, &oauth); err != nil {
			c.String(http.StatusInternalServerError, "JSON数据解析错误：", err.Error())
			c.Abort()
			return
		} else {
			user := oauth.Result
			if len(user.Cid) == 0 {
				c.String(http.StatusBadRequest, "param cid is not null")
				c.Abort()
				return
			} else if len(user.Address) == 0 {
				c.String(http.StatusBadRequest, "param address is not null")
				c.Abort()
				return
			}
			err = user.OauthLogin()
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code":    message.USER_OAUTH_LOGIN_FAIL_CODE,
					"message": err.Error(),
				})
				log.Println("user oauth login fail!")
				c.Abort()
				return
			} else {
				log.Println("user oauth login success!")
				c.JSON(http.StatusOK, gin.H{
					"code":    message.USER_OAUTH_LOGIN_SUCCESS_CODE,
					"message": message.USER_OAUTH_LOGIN_SUCCESS_MSG,
					"data":    oauth,
				})
			}
		}
	}
	defer response.Body.Close()
}

//用户登陆
func UserLogin(c *gin.Context) {
	cid := c.Request.FormValue("cid")
	adress := c.Request.FormValue("address")
	if len(cid) == 0 {
		c.String(http.StatusBadRequest, "param cid is not null")
		c.Abort()
		return
	} else if len(adress) == 0 {
		c.String(http.StatusBadRequest, "param address is not null")
		c.Abort()
		return
	}
	user := model.User{Cid: cid, Address: adress}
	err := user.Login()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		log.Println("user login fail!")
		return
	} else {
		log.Println("user login success!")
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": message.USER_OAUTH_LOGIN_SUCCESS_MSG,
		})
	}
}
