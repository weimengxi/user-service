package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"user-service/internal/model"
)

// 模拟用户数据存储
var users = []model.User{
	{ID: 1, Username: "admin", Email: "admin@example.com", Phone: "13800000001", Status: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 2, Username: "user1", Email: "user1@example.com", Phone: "13800000002", Status: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 3, Username: "user2", Email: "user2@example.com", Phone: "13800000003", Status: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
}
var nextID int64 = 4

// HealthCheck 健康检查
// @Summary      健康检查
// @Description  检查服务是否正常运行
// @Tags         System
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /health [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "user-service",
		"time":    time.Now().Format(time.RFC3339),
	})
}

// GetUsers 获取用户列表
// @Summary      获取用户列表
// @Description  分页获取用户列表
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        page   query     int  false  "页码"  default(1)
// @Param        size   query     int  false  "每页数量"  default(10)
// @Success      200    {object}  model.Response{data=model.UserListResponse}
// @Failure      400    {object}  model.ErrorResponse
// @Security     BearerAuth
// @Router       /users [get]
func GetUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	// 简单分页
	start := (page - 1) * size
	end := start + size
	if start > len(users) {
		start = len(users)
	}
	if end > len(users) {
		end = len(users)
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "success",
		Data: model.UserListResponse{
			Total: int64(len(users)),
			Page:  page,
			Size:  size,
			Users: users[start:end],
		},
	})
}

// GetUser 获取用户详情
// @Summary      获取用户详情
// @Description  根据用户ID获取用户详细信息
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "用户ID"
// @Success      200  {object}  model.Response{data=model.User}
// @Failure      404  {object}  model.ErrorResponse
// @Security     BearerAuth
// @Router       /users/{id} [get]
func GetUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    400,
			Message: "Invalid user ID",
		})
		return
	}

	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, model.Response{
				Code:    0,
				Message: "success",
				Data:    user,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, model.ErrorResponse{
		Code:    404,
		Message: "User not found",
	})
}

// CreateUser 创建用户
// @Summary      创建用户
// @Description  创建新用户账户
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user  body      model.CreateUserRequest  true  "用户信息"
// @Success      201   {object}  model.Response{data=model.User}
// @Failure      400   {object}  model.ErrorResponse
// @Security     BearerAuth
// @Router       /users [post]
func CreateUser(c *gin.Context) {
	var req model.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    400,
			Message: "Invalid request body",
			Details: err.Error(),
		})
		return
	}

	// 检查用户名是否已存在
	for _, user := range users {
		if user.Username == req.Username {
			c.JSON(http.StatusBadRequest, model.ErrorResponse{
				Code:    400,
				Message: "Username already exists",
			})
			return
		}
	}

	newUser := model.User{
		ID:        nextID,
		Username:  req.Username,
		Email:     req.Email,
		Phone:     req.Phone,
		Status:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	nextID++
	users = append(users, newUser)

	c.JSON(http.StatusCreated, model.Response{
		Code:    0,
		Message: "User created successfully",
		Data:    newUser,
	})
}

// UpdateUser 更新用户
// @Summary      更新用户信息
// @Description  根据用户ID更新用户信息
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id    path      int                      true  "用户ID"
// @Param        user  body      model.UpdateUserRequest  true  "用户信息"
// @Success      200   {object}  model.Response{data=model.User}
// @Failure      400   {object}  model.ErrorResponse
// @Failure      404   {object}  model.ErrorResponse
// @Security     BearerAuth
// @Router       /users/{id} [put]
func UpdateUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    400,
			Message: "Invalid user ID",
		})
		return
	}

	var req model.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    400,
			Message: "Invalid request body",
			Details: err.Error(),
		})
		return
	}

	for i, user := range users {
		if user.ID == id {
			if req.Username != "" {
				users[i].Username = req.Username
			}
			if req.Email != "" {
				users[i].Email = req.Email
			}
			if req.Phone != "" {
				users[i].Phone = req.Phone
			}
			if req.Avatar != "" {
				users[i].Avatar = req.Avatar
			}
			users[i].UpdatedAt = time.Now()

			c.JSON(http.StatusOK, model.Response{
				Code:    0,
				Message: "User updated successfully",
				Data:    users[i],
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, model.ErrorResponse{
		Code:    404,
		Message: "User not found",
	})
}

// DeleteUser 删除用户
// @Summary      删除用户
// @Description  根据用户ID删除用户
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "用户ID"
// @Success      200  {object}  model.Response
// @Failure      404  {object}  model.ErrorResponse
// @Security     BearerAuth
// @Router       /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    400,
			Message: "Invalid user ID",
		})
		return
	}

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, model.Response{
				Code:    0,
				Message: "User deleted successfully",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, model.ErrorResponse{
		Code:    404,
		Message: "User not found",
	})
}

// Login 用户登录
// @Summary      用户登录
// @Description  使用用户名和密码登录
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        credentials  body      model.LoginRequest  true  "登录凭证"
// @Success      200          {object}  model.Response{data=model.LoginResponse}
// @Failure      401          {object}  model.ErrorResponse
// @Router       /auth/login [post]
func Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    400,
			Message: "Invalid request body",
			Details: err.Error(),
		})
		return
	}

	// 简单的登录验证（实际应用中应该验证密码哈希）
	for _, user := range users {
		if user.Username == req.Username {
			c.JSON(http.StatusOK, model.Response{
				Code:    0,
				Message: "Login successful",
				Data: model.LoginResponse{
					Token:     "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MDAwMDAwMDB9.mock-token",
					ExpiresIn: 3600,
					User:      user,
				},
			})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, model.ErrorResponse{
		Code:    401,
		Message: "Invalid username or password",
	})
}

