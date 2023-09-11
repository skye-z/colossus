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
	// 算法类型
	Type string `json:"type"`
	// 复杂度
	Level int16 `json:"level"`
	// 公钥
	PublicKey string `json:"publicKey"`
	// 私钥
	PrivateKey string `json:"privateKey"`
	// 密码
	Passphrase string `json:"passphrase"`
	// 有效期限
	Period int64 `json:"period"`
}

type CertModel struct {
	DB *xorm.Engine
}

// 获取凭证列表
func (model CertModel) GetList(keyword string, page, num int) ([]Cert, error) {
	var (
		certs   []Cert
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
	err := session.Find(&certs)
	if err != nil {
		return nil, err
	}
	return certs, nil
}

// 获取指定凭证
func (model CertModel) GetItem(cert *Cert) error {
	has, err := model.DB.Get(cert)
	if !has {
		return err
	}
	return nil
}

// 新增命令
func (model CertModel) Add(cert *Cert) bool {
	_, err := model.DB.Insert(cert)
	return err == nil
}

// 编辑命令
func (model CertModel) Edit(cert *Cert) bool {
	if cert.Id == 0 {
		return false
	}
	_, err := model.DB.ID(cert.Id).AllCols().Update(cert)
	return err == nil
}

// 删除命令
func (model CertModel) Del(id int64) bool {
	if id == 0 {
		return false
	}
	cert := Cert{
		Id: id,
	}
	_, err := model.DB.Delete(cert)
	return err == nil
}
