// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSocialAuthPartial = "social_auth_partial"

// SocialAuthPartial mapped from table <social_auth_partial>
type SocialAuthPartial struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Token     string    `gorm:"column:token;not null" json:"token"`
	NextStep  int16     `gorm:"column:next_step;not null" json:"next_step"`
	Backend   string    `gorm:"column:backend;not null" json:"backend"`
	Data      string    `gorm:"column:data;not null" json:"data"`
	Timestamp time.Time `gorm:"column:timestamp;not null" json:"timestamp"`
}

// TableName SocialAuthPartial's table name
func (*SocialAuthPartial) TableName() string {
	return TableNameSocialAuthPartial
}
