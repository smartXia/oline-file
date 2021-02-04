// 自动生成模板Docs
package model

import (
	"online_file/global"
)

// 如果含有time.Time 请自行import time包
type Docs struct {
	global.GvaModel
	AccountId  string `json:"accountId" form:"accountId" gorm:"column:account_id;comment:中台用户唯一标识;type:varchar(32);size:32;"`
	Appkey     string `json:"appkey" form:"appkey" gorm:"column:appkey;comment:通道标识;type:varchar(32);size:32;"`
	Channel    int    `json:"channel" form:"channel" gorm:"column:channel;comment:通道标识;type:int;size:10;"`
	Content    string `json:"content" form:"content" gorm:"column:content;comment:;type:text(4294967295);size:4294967295;"`
	Guid       string `json:"guid" form:"guid" gorm:"column:guid;comment:文章唯一标识;type:varchar(32);size:32;"`
	InviteCode string `json:"inviteCode" form:"inviteCode" gorm:"column:invite_code;comment:邀请码;type:varchar(50);size:50;"`
	IsDeleted  int    `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:是否被删除;type:int;size:10;"`
	IsEditable int    `json:"isEditable" form:"isEditable" gorm:"column:is_editable;comment:是否可编辑;type:int;size:10;"`
	Name       string `json:"name" form:"name" gorm:"column:name;comment:标题;type:varchar(255);size:255;"`
	ParentId   string `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:父级标识;type:varchar(50);size:50;"`
	ShareMode  string `json:"shareMode" form:"shareMode" gorm:"column:share_mode;comment:分享 public/private;type:varchar(50);size:50;"`
	Sign       string `json:"sign" form:"sign" gorm:"column:sign;comment:时间戳;type:varchar(50);size:50;"`
	Type       string `json:"type" form:"type" gorm:"column:type;comment:类型;type:varchar(50);size:50;"`
}

func (Docs) TableName() string {
	return "docs"
}

// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type DocsWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	Docs   `json:"business"`
// }

// func (Docs) TableName() string {
// 	return "docs"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["docs"] = func() model.GVA_Workflow {
//   return new(model.DocsWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["docs" = func() interface{} {
// 	return new(model.Docs)
// }
