package dto

// 2.1

// GetCategoriesRequest 定义获取分类的请求结构体
type GetCategoriesRequest struct {
	UserID int `form:"user_id" binding:"required"`
}

// CategoryResponse 定义单个分类的响应结构体
type CategoryResponse struct {
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
	CreateTime   string `json:"create_time"`
}

// GetCategoriesResponse 定义获取分类的整体响应结构体
type GetCategoriesResponse struct {
	Categories []CategoryResponse `json:"categories"`
	Message    string             `json:"message"`
}

// 2.2

type CreateCategoryRequest struct {
	UserID       int    `json:"user_id" binding:"required"`
	CategoryName string `json:"category_name" binding:"required"`
}

type CreateCategoryResponse struct {
	CategoryID int    `json:"category_id"`
	Message    string `json:"message"`
}

// 2.3

// DeleteCategoryResponse 删除分类的响应结构体
type DeleteCategoryResponse struct {
	Message string `json:"message"`
}

// 2.4

// UpdateCategoryNameRequest 更新分类名称的请求结构体
type UpdateCategoryNameRequest struct {
	CategoryName string `json:"category_name" binding:"required"`
}

// UpdateCategoryNameResponse 更新分类名称的响应结构体
type UpdateCategoryNameResponse struct {
	Message string `json:"message"`
}

// 2.5.1

// SearchCategoryResponse 定义搜索分类的响应结构体
type SearchCategoryResponse struct {
	Results []CategoryResult `json:"results"`
}

// CategoryResult 定义单个分类的搜索结果结构体
type CategoryResult struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// 2.5.2

// SearchBookResponse 定义搜索书籍的响应结构体
type SearchBookResponse struct {
	Results []BookResult `json:"results"`
}

// BookResult 定义单个书籍的搜索结果结构体
type BookResult struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
