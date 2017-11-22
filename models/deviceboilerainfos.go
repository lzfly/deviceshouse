package models

import (
	"boilerserver/models/mymysql"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-sql-driver/mysql"
	"time"
)

type DeviceBoilerAInfo struct {
	Id            int64     `json:"id"`
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
	Water_Pump            string  `json:"water_pump"`
}

func NewDeviceBoilerAInfo(f *DeviceBoilerAInfoPostForm, t time.Time) *DeviceBoilerAInfo {
	DeviceBoilerAinfo := DeviceBoilerAInfo{
		Device_sn:     f.Device_sn,
		Start:     f.Start,
		Turn_Fire:     f.Turn_Fire,
		Stop:     f.Stop,
		Gas_Open:    f.Gas_Open,
		Gas_Feedback:    f.Gas_Feedback,
		Smoke_Loop:     f.Smoke_Loop,
		Steam_Pressure:     f.Steam_Pressure,
		Fan_Freq:     f.Fan_Freq,
		Freq_Feedback:    f.Freq_Feedback,
		Throttle_Open:    f.Throttle_Open,
		Throttle_Feedback:     f.Throttle_Feedback,
		Big_Fire:     f.Big_Fire,
		Small_Fire:     f.Small_Fire,
		Water_Pump:    f.Water_Pump,}

	return &DeviceBoilerAinfo
}

