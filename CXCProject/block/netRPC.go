/***************************************************
 ** @Desc : 节点网络操作
 ** @Time : 2020/1/9
 ** @Author : Administrator
 ** @File : netRPC.go
 ** @Last Modified by : Administrator
 ** @Last Modified time: 2020-01-09-13:10
 ** @Software: GoLand
****************************************************/
package block

import (
	"CXCProject/block/model"
	"encoding/json"
)

//ping网络连接（无效不适用）
func (b *CXCcoind) Ping() error {
	r, err := b.client.call("ping", nil)
	return handleError(err, &r)
}

//获取节点信息
func (b *CXCcoind) GetPeerInfo() (peerInfo []model.Peer, err error) {
	r, err := b.client.call("showpeer", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &peerInfo)
	return
}
