// router.go

package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"litlog-backend/api"
)

// InitializeRouter 配置所有路由
func InitializeRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// 2

	// 2.1 获取书架列表
	router.GET("/api/category", api.GetCategoriesHandler(db))
	// 2.2 新增分类
	router.POST("/api/category", api.CreateCategoryHandler(db))
	// 2.3 删除分类及内部书
	router.DELETE("/api/category/:category_id", api.DeleteCategoryHandler(db))
	// 2.4 更新分类名称的路由
	router.PUT("/api/category/:category_id", api.UpdateCategoryNameHandler(db))
	// 2.5.1 搜索分类的路由
	router.GET("/api/search/category", api.SearchCategoryHandler(db))
	// 2.5.2 搜索书籍的路由
	router.GET("/api/search/book", api.SearchBookHandler(db))

	// 4

	// 4.1.1 获取书籍详情的路由
	router.GET("/api/book/:book_id", api.GetBookDetailsHandler(db))
	// 4.1.2 获取书摘列表的路由
	router.GET("/api/book/:book_id/excerpts", api.GetBookExcerptsHandler(db))
	// 4.2 更新书名的路由
	router.PUT("/api/book/:book_id/name", api.UpdateBookNameHandler(db))
	// 4.3 更新书籍描述的路由
	router.PUT("/api/book/:book_id/description", api.UpdateBookDescriptionHandler(db))
	// 4.4 新增书摘的路由
	router.POST("/api/book/:book_id/excerpt", api.CreateBookExcerptHandler(db))
	// 4.5 删除书摘的路由
	router.DELETE("/api/book/:book_id/excerpt/:excerpt_id", api.DeleteBookExcerptHandler(db))
	// 4.6 恢复书摘的路由
	router.PUT("/api/book/:book_id/excerpt/:excerpt_id/restore", api.RestoreBookExcerptHandler(db))

	return router
}
