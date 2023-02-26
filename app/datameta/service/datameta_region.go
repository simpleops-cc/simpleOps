package service

import (
	"errors"

    "github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/datameta/models"
	"go-admin/app/datameta/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type DatametaRegion struct {
	service.Service
}

// GetPage 获取DatametaRegion列表
func (e *DatametaRegion) GetPage(c *dto.DatametaRegionGetPageReq, p *actions.DataPermission, list *[]models.DatametaRegion, count *int64) error {
	var err error
	var data models.DatametaRegion

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("DatametaRegionService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取DatametaRegion对象
func (e *DatametaRegion) Get(d *dto.DatametaRegionGetReq, p *actions.DataPermission, model *models.DatametaRegion) error {
	var data models.DatametaRegion

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetDatametaRegion error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建DatametaRegion对象
func (e *DatametaRegion) Insert(c *dto.DatametaRegionInsertReq) error {
    var err error
    var data models.DatametaRegion
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("DatametaRegionService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改DatametaRegion对象
func (e *DatametaRegion) Update(c *dto.DatametaRegionUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.DatametaRegion{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("DatametaRegionService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除DatametaRegion
func (e *DatametaRegion) Remove(d *dto.DatametaRegionDeleteReq, p *actions.DataPermission) error {
	var data models.DatametaRegion

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveDatametaRegion error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
