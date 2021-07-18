package serializer

import "backend/models"

// User 用户序列化器
type User struct {
	ID        int   `json:"id"`
	UserName  string `json:"user_name"`
	Nickname  string `json:"nickname"`
	Status    string `json:"status"`
	Avatar    string `json:"avatar"`
	Admin     bool   `json:"admin"`
	CreatedAt int  `json:"created_at"`
}

// UserResponse 单个用户序列化
type UserResponse struct {
	Data User `json:"user"`
}

// BuildUser 序列化用户
func BuildUser(user models.User) User {
	return User{
		ID:        user.ID,
		UserName:  user.UserName,
		Nickname:  user.Nickname,
		Status:    user.Status,
		Avatar:    user.Avatar,
		Admin:     user.SuperUser,
		CreatedAt: user.CreatedOn,
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user models.User) UserResponse {
	return UserResponse{
		Data: BuildUser(user),
	}
}
