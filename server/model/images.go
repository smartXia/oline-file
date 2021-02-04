// 自动生成模板Images
package model

import (
	"online_file/global"
)

// 如果含有time.Time 请自行import time包
type Images struct {
	global.GvaModel
	Appkey     string `json:"appkey" form:"appkey" gorm:"column:appkey;comment:中台标识;type:varchar(32);size:32;"`
	Channel    int    `json:"channel" form:"channel" gorm:"column:channel;comment:渠道标识;type:int;size:10;"`
	AccountId  string `json:"accountId" form:"accountId" gorm:"column:account_id;comment:账户id;type:varchar(32);size:32;"`
	FunctionId string `json:"functionId" form:"functionId" gorm:"column:function_id;comment:模块id;type:varchar(32);size:32;"`
	ImageType  string `json:"imageType" form:"imageType" gorm:"column:image_type;comment:图片类型:image&card;type:varchar(50);size:50;"`
	Name       string `json:"name" form:"name" gorm:"column:name;comment:描述;type:text;"`
	Url        string `json:"url" form:"url" gorm:"column:url;comment:图片url;type:varchar(255);size:255;"`
	Width      string `json:"width" form:"width" gorm:"column:width;comment:图片宽;type:varchar(32);size:32;"`
	Height     string `json:"height" form:"height" gorm:"column:height;comment:图片高;type:varchar(32);size:32;"`
	PositionX  string `json:"positionX" form:"positionX" gorm:"column:position_x;comment:图片展示x轴长度;type:varchar(32);size:32;"`
	PositionY  string `json:"positionY" form:"positionY" gorm:"column:position_y;comment:图片展示y轴长度;type:varchar(32);size:32;"`
	UniqueId   string `json:"uniqueId" form:"uniqueId" gorm:"column:unique_id;comment:图片唯一标识;type:varchar(32);size:32;"`
	Group      string `json:"group" form:"group" gorm:"column:group;comment:内置图片;type:varchar(255);size:255;"`
	Cover      string `json:"cover" form:"cover" gorm:"column:cover;comment:组封面图;type:varchar(255);size:255;"`
	TaskId     int    `json:"taskId" form:"taskId" gorm:"column:task_id;comment:图片来源哪个任务ID;type:int;size:10;"`
	TaskType   string `json:"taskType" form:"taskType" gorm:"column:task_type;comment:任务类型;type:varchar(32);size:32;"`
}

func (Images) TableName() string {
	return "images"
}

// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type ImagesWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	Images   `json:"business"`
// }

// func (Images) TableName() string {
// 	return "images"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["images"] = func() model.GVA_Workflow {
//   return new(model.ImagesWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["images"] = func() interface{} {
// 	return new(model.Images)
// }
