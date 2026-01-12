package function

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/xiaohangshuhub/admin/internal/users/domain/function"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// UpdateCmd 更新用户命令,包含更新用户所需的信息
type UpdateCmd struct {
	ID       uuid.UUID     `json:"id"`
	Name     string        // 用户名
	Route    string        // 昵称
	Icon     string        // 图标
	Desc     string        // 头像
	Weight   int32         // 权重
	FType    function.Type // 功能类型
	ParentID *uuid.UUID    // 父节点ID
}

type UpdateCmdHandler struct {
	*function.Manager
	*gorm.DB
	*zap.Logger
}

func NewUpdateCmdHandler(m *function.Manager, repo *gorm.DB, zap *zap.Logger) *UpdateCmdHandler {
	return &UpdateCmdHandler{
		Manager: m,
		DB:      repo,
		Logger:  zap,
	}
}

func (h *UpdateCmdHandler) Handle(ctx context.Context, cmd UpdateCmd) (bool, error) {

	uid, ok := ctx.Value("UserID").(uuid.UUID)

	if !ok {
		return false, errors.New("invalid user id in context")
	}

	account, err := h.Manager.Update(cmd.ID, cmd.Name, cmd.Route, cmd.Icon, cmd.Desc, cmd.Weight, cmd.FType, cmd.ParentID, uid)

	if err != nil {
		return false, err
	}

	tx := h.DB.Save(account)

	if tx.Error != nil {
		// TODO: 后续优化
		return false, tx.Error
	}

	return true, nil
}
