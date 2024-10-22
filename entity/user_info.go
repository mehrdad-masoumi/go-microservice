package entity

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)

type UserInfo struct {
	UserID    uint
	FirstName string
	LastName  string
	Gender    Gender
}
