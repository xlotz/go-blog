package models

import "go-blog/models/ctype"

// MenuModel 菜单表

type MenuModel struct {
	MODEL
	MenuTitle    string        `gorm:"size:32" json:"menu_title"` // 菜单
	MenuTitleEn  string        `gorm:"size:32" json:"menu_title_en"`
	Slogan       string        `gorm:"size:64" json:"slogan"`                                                                    // slogan
	Abstract     ctype.Array   `gorm:"type:string" json:"abstract"`                                                              // 简介
	AbstractTime int           `json:"abstract_time"`                                                                            // 简介的切换时间
	Banners      []BannerModel `gorm:"many2many:menu_image_models;JoinForeignKey:MenuID;JoinReferences:BannerID" json:"banners"` // 菜单的图片列表
	BannerTime   int           `json:"banner_time"`                                                                              // 菜单的图片切换时间
	Sort         int           `gorm:"size:10" json:"sort"`                                                                      // 顺序
}
