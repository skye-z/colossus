// 授权凭证
package model

import (
	"fmt"

	"xorm.io/xorm"
)

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

type CertModel struct {
	DB *xorm.Engine
}

// 获取凭证列表
func (model CertModel) GetList(keyword string, page, num int) ([]Code, error) {
	var (
		codes   []Code
		session *xorm.Session
	)
	// 拼接查询条件
	screen := ""
	if len(keyword) != 0 {
		screen += fmt.Sprintf("name LIKE \"%%%s%%\" ", keyword)
	}
	// 判断是否存在查询条件
	if len(screen) == 0 {
		session = model.DB.Limit(page*num, (page-1)*num)
	} else {
		session = model.DB.Where(screen[0:len(screen)-4]).Limit(page*num, (page-1)*num)
	}
	// 执行查询
	err := session.Find(&codes)
	if err != nil {
		return nil, err
	}
	return codes, nil
}

// 获取指定凭证
func (model CertModel) GetItem(cert *Cert) error {
	has, err := model.DB.Get(cert)
	if !has {
		return err
	}
	return nil
}
