package perm

import (
	"github.com/google/uuid"
)

// UserRole 用户角色关系
type UserRole struct {
	UserID uuid.UUID
	RoleID uuid.UUID
}
