package menu

import (
	"github.com/google/uuid"
	"github.com/xiaohangshuhub/admin/internal/users/domain/dic/status"
	"github.com/xiaohangshuhub/go-workit/pkg/ddd"
)

type Rule struct {
	ddd.AggregateRoot[uuid.UUID]               // ID
	Rule                         string        // 规则
	Name                         string        // 名称
	Type                         RuleType      // 规则类型
	ParentID                     uuid.UUID     // 父级ID
	Route                        string        //路由
	Icon                         string        // 图标
	Desc                         string        // 描述
	Weight                       int32         // 权重
	Status                       status.Status // 状态
}
