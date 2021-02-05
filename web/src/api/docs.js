import service from '@/utils/request'

// @Tags Docs
// @Summary 创建Docs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Docs true "创建Docs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /docs/createDocs [post]
export const createDocs = (data) => {
     return service({
         url: "/docs/createDocs",
         method: 'post',
         data
     })
 }


// @Tags Docs
// @Summary 删除Docs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Docs true "删除Docs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /docs/deleteDocs [delete]
 export const deleteDocs = (data) => {
     return service({
         url: "/docs/deleteDocs",
         method: 'delete',
         data
     })
 }

// @Tags Docs
// @Summary 删除Docs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Docs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /docs/deleteDocs [delete]
 export const deleteDocsByIds = (data) => {
     return service({
         url: "/docs/deleteDocsByIds",
         method: 'delete',
         data
     })
 }

// @Tags Docs
// @Summary 更新Docs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Docs true "更新Docs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /docs/updateDocs [put]
 export const updateDocs = (data) => {
     return service({
         url: "/docs/updateDocs",
         method: 'put',
         data
     })
 }


// @Tags Docs
// @Summary 用id查询Docs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Docs true "用id查询Docs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /docs/findDocs [get]
 export const findDocs = (params) => {
     return service({
         url: "/docs/findDocs",
         method: 'get',
         params
     })
 }


// @Tags Docs
// @Summary 分页获取Docs列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Docs列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /docs/getDocsList [get]
 export const getDocsList = (params) => {
     return service({
         url: "/docs/getDocsList",
         method: 'get',
         params
     })
 }