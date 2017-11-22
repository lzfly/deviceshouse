package controllers

import (
	"boilerserver/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"time"
	"strings"
)

type DeviceBoilerBController struct {
	BaseController
}

func (this *DeviceBoilerBController) Post() {
	form := models.DeviceBoilerBPostForm{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseDeviceBoilerBPost:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseDeviceBoilerBPost:", &form)

	regDate := time.Now()
	
    DeviceBoilerB1 := models.NewDeviceBoilerB(&form, regDate)
	if _, err1 := DeviceBoilerB1.FindByDeviceSN(form.Device_sn); err1 == nil{
        this.RetError(errDupUser)
		return
	}
	
	DeviceBoilerB := models.NewDeviceBoilerB(&form, regDate)
	beego.Debug("NewDeviceBoilerB:", DeviceBoilerB)

	if code, err := DeviceBoilerB.Insert(); err != nil {
		beego.Error("InsertDeviceBoilerB:", err)
		if code == models.ErrDupRows {
			this.RetError(errDupUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}

	DeviceBoilerB.ClearPass()

	this.Data["json"] = &models.DeviceBoilerBPostInfo{Status:0, Code:0, DeviceBoilerBInfo: DeviceBoilerB}
	this.ServeJSON()
}

func (this *DeviceBoilerBController) GetOne() {
	idStr := this.Ctx.Input.Param(":id")
	/*id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		beego.Debug("ParseDeviceBoilerBId:", err)
		this.RetError(errInputData)
		return
	}*/
	
	DeviceBoilerB := models.DeviceBoilerB{}
	if code, err := DeviceBoilerB.FindByDeviceSN(idStr); err != nil {
		beego.Error("FindDeviceBoilerBById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("DeviceBoilerBInfo:", &DeviceBoilerB)

	DeviceBoilerB.ClearPass()

	this.Data["json"] = &models.DeviceBoilerBGetOneInfo{Status:0, Code:0, DeviceBoilerBInfo: &DeviceBoilerB}
	this.ServeJSON()
}

func (this *DeviceBoilerBController) GetAll() {
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

	DeviceBoilerBs, err := models.GetAllDeviceBoilerBs(queryVal, queryOp, order,
		limit, offset)
	if err != nil {
		beego.Error("GetAllDeviceBoilerB:", err)
		this.RetError(errDatabase)
		return
	}
	beego.Debug("GetAllDeviceBoilerB:", &DeviceBoilerBs)

	for i, _ := range DeviceBoilerBs {
		DeviceBoilerBs[i].ClearPass()
	}

	this.Data["json"] = &models.DeviceBoilerBGetAllInfo{Status:0, Code:0, DeviceBoilerBsInfo: DeviceBoilerBs}
	this.ServeJSON()
}

func (this *DeviceBoilerBController) Put() {
	idStr := this.Ctx.Input.Param(":id")
	/*id, err := strconv.ParseInt(idStr, 0, 32)
	if err != nil {
		beego.Debug("ParseDeviceBoilerBId:", err)
		this.RetError(errInputData)
		return
	}*/
	
	form := models.DeviceBoilerBPutForm{}

	err = json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseDeviceBoilerBPut:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseDeviceBoilerBPut:", &form)

	DeviceBoilerB := models.DeviceBoilerB{}
	if code, err := DeviceBoilerB.UpdateByDeviceSN(idStr, &form); err != nil {
		beego.Error("UpdateDeviceBoilerBById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUserChange)
		return
	}

	if code, err := DeviceBoilerB.FindByDeviceSN(idStr); err != nil {
		beego.Error("FindDeviceBoilerBById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("NewDeviceBoilerBInfo:", &DeviceBoilerB)

	DeviceBoilerB.ClearPass()

	this.Data["json"] = &models.DeviceBoilerBPutInfo{Status:0, Code:0, DeviceBoilerBInfo: &DeviceBoilerB}
	this.ServeJSON()
}

func (this *DeviceBoilerBController) Delete() {
	idStr := this.Ctx.Input.Param(":id")
	/*id, err := strconv.ParseInt(idStr, 0, 32)
	if err != nil {
		beego.Debug("ParseDeviceBoilerBId:", err)
		this.RetError(errInputData)
		return
	}*/

	
	DeviceBoilerB := models.DeviceBoilerB{}
	if code, err := DeviceBoilerB.DeleteByDeviceSN(idStr); err != nil {
		beego.Error("DeleteDeviceBoilerBById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUser)
		return
	}
	this.Data["json"] = &models.DeviceBoilerBDeleteInfo{Status:0, Code:0}
	this.ServeJSON()
}
