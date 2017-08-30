package models

import (
	"sync"

	"github.com/asofdate/auth-core/entity"
	"github.com/hzwy23/dbobj"
	"github.com/hzwy23/utils"
	"github.com/hzwy23/utils/logger"
)

type RoleAndResourceModel struct {
	mres ResourceModel
	lock sync.RWMutex
}

func (this *RoleAndResourceModel) Delete(role_id string, resSlick []string) error {
	this.lock.Lock()
	defer this.lock.Unlock()

	var mp = make(map[string]bool)
	for _, val := range resSlick {
		mp[val] = true
	}

	// 获取已经拥有的角色
	all, err := this.Get(role_id)
	if err != nil {
		logger.Error(err)
		return err
	}

	tx, err := dbobj.Begin()
	if err != nil {
		logger.Error(err)
		return err
	}

	for _, val := range resSlick {
		var rst []entity.ResData
		this.dfs(all, val, &rst)
		if len(rst) == 0 {
			_, err = tx.Exec(sys_rdbms_093, role_id, val)
			if err != nil {
				logger.Error(err)
				tx.Rollback()
				return err
			}
		} else {
			var flag = true
			for _, s := range rst {
				if _, yes := mp[s.ResId]; !yes {
					flag = false
					break
				}
			}
			if flag {
				_, err = tx.Exec(sys_rdbms_093, role_id, val)
				if err != nil {
					logger.Error(err)
					tx.Rollback()
					return err
				}
			}
		}
	}
	return tx.Commit()
}

func (this *RoleAndResourceModel) Post(role_id string, resSlick []string) error {
	this.lock.Lock()
	defer this.lock.Unlock()
	// 获取所有资源
	all, err := this.mres.Get()
	if err != nil {
		logger.Error(err)
		return err
	}

	// 获取这个角色已经拥有的资源
	getted, err := this.Get(role_id)
	if err != nil {
		logger.Error(err)
		return err
	}

	var mp = make(map[string]bool)
	for _, val := range getted {
		mp[val.ResId] = true
	}

	tx, err := dbobj.Begin()
	if err != nil {
		logger.Error(err)
		return err
	}

	for _, val := range resSlick {
		var rst []entity.ResData
		this.parent(all, val, &rst)
		if len(rst) == 0 {
			if _, yes := mp[val]; !yes {
				mp[val] = true
				_, err = tx.Exec(sys_rdbms_074, utils.JoinCode(role_id, val), role_id, val)
				if err != nil {
					logger.Error(err)
					tx.Rollback()
					return err
				}
			}
		} else {
			for _, s := range rst {
				if _, yes := mp[s.ResId]; !yes {
					mp[s.ResId] = true
					_, err = tx.Exec(sys_rdbms_074, utils.JoinCode(role_id, s.ResId), role_id, s.ResId)
					if err != nil {
						logger.Error(err)
						tx.Rollback()
						return err
					}
				}
			}
		}
	}
	return tx.Commit()
}

// 查询没有获取到的资源信息
func (this *RoleAndResourceModel) UnGetted(role_id string) ([]entity.ResData, error) {

	var rst []entity.ResData

	// 获取角色已经拥有了的资源id
	role_res, err := this.get(role_id)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	// 获取所有的用户信息
	rst_res, err := this.mres.Get()
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	var diff = make(map[string]entity.ResData)
	for _, val := range rst_res {
		diff[val.ResId] = val
	}

	for _, val := range role_res {
		delete(diff, val.ResId)
	}

	// 修复差异项父节点
	tmp := this.searchParent(diff, rst_res)
	for len(tmp) != 0 {
		for _, val := range tmp {
			diff[val.ResId] = val
		}
		tmp = this.searchParent(diff, rst_res)
	}
	for _, val := range diff {
		rst = append(rst, val)
	}
	return rst, nil
}

