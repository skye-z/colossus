// 快捷代码
package model

import (
	"fmt"

	"xorm.io/xorm"
)

type Code struct {
	// 编号
	Id int64 `json:"id"`
	// 名称
	Name string `json:"name"`
	// 关键词
	Keyword string `json:"keyword"`
	// 运行平台
	Platform string `json:"platform"`
	// 操作系统
	System string `json:"system"`
	// 内容
	Content string `json:"content"`
}

type CodeModel struct {
	DB *xorm.Engine
}

// 获取命令列表
func (model CodeModel) GetList(keyword, platform, system string, page, num int) ([]Code, error) {
	var (
		codes   []Code
		session *xorm.Session
	)
	// 拼接查询条件
	screen := ""
	if len(keyword) != 0 {
		screen += fmt.Sprintf("(name LIKE \"%%%s%%\" OR keyword LIKE \"%%%s%%\")AND ", keyword, keyword)
	}
	if len(platform) != 0 {
		screen += fmt.Sprintf("platform = \"%s\" OR ", platform)
	}
	if len(system) != 0 {
		screen += fmt.Sprintf("system = \"%s\" OR ", system)
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

// 新增命令
func (model CodeModel) Add(code *Code) bool {
	_, err := model.DB.Insert(code)
	return err == nil
}

// 编辑命令
func (model CodeModel) Edit(code *Code) bool {
	if code.Id == 0 {
		return false
	}
	_, err := model.DB.ID(code.Id).AllCols().Update(code)
	return err == nil
}

// 删除命令
func (model CodeModel) Del(id int64) bool {
	if id == 0 {
		return false
	}
	code := Code{
		Id: id,
	}
	_, err := model.DB.Delete(code)
	return err == nil
}
