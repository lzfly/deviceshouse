package models

type DeviceBoilerAInfoPostForm struct {
	Device_sn            string `json:"device_sn"`
	Start                string  `json:"start"`
	Turn_Fire            string  `json:"turn_fire"`
	Stop                 string  `json:"stop"`
	Gas_Open             string  `json:"gas_open"`
    Gas_Feedback         string  `json:"gas_feedback"`
	Smoke_Loop           string  `json:"smoke_loop"`
	Steam_Pressure       string  `json:"steam_pressure"`
	Fan_Freq             string  `json:"fan_freq"`
	Freq_Feedback        string  `json:"freq_feedback"`
	Throttle_Open        string  `json:"throttle_open"`
	Throttle_Feedback    string  `json:"throttle_feedback"`
	Big_Fire             string  `json:"big_fire"`
	Small_Fire           string  `json:"small_fire"`
	Water_Pum            string  `json:"water_pum"`
}

type DeviceBoilerAInfoPostInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceBoilerAInfoInfo *DeviceBoilerAInfo `json:"DeviceBoilerAinfo"`
}

type DeviceBoilerAInfoGetOneInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceBoilerAInfoInfo *DeviceBoilerAInfo `json:"DeviceBoilerAinfo"`
}

type DeviceBoilerAInfoGetAllInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceBoilerAInfosInfo []DeviceBoilerAInfo `json:"DeviceBoilerAinfos"`
}

type DeviceBoilerAInfoPutForm struct {
	Device_sn            string `json:"device_sn"`
	Start                string  `json:"start"`
	Turn_Fire            string  `json:"turn_fire"`
	Stop                 string  `json:"stop"`
	Gas_Open             string  `json:"gas_open"`
    Gas_Feedback         string  `json:"gas_feedback"`
	Smoke_Loop           string  `json:"smoke_loop"`
	Steam_Pressure       string  `json:"steam_pressure"`
	Fan_Freq             string  `json:"fan_freq"`
	Freq_Feedback        string  `json:"freq_feedback"`
	Throttle_Open        string  `json:"throttle_open"`
	Throttle_Feedback    string  `json:"throttle_feedback"`
	Big_Fire             string  `json:"big_fire"`
	Small_Fire           string  `json:"small_fire"`
	Water_Pum            string  `json:"water_pum"`
}

type DeviceBoilerAInfoPutInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceBoilerAInfoInfo *DeviceBoilerAInfo `json:"DeviceBoilerAinfo"`
}

type DeviceBoilerAInfoDeleteInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
}
