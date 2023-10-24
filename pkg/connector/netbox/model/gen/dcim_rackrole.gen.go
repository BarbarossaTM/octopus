// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameDcimRackrole = "dcim_rackrole"

// DcimRackrole mapped from table <dcim_rackrole>
type DcimRackrole struct {
	Created         time.Time `gorm:"column:created" json:"created"`
	LastUpdated     time.Time `gorm:"column:last_updated" json:"last_updated"`
	CustomFieldData string    `gorm:"column:custom_field_data;not null" json:"custom_field_data"`
	ID              int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name            string    `gorm:"column:name;not null" json:"name"`
	Slug            string    `gorm:"column:slug;not null" json:"slug"`
	Color           string    `gorm:"column:color;not null" json:"color"`
	Description     string    `gorm:"column:description;not null" json:"description"`
}

// TableName DcimRackrole's table name
func (*DcimRackrole) TableName() string {
	return TableNameDcimRackrole
}
