package models

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	Name       string         `json:"name" gorm:"not null"`
	RealName   string         `json:"realname" gorm:"column:realname;not null"`
	AvatarURL  string         `json:"avatar_url" gorm:"column:avatar_url"` // 存储头像的URL
	AvatarBlob []byte         `json:"-" gorm:"column:avatar_blob"`
	Phone      string         `json:"phone"`
	Wechat     string         `json:"wechat"`
	Position   string         `json:"position"`
	Email      string         `json:"email"`
	Region     string         `json:"region"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// PersonResponse 用于返回给前端的结构体，包含头像的base64编码
type PersonResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	RealName  string    `json:"realname"`
	AvatarURL string	`json:"avatar_url"`
	Avatar    string    `json:"avatar"` // base64编码的头像
	Phone     string    `json:"phone"`
	Wechat    string    `json:"wechat"`
	Position  string    `json:"position"`
	Email     string    `json:"email"`
	Region    string    `json:"region"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Person) TableName() string {
	return "persons"
}
