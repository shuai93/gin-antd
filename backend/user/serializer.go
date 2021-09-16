package user

import (
	"backend/user/models"
	"backend/utils/response"
)

// User 用户序列化器
type Serializer struct {
	ID        int    `json:"id"`
	UserName  string `json:"name"`
	Nickname  string `json:"nick_name"`
	Status    string `json:"status"`
	Avatar    string `json:"avatar"`
	Admin     bool   `json:"admin"`
	Role      string `json:"role"`
	CreatedAt int    `json:"created_at"`
}

// UserResponse 单个用户序列化
type Response struct {
	Data models.User `json:"user"`
}

// BuildUser 序列化用户
func BuildUser(user models.User) Serializer {
	return Serializer{
		ID:       user.ID,
		UserName: user.Username,
		Nickname: user.Nickname,
		Status:   user.Status,
		Avatar:   user.Avatar,
		Admin:    user.SuperUser,
		Role:     user.Role,
		//CreatedAt: user.CreatedOn,
	}
}

// BuildUserResponse 序列化用户响应
func DefaultUserResponse(user models.User) response.Response {
	successResponse := response.Response{}.SuccessResponse()
	successResponse.Data = BuildUser(user)
	return successResponse
}
