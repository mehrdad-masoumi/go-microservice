package dto

type NodeInfo struct {
	LftReferralCode string `json:"lft_referral_code"`
	RgtReferralCode string `json:"rgt_referral_code"`
	Maintenance     uint   `json:"maintenance"`
	PackageId       uint   `json:"package_id"`
	Subset          uint   `json:"subset"`
}

type NodeCreateRequest struct {
	Username        string `json:"username" example:"mehrdad"`
	Email           string `json:"email" example:"mehrdad@gmail.com"`
	PhoneNumber     string `json:"phone_number" example:"09120246217"`
	Referral        string `json:"referral" example:"L100@2154"`
	WalletId        string `json:"wallet_id" example:"TZJjSqfXwTP5YgjdhLdRV1zqpDwDsgdZZE"`
	Password        string `json:"password" example:"123456"`
	ConfirmPassword string `json:"confirm_password" example:"123456"`
}

type NodeCreateResponse struct {
	NodeId uint `json:"id"`
}
