// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameIpamFhrpgroup = "ipam_fhrpgroup"

// IpamFhrpgroup mapped from table <ipam_fhrpgroup>
type IpamFhrpgroup struct {
	Created         time.Time `gorm:"column:created" json:"created"`
	LastUpdated     time.Time `gorm:"column:last_updated" json:"last_updated"`
	CustomFieldData string    `gorm:"column:custom_field_data;not null" json:"custom_field_data"`
	ID              int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	GroupID         int16     `gorm:"column:group_id;not null" json:"group_id"`
	Protocol        string    `gorm:"column:protocol;not null" json:"protocol"`
	AuthType        string    `gorm:"column:auth_type;not null" json:"auth_type"`
	AuthKey         string    `gorm:"column:auth_key;not null" json:"auth_key"`
	Description     string    `gorm:"column:description;not null" json:"description"`
}

// TableName IpamFhrpgroup's table name
func (*IpamFhrpgroup) TableName() string {
	return TableNameIpamFhrpgroup
}
