// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameExtrasCustomlink = "extras_customlink"

// ExtrasCustomlink mapped from table <extras_customlink>
type ExtrasCustomlink struct {
	ID            int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name          string    `gorm:"column:name;not null" json:"name"`
	LinkText      string    `gorm:"column:link_text;not null" json:"link_text"`
	LinkURL       string    `gorm:"column:link_url;not null" json:"link_url"`
	Weight        int16     `gorm:"column:weight;not null" json:"weight"`
	GroupName     string    `gorm:"column:group_name;not null" json:"group_name"`
	ButtonClass   string    `gorm:"column:button_class;not null" json:"button_class"`
	NewWindow     bool      `gorm:"column:new_window;not null" json:"new_window"`
	ContentTypeID int32     `gorm:"column:content_type_id;not null" json:"content_type_id"`
	Created       time.Time `gorm:"column:created" json:"created"`
	LastUpdated   time.Time `gorm:"column:last_updated" json:"last_updated"`
	Enabled       bool      `gorm:"column:enabled;not null" json:"enabled"`
}

// TableName ExtrasCustomlink's table name
func (*ExtrasCustomlink) TableName() string {
	return TableNameExtrasCustomlink
}
