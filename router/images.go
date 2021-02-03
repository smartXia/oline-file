package router

import (
	"github.com/gin-gonic/gin"
	"online_file/api/v1"
	"online_file/middleware"
)

func InitImagesRouter(Router *gin.RouterGroup) {
	ImagesRouter := Router.Group("xa_images").Use(middleware.OperationRecord())
	{
		ImagesRouter.POST("createImages", v1.CreateImages)   // 新建Images
		ImagesRouter.DELETE("deleteImages", v1.DeleteImages) // 删除Images
		ImagesRouter.DELETE("deleteImagesByIds", v1.DeleteImagesByIds) // 批量删除Images
		ImagesRouter.PUT("updateImages", v1.UpdateImages)    // 更新Images
		ImagesRouter.GET("findImages", v1.FindImages)        // 根据ID获取Images
		ImagesRouter.GET("getImagesList", v1.GetImagesList)  // 获取Images列表
	}
}
