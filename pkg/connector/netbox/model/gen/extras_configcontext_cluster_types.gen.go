// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameExtrasConfigcontextClusterType = "extras_configcontext_cluster_types"

// ExtrasConfigcontextClusterType mapped from table <extras_configcontext_cluster_types>
type ExtrasConfigcontextClusterType struct {
	ID              int64 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ConfigcontextID int64 `gorm:"column:configcontext_id;not null" json:"configcontext_id"`
	ClustertypeID   int64 `gorm:"column:clustertype_id;not null" json:"clustertype_id"`
}

// TableName ExtrasConfigcontextClusterType's table name
func (*ExtrasConfigcontextClusterType) TableName() string {
	return TableNameExtrasConfigcontextClusterType
}
