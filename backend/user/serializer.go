package user

import (
	"backend/models"
	"backend/utils/response"
)

// User 用户序列化器
type User struct {
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
type UserResponse struct {
	Data User `json:"user"`
}

func UserDemo() models.User {
	user := models.User{
		Username:  "Bob",
		Nickname:  "Bob",
		Status:    "active",
		Avatar:    "https://pic4.zhimg.com/80/v2-867a95c44703177811f2590b09396113_1440w.jpg?source=1940ef5c",
		Role:      "admin",
		SuperUser: true,
	}
	user.ID = 1
	//user.CreatedOn = int(time.Now().Unix())
	return user

}

// BuildUser 序列化用户
func BuildUser(user models.User) User {
	return User{
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
	response := response.Response{}.SuccessResponse()
	response.Data = BuildUser(user)
	return response
}
