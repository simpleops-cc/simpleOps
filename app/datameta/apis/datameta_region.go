package apis

import (
    "fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/datameta/models"
	"go-admin/app/datameta/service"
	"go-admin/app/datameta/service/dto"
	"go-admin/common/actions"
)

type DatametaRegion struct {
	api.Api
}

// GetPage 获取DatametaRegion列表
// @Summary 获取DatametaRegion列表
// @Description 获取DatametaRegion列表
// @Tags DatametaRegion
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.DatametaRegion}} "{"code": 200, "data": [...]}"
// @Router /api/v1/datameta-region [get]
// @Security Bearer
func (e DatametaRegion) GetPage(c *gin.Context) {
    req := dto.DatametaRegionGetPageReq{}
    s := service.DatametaRegion{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
   	if err != nil {
   		e.Logger.Error(err)
   		e.Error(500, err, err.Error())
   		return
   	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.DatametaRegion, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取DatametaRegion失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取DatametaRegion
// @Summary 获取DatametaRegion
// @Description 获取DatametaRegion
// @Tags DatametaRegion
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.DatametaRegion} "{"code": 200, "data": [...]}"
// @Router /api/v1/datameta-region/{id} [get]
// @Security Bearer
func (e DatametaRegion) Get(c *gin.Context) {
	req := dto.DatametaRegionGetReq{}
	s := service.DatametaRegion{}
    err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.DatametaRegion

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取DatametaRegion失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建DatametaRegion
// @Summary 创建DatametaRegion
// @Description 创建DatametaRegion
// @Tags DatametaRegion
// @Accept application/json
// @Product application/json
// @Param data body dto.DatametaRegionInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/datameta-region [post]
// @Security Bearer
func (e DatametaRegion) Insert(c *gin.Context) {
    req := dto.DatametaRegionInsertReq{}
    s := service.DatametaRegion{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建DatametaRegion失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改DatametaRegion
// @Summary 修改DatametaRegion
// @Description 修改DatametaRegion
// @Tags DatametaRegion
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.DatametaRegionUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/datameta-region/{id} [put]
// @Security Bearer
func (e DatametaRegion) Update(c *gin.Context) {
    req := dto.DatametaRegionUpdateReq{}
    s := service.DatametaRegion{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改DatametaRegion失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除DatametaRegion
// @Summary 删除DatametaRegion
// @Description 删除DatametaRegion
// @Tags DatametaRegion
// @Param data body dto.DatametaRegionDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/datameta-region [delete]
// @Security Bearer
func (e DatametaRegion) Delete(c *gin.Context) {
    s := service.DatametaRegion{}
    req := dto.DatametaRegionDeleteReq{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除DatametaRegion失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
