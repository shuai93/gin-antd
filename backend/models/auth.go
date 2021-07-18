package models

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *Auth) TableName() string {
	return "blog_auth"
}

func CheckAuth(username, password string) (bool, int) {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	if auth.ID > 0 {
		return true, auth.ID
	}

	return false, 0
}