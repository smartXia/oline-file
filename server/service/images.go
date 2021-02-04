package service

import (
	"online_file/global"
	"online_file/model"
	"online_file/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateImages
//@description: 创建Images记录
//@param: images model.Images
//@return: err error

func CreateImages(images model.Images) (err error) {
	err = global.GVA_DB.Create(&images).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteImages
//@description: 删除Images记录
//@param: images model.Images
//@return: err error

func DeleteImages(images model.Images) (err error) {
	err = global.GVA_DB.Delete(&images).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteImagesByIds
//@description: 批量删除Images记录
//@param: ids request.IdsReq
//@return: err error

func DeleteImagesByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Images{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateImages
//@description: 更新Images记录
//@param: images *model.Images
//@return: err error

func UpdateImages(images model.Images) (err error) {
	err = global.GVA_DB.Save(&images).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetImages
//@description: 根据id获取Images记录
//@param: id uint
//@return: err error, images model.Images

func GetImages(id uint) (err error, images model.Images) {
	err = global.GVA_DB.Where("id = ?", id).First(&images).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetImagesInfoList
//@description: 分页获取Images记录
//@param: info request.ImagesSearch
//@return: err error, list interface{}, total int64

func GetImagesInfoList(info request.ImagesSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Images{})
	var imagess []model.Images
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.TaskId != 0 {
		db = db.Where("`task_id` = ?", info.TaskId)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&imagess).Error
	return err, imagess, total
}
