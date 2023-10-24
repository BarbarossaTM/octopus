// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTenancyTenantgroup = "tenancy_tenantgroup"

// TenancyTenantgroup mapped from table <tenancy_tenantgroup>
type TenancyTenantgroup struct {
	Created         time.Time `gorm:"column:created" json:"created"`
	LastUpdated     time.Time `gorm:"column:last_updated" json:"last_updated"`
	CustomFieldData string    `gorm:"column:custom_field_data;not null" json:"custom_field_data"`
	ID              int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name            string    `gorm:"column:name;not null" json:"name"`
	Slug            string    `gorm:"column:slug;not null" json:"slug"`
	Description     string    `gorm:"column:description;not null" json:"description"`
	Lft             int32     `gorm:"column:lft;not null" json:"lft"`
	Rght            int32     `gorm:"column:rght;not null" json:"rght"`
	TreeID          int32     `gorm:"column:tree_id;not null" json:"tree_id"`
	Level           int32     `gorm:"column:level;not null" json:"level"`
	ParentID        int64     `gorm:"column:parent_id" json:"parent_id"`
}

// TableName TenancyTenantgroup's table name
func (*TenancyTenantgroup) TableName() string {
	return TableNameTenancyTenantgroup
}
