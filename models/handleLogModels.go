package models

import (
	"github.com/asofdate/auth-core/entity"
	"github.com/hzwy23/dbobj"
	"github.com/hzwy23/utils/logger"
	"github.com/hzwy23/utils/validator"
)

type HandleLogMode struct {
}

func (this *HandleLogMode) Post(log_buf []entity.HandleLogBuf) {
	tx, err := dbobj.Begin()
	if err != nil {
		logger.Error(err)
	}

	for _, val := range log_buf {
		_, err := tx.Exec(sys_rdbms_051, val.User_id, val.Client_ip, val.Ret_status, val.Req_method, val.Req_url, val.Req_body)
		if err != nil {
			tx.Rollback()
			logger.Error("同步日志信息到数据库失败")
			return
		}
	}
	err = tx.Commit()
	if err != nil {
		logger.Error("同步日志信息到数据库失败")
	}
}

func (this HandleLogMode) Download() ([]entity.HandleLogs, error) {
	var rst []entity.HandleLogs
	rows, err := dbobj.Query(sys_rdbms_012)
	if err != nil {
		logger.Error(err)
		return rst, err
	}
	err = dbobj.Scan(rows, &rst)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return rst, nil
}

func (this HandleLogMode) getTotal() (total int64, err error) {
	err = dbobj.QueryRow(sys_rdbms_030).Scan(&total)
	return
}

func (this HandleLogMode) Get(offset, limit string) ([]entity.HandleLogs, int64, error) {
	var rst []entity.HandleLogs
	rows, err := dbobj.Query(sys_rdbms_029, offset, limit)
	if err != nil {
		logger.Error(err)
		return nil, 0, err
	}
	err = dbobj.Scan(rows, &rst)
	if err != nil {
		logger.Error(err)
		return nil, 0, err
	}
	total, err := this.getTotal()
	if err != nil {
		logger.Error(err)
		return nil, 0, err
	}
	return rst, total, nil
}

func (this HandleLogMode) Search(userid, start, end, offset, limit string) ([]entity.HandleLogs, int64, error) {
	var rst []entity.HandleLogs
	var cnt int64 = 0
	if len(userid) != 0 && validator.IsDate(start) && validator.IsDate(end) {
		rows, err := dbobj.Query(sys_rdbms_031, userid, start, end, offset, limit)
		defer rows.Close()
		if err != nil {
			return nil, cnt, err
		}
		err = dbobj.Scan(rows, &rst)
		if err != nil {
			logger.Error(err)
			return nil, cnt, err
		}
		dbobj.QueryForObject(sys_rdbms_034, dbobj.PackArgs(userid, start, end), &cnt)

	} else if len(userid) != 0 && validator.IsDate(start) {

		rows, err := dbobj.Query(sys_rdbms_032, userid, start, offset, limit)
		defer rows.Close()
		if err != nil {
			logger.Error(err)
			return nil, cnt, err
		}
		err = dbobj.Scan(rows, &rst)
		if err != nil {
			logger.Error(err)
			return nil, cnt, err
		}
		dbobj.QueryForObject(sys_rdbms_083, dbobj.PackArgs(userid, start), &cnt)

	} else if len(userid) != 0 && validator.IsDate(end) {

		rows, err := dbobj.Query(sys_rdbms_031, userid, start, end, offset, limit)
		defer rows.Close()
		if err != nil {
			logger.Error(err)
			return nil, cnt, err
		}
		err = dbobj.Scan(rows, &rst)
		if err != nil {
			logger.Error(err)
			return nil, cnt, err
		}

		dbobj.QueryForObject(sys_rdbms_034, dbobj.PackArgs(userid, start, end), &cnt)

	} else if validator.IsDate(start) && validator.IsDate(end) {
		rows, err := dbobj.Query(sys_rdbms_033, start, end, offset, limit)
		defer rows.Close()
		if err != nil {
			logger.Error(err)
			return nil, cnt, err
		}
		err = dbobj.Scan(rows, &rst)
		if err != nil {
			logger.Error(err)
			return nil, cnt, err
		}

		dbobj.QueryForObject(sys_rdbms_086, dbobj.PackArgs(start, end), &cnt)

	} else if validator.IsDate(start) {
		rows, err := dbobj.Query(sys_rdbms_035, start, offset, limit)
		defer rows.Close()
		if err != nil {
			logger.Error(err)
			return nil, cnt, err
		}
		err = dbobj.Scan(rows, &rst)
		if err != nil {
			logger.Error(err)
			return nil, cnt, err
		}

		dbobj.QueryForObject(sys_rdbms_087, dbobj.PackArgs(start), &cnt)

	} else if validator.IsDate(end) {
		rows, err := dbobj.Query(sys_rdbms_039, end, offset, limit)
		defer rows.Close()
		if err != nil {
			logger.Error(err)
			return nil, cnt, err
		}
		err = dbobj.Scan(rows, &rst)
		if err != nil {
			logger.Error(err)
			return nil, cnt, err
		}

		dbobj.QueryForObject(sys_rdbms_003, dbobj.PackArgs(end), &cnt)
	} else if len(userid) != 0 {
		rows, err := dbobj.Query(sys_rdbms_040, userid, offset, limit)
		defer rows.Close()
		if err != nil {
			logger.Error(err)
			return nil, cnt, err
		}
		err = dbobj.Scan(rows, &rst)
		if err != nil {
			logger.Error(err)
			return nil, cnt, err
		}

		dbobj.QueryForObject(sys_rdbms_001, dbobj.PackArgs(userid), &cnt)

	} else {
		rows, err := dbobj.Query(sys_rdbms_042, offset, limit)
		defer rows.Close()
		if err != nil {
			logger.Error(err)
			return nil, cnt, err
		}
		err = dbobj.Scan(rows, &rst)
		if err != nil {
			logger.Error(err)
			return nil, cnt, err
		}

		dbobj.QueryForObject(sys_rdbms_106, dbobj.PackArgs(), &cnt)

	}
	return rst, cnt, nil
}
