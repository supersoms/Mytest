/***************************************************
 ** @Desc :
 ** @Time : 2020/1/8
 ** @Author : Administrator
 ** @File : chain.go
 ** @Last Modified by : Administrator
 ** @Last Modified time: 2020-01-08-18:06
 ** @Software: GoLand
****************************************************/
package model

//区块同步信息
type BlockChain struct {
	Blocks        uint64 `json:"blocks"`        //已同步的区块数量
	Headers       string `json:"headers"`       //区块头的数量
	Bestblockhash string `json:"bestblockhash"` //最新的区块HASH
}
