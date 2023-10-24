// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameDcimCablepath = "dcim_cablepath"

// DcimCablepath mapped from table <dcim_cablepath>
type DcimCablepath struct {
	ID                int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	OriginID          int64  `gorm:"column:origin_id;not null" json:"origin_id"`
	DestinationID     int64  `gorm:"column:destination_id" json:"destination_id"`
	Path              string `gorm:"column:path;not null" json:"path"`
	IsActive          bool   `gorm:"column:is_active;not null" json:"is_active"`
	IsSplit           bool   `gorm:"column:is_split;not null" json:"is_split"`
	DestinationTypeID int32  `gorm:"column:destination_type_id" json:"destination_type_id"`
	OriginTypeID      int32  `gorm:"column:origin_type_id;not null" json:"origin_type_id"`
}

// TableName DcimCablepath's table name
func (*DcimCablepath) TableName() string {
	return TableNameDcimCablepath
}
