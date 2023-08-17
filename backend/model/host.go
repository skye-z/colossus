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
	Period string `json:"period"`

	// 访问地址
	Address string `json:"address"`
	// 连接端口
	Port int `json:"port"`
	// 授权方式
	AuthType string `json:"type"`
	// 登录用户
	User string `json:"user"`
	// 证书
	Key string `json:"-"`
	// 密钥
	Secret string `json:"-"`
}

type HostModel struct {
	DB *xorm.Engine
}

func (model HostModel) GetItem(host *Host) error {
	has, err := model.DB.Get(host)
	if !has {
		return err
	}
	return nil
}

func (model HostModel) GetList(name string, platform string, system string, region string, usage string, period string, address string, page int, num int) ([]Host, error) {
	var (
		hosts   []Host
		session *xorm.Session
	)
	// 拼接查询条件
	screen := ""
	if len(name) != 0 {
		screen += fmt.Sprintf("name like %%%s%% AND ", name)
	}
	if len(platform) != 0 {
		screen += fmt.Sprintf("platform = %s AND ", platform)
	}
	if len(system) != 0 {
		screen += fmt.Sprintf("system = %s AND ", system)
	}
	if len(region) != 0 {
		screen += fmt.Sprintf("region = %s AND ", region)
	}
	if len(usage) != 0 {
		screen += fmt.Sprintf("usage = %s AND ", usage)
	}
	if len(period) != 0 {
		screen += fmt.Sprintf("period = %s AND ", period)
	}
	if len(address) != 0 {
		screen += fmt.Sprintf("address like %%%s%% AND ", address)
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
