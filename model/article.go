package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         int    `json:"state"`
}

func (o Article) GetAllArticle() (result []Article, err error) {
	err = db.Model(o).Find(&result).Error
	return
}
