// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameExtrasTag = "extras_tag"

// ExtrasTag mapped from table <extras_tag>
type ExtrasTag struct {
	Name        string    `gorm:"column:name;not null" json:"name"`
	Slug        string    `gorm:"column:slug;not null" json:"slug"`
//	Created     time.Time `gorm:"column:created" json:"created"`
//	LastUpdated time.Time `gorm:"column:last_updated" json:"last_updated"`
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
//	Color       string    `gorm:"column:color;not null" json:"color"`
//	Description string    `gorm:"column:description;not null" json:"description"`
}

// TableName ExtrasTag's table name
func (*ExtrasTag) TableName() string {
	return TableNameExtrasTag
}
