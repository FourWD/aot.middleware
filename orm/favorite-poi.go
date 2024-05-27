package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type FavoritePoi struct { // ปลายทางใช้แสดงที่ระบบขายหน้าแรก
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	PoiID         string `json:"poi_id" query:"poi_id" gorm:"type:varchar(36);"`
	FleetClientID string `json:"fleet_client_id"  query:"fleet_client_id" gorm:"type:varchar(10);"`
	orm.GormRowOrder
}
