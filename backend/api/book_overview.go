package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"litlog-backend/dto"
	"litlog-backend/repo"
	"net/http"
	"strconv"
)

// 4.1.1

// GetBookDetailsHandler 处理获取书籍详情的请求
func GetBookDetailsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从路径参数中获取 book_id
		bookIDParam := c.Param("book_id")
		bookID, err := strconv.Atoi(bookIDParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book_id"})
			return
		}

		// 从查询参数中获取 user_id
		userIDParam := c.Query("user_id")
		userID, err := strconv.Atoi(userIDParam)
		if err != nil || userIDParam == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing user_id"})
			return
		}

		// 获取书籍详情
		book, err := repo.GetBookDetails(db, userID, bookID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve book details"})
			return
		}

		// 构造响应数据
		response := dto.BookDetailsResponse{
			BookID:          book.BookID,
			BookName:        book.BookName,
			BookAuthor:      book.BookAuthor,
			BookRating:      book.BookRating,
			BookDescription: book.BookDescription,
			CreateTime:      book.CreateTime,
		}

		// 返回响应
		c.JSON(http.StatusOK, response)
	}
}

// 4.1.2

// GetBookExcerptsHandler 处理获取书摘列表的请求
func GetBookExcerptsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从路径参数中获取 book_id
		bookIDParam := c.Param("book_id")
		bookID, err := strconv.Atoi(bookIDParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book_id"})
			return
		}

		// 获取书摘列表
		excerpts, err := repo.GetBookExcerpts(db, bookID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve book excerpts"})
			return
		}

		// 构造响应数据
		var excerptDTOs []dto.ExcerptDTO
		for _, excerpt := range excerpts {
			excerptDTOs = append(excerptDTOs, dto.ExcerptDTO{
				ExcerptID:  excerpt.ExcerptID,
				Content:    excerpt.ExcerptContent,
				Thought:    excerpt.ExcerptThought,
				Page:       excerpt.ExcerptPage,
				DeleteTime: excerpt.DeletedTime,
				CreateTime: excerpt.CreateTime,
			})
		}

		response := dto.BookExcerptResponse{
			Excerpts: excerptDTOs,
		}

		// 返回响应
		c.JSON(http.StatusOK, response)
	}
}

// 4.2

// UpdateBookNameHandler 处理更新书名的请求
func UpdateBookNameHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从路径参数中获取 book_id
		bookIDParam := c.Param("book_id")
		bookID, err := strconv.Atoi(bookIDParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book_id"})
			return
		}

		// 从查询参数中获取 user_id
		userIDParam := c.Query("user_id")
		userID, err := strconv.Atoi(userIDParam)
		if err != nil || userIDParam == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing user_id"})
			return
		}

		// 绑定请求体中的 JSON 数据到 UpdateBookNameRequest 结构体
		var req dto.UpdateBookNameRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		// 调用 repo.UpdateBookName 更新书名
		if err := repo.UpdateBookName(db, userID, bookID, req.BookName); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book name"})
			return
		}

		// 返回成功响应
		c.JSON(http.StatusOK, dto.UpdateBookNameResponse{
			Message: "书名更新成功",
		})
	}
}

// 4.3

// UpdateBookDescriptionHandler 处理更新书籍描述的请求
func UpdateBookDescriptionHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从路径参数中获取 book_id
		bookIDParam := c.Param("book_id")
		bookID, err := strconv.Atoi(bookIDParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book_id"})
			return
		}

		// 从查询参数中获取 user_id
		userIDParam := c.Query("user_id")
		userID, err := strconv.Atoi(userIDParam)
		if err != nil || userIDParam == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing user_id"})
			return
		}

		// 绑定请求体中的 JSON 数据到 UpdateBookDescriptionRequest 结构体
		var req dto.UpdateBookDescriptionRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		// 调用 repo.UpdateBookDescription 更新书籍描述
		if err := repo.UpdateBookDescription(db, userID, bookID, req.BookDescription); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book description"})
			return
		}

		// 返回成功响应
		c.JSON(http.StatusOK, dto.UpdateBookDescriptionResponse{
			Message: "书籍描述更新成功",
		})
	}
}

// 4.4

// CreateBookExcerptHandler 处理新增书摘的请求
func CreateBookExcerptHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从路径参数中获取 book_id
		bookIDParam := c.Param("book_id")
		bookID, err := strconv.Atoi(bookIDParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book_id"})
			return
		}

		// 从请求体中绑定 JSON 数据到 CreateBookExcerptRequest 结构体，仅需要 user_id
		var req dto.CreateBookExcerptRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		// 调用 repo.CreateBookExcerpt 创建新的书摘
		excerptID, err := repo.CreateBookExcerpt(db, bookID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book excerpt"})
			return
		}

		// 返回成功响应
		c.JSON(http.StatusCreated, dto.CreateBookExcerptResponse{
			ExcerptID: excerptID,
			Message:   "书摘新增成功",
		})
	}
}

// 4.5

// DeleteBookExcerptHandler 处理删除书摘的请求
func DeleteBookExcerptHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从路径参数中获取 book_id 和 excerpt_id
		bookIDParam := c.Param("book_id")
		bookID, err := strconv.Atoi(bookIDParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book_id"})
			return
		}

		excerptIDParam := c.Param("excerpt_id")
		excerptID, err := strconv.Atoi(excerptIDParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid excerpt_id"})
			return
		}

		// 从查询参数中获取 user_id
		userIDParam := c.Query("user_id")
		userID, err := strconv.Atoi(userIDParam)
		if err != nil || userIDParam == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing user_id"})
			return
		}

		// 调用 repo.DeleteBookExcerpt 设置书摘的删除时间
		if err := repo.DeleteBookExcerpt(db, userID, bookID, excerptID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book excerpt"})
			return
		}

		// 返回成功响应
		c.JSON(http.StatusOK, dto.DeleteBookExcerptResponse{
			Message: "书摘删除成功",
		})
	}
}

// 4.6

// RestoreBookExcerptHandler 处理恢复书摘的请求
func RestoreBookExcerptHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从路径参数中获取 book_id 和 excerpt_id
		bookIDParam := c.Param("book_id")
		bookID, err := strconv.Atoi(bookIDParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book_id"})
			return
		}

		excerptIDParam := c.Param("excerpt_id")
		excerptID, err := strconv.Atoi(excerptIDParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid excerpt_id"})
			return
		}

		// 从查询参数中获取 user_id
		userIDParam := c.Query("user_id")
		userID, err := strconv.Atoi(userIDParam)
		if err != nil || userIDParam == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing user_id"})
			return
		}

		// 调用 repo.RestoreBookExcerpt 将 DeletedTime 设置为 NULL
		if err := repo.RestoreBookExcerpt(db, userID, bookID, excerptID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to restore book excerpt"})
			return
		}

		// 返回成功响应
		c.JSON(http.StatusOK, dto.RestoreBookExcerptResponse{
			Message: "书摘恢复成功",
		})
	}
}
