// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameDcimPowerfeed = "dcim_powerfeed"

// DcimPowerfeed mapped from table <dcim_powerfeed>
type DcimPowerfeed struct {
	Created         time.Time `gorm:"column:created" json:"created"`
	LastUpdated     time.Time `gorm:"column:last_updated" json:"last_updated"`
	CustomFieldData string    `gorm:"column:custom_field_data;not null" json:"custom_field_data"`
	ID              int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	LinkPeerID      int64     `gorm:"column:_link_peer_id" json:"_link_peer_id"`
	MarkConnected   bool      `gorm:"column:mark_connected;not null" json:"mark_connected"`
	Name            string    `gorm:"column:name;not null" json:"name"`
	Status          string    `gorm:"column:status;not null" json:"status"`
	Type            string    `gorm:"column:type;not null" json:"type"`
	Supply          string    `gorm:"column:supply;not null" json:"supply"`
	Phase           string    `gorm:"column:phase;not null" json:"phase"`
	Voltage         int16     `gorm:"column:voltage;not null" json:"voltage"`
	Amperage        int16     `gorm:"column:amperage;not null" json:"amperage"`
	MaxUtilization  int16     `gorm:"column:max_utilization;not null" json:"max_utilization"`
	AvailablePower  int32     `gorm:"column:available_power;not null" json:"available_power"`
	Comments        string    `gorm:"column:comments;not null" json:"comments"`
	LinkPeerTypeID  int32     `gorm:"column:_link_peer_type_id" json:"_link_peer_type_id"`
	PathID          int64     `gorm:"column:_path_id" json:"_path_id"`
	CableID         int64     `gorm:"column:cable_id" json:"cable_id"`
	PowerPanelID    int64     `gorm:"column:power_panel_id;not null" json:"power_panel_id"`
	RackID          int64     `gorm:"column:rack_id" json:"rack_id"`
}

// TableName DcimPowerfeed's table name
func (*DcimPowerfeed) TableName() string {
	return TableNameDcimPowerfeed
}
