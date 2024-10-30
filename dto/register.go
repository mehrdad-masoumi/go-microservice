package dto

type RegisterRequest struct {
	Username        string `json:"username" example:"mehrdad"`
	Email           string `json:"email" example:"mehrdad@gmail.com"`
	PhoneNumber     string `json:"phone_number" example:"09120246217"`
	Password        string `json:"password" example:"123456@mmMM.$"`
	ConfirmPassword string `json:"confirm_password" example:"123456@mmMM.$"`
	Referral        string `json:"referral" example:"L100@1245"`
	WalletId        string `json:"wallet_id" example:"TZJjSqfXwTP5YgjdhLdRV1zqpDwDsgdZZE"`
}
