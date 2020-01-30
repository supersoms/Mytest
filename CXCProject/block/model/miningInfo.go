package model

//挖矿信息
type MiningInfo struct {
	Blocks           uint64  `json:"blocks"`           //当前区块
	CurrentBlocksize uint64  `json:"currentblocksize"` //当前区块大小
	CurrentBlockTx   uint64  `json:"currentblocktx"`   //当前区块交易
	Difficulty       float64 `json:"difficulty"`       //当前难度
	Errors           string  `json:"errors"`           //错误信息
	GenProcLimit     int32   `json:"genproclimit"`     //处理器对生成的限制。-1如果没有创世区块。
	PooledtTx        uint64  `json:"pooledtx"`         // 挖矿 mem池的大小
	Testnet          bool    `json:"testnet"`          // 是否使用测试网络
	Generate         bool    `json:"generate"`         //是否开启挖矿 getgenerate or setgenerate
	NetworkHashps    uint64  `json:"networkhashps"`    //网络hash率
	HashesPersec     uint64  `json:"hashespersec"`     //节点 hash率
}
