package encrypted_middleware

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/encrypted_middleware"
	encrypted_middlewareReq "github.com/flipped-aurora/gin-vue-admin/server/model/encrypted_middleware/request"
)

type UnitService struct {
}

// CreateUnit 创建Unit记录
// Author [piexlmax](https://github.com/piexlmax)
func (unitService *UnitService) CreateUnit(unit encrypted_middleware.Unit) (err error) {
	err = global.GVA_DB.Create(&unit).Error
	return err
}

// DeleteUnit 删除Unit记录
// Author [piexlmax](https://github.com/piexlmax)
func (unitService *UnitService) DeleteUnit(unit encrypted_middleware.Unit) (err error) {
	err = global.GVA_DB.Delete(&unit).Error
	return err
}

// DeleteUnitByIds 批量删除Unit记录
// Author [piexlmax](https://github.com/piexlmax)
func (unitService *UnitService) DeleteUnitByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]encrypted_middleware.Unit{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateUnit 更新Unit记录
// Author [piexlmax](https://github.com/piexlmax)
func (unitService *UnitService) UpdateUnit(unit encrypted_middleware.Unit) (err error) {
	err = global.GVA_DB.Save(&unit).Error
	return err
}

// GetUnit 根据id获取Unit记录
// Author [piexlmax](https://github.com/piexlmax)
func (unitService *UnitService) GetUnit(id uint) (unit encrypted_middleware.Unit, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&unit).Error
	return
}

// GetUnitInfoList 分页获取Unit记录
// Author [piexlmax](https://github.com/piexlmax)
func (unitService *UnitService) GetUnitInfoList(info encrypted_middlewareReq.UnitSearch) (list []encrypted_middleware.Unit, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&encrypted_middleware.Unit{})
	var units []encrypted_middleware.Unit
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&units).Error
	return units, total, err
}
