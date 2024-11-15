package repo

import "time"

// User 1. 用户表 (user)
type User struct {
	UserID       int    `gorm:"primaryKey;column:user_id"`
	UserName     string `gorm:"column:user_name"`
	UserPassword string `gorm:"column:user_password"`
}

// TableName 自定义表名为 "user"
func (User) TableName() string {
	return "user"
}

// Category 2. 书架表 (category)
type Category struct {
	CategoryID   int       `gorm:"primaryKey;column:category_id"`
	UserID       int       `gorm:"column:user_id"`
	CategoryName string    `gorm:"column:category_name"`
	CreateTime   time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
}

// TableName 自定义表名为 "category"
func (Category) TableName() string {
	return "category"
}

// Book 3. 图书表 (book)
type Book struct {
	BookID          int        `gorm:"primaryKey;column:book_id"`
	UserID          int        `gorm:"column:user_id"`
	CategoryID      int        `gorm:"column:category_id"`
	BookName        string     `gorm:"column:book_name"`
	BookAuthor      string     `gorm:"column:book_author"`
	BookRating      int        `gorm:"column:book_rating;default:0"`
	BookDescription string     `gorm:"column:book_description"`
	DeletedTime     *time.Time `gorm:"column:deleted_time"`
	AccessTime      *time.Time `gorm:"column:access_time"`
	CreateTime      time.Time  `gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
}

// TableName 自定义表名为 "book"
func (Book) TableName() string {
	return "book"
}

// BookExcerpt 4. 书摘表 (book_excerpt)
type BookExcerpt struct {
	ExcerptID      int       `gorm:"primaryKey;column:excerpt_id"`
	BookID         int       `gorm:"column:book_id"`
	ExcerptContent string    `gorm:"column:excerpt_content"`
	ExcerptThought string    `gorm:"column:excerpt_thought"`
	ExcerptPage    int       `gorm:"column:excerpt_page"`
	DeletedTime    time.Time `gorm:"column:deleted_time"`
	CreateTime     time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
}

// TableName 自定义表名为 "book_excerpt"
func (BookExcerpt) TableName() string {
	return "book_excerpt"
}

// BookReview 5. 书评表 (book_review)
type BookReview struct {
	ReviewID      int        `gorm:"primaryKey;column:review_id"`
	BookID        int        `gorm:"column:book_id"`
	ReviewThought string     `gorm:"column:review_thought"`
	DeletedTime   *time.Time `gorm:"column:deleted_time"`
	CreateTime    time.Time  `gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
}

// TableName 自定义表名为 "book_review"
func (BookReview) TableName() string {
	return "book_review"
}
