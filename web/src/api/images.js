import service from '@/utils/request'

// @Tags Images
// @Summary 创建Images
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Images true "创建Images"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /images/createImages [post]
export const createImages = (data) => {
     return service({
         url: "/images/createImages",
         method: 'post',
         data
     })
 }


// @Tags Images
// @Summary 删除Images
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Images true "删除Images"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /images/deleteImages [delete]
 export const deleteImages = (data) => {
     return service({
         url: "/images/deleteImages",
         method: 'delete',
         data
     })
 }

// @Tags Images
// @Summary 删除Images
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Images"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /images/deleteImages [delete]
 export const deleteImagesByIds = (data) => {
     return service({
         url: "/images/deleteImagesByIds",
         method: 'delete',
         data
     })
 }

// @Tags Images
// @Summary 更新Images
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Images true "更新Images"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /images/updateImages [put]
 export const updateImages = (data) => {
     return service({
         url: "/images/updateImages",
         method: 'put',
         data
     })
 }


// @Tags Images
// @Summary 用id查询Images
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Images true "用id查询Images"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /images/findImages [get]
 export const findImages = (params) => {
     return service({
         url: "/images/findImages",
         method: 'get',
         params
     })
 }


// @Tags Images
// @Summary 分页获取Images列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Images列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /images/getImagesList [get]
 export const getImagesList = (params) => {
     return service({
         url: "/images/getImagesList",
         method: 'get',
         params
     })
 }