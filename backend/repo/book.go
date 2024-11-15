package repo

import "gorm.io/gorm"

// DeleteBooksByCategory 删除指定分类下的所有书籍
func DeleteBooksByCategory(db *gorm.DB, categoryID int) error {
	return db.Where("category_id = ?", categoryID).Delete(&Book{}).Error
}

// SearchBooks 根据书名关键字搜索书籍
func SearchBooks(db *gorm.DB, userID int, keyword string) ([]Book, error) {
	var books []Book
	// 模糊查询书名
	err := db.Where("user_id = ? AND book_name LIKE ?", userID, "%"+keyword+"%").Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

// 4.1.1

// GetBookDetails 获取书籍的基本信息
func GetBookDetails(db *gorm.DB, userID int, bookID int) (*Book, error) {
	var book Book
	err := db.Where("user_id = ? AND book_id = ?", userID, bookID).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

// 4.2

// UpdateBookName 根据 book_id 更新书名
func UpdateBookName(db *gorm.DB, userID int, bookID int, newBookName string) error {
	return db.Model(&Book{}).
		Where("user_id = ? AND book_id = ?", userID, bookID).
		Update("book_name", newBookName).Error
}

// 4.3

// UpdateBookDescription 根据 book_id 更新书籍描述
func UpdateBookDescription(db *gorm.DB, userID int, bookID int, newDescription string) error {
	return db.Model(&Book{}).
		Where("user_id = ? AND book_id = ?", userID, bookID).
		Update("book_description", newDescription).Error
}
