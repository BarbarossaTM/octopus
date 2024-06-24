// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameExtrasTag = "extras_tag"

// ExtrasTag mapped from table <extras_tag>
type ExtrasTag struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name        string    `gorm:"column:name;not null" json:"name"`
	Slug        string    `gorm:"column:slug;not null" json:"slug"`
	// Color       string    `gorm:"column:color;not null" json:"color"`
	// Description string    `gorm:"column:description;not null" json:"description"`
	// Created     time.Time `gorm:"column:created" json:"created"`
	// LastUpdated time.Time `gorm:"column:last_updated" json:"last_updated"`
}

// TableName ExtrasTag's table name
func (*ExtrasTag) TableName() string {
	return TableNameExtrasTag
}
