package user

type User struct {
	ID           int64  `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"passwordHash"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	TG_ID        int64  `json:"telegram_id"`
}

type TgUser struct {
	TG_ID       int64  `json:"telegram_id"`
	TG_USERNAME string `json:"tg_username"`
}

type TgAdmin struct {
	TG_ID     int64  `json:"telegram_id"`
	ADMIN_PWD string `json:"admin_password"`
}
