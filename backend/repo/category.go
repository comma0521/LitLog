package repo

import (
	"gorm.io/gorm"
	"time"
)

// 2.1

// GetCategoryByUserID 根据 user_id 查询该用户的所有书架
func GetCategoryByUserID(db *gorm.DB, userID int) ([]Category, error) {
	var categories []Category
	err := db.Where("user_id = ?", userID).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// 2.2

// CreateCategory 在数据库中新增一个分类
func CreateCategory(db *gorm.DB, userID int, categoryName string) (int, error) {
	category := Category{
		UserID:       userID,
		CategoryName: categoryName,
		CreateTime:   time.Now(),
	}
	if err := db.Create(&category).Error; err != nil {
		return 0, err
	}
	return category.CategoryID, nil
}

// 2.3

// DeleteCategory 删除指定的分类
func DeleteCategory(db *gorm.DB, userID int, categoryID int) error {
	return db.Where("user_id = ? AND category_id = ?", userID, categoryID).Delete(&Category{}).Error
}

// 2.4

// UpdateCategoryName 更新指定分类的名称
func UpdateCategoryName(db *gorm.DB, userID int, categoryID int, newName string) error {
	return db.Model(&Category{}).
		Where("user_id = ? AND category_id = ?", userID, categoryID).
		Update("category_name", newName).Error
}

// 2.5.1

// SearchCategories 根据分类名搜索分类
func SearchCategories(db *gorm.DB, userID int, keyword string) ([]Category, error) {
	var categories []Category
	// 模糊查询分类名称
	err := db.Where("user_id = ? AND category_name LIKE ?", userID, "%"+keyword+"%").Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}
