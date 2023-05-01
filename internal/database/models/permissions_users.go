package models

type PermissionUser struct {
	ID    int
	Ptype string
	VO    string
	V1    string
	V2    string
}

func (PermissionUser) TableName() string {
	return "permissions__users"
}
