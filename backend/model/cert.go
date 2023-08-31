// 授权凭证
package model

type Cert struct {
	// 编号
	Id int64 `json:"id"`
	// 名称
	Name string `json:"name"`
	// 公钥
	PublicKey string `json:"publicKey"`
	// 私钥
	PrivateKey string `json:"privateKey"`
}
