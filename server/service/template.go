package service

import (
	"online_file/global"
	"online_file/model"
	"online_file/model/request"
)

func CreateTemplate(template model.Template) (err error) {
	err = global.GVA_DB.Create(&template).Error
	return err
}
func DeleteTemplate(template model.Template) (err error) {
	err = global.GVA_DB.Delete(&template).Error
	return err
}
func DeleteTemplateByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Template{}, "id in ?", ids.Ids).Error
	return err
}
func updateTemplate(template model.Template) (err error) {
	err = global.GVA_DB.Create(&template).Error
	return err
}
func getTemplate(id uint) (err error, template model.Template) {
	err = global.GVA_DB.Where("id = ?", id).First(&template).Error
	return
}
func getTemplateList(info request.ImagesSearch) (err error, list interface{}, total int64) {

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.Template{})
	var template []model.Template
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&template).Error
	return err, template, total
}
