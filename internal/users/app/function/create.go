package function

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/xiaohangshuhub/admin/internal/users/domain/function"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// CreateCmd 创建用户命令,包含创建用户所需的信息
type CreateCmd struct {
	Name     string        // 用户名
	Route    string        // 昵称
	Icon     string        // 图标
	Desc     string        // 头像
	Weight   int32         // 权重
	FType    function.Type // 功能类型
	ParentID *uuid.UUID    // 父节点ID
}

type CreateCmdHandler struct {
	*function.Manager
	*gorm.DB
	*zap.Logger
}

func NewCreateCmdHandler(manager *function.Manager, db *gorm.DB, zap *zap.Logger) *CreateCmdHandler {
	return &CreateCmdHandler{
		Manager: manager,
		DB:      db,
		Logger:  zap,
	}
}

func (c *CreateCmdHandler) Handle(ctx context.Context, cmd CreateCmd) (bool, error) {

	uid, ok := ctx.Value("UserID").(uuid.UUID)

	if !ok {
		return false, errors.New("invalid user id in context")
	}
	u, err := c.Manager.Create(cmd.Name, cmd.Route, cmd.Icon, cmd.Desc, cmd.Weight, cmd.FType, cmd.ParentID, uid)

	if err != nil {
		return false, err
	}

	if err := c.DB.Create(u).Error; err != nil {
		return false, err
	}

	return true, nil
}
