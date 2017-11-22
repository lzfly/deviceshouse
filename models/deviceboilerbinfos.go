package models

import (
	"boilerserver/models/mymysql"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-sql-driver/mysql"
	"time"
)

type DeviceBoilerBInfo struct {
	Id            int64     `json:"id"`
	Device_sn            string `json:"device_sn"`
	Start_Temp           string  `json:"start_temp"`
	Target_Temp          string  `json:"target_temp"`
	Stop_Temp            string  `json:"stop_temp"`
	Out_Water_Temp       string  `json:"out_water_temp"`
    Smoke_Temp           string  `json:"smoke_temp"`
	Load                 string  `json:"load"`
	Gas                  string  `json:"gas"`
	Throttle             string  `json:"throttle"`
	Smoke                string  `json:"smoke"`
	Freq                 string  `json:"freq"`
	Run_State            string  `json:"run_state"`
}

func NewDeviceBoilerBInfo(f *DeviceBoilerBInfoPostForm, t time.Time) *DeviceBoilerBInfo {
	DeviceBoilerBinfo := DeviceBoilerBInfo{
		Device_sn:     f.Device_sn,
		Start_Temp:     f.Start_Temp,
		Target_Temp:     f.Target_Temp,
		Stop_Temp:     f.Stop_Temp,
		Out_Water_Temp:    f.Out_Water_Temp,
		Smoke_Temp:    f.Smoke_Temp,
		Load:     f.Load,
		Gas:     f.Gas,
		Throttle:     f.Throttle,
		Smoke:    f.Smoke,
		Freq:    f.Freq,
		Run_State:    f.Run_State,}

	return &DeviceBoilerBinfo
}

