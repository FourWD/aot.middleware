package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type FavoritePoi struct { // ปลายทางใช้แสดงที่ระบบขายหน้าแรก
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	PoiID         string `db:"poi_id" json:"poi_id" gorm:"type:varchar(36);"`
	FleetClientID string `db:"fleet_client_id"  json:"fleet_client_id" gorm:"type:varchar(10);"`
	RowOrder      int    `db:"row_order" json:"row_order" gorm:"type:int"`
}
