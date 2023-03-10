package dto

import (
	"go-admin/app/datameta/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type DatametaRegionGetPageReq struct {
	dto.Pagination `search:"-"`
	DatametaRegionOrder
}

type DatametaRegionOrder struct {
	Id               string `form:"idOrder"  search:"type:order;column:id;table:datameta_region"`
	Name             string `form:"nameOrder"  search:"type:order;column:name;table:datameta_region"`
	CloudPlatform    string `form:"cloudPlatformOrder"  search:"type:order;column:cloud_platform;table:datameta_region"`
	RegionId         string `form:"regionIdOrder"  search:"type:order;column:region_id;table:datameta_region"`
	RegionName       string `form:"regionNameOrder"  search:"type:order;column:region_name;table:datameta_region"`
	RegionEndpoint   string `form:"regionEndpointOrder"  search:"type:order;column:region_endpoint;table:datameta_region"`
	AccessKeyId      string `form:"accessKeyIdOrder"  search:"type:order;column:access_key_id;table:datameta_region"`
	AccessKeySecret  string `form:"accessKeySecretOrder"  search:"type:order;column:access_key_secret;table:datameta_region"`
	Status           string `form:"statusOrder"  search:"type:order;column:status;table:datameta_region"`
	Schedule         string `form:"scheduleOrder"  search:"type:order;column:schedule;table:datameta_region"`
	SchedulerContent string `form:"schedulerContentOrder"  search:"type:order;column:scheduler_content;table:datameta_region"`
	CreatedAt        string `form:"createdAtOrder"  search:"type:order;column:created_at;table:datameta_region"`
	UpdatedAt        string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:datameta_region"`
	DeletedAt        string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:datameta_region"`
	CreateBy         string `form:"createByOrder"  search:"type:order;column:create_by;table:datameta_region"`
	UpdateBy         string `form:"updateByOrder"  search:"type:order;column:update_by;table:datameta_region"`
}

func (m *DatametaRegionGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type DatametaRegionInsertReq struct {
	Id               int    `json:"-" comment:"????????????"` // ????????????
	Name             string `json:"name" comment:"?????????????????????"`
	CloudPlatform    string `json:"cloudPlatform" comment:"????????? ali,huawei,tencent,aws...."`
	RegionId         string `json:"regionId" comment:"region id"`
	RegionName       string `json:"regionName" comment:"region name"`
	RegionEndpoint   string `json:"regionEndpoint" comment:"region endpoint"`
	AccessKeyId      string `json:"accessKeyId" comment:"ak_id"`
	AccessKeySecret  string `json:"accessKeySecret" comment:"ak_secret"`
	Status           string `json:"status" comment:"region?????? 0?????? 1??????"`
	Schedule         string `json:"schedule" comment:"????????????????????????"`
	SchedulerContent string `json:"schedulerContent" comment:"schedule???????????????????????????"`
	common.ControlBy
}

func (s *DatametaRegionInsertReq) Generate(model *models.DatametaRegion) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.CloudPlatform = s.CloudPlatform
	model.RegionId = s.RegionId
	model.RegionName = s.RegionName
	model.RegionEndpoint = s.RegionEndpoint
	model.AccessKeyId = s.AccessKeyId
	model.AccessKeySecret = s.AccessKeySecret
	model.Status = s.Status
	model.Schedule = s.Schedule
	model.SchedulerContent = s.SchedulerContent
	model.CreateBy = s.CreateBy // ?????????????????????????????????????????????
}

func (s *DatametaRegionInsertReq) GetId() interface{} {
	return s.Id
}

type DatametaRegionUpdateReq struct {
	Id               int    `uri:"id" comment:"????????????"` // ????????????
	Name             string `json:"name" comment:"?????????????????????"`
	CloudPlatform    string `json:"cloudPlatform" comment:"????????? ali,huawei,tencent,aws...."`
	RegionId         string `json:"regionId" comment:"region id"`
	RegionName       string `json:"regionName" comment:"region name"`
	RegionEndpoint   string `json:"regionEndpoint" comment:"region endpoint"`
	AccessKeyId      string `json:"accessKeyId" comment:"ak_id"`
	AccessKeySecret  string `json:"accessKeySecret" comment:"ak_secret"`
	Status           string `json:"status" comment:"region?????? 0?????? 1??????"`
	Schedule         string `json:"schedule" comment:"????????????????????????"`
	SchedulerContent string `json:"schedulerContent" comment:"schedule???????????????????????????"`
	common.ControlBy
}

func (s *DatametaRegionUpdateReq) Generate(model *models.DatametaRegion) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.CloudPlatform = s.CloudPlatform
	model.RegionId = s.RegionId
	model.RegionName = s.RegionName
	model.RegionEndpoint = s.RegionEndpoint
	model.AccessKeyId = s.AccessKeyId
	model.AccessKeySecret = s.AccessKeySecret
	model.Status = s.Status
	model.Schedule = s.Schedule
	model.SchedulerContent = s.SchedulerContent
	model.UpdateBy = s.UpdateBy // ?????????????????????????????????????????????
}

func (s *DatametaRegionUpdateReq) GetId() interface{} {
	return s.Id
}

// DatametaRegionGetReq ????????????????????????
type DatametaRegionGetReq struct {
	Id int `uri:"id"`
}

func (s *DatametaRegionGetReq) GetId() interface{} {
	return s.Id
}

// DatametaRegionDeleteReq ????????????????????????
type DatametaRegionDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *DatametaRegionDeleteReq) GetId() interface{} {
	return s.Ids
}
