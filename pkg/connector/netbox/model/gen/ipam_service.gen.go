// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameIpamService = "ipam_service"

// IpamService mapped from table <ipam_service>
type IpamService struct {
	Created          time.Time `gorm:"column:created" json:"created"`
	LastUpdated      time.Time `gorm:"column:last_updated" json:"last_updated"`
	CustomFieldData  string    `gorm:"column:custom_field_data;not null" json:"custom_field_data"`
	ID               int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name             string    `gorm:"column:name;not null" json:"name"`
	Protocol         string    `gorm:"column:protocol;not null" json:"protocol"`
	Ports            string    `gorm:"column:ports;not null" json:"ports"`
	Description      string    `gorm:"column:description;not null" json:"description"`
	DeviceID         int64     `gorm:"column:device_id" json:"device_id"`
	VirtualMachineID int64     `gorm:"column:virtual_machine_id" json:"virtual_machine_id"`
}

// TableName IpamService's table name
func (*IpamService) TableName() string {
	return TableNameIpamService
}
