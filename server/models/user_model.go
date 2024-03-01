package models

import "go-blog/models/ctype"

// 用户表

type UserModel struct {
	MODEL
	NickName string         `gorm:"size:36" json:"nick_name"`          // 昵称
	UserName string         `gorm:"size:36" json:"user_name"`          // 用户名
	Password string         `gorm:"size:128" json:"password"`          // 密码
	Avatar   string         `gorm:"size:256" json:"avatar"`            // 头像
	Email    string         `gorm:"size:128" json:"email"`             // 邮箱
	Tel      string         `gorm:"size:18" json:"tel"`                // 手机
	Addr     string         `gorm:"size:64" json:"addr"`               // 地址
	Token    string         `gorm:"size:64" json:"token"`              // token
	IP       string         `gorm:"size:18" json:"ip"`                 // ip
	Role     ctype.Role     `gorm:"size:4;default:1" json:"role"`      // 权限、角色
	SignType ctype.SignType `gorm:"type=smallint(6)" json:"sign_type"` // 注册来源
	//ArticleModels  []ArticleModel `gorm:"foreignKey:AuthID" json:"-"`                                                       // 发表的文章列表
	//CollectsModels []ArticleModel `gorm:"many2many:auth2_collects;joinForeignKey:AuthID;JoinReferences:ArticleID" json:"-"` // 收藏的文件列表
}
