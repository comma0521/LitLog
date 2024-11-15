package dto

import "time"

// 4.1.1

// BookDetailsResponse 定义书籍详情的响应结构体
type BookDetailsResponse struct {
	BookID          int       `json:"id"`
	BookName        string    `json:"name"`
	BookAuthor      string    `json:"author"`
	BookRating      int       `json:"rating"`
	BookDescription string    `json:"description"`
	CreateTime      time.Time `json:"create_time"`
}

// 4.1.2

// BookExcerptResponse 定义书摘列表的响应结构体
type BookExcerptResponse struct {
	Excerpts []ExcerptDTO `json:"excerpts"`
}

// ExcerptDTO 定义单个书摘的信息
type ExcerptDTO struct {
	ExcerptID  int       `json:"id"`
	Content    string    `json:"content"`
	Thought    string    `json:"thought"`
	Page       int       `json:"page"`
	DeleteTime time.Time `json:"delete_time"`
	CreateTime time.Time `json:"create_time"`
}

// 4.2

// UpdateBookNameRequest 更新书名的请求结构体
type UpdateBookNameRequest struct {
	BookName string `json:"book_name" binding:"required"`
}

// UpdateBookNameResponse 更新书名的响应结构体
type UpdateBookNameResponse struct {
	Message string `json:"message"`
}

// 4.3

// UpdateBookDescriptionRequest 更新书籍描述的请求结构体
type UpdateBookDescriptionRequest struct {
	BookDescription string `json:"book_description" binding:"required"`
}

// UpdateBookDescriptionResponse 更新书籍描述的响应结构体
type UpdateBookDescriptionResponse struct {
	Message string `json:"message"`
}

// 4.4

// CreateBookExcerptRequest 新增书摘的请求结构体，仅包含 user_id
type CreateBookExcerptRequest struct {
	UserID int `json:"user_id" binding:"required"`
}

// CreateBookExcerptResponse 新增书摘的响应结构体
type CreateBookExcerptResponse struct {
	ExcerptID int    `json:"excerpt_id"`
	Message   string `json:"message"`
}

// 4.5

// DeleteBookExcerptResponse 删除书摘的响应结构体
type DeleteBookExcerptResponse struct {
	Message string `json:"message"`
}

// 4.6

// RestoreBookExcerptResponse 恢复书摘的响应结构体
type RestoreBookExcerptResponse struct {
	Message string `json:"message"`
}
