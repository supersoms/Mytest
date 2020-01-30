/***************************************************
 ** @Desc :链上发行的资产
 ** @Time : 2020/1/9
 ** @Author : Administrator
 ** @File : asset.go
 ** @Last Modified by : Administrator
 ** @Last Modified time: 2020-01-09-17:32
 ** @Software: GoLand
****************************************************/

package model

//钱包地址资产信息
type AssetsInfo struct {
	Name     string   `json:"name"`     //资产名称         "name" : "ETH",
	Selltxid string   `json:"selltxid"` //发行TXID         "selltxid" : "0f09931f3ebd9d744dfcbf6e8cad640e8332a03fbf00c288a34b322f7d6d88ff",
	Assetref string   `json:"assetref"` //资产序号         				   "assetref" : "96-300-2319",
	Multiple uint     `json:"multiple"` //倍数        	                   "multiple" : 1000000,
	Units    float64  `json:"units"`    //计量单位         				   "units" : 0.000001,
	Open     bool     `json:"open"`     //若open为true,则可增发             "open" : true,
	Restrict Restrict `json:"restrict"` //限制条件         "restrict" : {"send" : false,   "receive" : false  },
	Details  []string `json:"details"`  //详细说明         "details" : {      },
	Sellqty  float64  `json:"sellqty"`  //发行数量         "sellqty" : 2310.411745,
	Sellraw  uint     `json:"sellraw"`  //发行数量原始数据  "sellraw" : 2310411745,
	Ordered  bool     `json:"ordered"`  //是否可以预定      "ordered" : false
}

//账户资产详情
type Asset struct {
	Assetref string `json:"assetref"` //资产序号
	QTY      string `json:"qty"`      //数量        qty
	RAW      uint   `json:"raw"`      //原始数据    raw,
}

//限制条件
type Restrict struct {
	Send    bool `json:"send"`
	Receive bool `json:"receive"`
}
