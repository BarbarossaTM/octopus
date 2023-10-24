// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTenancyContact = "tenancy_contact"

// TenancyContact mapped from table <tenancy_contact>
type TenancyContact struct {
	Created         time.Time `gorm:"column:created" json:"created"`
	LastUpdated     time.Time `gorm:"column:last_updated" json:"last_updated"`
	CustomFieldData string    `gorm:"column:custom_field_data;not null" json:"custom_field_data"`
	ID              int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name            string    `gorm:"column:name;not null" json:"name"`
	Title           string    `gorm:"column:title;not null" json:"title"`
	Phone           string    `gorm:"column:phone;not null" json:"phone"`
	Email           string    `gorm:"column:email;not null" json:"email"`
	Address         string    `gorm:"column:address;not null" json:"address"`
	Comments        string    `gorm:"column:comments;not null" json:"comments"`
	GroupID         int64     `gorm:"column:group_id" json:"group_id"`
	Link            string    `gorm:"column:link;not null" json:"link"`
}

// TableName TenancyContact's table name
func (*TenancyContact) TableName() string {
	return TableNameTenancyContact
}
