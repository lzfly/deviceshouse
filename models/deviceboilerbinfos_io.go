package models

type DeviceBoilerBInfoPostForm struct {
	Device_sn            string `json:"device_sn"`
	Start_Temp           string  `json:"start_temp"`
	Target_Temp          string  `json:"target_temp"`
	Stop_Temp            string  `json:"stop_temp"`
	Out_Water_Temp       string  `json:"out_water_temp"`
	Back_Water_Temp      string  `json:"back_water_temp"`
    Smoke_Temp           string  `json:"smoke_temp"`
	Boiler_Load          string  `json:"boiler_load"`
	Gas                  string  `json:"gas"`
	Throttle             string  `json:"throttle"`
	Smoke                string  `json:"smoke"`
	Freq                 string  `json:"freq"`
	Run_State            string  `json:"run_state"`
}

type DeviceBoilerBInfoPostInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceBoilerBInfoInfo *DeviceBoilerBInfo `json:"DeviceBoilerBinfo"`
}

type DeviceBoilerBInfoGetOneInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceBoilerBInfoInfo *DeviceBoilerBInfo `json:"DeviceBoilerBinfo"`
}

type DeviceBoilerBInfoGetAllInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceBoilerBInfosInfo []DeviceBoilerBInfo `json:"DeviceBoilerBinfos"`
}

type DeviceBoilerBInfoPutForm struct {
	Device_sn            string `json:"device_sn"`
	Start_Temp           string  `json:"start_temp"`
	Target_Temp          string  `json:"target_temp"`
	Stop_Temp            string  `json:"stop_temp"`
	Out_Water_Temp       string  `json:"out_water_temp"`
	Back_Water_Temp      string  `json:"back_water_temp"`
    Smoke_Temp           string  `json:"smoke_temp"`
	Boiler_Load          string  `json:"boiler_load"`
	Gas                  string  `json:"gas"`
	Throttle             string  `json:"throttle"`
	Smoke                string  `json:"smoke"`
	Freq                 string  `json:"freq"`
	Run_State            string  `json:"run_state"`
}

type DeviceBoilerBInfoPutInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceBoilerBInfoInfo *DeviceBoilerBInfo `json:"DeviceBoilerBinfo"`
}

type DeviceBoilerBInfoDeleteInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
}
