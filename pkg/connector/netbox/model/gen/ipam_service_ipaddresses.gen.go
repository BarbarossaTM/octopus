// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameIpamServiceIpaddress = "ipam_service_ipaddresses"

// IpamServiceIpaddress mapped from table <ipam_service_ipaddresses>
type IpamServiceIpaddress struct {
	ID          int64 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ServiceID   int64 `gorm:"column:service_id;not null" json:"service_id"`
	IpaddressID int64 `gorm:"column:ipaddress_id;not null" json:"ipaddress_id"`
}

// TableName IpamServiceIpaddress's table name
func (*IpamServiceIpaddress) TableName() string {
	return TableNameIpamServiceIpaddress
}
