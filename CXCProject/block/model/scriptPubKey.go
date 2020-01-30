package model

//签名公钥公钥
type ScriptPubKey struct {
	ID        int       `orm:"column(id);pk,auto" description:"ID"  json:"id,omitempty"`
	Asm       string    `orm:"column(asm);size(100)" description:"asm" json:"asm"`
	Hex       string    `orm:"column(hex);size(100)" description:"hex" json:"hex"`
	ReqSigs   int       `orm:"column(reqsigs);" description:"返回的签名" json:"reqSigs,omitempty"`
	Type      string    `orm:"column(type);size(100)" description:"类型pubKey" json:"type"`
	Addresses []Address `orm:"column(address);" description:"address" json:"addresses,omitempty"`
}
