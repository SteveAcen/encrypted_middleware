// 自动生成模板Unit
package encrypted_middleware

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Unit 结构体
type Unit struct {
	global.GVA_MODEL
	UnitName *int   `json:"unitName" form:"unitName" gorm:"column:unit_name;comment:;"`
	DeviceId string `json:"deviceId" form:"deviceId" gorm:"column:device_id;comment:;"`
}

// TableName Unit 表名
func (Unit) TableName() string {
	return "unit"
}
