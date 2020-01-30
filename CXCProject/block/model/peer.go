package model

//钱包同步节点信息
type Peer struct {
	Id             int    `json:"id"`             //ID
	Addr           string `json:"addr" `          //同步的IP地址和端口号
	Addrlocal      string `json:"addrlocal" `     //本地钱包地址和端口号
	Services       string `json:"services"`       //服务器
	Lastsend       uint64 `json:"lastsend"`       //最后发送一次的时间
	Lastrecv       uint64 `json:"lastrecv"`       //最后一次接受时间
	Bytessent      uint64 `json:"bytessent"`      //发送的总字节数
	Bytesrecv      uint64 `json:"bytesrecv"`      //接收的总字节数
	Conntime       uint64 `json:"conntime"`       //连接时间
	Pingtime       uint64 `json:"pingtime"`       //ping时间
	PingWait       uint64 `json:"ping_wait"`      //ping 等待时间
	Version        uint32 `json:"version"`        //同级版本，比如7001
	Subver         string `json:"subver"`         //版本
	Inbound        bool   `json:"inbound"`        // 转入 或者转出
	Startingheight int32  `json:"startingheight"` //节点的起始高度
	Banscore       int32  `json:"banscore"`       //交易数量
	Syncnode       bool   `json:"syncnode"`       //是否进行节点同步
	Synced_headers int 		`json:"synced_headers"` //同步区块头
	Synced_blocks  int		`json:"synced_blocks"`  //同步区块
}
/*
[
    {

        "handlocal" : null,
        "handshake" : null,
        "inflight" : [
        ],
        "whitelisted" : false
    }
]

*/