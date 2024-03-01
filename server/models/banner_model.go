package models

type BannerModel struct {
	MODEL
	Path string `json:"path"`                // 图片路径
	Hash string `json:"hash"`                // 图片的hash 值
	Name string `gorm:"size:38" json:"name"` // 图片名称
	//ArticleModels []ArticleModel `gorm:"foreignKey:CoverID" json:"-"`
}
