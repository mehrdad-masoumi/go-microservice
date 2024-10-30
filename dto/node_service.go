package dto

type NodeInfo struct {
	LftReferralCode string `json:"lft_referral_code"`
	RgtReferralCode string `json:"rgt_referral_code"`
	Maintenance     uint   `json:"maintenance"`
	PackageId       uint   `json:"package_id"`
	Subset          uint   `json:"subset"`
}

type NodeCreateRequest struct {
	UserID   uint   `json:"user_id"`
	Referral string `json:"referral" example:"L100@2154"`
	WalletId string `json:"wallet_id" example:"TZJjSqfXwTP5YgjdhLdRV1zqpDwDsgdZZE"`
}

type NodeCreateResponse struct {
	ID uint `json:"id"`
}
