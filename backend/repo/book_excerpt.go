package repo

import (
	"gorm.io/gorm"
	"time"
)

// 4.1.2

// GetBookExcerpts 根据 book_id 获取书摘列表（是否已经删除的都在）
func GetBookExcerpts(db *gorm.DB, bookID int) ([]BookExcerpt, error) {
	var excerpts []BookExcerpt
	err := db.Where("book_id = ?", bookID).Find(&excerpts).Error
	if err != nil {
		return nil, err
	}
	return excerpts, nil
}

// 4.4

// CreateBookExcerpt 创建新的书摘，仅初始化 book_id，其他字段为空
func CreateBookExcerpt(db *gorm.DB, bookID int) (int, error) {
	excerpt := BookExcerpt{
		BookID:     bookID,
		CreateTime: time.Now(),
	}

	if err := db.Create(&excerpt).Error; err != nil {
		return 0, err
	}

	return excerpt.ExcerptID, nil
}

// 4.5

// DeleteBookExcerpt 设置书摘的 DeletedTime 以表示删除
func DeleteBookExcerpt(db *gorm.DB, userID int, bookID int, excerptID int) error {
	return db.Model(&BookExcerpt{}).
		Where("book_id = ? AND excerpt_id = ?", bookID, excerptID).
		Update("deleted_time", time.Now()).Error
}

// 4.6

// RestoreBookExcerpt 将书摘的 DeletedTime 字段设置为 NULL，以表示恢复
func RestoreBookExcerpt(db *gorm.DB, userID int, bookID int, excerptID int) error {
	return db.Model(&BookExcerpt{}).
		Where("book_id = ? AND excerpt_id = ?", bookID, excerptID).
		Update("deleted_time", nil).Error
}
