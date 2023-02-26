package models

import (
	"go-admin/common/models"
)

type DatametaRegion struct {
	models.Model

	Name             string `json:"name" gorm:"type:varchar(255);comment:此条记录的名字"`
	CloudPlatform    string `json:"cloudPlatform" gorm:"type:varchar(255);comment:云平台 ali,huawei,tencent,aws...."`
	RegionId         string `json:"regionId" gorm:"type:varchar(255);comment:region id"`
	RegionName       string `json:"regionName" gorm:"type:varchar(255);comment:region name"`
	RegionEndpoint   string `json:"regionEndpoint" gorm:"type:varchar(255);comment:region endpoint"`
	AccessKeyId      string `json:"accessKeyId" gorm:"type:varchar(255);comment:ak_id"`
	AccessKeySecret  string `json:"accessKeySecret" gorm:"type:varchar(255);comment:ak_secret"`
	Status           string `json:"status" gorm:"type:tinyint;comment:region状态 0正常 1禁用"`
	Schedule         string `json:"schedule" gorm:"type:varchar(255);comment:实例拉取调度策略"`
	SchedulerContent string `json:"schedulerContent" gorm:"type:json;comment:schedule调度更新的实例类型"`
	models.ModelTime
	models.ControlBy
}

func (DatametaRegion) TableName() string {
	return "datameta_region"
}

func (e *DatametaRegion) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *DatametaRegion) GetId() interface{} {
	return e.Id
}
