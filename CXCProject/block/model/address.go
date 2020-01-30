/***************************************************
 ** @Desc :
 ** @Time : 2020/1/8
 ** @Author : Administrator
 ** @File : address.go
 ** @Last Modified by : Administrator
 ** @Last Modified time: 2020-01-08-23:31
 ** @Software: GoLand
****************************************************/
package model

//地址信息
type Address struct {
	Address      string `json:"address"`                //验证的地址
	IsMine       bool   `json:"ismine"`                 //是否自己的地址
	IsScript     bool   `json:"isscript"`               //是否是脚本
	PubKey       string `json:"pubkey"`                 //原始公钥的十六进制值
	IsCompressed bool   `json:"iscompressed"`           //是否额外创建地址
	Synchronized bool   `json:"synchronized,omitempty"` //是否同步完成
	IsValid      bool   `json:"isvalid,omitempty"`      //字段为false，则为无效地址。
	Iswatchonly  bool   `json:"iswatchonly"`            //地址只能查看
}

//钱包地址列表
type ListAddressResult struct {
	Address string
	Amount  float64
	Account string
}
