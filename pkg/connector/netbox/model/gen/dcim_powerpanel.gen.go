// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameDcimPowerpanel = "dcim_powerpanel"

// DcimPowerpanel mapped from table <dcim_powerpanel>
type DcimPowerpanel struct {
	Created         time.Time `gorm:"column:created" json:"created"`
	LastUpdated     time.Time `gorm:"column:last_updated" json:"last_updated"`
	CustomFieldData string    `gorm:"column:custom_field_data;not null" json:"custom_field_data"`
	ID              int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name            string    `gorm:"column:name;not null" json:"name"`
	LocationID      int64     `gorm:"column:location_id" json:"location_id"`
	SiteID          int64     `gorm:"column:site_id;not null" json:"site_id"`
}

// TableName DcimPowerpanel's table name
func (*DcimPowerpanel) TableName() string {
	return TableNameDcimPowerpanel
}
