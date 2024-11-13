package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建一个新的 Gin 路由器
	r := gin.Default()

	// 获取书架列表的处理函数
	r.GET("/api/shelf", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"categories": []gin.H{
				{
					"category_id":          1,
					"user_id":              1,
					"category_name":        "小说",
					"category_description": "包含各种小说类型",
					"create_time":          "2023-01-01T10:00:00Z",
				},
				{
					"category_id":          2,
					"user_id":              1,
					"category_name":        "技术书籍",
					"category_description": "与技术相关的书籍",
					"create_time":          "2023-02-01T10:00:00Z",
				},
			},
		})
	})

	// 新增分类的处理函数
	r.POST("/api/shelf/category", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"category_id": 3,
			"message":     "分类创建成功",
		})
	})

	// 删除分类的处理函数
	r.DELETE("/api/shelf/category/:category_id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "分类删除成功",
		})
	})

	// 修改分类名称的处理函数
	r.PUT("/api/shelf/category/:category_id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "分类名称修改成功",
		})
	})

	// 查询分类的处理函数
	r.GET("/api/shelf/search/category", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"results": []gin.H{
				{
					"category_id":   1,
					"category_name": "文学",
				},
			},
		})
	})

	// 查询书籍的处理函数
	r.GET("/api/shelf/search/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"results": []gin.H{
				{
					"book_id":   101,
					"book_name": "红楼梦",
				},
			},
		})
	})

	// 获取书籍详情的处理函数
	r.GET("/api/book/:book_id", func(c *gin.Context) {
		bookID := c.Param("book_id")
		c.JSON(http.StatusOK, gin.H{
			"book_id":          bookID,
			"book_name":        "红楼梦",
			"book_description": "一本中国古典名著。",
			"rating":           4,
			"create_time":      "2024-11-01",
			"excerpts": []gin.H{
				{
					"excerpt_id":      101,
					"excerpt_content": "满纸荒唐言，一把辛酸泪。",
					"thought":         "非常感慨。",
					"page":            12,
					"deleted_time":    nil,
				},
			},
			"reviews": []gin.H{
				{
					"review_id":      201,
					"review_content": "这本书让我重新思考人生的意义。",
				},
			},
		})
	})

	// 修改书名的处理函数
	r.PUT("/api/book/:book_id/name", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "书名更新成功",
		})
	})

	// 修改书描述的处理函数
	r.PUT("/api/book/:book_id/description", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "书籍描述更新成功",
		})
	})

	// 启动 Gin 路由器
	r.Run(":8080")
}
