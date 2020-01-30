package model

//Vout 表示输出数量
type Vout struct {
	ID           int          `orm:"column(id);pk,auto" description:"ID"  json:"id,omitempty"`
	Value        float64      `orm:"column(value);" description:"输出数量" json:"value"` //输出数量
	N            int          `orm:"column(n);p" description:"N"json:"n"`
	ScriptPubKey ScriptPubKey `json:"scriptPubKey"`    //公钥
	Data         []string     `json:"data,omitempty" ` //数据
	Assets       []AssetsInfo `json:"assets"`          //输出资产
}
