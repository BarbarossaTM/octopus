// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameVirtualizationVminterfaceTaggedVlan = "virtualization_vminterface_tagged_vlans"

// VirtualizationVminterfaceTaggedVlan mapped from table <virtualization_vminterface_tagged_vlans>
type VirtualizationVminterfaceTaggedVlan struct {
	ID            int64 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	VminterfaceID int64 `gorm:"column:vminterface_id;not null" json:"vminterface_id"`
	VlanID        int64 `gorm:"column:vlan_id;not null" json:"vlan_id"`
}

// TableName VirtualizationVminterfaceTaggedVlan's table name
func (*VirtualizationVminterfaceTaggedVlan) TableName() string {
	return TableNameVirtualizationVminterfaceTaggedVlan
}
