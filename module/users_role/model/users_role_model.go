package model

// UserRole Struct
type UserRole struct {
	UserRoleID   int64  `json:"user_role_id"`
	UserRoleName string `json:"user_role_name"`
}

// UserRoleList list
type UserRoleList []UserRole

// NewUserRole func
func NewUserRole() *UserRole {
	return &UserRole{}
}
