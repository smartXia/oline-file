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

// @Tags Docs
// @Summary 创建Docs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Docs true "创建Docs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /docs/createDocs [post]
func CreateDocs(c *gin.Context) {
	var docs model.Docs
	_ = c.ShouldBindJSON(&docs)
	if err := service.CreateDocs(docs); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Docs
// @Summary 删除Docs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Docs true "删除Docs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /docs/deleteDocs [delete]
func DeleteDocs(c *gin.Context) {
	var docs model.Docs
	_ = c.ShouldBindJSON(&docs)
	if err := service.DeleteDocs(docs); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Docs
// @Summary 批量删除Docs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Docs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /docs/deleteDocsByIds [delete]
func DeleteDocsByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteDocsByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Docs
// @Summary 更新Docs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Docs true "更新Docs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /docs/updateDocs [put]
func UpdateDocs(c *gin.Context) {
	var docs model.Docs
	_ = c.ShouldBindJSON(&docs)
	if err := service.UpdateDocs(docs); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Docs
// @Summary 用id查询Docs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Docs true "用id查询Docs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /docs/findDocs [get]
func FindDocs(c *gin.Context) {
	var docs model.Docs
	_ = c.ShouldBindQuery(&docs)
	if err, redocs := service.GetDocs(docs.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redocs": redocs}, c)
	}
}

// @Tags Docs
// @Summary 分页获取Docs列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DocsSearch true "分页获取Docs列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /docs/getDocsList [get]
func GetDocsList(c *gin.Context) {
	var pageInfo request.DocsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetDocsInfoList(pageInfo); err != nil {
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
