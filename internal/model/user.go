package model

import "time"

// User 用户模型
type User struct {
	ID        int64     `json:"id" example:"1"`
	Username  string    `json:"username" example:"johndoe"`
	Email     string    `json:"email" example:"john@example.com"`
	Phone     string    `json:"phone" example:"13800138000"`
	Avatar    string    `json:"avatar" example:"https://example.com/avatar.png"`
	Status    int       `json:"status" example:"1" enums:"0,1,2"`
	CreatedAt time.Time `json:"created_at" example:"2024-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2024-01-01T00:00:00Z"`
}

// CreateUserRequest 创建用户请求
type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50" example:"johndoe"`
	Email    string `json:"email" binding:"required,email" example:"john@example.com"`
	Password string `json:"password" binding:"required,min=6" example:"password123"`
	Phone    string `json:"phone" example:"13800138000"`
}

// UpdateUserRequest 更新用户请求
type UpdateUserRequest struct {
	Username string `json:"username" binding:"omitempty,min=3,max=50" example:"johndoe"`
	Email    string `json:"email" binding:"omitempty,email" example:"john@example.com"`
	Phone    string `json:"phone" example:"13800138000"`
	Avatar   string `json:"avatar" example:"https://example.com/avatar.png"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required" example:"johndoe"`
	Password string `json:"password" binding:"required" example:"password123"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token     string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	ExpiresIn int64  `json:"expires_in" example:"3600"`
	User      User   `json:"user"`
}

// UserListResponse 用户列表响应
type UserListResponse struct {
	Total int64  `json:"total" example:"100"`
	Page  int    `json:"page" example:"1"`
	Size  int    `json:"size" example:"10"`
	Users []User `json:"users"`
}

// Response 通用响应
type Response struct {
	Code    int         `json:"code" example:"0"`
	Message string      `json:"message" example:"success"`
	Data    interface{} `json:"data"`
}

// ErrorResponse 错误响应
type ErrorResponse struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Bad Request"`
	Details string `json:"details,omitempty" example:"Invalid parameters"`
}

