// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTaggitTaggeditem = "taggit_taggeditem"

// TaggitTaggeditem mapped from table <taggit_taggeditem>
type TaggitTaggeditem struct {
	ID            int32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ObjectID      int32 `gorm:"column:object_id;not null" json:"object_id"`
	ContentTypeID int32 `gorm:"column:content_type_id;not null" json:"content_type_id"`
	TagID         int32 `gorm:"column:tag_id;not null" json:"tag_id"`
}

// TableName TaggitTaggeditem's table name
func (*TaggitTaggeditem) TableName() string {
	return TableNameTaggitTaggeditem
}
