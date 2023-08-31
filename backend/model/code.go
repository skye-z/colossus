// 快捷代码
package model

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
