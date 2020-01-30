/***************************************************
 ** @Desc :CXC节点工具操作类
 ** @Time : 2020/1/9
 ** @Author : Administrator
 ** @File : tool.go
 ** @Last Modified by : Administrator
 ** @Last Modified time: 2020-01-09-11:16
 ** @Software: GoLand
****************************************************/
package block

import (
	"CXCProject/block/model"
	"encoding/json"
	"errors"
	"log"
)

//获取节点信息
func (b *CXCcoind) GetInfo() (i model.Info, err error) {
	r, err := b.client.call("showinfo", nil)

	if err = handleError(err, &r); err != nil {
		log.Println("ShowInfo Error:", err.Error())
		return
	}
	//fmt.Printf("%s",r.Result)
	err = json.Unmarshal(r.Result, &i)
	return
}

// 验证检测账户地址的有效性
//请求问题：ValidateAddress Error: HTTP error: 500 Internal Server Error
func (b *CXCcoind) ValidateAddress(address string) (vadd model.Address, err error) {
	if len(address) < 34 {
		err = errors.New("请输入完整的验证地址")
		return
	}

	r, err := b.client.call("validaddr", []interface{}{address})
	//log.Printf("请求结果：%s",r.Result)
	if err = handleError(err, &r); err != nil {
		log.Println("ValidateAddress Error:", err.Error())
		return
	}
	err = json.Unmarshal(r.Result, &vadd)
	return
}
