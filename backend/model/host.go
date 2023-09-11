// 主机
package model

import (
	"fmt"

	"xorm.io/xorm"
)

type Host struct {
	// 编号
	Id int64 `json:"id"`
	// 名称
	Name string `json:"name"`

	// 运行平台
	Platform string `json:"platform"`
	// 操作系统
	System string `json:"system"`
	// 所在地区
	Region string `json:"region"`
	// 主机用途
	Usage string `json:"usage"`
	// 有效期限
	Period int64 `json:"period"`

	// 访问地址
	Address string `json:"address"`
	// 连接端口
	Port int `json:"port"`
	// 授权方式
	AuthType string `json:"type"`
	// 登录用户
	User string `json:"user"`
	// 证书
	Cert int64 `json:"cert"`
	// 密钥
	Secret string `json:"-"`
}

type AddHost struct {
	// 编号
	Id int64 `json:"id"`
	// 名称
	Name string `json:"name"`

	// 运行平台
	Platform string `json:"platform"`
	// 操作系统
	System string `json:"system"`
	// 所在地区
	Region string `json:"region"`
	// 主机用途
	Usage string `json:"usage"`
	// 有效期限
	Period int64 `json:"period"`

	// 访问地址
	Address string `json:"address"`
	// 连接端口
	Port int `json:"port"`
	// 授权方式
	AuthType string `json:"type"`
	// 登录用户
	User string `json:"user"`
	// 证书
	Cert int64 `json:"cert"`
	// 密钥
	Secret string `json:"secret"`
}

type HostModel struct {
	DB *xorm.Engine
}

// 获取指定主机
func (model HostModel) GetItem(host *Host) error {
	has, err := model.DB.Get(host)
	if !has {
		return err
	}
	return nil
}

// 获取主机列表
func (model HostModel) GetList(keyword, platform, system, region, usage, period string, page, num int) ([]Host, error) {
	var (
		hosts   []Host
		session *xorm.Session
	)
	// 拼接查询条件
	screen := ""
	if len(keyword) != 0 {
		screen += fmt.Sprintf("(name LIKE \"%%%s%%\" OR address LIKE \"%%%s%%\")AND ", keyword, keyword)
	}
	if len(platform) != 0 {
		screen += fmt.Sprintf("platform = \"%s\" AND ", platform)
	}
	if len(system) != 0 {
		screen += fmt.Sprintf("system = \"%s\" AND ", system)
	}
	if len(region) != 0 {
		screen += fmt.Sprintf("region = \"%s\" AND ", region)
	}
	if len(usage) != 0 {
		screen += fmt.Sprintf("usage = \"%s\" AND ", usage)
	}
	if len(period) != 0 {
		screen += fmt.Sprintf("period = \"%s\" AND ", period)
	}
	// 判断是否存在查询条件
	if len(screen) == 0 {
		session = model.DB.Limit(page*num, (page-1)*num)
	} else {
		session = model.DB.Where(screen[0:len(screen)-4]).Limit(page*num, (page-1)*num)
	}
	// 执行查询
	err := session.Find(&hosts)
	if err != nil {
		return nil, err
	}
	return hosts, nil
}

// 添加主机
func (model HostModel) Add(host *Host) bool {
	_, err := model.DB.Insert(host)
	return err == nil
}

// 编辑主机
func (model HostModel) Edit(host *Host) bool {
	if host.Id == 0 {
		return false
	}
	_, err := model.DB.ID(host.Id).AllCols().Update(host)
	return err == nil
}

// 删除主机
func (model HostModel) Del(id int64) bool {
	if id == 0 {
		return false
	}
	host := Host{
		Id: id,
	}
	_, err := model.DB.Delete(host)
	return err == nil
}
