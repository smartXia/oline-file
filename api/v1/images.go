package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"online_file/global"
	"online_file/model"
	"online_file/model/request"
	"online_file/model/response"
	"online_file/service"
)

// @Tags Images
// @Summary 创建Images
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Images true "创建Images"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /images/createImages [post]
func CreateImages(c *gin.Context) {
	var images model.Images
	_ = c.ShouldBindJSON(&images)
	if err := service.CreateImages(images); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Images
// @Summary 删除Images
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Images true "删除Images"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /images/deleteImages [delete]
func DeleteImages(c *gin.Context) {
	var images model.Images
	_ = c.ShouldBindJSON(&images)
	if err := service.DeleteImages(images); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Images
// @Summary 批量删除Images
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Images"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /images/deleteImagesByIds [delete]
func DeleteImagesByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteImagesByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Images
// @Summary 更新Images
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Images true "更新Images"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /images/updateImages [put]
func UpdateImages(c *gin.Context) {
	var images model.Images
	_ = c.ShouldBindJSON(&images)
	if err := service.UpdateImages(images); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Images
// @Summary 用id查询Images
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Images true "用id查询Images"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /images/findImages [get]
func FindImages(c *gin.Context) {
	var images model.Images
	_ = c.ShouldBindQuery(&images)
	if err, reimages := service.GetImages(images.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reimages": reimages}, c)
	}
}

// @Tags Images
// @Summary 分页获取Images列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ImagesSearch true "分页获取Images列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /images/getImagesList [get]
func GetImagesList(c *gin.Context) {
	var pageInfo request.ImagesSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetImagesInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
