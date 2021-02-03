package router

import (
	"github.com/gin-gonic/gin"
	"online_file/api/v1"
	"online_file/middleware"
)

func InitDocsRouter(Router *gin.RouterGroup) {
	DocsRouter := Router.Group("docs").Use(middleware.OperationRecord())
	{
		DocsRouter.POST("createDocs", v1.CreateDocs)             // 新建Docs
		DocsRouter.DELETE("deleteDocs", v1.DeleteDocs)           // 删除Docs
		DocsRouter.DELETE("deleteDocsByIds", v1.DeleteDocsByIds) // 批量删除Docs
		DocsRouter.PUT("updateDocs", v1.UpdateDocs)              // 更新Docs
		DocsRouter.GET("findDocs", v1.FindDocs)                  // 根据ID获取Docs
		DocsRouter.GET("getDocsList", v1.GetDocsList)            // 获取Docs列表
	}
}
