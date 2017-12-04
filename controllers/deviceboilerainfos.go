package controllers

import (
	"deviceshouse/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"time"
)

type DeviceBoilerAInfoController struct {
	BaseController
}

func (this *DeviceBoilerAInfoController) Post() {
	form := models.DeviceBoilerAInfoPostForm{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseDeviceBoilerAInfoPost:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseDeviceBoilerAInfoPost:", &form)

	regDate := time.Now()
	
    DeviceBoilerAInfo1 := models.NewDeviceBoilerAInfo(&form, regDate)
	if _, err1 := DeviceBoilerAInfo1.FindByDeviceSN(form.Device_sn); err1 == nil{
        this.RetError(errDupUser)
		return
	}
	
	DeviceBoilerAInfo := models.NewDeviceBoilerAInfo(&form, regDate)
	beego.Debug("NewDeviceBoilerAInfo:", DeviceBoilerAInfo)

	if code, err := DeviceBoilerAInfo.Insert(); err != nil {
		beego.Error("InsertDeviceBoilerAInfo:", err)
		if code == models.ErrDupRows {
			this.RetError(errDupUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}

	DeviceBoilerAInfo.ClearPass()

	this.Data["json"] = &models.DeviceBoilerAInfoPostInfo{Status:0, Code:0, DeviceBoilerAInfoInfo: DeviceBoilerAInfo}
	this.ServeJSON()
}

func (this *DeviceBoilerAInfoController) GetOne() {
	idStr := this.Ctx.Input.Param(":id")
	/*id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		beego.Debug("ParseDeviceBoilerAInfoId:", err)
		this.RetError(errInputData)
		return
	}*/
	
	DeviceBoilerAInfo := models.DeviceBoilerAInfo{}
	if code, err := DeviceBoilerAInfo.FindByDeviceSN(idStr); err != nil {
		beego.Error("FindDeviceBoilerAInfoById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("DeviceBoilerAInfoInfo:", &DeviceBoilerAInfo)

	DeviceBoilerAInfo.ClearPass()

	this.Data["json"] = &models.DeviceBoilerAInfoGetOneInfo{Status:0, Code:0, DeviceBoilerAInfoInfo: &DeviceBoilerAInfo}
	this.ServeJSON()
}

func (this *DeviceBoilerAInfoController) GetAll() {
	queryVal, queryOp, err := this.ParseQueryParm()
	if err != nil {
		beego.Debug("ParseQuery:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("QueryVal:", queryVal)
	beego.Debug("QueryOp:", queryOp)

	order, err := this.ParseOrderParm()
	if err != nil {
		beego.Debug("ParseOrder:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("Order:", order)

	limit, err := this.ParseLimitParm()
	/*
		if err != nil {
			beego.Debug("ParseLimit:", err)
			this.RetError(errInputData)
			return
		}
	*/
	beego.Debug("Limit:", limit)

	offset, err := this.ParseOffsetParm()
	/*
		if err != nil {
			beego.Debug("ParseOffset:", err)
			this.RetError(errInputData)
			return
		}
	*/
	beego.Debug("Offset:", offset)

	DeviceBoilerAInfos, err := models.GetAllDeviceBoilerAInfos(queryVal, queryOp, order,
		limit, offset)
	if err != nil {
		beego.Error("GetAllDeviceBoilerAInfo:", err)
		this.RetError(errDatabase)
		return
	}
	beego.Debug("GetAllDeviceBoilerAInfo:", &DeviceBoilerAInfos)

	for i, _ := range DeviceBoilerAInfos {
		DeviceBoilerAInfos[i].ClearPass()
	}

	this.Data["json"] = &models.DeviceBoilerAInfoGetAllInfo{Status:0, Code:0, DeviceBoilerAInfosInfo: DeviceBoilerAInfos}
	this.ServeJSON()
}

func (this *DeviceBoilerAInfoController) Put() {
	idStr := this.Ctx.Input.Param(":id")
	/*id, err := strconv.ParseInt(idStr, 0, 32)
	if err != nil {
		beego.Debug("ParseDeviceBoilerAInfoId:", err)
		this.RetError(errInputData)
		return
	}*/
	
	form := models.DeviceBoilerAInfoPutForm{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseDeviceBoilerAInfoPut:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseDeviceBoilerAInfoPut:", &form)

	DeviceBoilerAInfo := models.DeviceBoilerAInfo{}
	if code, err := DeviceBoilerAInfo.UpdateByDeviceSN(idStr, &form); err != nil {
		beego.Error("UpdateDeviceBoilerAInfoById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUserChange)
		return
	}

	if code, err := DeviceBoilerAInfo.FindByDeviceSN(idStr); err != nil {
		beego.Error("FindDeviceBoilerAInfoById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("NewDeviceBoilerAInfoInfo:", &DeviceBoilerAInfo)

	DeviceBoilerAInfo.ClearPass()

	this.Data["json"] = &models.DeviceBoilerAInfoPutInfo{Status:0, Code:0, DeviceBoilerAInfoInfo: &DeviceBoilerAInfo}
	this.ServeJSON()
}

func (this *DeviceBoilerAInfoController) Delete() {
	idStr := this.Ctx.Input.Param(":id")
	/*id, err := strconv.ParseInt(idStr, 0, 32)
	if err != nil {
		beego.Debug("ParseDeviceBoilerAInfoId:", err)
		this.RetError(errInputData)
		return
	}*/

	
	DeviceBoilerAInfo := models.DeviceBoilerAInfo{}
	if code, err := DeviceBoilerAInfo.DeleteByDeviceSN(idStr); err != nil {
		beego.Error("DeleteDeviceBoilerAInfoById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUser)
		return
	}
	this.Data["json"] = &models.DeviceBoilerAInfoDeleteInfo{Status:0, Code:0}
	this.ServeJSON()
}
