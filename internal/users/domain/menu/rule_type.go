package menu

type RuleType int8

const (
	Menu     RuleType = iota + 1 // 菜单
	MenuItem                     // 菜单项
	Action                       // 操作
)
