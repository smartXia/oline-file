package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"online_file/global"
	"online_file/model"
	"online_file/model/response"
	"online_file/service"
)

// @Tags Template
// @Summary 创建模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Images true "创建Images"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /images/createImages [post]
func CreateTemplate(c *gin.Context) {
	var template model.Template
	_ = c.ShouldBindJSON(&template)
	if err := service.CreateTemplate(template); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
