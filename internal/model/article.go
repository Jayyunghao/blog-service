package model

import "github.com/jinzhu/gorm"

func (a Article) Count(db *gorm.DB) (int, error) {
	var count int
	if a.Title != "" {
		searchStr := "%" + a.Title + "%"
		db = db.Where("title like ?", searchStr)
	}
	if a.Desc != "" {
		searchStr := "%" + a.Desc + "%"
		db = db.Where("desc like ?", searchStr)
	}
	if a.Content != "" {
		searchStr := "%" + a.Content + "%"
		db =db.Where("content like ?", searchStr)
	}
	db = db.Where("state = ?", a.State)
	if err := db.Model(&a).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (a Article) List(db *gorm.DB, pageOffset, pageSize int) ([]*Article, error) {
	var articles []*Article
	var err error
	if pageOffset > 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if a.Title != "" {
		searchStr := "%" + a.Title + "%"
		db = db.Where("title like ?", searchStr)
	}
	if a.Desc != "" {
		searchStr := "%" + a.Desc + "%"
		db = db.Where("desc like ?", searchStr)
	}
	db = db.Where("state = ?", a.State)
	if err = db.Model(&a).Where("is_del = ?", 0).Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, err
}

func (a Article) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a Article) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(&Article{}).Where("id = ? AND is_del = ?", a.ID, 0).Update(values).Error; err != nil {
		return err
	}
	return nil
}

func (a Article) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", a.Model.ID, 0).Delete(&a).Error
}