func (r *DeviceBoilerBInfo) Insert() (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("INSERT INTO dev_DeviceBoilerBinfo(device_sn, start_temp, target_temp, stop_temp, out_water_temp, smoke_temp, load, gas, throttle, smoke, freq, run_state) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	//if result, err := st.Exec(
	if _, err := st.Exec(r.Device_sn, r.Start_Temp, r.Target_Temp, r.Stop_Temp, r.Out_Water_Temp, r.Smoke_Temp, r.Load, r.Gas, r.Throttle, r.Smoke, r.Freq, r.Run_State); err != nil {
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

func (r *DeviceBoilerBInfo) FindById(id int64) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("SELECT id, device_sn, start_temp, target_temp, stop_temp, out_water_temp, smoke_temp, load, gas, throttle, smoke, freq, run_state FROM dev_DeviceBoilerBinfo WHERE id = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	row := st.QueryRow(id)
		
	var tmpId             sql.NullInt64 
	var tmpDevice_sn      sql.NullString 
	var tmpStart_Temp     sql.NullString
	var tmpTarget_Temp    sql.NullString
    var tmpStop_Temp      sql.NullString  
    var tmpOut_Water_Temp sql.NullString 
	var tmpSmoke_Temp     sql.NullString 
	var tmpLoad           sql.NullString
	var tmpGas            sql.NullString
    var tmpThrottle       sql.NullString  
    var tmpSmoke          sql.NullString 
	var tmpFreq           sql.NullString
	var tmpRun_State         sql.NullString



	if err := row.Scan(&tmpId, &tmpDevice_sn, &tmpStart_Temp, &tmpTarget_Temp, &tmpStop_Temp, &tmpOut_Water_Temp, &tmpSmoke_Temp,  &tmpLoad, &tmpGas, &tmpThrottle, &tmpSmoke, &tmpFreq, &Run_State); err != nil {
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
		if tmpStart_Temp.Valid {
			r.Start_Temp = tmpStart_Temp.String
		}
		if tmpTarget_Temp.Valid {
			r.Target_Temp = tmpTarget_Temp.String
		}
		if tmpStop_Temp.Valid {
			r.Stop_Temp = tmpStop_Temp.String
		}
		if tmpOut_Water_Temp.Valid {
			r.Out_Water_Temp = tmpOut_Water_Temp.String
		}
		if tmpSmoke_Temp.Valid {
			r.Smoke_Temp = tmpSmoke_Temp.String
		}
		if tmpLoad.Valid {
			r.Load = tmpLoad.String
		}
		if tmpGas.Valid {
			r.Gas = tmpGas.String
		}
		if tmpThrottle.Valid {
			r.Throttle = tmpThrottle.String
		}
		if tmpSmoke.Valid {
			r.Smoke = tmpSmoke.String
		}
		if tmpFreq.Valid {
			r.Freq = tmpFreq.String
		}
		if tmpRun_State.Valid {
			r.Run_State = tmpRun_State.String
		}


	return 0, nil
}

func (r *DeviceBoilerBInfo) FindByDeviceSN(device_sn string) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("SELECT id, device_sn, start_temp, target_temp, stop_temp, out_water_temp, smoke_temp, load, gas, throttle, smoke, freq, run_state FROM dev_DeviceBoilerBinfo WHERE device_sn = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	row := st.QueryRow(device_sn)

	var tmpId             sql.NullInt64 
	var tmpDevice_sn      sql.NullString 
	var tmpStart_Temp     sql.NullString
	var tmpTarget_Temp    sql.NullString
    var tmpStop_Temp      sql.NullString  
    var tmpOut_Water_Temp sql.NullString 
	var tmpSmoke_Temp     sql.NullString 
	var tmpLoad           sql.NullString
	var tmpGas            sql.NullString
    var tmpThrottle       sql.NullString  
    var tmpSmoke          sql.NullString 
	var tmpFreq           sql.NullString
	var Run_State         sql.NullString


	if err := row.Scan(&tmpId, &tmpDevice_sn, &tmpStart_Temp, &tmpTarget_Temp, &tmpStop_Temp, &tmpOut_Water_Temp, &tmpSmoke_Temp,  &tmpLoad, &tmpGas, &tmpThrottle, &tmpSmoke, &tmpFreq, &Run_State); err != nil {
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
		if tmpStart_Temp.Valid {
			r.Start_Temp = tmpStart_Temp.String
		}
		if tmpTarget_Temp.Valid {
			r.Target_Temp = tmpTarget_Temp.String
		}
		if tmpStop_Temp.Valid {
			r.Stop_Temp = tmpStop_Temp.String
		}
		if tmpOut_Water_Temp.Valid {
			r.Out_Water_Temp = tmpOut_Water_Temp.String
		}
		if tmpSmoke_Temp.Valid {
			r.Smoke_Temp = tmpSmoke_Temp.String
		}
		if tmpLoad.Valid {
			r.Load = tmpLoad.String
		}
		if tmpGas.Valid {
			r.Gas = tmpGas.String
		}
		if tmpThrottle.Valid {
			r.Throttle = tmpThrottle.String
		}
		if tmpSmoke.Valid {
			r.Smoke = tmpSmoke.String
		}
		if tmpFreq.Valid {
			r.Freq = tmpFreq.String
		}
		if tmpRun_State.Valid {
			r.Run_State = tmpRun_State.String
		}

	return 0, nil
}

func (r *DeviceBoilerBInfo) ClearPass() {
	
}

func GetAllDeviceBoilerBInfos(queryVal map[string]string, queryOp map[string]string,
	order map[string]string, limit int64,
	offset int64) (records []DeviceBoilerBInfo, err error) {
	sqlStr := "SELECT id, device_sn, start_temp, target_temp, stop_temp, out_water_temp, smoke_temp, load, gas, throttle, smoke, freq, run_state FROM dev_DeviceBoilerBinfo"
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

	records = make([]DeviceBoilerBInfo, 0, limit)
	for rows.Next() {

		var tmpId             sql.NullInt64 
		var tmpDevice_sn      sql.NullString 
		var tmpStart_Temp     sql.NullString
		var tmpTarget_Temp    sql.NullString
		var tmpStop_Temp      sql.NullString  
		var tmpOut_Water_Temp sql.NullString 
		var tmpSmoke_Temp     sql.NullString 
		var tmpLoad           sql.NullString
		var tmpGas            sql.NullString
		var tmpThrottle       sql.NullString  
		var tmpSmoke          sql.NullString 
		var tmpFreq           sql.NullString
		var Run_State         sql.NullString
		
		if err := rows.Scan(&tmpId, &tmpDevice_sn, &tmpStart_Temp, &tmpTarget_Temp, &tmpStop_Temp, &tmpOut_Water_Temp, &tmpSmoke_Temp,  &tmpLoad, &tmpGas, &tmpThrottle, &tmpSmoke, &tmpFreq, &Run_State); err != nil {
			return nil, err
		}

        r := DeviceBoilerBInfo{}
		if tmpId.Valid {
			r.Id = tmpId.Int64
		}
		if tmpDevice_sn.Valid {
			r.Device_sn = tmpDevice_sn.String
		}
		if tmpStart_Temp.Valid {
			r.Start_Temp = tmpStart_Temp.String
		}
		if tmpTarget_Temp.Valid {
			r.Target_Temp = tmpTarget_Temp.String
		}
		if tmpStop_Temp.Valid {
			r.Stop_Temp = tmpStop_Temp.String
		}
		if tmpOut_Water_Temp.Valid {
			r.Out_Water_Temp = tmpOut_Water_Temp.String
		}
		if tmpSmoke_Temp.Valid {
			r.Smoke_Temp = tmpSmoke_Temp.String
		}
		if tmpLoad.Valid {
			r.Load = tmpLoad.String
		}
		if tmpGas.Valid {
			r.Gas = tmpGas.String
		}
		if tmpThrottle.Valid {
			r.Throttle = tmpThrottle.String
		}
		if tmpSmoke.Valid {
			r.Smoke = tmpSmoke.String
		}
		if tmpFreq.Valid {
			r.Freq = tmpFreq.String
		}
		if tmpRun_State.Valid {
			r.Run_State = tmpRun_State.String
		}
		
		records = append(records, r)
		}
		if err := rows.Err(); err != nil {
			return nil, err
		}

	return records, nil
}

func (r *DeviceBoilerBInfo) UpdateById(id int64, f *DeviceBoilerBInfoPutForm) (code int, err error) {
	db := mymysql.Conn()
		
	if len(f.Device_sn) > 0 {
		st, err1 := db.Prepare("UPDATE dev_DeviceBoilerBinfo SET device_sn = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Device_sn, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	if len(f.Start_Temp) > 0 {
		st, err1 := db.Prepare("UPDATE dev_DeviceBoilerBinfo SET start_temp = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Start_Temp, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Target_Temp) > 0 {
		st, err1 := db.Prepare("UPDATE dev_DeviceBoilerBinfo SET target_temp = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Target_Temp, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Stop_Temp) > 0 {
		st, err1 := db.Prepare("UPDATE dev_DeviceBoilerBinfo SET stop_temp = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Stop_Temp, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	if len(f.Out_Water_Temp) > 0 {
		st, err1 := db.Prepare("UPDATE dev_DeviceBoilerBinfo SET out_water_temp = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Out_Water_Temp, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	if len(f.Smoke_Temp) > 0 {
		st, err1 := db.Prepare("UPDATE dev_DeviceBoilerBinfo SET smoke_temp = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Smoke_Temp, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	
	if len(f.Load) > 0 {
		st, err1 := db.Prepare("UPDATE dev_DeviceBoilerBinfo SET load = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Load, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	if len(f.Gas) > 0 {
		st, err1 := db.Prepare("UPDATE dev_DeviceBoilerBinfo SET gas = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Gas, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Throttle) > 0 {
		st, err1 := db.Prepare("UPDATE dev_DeviceBoilerBinfo SET throttle = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Throttle, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Smoke) > 0 {
		st, err1 := db.Prepare("UPDATE dev_DeviceBoilerBinfo SET smoke = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Smoke, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	if len(f.Freq) > 0 {
		st, err1 := db.Prepare("UPDATE dev_DeviceBoilerBinfo SET freq = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Freq, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	if len(f.Run_State) > 0 {
		st, err1 := db.Prepare("UPDATE dev_DeviceBoilerBinfo SET run_state = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Run_State, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	
	
	return 0, nil

}

func (r *DeviceBoilerBInfo) UpdateByDeviceSN(device_sn string, f *DeviceBoilerBInfoPutForm) (code int, err error) {
	db := mymysql.Conn()



	if len(f.Device_sn) > 0 {
		st, err1 := db.Prepare("UPDATE dev_DeviceBoilerBinfo SET device_sn = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Device_sn, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	if len(f.Start_Temp) > 0 {
		st, err1 := db.Prepare("UPDATE dev_DeviceBoilerBinfo SET start_temp = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Start_Temp, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Target_Temp) > 0 {
		st, err1 := db.Prepare("UPDATE dev_DeviceBoilerBinfo SET target_temp = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Target_Temp, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Stop_Temp) > 0 {
		st, err1 := db.Prepare("UPDATE dev_DeviceBoilerBinfo SET stop_temp = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Stop_Temp, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	if len(f.Out_Water_Temp) > 0 {
		st, err1 := db.Prepare("UPDATE dev_DeviceBoilerBinfo SET out_water_temp = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Out_Water_Temp, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	if len(f.Smoke_Temp) > 0 {
		st, err1 := db.Prepare("UPDATE dev_DeviceBoilerBinfo SET smoke_temp = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Smoke_Temp, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	
	if len(f.Load) > 0 {
		st, err1 := db.Prepare("UPDATE dev_DeviceBoilerBinfo SET load = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Load, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	if len(f.Gas) > 0 {
		st, err1 := db.Prepare("UPDATE dev_DeviceBoilerBinfo SET gas = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Gas, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Throttle) > 0 {
		st, err1 := db.Prepare("UPDATE dev_DeviceBoilerBinfo SET throttle = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Throttle, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}
	if len(f.Smoke) > 0 {
		st, err1 := db.Prepare("UPDATE dev_DeviceBoilerBinfo SET smoke = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Smoke, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	if len(f.Freq) > 0 {
		st, err1 := db.Prepare("UPDATE dev_DeviceBoilerBinfo SET freq = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Freq, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	if len(f.Run_State) > 0 {
		st, err1 := db.Prepare("UPDATE dev_DeviceBoilerBinfo SET run_state = ? WHERE id = ?")
		if err1 != nil {
			return ErrDatabase, err1
		}
		defer st.Close()

		_, err2 := st.Exec(f.Run_State, id)
		if err2 != nil {
			return ErrDatabase, err2
		}
	}

	return 0, nil

}

func (r *DeviceBoilerBInfo) DeleteById(id int64) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("DELETE FROM dev_DeviceBoilerBinfo WHERE id = ?")
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


func (r *DeviceBoilerBInfo) DeleteByDeviceSN(device_sn string) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("DELETE FROM dev_DeviceBoilerBinfo WHERE device_sn = ?")
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
