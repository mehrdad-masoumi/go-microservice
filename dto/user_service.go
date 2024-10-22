package dto

type UserInfo struct {
	Id          uint   `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type UserCreateRequest struct {
	Username        string `json:"username" example:"mehrdad"`
	Email           string `json:"email" example:"mehrdad@gmail.com"`
	PhoneNumber     string `json:"phone_number" example:"09120246217"`
	Password        string `json:"password" example:"123456"`
	ConfirmPassword string `json:"confirm_password" example:"123456"`
}

type UserCreateResponse struct {
	ID uint `json:"id"`
}
