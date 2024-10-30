package entity

type Line string

const (
	Left  Line = "L"
	Right Line = "R"
)

type Node struct {
	ID          uint
	ParentId    uint
	Ancestry    string
	Line        string
	LftReferral string
	RgtReferral string
}
