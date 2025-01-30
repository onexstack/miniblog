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

const TableNamePostM = "post"

// PostM mapped from table <post>
type PostM struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID    string    `gorm:"column:userID;not null;comment:用户唯一 ID" json:"userID"`                                  // 用户唯一 ID
	PostID    string    `gorm:"column:postID;not null;uniqueIndex:idx_post_postID;comment:博文唯一 ID" json:"postID"`      // 博文唯一 ID
	Title     string    `gorm:"column:title;not null;comment:博文标题" json:"title"`                                       // 博文标题
	Content   string    `gorm:"column:content;not null;comment:博文内容" json:"content"`                                   // 博文内容
	CreatedAt time.Time `gorm:"column:createdAt;not null;default:current_timestamp;comment:博文创建时间" json:"createdAt"`   // 博文创建时间
	UpdatedAt time.Time `gorm:"column:updatedAt;not null;default:current_timestamp;comment:博文最后修改时间" json:"updatedAt"` // 博文最后修改时间
}

// TableName PostM's table name
func (*PostM) TableName() string {
	return TableNamePostM
}
