package perm

import "github.com/google/uuid"

// RoleFunc 角色功能关系
type RoleFunc struct {
	RoleId uuid.UUID
	FuncId uuid.UUID
}
