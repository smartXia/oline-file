package model

// 自动生成模板

import (
	"online_file/global"
)

// 如果含有time.Time 请自行import time包
type Template struct {
	global.GvaModel
	AccountId string `json:"accountId" form:"accountId" gorm:"column:account_id;comment:中台用户唯一标识;type:varchar(32);size:32;"`
	Appkey    string `json:"appkey" form:"appkey" gorm:"column:appkey;comment:通道标识;type:varchar(32);size:32;"`
	Channel   int    `json:"channel" form:"channel" gorm:"column:channel;comment:通道标识;type:int;size:10;"`
	Name      string `json:"name" form:"name" gorm:"column:name;comment:标题;type:varchar(255);size:255;"`
	Type      string `json:"type" form:"type" gorm:"column:type;comment:类型;type:int;size:4;"`
	Content   string `json:"content" form:"content" gorm:"column:content;comment:;type:text(4294967295);size:4294967295;"`
	Guid      string `json:"guid" form:"guid" gorm:"column:guid;comment:文章唯一标识;type:varchar(32);size:32;"`
	State     string `json:"state" form:"state" gorm:"column:state;comment:state ;type:int;size:4;"`
	Sign      string `json:"sign" form:"sign" gorm:"column:sign;comment:时间戳;type:varchar(50);size:50;"`
}

func (Template) TableName() string {
	return "template"
}
