// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameDcimPlatform = "dcim_platform"

// DcimPlatform mapped from table <dcim_platform>
type DcimPlatform struct {
//	Created         time.Time `gorm:"column:created" json:"created"`
//	LastUpdated     time.Time `gorm:"column:last_updated" json:"last_updated"`
//	CustomFieldData string    `gorm:"column:custom_field_data;not null" json:"custom_field_data"`
	ID              int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name            string    `gorm:"column:name;not null" json:"name"`
	Slug            string    `gorm:"column:slug;not null" json:"slug"`
//	NapalmDriver    string    `gorm:"column:napalm_driver;not null" json:"napalm_driver"`
//	NapalmArgs      string    `gorm:"column:napalm_args" json:"napalm_args"`
//	Description     string    `gorm:"column:description;not null" json:"description"`
	ManufacturerID  int64     `gorm:"column:manufacturer_id" json:"manufacturer_id"`
}

// TableName DcimPlatform's table name
func (*DcimPlatform) TableName() string {
	return TableNameDcimPlatform
}
