package model

//格式化数据
type Work struct {
	Midstate string `json:"midstate"` //状态
	Data     string `json:"data"`     //数据
	Hash1    string `json:"hash1"`    //hash
	Target   string `json:"target"`   //目标
}
