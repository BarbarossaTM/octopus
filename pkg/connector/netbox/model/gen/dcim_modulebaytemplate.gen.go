// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameDcimModulebaytemplate = "dcim_modulebaytemplate"

// DcimModulebaytemplate mapped from table <dcim_modulebaytemplate>
type DcimModulebaytemplate struct {
	Created      time.Time `gorm:"column:created" json:"created"`
	LastUpdated  time.Time `gorm:"column:last_updated" json:"last_updated"`
	ID           int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name         string    `gorm:"column:name;not null" json:"name"`
	Label        string    `gorm:"column:label;not null" json:"label"`
	Position     string    `gorm:"column:position;not null" json:"position"`
	Description  string    `gorm:"column:description;not null" json:"description"`
	DeviceTypeID int64     `gorm:"column:device_type_id;not null" json:"device_type_id"`
}

// TableName DcimModulebaytemplate's table name
func (*DcimModulebaytemplate) TableName() string {
	return TableNameDcimModulebaytemplate
}
