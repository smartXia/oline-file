package service

import (
	"online_file/global"
	"online_file/model"
	"online_file/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateDocs
//@description: 创建Docs记录
//@param: docs model.Docs
//@return: err error

func CreateDocs(docs model.Docs) (err error) {
	err = global.GVA_DB.Create(&docs).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDocs
//@description: 删除Docs记录
//@param: docs model.Docs
//@return: err error

func DeleteDocs(docs model.Docs) (err error) {
	err = global.GVA_DB.Delete(&docs).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDocsByIds
//@description: 批量删除Docs记录
//@param: ids request.IdsReq
//@return: err error

func DeleteDocsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Docs{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateDocs
//@description: 更新Docs记录
//@param: docs *model.Docs
//@return: err error

func UpdateDocs(docs model.Docs) (err error) {
	err = global.GVA_DB.Save(&docs).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDocs
//@description: 根据id获取Docs记录
//@param: id uint
//@return: err error, docs model.Docs

func GetDocs(id uint) (err error, docs model.Docs) {
	err = global.GVA_DB.Where("id = ?", id).First(&docs).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDocsInfoList
//@description: 分页获取Docs记录
//@param: info request.DocsSearch
//@return: err error, list interface{}, total int64

func GetDocsInfoList(info request.DocsSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Docs{})
	var docss []model.Docs
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&docss).Error
	return err, docss, total
}
