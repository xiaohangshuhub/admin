package role

import (
	"github.com/google/uuid"
	"github.com/xiaohangshuhub/admin/internal/users/domain/dic/status"
	"github.com/xiaohangshuhub/go-workit/pkg/ddd"
)

type Role struct {
	ddd.AggregateRoot[uuid.UUID]               // ID
	Role                         string        // 角色
	Name                         string        // 名称
	ParentID                     uuid.UUID     // 父级ID
	Permissions                  []uuid.UUID   // 权限
	Status                       status.Status // 状态
}
