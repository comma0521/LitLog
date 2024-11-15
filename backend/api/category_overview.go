package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"litlog-backend/dto"
	"litlog-backend/repo"
	"net/http"
	"strconv"
	"time"
)

// 2.1

// GetCategoriesHandler 处理查询用户所有书架的请求
func GetCategoriesHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 解析请求参数
		var req dto.GetCategoriesRequest
		if err := c.ShouldBindQuery(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request parameters"})
			return
		}

		// 调用 repo 中的 GetCategoryByUserID 函数查询书架
		categories, err := repo.GetCategoryByUserID(db, req.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve categories"})
			return
		}

		// 构造响应数据
		var categoryResponses []dto.CategoryResponse
		for _, category := range categories {
			categoryResponses = append(categoryResponses, dto.CategoryResponse{
				CategoryID:   category.CategoryID,
				CategoryName: category.CategoryName,
				CreateTime:   category.CreateTime.Format(time.RFC3339),
			})
		}

		// 返回查询结果
		c.JSON(http.StatusOK, dto.GetCategoriesResponse{
			Categories: categoryResponses,
			Message:    "Categories retrieved successfully",
		})
	}
}

// 2.2

// CreateCategoryHandler 处理新增分类的请求
func CreateCategoryHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.CreateCategoryRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request parameters"})
			return
		}

		// 调用 CreateCategory 函数与数据库交互
		categoryID, err := repo.CreateCategory(db, req.UserID, req.CategoryName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
			return
		}

		// 返回成功响应
		c.JSON(http.StatusCreated, dto.CreateCategoryResponse{
			CategoryID: categoryID,
			Message:    "分类创建成功",
		})
	}
}

// 2.3

// DeleteCategoryHandler 处理删除分类的请求
func DeleteCategoryHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从路径参数中获取 category_id
		categoryIDParam := c.Param("category_id")
		categoryID, err := strconv.Atoi(categoryIDParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category_id"})
			return
		}

		// 从查询参数中获取 user_id
		userIDParam := c.Query("user_id")
		userID, err := strconv.Atoi(userIDParam)
		if err != nil || userIDParam == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing user_id"})
			return
		}

		// 使用事务删除分类及其下的所有书籍
		err = db.Transaction(func(tx *gorm.DB) error {
			//// 先删除分类下的所有书籍
			//if err := repo.DeleteBooksByCategory(tx, categoryID); err != nil {
			//	return err
			//}
			// 然后删除分类
			if err := repo.DeleteCategory(tx, userID, categoryID); err != nil {
				return err
			}
			return nil
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category and books"})
			return
		}

		// 返回删除成功的响应
		c.JSON(http.StatusOK, dto.DeleteCategoryResponse{
			Message: "分类删除成功",
		})
	}
}

// UpdateCategoryNameHandler 处理更新分类名称的请求
func UpdateCategoryNameHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从路径参数中获取 category_id
		categoryIDParam := c.Param("category_id")
		categoryID, err := strconv.Atoi(categoryIDParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category_id"})
			return
		}

		// 从查询参数中获取 user_id
		userIDParam := c.Query("user_id")
		userID, err := strconv.Atoi(userIDParam)
		if err != nil || userIDParam == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing user_id"})
			return
		}

		// 绑定请求体中的 JSON 数据到 UpdateCategoryNameRequest 结构体
		var req dto.UpdateCategoryNameRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		// 调用 repo.UpdateCategoryName 更新分类名称
		if err := repo.UpdateCategoryName(db, userID, categoryID, req.CategoryName); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category name"})
			return
		}

		// 返回成功响应
		c.JSON(http.StatusOK, dto.UpdateCategoryNameResponse{
			Message: "分类名称更新成功",
		})
	}
}

// 2.5.1

// SearchCategoryHandler 处理搜索分类的请求
func SearchCategoryHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取查询参数 user_id
		userIDParam := c.Query("user_id")
		userID, err := strconv.Atoi(userIDParam)
		if err != nil || userIDParam == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing user_id"})
			return
		}

		// 获取查询参数 keyword
		keyword := c.Query("keyword")
		if keyword == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Keyword is required"})
			return
		}

		// 调用 repo.SearchCategories 查询分类
		categories, err := repo.SearchCategories(db, userID, keyword)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search categories"})
			return
		}

		// 构造响应数据
		var results []dto.CategoryResult
		for _, category := range categories {
			results = append(results, dto.CategoryResult{
				ID:   category.CategoryID,
				Name: category.CategoryName,
			})
		}

		// 返回响应
		c.JSON(http.StatusOK, dto.SearchCategoryResponse{
			Results: results,
		})
	}
}

// 2.5.2

// SearchBookHandler 处理搜索书籍的请求
func SearchBookHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取查询参数 user_id
		userIDParam := c.Query("user_id")
		userID, err := strconv.Atoi(userIDParam)
		if err != nil || userIDParam == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing user_id"})
			return
		}

		// 获取查询参数 keyword
		keyword := c.Query("keyword")
		if keyword == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Keyword is required"})
			return
		}

		// 调用 repo.SearchBooks 查询书籍
		books, err := repo.SearchBooks(db, userID, keyword)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search books"})
			return
		}

		// 构造响应数据
		var results []dto.BookResult
		for _, book := range books {
			results = append(results, dto.BookResult{
				ID:   book.BookID,
				Name: book.BookName,
			})
		}

		// 返回响应
		c.JSON(http.StatusOK, dto.SearchBookResponse{
			Results: results,
		})
	}
}
