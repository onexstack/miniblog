// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/miniblog. The professional
// version of this repository is https://github.com/onexstack/onex.

// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUserM = "user"

// UserM mapped from table <user>
type UserM struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID    string    `gorm:"column:userID;not null;uniqueIndex:idx_user_userID;comment:用户唯一 ID" json:"userID"`       // 用户唯一 ID
	Username  string    `gorm:"column:username;not null;uniqueIndex:idx_user_username;comment:用户名（唯一）" json:"username"` // 用户名（唯一）
	Password  string    `gorm:"column:password;not null;comment:用户密码（加密后）" json:"password"`                             // 用户密码（加密后）
	Nickname  string    `gorm:"column:nickname;not null;comment:用户昵称" json:"nickname"`                                  // 用户昵称
	Email     string    `gorm:"column:email;not null;comment:用户电子邮箱地址" json:"email"`                                    // 用户电子邮箱地址
	Phone     string    `gorm:"column:phone;not null;uniqueIndex:idx_user_phone;comment:用户手机号" json:"phone"`            // 用户手机号
	CreatedAt time.Time `gorm:"column:createdAt;not null;default:current_timestamp;comment:用户创建时间" json:"createdAt"`    // 用户创建时间
	UpdatedAt time.Time `gorm:"column:updatedAt;not null;default:current_timestamp;comment:用户最后修改时间" json:"updatedAt"`  // 用户最后修改时间
}

// TableName UserM's table name
func (*UserM) TableName() string {
	return TableNameUserM
}
