package model

//获取节点信息
type Info struct {
	/*	Version         uint32  `json:"version"`                  //服务器版
		Protocolversion uint32  `json:"protocolversion"`          //协议版本
		Walletversion   uint32  `json:"walletversion"`            //钱包版本
		Balance         float64 `json:"balance"`                  //钱包总余额
		Blocks          uint32  `json:"blocks"`                   //当前服务器中块数
		Timeoffset      int32   `json:"timeoffset"`               //时间偏移值
		Connections     uint32  `json:"connections"`              // 连接数量
		Proxy           string  `json:"proxy,omitempty"`          //使用代理服务器地址端口
		Difficulty      float64 `json:"difficulty"`               // 当前难度
		Testnet         bool    `json:"testnet"`                  // 是否使用测试网络
		Keypoololdest   uint64  `json:"keypoololdest"`            //矿池中最晚生成密钥的预计时间
		KeypoolSize     uint32  `json:"keypoolsize,omitempty"`    // 预计挖矿生成多少个币
		Errors          string  `json:"errors"`                   // 错误信息*/

	Paytxfee      string `json:"paytxfee"`                 // 交易手续费设置 btc/kb
	Rpcport       int    `json:"rpcport"`                  //RPC远程端口
	Relayfee      string `json:"relayfee"`                 //交易手续费 btc/kb
	UnlockedUntil int64  `json:"unlocked_until,omitempty"` //解锁钱包交易
	Maxout        int64  `json:"maxout"`                   //最大交易输出
}

//获取挖掘信息的响应
/*type Info struct {
	Version         uint32  `json:"version"`                  //服务器版
	Protocolversion uint32  `json:"protocolversion"`          //协议版本
	Walletversion   uint32  `json:"walletversion"`            //钱包版本
	Balance         float64 `json:"balance"`                  //钱包总余额
	Blocks          uint32  `json:"blocks"`                   //当前服务器中块数
	Timeoffset      int32   `json:"timeoffset"`               //时间偏移值
	Connections     uint32  `json:"connections"`              // 连接数量
	Proxy           string  `json:"proxy,omitempty"`          //使用代理服务器地址端口
	Difficulty      float64 `json:"difficulty"`               // 当前难度
	Testnet         bool    `json:"testnet"`                  // 是否使用测试网络
	Keypoololdest   uint64  `json:"keypoololdest"`            //矿池中最晚生成密钥的预计时间
	KeypoolSize     uint32  `json:"keypoolsize,omitempty"`    // 预计挖矿生成多少个币
	UnlockedUntil   int64   `json:"unlocked_until,omitempty"` //解锁钱包交易
	Paytxfee        float64 `json:"paytxfee"`                 // 交易手续费设置 btc/kb
	Relayfee        float64 `json:"relayfee"`                 //交易手续费 btc/kb
	Errors          string  `json:"errors"`                   // 错误信息
}*/