// 查询角色已经拥有的资源信息
func (this RoleAndResourceModel) Get(role_id string) ([]entity.ResData, error) {

	var rst []entity.ResData

	role_res, err := this.get(role_id)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	rst_res, err := this.mres.Get()
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	for _, val := range role_res {
		for _, res := range rst_res {
			if val.ResId == res.ResId {
				var one entity.ResData
				one.ResId = res.ResId
				one.ResName = res.ResName
				one.ResUpid = res.ResUpid
				rst = append(rst, one)
				break
			}
		}
	}
	return rst, nil
}

// 获取某些角色,指定资源的所有下级资源
func (this *RoleAndResourceModel) Gets(roles []string, res_id ...string) ([]entity.ResData, error) {

	var rst []entity.ResData
	var role_res map[string]string = make(map[string]string)
	for _, val := range roles {
		tmp, err := this.get(val)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		for _, p := range tmp {
			role_res[p.ResId] = ""
		}
	}

	var rst_res []entity.ResData
	if len(res_id) == 1 {
		var err error
		rst_res, err = this.mres.GetChildren(res_id[0])
		if err != nil {
			logger.Error(err)
			return nil, err
		}
	} else if len(res_id) == 2 {
		tmp, err := this.mres.GetChildren(res_id[0])
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		for _, val := range tmp {
			if val.Restype == res_id[1] {
				rst_res = append(rst_res, val)
			}
		}
	} else {
		var err error
		rst_res, err = this.mres.Get()
		if err != nil {
			logger.Error(err)
			return nil, err
		}
	}

	for _, res := range rst_res {
		if _, ok := role_res[res.ResId]; ok {
			var one entity.ResData
			one.ResId = res.ResId
			one.ResName = res.ResName
			one.ResUpid = res.ResUpid
			one.InnerFlag = res.InnerFlag
			one.Restype = res.Restype
			one.ResAttr = res.ResAttr
			one.ServiceCd = res.ServiceCd
			rst = append(rst, one)
		}
	}
	return rst, nil
}

func (this *RoleAndResourceModel) parent(all []entity.ResData, resId string, ret *[]entity.ResData) {
	for _, val := range all {
		if val.ResId == resId {
			*ret = append(*ret, val)
			if val.ResUpid != val.ResId {
				this.parent(all, val.ResUpid, ret)
			}
		}
	}
}

// 查找所有的父级资源信息
func (this RoleAndResourceModel) searchParent(diff map[string]entity.ResData, all []entity.ResData) []entity.ResData {
	var ret []entity.ResData
	for _, val := range diff {
		if _, ok := diff[val.ResUpid]; !ok {
			for _, vl := range all {
				if vl.ResId == val.ResUpid {
					ret = append(ret, vl)
				}
			}
		}
	}
	return ret
}

func (this *RoleAndResourceModel) dfs(rst []entity.ResData, resId string, ret *[]entity.ResData) {
	for _, val := range rst {
		if resId == val.ResUpid {
			*ret = append(*ret, val)
			if val.ResId != val.ResUpid {
				this.dfs(rst, val.ResId, ret)
			}
		}
	}
}

// 获取指定角色拥有的资源ID列表
func (this *RoleAndResourceModel) get(role_id string) ([]entity.RoleResourceRelData, error) {

	var rst []entity.RoleResourceRelData
	rows, err := dbobj.Query(sys_rdbms_100, role_id)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	err = dbobj.Scan(rows, &rst)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return rst, nil
}

func (this *RoleAndResourceModel) CheckResIDAuth(userId string, resId string) bool {
	cnt := 0
	err := dbobj.QueryRow(sys_rdbms_108, userId, resId).Scan(&cnt)
	if err != nil {
		logger.Error(err)
		return false
	}
	if cnt > 0 {
		return true
	}
	return false
}

func (this *RoleAndResourceModel) CheckUrlAuth(userId string, url string) bool {
	cnt := 0
	err := dbobj.QueryRow(sys_rdbms_098, userId, url).Scan(&cnt)
	if err != nil {
		logger.Error(err)
		return false
	}
	if cnt > 0 {
		return true
	}
	logger.Error("insufficient privileges", "user id is :", userId, "api is :", url)
	return false
}
