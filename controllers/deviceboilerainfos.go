package controllers

import (
	"boilerserver/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"time"
	"strings"
)

type DeviceBoilerAController struct {
	BaseController
}

func (this *DeviceBoilerAController) Post() {
	form := models.DeviceBoilerAPostForm{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseDeviceBoilerAPost:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseDeviceBoilerAPost:", &form)

	regDate := time.Now()
	
    DeviceBoilerA1 := models.NewDeviceBoilerA(&form, regDate)
	if _, err1 := DeviceBoilerA1.FindByDeviceSN(form.Device_sn); err1 == nil{
        this.RetError(errDupUser)
		return
	}
	
	DeviceBoilerA := models.NewDeviceBoilerA(&form, regDate)
	beego.Debug("NewDeviceBoilerA:", DeviceBoilerA)

	if code, err := DeviceBoilerA.Insert(); err != nil {
		beego.Error("InsertDeviceBoilerA:", err)
		if code == models.ErrDupRows {
			this.RetError(errDupUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}

	DeviceBoilerA.ClearPass()

	this.Data["json"] = &models.DeviceBoilerAPostInfo{Status:0, Code:0, DeviceBoilerAInfo: DeviceBoilerA}
	this.ServeJSON()
}

func (this *DeviceBoilerAController) GetOne() {
	idStr := this.Ctx.Input.Param(":id")
	/*id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		beego.Debug("ParseDeviceBoilerAId:", err)
		this.RetError(errInputData)
		return
	}*/
	
	DeviceBoilerA := models.DeviceBoilerA{}
	if code, err := DeviceBoilerA.FindByDeviceSN(idStr); err != nil {
		beego.Error("FindDeviceBoilerAById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("DeviceBoilerAInfo:", &DeviceBoilerA)

	DeviceBoilerA.ClearPass()

	this.Data["json"] = &models.DeviceBoilerAGetOneInfo{Status:0, Code:0, DeviceBoilerAInfo: &DeviceBoilerA}
	this.ServeJSON()
}

func (this *DeviceBoilerAController) GetAll() {
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

	DeviceBoilerAs, err := models.GetAllDeviceBoilerAs(queryVal, queryOp, order,
		limit, offset)
	if err != nil {
		beego.Error("GetAllDeviceBoilerA:", err)
		this.RetError(errDatabase)
		return
	}
	beego.Debug("GetAllDeviceBoilerA:", &DeviceBoilerAs)

	for i, _ := range DeviceBoilerAs {
		DeviceBoilerAs[i].ClearPass()
	}

	this.Data["json"] = &models.DeviceBoilerAGetAllInfo{Status:0, Code:0, DeviceBoilerAsInfo: DeviceBoilerAs}
	this.ServeJSON()
}

func (this *DeviceBoilerAController) Put() {
	idStr := this.Ctx.Input.Param(":id")
	/*id, err := strconv.ParseInt(idStr, 0, 32)
	if err != nil {
		beego.Debug("ParseDeviceBoilerAId:", err)
		this.RetError(errInputData)
		return
	}*/
	
	form := models.DeviceBoilerAPutForm{}

	err = json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseDeviceBoilerAPut:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseDeviceBoilerAPut:", &form)

	DeviceBoilerA := models.DeviceBoilerA{}
	if code, err := DeviceBoilerA.UpdateByDeviceSN(idStr, &form); err != nil {
		beego.Error("UpdateDeviceBoilerAById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUserChange)
		return
	}

	if code, err := DeviceBoilerA.FindByDeviceSN(idStr); err != nil {
		beego.Error("FindDeviceBoilerAById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("NewDeviceBoilerAInfo:", &DeviceBoilerA)

	DeviceBoilerA.ClearPass()

	this.Data["json"] = &models.DeviceBoilerAPutInfo{Status:0, Code:0, DeviceBoilerAInfo: &DeviceBoilerA}
	this.ServeJSON()
}

func (this *DeviceBoilerAController) Delete() {
	idStr := this.Ctx.Input.Param(":id")
	/*id, err := strconv.ParseInt(idStr, 0, 32)
	if err != nil {
		beego.Debug("ParseDeviceBoilerAId:", err)
		this.RetError(errInputData)
		return
	}*/

	
	DeviceBoilerA := models.DeviceBoilerA{}
	if code, err := DeviceBoilerA.DeleteByDeviceSN(idStr); err != nil {
		beego.Error("DeleteDeviceBoilerAById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUser)
		return
	}
	this.Data["json"] = &models.DeviceBoilerADeleteInfo{Status:0, Code:0}
	this.ServeJSON()
}
