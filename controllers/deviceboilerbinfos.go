package controllers

import (
	"boilerserver/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"time"
)

type DeviceBoilerBInfoController struct {
	BaseController
}

func (this *DeviceBoilerBInfoController) Post() {
	form := models.DeviceBoilerBInfoPostForm{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseDeviceBoilerBInfoPost:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseDeviceBoilerBInfoPost:", &form)

	regDate := time.Now()
	
    DeviceBoilerBInfo1 := models.NewDeviceBoilerBInfo(&form, regDate)
	if _, err1 := DeviceBoilerBInfo1.FindByDeviceSN(form.Device_sn); err1 == nil{
        this.RetError(errDupUser)
		return
	}
	
	DeviceBoilerBInfo := models.NewDeviceBoilerBInfo(&form, regDate)
	beego.Debug("NewDeviceBoilerBInfo:", DeviceBoilerBInfo)

	if code, err := DeviceBoilerBInfo.Insert(); err != nil {
		beego.Error("InsertDeviceBoilerBInfo:", err)
		if code == models.ErrDupRows {
			this.RetError(errDupUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}

	DeviceBoilerBInfo.ClearPass()

	this.Data["json"] = &models.DeviceBoilerBInfoPostInfo{Status:0, Code:0, DeviceBoilerBInfoInfo: DeviceBoilerBInfo}
	this.ServeJSON()
}

func (this *DeviceBoilerBInfoController) GetOne() {
	idStr := this.Ctx.Input.Param(":id")
	/*id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		beego.Debug("ParseDeviceBoilerBInfoId:", err)
		this.RetError(errInputData)
		return
	}*/
	
	DeviceBoilerBInfo := models.DeviceBoilerBInfo{}
	if code, err := DeviceBoilerBInfo.FindByDeviceSN(idStr); err != nil {
		beego.Error("FindDeviceBoilerBInfoById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("DeviceBoilerBInfoInfo:", &DeviceBoilerBInfo)

	DeviceBoilerBInfo.ClearPass()

	this.Data["json"] = &models.DeviceBoilerBInfoGetOneInfo{Status:0, Code:0, DeviceBoilerBInfoInfo: &DeviceBoilerBInfo}
	this.ServeJSON()
}

func (this *DeviceBoilerBInfoController) GetAll() {
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

	DeviceBoilerBInfos, err := models.GetAllDeviceBoilerBInfos(queryVal, queryOp, order,
		limit, offset)
	if err != nil {
		beego.Error("GetAllDeviceBoilerBInfo:", err)
		this.RetError(errDatabase)
		return
	}
	beego.Debug("GetAllDeviceBoilerBInfo:", &DeviceBoilerBInfos)

	for i, _ := range DeviceBoilerBInfos {
		DeviceBoilerBInfos[i].ClearPass()
	}

	this.Data["json"] = &models.DeviceBoilerBInfoGetAllInfo{Status:0, Code:0, DeviceBoilerBInfosInfo: DeviceBoilerBInfos}
	this.ServeJSON()
}

func (this *DeviceBoilerBInfoController) Put() {
	idStr := this.Ctx.Input.Param(":id")
	/*id, err := strconv.ParseInt(idStr, 0, 32)
	if err != nil {
		beego.Debug("ParseDeviceBoilerBInfoId:", err)
		this.RetError(errInputData)
		return
	}*/
	
	form := models.DeviceBoilerBInfoPutForm{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseDeviceBoilerBInfoPut:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseDeviceBoilerBInfoPut:", &form)

	DeviceBoilerBInfo := models.DeviceBoilerBInfo{}
	if code, err := DeviceBoilerBInfo.UpdateByDeviceSN(idStr, &form); err != nil {
		beego.Error("UpdateDeviceBoilerBInfoById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUserChange)
		return
	}

	if code, err := DeviceBoilerBInfo.FindByDeviceSN(idStr); err != nil {
		beego.Error("FindDeviceBoilerBInfoById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("NewDeviceBoilerBInfoInfo:", &DeviceBoilerBInfo)

	DeviceBoilerBInfo.ClearPass()

	this.Data["json"] = &models.DeviceBoilerBInfoPutInfo{Status:0, Code:0, DeviceBoilerBInfoInfo: &DeviceBoilerBInfo}
	this.ServeJSON()
}

func (this *DeviceBoilerBInfoController) Delete() {
	idStr := this.Ctx.Input.Param(":id")
	/*id, err := strconv.ParseInt(idStr, 0, 32)
	if err != nil {
		beego.Debug("ParseDeviceBoilerBInfoId:", err)
		this.RetError(errInputData)
		return
	}*/

	
	DeviceBoilerBInfo := models.DeviceBoilerBInfo{}
	if code, err := DeviceBoilerBInfo.DeleteByDeviceSN(idStr); err != nil {
		beego.Error("DeleteDeviceBoilerBInfoById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUser)
		return
	}
	this.Data["json"] = &models.DeviceBoilerBInfoDeleteInfo{Status:0, Code:0}
	this.ServeJSON()
}