func (r *DeviceBoilerAInfo) Insert() (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("INSERT INTO dev_boiler_a_info(device_sn, start, turn_fire, stop, gas_open, gas_feedback, smoke_loop, steam_pressure, fan_freq, freq_feedback, throttle_open, throttle_feedback, big_fire, small_fire, water_pump) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	//if result, err := st.Exec(
	if _, err := st.Exec(r.Device_sn, r.Start, r.Turn_Fire, r.Stop, r.Gas_Open, r.Gas_Feedback, r.Smoke_Loop, r.Steam_Pressure, r.Fan_Freq, r.Freq_Feedback, r.Throttle_Open, r.Throttle_Feedback, r.Big_Fire, r.Small_Fire, r.Water_Pump); err != nil {
		if e, ok := err.(*mysql.MySQLError); ok {
			//Duplicate key
			if e.Number == 1062 {
				return ErrDupRows, err
			} else {
				return ErrDatabase, err
			}
		} else {
			return ErrDatabase, err
		}
	} else {
		//r.Id, _ = result.LastInsertId()
	}

	return 0, nil
}

func (r *DeviceBoilerAInfo) FindById(id int64) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("SELECT id, device_sn, start, turn_fire, stop, gas_open, gas_feedback, smoke_loop, steam_pressure, fan_freq, freq_feedback, throttle_open, throttle_feedback, big_fire, small_fire, water_pump FROM dev_boiler_a_info WHERE id = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	row := st.QueryRow(id)
	
	var tmpId           sql.NullInt64 
	var tmpDevice_sn    sql.NullString 
	var tmpStart        sql.NullString
	var tmpTurn_Fire    sql.NullString
    var tmpStop         sql.NullString  
    var tmpGas_Open     sql.NullString 
	var tmpGas_Feedback     sql.NullString 
	var tmpSmoke_Loop   sql.NullString
	var tmpSteam_Pressure    sql.NullString
    var tmpFan_Freq     sql.NullString  
    var tmpFreq_Feedback   sql.NullString 
	var tmpThrottle_Open  sql.NullString
	var tmpThrottle_Feedback    sql.NullString
    var tmpBig_Fire     sql.NullString  
    var tmpSmall_Fire   sql.NullString 
	var tmpWater_Pump    sql.NullString 


	if err := row.Scan(&tmpId, &tmpDevice_sn, &tmpStart, &tmpTurn_Fire, &tmpStop, &tmpGas_Open, &tmpGas_Feedback,  &tmpSmoke_Loop, &tmpSteam_Pressure, &tmpFan_Freq, &tmpFreq_Feedback, &tmpThrottle_Open, &tmpThrottle_Feedback,  &tmpBig_Fire, &tmpSmall_Fire, &tmpWater_Pump); err != nil {
		if err == sql.ErrNoRows {
			return ErrNotFound, err
		} else {
			return ErrDatabase, err
		}
	}

	if tmpId.Valid {
		r.Id = tmpId.Int64
	}
		if tmpDevice_sn.Valid {
			r.Device_sn = tmpDevice_sn.String
		}
		if tmpStart.Valid {
			r.Start = tmpStart.String
		}
		if tmpTurn_Fire.Valid {
			r.Turn_Fire = tmpTurn_Fire.String
		}
		if tmpStop.Valid {
			r.Stop = tmpStop.String
		}
		if tmpGas_Open.Valid {
			r.Gas_Open = tmpGas_Open.String
		}
		if tmpGas_Feedback.Valid {
			r.Gas_Feedback = tmpGas_Feedback.String
		}
		if tmpSmoke_Loop.Valid {
			r.Smoke_Loop = tmpSmoke_Loop.String
		}
		if tmpSteam_Pressure.Valid {
			r.Steam_Pressure = tmpSteam_Pressure.String
		}
		if tmpFan_Freq.Valid {
			r.Fan_Freq = tmpFan_Freq.String
		}
		if tmpFreq_Feedback.Valid {
			r.Freq_Feedback = tmpFreq_Feedback.String
		}
		if tmpThrottle_Open.Valid {
			r.Throttle_Open = tmpThrottle_Open.String
		}
		if tmpThrottle_Feedback.Valid {
			r.Throttle_Feedback = tmpThrottle_Feedback.String
		}
		if tmpBig_Fire.Valid {
			r.Big_Fire = tmpBig_Fire.String
		}
		if tmpSmall_Fire.Valid {
			r.Small_Fire = tmpSmall_Fire.String
		}
		if tmpWater_Pump.Valid {
			r.Water_Pump = tmpWater_Pump.String
		}

	return 0, nil
}

func (r *DeviceBoilerAInfo) FindByDeviceSN(device_sn string) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("SELECT id, device_sn, start, turn_fire, stop, gas_open, gas_feedback, smoke_loop, steam_pressure, fan_freq, freq_feedback, throttle_open, throttle_feedback, big_fire, small_fire, water_pump FROM dev_boiler_a_info WHERE device_sn = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	row := st.QueryRow(device_sn)

	var tmpId           sql.NullInt64 
	var tmpDevice_sn    sql.NullString 
	var tmpStart        sql.NullString
	var tmpTurn_Fire    sql.NullString
    var tmpStop         sql.NullString  
    var tmpGas_Open     sql.NullString 
	var tmpGas_Feedback     sql.NullString 
	var tmpSmoke_Loop   sql.NullString
	var tmpSteam_Pressure    sql.NullString
    var tmpFan_Freq     sql.NullString  
    var tmpFreq_Feedback   sql.NullString 
	var tmpThrottle_Open  sql.NullString
	var tmpThrottle_Feedback    sql.NullString
    var tmpBig_Fire     sql.NullString  
    var tmpSmall_Fire   sql.NullString 
	var tmpWater_Pump    sql.NullString 


	if err := row.Scan(&tmpId, &tmpDevice_sn, &tmpStart, &tmpTurn_Fire, &tmpStop, &tmpGas_Open, &tmpGas_Feedback,  &tmpSmoke_Loop, &tmpSteam_Pressure, &tmpFan_Freq, &tmpFreq_Feedback, &tmpThrottle_Open, &tmpThrottle_Feedback,  &tmpBig_Fire, &tmpSmall_Fire, &tmpWater_Pump); err != nil {
		if err == sql.ErrNoRows {
			return ErrNotFound, err
		} else {
			return ErrDatabase, err
		}
	}

	if tmpId.Valid {
		r.Id = tmpId.Int64
	}
		if tmpDevice_sn.Valid {
			r.Device_sn = tmpDevice_sn.String
		}
		if tmpStart.Valid {
			r.Start = tmpStart.String
		}
		if tmpTurn_Fire.Valid {
			r.Turn_Fire = tmpTurn_Fire.String
		}
		if tmpStop.Valid {
			r.Stop = tmpStop.String
		}
		if tmpGas_Open.Valid {
			r.Gas_Open = tmpGas_Open.String
		}
		if tmpGas_Feedback.Valid {
			r.Gas_Feedback = tmpGas_Feedback.String
		}
		if tmpSmoke_Loop.Valid {
			r.Smoke_Loop = tmpSmoke_Loop.String
		}
		if tmpSteam_Pressure.Valid {
			r.Steam_Pressure = tmpSteam_Pressure.String
		}
		if tmpFan_Freq.Valid {
			r.Fan_Freq = tmpFan_Freq.String
		}
		if tmpFreq_Feedback.Valid {
			r.Freq_Feedback = tmpFreq_Feedback.String
		}
		if tmpThrottle_Open.Valid {
			r.Throttle_Open = tmpThrottle_Open.String
		}
		if tmpThrottle_Feedback.Valid {
			r.Throttle_Feedback = tmpThrottle_Feedback.String
		}
		if tmpBig_Fire.Valid {
			r.Big_Fire = tmpBig_Fire.String
		}
		if tmpSmall_Fire.Valid {
			r.Small_Fire = tmpSmall_Fire.String
		}
		if tmpWater_Pump.Valid {
			r.Water_Pump = tmpWater_Pump.String
		}

	return 0, nil
}

func (r *DeviceBoilerAInfo) ClearPass() {
	
}

func GetAllDeviceBoilerAInfos(queryVal map[string]string, queryOp map[string]string,
	order map[string]string, limit int64,
	offset int64) (records []DeviceBoilerAInfo, err error) {
	sqlStr := "SELECT id, device_sn, start, turn_fire, stop, gas_open, gas_feedback, smoke_loop, steam_pressure, fan_freq, freq_feedback, throttle_open, throttle_feedback, big_fire, small_fire, water_pump FROM dev_boiler_a_info"
	if len(queryVal) > 0 {
		sqlStr += " WHERE "
		first := true
		for k, v := range queryVal {
			if !first {
				sqlStr += " AND "
			} else {
				first = false
			}

			sqlStr += k
			sqlStr += " "
			sqlStr += queryOp[k]
			sqlStr += " '"
			sqlStr += v
			sqlStr += "'"
		}
	}
	if len(order) > 0 {
		sqlStr += " ORDER BY "
		first := true
		for k, v := range order {
			if !first {
				sqlStr += ", "
			} else {
				first = false
			}

			sqlStr += k
			sqlStr += " "
			sqlStr += v
		}
	}
	sqlStr += " LIMIT " + fmt.Sprintf("%d", limit)
	if offset > 0 {
		sqlStr += " OFFSET " + fmt.Sprintf("%d", offset)
	}
	beego.Debug("sqlStr:", sqlStr)

	db := mymysql.Conn()

	st, err := db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	defer st.Close()

	rows, err := st.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	records = make([]DeviceBoilerAInfo, 0, limit)
	for rows.Next() {

		var tmpId           sql.NullInt64 
		var tmpDevice_sn    sql.NullString 
		var tmpStart        sql.NullString
		var tmpTurn_Fire    sql.NullString
		var tmpStop         sql.NullString  
		var tmpGas_Open     sql.NullString 
		var tmpGas_Feedback     sql.NullString 
		var tmpSmoke_Loop   sql.NullString
		var tmpSteam_Pressure    sql.NullString
		var tmpFan_Freq     sql.NullString  
		var tmpFreq_Feedback   sql.NullString 
		var tmpThrottle_Open  sql.NullString
		var tmpThrottle_Feedback    sql.NullString
		var tmpBig_Fire     sql.NullString  
		var tmpSmall_Fire   sql.NullString 
		var tmpWater_Pump    sql.NullString
		
		if err := rows.Scan(&tmpId, &tmpDevice_sn, &tmpStart, &tmpTurn_Fire, &tmpStop, &tmpGas_Open, &tmpGas_Feedback,  &tmpSmoke_Loop, &tmpSteam_Pressure, &tmpFan_Freq, &tmpFreq_Feedback, &tmpThrottle_Open, &tmpThrottle_Feedback,  &tmpBig_Fire, &tmpSmall_Fire, &tmpWater_Pump); err != nil {
			return nil, err
		}

        r := DeviceBoilerAInfo{}
		if tmpId.Valid {
			r.Id = tmpId.Int64
		}
		if tmpDevice_sn.Valid {
			r.Device_sn = tmpDevice_sn.String
		}
		if tmpStart.Valid {
			r.Start = tmpStart.String
		}
		if tmpTurn_Fire.Valid {
			r.Turn_Fire = tmpTurn_Fire.String
		}
		if tmpStop.Valid {
			r.Stop = tmpStop.String
		}
		if tmpGas_Open.Valid {
			r.Gas_Open = tmpGas_Open.String
		}
		if tmpGas_Feedback.Valid {
			r.Gas_Feedback = tmpGas_Feedback.String
		}
		if tmpSmoke_Loop.Valid {
			r.Smoke_Loop = tmpSmoke_Loop.String
		}
		if tmpSteam_Pressure.Valid {
			r.Steam_Pressure = tmpSteam_Pressure.String
		}
		if tmpFan_Freq.Valid {
			r.Fan_Freq = tmpFan_Freq.String
		}
		if tmpFreq_Feedback.Valid {
			r.Freq_Feedback = tmpFreq_Feedback.String
		}
		if tmpThrottle_Open.Valid {
			r.Throttle_Open = tmpThrottle_Open.String
		}
		if tmpThrottle_Feedback.Valid {
			r.Throttle_Feedback = tmpThrottle_Feedback.String
		}
		if tmpBig_Fire.Valid {
			r.Big_Fire = tmpBig_Fire.String
		}
		if tmpSmall_Fire.Valid {
			r.Small_Fire = tmpSmall_Fire.String
		}
		if tmpWater_Pump.Valid {
			r.Water_Pump = tmpWater_Pump.String
		}
		
		records = append(records, r)
		}
		if err := rows.Err(); err != nil {
			return nil, err
		}

	return records, nil
}

func (r *DeviceBoilerAInfo) UpdateById(id int64, f *DeviceBoilerAInfoPutForm) (code int, err error) {
	db := mymysql.Conn()
	
	if len(f.Device_sn) > 0 {
		st, err1 := db.Prepare("UPDATE dev_boiler_a_info SET device_sn = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Device_sn, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	if len(f.Start) > 0 {
		st, err1 := db.Prepare("UPDATE dev_boiler_a_info SET start = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Start, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Turn_Fire) > 0 {
		st, err1 := db.Prepare("UPDATE dev_boiler_a_info SET turn_fire = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Turn_Fire, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Stop) > 0 {
		st, err1 := db.Prepare("UPDATE dev_boiler_a_info SET stop = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Stop, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	if len(f.Gas_Open) > 0 {
		st, err1 := db.Prepare("UPDATE dev_boiler_a_info SET gas_open = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Gas_Open, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	if len(f.Gas_Feedback) > 0 {
		st, err1 := db.Prepare("UPDATE dev_boiler_a_info SET gas_feedback = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Gas_Feedback, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	
	if len(f.Smoke_Loop) > 0 {
		st, err1 := db.Prepare("UPDATE dev_boiler_a_info SET smoke_loop = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Smoke_Loop, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	if len(f.Steam_Pressure) > 0 {
		st, err1 := db.Prepare("UPDATE dev_boiler_a_info SET steam_pressure = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Steam_Pressure, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Fan_Freq) > 0 {
		st, err1 := db.Prepare("UPDATE dev_boiler_a_info SET fan_freq = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Fan_Freq, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Freq_Feedback) > 0 {
		st, err1 := db.Prepare("UPDATE dev_boiler_a_info SET freq_feedback = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Freq_Feedback, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	if len(f.Throttle_Open) > 0 {
		st, err1 := db.Prepare("UPDATE dev_boiler_a_info SET throttle_open = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Throttle_Open, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	if len(f.Throttle_Feedback) > 0 {
		st, err1 := db.Prepare("UPDATE dev_boiler_a_info SET throttle_feedback = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Throttle_Feedback, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.Big_Fire) > 0 {
		st, err1 := db.Prepare("UPDATE dev_boiler_a_info SET big_fire = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Big_Fire, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	if len(f.Small_Fire) > 0 {
		st, err1 := db.Prepare("UPDATE dev_boiler_a_info SET small_fire = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Small_Fire, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	if len(f.Water_Pump) > 0 {
		st, err1 := db.Prepare("UPDATE dev_boiler_a_info SET water_pump = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Water_Pump, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	return 0, nil

}

func (r *DeviceBoilerAInfo) UpdateByDeviceSN(device_sn string, f *DeviceBoilerAInfoPutForm) (code int, err error) {
	db := mymysql.Conn()



	if len(f.Start) > 0 {
		st, err1 := db.Prepare("UPDATE dev_boiler_a_info SET start = ? WHERE device_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Start, device_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Turn_Fire) > 0 {
		st, err1 := db.Prepare("UPDATE dev_boiler_a_info SET turn_fire = ? WHERE device_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Turn_Fire, device_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Stop) > 0 {
		st, err1 := db.Prepare("UPDATE dev_boiler_a_info SET stop = ? WHERE device_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Stop, device_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	if len(f.Gas_Open) > 0 {
		st, err1 := db.Prepare("UPDATE dev_boiler_a_info SET gas_open = ? WHERE device_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Gas_Open, device_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	if len(f.Gas_Feedback) > 0 {
		st, err1 := db.Prepare("UPDATE dev_boiler_a_info SET gas_feedback = ? WHERE device_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Gas_Feedback, device_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	
	if len(f.Smoke_Loop) > 0 {
		st, err1 := db.Prepare("UPDATE dev_boiler_a_info SET smoke_loop = ? WHERE device_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Smoke_Loop, device_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	if len(f.Steam_Pressure) > 0 {
		st, err1 := db.Prepare("UPDATE dev_boiler_a_info SET steam_pressure = ? WHERE device_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Steam_Pressure, device_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Fan_Freq) > 0 {
		st, err1 := db.Prepare("UPDATE dev_boiler_a_info SET fan_freq = ? WHERE device_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Fan_Freq, device_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Freq_Feedback) > 0 {
		st, err1 := db.Prepare("UPDATE dev_boiler_a_info SET freq_feedback = ? WHERE device_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Freq_Feedback, device_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	if len(f.Throttle_Open) > 0 {
		st, err1 := db.Prepare("UPDATE dev_boiler_a_info SET throttle_open = ? WHERE device_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Throttle_Open, device_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	if len(f.Throttle_Feedback) > 0 {
		st, err1 := db.Prepare("UPDATE dev_boiler_a_info SET throttle_feedback = ? WHERE device_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Throttle_Feedback, device_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	if len(f.Big_Fire) > 0 {
		st, err1 := db.Prepare("UPDATE dev_boiler_a_info SET big_fire = ? WHERE device_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Big_Fire, device_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	if len(f.Small_Fire) > 0 {
		st, err1 := db.Prepare("UPDATE dev_boiler_a_info SET small_fire = ? WHERE device_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Small_Fire, device_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	if len(f.Water_Pump) > 0 {
		st, err1 := db.Prepare("UPDATE dev_boiler_a_info SET water_pump = ? WHERE device_sn = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Water_Pump, device_sn)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	return 0, nil

}

func (r *DeviceBoilerAInfo) DeleteById(id int64) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("DELETE FROM dev_boiler_a_info WHERE id = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	result, err := st.Exec(id)
	if err != nil {
		return ErrDatabase, err
	}

	num, _ := result.RowsAffected()
	if num > 0 {
		return 0, nil
	} else {
		return ErrNotFound, nil
	}
}


func (r *DeviceBoilerAInfo) DeleteByDeviceSN(device_sn string) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("DELETE FROM dev_boiler_a_info WHERE device_sn = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	result, err := st.Exec(device_sn)
	if err != nil {
		return ErrDatabase, err
	}

	num, _ := result.RowsAffected()
	if num > 0 {
		return 0, nil
	} else {
		return ErrNotFound, nil
	}
}
